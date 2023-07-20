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

type EIPAssociationInitParameters struct {

	// The allocation ID. This is required for EC2-VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EIP
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	AllocationIDRef *v1.Reference `json:"allocationIdRef,omitempty" tf:"-"`

	AllocationIDSelector *v1.Selector `json:"allocationIdSelector,omitempty" tf:"-"`

	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to true in VPC.
	AllowReassociation *bool `json:"allowReassociation,omitempty" tf:"allow_reassociation,omitempty"`

	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The Elastic IP address. This is required for EC2-Classic.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type EIPAssociationObservation struct {

	// The allocation ID. This is required for EC2-VPC.
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to true in VPC.
	AllowReassociation *bool `json:"allowReassociation,omitempty" tf:"allow_reassociation,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The Elastic IP address. This is required for EC2-Classic.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`
}

type EIPAssociationParameters struct {

	// The allocation ID. This is required for EC2-VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EIP
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	AllocationID *string `json:"allocationId,omitempty" tf:"allocation_id,omitempty"`

	// Reference to a EIP in ec2 to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDRef *v1.Reference `json:"allocationIdRef,omitempty" tf:"-"`

	// Selector for a EIP in ec2 to populate allocationId.
	// +kubebuilder:validation:Optional
	AllocationIDSelector *v1.Selector `json:"allocationIdSelector,omitempty" tf:"-"`

	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to true in VPC.
	AllowReassociation *bool `json:"allowReassociation,omitempty" tf:"allow_reassociation,omitempty"`

	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ec2 to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The Elastic IP address. This is required for EC2-Classic.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// EIPAssociationSpec defines the desired state of EIPAssociation
type EIPAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EIPAssociationInitParameters `json:"initProvider,omitempty"`
}

// EIPAssociationStatus defines the observed state of EIPAssociation.
type EIPAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EIPAssociation is the Schema for the EIPAssociations API. Provides an AWS EIP Association
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EIPAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EIPAssociationSpec   `json:"spec"`
	Status            EIPAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPAssociationList contains a list of EIPAssociations
type EIPAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIPAssociation `json:"items"`
}

// Repository type metadata.
var (
	EIPAssociation_Kind             = "EIPAssociation"
	EIPAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIPAssociation_Kind}.String()
	EIPAssociation_KindAPIVersion   = EIPAssociation_Kind + "." + CRDGroupVersion.String()
	EIPAssociation_GroupVersionKind = CRDGroupVersion.WithKind(EIPAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&EIPAssociation{}, &EIPAssociationList{})
}
