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

type SMSChannelInitParameters struct {

	// Whether the channel is enabled or disabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Sender identifier of your messages.
	SenderID *string `json:"senderId,omitempty" tf:"sender_id,omitempty"`

	// The Short Code registered with the phone provider.
	ShortCode *string `json:"shortCode,omitempty" tf:"short_code,omitempty"`
}

type SMSChannelObservation struct {

	// The application ID.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Whether the channel is enabled or disabled. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond *float64 `json:"promotionalMessagesPerSecond,omitempty" tf:"promotional_messages_per_second,omitempty"`

	// Sender identifier of your messages.
	SenderID *string `json:"senderId,omitempty" tf:"sender_id,omitempty"`

	// The Short Code registered with the phone provider.
	ShortCode *string `json:"shortCode,omitempty" tf:"short_code,omitempty"`

	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond *float64 `json:"transactionalMessagesPerSecond,omitempty" tf:"transactional_messages_per_second,omitempty"`
}

type SMSChannelParameters struct {

	// The application ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/pinpoint/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("application_id",true)
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a App in pinpoint to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a App in pinpoint to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Whether the channel is enabled or disabled. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Sender identifier of your messages.
	// +kubebuilder:validation:Optional
	SenderID *string `json:"senderId,omitempty" tf:"sender_id,omitempty"`

	// The Short Code registered with the phone provider.
	// +kubebuilder:validation:Optional
	ShortCode *string `json:"shortCode,omitempty" tf:"short_code,omitempty"`
}

// SMSChannelSpec defines the desired state of SMSChannel
type SMSChannelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SMSChannelParameters `json:"forProvider"`
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
	InitProvider SMSChannelInitParameters `json:"initProvider,omitempty"`
}

// SMSChannelStatus defines the observed state of SMSChannel.
type SMSChannelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SMSChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SMSChannel is the Schema for the SMSChannels API. Provides a Pinpoint SMS Channel resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SMSChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SMSChannelSpec   `json:"spec"`
	Status            SMSChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SMSChannelList contains a list of SMSChannels
type SMSChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SMSChannel `json:"items"`
}

// Repository type metadata.
var (
	SMSChannel_Kind             = "SMSChannel"
	SMSChannel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SMSChannel_Kind}.String()
	SMSChannel_KindAPIVersion   = SMSChannel_Kind + "." + CRDGroupVersion.String()
	SMSChannel_GroupVersionKind = CRDGroupVersion.WithKind(SMSChannel_Kind)
)

func init() {
	SchemeBuilder.Register(&SMSChannel{}, &SMSChannelList{})
}
