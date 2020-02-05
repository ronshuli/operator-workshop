package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RonHelloWorldSpec defines the desired state of RonHelloWorld
type RonHelloWorldSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message string `json:"message"`
}

// RonHelloWorldStatus defines the observed state of RonHelloWorld
type RonHelloWorldStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RonHelloWorld is the Schema for the ronhelloworlds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ronhelloworlds,scope=Namespaced
type RonHelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RonHelloWorldSpec   `json:"spec,omitempty"`
	Status RonHelloWorldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RonHelloWorldList contains a list of RonHelloWorld
type RonHelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RonHelloWorld `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RonHelloWorld{}, &RonHelloWorldList{})
}
