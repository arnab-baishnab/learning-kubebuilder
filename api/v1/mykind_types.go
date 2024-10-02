/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// MyKind is the Schema for the mykinds API
type MyKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyKindSpec   `json:"spec,omitempty"`
	Status MyKindStatus `json:"status,omitempty"`
}

// MyKindSpec defines the desired state of MyKind
type MyKindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Quantity of instances
	// +optional
	DeploymentName string `json:"deploymentName,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	Replicas  *int32        `json:"replicas,omitempty"`
	Container ContainerSpec `json:"mycontainer"`

	// +optional
	Service ServiceSpec `json:"service,omitempty"`
}

type ContainerSpec struct {
	Image string `json:"image,omitempty"`
	Port  int32  `json:"port,omitempty"`
}

type ServiceSpec struct {
	// +optional
	ServiceName string `json:"serviceName"`
	ServiceType string `json:"serviceType,omitempty"`
	// +optional
	ServiceNodePort int32 `json:"serviceNodePort,omitempty"`
}

// MyKindStatus defines the observed state of MyKind
type MyKindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +optional
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +kubebuilder:object:root=true

// MyKindList contains a list of MyKind
type MyKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyKind{}, &MyKindList{})
}
