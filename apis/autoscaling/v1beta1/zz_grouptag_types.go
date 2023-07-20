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

type GroupTagInitParameters struct {

	// Name of the Autoscaling Group to apply the tag to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Tag to create. The tag block is documented below.
	Tag []GroupTagTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type GroupTagObservation struct {

	// Name of the Autoscaling Group to apply the tag to.
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// ASG name and key, separated by a comma (,)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Tag to create. The tag block is documented below.
	Tag []GroupTagTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type GroupTagParameters struct {

	// Name of the Autoscaling Group to apply the tag to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/autoscaling/v1beta1.AutoscalingGroup
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup in autoscaling to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Tag to create. The tag block is documented below.
	Tag []GroupTagTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type GroupTagTagInitParameters struct {

	// Tag name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether to propagate the tags to instances launched by the ASG.
	PropagateAtLaunch *bool `json:"propagateAtLaunch,omitempty" tf:"propagate_at_launch,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GroupTagTagObservation struct {

	// Tag name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether to propagate the tags to instances launched by the ASG.
	PropagateAtLaunch *bool `json:"propagateAtLaunch,omitempty" tf:"propagate_at_launch,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GroupTagTagParameters struct {

	// Tag name.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Whether to propagate the tags to instances launched by the ASG.
	PropagateAtLaunch *bool `json:"propagateAtLaunch,omitempty" tf:"propagate_at_launch,omitempty"`

	// Tag value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// GroupTagSpec defines the desired state of GroupTag
type GroupTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupTagParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider GroupTagInitParameters `json:"initProvider,omitempty"`
}

// GroupTagStatus defines the observed state of GroupTag.
type GroupTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupTag is the Schema for the GroupTags API. Manages an individual Autoscaling Group tag
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GroupTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tag) || has(self.initProvider.tag)",message="%!s(MISSING) is a required parameter"
	Spec   GroupTagSpec   `json:"spec"`
	Status GroupTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupTagList contains a list of GroupTags
type GroupTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupTag `json:"items"`
}

// Repository type metadata.
var (
	GroupTag_Kind             = "GroupTag"
	GroupTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupTag_Kind}.String()
	GroupTag_KindAPIVersion   = GroupTag_Kind + "." + CRDGroupVersion.String()
	GroupTag_GroupVersionKind = CRDGroupVersion.WithKind(GroupTag_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupTag{}, &GroupTagList{})
}
