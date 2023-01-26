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
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	samplev1 "github.com/zhengyansheng/sample-operator/api/v1"
)

const (
	selectorKey = "sample.zhengyansheng.com/name"
)

// NginxReconciler reconciles a Nginx object
type NginxReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
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
	pod := corev1.Pod{}
	err := r.Get(ctx, req.NamespacedName, &pod)
	if errors.IsNotFound(err) {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// reference: https://yash-kukreja-98.medium.com/develop-on-kubernetes-series-demystifying-the-for-vs-owns-vs-watches-controller-builders-in-c11ab32a046e
	// channel which will act as the source for reconciliation
	reconciliationSourceChannel := make(chan event.GenericEvent)

	// running the time checker as a goroutine
	// this goroutine will periodically check the local time and trigger a reconciliation via the above channel at 12:00AM
	// for MyCustomResource with name: "foo" and namespace: "default"
	go func() {
		// every 45 seconds, run the time-check and trigger an event if the time is 12:00AM
		ticker := time.NewTicker(2 * time.Minute)
		defer ticker.Stop()
		for _ = range ticker.C {
			//now := time.Now()
			//// if 12:00AM, trigger the reconciliation by firing a Generic event against the channel registered as a source to `Watches()` containing the object to reconcile
			//if now.Local().Hour() == 0 && now.Local().Minute() == 0 {
			//	object := samplev1.Nginx{ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"}}
			//	reconciliationSourceChannel <- event.GenericEvent{Object: &object}
			//}

			object := samplev1.Nginx{ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"}}
			reconciliationSourceChannel <- event.GenericEvent{Object: &object}
			r.EventRecorder.Event(&object, corev1.EventTypeNormal, "Upgrade", "trigger event")
		}
	}()

	return ctrl.NewControllerManagedBy(mgr).
		For(&samplev1.Nginx{}).
		Watches(&source.Channel{Source: reconciliationSourceChannel}, &handler.EnqueueRequestForObject{}).
		Watches(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
