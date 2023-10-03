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

type BackupPolicyBackupPolicyInitParameters struct {

	// A status of the backup policy. Valid values: ENABLED, DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BackupPolicyBackupPolicyObservation struct {

	// A status of the backup policy. Valid values: ENABLED, DISABLED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BackupPolicyBackupPolicyParameters struct {

	// A status of the backup policy. Valid values: ENABLED, DISABLED.
	// +kubebuilder:validation:Optional
	Status *string `json:"status" tf:"status,omitempty"`
}

type BackupPolicyInitParameters struct {

	// A backup_policy object (documented below).
	BackupPolicy []BackupPolicyBackupPolicyInitParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`
}

type BackupPolicyObservation struct {

	// A backup_policy object (documented below).
	BackupPolicy []BackupPolicyBackupPolicyObservation `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// The ID of the EFS file system.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// The ID that identifies the file system (e.g., fs-ccfc0d65).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupPolicyParameters struct {

	// A backup_policy object (documented below).
	// +kubebuilder:validation:Optional
	BackupPolicy []BackupPolicyBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// The ID of the EFS file system.
	// +crossplane:generate:reference:type=FileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BackupPolicySpec defines the desired state of BackupPolicy
type BackupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyParameters `json:"forProvider"`
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
	InitProvider BackupPolicyInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyStatus defines the observed state of BackupPolicy.
type BackupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicy is the Schema for the BackupPolicys API. Provides an Elastic File System (EFS) Backup Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backupPolicy) || (has(self.initProvider) && has(self.initProvider.backupPolicy))",message="spec.forProvider.backupPolicy is a required parameter"
	Spec   BackupPolicySpec   `json:"spec"`
	Status BackupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyList contains a list of BackupPolicys
type BackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicy `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicy_Kind             = "BackupPolicy"
	BackupPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicy_Kind}.String()
	BackupPolicy_KindAPIVersion   = BackupPolicy_Kind + "." + CRDGroupVersion.String()
	BackupPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicy{}, &BackupPolicyList{})
}
