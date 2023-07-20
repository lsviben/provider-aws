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

type ThingPrincipalAttachmentInitParameters struct {

	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iot/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	PrincipalRef *v1.Reference `json:"principalRef,omitempty" tf:"-"`

	PrincipalSelector *v1.Selector `json:"principalSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The name of the thing.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iot/v1beta1.Thing
	Thing *string `json:"thing,omitempty" tf:"thing,omitempty"`

	ThingRef *v1.Reference `json:"thingRef,omitempty" tf:"-"`

	ThingSelector *v1.Selector `json:"thingSelector,omitempty" tf:"-"`
}

type ThingPrincipalAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// The name of the thing.
	Thing *string `json:"thing,omitempty" tf:"thing,omitempty"`
}

type ThingPrincipalAttachmentParameters struct {

	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iot/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Reference to a Certificate in iot to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalRef *v1.Reference `json:"principalRef,omitempty" tf:"-"`

	// Selector for a Certificate in iot to populate principal.
	// +kubebuilder:validation:Optional
	PrincipalSelector *v1.Selector `json:"principalSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The name of the thing.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iot/v1beta1.Thing
	Thing *string `json:"thing,omitempty" tf:"thing,omitempty"`

	// Reference to a Thing in iot to populate thing.
	// +kubebuilder:validation:Optional
	ThingRef *v1.Reference `json:"thingRef,omitempty" tf:"-"`

	// Selector for a Thing in iot to populate thing.
	// +kubebuilder:validation:Optional
	ThingSelector *v1.Selector `json:"thingSelector,omitempty" tf:"-"`
}

// ThingPrincipalAttachmentSpec defines the desired state of ThingPrincipalAttachment
type ThingPrincipalAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ThingPrincipalAttachmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ThingPrincipalAttachmentInitParameters `json:"initProvider,omitempty"`
}

// ThingPrincipalAttachmentStatus defines the observed state of ThingPrincipalAttachment.
type ThingPrincipalAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ThingPrincipalAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ThingPrincipalAttachment is the Schema for the ThingPrincipalAttachments API. Provides AWS IoT Thing Principal attachment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ThingPrincipalAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThingPrincipalAttachmentSpec   `json:"spec"`
	Status            ThingPrincipalAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingPrincipalAttachmentList contains a list of ThingPrincipalAttachments
type ThingPrincipalAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ThingPrincipalAttachment `json:"items"`
}

// Repository type metadata.
var (
	ThingPrincipalAttachment_Kind             = "ThingPrincipalAttachment"
	ThingPrincipalAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThingPrincipalAttachment_Kind}.String()
	ThingPrincipalAttachment_KindAPIVersion   = ThingPrincipalAttachment_Kind + "." + CRDGroupVersion.String()
	ThingPrincipalAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ThingPrincipalAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ThingPrincipalAttachment{}, &ThingPrincipalAttachmentList{})
}
