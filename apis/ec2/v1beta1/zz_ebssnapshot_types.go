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

type EBSSnapshotObservation struct {

	// Amazon Resource Name (ARN) of the EBS Snapshot.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyID *string `json:"dataEncryptionKeyId,omitempty" tf:"data_encryption_key_id,omitempty"`

	// Whether the snapshot is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The snapshot ID (e.g., snap-59fcb34e).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Value from an Amazon-maintained list (amazon, aws-marketplace, microsoft) of snapshot owners.
	OwnerAlias *string `json:"ownerAlias,omitempty" tf:"owner_alias,omitempty"`

	// The AWS account ID of the EBS snapshot owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The size of the drive in GiBs.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
}

type EBSSnapshotParameters struct {

	// A description of what the snapshot is.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost on which to create a local snapshot.
	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Indicates whether to permanently restore an archived snapshot.
	// +kubebuilder:validation:Optional
	PermanentRestore *bool `json:"permanentRestore,omitempty" tf:"permanent_restore,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the storage tier. Valid values are archive and standard. Default value is standard.
	// +kubebuilder:validation:Optional
	StorageTier *string `json:"storageTier,omitempty" tf:"storage_tier,omitempty"`

	// A map of tags to assign to the snapshot. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
	// +kubebuilder:validation:Optional
	TemporaryRestoreDays *float64 `json:"temporaryRestoreDays,omitempty" tf:"temporary_restore_days,omitempty"`

	// The Volume ID of which to make a snapshot.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.EBSVolume
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a EBSVolume in ec2 to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// EBSSnapshotSpec defines the desired state of EBSSnapshot
type EBSSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EBSSnapshotParameters `json:"forProvider"`
}

// EBSSnapshotStatus defines the observed state of EBSSnapshot.
type EBSSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EBSSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshot is the Schema for the EBSSnapshots API. Provides an elastic block storage snapshot resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EBSSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EBSSnapshotSpec   `json:"spec"`
	Status            EBSSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EBSSnapshotList contains a list of EBSSnapshots
type EBSSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EBSSnapshot `json:"items"`
}

// Repository type metadata.
var (
	EBSSnapshot_Kind             = "EBSSnapshot"
	EBSSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EBSSnapshot_Kind}.String()
	EBSSnapshot_KindAPIVersion   = EBSSnapshot_Kind + "." + CRDGroupVersion.String()
	EBSSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(EBSSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&EBSSnapshot{}, &EBSSnapshotList{})
}
