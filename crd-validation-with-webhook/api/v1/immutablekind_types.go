/*
Copyright 2022 rewanthtammana.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImmutableKindSpec defines the desired state of ImmutableKind
type ImmutableKindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ImmutableKind. Edit immutablekind_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ImmutableKindStatus defines the observed state of ImmutableKind
type ImmutableKindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ImmutableKind is the Schema for the immutablekinds API
type ImmutableKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImmutableKindSpec   `json:"spec,omitempty"`
	Status ImmutableKindStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImmutableKindList contains a list of ImmutableKind
type ImmutableKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImmutableKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImmutableKind{}, &ImmutableKindList{})
}
