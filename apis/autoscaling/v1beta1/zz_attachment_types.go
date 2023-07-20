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

type AttachmentInitParameters struct {

	// ARN of an ALB Target Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ALBTargetGroupArn *string `json:"albTargetGroupArn,omitempty" tf:"alb_target_group_arn,omitempty"`

	ALBTargetGroupArnRef *v1.Reference `json:"albTargetGroupArnRef,omitempty" tf:"-"`

	ALBTargetGroupArnSelector *v1.Selector `json:"albTargetGroupArnSelector,omitempty" tf:"-"`

	// Name of ASG to associate with the ELB.
	// +crossplane:generate:reference:type=AutoscalingGroup
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Name of the ELB.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	ELBRef *v1.Reference `json:"elbRef,omitempty" tf:"-"`

	ELBSelector *v1.Selector `json:"elbSelector,omitempty" tf:"-"`

	// ARN of a load balancer target group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`

	LBTargetGroupArnRef *v1.Reference `json:"lbTargetGroupArnRef,omitempty" tf:"-"`

	LBTargetGroupArnSelector *v1.Selector `json:"lbTargetGroupArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type AttachmentObservation struct {

	// ARN of an ALB Target Group.
	ALBTargetGroupArn *string `json:"albTargetGroupArn,omitempty" tf:"alb_target_group_arn,omitempty"`

	// Name of ASG to associate with the ELB.
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Name of the ELB.
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of a load balancer target group.
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`
}

type AttachmentParameters struct {

	// ARN of an ALB Target Group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	ALBTargetGroupArn *string `json:"albTargetGroupArn,omitempty" tf:"alb_target_group_arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate albTargetGroupArn.
	// +kubebuilder:validation:Optional
	ALBTargetGroupArnRef *v1.Reference `json:"albTargetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate albTargetGroupArn.
	// +kubebuilder:validation:Optional
	ALBTargetGroupArnSelector *v1.Selector `json:"albTargetGroupArnSelector,omitempty" tf:"-"`

	// Name of ASG to associate with the ELB.
	// +crossplane:generate:reference:type=AutoscalingGroup
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Name of the ELB.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	// Reference to a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBRef *v1.Reference `json:"elbRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBSelector *v1.Selector `json:"elbSelector,omitempty" tf:"-"`

	// ARN of a load balancer target group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnRef *v1.Reference `json:"lbTargetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnSelector *v1.Selector `json:"lbTargetGroupArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AttachmentInitParameters `json:"initProvider,omitempty"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attachment is the Schema for the Attachments API. Provides an AutoScaling Group Attachment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
