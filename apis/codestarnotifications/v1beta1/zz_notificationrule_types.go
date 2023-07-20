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

type NotificationRuleInitParameters struct {

	// The level of detail to include in the notifications for this resource. Possible values are BASIC and FULL.
	DetailType *string `json:"detailType,omitempty" tf:"detail_type,omitempty"`

	// A list of event types associated with this notification rule.
	// For list of allowed events see here.
	EventTypeIds []*string `json:"eventTypeIds,omitempty" tf:"event_type_ids,omitempty"`

	// The name of notification rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN of the resource to associate with the notification rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codecommit/v1beta1.Repository
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	ResourceRef *v1.Reference `json:"resourceRef,omitempty" tf:"-"`

	ResourceSelector *v1.Selector `json:"resourceSelector,omitempty" tf:"-"`

	// The status of the notification rule. Possible values are ENABLED and DISABLED, default is ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Target []TargetInitParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type NotificationRuleObservation struct {

	// The codestar notification rule ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The level of detail to include in the notifications for this resource. Possible values are BASIC and FULL.
	DetailType *string `json:"detailType,omitempty" tf:"detail_type,omitempty"`

	// A list of event types associated with this notification rule.
	// For list of allowed events see here.
	EventTypeIds []*string `json:"eventTypeIds,omitempty" tf:"event_type_ids,omitempty"`

	// The codestar notification rule ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of notification rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ARN of the resource to associate with the notification rule.
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// The status of the notification rule. Possible values are ENABLED and DISABLED, default is ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Target []TargetObservation `json:"target,omitempty" tf:"target,omitempty"`
}

type NotificationRuleParameters struct {

	// The level of detail to include in the notifications for this resource. Possible values are BASIC and FULL.
	DetailType *string `json:"detailType,omitempty" tf:"detail_type,omitempty"`

	// A list of event types associated with this notification rule.
	// For list of allowed events see here.
	EventTypeIds []*string `json:"eventTypeIds,omitempty" tf:"event_type_ids,omitempty"`

	// The name of notification rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The ARN of the resource to associate with the notification rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codecommit/v1beta1.Repository
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Resource *string `json:"resource,omitempty" tf:"resource,omitempty"`

	// Reference to a Repository in codecommit to populate resource.
	// +kubebuilder:validation:Optional
	ResourceRef *v1.Reference `json:"resourceRef,omitempty" tf:"-"`

	// Selector for a Repository in codecommit to populate resource.
	// +kubebuilder:validation:Optional
	ResourceSelector *v1.Selector `json:"resourceSelector,omitempty" tf:"-"`

	// The status of the notification rule. Possible values are ENABLED and DISABLED, default is ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Target []TargetParameters `json:"target,omitempty" tf:"target,omitempty"`
}

type TargetInitParameters struct {

	// The ARN of notification rule target. For example, a SNS Topic ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	AddressRef *v1.Reference `json:"addressRef,omitempty" tf:"-"`

	AddressSelector *v1.Selector `json:"addressSelector,omitempty" tf:"-"`

	// The type of the notification target. Default value is SNS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetObservation struct {

	// The ARN of notification rule target. For example, a SNS Topic ARN.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The status of the notification rule. Possible values are ENABLED and DISABLED, default is ENABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The type of the notification target. Default value is SNS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetParameters struct {

	// The ARN of notification rule target. For example, a SNS Topic ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Reference to a Topic in sns to populate address.
	// +kubebuilder:validation:Optional
	AddressRef *v1.Reference `json:"addressRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate address.
	// +kubebuilder:validation:Optional
	AddressSelector *v1.Selector `json:"addressSelector,omitempty" tf:"-"`

	// The type of the notification target. Default value is SNS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// NotificationRuleSpec defines the desired state of NotificationRule
type NotificationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider NotificationRuleInitParameters `json:"initProvider,omitempty"`
}

// NotificationRuleStatus defines the observed state of NotificationRule.
type NotificationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRule is the Schema for the NotificationRules API. Provides a CodeStar Notifications Rule
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type NotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.detailType) || has(self.initProvider.detailType)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.eventTypeIds) || has(self.initProvider.eventTypeIds)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	Spec   NotificationRuleSpec   `json:"spec"`
	Status NotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRuleList contains a list of NotificationRules
type NotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRule `json:"items"`
}

// Repository type metadata.
var (
	NotificationRule_Kind             = "NotificationRule"
	NotificationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRule_Kind}.String()
	NotificationRule_KindAPIVersion   = NotificationRule_Kind + "." + CRDGroupVersion.String()
	NotificationRule_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRule{}, &NotificationRuleList{})
}
