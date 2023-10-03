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

type APICacheInitParameters struct {

	// Caching behavior. Valid values are FULL_REQUEST_CACHING and PER_RESOLVER_CACHING.
	APICachingBehavior *string `json:"apiCachingBehavior,omitempty" tf:"api_caching_behavior,omitempty"`

	// At-rest encryption flag for cache. You cannot update this setting after creation.
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// TTL in seconds for cache entries.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`

	// Cache instance type. Valid values are SMALL, MEDIUM, LARGE, XLARGE, LARGE_2X, LARGE_4X, LARGE_8X, LARGE_12X, T2_SMALL, T2_MEDIUM, R4_LARGE, R4_XLARGE, R4_2XLARGE, R4_4XLARGE, R4_8XLARGE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type APICacheObservation struct {

	// Caching behavior. Valid values are FULL_REQUEST_CACHING and PER_RESOLVER_CACHING.
	APICachingBehavior *string `json:"apiCachingBehavior,omitempty" tf:"api_caching_behavior,omitempty"`

	// GraphQL API ID.
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// At-rest encryption flag for cache. You cannot update this setting after creation.
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// AppSync API ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TTL in seconds for cache entries.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`

	// Cache instance type. Valid values are SMALL, MEDIUM, LARGE, XLARGE, LARGE_2X, LARGE_4X, LARGE_8X, LARGE_12X, T2_SMALL, T2_MEDIUM, R4_LARGE, R4_XLARGE, R4_2XLARGE, R4_4XLARGE, R4_8XLARGE.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type APICacheParameters struct {

	// Caching behavior. Valid values are FULL_REQUEST_CACHING and PER_RESOLVER_CACHING.
	// +kubebuilder:validation:Optional
	APICachingBehavior *string `json:"apiCachingBehavior,omitempty" tf:"api_caching_behavior,omitempty"`

	// GraphQL API ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// At-rest encryption flag for cache. You cannot update this setting after creation.
	// +kubebuilder:validation:Optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// TTL in seconds for cache entries.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	// +kubebuilder:validation:Optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`

	// Cache instance type. Valid values are SMALL, MEDIUM, LARGE, XLARGE, LARGE_2X, LARGE_4X, LARGE_8X, LARGE_12X, T2_SMALL, T2_MEDIUM, R4_LARGE, R4_XLARGE, R4_2XLARGE, R4_4XLARGE, R4_8XLARGE.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// APICacheSpec defines the desired state of APICache
type APICacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APICacheParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It is not honored
	// unless the relevant Crossplane feature flag is enabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider APICacheInitParameters `json:"initProvider,omitempty"`
}

// APICacheStatus defines the observed state of APICache.
type APICacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APICacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APICache is the Schema for the APICaches API. Provides an AppSync API Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APICache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiCachingBehavior) || (has(self.initProvider) && has(self.initProvider.apiCachingBehavior))",message="spec.forProvider.apiCachingBehavior is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ttl) || (has(self.initProvider) && has(self.initProvider.ttl))",message="spec.forProvider.ttl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   APICacheSpec   `json:"spec"`
	Status APICacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APICacheList contains a list of APICaches
type APICacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APICache `json:"items"`
}

// Repository type metadata.
var (
	APICache_Kind             = "APICache"
	APICache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APICache_Kind}.String()
	APICache_KindAPIVersion   = APICache_Kind + "." + CRDGroupVersion.String()
	APICache_GroupVersionKind = CRDGroupVersion.WithKind(APICache_Kind)
)

func init() {
	SchemeBuilder.Register(&APICache{}, &APICacheList{})
}
