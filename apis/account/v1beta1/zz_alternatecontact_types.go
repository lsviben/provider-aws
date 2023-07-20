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

type AlternateContactInitParameters struct {

	// ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Type of the alternate contact. Allowed values are: BILLING, OPERATIONS, SECURITY.
	AlternateContactType *string `json:"alternateContactType,omitempty" tf:"alternate_contact_type,omitempty"`

	// An email address for the alternate contact.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// Name of the alternate contact.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Phone number for the alternate contact.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Title for the alternate contact.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AlternateContactObservation struct {

	// ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Type of the alternate contact. Allowed values are: BILLING, OPERATIONS, SECURITY.
	AlternateContactType *string `json:"alternateContactType,omitempty" tf:"alternate_contact_type,omitempty"`

	// An email address for the alternate contact.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the alternate contact.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Phone number for the alternate contact.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// Title for the alternate contact.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AlternateContactParameters struct {

	// ID of the target account when managing member accounts. Will manage current user's account by default if omitted.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Type of the alternate contact. Allowed values are: BILLING, OPERATIONS, SECURITY.
	AlternateContactType *string `json:"alternateContactType,omitempty" tf:"alternate_contact_type,omitempty"`

	// An email address for the alternate contact.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// Name of the alternate contact.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Phone number for the alternate contact.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Title for the alternate contact.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// AlternateContactSpec defines the desired state of AlternateContact
type AlternateContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlternateContactParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AlternateContactInitParameters `json:"initProvider,omitempty"`
}

// AlternateContactStatus defines the observed state of AlternateContact.
type AlternateContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlternateContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlternateContact is the Schema for the AlternateContacts API. Manages the specified alternate contact attached to an AWS Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AlternateContact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.emailAddress) || has(self.initProvider.emailAddress)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.phoneNumber) || has(self.initProvider.phoneNumber)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || has(self.initProvider.title)",message="%!s(MISSING) is a required parameter"
	Spec   AlternateContactSpec   `json:"spec"`
	Status AlternateContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlternateContactList contains a list of AlternateContacts
type AlternateContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlternateContact `json:"items"`
}

// Repository type metadata.
var (
	AlternateContact_Kind             = "AlternateContact"
	AlternateContact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlternateContact_Kind}.String()
	AlternateContact_KindAPIVersion   = AlternateContact_Kind + "." + CRDGroupVersion.String()
	AlternateContact_GroupVersionKind = CRDGroupVersion.WithKind(AlternateContact_Kind)
)

func init() {
	SchemeBuilder.Register(&AlternateContact{}, &AlternateContactList{})
}
