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

type UserInitParameters struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// The ARN of the created ElastiCache User.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Passwords used for this user. You can create up to two passwords for each user.
	PasswordsSecretRef *[]v1.SecretKeySelector `json:"passwordsSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The username of the user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserObservation struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// The ARN of the created ElastiCache User.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The username of the user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type UserParameters struct {

	// Access permissions string used for this user. See Specifying Permissions Using an Access String for more details.
	AccessString *string `json:"accessString,omitempty" tf:"access_string,omitempty"`

	// The ARN of the created ElastiCache User.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `json:"noPasswordRequired,omitempty" tf:"no_password_required,omitempty"`

	// Passwords used for this user. You can create up to two passwords for each user.
	PasswordsSecretRef *[]v1.SecretKeySelector `json:"passwordsSecretRef,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The username of the user.
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API. Provides an ElastiCache user.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessString) || has(self.initProvider.accessString)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || has(self.initProvider.engine)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || has(self.initProvider.userName)",message="%!s(MISSING) is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
