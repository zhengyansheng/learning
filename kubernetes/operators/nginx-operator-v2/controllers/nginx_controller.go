/*
Copyright 2022 zhengyansheng.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	samplev1 "github.com/zhengyansheng/sample-operator/api/v1"
)

const (
	selectorKey = "sample.zhengyansheng.com/name"
)

// NginxReconciler reconciles a Nginx object
type NginxReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	//增加事件记录
	EventRecorder record.EventRecorder
}

//+kubebuilder:rbac:groups=apps.zhengyansheng.com,resources=nginxes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.zhengyansheng.com,resources=nginxes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.zhengyansheng.com,resources=nginxes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Nginx object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *NginxReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	defer utilruntime.HandleCrash()
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	klog.Infof("%s/ %s", req.Namespace, req.Name)

	// 生成 Nginx struct
	ngx := &samplev1.Nginx{}
	err := r.Client.Get(ctx, req.NamespacedName, ngx)
	if errors.IsNotFound(err) || !ngx.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}
	if err != nil {
		return ctrl.Result{}, err
	}
	// if resource is deleting, so continue
	if ngx.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	// create deployment
	foundDeployment := appsv1.Deployment{}
	err = r.Client.Get(ctx, req.NamespacedName, &foundDeployment)
	if errors.IsNotFound(err) {
		foundDeployment, err := r.buildDeployment(ngx)
		if err != nil {
			klog.Error(err, "failed to build Deployment resource")
			return ctrl.Result{}, err
		}
		if err = r.Create(ctx, foundDeployment); err != nil {
			klog.Error(err, "failed to create Deployment resource")
			return ctrl.Result{}, err
		}

		r.EventRecorder.Eventf(foundDeployment, corev1.EventTypeNormal, "Created", "Created deployment %v", foundDeployment.Name)
		klog.Info("created Deployment resource for nginx")

		// 关联 annotations
		r.EventRecorder.Event(ngx, corev1.EventTypeNormal, "Relation", "relation annotations")
		ngxSpecByte, err := json.Marshal(ngx.Spec)
		if err != nil {
			klog.Warning(err, "json marshal err")
			return ctrl.Result{}, err
		}
		if ngx.Annotations == nil {
			ngx.Annotations["spec"] = string(ngxSpecByte)
		} else {
			ngx.Annotations = map[string]string{"spec": string(ngxSpecByte)}
		}
		r.EventRecorder.Event(ngx, corev1.EventTypeNormal, "Updated", "Update deployment nginx")
		if err := r.Update(ctx, ngx); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}
	if err != nil {
		klog.Error(err, "failed to get Deployment for nginx resource")
		return ctrl.Result{}, err
	}

	oldSpec := samplev1.NginxSpec{}
	if err := json.Unmarshal([]byte(ngx.Annotations["spec"]), &oldSpec); err != nil {
		klog.Warning(err, "json unmarshal annotations spec err")
		return ctrl.Result{}, err
	}

	// 深度比较spec部分是否变动，如果变动就更新deployment
	klog.Infof("old spec: %+v\n", oldSpec)
	klog.Infof("new spec: %+v\n", ngx.Spec)
	if !reflect.DeepEqual(oldSpec, ngx.Spec) {
		klog.Info("nginx spec modified")
		// 渲染新的deployment
		newDeployment, err := r.buildDeployment(ngx)
		if err != nil {
			klog.Warning(err, "build deployment err")
			return ctrl.Result{}, err
		}

		// 当前运行的deployment
		oldDeployment := appsv1.Deployment{}
		err = r.Client.Get(ctx, req.NamespacedName, &oldDeployment)
		if err != nil {
			klog.Warning(err, "get old deployment err")
			return ctrl.Result{}, err
		}

		ngxSpecByte, err := json.Marshal(oldDeployment.Spec)
		if err != nil {
			klog.Warning(err, "json marshal nginx annotations spec err")
			return reconcile.Result{}, err
		}
		ngx.Annotations = map[string]string{"spec": string(ngxSpecByte)}

		oldDeployment.Spec = newDeployment.Spec
		klog.Infof("old deployment replicas: %+v", oldDeployment.Spec.Replicas)
		klog.Infof("old deployment containers: %+v", oldDeployment.Spec.Template.Spec.Containers[0].Image)
		if err := r.Update(ctx, &oldDeployment); err != nil {
			klog.Warning(err, "update deployment err")
			return reconcile.Result{}, err
		}

		klog.Info("update nginx annotations spec")
		if err := r.Update(ctx, ngx); err != nil {
			klog.Warning(err, "update nginx spec err")
			return ctrl.Result{}, err
		}
	}
	klog.Info("resource status synced")

	return ctrl.Result{}, nil
}

func (r *NginxReconciler) buildDeployment(c *samplev1.Nginx) (*appsv1.Deployment, error) {
	deployMent := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(c, schema.GroupVersionKind{
					Group:   samplev1.GroupVersion.Group,
					Version: samplev1.GroupVersion.Version,
					Kind:    "Nginx",
				}),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &c.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					selectorKey: c.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						selectorKey: c.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            c.Name,
							Image:           c.Spec.Image,
							ImagePullPolicy: "IfNotPresent",
						},
					},
				},
			},
		},
	}

	// owner reference
	// https://zhuanlan.zhihu.com/p/67406200
	if err := controllerutil.SetControllerReference(c, deployMent, r.Scheme); err != nil {
		return deployMent, err
	}
	return deployMent, nil
}

func (r *NginxReconciler) deleteHandler(event event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
	klog.Info("delete handle by callback")
	for _, ownerReference := range event.Object.GetOwnerReferences() {
		//if ownerReference.Kind == "Nginx" && ownerReference.Name == "redis" {
		klog.Info(">>>: ", ownerReference.Kind, ownerReference.Name, ownerReference.Controller)
		if ownerReference.Kind == "Nginx" {
			limitingInterface.Add(reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: event.Object.GetNamespace(),
					Name:      ownerReference.Name,
				},
			})
		}
	}
}

func (r *NginxReconciler) updateHandler(event event.UpdateEvent, limitingInterface workqueue.RateLimitingInterface) {

}

//func (r *NginxReconciler) deleteRelationResource(ctx context.Context, req ctrl.Request, ngx *samplev1.Nginx) error {
//	klog.Info("delete relation resource")
//	// delete deployment
//	instance := appsv1.Deployment{}
//	err := r.Client.Get(ctx, runtimeObjectKey(req), &instance)
//	if err != nil {
//		return err
//	}
//	err = r.Client.Delete(ctx, &instance)
//	if err != nil {
//		return err
//	}
//	r.EventRecorder.Event(ngx, corev1.EventTypeNormal, "Updated", "Delete deployment")
//	return nil
//}

func (r *NginxReconciler) getDeploymentSpec(c *samplev1.Nginx, labels map[string]string) appsv1.DeploymentSpec {
	return appsv1.DeploymentSpec{
		Replicas: &c.Spec.Replicas,
		Selector: metav1.SetAsLabelSelector(labels),
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            c.Name,
						Image:           c.Spec.Image,
						ImagePullPolicy: "IfNotPresent",
					},
				},
			},
		},
	}

}

func (r *NginxReconciler) getSpecFromDeployment(deploy *appsv1.Deployment) appsv1.DeploymentSpec {
	container := deploy.Spec.Template.Spec.Containers[0]
	return appsv1.DeploymentSpec{
		Replicas: deploy.Spec.Replicas,
		Selector: deploy.Spec.Selector,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: deploy.Spec.Template.Labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            container.Name,
						Image:           container.Image,
						ImagePullPolicy: "IfNotPresent",
					},
				},
			},
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *NginxReconciler) SetupWithManager(mgr ctrl.Manager) error {

	// reference https://github.com/kubernetes-sigs/kubebuilder/issues/549
	r.EventRecorder = mgr.GetEventRecorderFor("nginx")

	return ctrl.NewControllerManagedBy(mgr).
		For(&samplev1.Nginx{}). // 注意这里需要使用指针类型的Kind，因为其start方法接收者为指针类型
		// 使用 CR 创建 deployment 时，可以为他塞入一个从属关系，类似于 Pod 资源的Metadata 里会有一个OnwerReference字段
		Owns(&appsv1.Deployment{}).
		//Watches(&source.Kind{Type: &samplev1.Nginx{}}, handler.Funcs{
		//	DeleteFunc: r.deleteHandler,
		//	UpdateFunc: r.updateHandler,
		//}).
		Complete(r)
}
