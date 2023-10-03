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

type TransitGatewayRouteTableAssociationInitParameters struct {
}

type TransitGatewayRouteTableAssociationObservation struct {

	// EC2 Transit Gateway Route Table identifier combined with EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the resource
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Type of the resource
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`
}

type TransitGatewayRouteTableAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Identifier of EC2 Transit Gateway Attachment.
	// +crossplane:generate:reference:type=TransitGatewayVPCAttachment
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDRef *v1.Reference `json:"transitGatewayAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment to populate transitGatewayAttachmentId.
	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentIDSelector *v1.Selector `json:"transitGatewayAttachmentIdSelector,omitempty" tf:"-"`

	// Identifier of EC2 Transit Gateway Route Table.
	// +crossplane:generate:reference:type=TransitGatewayRouteTable
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableID *string `json:"transitGatewayRouteTableId,omitempty" tf:"transit_gateway_route_table_id,omitempty"`

	// Reference to a TransitGatewayRouteTable to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDRef *v1.Reference `json:"transitGatewayRouteTableIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayRouteTable to populate transitGatewayRouteTableId.
	// +kubebuilder:validation:Optional
	TransitGatewayRouteTableIDSelector *v1.Selector `json:"transitGatewayRouteTableIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteTableAssociationSpec defines the desired state of TransitGatewayRouteTableAssociation
type TransitGatewayRouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteTableAssociationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It is not honored
	// unless the relevant Crossplane feature flag is enabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TransitGatewayRouteTableAssociationInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayRouteTableAssociationStatus defines the observed state of TransitGatewayRouteTableAssociation.
type TransitGatewayRouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTableAssociation is the Schema for the TransitGatewayRouteTableAssociations API. Manages an EC2 Transit Gateway Route Table association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTableAssociationSpec   `json:"spec"`
	Status            TransitGatewayRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTableAssociationList contains a list of TransitGatewayRouteTableAssociations
type TransitGatewayRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTableAssociation_Kind             = "TransitGatewayRouteTableAssociation"
	TransitGatewayRouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteTableAssociation_Kind}.String()
	TransitGatewayRouteTableAssociation_KindAPIVersion   = TransitGatewayRouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTableAssociation{}, &TransitGatewayRouteTableAssociationList{})
}
