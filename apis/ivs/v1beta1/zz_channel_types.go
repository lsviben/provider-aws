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

type ChannelInitParameters struct {

	// If true, channel is private (enabled for playback authorization).
	Authorized *bool `json:"authorized,omitempty" tf:"authorized,omitempty"`

	// Channel latency mode. Valid values: NORMAL, LOW.
	LatencyMode *string `json:"latencyMode,omitempty" tf:"latency_mode,omitempty"`

	// Channel name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Recording configuration ARN.
	RecordingConfigurationArn *string `json:"recordingConfigurationArn,omitempty" tf:"recording_configuration_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Channel type, which determines the allowable resolution and bitrate. Valid values: STANDARD, BASIC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ChannelObservation struct {

	// ARN of the Channel.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// If true, channel is private (enabled for playback authorization).
	Authorized *bool `json:"authorized,omitempty" tf:"authorized,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Channel ingest endpoint, part of the definition of an ingest server, used when setting up streaming software.
	IngestEndpoint *string `json:"ingestEndpoint,omitempty" tf:"ingest_endpoint,omitempty"`

	// Channel latency mode. Valid values: NORMAL, LOW.
	LatencyMode *string `json:"latencyMode,omitempty" tf:"latency_mode,omitempty"`

	// Channel name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Channel playback URL.
	PlaybackURL *string `json:"playbackUrl,omitempty" tf:"playback_url,omitempty"`

	// Recording configuration ARN.
	RecordingConfigurationArn *string `json:"recordingConfigurationArn,omitempty" tf:"recording_configuration_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Channel type, which determines the allowable resolution and bitrate. Valid values: STANDARD, BASIC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ChannelParameters struct {

	// If true, channel is private (enabled for playback authorization).
	Authorized *bool `json:"authorized,omitempty" tf:"authorized,omitempty"`

	// Channel latency mode. Valid values: NORMAL, LOW.
	LatencyMode *string `json:"latencyMode,omitempty" tf:"latency_mode,omitempty"`

	// Channel name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Recording configuration ARN.
	RecordingConfigurationArn *string `json:"recordingConfigurationArn,omitempty" tf:"recording_configuration_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Channel type, which determines the allowable resolution and bitrate. Valid values: STANDARD, BASIC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ChannelSpec defines the desired state of Channel
type ChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ChannelParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ChannelInitParameters `json:"initProvider,omitempty"`
}

// ChannelStatus defines the observed state of Channel.
type ChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Channel is the Schema for the Channels API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Channel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelSpec   `json:"spec"`
	Status            ChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChannelList contains a list of Channels
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Channel `json:"items"`
}

// Repository type metadata.
var (
	Channel_Kind             = "Channel"
	Channel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Channel_Kind}.String()
	Channel_KindAPIVersion   = Channel_Kind + "." + CRDGroupVersion.String()
	Channel_GroupVersionKind = CRDGroupVersion.WithKind(Channel_Kind)
)

func init() {
	SchemeBuilder.Register(&Channel{}, &ChannelList{})
}
