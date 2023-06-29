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

type LagObservation struct {

	// The ARN of the LAG.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of an existing dedicated connection to migrate to the LAG.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	ConnectionsBandwidth *string `json:"connectionsBandwidth,omitempty" tf:"connections_bandwidth,omitempty"`

	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are not recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`

	// The ID of the LAG.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`

	// The AWS Direct Connect location in which the LAG should be allocated. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the LAG.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the AWS account that owns the LAG.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The name of the service provider associated with the LAG.
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LagParameters struct {

	// The ID of an existing dedicated connection to migrate to the LAG.
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps and 100Gbps. Case sensitive.
	// +kubebuilder:validation:Optional
	ConnectionsBandwidth *string `json:"connectionsBandwidth,omitempty" tf:"connections_bandwidth,omitempty"`

	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are not recoverable.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// The AWS Direct Connect location in which the LAG should be allocated. See DescribeLocations for the list of AWS Direct Connect locations. Use locationCode.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the LAG.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the service provider associated with the LAG.
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LagSpec defines the desired state of Lag
type LagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LagParameters `json:"forProvider"`
}

// LagStatus defines the observed state of Lag.
type LagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Lag is the Schema for the Lags API. Provides a Direct Connect LAG.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Lag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionsBandwidth)",message="connectionsBandwidth is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   LagSpec   `json:"spec"`
	Status LagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LagList contains a list of Lags
type LagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lag `json:"items"`
}

// Repository type metadata.
var (
	Lag_Kind             = "Lag"
	Lag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Lag_Kind}.String()
	Lag_KindAPIVersion   = Lag_Kind + "." + CRDGroupVersion.String()
	Lag_GroupVersionKind = CRDGroupVersion.WithKind(Lag_Kind)
)

func init() {
	SchemeBuilder.Register(&Lag{}, &LagList{})
}
