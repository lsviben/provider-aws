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

type HostInitParameters struct {

	// The name of the host to be created. The name must be unique in the calling AWS account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The endpoint of the infrastructure to be represented by the host after it is created.
	ProviderEndpoint *string `json:"providerEndpoint,omitempty" tf:"provider_endpoint,omitempty"`

	// The name of the external provider where your third-party code repository is configured.
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
	VPCConfiguration []VPCConfigurationInitParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type HostObservation struct {

	// The CodeStar Host ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The CodeStar Host ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the host to be created. The name must be unique in the calling AWS account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The endpoint of the infrastructure to be represented by the host after it is created.
	ProviderEndpoint *string `json:"providerEndpoint,omitempty" tf:"provider_endpoint,omitempty"`

	// The name of the external provider where your third-party code repository is configured.
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// The CodeStar Host status. Possible values are PENDING, AVAILABLE, VPC_CONFIG_DELETING, VPC_CONFIG_INITIALIZING, and VPC_CONFIG_FAILED_INITIALIZATION.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
	VPCConfiguration []VPCConfigurationObservation `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type HostParameters struct {

	// The name of the host to be created. The name must be unique in the calling AWS account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The endpoint of the infrastructure to be represented by the host after it is created.
	ProviderEndpoint *string `json:"providerEndpoint,omitempty" tf:"provider_endpoint,omitempty"`

	// The name of the external provider where your third-party code repository is configured.
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type VPCConfigurationInitParameters struct {

	// he ID of the security group or security groups associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	TLSCertificate *string `json:"tlsCertificate,omitempty" tf:"tls_certificate,omitempty"`

	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigurationObservation struct {

	// he ID of the security group or security groups associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	TLSCertificate *string `json:"tlsCertificate,omitempty" tf:"tls_certificate,omitempty"`

	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigurationParameters struct {

	// he ID of the security group or security groups associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the infrastructure where your provider type is installed.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	TLSCertificate *string `json:"tlsCertificate,omitempty" tf:"tls_certificate,omitempty"`

	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider HostInitParameters `json:"initProvider,omitempty"`
}

// HostStatus defines the observed state of Host.
type HostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Host is the Schema for the Hosts API. Provides a CodeStar Host
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerEndpoint) || has(self.initProvider.providerEndpoint)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerType) || has(self.initProvider.providerType)",message="%!s(MISSING) is a required parameter"
	Spec   HostSpec   `json:"spec"`
	Status HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostList contains a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// Repository type metadata.
var (
	Host_Kind             = "Host"
	Host_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Host_Kind}.String()
	Host_KindAPIVersion   = Host_Kind + "." + CRDGroupVersion.String()
	Host_GroupVersionKind = CRDGroupVersion.WithKind(Host_Kind)
)

func init() {
	SchemeBuilder.Register(&Host{}, &HostList{})
}
