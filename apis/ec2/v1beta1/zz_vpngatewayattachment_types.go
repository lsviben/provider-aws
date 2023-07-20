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

type VPNGatewayAttachmentInitParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Private Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`
}

type VPNGatewayAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The ID of the Virtual Private Gateway.
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`
}

type VPNGatewayAttachmentParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ID of the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The ID of the Virtual Private Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPNGateway
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	VPNGatewayID *string `json:"vpnGatewayId,omitempty" tf:"vpn_gateway_id,omitempty"`

	// Reference to a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDRef *v1.Reference `json:"vpnGatewayIdRef,omitempty" tf:"-"`

	// Selector for a VPNGateway in ec2 to populate vpnGatewayId.
	// +kubebuilder:validation:Optional
	VPNGatewayIDSelector *v1.Selector `json:"vpnGatewayIdSelector,omitempty" tf:"-"`
}

// VPNGatewayAttachmentSpec defines the desired state of VPNGatewayAttachment
type VPNGatewayAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayAttachmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider VPNGatewayAttachmentInitParameters `json:"initProvider,omitempty"`
}

// VPNGatewayAttachmentStatus defines the observed state of VPNGatewayAttachment.
type VPNGatewayAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayAttachment is the Schema for the VPNGatewayAttachments API. Provides a Virtual Private Gateway attachment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPNGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewayAttachmentSpec   `json:"spec"`
	Status            VPNGatewayAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayAttachmentList contains a list of VPNGatewayAttachments
type VPNGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGatewayAttachment `json:"items"`
}

// Repository type metadata.
var (
	VPNGatewayAttachment_Kind             = "VPNGatewayAttachment"
	VPNGatewayAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGatewayAttachment_Kind}.String()
	VPNGatewayAttachment_KindAPIVersion   = VPNGatewayAttachment_Kind + "." + CRDGroupVersion.String()
	VPNGatewayAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VPNGatewayAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGatewayAttachment{}, &VPNGatewayAttachmentList{})
}
