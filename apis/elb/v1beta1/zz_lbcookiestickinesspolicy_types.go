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

type LBCookieStickinessPolicyInitParameters struct {

	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod *float64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The name of the stickiness policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LBCookieStickinessPolicyObservation struct {

	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod *float64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The load balancer to which the policy
	// should be attached.
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// The name of the stickiness policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type LBCookieStickinessPolicyParameters struct {

	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	// +kubebuilder:validation:Optional
	CookieExpirationPeriod *float64 `json:"cookieExpirationPeriod,omitempty" tf:"cookie_expiration_period,omitempty"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	// +kubebuilder:validation:Optional
	LBPort *float64 `json:"lbPort,omitempty" tf:"lb_port,omitempty"`

	// The load balancer to which the policy
	// should be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LoadBalancer *string `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// Reference to a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerRef *v1.Reference `json:"loadBalancerRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancer.
	// +kubebuilder:validation:Optional
	LoadBalancerSelector *v1.Selector `json:"loadBalancerSelector,omitempty" tf:"-"`

	// The name of the stickiness policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBCookieStickinessPolicySpec defines the desired state of LBCookieStickinessPolicy
type LBCookieStickinessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBCookieStickinessPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LBCookieStickinessPolicyInitParameters `json:"initProvider,omitempty"`
}

// LBCookieStickinessPolicyStatus defines the observed state of LBCookieStickinessPolicy.
type LBCookieStickinessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBCookieStickinessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBCookieStickinessPolicy is the Schema for the LBCookieStickinessPolicys API. Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBCookieStickinessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lbPort) || (has(self.initProvider) && has(self.initProvider.lbPort))",message="spec.forProvider.lbPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LBCookieStickinessPolicySpec   `json:"spec"`
	Status LBCookieStickinessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBCookieStickinessPolicyList contains a list of LBCookieStickinessPolicys
type LBCookieStickinessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBCookieStickinessPolicy `json:"items"`
}

// Repository type metadata.
var (
	LBCookieStickinessPolicy_Kind             = "LBCookieStickinessPolicy"
	LBCookieStickinessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBCookieStickinessPolicy_Kind}.String()
	LBCookieStickinessPolicy_KindAPIVersion   = LBCookieStickinessPolicy_Kind + "." + CRDGroupVersion.String()
	LBCookieStickinessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LBCookieStickinessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LBCookieStickinessPolicy{}, &LBCookieStickinessPolicyList{})
}
