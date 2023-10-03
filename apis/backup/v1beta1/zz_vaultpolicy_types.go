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

type VaultPolicyInitParameters struct {

	// The backup vault access policy document in JSON format.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type VaultPolicyObservation struct {

	// The ARN of the vault.
	BackupVaultArn *string `json:"backupVaultArn,omitempty" tf:"backup_vault_arn,omitempty"`

	// Name of the backup vault to add policy for.
	BackupVaultName *string `json:"backupVaultName,omitempty" tf:"backup_vault_name,omitempty"`

	// The name of the vault.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The backup vault access policy document in JSON format.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type VaultPolicyParameters struct {

	// Name of the backup vault to add policy for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/backup/v1beta1.Vault
	// +kubebuilder:validation:Optional
	BackupVaultName *string `json:"backupVaultName,omitempty" tf:"backup_vault_name,omitempty"`

	// Reference to a Vault in backup to populate backupVaultName.
	// +kubebuilder:validation:Optional
	BackupVaultNameRef *v1.Reference `json:"backupVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in backup to populate backupVaultName.
	// +kubebuilder:validation:Optional
	BackupVaultNameSelector *v1.Selector `json:"backupVaultNameSelector,omitempty" tf:"-"`

	// The backup vault access policy document in JSON format.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// VaultPolicySpec defines the desired state of VaultPolicy
type VaultPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultPolicyParameters `json:"forProvider"`
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
	InitProvider VaultPolicyInitParameters `json:"initProvider,omitempty"`
}

// VaultPolicyStatus defines the observed state of VaultPolicy.
type VaultPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultPolicy is the Schema for the VaultPolicys API. Provides an AWS Backup vault policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VaultPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   VaultPolicySpec   `json:"spec"`
	Status VaultPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultPolicyList contains a list of VaultPolicys
type VaultPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultPolicy `json:"items"`
}

// Repository type metadata.
var (
	VaultPolicy_Kind             = "VaultPolicy"
	VaultPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultPolicy_Kind}.String()
	VaultPolicy_KindAPIVersion   = VaultPolicy_Kind + "." + CRDGroupVersion.String()
	VaultPolicy_GroupVersionKind = CRDGroupVersion.WithKind(VaultPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultPolicy{}, &VaultPolicyList{})
}
