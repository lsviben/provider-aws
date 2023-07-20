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

type LambdaFunctionAssociationInitParameters struct {

	// Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	FunctionArnRef *v1.Reference `json:"functionArnRef,omitempty" tf:"-"`

	FunctionArnSelector *v1.Selector `json:"functionArnSelector,omitempty" tf:"-"`

	// The identifier of the Amazon Connect instance. You can find the instanceId in the ARN of the instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type LambdaFunctionAssociationObservation struct {

	// Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// The Amazon Connect instance ID and Lambda Function ARN separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identifier of the Amazon Connect instance. You can find the instanceId in the ARN of the instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type LambdaFunctionAssociationParameters struct {

	// Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// Reference to a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnRef *v1.Reference `json:"functionArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate functionArn.
	// +kubebuilder:validation:Optional
	FunctionArnSelector *v1.Selector `json:"functionArnSelector,omitempty" tf:"-"`

	// The identifier of the Amazon Connect instance. You can find the instanceId in the ARN of the instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// LambdaFunctionAssociationSpec defines the desired state of LambdaFunctionAssociation
type LambdaFunctionAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LambdaFunctionAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider LambdaFunctionAssociationInitParameters `json:"initProvider,omitempty"`
}

// LambdaFunctionAssociationStatus defines the observed state of LambdaFunctionAssociation.
type LambdaFunctionAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LambdaFunctionAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunctionAssociation is the Schema for the LambdaFunctionAssociations API. Provides details about a specific Connect Lambda Function Association.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LambdaFunctionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaFunctionAssociationSpec   `json:"spec"`
	Status            LambdaFunctionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunctionAssociationList contains a list of LambdaFunctionAssociations
type LambdaFunctionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaFunctionAssociation `json:"items"`
}

// Repository type metadata.
var (
	LambdaFunctionAssociation_Kind             = "LambdaFunctionAssociation"
	LambdaFunctionAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LambdaFunctionAssociation_Kind}.String()
	LambdaFunctionAssociation_KindAPIVersion   = LambdaFunctionAssociation_Kind + "." + CRDGroupVersion.String()
	LambdaFunctionAssociation_GroupVersionKind = CRDGroupVersion.WithKind(LambdaFunctionAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&LambdaFunctionAssociation{}, &LambdaFunctionAssociationList{})
}
