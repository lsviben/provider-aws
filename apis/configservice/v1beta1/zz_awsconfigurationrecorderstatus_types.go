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

type AWSConfigurationRecorderStatusInitParameters struct {

	// Whether the configuration recorder should be enabled or disabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type AWSConfigurationRecorderStatusObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the configuration recorder should be enabled or disabled.
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type AWSConfigurationRecorderStatusParameters struct {

	// Whether the configuration recorder should be enabled or disabled.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AWSConfigurationRecorderStatusSpec defines the desired state of AWSConfigurationRecorderStatus
type AWSConfigurationRecorderStatusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AWSConfigurationRecorderStatusParameters `json:"forProvider"`
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
	InitProvider AWSConfigurationRecorderStatusInitParameters `json:"initProvider,omitempty"`
}

// AWSConfigurationRecorderStatusStatus defines the observed state of AWSConfigurationRecorderStatus.
type AWSConfigurationRecorderStatusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AWSConfigurationRecorderStatusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AWSConfigurationRecorderStatus is the Schema for the AWSConfigurationRecorderStatuss API. Manages status of an AWS Config Configuration Recorder.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AWSConfigurationRecorderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isEnabled) || (has(self.initProvider) && has(self.initProvider.isEnabled))",message="spec.forProvider.isEnabled is a required parameter"
	Spec   AWSConfigurationRecorderStatusSpec   `json:"spec"`
	Status AWSConfigurationRecorderStatusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSConfigurationRecorderStatusList contains a list of AWSConfigurationRecorderStatuss
type AWSConfigurationRecorderStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSConfigurationRecorderStatus `json:"items"`
}

// Repository type metadata.
var (
	AWSConfigurationRecorderStatus_Kind             = "AWSConfigurationRecorderStatus"
	AWSConfigurationRecorderStatus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AWSConfigurationRecorderStatus_Kind}.String()
	AWSConfigurationRecorderStatus_KindAPIVersion   = AWSConfigurationRecorderStatus_Kind + "." + CRDGroupVersion.String()
	AWSConfigurationRecorderStatus_GroupVersionKind = CRDGroupVersion.WithKind(AWSConfigurationRecorderStatus_Kind)
)

func init() {
	SchemeBuilder.Register(&AWSConfigurationRecorderStatus{}, &AWSConfigurationRecorderStatusList{})
}
