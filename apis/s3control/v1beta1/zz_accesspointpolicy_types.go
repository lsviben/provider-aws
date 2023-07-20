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

type AccessPointPolicyInitParameters struct {

	// The ARN of the access point that you want to associate with the specified policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3control/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	AccessPointArn *string `json:"accessPointArn,omitempty" tf:"access_point_arn,omitempty"`

	AccessPointArnRef *v1.Reference `json:"accessPointArnRef,omitempty" tf:"-"`

	AccessPointArnSelector *v1.Selector `json:"accessPointArnSelector,omitempty" tf:"-"`

	// The policy that you want to apply to the specified access point.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type AccessPointPolicyObservation struct {

	// The ARN of the access point that you want to associate with the specified policy.
	AccessPointArn *string `json:"accessPointArn,omitempty" tf:"access_point_arn,omitempty"`

	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy *bool `json:"hasPublicAccessPolicy,omitempty" tf:"has_public_access_policy,omitempty"`

	// The AWS account ID and access point name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy that you want to apply to the specified access point.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type AccessPointPolicyParameters struct {

	// The ARN of the access point that you want to associate with the specified policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3control/v1beta1.AccessPoint
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	AccessPointArn *string `json:"accessPointArn,omitempty" tf:"access_point_arn,omitempty"`

	// Reference to a AccessPoint in s3control to populate accessPointArn.
	// +kubebuilder:validation:Optional
	AccessPointArnRef *v1.Reference `json:"accessPointArnRef,omitempty" tf:"-"`

	// Selector for a AccessPoint in s3control to populate accessPointArn.
	// +kubebuilder:validation:Optional
	AccessPointArnSelector *v1.Selector `json:"accessPointArnSelector,omitempty" tf:"-"`

	// The policy that you want to apply to the specified access point.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// AccessPointPolicySpec defines the desired state of AccessPointPolicy
type AccessPointPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPointPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AccessPointPolicyInitParameters `json:"initProvider,omitempty"`
}

// AccessPointPolicyStatus defines the observed state of AccessPointPolicy.
type AccessPointPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPointPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointPolicy is the Schema for the AccessPointPolicys API. Provides a resource to manage an S3 Access Point resource policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessPointPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || has(self.initProvider.policy)",message="%!s(MISSING) is a required parameter"
	Spec   AccessPointPolicySpec   `json:"spec"`
	Status AccessPointPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointPolicyList contains a list of AccessPointPolicys
type AccessPointPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPointPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccessPointPolicy_Kind             = "AccessPointPolicy"
	AccessPointPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPointPolicy_Kind}.String()
	AccessPointPolicy_KindAPIVersion   = AccessPointPolicy_Kind + "." + CRDGroupVersion.String()
	AccessPointPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccessPointPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPointPolicy{}, &AccessPointPolicyList{})
}
