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

type EventSubscriptionInitParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run aws redshift describe-event-categories.
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The event severity to be published by the notification subscription. Valid options are INFO or ERROR. Default value of INFO.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// The ARN of the SNS topic to send events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are cluster, cluster-parameter-group, cluster-security-group, cluster-snapshot, or scheduled-action. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventSubscriptionObservation struct {

	// Amazon Resource Name (ARN) of the Redshift event notification subscription
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS customer account associated with the Redshift event notification subscription
	CustomerAwsID *string `json:"customerAwsId,omitempty" tf:"customer_aws_id,omitempty"`

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run aws redshift describe-event-categories.
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// The name of the Redshift event notification subscription
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The event severity to be published by the notification subscription. Valid options are INFO or ERROR. Default value of INFO.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// The ARN of the SNS topic to send events to.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are cluster, cluster-parameter-group, cluster-security-group, cluster-snapshot, or scheduled-action. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventSubscriptionParameters struct {

	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run aws redshift describe-event-categories.
	EventCategories []*string `json:"eventCategories,omitempty" tf:"event_categories,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The event severity to be published by the notification subscription. Valid options are INFO or ERROR. Default value of INFO.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// The ARN of the SNS topic to send events to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`

	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a source_type must also be specified.
	SourceIds []*string `json:"sourceIds,omitempty" tf:"source_ids,omitempty"`

	// The type of source that will be generating the events. Valid options are cluster, cluster-parameter-group, cluster-security-group, cluster-snapshot, or scheduled-action. If not set, all sources will be subscribed to.
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventSubscriptionSpec defines the desired state of EventSubscription
type EventSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSubscriptionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EventSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// EventSubscriptionStatus defines the observed state of EventSubscription.
type EventSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscription is the Schema for the EventSubscriptions API. Provides a Redshift event subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSubscriptionSpec   `json:"spec"`
	Status            EventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSubscriptionList contains a list of EventSubscriptions
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

// Repository type metadata.
var (
	EventSubscription_Kind             = "EventSubscription"
	EventSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSubscription_Kind}.String()
	EventSubscription_KindAPIVersion   = EventSubscription_Kind + "." + CRDGroupVersion.String()
	EventSubscription_GroupVersionKind = CRDGroupVersion.WithKind(EventSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
