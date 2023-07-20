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

type DomainIdentityInitParameters struct {

	// The domain name to assign to SES
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type DomainIdentityObservation struct {

	// The ARN of the domain identity.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The domain name to assign to SES
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A code which when added to the domain as a TXT record
	// will signal to SES that the owner of the domain has authorised SES to act on
	// their behalf. The domain identity will be in state "verification pending"
	// until this is done.  Find out more about verifying domains in Amazon
	// SES in the AWS SES
	// docs.
	VerificationToken *string `json:"verificationToken,omitempty" tf:"verification_token,omitempty"`
}

type DomainIdentityParameters struct {

	// The domain name to assign to SES
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// DomainIdentitySpec defines the desired state of DomainIdentity
type DomainIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainIdentityParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DomainIdentityInitParameters `json:"initProvider,omitempty"`
}

// DomainIdentityStatus defines the observed state of DomainIdentity.
type DomainIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentity is the Schema for the DomainIdentitys API. Provides an SES domain identity resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || has(self.initProvider.domain)",message="%!s(MISSING) is a required parameter"
	Spec   DomainIdentitySpec   `json:"spec"`
	Status DomainIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainIdentityList contains a list of DomainIdentitys
type DomainIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainIdentity `json:"items"`
}

// Repository type metadata.
var (
	DomainIdentity_Kind             = "DomainIdentity"
	DomainIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainIdentity_Kind}.String()
	DomainIdentity_KindAPIVersion   = DomainIdentity_Kind + "." + CRDGroupVersion.String()
	DomainIdentity_GroupVersionKind = CRDGroupVersion.WithKind(DomainIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainIdentity{}, &DomainIdentityList{})
}
