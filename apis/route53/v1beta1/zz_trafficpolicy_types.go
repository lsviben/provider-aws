/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TrafficPolicyObservation struct {

	// Comment for the traffic policy.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the AWS Route53 Traffic Policy document format
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// ID of the traffic policy
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the traffic policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Version number of the traffic policy. This value is automatically incremented by AWS after each update of this resource.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type TrafficPolicyParameters struct {

	// Comment for the traffic policy.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the AWS Route53 Traffic Policy document format
	// +kubebuilder:validation:Optional
	Document *string `json:"document,omitempty" tf:"document,omitempty"`

	// Name of the traffic policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// TrafficPolicySpec defines the desired state of TrafficPolicy
type TrafficPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficPolicyParameters `json:"forProvider"`
}

// TrafficPolicyStatus defines the observed state of TrafficPolicy.
type TrafficPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicy is the Schema for the TrafficPolicys API. Manages a Route53 Traffic Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TrafficPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.document)",message="document is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   TrafficPolicySpec   `json:"spec"`
	Status TrafficPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficPolicyList contains a list of TrafficPolicys
type TrafficPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPolicy `json:"items"`
}

// Repository type metadata.
var (
	TrafficPolicy_Kind             = "TrafficPolicy"
	TrafficPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficPolicy_Kind}.String()
	TrafficPolicy_KindAPIVersion   = TrafficPolicy_Kind + "." + CRDGroupVersion.String()
	TrafficPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TrafficPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficPolicy{}, &TrafficPolicyList{})
}
