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

type DiscovererInitParameters struct {

	// The description of the discoverer. Maximum of 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN of the event bus to discover event schemas on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Bus
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`

	SourceArnRef *v1.Reference `json:"sourceArnRef,omitempty" tf:"-"`

	SourceArnSelector *v1.Selector `json:"sourceArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DiscovererObservation struct {

	// The Amazon Resource Name (ARN) of the discoverer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the discoverer. Maximum of 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the discoverer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the event bus to discover event schemas on.
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DiscovererParameters struct {

	// The description of the discoverer. Maximum of 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN of the event bus to discover event schemas on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Bus
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`

	// Reference to a Bus in cloudwatchevents to populate sourceArn.
	// +kubebuilder:validation:Optional
	SourceArnRef *v1.Reference `json:"sourceArnRef,omitempty" tf:"-"`

	// Selector for a Bus in cloudwatchevents to populate sourceArn.
	// +kubebuilder:validation:Optional
	SourceArnSelector *v1.Selector `json:"sourceArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DiscovererSpec defines the desired state of Discoverer
type DiscovererSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiscovererParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DiscovererInitParameters `json:"initProvider,omitempty"`
}

// DiscovererStatus defines the observed state of Discoverer.
type DiscovererStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiscovererObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Discoverer is the Schema for the Discoverers API. Provides an EventBridge Schema Discoverer resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Discoverer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiscovererSpec   `json:"spec"`
	Status            DiscovererStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiscovererList contains a list of Discoverers
type DiscovererList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Discoverer `json:"items"`
}

// Repository type metadata.
var (
	Discoverer_Kind             = "Discoverer"
	Discoverer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Discoverer_Kind}.String()
	Discoverer_KindAPIVersion   = Discoverer_Kind + "." + CRDGroupVersion.String()
	Discoverer_GroupVersionKind = CRDGroupVersion.WithKind(Discoverer_Kind)
)

func init() {
	SchemeBuilder.Register(&Discoverer{}, &DiscovererList{})
}
