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

type NetworkInsightsPathObservation struct {

	// ARN of the Network Insights Path.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ID of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// IP address of the destination resource.
	DestinationIP *string `json:"destinationIp,omitempty" tf:"destination_ip,omitempty"`

	// Destination port to analyze access to.
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// ID of the Network Insights Path.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Protocol to use for analysis. Valid options are tcp or udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// ID of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// IP address of the source resource.
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NetworkInsightsPathParameters struct {

	// ID of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// IP address of the destination resource.
	// +kubebuilder:validation:Optional
	DestinationIP *string `json:"destinationIp,omitempty" tf:"destination_ip,omitempty"`

	// Destination port to analyze access to.
	// +kubebuilder:validation:Optional
	DestinationPort *float64 `json:"destinationPort,omitempty" tf:"destination_port,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate destination.
	// +kubebuilder:validation:Optional
	DestinationRef *v1.Reference `json:"destinationRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate destination.
	// +kubebuilder:validation:Optional
	DestinationSelector *v1.Selector `json:"destinationSelector,omitempty" tf:"-"`

	// Protocol to use for analysis. Valid options are tcp or udp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ID of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// IP address of the source resource.
	// +kubebuilder:validation:Optional
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// Reference to a NetworkInterface in ec2 to populate source.
	// +kubebuilder:validation:Optional
	SourceRef *v1.Reference `json:"sourceRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface in ec2 to populate source.
	// +kubebuilder:validation:Optional
	SourceSelector *v1.Selector `json:"sourceSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NetworkInsightsPathSpec defines the desired state of NetworkInsightsPath
type NetworkInsightsPathSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInsightsPathParameters `json:"forProvider"`
}

// NetworkInsightsPathStatus defines the observed state of NetworkInsightsPath.
type NetworkInsightsPathStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInsightsPathObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInsightsPath is the Schema for the NetworkInsightsPaths API. Provides a Network Insights Path resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NetworkInsightsPath struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol)",message="protocol is a required parameter"
	Spec   NetworkInsightsPathSpec   `json:"spec"`
	Status NetworkInsightsPathStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInsightsPathList contains a list of NetworkInsightsPaths
type NetworkInsightsPathList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInsightsPath `json:"items"`
}

// Repository type metadata.
var (
	NetworkInsightsPath_Kind             = "NetworkInsightsPath"
	NetworkInsightsPath_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInsightsPath_Kind}.String()
	NetworkInsightsPath_KindAPIVersion   = NetworkInsightsPath_Kind + "." + CRDGroupVersion.String()
	NetworkInsightsPath_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInsightsPath_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInsightsPath{}, &NetworkInsightsPathList{})
}
