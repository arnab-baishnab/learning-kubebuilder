package controller

import (
	crdv1 "github.com/arnab-baishnab/learning-kubebuilder/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	_ "strings"
)

func (r *MyKindReconciler) CheckDeployment() error {

	deploy := &appsv1.Deployment{}

	if err := r.Client.Get(r.ctx, types.NamespacedName{
		Name:      r.myKind.DeploymentName(),
		Namespace: "demo",
	}, deploy); err != nil {
		if errors.IsNotFound(err) {
			r.Log.Info("Creating a new Deployment", "Namespace", r.myKind.Namespace)
			deploy := r.NewDeployment()
			if err := r.Client.Create(r.ctx, deploy); err != nil {
				return err
			}
			r.Log.Info("Created Deployment", "Namespace", deploy.Namespace, "Name", deploy.Name)
			return nil
		}
		return err
	}
	if r.myKind.Spec.Replicas != nil && *deploy.Spec.Replicas != *r.myKind.Spec.Replicas {
		r.Log.Info("replica mismatch...")
		*deploy.Spec.Replicas = *r.myKind.Spec.Replicas
		if err := r.Client.Update(r.ctx, deploy); err != nil {
			r.Log.Error(err, "Failed to update Deployment", "Namespace", deploy.Namespace, "Name", deploy.Name)
			return err
		}
	}

	return nil
}

func (r *MyKindReconciler) NewDeployment() *appsv1.Deployment {
	r.Log.Info("New Deployment is called")
	labels := map[string]string{
		"controller": r.myKind.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      r.myKind.DeploymentName(),
			Namespace: r.myKind.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(r.myKind, crdv1.GroupVersion.WithKind(ourKind)),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Replicas: r.myKind.Spec.Replicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "mycontainer",
							Image: r.myKind.Spec.Container.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: r.myKind.Spec.Container.Port,
								},
							},
						},
					},
				},
			},
		},
	}
}
