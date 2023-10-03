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

type VaultNotificationsInitParameters struct {

	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	BackupVaultEvents []*string `json:"backupVaultEvents,omitempty" tf:"backup_vault_events,omitempty"`
}

type VaultNotificationsObservation struct {

	// The ARN of the vault.
	BackupVaultArn *string `json:"backupVaultArn,omitempty" tf:"backup_vault_arn,omitempty"`

	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	BackupVaultEvents []*string `json:"backupVaultEvents,omitempty" tf:"backup_vault_events,omitempty"`

	// Name of the backup vault to add notifications for.
	BackupVaultName *string `json:"backupVaultName,omitempty" tf:"backup_vault_name,omitempty"`

	// The name of the vault.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`
}

type VaultNotificationsParameters struct {

	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	// +kubebuilder:validation:Optional
	BackupVaultEvents []*string `json:"backupVaultEvents,omitempty" tf:"backup_vault_events,omitempty"`

	// Name of the backup vault to add notifications for.
	// +crossplane:generate:reference:type=Vault
	// +kubebuilder:validation:Optional
	BackupVaultName *string `json:"backupVaultName,omitempty" tf:"backup_vault_name,omitempty"`

	// Reference to a Vault to populate backupVaultName.
	// +kubebuilder:validation:Optional
	BackupVaultNameRef *v1.Reference `json:"backupVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault to populate backupVaultName.
	// +kubebuilder:validation:Optional
	BackupVaultNameSelector *v1.Selector `json:"backupVaultNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// Reference to a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnRef *v1.Reference `json:"snsTopicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate snsTopicArn.
	// +kubebuilder:validation:Optional
	SnsTopicArnSelector *v1.Selector `json:"snsTopicArnSelector,omitempty" tf:"-"`
}

// VaultNotificationsSpec defines the desired state of VaultNotifications
type VaultNotificationsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultNotificationsParameters `json:"forProvider"`
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
	InitProvider VaultNotificationsInitParameters `json:"initProvider,omitempty"`
}

// VaultNotificationsStatus defines the observed state of VaultNotifications.
type VaultNotificationsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultNotificationsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultNotifications is the Schema for the VaultNotificationss API. Provides an AWS Backup vault notifications resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VaultNotifications struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backupVaultEvents) || (has(self.initProvider) && has(self.initProvider.backupVaultEvents))",message="spec.forProvider.backupVaultEvents is a required parameter"
	Spec   VaultNotificationsSpec   `json:"spec"`
	Status VaultNotificationsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultNotificationsList contains a list of VaultNotificationss
type VaultNotificationsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultNotifications `json:"items"`
}

// Repository type metadata.
var (
	VaultNotifications_Kind             = "VaultNotifications"
	VaultNotifications_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultNotifications_Kind}.String()
	VaultNotifications_KindAPIVersion   = VaultNotifications_Kind + "." + CRDGroupVersion.String()
	VaultNotifications_GroupVersionKind = CRDGroupVersion.WithKind(VaultNotifications_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultNotifications{}, &VaultNotificationsList{})
}
