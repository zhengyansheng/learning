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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	elasticwebv1 "github.com/zhengyansheng/sample-operator/elasticweb-operator/api/v1"
)

const (
	appName       = "elastic-app"
	containerPort = 8080
	cpuRequest    = "100m"
	cpuLimit      = "100m"
	memRequest    = "512Mi"
	memLimit      = "512Mi"

	selectorKey     = "app"
	defaultPortName = "http"
)

// ElasticWebReconciler reconciles a ElasticWeb object
type ElasticWebReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=elasticweb.zhengyansheng.com,resources=elasticwebs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=elasticweb.zhengyansheng.com,resources=elasticwebs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=elasticweb.zhengyansheng.com,resources=elasticwebs/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;update;patch;watch;create;delete;list
//+kubebuilder:rbac:groups=apps,resources=services,verbs=get;update;patch;watch;create;delete;list

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ElasticWeb object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ElasticWebReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	instance := &elasticwebv1.ElasticWeb{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		//return ctrl.Result{RequeueAfter: time.Second * 10}, err
		return ctrl.Result{}, err
	}

	// deployment
	foundDeployment := &appsv1.Deployment{}
	err = r.Get(ctx, req.NamespacedName, foundDeployment)
	if err != nil {

		if errors.IsNotFound(err) {
			// deployment 不存在时，recreate

			if *(instance.Spec.TotalQPS) < 1 {
				return ctrl.Result{}, nil
			}

			// create deployment
			if err = r.createDeploymentIfNotExists(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}

			// create svc
			if err = r.createServiceIfNotExists(ctx, req, instance); err != nil {
				return ctrl.Result{}, err
			}

			// create ingress
			if err = r.createIngressIfNotExists(ctx, req, instance); err != nil {
				return ctrl.Result{}, err
			}

			// update status
			if err = r.updateStatus(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	expectReplicas := getExpectReplicas(instance)

	realReplicas := *foundDeployment.Spec.Replicas
	klog.Infof("expectReplicas [%d], realReplicas [%d]", expectReplicas, realReplicas)

	if expectReplicas == realReplicas {
		// 如果相等 直接返回
		return ctrl.Result{}, nil
	}

	// 如果不相等，则调整
	*(foundDeployment.Spec.Replicas) = expectReplicas
	if err := r.Update(ctx, foundDeployment); err != nil {
		return ctrl.Result{}, err
	}

	if err := r.updateStatus(ctx, instance); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ElasticWebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&elasticwebv1.ElasticWeb{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&networkv1.Ingress{}).
		Complete(r)
}
