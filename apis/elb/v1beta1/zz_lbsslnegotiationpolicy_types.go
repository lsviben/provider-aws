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

type AttributeInitParameters struct {

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the attribute
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AttributeObservation struct {

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the attribute
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AttributeParameters struct {

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the attribute
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LBSSLNegotiationPolicyInitParameters struct {

	// An SSL Negotiation policy attribute. Each has two properties:
	Attribute []AttributeInitParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The load balancer to which the policy
	// should be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	LoadBalancerRef *v1.Reference `json:"loadBalancerRef,omitempty" tf:"-"`

	LoadBalancerSelector *v1.Selector `json:"loadBalancerSelector,omitempty" tf:"-"`

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type LBSSLNegotiationPolicyObservation struct {

	// An SSL Negotiation policy attribute. Each has two properties:
	Attribute []AttributeObservation `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The load balancer to which the policy
	// should be attached.
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

type LBSSLNegotiationPolicyParameters struct {

	// An SSL Negotiation policy attribute. Each has two properties:
	Attribute []AttributeParameters `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The load balancer to which the policy
	// should be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// Reference to a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerRef *v1.Reference `json:"loadBalancerRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerSelector *v1.Selector `json:"loadBalancerSelector,omitempty" tf:"-"`

	// The name of the SSL negotiation policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`
}

// LBSSLNegotiationPolicySpec defines the desired state of LBSSLNegotiationPolicy
type LBSSLNegotiationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBSSLNegotiationPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider LBSSLNegotiationPolicyInitParameters `json:"initProvider,omitempty"`
}

// LBSSLNegotiationPolicyStatus defines the observed state of LBSSLNegotiationPolicy.
type LBSSLNegotiationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBSSLNegotiationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBSSLNegotiationPolicy is the Schema for the LBSSLNegotiationPolicys API. Provides a load balancer SSL negotiation policy, which allows an ELB to control which ciphers and protocols are supported during SSL negotiations between a client and a load balancer.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBSSLNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lbPort) || has(self.initProvider.lbPort)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	Spec   LBSSLNegotiationPolicySpec   `json:"spec"`
	Status LBSSLNegotiationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBSSLNegotiationPolicyList contains a list of LBSSLNegotiationPolicys
type LBSSLNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBSSLNegotiationPolicy `json:"items"`
}

// Repository type metadata.
var (
	LBSSLNegotiationPolicy_Kind             = "LBSSLNegotiationPolicy"
	LBSSLNegotiationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBSSLNegotiationPolicy_Kind}.String()
	LBSSLNegotiationPolicy_KindAPIVersion   = LBSSLNegotiationPolicy_Kind + "." + CRDGroupVersion.String()
	LBSSLNegotiationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LBSSLNegotiationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LBSSLNegotiationPolicy{}, &LBSSLNegotiationPolicyList{})
}
