package controller

import (
	"fmt"
	crdv1 "github.com/arnab-baishnab/learning-kubebuilder/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (r *MyKindReconciler) CheckService() error {
	srv := &corev1.Service{}
	if err := r.Client.Get(r.ctx, types.NamespacedName{
		Name:      r.myKind.ServiceName(),
		Namespace: r.myKind.Namespace,
	}, srv); err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info("Bookserver service not found")
			if err := r.Client.Create(r.ctx, r.newService()); err != nil {
				r.Log.Error(err, "failed to create service")
				return err
			}
			r.Log.Info("created service")
			return nil
		}
		return err
	}
	return nil
}

func (r *MyKindReconciler) newService() *corev1.Service {
	fmt.Println("New Service is called")
	labels := map[string]string{
		"controller": r.myKind.Name,
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      r.myKind.ServiceName(),
			Namespace: r.myKind.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(r.myKind, crdv1.GroupVersion.WithKind(ourKind)),
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     getServiceType(r.myKind.Spec.Service.ServiceName),
			Ports: []corev1.ServicePort{
				{
					Port:       r.myKind.Spec.Container.Port,
					NodePort:   r.myKind.Spec.Service.ServiceNodePort,
					TargetPort: intstr.FromInt32(r.myKind.Spec.Container.Port),
				},
			},
		},
	}
}
