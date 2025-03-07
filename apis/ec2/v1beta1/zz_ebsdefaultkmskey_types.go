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

type EBSDefaultKMSKeyInitParameters struct {
}

type EBSDefaultKMSKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn *string `json:"keyArn,omitempty" tf:"key_arn,omitempty"`
}

type EBSDefaultKMSKeyParameters struct {

	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	KeyArn *string `json:"keyArn,omitempty" tf:"key_arn,omitempty"`

	// Reference to a Key in kms to populate keyArn.
	// +kubebuilder:validation:Optional
	KeyArnRef *v1.Reference `json:"keyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyArn.
	// +kubebuilder:validation:Optional
	KeyArnSelector *v1.Selector `json:"keyArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// EBSDefaultKMSKeySpec defines the desired state of EBSDefaultKMSKey
type EBSDefaultKMSKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSDefaultKMSKeyParameters `json:"forProvider"`
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
	InitProvider EBSDefaultKMSKeyInitParameters `json:"initProvider,omitempty"`
}

// EBSDefaultKMSKeyStatus defines the observed state of EBSDefaultKMSKey.
type EBSDefaultKMSKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSDefaultKMSKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EBSDefaultKMSKey is the Schema for the EBSDefaultKMSKeys API. Manages the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EBSDefaultKMSKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EBSDefaultKMSKeySpec   `json:"spec"`
	Status            EBSDefaultKMSKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSDefaultKMSKeyList contains a list of EBSDefaultKMSKeys
type EBSDefaultKMSKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSDefaultKMSKey `json:"items"`
}

// Repository type metadata.
var (
	EBSDefaultKMSKey_Kind             = "EBSDefaultKMSKey"
	EBSDefaultKMSKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSDefaultKMSKey_Kind}.String()
	EBSDefaultKMSKey_KindAPIVersion   = EBSDefaultKMSKey_Kind + "." + CRDGroupVersion.String()
	EBSDefaultKMSKey_GroupVersionKind = CRDGroupVersion.WithKind(EBSDefaultKMSKey_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSDefaultKMSKey{}, &EBSDefaultKMSKeyList{})
}
