package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

const schedulerName = "random-scheduler"

type predicateFunc func(node *corev1.Node, pod *corev1.Pod) bool
type priorityFunc func(node *corev1.Node, pod *corev1.Pod) int

type Scheduler struct {
	clientset  *kubernetes.Clientset
	podQueue   chan *corev1.Pod
	nodeLister listersv1.NodeLister
	predicates []predicateFunc
	priorities []priorityFunc
}

func NewScheduler(podQueue chan *corev1.Pod, quit chan struct{}) *Scheduler {
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatal(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	// 创建一个 Scheduler 对象
	scheduler := &Scheduler{
		clientset:  clientSet,
		podQueue:   podQueue,
		nodeLister: initInformers(clientSet, podQueue, quit),
	}
	// 注册调度器
	scheduler.registerPredicates()
	scheduler.registerPriorities()
	return scheduler
}

func initInformers(clientSet *kubernetes.Clientset, podQueue chan *corev1.Pod, quit chan struct{}) listersv1.NodeLister {
	factory := informers.NewSharedInformerFactory(clientSet, 0)

	nodeInformer := factory.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			node, ok := obj.(*corev1.Node)
			if !ok {
				klog.Warningf("this is not a node")
				return
			}
			klog.Infof("New Node Added to Store: %s", node.GetName())
		},
	})

	podInformer := factory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				klog.Warningf("this is not a pod")
				return
			}
			if pod.Spec.NodeName == "" && pod.Spec.SchedulerName == schedulerName {
				podQueue <- pod
			}
		},
	})

	factory.Start(quit)
	return nodeInformer.Lister()
}

func (s *Scheduler) Run() {
	ctx := context.Background()

	// 启动一个 goroutine，用于处理 podQueue 中的 pod
	s.schedule(ctx)
}

func (s *Scheduler) schedule(ctx context.Context) {
	for {
		// 从 podQueue 中获取一个 pod
		pod := <-s.podQueue
		// 执行调度
		s.scheduleOne(ctx, pod)
	}
}

func (s *Scheduler) scheduleOne(ctx context.Context, pod *corev1.Pod) {
	// 获取所有的节点
	nodes, err := s.nodeLister.List(labels.Everything())
	if err != nil {
		// 获取失败，将 pod 重新放入队列
		s.podQueue <- pod
		return
	}
	// 执行预选
	filteredNodes := s.runPredicates(pod, nodes)

	// 执行优选
	priorityNodes := s.runPriorities(pod, filteredNodes)

	// 选择一个节点
	node := s.selectNode(pod, priorityNodes)

	// 绑定 pod 和 node
	err = s.bindPod(ctx, pod, node)
	if err != nil {
		klog.Errorf("bind pod %s to node %s failed: %v", pod.GetName(), node.GetName(), err)
	}

	// 启动一个 goroutine，用于处理事件
	message := fmt.Sprintf("Placed pod [%s/%s] on %s\n", pod.Namespace, pod.Name, node)
	go s.watchEvent(ctx, pod, message)
}

func (s *Scheduler) runPredicates(pod *corev1.Pod, nodes []*corev1.Node) []*corev1.Node {
	filteredNodes := nodes
	for _, predicate := range s.predicates {
		filteredNodes = s.runPredicate(predicate, pod, filteredNodes)
	}
	return filteredNodes
}

func (s *Scheduler) runPredicate(predicate predicateFunc, pod *corev1.Pod, nodes []*corev1.Node) []*corev1.Node {
	filteredNodes := make([]*corev1.Node, 0, len(nodes))
	for _, node := range nodes {
		if predicate(node, pod) {
			filteredNodes = append(filteredNodes, node)
		}
	}
	return filteredNodes
}

func (s *Scheduler) runPriorities(pod *corev1.Pod, nodes []*corev1.Node) map[*corev1.Node]int {
	var priorities map[*corev1.Node]int
	for _, priority := range s.priorities {
		for node, score := range s.runPriority(priority, pod, nodes) {
			if v, ok := priorities[node]; ok {
				priorities[node] = v + score
			}
			priorities[node] += score
		}
	}
	return priorities
}

func (s *Scheduler) runPriority(priority priorityFunc, pod *corev1.Pod, nodes []*corev1.Node) map[*corev1.Node]int {
	priorities := make(map[*corev1.Node]int)
	for _, node := range nodes {
		priorities[node] = priority(node, pod)
	}
	return priorities
}

func (s *Scheduler) selectNode(pod *corev1.Pod, priorities map[*corev1.Node]int) *corev1.Node {
	var maxP int
	var bestNode *corev1.Node
	for node, p := range priorities {
		if p > maxP {
			maxP = p
			bestNode = node
		}
	}
	return bestNode
}

func (s *Scheduler) bindPod(ctx context.Context, pod *corev1.Pod, node *corev1.Node) error {
	// 绑定 pod 和 node 的逻辑
	return s.clientset.CoreV1().Pods(pod.Namespace).Bind(ctx, &corev1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
		Target: corev1.ObjectReference{
			APIVersion: "v1",
			Kind:       "Node",
			Name:       node.Name,
		},
	}, metav1.CreateOptions{})

}

func (s *Scheduler) watchEvent(ctx context.Context, p *corev1.Pod, message string) error {
	// 监听事件的逻辑
	timestamp := time.Now().UTC()
	e := &corev1.Event{
		Count:          1,
		Message:        message,
		Reason:         "Scheduled",
		LastTimestamp:  metav1.NewTime(timestamp),
		FirstTimestamp: metav1.NewTime(timestamp),
		Type:           "Normal",
		Source: corev1.EventSource{
			Component: schedulerName,
		},
		InvolvedObject: corev1.ObjectReference{
			Kind:      "Pod",
			Name:      p.Name,
			Namespace: p.Namespace,
			UID:       p.UID,
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: p.Name + "-",
		},
	}
	_, err := s.clientset.CoreV1().Events(p.Namespace).Create(ctx, e, metav1.CreateOptions{})
	return err
}

func (s *Scheduler) watchNodes() {
	// 监听节点的逻辑
}

func randomPredicate(node *corev1.Node, pod *corev1.Pod) bool {
	r := rand.Intn(2)
	return r == 0
}

func randomPriority(node *corev1.Node, pod *corev1.Pod) int {
	return rand.Intn(100)
}

func (s *Scheduler) registerPredicates() {
	// 注册预选的逻辑
	s.registerPredicate(randomPredicate)
}

func (s *Scheduler) registerPriorities() {
	// 注册优选的逻辑
	s.registerPriority(randomPriority)
}

func (s *Scheduler) registerPredicate(predicate predicateFunc) {
	s.predicates = append(s.predicates, predicate)
}

func (s *Scheduler) registerPriority(priority priorityFunc) {
	s.priorities = append(s.priorities, priority)
}

func (s *Scheduler) registerPredicateFunc(predicate predicateFunc) {
	s.registerPredicate(func(node *corev1.Node, pod *corev1.Pod) bool {
		return predicate(node, pod)
	})
}

func (s *Scheduler) registerPriorityFunc(priority priorityFunc) {
	s.registerPriority(func(node *corev1.Node, pod *corev1.Pod) int {
		return priority(node, pod)
	})
}

func main() {
	klog.Info("I'm a scheduler!")
	rand.Seed(time.Now().Unix())

	podQueue := make(chan *corev1.Pod, 300)
	defer close(podQueue)

	quit := make(chan struct{})
	defer close(quit)

	// 创建一个 Scheduler 对象
	scheduler := NewScheduler(podQueue, quit)
	// 启动 Scheduler
	scheduler.Run()
}
