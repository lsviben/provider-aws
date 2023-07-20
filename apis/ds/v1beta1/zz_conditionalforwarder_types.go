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

type ConditionalForwarderInitParameters struct {

	// A list of forwarder IP addresses.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// ID of directory.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The fully qualified domain name of the remote domain for which forwarders will be used.
	RemoteDomainName *string `json:"remoteDomainName,omitempty" tf:"remote_domain_name,omitempty"`
}

type ConditionalForwarderObservation struct {

	// A list of forwarder IP addresses.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// ID of directory.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fully qualified domain name of the remote domain for which forwarders will be used.
	RemoteDomainName *string `json:"remoteDomainName,omitempty" tf:"remote_domain_name,omitempty"`
}

type ConditionalForwarderParameters struct {

	// A list of forwarder IP addresses.
	DNSIps []*string `json:"dnsIps,omitempty" tf:"dns_ips,omitempty"`

	// ID of directory.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ds/v1beta1.Directory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDRef *v1.Reference `json:"directoryIdRef,omitempty" tf:"-"`

	// Selector for a Directory in ds to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryIDSelector *v1.Selector `json:"directoryIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The fully qualified domain name of the remote domain for which forwarders will be used.
	RemoteDomainName *string `json:"remoteDomainName,omitempty" tf:"remote_domain_name,omitempty"`
}

// ConditionalForwarderSpec defines the desired state of ConditionalForwarder
type ConditionalForwarderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConditionalForwarderParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ConditionalForwarderInitParameters `json:"initProvider,omitempty"`
}

// ConditionalForwarderStatus defines the observed state of ConditionalForwarder.
type ConditionalForwarderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConditionalForwarderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionalForwarder is the Schema for the ConditionalForwarders API. Provides a conditional forwarder for managed Microsoft AD in AWS Directory Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConditionalForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dnsIps) || has(self.initProvider.dnsIps)",message="%!s(MISSING) is a required parameter"
	Spec   ConditionalForwarderSpec   `json:"spec"`
	Status ConditionalForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionalForwarderList contains a list of ConditionalForwarders
type ConditionalForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConditionalForwarder `json:"items"`
}

// Repository type metadata.
var (
	ConditionalForwarder_Kind             = "ConditionalForwarder"
	ConditionalForwarder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConditionalForwarder_Kind}.String()
	ConditionalForwarder_KindAPIVersion   = ConditionalForwarder_Kind + "." + CRDGroupVersion.String()
	ConditionalForwarder_GroupVersionKind = CRDGroupVersion.WithKind(ConditionalForwarder_Kind)
)

func init() {
	SchemeBuilder.Register(&ConditionalForwarder{}, &ConditionalForwarderList{})
}
