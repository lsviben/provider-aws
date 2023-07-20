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

type RuleAssociationInitParameters struct {

	// A name for the association that you're creating between a resolver rule and a VPC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the resolver rule that you want to associate with the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53resolver/v1beta1.Rule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ResolverRuleID *string `json:"resolverRuleId,omitempty" tf:"resolver_rule_id,omitempty"`

	ResolverRuleIDRef *v1.Reference `json:"resolverRuleIdRef,omitempty" tf:"-"`

	ResolverRuleIDSelector *v1.Selector `json:"resolverRuleIdSelector,omitempty" tf:"-"`

	// The ID of the VPC that you want to associate the resolver rule with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RuleAssociationObservation struct {

	// The ID of the resolver rule association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A name for the association that you're creating between a resolver rule and a VPC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleID *string `json:"resolverRuleId,omitempty" tf:"resolver_rule_id,omitempty"`

	// The ID of the VPC that you want to associate the resolver rule with.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type RuleAssociationParameters struct {

	// A name for the association that you're creating between a resolver rule and a VPC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the resolver rule that you want to associate with the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53resolver/v1beta1.Rule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ResolverRuleID *string `json:"resolverRuleId,omitempty" tf:"resolver_rule_id,omitempty"`

	// Reference to a Rule in route53resolver to populate resolverRuleId.
	// +kubebuilder:validation:Optional
	ResolverRuleIDRef *v1.Reference `json:"resolverRuleIdRef,omitempty" tf:"-"`

	// Selector for a Rule in route53resolver to populate resolverRuleId.
	// +kubebuilder:validation:Optional
	ResolverRuleIDSelector *v1.Selector `json:"resolverRuleIdSelector,omitempty" tf:"-"`

	// The ID of the VPC that you want to associate the resolver rule with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RuleAssociationSpec defines the desired state of RuleAssociation
type RuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider RuleAssociationInitParameters `json:"initProvider,omitempty"`
}

// RuleAssociationStatus defines the observed state of RuleAssociation.
type RuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleAssociation is the Schema for the RuleAssociations API. Provides a Route53 Resolver rule association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleAssociationSpec   `json:"spec"`
	Status            RuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleAssociationList contains a list of RuleAssociations
type RuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	RuleAssociation_Kind             = "RuleAssociation"
	RuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleAssociation_Kind}.String()
	RuleAssociation_KindAPIVersion   = RuleAssociation_Kind + "." + CRDGroupVersion.String()
	RuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(RuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleAssociation{}, &RuleAssociationList{})
}
