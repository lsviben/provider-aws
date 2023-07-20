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

type StackInitParameters struct {

	// A list of capabilities.
	// Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, or CAPABILITY_AUTO_EXPAND
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with on_failure.
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// Stack name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []*string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`

	// Action to be taken if stack creation fails. This must be
	// one of: DO_NOTHING, ROLLBACK, or DELETE. Conflicts with disable_rollback.
	OnFailure *string `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Structure containing the stack policy body.
	// Conflicts w/ policy_url.
	PolicyBody *string `json:"policyBody,omitempty" tf:"policy_body,omitempty"`

	// Location of a file containing the stack policy.
	// Conflicts w/ policy_body.
	PolicyURL *string `json:"policyUrl,omitempty" tf:"policy_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// The amount of time that can pass before the stack status becomes CREATE_FAILED.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes,omitempty" tf:"timeout_in_minutes,omitempty"`
}

type StackObservation struct {

	// A list of capabilities.
	// Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, or CAPABILITY_AUTO_EXPAND
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with on_failure.
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// A unique identifier of the stack.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Stack name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []*string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`

	// Action to be taken if stack creation fails. This must be
	// one of: DO_NOTHING, ROLLBACK, or DELETE. Conflicts with disable_rollback.
	OnFailure *string `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// A map of outputs from the stack.
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Structure containing the stack policy body.
	// Conflicts w/ policy_url.
	PolicyBody *string `json:"policyBody,omitempty" tf:"policy_body,omitempty"`

	// Location of a file containing the stack policy.
	// Conflicts w/ policy_body.
	PolicyURL *string `json:"policyUrl,omitempty" tf:"policy_url,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// The amount of time that can pass before the stack status becomes CREATE_FAILED.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes,omitempty" tf:"timeout_in_minutes,omitempty"`
}

type StackParameters struct {

	// A list of capabilities.
	// Valid values: CAPABILITY_IAM, CAPABILITY_NAMED_IAM, or CAPABILITY_AUTO_EXPAND
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with on_failure.
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// Stack name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []*string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`

	// Action to be taken if stack creation fails. This must be
	// one of: DO_NOTHING, ROLLBACK, or DELETE. Conflicts with disable_rollback.
	OnFailure *string `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Structure containing the stack policy body.
	// Conflicts w/ policy_url.
	PolicyBody *string `json:"policyBody,omitempty" tf:"policy_body,omitempty"`

	// Location of a file containing the stack policy.
	// Conflicts w/ policy_body.
	PolicyURL *string `json:"policyUrl,omitempty" tf:"policy_url,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// The amount of time that can pass before the stack status becomes CREATE_FAILED.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes,omitempty" tf:"timeout_in_minutes,omitempty"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider StackInitParameters `json:"initProvider,omitempty"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stack is the Schema for the Stacks API. Provides a CloudFormation Stack resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSpec   `json:"spec"`
	Status            StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
