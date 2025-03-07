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

type AssessmentTemplateInitParameters struct {

	// The duration of the inspector run.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A block that enables sending notifications about a specified assessment template event to a designated SNS topic. See Event Subscriptions for details.
	EventSubscription []EventSubscriptionInitParameters `json:"eventSubscription,omitempty" tf:"event_subscription,omitempty"`

	// The name of the assessment template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The rules to be used during the run.
	RulesPackageArns []*string `json:"rulesPackageArns,omitempty" tf:"rules_package_arns,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AssessmentTemplateObservation struct {

	// The template assessment ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The duration of the inspector run.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A block that enables sending notifications about a specified assessment template event to a designated SNS topic. See Event Subscriptions for details.
	EventSubscription []EventSubscriptionObservation `json:"eventSubscription,omitempty" tf:"event_subscription,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the assessment template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The rules to be used during the run.
	RulesPackageArns []*string `json:"rulesPackageArns,omitempty" tf:"rules_package_arns,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The assessment target ARN to attach the template to.
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`
}

type AssessmentTemplateParameters struct {

	// The duration of the inspector run.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// A block that enables sending notifications about a specified assessment template event to a designated SNS topic. See Event Subscriptions for details.
	// +kubebuilder:validation:Optional
	EventSubscription []EventSubscriptionParameters `json:"eventSubscription,omitempty" tf:"event_subscription,omitempty"`

	// The name of the assessment template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The rules to be used during the run.
	// +kubebuilder:validation:Optional
	RulesPackageArns []*string `json:"rulesPackageArns,omitempty" tf:"rules_package_arns,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The assessment target ARN to attach the template to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/inspector/v1beta1.AssessmentTarget
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	// Reference to a AssessmentTarget in inspector to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnRef *v1.Reference `json:"targetArnRef,omitempty" tf:"-"`

	// Selector for a AssessmentTarget in inspector to populate targetArn.
	// +kubebuilder:validation:Optional
	TargetArnSelector *v1.Selector `json:"targetArnSelector,omitempty" tf:"-"`
}

type EventSubscriptionInitParameters struct {

	// The event for which you want to receive SNS notifications. Valid values are ASSESSMENT_RUN_STARTED, ASSESSMENT_RUN_COMPLETED, ASSESSMENT_RUN_STATE_CHANGED, and FINDING_REPORTED.
	Event *string `json:"event,omitempty" tf:"event,omitempty"`
}

type EventSubscriptionObservation struct {

	// The event for which you want to receive SNS notifications. Valid values are ASSESSMENT_RUN_STARTED, ASSESSMENT_RUN_COMPLETED, ASSESSMENT_RUN_STATE_CHANGED, and FINDING_REPORTED.
	Event *string `json:"event,omitempty" tf:"event,omitempty"`

	// The ARN of the SNS topic to which notifications are sent.
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type EventSubscriptionParameters struct {

	// The event for which you want to receive SNS notifications. Valid values are ASSESSMENT_RUN_STARTED, ASSESSMENT_RUN_COMPLETED, ASSESSMENT_RUN_STATE_CHANGED, and FINDING_REPORTED.
	// +kubebuilder:validation:Optional
	Event *string `json:"event" tf:"event,omitempty"`

	// The ARN of the SNS topic to which notifications are sent.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// AssessmentTemplateSpec defines the desired state of AssessmentTemplate
type AssessmentTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssessmentTemplateParameters `json:"forProvider"`
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
	InitProvider AssessmentTemplateInitParameters `json:"initProvider,omitempty"`
}

// AssessmentTemplateStatus defines the observed state of AssessmentTemplate.
type AssessmentTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssessmentTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AssessmentTemplate is the Schema for the AssessmentTemplates API. Provides an Inspector Classic Assessment Template.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AssessmentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.duration) || (has(self.initProvider) && has(self.initProvider.duration))",message="spec.forProvider.duration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rulesPackageArns) || (has(self.initProvider) && has(self.initProvider.rulesPackageArns))",message="spec.forProvider.rulesPackageArns is a required parameter"
	Spec   AssessmentTemplateSpec   `json:"spec"`
	Status AssessmentTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssessmentTemplateList contains a list of AssessmentTemplates
type AssessmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssessmentTemplate `json:"items"`
}

// Repository type metadata.
var (
	AssessmentTemplate_Kind             = "AssessmentTemplate"
	AssessmentTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AssessmentTemplate_Kind}.String()
	AssessmentTemplate_KindAPIVersion   = AssessmentTemplate_Kind + "." + CRDGroupVersion.String()
	AssessmentTemplate_GroupVersionKind = CRDGroupVersion.WithKind(AssessmentTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&AssessmentTemplate{}, &AssessmentTemplateList{})
}
