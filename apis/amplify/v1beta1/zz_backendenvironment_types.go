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

type BackendEnvironmentInitParameters struct {

	// Name of deployment artifacts.
	DeploymentArtifacts *string `json:"deploymentArtifacts,omitempty" tf:"deployment_artifacts,omitempty"`

	// AWS CloudFormation stack name of a backend environment.
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`
}

type BackendEnvironmentObservation struct {

	// Unique ID for an Amplify app.
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// ARN for a backend environment that is part of an Amplify app.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Name of deployment artifacts.
	DeploymentArtifacts *string `json:"deploymentArtifacts,omitempty" tf:"deployment_artifacts,omitempty"`

	// Unique ID of the Amplify backend environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// AWS CloudFormation stack name of a backend environment.
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`
}

type BackendEnvironmentParameters struct {

	// Unique ID for an Amplify app.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/amplify/v1beta1.App
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Reference to a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDRef *v1.Reference `json:"appIdRef,omitempty" tf:"-"`

	// Selector for a App in amplify to populate appId.
	// +kubebuilder:validation:Optional
	AppIDSelector *v1.Selector `json:"appIdSelector,omitempty" tf:"-"`

	// Name of deployment artifacts.
	// +kubebuilder:validation:Optional
	DeploymentArtifacts *string `json:"deploymentArtifacts,omitempty" tf:"deployment_artifacts,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// AWS CloudFormation stack name of a backend environment.
	// +kubebuilder:validation:Optional
	StackName *string `json:"stackName,omitempty" tf:"stack_name,omitempty"`
}

// BackendEnvironmentSpec defines the desired state of BackendEnvironment
type BackendEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendEnvironmentParameters `json:"forProvider"`
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
	InitProvider BackendEnvironmentInitParameters `json:"initProvider,omitempty"`
}

// BackendEnvironmentStatus defines the observed state of BackendEnvironment.
type BackendEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackendEnvironment is the Schema for the BackendEnvironments API. Provides an Amplify Backend Environment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BackendEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendEnvironmentSpec   `json:"spec"`
	Status            BackendEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendEnvironmentList contains a list of BackendEnvironments
type BackendEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendEnvironment `json:"items"`
}

// Repository type metadata.
var (
	BackendEnvironment_Kind             = "BackendEnvironment"
	BackendEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendEnvironment_Kind}.String()
	BackendEnvironment_KindAPIVersion   = BackendEnvironment_Kind + "." + CRDGroupVersion.String()
	BackendEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(BackendEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendEnvironment{}, &BackendEnvironmentList{})
}
