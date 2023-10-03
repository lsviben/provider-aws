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

type ApprovalRuleTemplateInitParameters struct {

	// The content of the approval rule template. Maximum of 3000 characters.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The description of the approval rule template. Maximum of 1000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ApprovalRuleTemplateObservation struct {

	// The ID of the approval rule template
	ApprovalRuleTemplateID *string `json:"approvalRuleTemplateId,omitempty" tf:"approval_rule_template_id,omitempty"`

	// The content of the approval rule template. Maximum of 3000 characters.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The date the approval rule template was created, in RFC3339 format.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The description of the approval rule template. Maximum of 1000 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date the approval rule template was most recently changed, in RFC3339 format.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date,omitempty"`

	// The Amazon Resource Name (ARN) of the user who made the most recent changes to the approval rule template.
	LastModifiedUser *string `json:"lastModifiedUser,omitempty" tf:"last_modified_user,omitempty"`

	// The SHA-256 hash signature for the content of the approval rule template.
	RuleContentSha256 *string `json:"ruleContentSha256,omitempty" tf:"rule_content_sha256,omitempty"`
}

type ApprovalRuleTemplateParameters struct {

	// The content of the approval rule template. Maximum of 3000 characters.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The description of the approval rule template. Maximum of 1000 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ApprovalRuleTemplateSpec defines the desired state of ApprovalRuleTemplate
type ApprovalRuleTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApprovalRuleTemplateParameters `json:"forProvider"`
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
	InitProvider ApprovalRuleTemplateInitParameters `json:"initProvider,omitempty"`
}

// ApprovalRuleTemplateStatus defines the observed state of ApprovalRuleTemplate.
type ApprovalRuleTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApprovalRuleTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApprovalRuleTemplate is the Schema for the ApprovalRuleTemplates API. Provides a CodeCommit Approval Rule Template Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ApprovalRuleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	Spec   ApprovalRuleTemplateSpec   `json:"spec"`
	Status ApprovalRuleTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprovalRuleTemplateList contains a list of ApprovalRuleTemplates
type ApprovalRuleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApprovalRuleTemplate `json:"items"`
}

// Repository type metadata.
var (
	ApprovalRuleTemplate_Kind             = "ApprovalRuleTemplate"
	ApprovalRuleTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApprovalRuleTemplate_Kind}.String()
	ApprovalRuleTemplate_KindAPIVersion   = ApprovalRuleTemplate_Kind + "." + CRDGroupVersion.String()
	ApprovalRuleTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ApprovalRuleTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ApprovalRuleTemplate{}, &ApprovalRuleTemplateList{})
}
