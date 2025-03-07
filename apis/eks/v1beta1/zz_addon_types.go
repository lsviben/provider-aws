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

type AddonInitParameters struct {

	// on. The name must match one of
	// the names returned by describe-addon-versions.
	AddonName *string `json:"addonName,omitempty" tf:"addon_name,omitempty"`

	// on. The version must
	// match one of the versions returned by describe-addon-versions.
	AddonVersion *string `json:"addonVersion,omitempty" tf:"addon_version,omitempty"`

	// custom configuration values for addons with single JSON string. This JSON string value must match the JSON schema derived from describe-addon-configuration.
	ConfigurationValues *string `json:"configurationValues,omitempty" tf:"configuration_values,omitempty"`

	// Indicates if you want to preserve the created resources when deleting the EKS add-on.
	Preserve *bool `json:"preserve,omitempty" tf:"preserve,omitempty"`

	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are NONE, OVERWRITE and PRESERVE. For more details check UpdateAddon API Docs.
	ResolveConflicts *string `json:"resolveConflicts,omitempty" tf:"resolve_conflicts,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AddonObservation struct {

	// on. The name must match one of
	// the names returned by describe-addon-versions.
	AddonName *string `json:"addonName,omitempty" tf:"addon_name,omitempty"`

	// on. The version must
	// match one of the versions returned by describe-addon-versions.
	AddonVersion *string `json:"addonVersion,omitempty" tf:"addon_version,omitempty"`

	// Amazon Resource Name (ARN) of the EKS add-on.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// 100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (^[0-9A-Za-z][A-Za-z0-9\-_]+$).
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// custom configuration values for addons with single JSON string. This JSON string value must match the JSON schema derived from describe-addon-configuration.
	ConfigurationValues *string `json:"configurationValues,omitempty" tf:"configuration_values,omitempty"`

	// Date and time in RFC3339 format that the EKS add-on was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// EKS Cluster name and EKS Addon name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date and time in RFC3339 format that the EKS add-on was updated.
	ModifiedAt *string `json:"modifiedAt,omitempty" tf:"modified_at,omitempty"`

	// Indicates if you want to preserve the created resources when deleting the EKS add-on.
	Preserve *bool `json:"preserve,omitempty" tf:"preserve,omitempty"`

	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are NONE, OVERWRITE and PRESERVE. For more details check UpdateAddon API Docs.
	ResolveConflicts *string `json:"resolveConflicts,omitempty" tf:"resolve_conflicts,omitempty"`

	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see Amazon EKS node IAM role
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn *string `json:"serviceAccountRoleArn,omitempty" tf:"service_account_role_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Key-value map of resource tags, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AddonParameters struct {

	// on. The name must match one of
	// the names returned by describe-addon-versions.
	// +kubebuilder:validation:Optional
	AddonName *string `json:"addonName,omitempty" tf:"addon_name,omitempty"`

	// on. The version must
	// match one of the versions returned by describe-addon-versions.
	// +kubebuilder:validation:Optional
	AddonVersion *string `json:"addonVersion,omitempty" tf:"addon_version,omitempty"`

	// 100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (^[0-9A-Za-z][A-Za-z0-9\-_]+$).
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// custom configuration values for addons with single JSON string. This JSON string value must match the JSON schema derived from describe-addon-configuration.
	// +kubebuilder:validation:Optional
	ConfigurationValues *string `json:"configurationValues,omitempty" tf:"configuration_values,omitempty"`

	// Indicates if you want to preserve the created resources when deleting the EKS add-on.
	// +kubebuilder:validation:Optional
	Preserve *bool `json:"preserve,omitempty" tf:"preserve,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are NONE, OVERWRITE and PRESERVE. For more details check UpdateAddon API Docs.
	// +kubebuilder:validation:Optional
	ResolveConflicts *string `json:"resolveConflicts,omitempty" tf:"resolve_conflicts,omitempty"`

	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see Amazon EKS node IAM role
	// in the Amazon EKS User Guide.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceAccountRoleArn *string `json:"serviceAccountRoleArn,omitempty" tf:"service_account_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceAccountRoleArn.
	// +kubebuilder:validation:Optional
	ServiceAccountRoleArnRef *v1.Reference `json:"serviceAccountRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceAccountRoleArn.
	// +kubebuilder:validation:Optional
	ServiceAccountRoleArnSelector *v1.Selector `json:"serviceAccountRoleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AddonSpec defines the desired state of Addon
type AddonSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddonParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AddonInitParameters `json:"initProvider,omitempty"`
}

// AddonStatus defines the observed state of Addon.
type AddonStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddonObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Addon is the Schema for the Addons API. Manages an EKS add-on
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Addon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addonName) || (has(self.initProvider) && has(self.initProvider.addonName))",message="spec.forProvider.addonName is a required parameter"
	Spec   AddonSpec   `json:"spec"`
	Status AddonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddonList contains a list of Addons
type AddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Addon `json:"items"`
}

// Repository type metadata.
var (
	Addon_Kind             = "Addon"
	Addon_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Addon_Kind}.String()
	Addon_KindAPIVersion   = Addon_Kind + "." + CRDGroupVersion.String()
	Addon_GroupVersionKind = CRDGroupVersion.WithKind(Addon_Kind)
)

func init() {
	SchemeBuilder.Register(&Addon{}, &AddonList{})
}
