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

type EndpointAccessInitParameters struct {

	// The name of the workgroup.
	WorkgroupName *string `json:"workgroupName,omitempty" tf:"workgroup_name,omitempty"`
}

type EndpointAccessObservation struct {

	// The DNS address of the VPC endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Amazon Resource Name (ARN) of the Redshift Serverless Endpoint Access.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Redshift Endpoint Access Name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The port that Amazon Redshift Serverless listens on.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// An array of VPC subnet IDs to associate with the endpoint.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The VPC endpoint or the Redshift Serverless workgroup. See VPC Endpoint below.
	VPCEndpoint []VPCEndpointObservation `json:"vpcEndpoint,omitempty" tf:"vpc_endpoint,omitempty"`

	// An array of security group IDs to associate with the workgroup.
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// The name of the workgroup.
	WorkgroupName *string `json:"workgroupName,omitempty" tf:"workgroup_name,omitempty"`
}

type EndpointAccessParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// An array of VPC subnet IDs to associate with the endpoint.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// An array of security group IDs to associate with the workgroup.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// The name of the workgroup.
	// +kubebuilder:validation:Optional
	WorkgroupName *string `json:"workgroupName,omitempty" tf:"workgroup_name,omitempty"`
}

type NetworkInterfaceInitParameters struct {
}

type NetworkInterfaceObservation struct {

	// The availability Zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The unique identifier of the network interface.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// The IPv4 address of the network interface within the subnet.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The unique identifier of the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceParameters struct {
}

type VPCEndpointInitParameters struct {
}

type VPCEndpointObservation struct {

	// The network interfaces of the endpoint.. See Network Interface below.
	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The DNS address of the VPC endpoint.
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// The port that Amazon Redshift Serverless listens on.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCEndpointParameters struct {
}

// EndpointAccessSpec defines the desired state of EndpointAccess
type EndpointAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointAccessParameters `json:"forProvider"`
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
	InitProvider EndpointAccessInitParameters `json:"initProvider,omitempty"`
}

// EndpointAccessStatus defines the observed state of EndpointAccess.
type EndpointAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointAccess is the Schema for the EndpointAccesss API. Provides a Redshift Serverless Endpoint Access resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EndpointAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workgroupName) || (has(self.initProvider) && has(self.initProvider.workgroupName))",message="spec.forProvider.workgroupName is a required parameter"
	Spec   EndpointAccessSpec   `json:"spec"`
	Status EndpointAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointAccessList contains a list of EndpointAccesss
type EndpointAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointAccess `json:"items"`
}

// Repository type metadata.
var (
	EndpointAccess_Kind             = "EndpointAccess"
	EndpointAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointAccess_Kind}.String()
	EndpointAccess_KindAPIVersion   = EndpointAccess_Kind + "." + CRDGroupVersion.String()
	EndpointAccess_GroupVersionKind = CRDGroupVersion.WithKind(EndpointAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointAccess{}, &EndpointAccessList{})
}
