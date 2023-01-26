package controllers

import (
	"context"
	"fmt"

	elasticwebv1 "github.com/zhengyansheng/sample-operator/elasticweb-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func getExpectReplicas(web *elasticwebv1.ElasticWeb) int32 {
	// 单个Pod 的QPS
	singlePodQPS := *(web.Spec.SinglePodQPS)

	// 期望的QPS
	totalQPS := *(web.Spec.TotalQPS)

	// 期望的Pod个数
	replicas := totalQPS / singlePodQPS
	if totalQPS%singlePodQPS > 0 {
		replicas++
	}
	return replicas
}

func (r *ElasticWebReconciler) createDeploymentIfNotExists(ctx context.Context, web *elasticwebv1.ElasticWeb) error {
	// calculate expect replicas
	expectReplicas := getExpectReplicas(web)
	klog.Infof("expect replicas %d", expectReplicas)

	// new deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: web.Namespace,
			Name:      web.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &expectReplicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					selectorKey: appName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						selectorKey: appName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            appName,
							Image:           web.Spec.Image,
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolSCTP,
									ContainerPort: containerPort,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									"cpu":    resource.MustParse(cpuRequest),
									"memory": resource.MustParse(memRequest),
								},
								Limits: corev1.ResourceList{
									"cpu":    resource.MustParse(cpuLimit),
									"memory": resource.MustParse(memLimit),
								},
							},
						},
					},
				},
			},
		},
	}
	klog.Info("set deployment reference")
	if err := controllerutil.SetControllerReference(web, deployment, r.Scheme); err != nil {
		klog.Error(err, "set deployment controller reference err")
		return err
	}

	// create deployment
	if err := r.Create(ctx, deployment); err != nil {
		klog.Error(err, "create deployment err")
		return err
	}
	klog.Info("create deployment success")
	return nil

}

func (r *ElasticWebReconciler) createServiceIfNotExists(ctx context.Context, req ctrl.Request, web *elasticwebv1.ElasticWeb) error {
	service := &corev1.Service{}
	err := r.Get(ctx, req.NamespacedName, service)
	if err == nil {
		// 如果service存在，则退出
		return nil
	}
	if !errors.IsNotFound(err) {
		// 如果错误不是not found，则表示异常，直接返回错误
		return err
	}

	// new service
	service = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: web.Namespace,
			Name:      web.Name,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:     defaultPortName,
				Port:     8080,
				NodePort: web.Spec.Port,
			}},
			Selector: map[string]string{
				selectorKey: appName,
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	klog.Info("set reference by service")
	if err := controllerutil.SetControllerReference(web, service, r.Scheme); err != nil {
		klog.Error(err, "set service controller reference err")
		return err
	}

	// create service
	if err := r.Create(ctx, service); err != nil {
		klog.Error(err, "create service err")
		return err
	}
	klog.Info("create service success")
	return nil
}

func (r *ElasticWebReconciler) createIngressIfNotExists(ctx context.Context, req ctrl.Request, web *elasticwebv1.ElasticWeb) error {
	ingress := &networkv1.Ingress{}
	err := r.Get(ctx, req.NamespacedName, ingress)
	if err == nil {
		// 如果service存在，则退出
		return nil
	}
	if !errors.IsNotFound(err) {
		// 如果错误不是not found，则表示异常，直接返回错误
		return err
	}

	// new ingress
	pathType := networkv1.PathTypeImplementationSpecific
	ingress = &networkv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: web.Namespace,
			Name:      web.Name,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(web, schema.GroupVersionKind{
					Group:   elasticwebv1.GroupVersion.Group,
					Version: elasticwebv1.GroupVersion.Version,
					Kind:    "ElasticWeb",
				}),
			},
		},
		Spec: networkv1.IngressSpec{
			Rules: []networkv1.IngressRule{
				{
					Host: fmt.Sprintf("%v.zhengyansheng.com", web.Name),
					IngressRuleValue: networkv1.IngressRuleValue{
						HTTP: &networkv1.HTTPIngressRuleValue{
							Paths: []networkv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networkv1.IngressBackend{
										Service: &networkv1.IngressServiceBackend{
											Name: web.Name,
											Port: networkv1.ServiceBackendPort{
												//Name:   defaultPortName,
												Number: int32(containerPort),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	klog.Info("set reference by ingress")
	if err := controllerutil.SetControllerReference(web, ingress, r.Scheme); err != nil {
		klog.Error(err, "set ingress controller reference err")
		return err
	}

	// create ingress
	if err := r.Create(ctx, ingress); err != nil {
		klog.Error(err, "create ingress err")
		return err
	}
	klog.Info("create ingress success")
	return nil
}

func (r *ElasticWebReconciler) updateStatus(ctx context.Context, web *elasticwebv1.ElasticWeb) error {
	singlePodQPS := *(web.Spec.SinglePodQPS)

	replicas := getExpectReplicas(web)

	if web.Status.RealQPS == nil {
		web.Status.RealQPS = new(int32)
	}

	*(web.Status.RealQPS) = singlePodQPS * replicas
	klog.Infof("singlePodQPS [%d], replicas [%d], realQPS[%d]", singlePodQPS, replicas, *(web.Status.RealQPS))

	if err := r.Update(ctx, web); err != nil {
		klog.Error(err, "update web err")
		return err
	}
	return nil
}
