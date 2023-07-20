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

type VPCLinkInitParameters struct {

	// Description of the VPC link.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name used to label and identify the VPC link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TargetArnRefs []v1.Reference `json:"targetArnRefs,omitempty" tf:"-"`

	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`

	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LB
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +crossplane:generate:reference:refFieldName=TargetArnRefs
	// +crossplane:generate:reference:selectorFieldName=TargetArnSelector
	TargetArns []*string `json:"targetArns,omitempty" tf:"target_arns,omitempty"`
}

type VPCLinkObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the VPC link.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Identifier of the VpcLink.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name used to label and identify the VPC link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArns []*string `json:"targetArns,omitempty" tf:"target_arns,omitempty"`
}

type VPCLinkParameters struct {

	// Description of the VPC link.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name used to label and identify the VPC link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to LB in elbv2 to populate targetArns.
	// +kubebuilder:validation:Optional
	TargetArnRefs []v1.Reference `json:"targetArnRefs,omitempty" tf:"-"`

	// Selector for a list of LB in elbv2 to populate targetArns.
	// +kubebuilder:validation:Optional
	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`

	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LB
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +crossplane:generate:reference:refFieldName=TargetArnRefs
	// +crossplane:generate:reference:selectorFieldName=TargetArnSelector
	TargetArns []*string `json:"targetArns,omitempty" tf:"target_arns,omitempty"`
}

// VPCLinkSpec defines the desired state of VPCLink
type VPCLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCLinkParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider VPCLinkInitParameters `json:"initProvider,omitempty"`
}

// VPCLinkStatus defines the observed state of VPCLink.
type VPCLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCLink is the Schema for the VPCLinks API. Provides an API Gateway VPC Link.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	Spec   VPCLinkSpec   `json:"spec"`
	Status VPCLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCLinkList contains a list of VPCLinks
type VPCLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCLink `json:"items"`
}

// Repository type metadata.
var (
	VPCLink_Kind             = "VPCLink"
	VPCLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCLink_Kind}.String()
	VPCLink_KindAPIVersion   = VPCLink_Kind + "." + CRDGroupVersion.String()
	VPCLink_GroupVersionKind = CRDGroupVersion.WithKind(VPCLink_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCLink{}, &VPCLinkList{})
}
