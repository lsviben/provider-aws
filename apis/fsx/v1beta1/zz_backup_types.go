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

type BackupInitParameters struct {

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the volume to back up. Required if backing up a ONTAP Volume.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type BackupObservation struct {

	// Amazon Resource Name of the backup.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the file system to back up. Required if backing up Lustre or Windows file systems.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Identifier of the backup, e.g., fs-12345678
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// AWS account identifier that created the file system.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of the file system backup.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the volume to back up. Required if backing up a ONTAP Volume.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type BackupParameters struct {

	// The ID of the file system to back up. Required if backing up Lustre or Windows file systems.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/fsx/v1beta1.LustreFileSystem
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a LustreFileSystem in fsx to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a LustreFileSystem in fsx to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the volume to back up. Required if backing up a ONTAP Volume.
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupParameters `json:"forProvider"`
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
	InitProvider BackupInitParameters `json:"initProvider,omitempty"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Backup is the Schema for the Backups API. Manages a FSx Backup.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSpec   `json:"spec"`
	Status            BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backups
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// Repository type metadata.
var (
	Backup_Kind             = "Backup"
	Backup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backup_Kind}.String()
	Backup_KindAPIVersion   = Backup_Kind + "." + CRDGroupVersion.String()
	Backup_GroupVersionKind = CRDGroupVersion.WithKind(Backup_Kind)
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
