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

type ObjectLambdaAccessPointPolicyInitParameters struct {

	// The AWS account ID for the account that owns the Object Lambda Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The Object Lambda Access Point resource policy document.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ObjectLambdaAccessPointPolicyObservation struct {

	// The AWS account ID for the account that owns the Object Lambda Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy *bool `json:"hasPublicAccessPolicy,omitempty" tf:"has_public_access_policy,omitempty"`

	// The AWS account ID and access point name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Object Lambda Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Object Lambda Access Point resource policy document.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type ObjectLambdaAccessPointPolicyParameters struct {

	// The AWS account ID for the account that owns the Object Lambda Access Point.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The name of the Object Lambda Access Point.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3control/v1beta1.ObjectLambdaAccessPoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a ObjectLambdaAccessPoint in s3control to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a ObjectLambdaAccessPoint in s3control to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The Object Lambda Access Point resource policy document.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ObjectLambdaAccessPointPolicySpec defines the desired state of ObjectLambdaAccessPointPolicy
type ObjectLambdaAccessPointPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectLambdaAccessPointPolicyParameters `json:"forProvider"`
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
	InitProvider ObjectLambdaAccessPointPolicyInitParameters `json:"initProvider,omitempty"`
}

// ObjectLambdaAccessPointPolicyStatus defines the observed state of ObjectLambdaAccessPointPolicy.
type ObjectLambdaAccessPointPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectLambdaAccessPointPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectLambdaAccessPointPolicy is the Schema for the ObjectLambdaAccessPointPolicys API. Provides a resource to manage an S3 Object Lambda Access Point resource policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ObjectLambdaAccessPointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   ObjectLambdaAccessPointPolicySpec   `json:"spec"`
	Status ObjectLambdaAccessPointPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectLambdaAccessPointPolicyList contains a list of ObjectLambdaAccessPointPolicys
type ObjectLambdaAccessPointPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectLambdaAccessPointPolicy `json:"items"`
}

// Repository type metadata.
var (
	ObjectLambdaAccessPointPolicy_Kind             = "ObjectLambdaAccessPointPolicy"
	ObjectLambdaAccessPointPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectLambdaAccessPointPolicy_Kind}.String()
	ObjectLambdaAccessPointPolicy_KindAPIVersion   = ObjectLambdaAccessPointPolicy_Kind + "." + CRDGroupVersion.String()
	ObjectLambdaAccessPointPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ObjectLambdaAccessPointPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectLambdaAccessPointPolicy{}, &ObjectLambdaAccessPointPolicyList{})
}
