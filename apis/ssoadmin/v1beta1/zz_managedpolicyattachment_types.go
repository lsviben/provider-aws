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

type ManagedPolicyAttachmentInitParameters struct {

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn *string `json:"managedPolicyArn,omitempty" tf:"managed_policy_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssoadmin/v1beta1.PermissionSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`

	PermissionSetArnRef *v1.Reference `json:"permissionSetArnRef,omitempty" tf:"-"`

	PermissionSetArnSelector *v1.Selector `json:"permissionSetArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type ManagedPolicyAttachmentObservation struct {

	// The Amazon Resource Names (ARNs) of the Managed Policy, Permission Set, and SSO Instance, separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn *string `json:"managedPolicyArn,omitempty" tf:"managed_policy_arn,omitempty"`

	// The name of the IAM Managed Policy.
	ManagedPolicyName *string `json:"managedPolicyName,omitempty" tf:"managed_policy_name,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`
}

type ManagedPolicyAttachmentParameters struct {

	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `json:"instanceArn,omitempty" tf:"instance_arn,omitempty"`

	// The IAM managed policy Amazon Resource Name (ARN) to be attached to the Permission Set.
	ManagedPolicyArn *string `json:"managedPolicyArn,omitempty" tf:"managed_policy_arn,omitempty"`

	// The Amazon Resource Name (ARN) of the Permission Set.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ssoadmin/v1beta1.PermissionSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	PermissionSetArn *string `json:"permissionSetArn,omitempty" tf:"permission_set_arn,omitempty"`

	// Reference to a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnRef *v1.Reference `json:"permissionSetArnRef,omitempty" tf:"-"`

	// Selector for a PermissionSet in ssoadmin to populate permissionSetArn.
	// +kubebuilder:validation:Optional
	PermissionSetArnSelector *v1.Selector `json:"permissionSetArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// ManagedPolicyAttachmentSpec defines the desired state of ManagedPolicyAttachment
type ManagedPolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPolicyAttachmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ManagedPolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// ManagedPolicyAttachmentStatus defines the observed state of ManagedPolicyAttachment.
type ManagedPolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPolicyAttachment is the Schema for the ManagedPolicyAttachments API. Manages an IAM managed policy for a Single Sign-On (SSO) Permission Set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ManagedPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedPolicyAttachmentSpec   `json:"spec"`
	Status            ManagedPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPolicyAttachmentList contains a list of ManagedPolicyAttachments
type ManagedPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	ManagedPolicyAttachment_Kind             = "ManagedPolicyAttachment"
	ManagedPolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPolicyAttachment_Kind}.String()
	ManagedPolicyAttachment_KindAPIVersion   = ManagedPolicyAttachment_Kind + "." + CRDGroupVersion.String()
	ManagedPolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPolicyAttachment{}, &ManagedPolicyAttachmentList{})
}
