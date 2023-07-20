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

type TransitGatewayRouteTableInitParameters_2 struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=TransitGateway
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

type TransitGatewayRouteTableObservation_2 struct {

	// EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable *bool `json:"defaultAssociationRouteTable,omitempty" tf:"default_association_route_table,omitempty"`

	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable *bool `json:"defaultPropagationRouteTable,omitempty" tf:"default_propagation_route_table,omitempty"`

	// EC2 Transit Gateway Route Table identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`
}

type TransitGatewayRouteTableParameters_2 struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=TransitGateway
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

// TransitGatewayRouteTableSpec defines the desired state of TransitGatewayRouteTable
type TransitGatewayRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayRouteTableParameters_2 `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider TransitGatewayRouteTableInitParameters_2 `json:"initProvider,omitempty"`
}

// TransitGatewayRouteTableStatus defines the observed state of TransitGatewayRouteTable.
type TransitGatewayRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayRouteTableObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTable is the Schema for the TransitGatewayRouteTables API. Manages an EC2 Transit Gateway Route Table
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayRouteTableSpec   `json:"spec"`
	Status            TransitGatewayRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayRouteTableList contains a list of TransitGatewayRouteTables
type TransitGatewayRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayRouteTable `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayRouteTable_Kind             = "TransitGatewayRouteTable"
	TransitGatewayRouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayRouteTable_Kind}.String()
	TransitGatewayRouteTable_KindAPIVersion   = TransitGatewayRouteTable_Kind + "." + CRDGroupVersion.String()
	TransitGatewayRouteTable_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayRouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayRouteTable{}, &TransitGatewayRouteTableList{})
}
