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

type AccountPasswordPolicyInitParameters struct {

	// Whether to allow users to change their own password
	AllowUsersToChangePassword *bool `json:"allowUsersToChangePassword,omitempty" tf:"allow_users_to_change_password,omitempty"`

	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry *bool `json:"hardExpiry,omitempty" tf:"hard_expiry,omitempty"`

	// The number of days that an user password is valid.
	MaxPasswordAge *float64 `json:"maxPasswordAge,omitempty" tf:"max_password_age,omitempty"`

	// Minimum length to require for user passwords.
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *float64 `json:"passwordReusePrevention,omitempty" tf:"password_reuse_prevention,omitempty"`

	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `json:"requireLowercaseCharacters,omitempty" tf:"require_lowercase_characters,omitempty"`

	// Whether to require numbers for user passwords.
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`

	// Whether to require symbols for user passwords.
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`

	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `json:"requireUppercaseCharacters,omitempty" tf:"require_uppercase_characters,omitempty"`
}

type AccountPasswordPolicyObservation struct {

	// Whether to allow users to change their own password
	AllowUsersToChangePassword *bool `json:"allowUsersToChangePassword,omitempty" tf:"allow_users_to_change_password,omitempty"`

	// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
	ExpirePasswords *bool `json:"expirePasswords,omitempty" tf:"expire_passwords,omitempty"`

	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry *bool `json:"hardExpiry,omitempty" tf:"hard_expiry,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The number of days that an user password is valid.
	MaxPasswordAge *float64 `json:"maxPasswordAge,omitempty" tf:"max_password_age,omitempty"`

	// Minimum length to require for user passwords.
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *float64 `json:"passwordReusePrevention,omitempty" tf:"password_reuse_prevention,omitempty"`

	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `json:"requireLowercaseCharacters,omitempty" tf:"require_lowercase_characters,omitempty"`

	// Whether to require numbers for user passwords.
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`

	// Whether to require symbols for user passwords.
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`

	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `json:"requireUppercaseCharacters,omitempty" tf:"require_uppercase_characters,omitempty"`
}

type AccountPasswordPolicyParameters struct {

	// Whether to allow users to change their own password
	AllowUsersToChangePassword *bool `json:"allowUsersToChangePassword,omitempty" tf:"allow_users_to_change_password,omitempty"`

	// Whether users are prevented from setting a new password after their password has expired (i.e., require administrator reset)
	HardExpiry *bool `json:"hardExpiry,omitempty" tf:"hard_expiry,omitempty"`

	// The number of days that an user password is valid.
	MaxPasswordAge *float64 `json:"maxPasswordAge,omitempty" tf:"max_password_age,omitempty"`

	// Minimum length to require for user passwords.
	MinimumPasswordLength *float64 `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`

	// The number of previous passwords that users are prevented from reusing.
	PasswordReusePrevention *float64 `json:"passwordReusePrevention,omitempty" tf:"password_reuse_prevention,omitempty"`

	// Whether to require lowercase characters for user passwords.
	RequireLowercaseCharacters *bool `json:"requireLowercaseCharacters,omitempty" tf:"require_lowercase_characters,omitempty"`

	// Whether to require numbers for user passwords.
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`

	// Whether to require symbols for user passwords.
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`

	// Whether to require uppercase characters for user passwords.
	RequireUppercaseCharacters *bool `json:"requireUppercaseCharacters,omitempty" tf:"require_uppercase_characters,omitempty"`
}

// AccountPasswordPolicySpec defines the desired state of AccountPasswordPolicy
type AccountPasswordPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountPasswordPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AccountPasswordPolicyInitParameters `json:"initProvider,omitempty"`
}

// AccountPasswordPolicyStatus defines the observed state of AccountPasswordPolicy.
type AccountPasswordPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountPasswordPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountPasswordPolicy is the Schema for the AccountPasswordPolicys API. Manages Password Policy for the AWS Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccountPasswordPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountPasswordPolicySpec   `json:"spec"`
	Status            AccountPasswordPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountPasswordPolicyList contains a list of AccountPasswordPolicys
type AccountPasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountPasswordPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccountPasswordPolicy_Kind             = "AccountPasswordPolicy"
	AccountPasswordPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccountPasswordPolicy_Kind}.String()
	AccountPasswordPolicy_KindAPIVersion   = AccountPasswordPolicy_Kind + "." + CRDGroupVersion.String()
	AccountPasswordPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccountPasswordPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccountPasswordPolicy{}, &AccountPasswordPolicyList{})
}
