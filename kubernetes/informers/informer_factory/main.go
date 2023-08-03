package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/discovery"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/metadata/metadatainformer"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/controller-manager/pkg/informerfactory"
	"k8s.io/klog/v2"
)

const (
	addEvent eventType = iota
	updateEvent
	deleteEvent
)

var (
	resources map[schema.GroupVersionResource]struct{}
)

type eventType int

func (e eventType) String() string {
	switch e {
	case addEvent:
		return "add"
	case updateEvent:
		return "update"
	case deleteEvent:
		return "delete"
	default:
		return fmt.Sprintf("unknown(%d)", int(e))
	}
}

type event struct {
	// virtual indicates this event did not come from an informer, but was constructed artificially
	virtual   bool
	eventType eventType
	obj       interface{}
	// the update event comes with an old object, but it's not used by the garbage collector.
	oldObj interface{}
	gvk    schema.GroupVersionKind
}

type GraphBuilder struct {
	sharedInformers informerfactory.InformerFactory
	restMapper      meta.RESTMapper
	graphChanges    workqueue.RateLimitingInterface
	monitors        monitors
	stopCh          <-chan struct{}
}

func NewGraphBuilder(informer informerfactory.InformerFactory, restMapper meta.RESTMapper) *GraphBuilder {
	return &GraphBuilder{
		sharedInformers: informer,
		restMapper:      restMapper,
		graphChanges:    workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "garbage_collector_graph_changes"),
	}
}

type monitors map[schema.GroupVersionResource]*monitor

// monitor runs a Controller with a local stop channel.
type monitor struct {
	controller cache.Controller
	store      cache.Store

	// stopCh stops Controller. If stopCh is nil, the monitor is considered to be not yet started.
	// stopCh 停止控制器。如果stopCh为nil，则认为监视器尚未启动。
	stopCh chan struct{}
}

func (m *monitor) Run() {
	m.controller.Run(m.stopCh)
}

func (gb *GraphBuilder) Run() {
	gb.syncMonitor()
	gb.startMonitor()
}

func (gb *GraphBuilder) syncMonitor() {
	current := monitors{}
	for resource := range resources {

		// 获取资源对象的kind
		kind, err := gb.restMapper.KindFor(resource)
		if err != nil {
			panic(err)
		}
		c, s, err := gb.controllerFor(resource, kind)
		if err != nil {
			panic(err)
		}
		current[resource] = &monitor{store: s, controller: c}
	}
	gb.monitors = current
}

func (gb *GraphBuilder) runProcessGraphChanges() {
	for gb.processGraphChanges() {
	}
}

func (gb *GraphBuilder) processGraphChanges() bool {
	// 从graphChanges队列中获取事件
	item, quit := gb.graphChanges.Get()
	if quit {
		return false
	}
	// 事件处理完成后，调用Done方法，将事件从graphChanges队列中移除
	defer gb.graphChanges.Done(item)

	event, ok := item.(*event)
	if !ok {
		utilruntime.HandleError(fmt.Errorf("expect a *event, got %v", item))
		return true
	}
	obj := event.obj
	// 获取对象的accessor
	accessor, err := meta.Accessor(obj)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("cannot access obj: %v", err))
		return true
	}
	klog.Infof("processing %s event for %s %s/%s", event.eventType, event.gvk, accessor.GetNamespace(), accessor.GetName())

	switch {
	case event.eventType == addEvent:
	case event.eventType == updateEvent:
	case event.eventType == deleteEvent:
	}
	return true
}

func (gb *GraphBuilder) controllerFor(resource schema.GroupVersionResource, kind schema.GroupVersionKind) (cache.Controller, cache.Store, error) {

	// 1. 将 新增、更改、删除的资源对象构建为event结构体，放入GraphBuilder的graphChanges队列里
	handlers := cache.ResourceEventHandlerFuncs{
		// add the event to the dependencyGraphBuilder's graphChanges.
		// 添加事件到dependencyGraphBuilder的graphChanges。
		AddFunc: func(obj interface{}) {
			event := &event{
				eventType: addEvent,
				obj:       obj,
				gvk:       kind,
			}
			gb.graphChanges.Add(event)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			// TODO: check if there are differences in the ownerRefs,
			// finalizers, and DeletionTimestamp; if not, ignore the update.
			event := &event{
				eventType: updateEvent,
				obj:       newObj,
				oldObj:    oldObj,
				gvk:       kind,
			}
			gb.graphChanges.Add(event)
		},
		DeleteFunc: func(obj interface{}) { // 执行 delete 操作会被监听到
			// delta fifo may wrap the object in a cache.DeletedFinalStateUnknown, unwrap it
			// delta fifo可以将对象包装在cache.DeletedFinalStateUnknown中，解包它
			if deletedFinalStateUnknown, ok := obj.(cache.DeletedFinalStateUnknown); ok {
				obj = deletedFinalStateUnknown.Obj
			}
			event := &event{
				eventType: deleteEvent,
				obj:       obj,
				gvk:       kind,
			}
			gb.graphChanges.Add(event)
		},
	}
	// 2. 从sharedInformers中获取资源对象的informer，添加事件处理函数
	shared, err := gb.sharedInformers.ForResource(resource)
	if err != nil {
		// 获取资源对象时出错会到这里,比如非k8s内置RedisCluster、clusterbases、clusters、esclusters、volumeproviders、stsmasters、appapps、mysqlclusters、brokerclusters、clustertemplates;
		// 内置的networkPolicies、apiservices、customresourcedefinitions
		klog.Infof("unable to use a shared informer for resource %q, kind %q: %v", resource.String(), kind.String(), err)
		return nil, nil, err
	}

	klog.V(4).Infof("using a shared informer for resource %q, kind %q", resource.String(), kind.String())

	// need to clone because it's from a shared cache
	// 不需要克隆，因为它不是来自共享缓存
	shared.Informer().AddEventHandlerWithResyncPeriod(handlers, 0) // 0

	return shared.Informer().GetController(), shared.Informer().GetStore(), nil
}

func (gb *GraphBuilder) startMonitor() error {
	started := 0
	var gvrs []string
	for gvr, monitor := range gb.monitors {
		if monitor.stopCh == nil {
			monitor.stopCh = make(chan struct{})
			gb.sharedInformers.Start(gb.stopCh)
			klog.Infof("-----> start monitor gvr: %+v,", gvr.String())
			gvrs = append(gvrs, gvr.String())
			go monitor.Run()
			started++
		}
	}
	gb.monitors = nil
	return nil
}

func init() {
	resources = make(map[schema.GroupVersionResource]struct{})
	resources[schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}] = struct{}{}
	resources[schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}] = struct{}{}
}

func main() {
	// client
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	metadataClient, err := metadata.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	// 实力化 share informer factory
	// ResyncPeriod: 默认的同步周期
	sharedInformers := informers.NewSharedInformerFactory(clientSet, 0)
	metadataInformers := metadatainformer.NewSharedInformerFactory(metadataClient, 0)
	sharedInformerFactory := informerfactory.NewInformerFactory(sharedInformers, metadataInformers)

	// Use a discovery client capable of being refreshed.
	cachedClient := cacheddiscovery.NewMemCacheClient(discoveryClient)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedClient)

	gb := NewGraphBuilder(sharedInformerFactory, restMapper)
	gb.Run()

	// 开启消费者
	stopCh := make(chan struct{})
	wait.Until(gb.runProcessGraphChanges, 1*time.Second, stopCh)

	<-stopCh
}
