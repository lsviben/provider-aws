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

type PublicVirtualInterfaceInitParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	// +crossplane:generate:reference:type=Connection
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VLAN ID.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type PublicVirtualInterfaceObservation struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	AmazonSideAsn *string `json:"amazonSideAsn,omitempty" tf:"amazon_side_asn,omitempty"`

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The VLAN ID.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type PublicVirtualInterfaceParameters struct {

	// The address family for the BGP peer. ipv4  or ipv6.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `json:"amazonAddress,omitempty" tf:"amazon_address,omitempty"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// The authentication key for BGP configuration.
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	// +crossplane:generate:reference:type=Connection
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a Connection to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// The name for the virtual interface.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VLAN ID.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// PublicVirtualInterfaceSpec defines the desired state of PublicVirtualInterface
type PublicVirtualInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicVirtualInterfaceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider PublicVirtualInterfaceInitParameters `json:"initProvider,omitempty"`
}

// PublicVirtualInterfaceStatus defines the observed state of PublicVirtualInterface.
type PublicVirtualInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicVirtualInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicVirtualInterface is the Schema for the PublicVirtualInterfaces API. Provides a Direct Connect public virtual interface resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PublicVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || has(self.initProvider.addressFamily)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bgpAsn) || has(self.initProvider.bgpAsn)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routeFilterPrefixes) || has(self.initProvider.routeFilterPrefixes)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vlan) || has(self.initProvider.vlan)",message="%!s(MISSING) is a required parameter"
	Spec   PublicVirtualInterfaceSpec   `json:"spec"`
	Status PublicVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicVirtualInterfaceList contains a list of PublicVirtualInterfaces
type PublicVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicVirtualInterface `json:"items"`
}

// Repository type metadata.
var (
	PublicVirtualInterface_Kind             = "PublicVirtualInterface"
	PublicVirtualInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicVirtualInterface_Kind}.String()
	PublicVirtualInterface_KindAPIVersion   = PublicVirtualInterface_Kind + "." + CRDGroupVersion.String()
	PublicVirtualInterface_GroupVersionKind = CRDGroupVersion.WithKind(PublicVirtualInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicVirtualInterface{}, &PublicVirtualInterfaceList{})
}
