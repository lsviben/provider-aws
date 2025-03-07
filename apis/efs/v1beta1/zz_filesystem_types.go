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

type FileSystemInitParameters struct {

	// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See user guide for more information.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. See Elastic File System
	// user guide for more information.
	CreationToken *string `json:"creationToken,omitempty" tf:"creation_token,omitempty"`

	// If true, the disk will be encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// A file system lifecycle policy object (documented below).
	LifecyclePolicy []LifecyclePolicyInitParameters `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy,omitempty"`

	// The file system performance mode. Can be either "generalPurpose" or "maxIO" (Default: "generalPurpose").
	PerformanceMode *string `json:"performanceMode,omitempty" tf:"performance_mode,omitempty"`

	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with throughput_mode set to provisioned.
	ProvisionedThroughputInMibps *float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Throughput mode for the file system. Defaults to bursting. Valid values: bursting, provisioned, or elastic. When using provisioned, also set provisioned_throughput_in_mibps.
	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`
}

type FileSystemObservation struct {

	// Amazon Resource Name of the file system.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the Availability Zone in which the file system's One Zone storage classes exist.
	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See user guide for more information.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. See Elastic File System
	// user guide for more information.
	CreationToken *string `json:"creationToken,omitempty" tf:"creation_token,omitempty"`

	// The DNS name for the filesystem per documented convention.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// If true, the disk will be encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ID that identifies the file system (e.g., fs-ccfc0d65).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// A file system lifecycle policy object (documented below).
	LifecyclePolicy []LifecyclePolicyObservation `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy,omitempty"`

	// The current number of mount targets that the file system has.
	NumberOfMountTargets *float64 `json:"numberOfMountTargets,omitempty" tf:"number_of_mount_targets,omitempty"`

	// The AWS account that created the file system. If the file system was createdby an IAM user, the parent account to which the user belongs is the owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The file system performance mode. Can be either "generalPurpose" or "maxIO" (Default: "generalPurpose").
	PerformanceMode *string `json:"performanceMode,omitempty" tf:"performance_mode,omitempty"`

	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with throughput_mode set to provisioned.
	ProvisionedThroughputInMibps *float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps,omitempty"`

	// The latest known metered size (in bytes) of data stored in the file system, the value is not the exact size that the file system was at any point in time. See Size In Bytes.
	SizeInBytes []SizeInBytesObservation `json:"sizeInBytes,omitempty" tf:"size_in_bytes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Throughput mode for the file system. Defaults to bursting. Valid values: bursting, provisioned, or elastic. When using provisioned, also set provisioned_throughput_in_mibps.
	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`
}

type FileSystemParameters struct {

	// the AWS Availability Zone in which to create the file system. Used to create a file system that uses One Zone storage classes. See user guide for more information.
	// +kubebuilder:validation:Optional
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// A unique name (a maximum of 64 characters are allowed)
	// used as reference when creating the Elastic File System to ensure idempotent file
	// system creation. See Elastic File System
	// user guide for more information.
	// +kubebuilder:validation:Optional
	CreationToken *string `json:"creationToken,omitempty" tf:"creation_token,omitempty"`

	// If true, the disk will be encrypted.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// A file system lifecycle policy object (documented below).
	// +kubebuilder:validation:Optional
	LifecyclePolicy []LifecyclePolicyParameters `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy,omitempty"`

	// The file system performance mode. Can be either "generalPurpose" or "maxIO" (Default: "generalPurpose").
	// +kubebuilder:validation:Optional
	PerformanceMode *string `json:"performanceMode,omitempty" tf:"performance_mode,omitempty"`

	// The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with throughput_mode set to provisioned.
	// +kubebuilder:validation:Optional
	ProvisionedThroughputInMibps *float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Throughput mode for the file system. Defaults to bursting. Valid values: bursting, provisioned, or elastic. When using provisioned, also set provisioned_throughput_in_mibps.
	// +kubebuilder:validation:Optional
	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`
}

type LifecyclePolicyInitParameters struct {

	// Indicates how long it takes to transition files to the IA storage class. Valid values: AFTER_1_DAY, AFTER_7_DAYS, AFTER_14_DAYS, AFTER_30_DAYS, AFTER_60_DAYS, or AFTER_90_DAYS.
	TransitionToIa *string `json:"transitionToIa,omitempty" tf:"transition_to_ia,omitempty"`

	// Describes the policy used to transition a file from infequent access storage to primary storage. Valid values: AFTER_1_ACCESS.
	TransitionToPrimaryStorageClass *string `json:"transitionToPrimaryStorageClass,omitempty" tf:"transition_to_primary_storage_class,omitempty"`
}

type LifecyclePolicyObservation struct {

	// Indicates how long it takes to transition files to the IA storage class. Valid values: AFTER_1_DAY, AFTER_7_DAYS, AFTER_14_DAYS, AFTER_30_DAYS, AFTER_60_DAYS, or AFTER_90_DAYS.
	TransitionToIa *string `json:"transitionToIa,omitempty" tf:"transition_to_ia,omitempty"`

	// Describes the policy used to transition a file from infequent access storage to primary storage. Valid values: AFTER_1_ACCESS.
	TransitionToPrimaryStorageClass *string `json:"transitionToPrimaryStorageClass,omitempty" tf:"transition_to_primary_storage_class,omitempty"`
}

type LifecyclePolicyParameters struct {

	// Indicates how long it takes to transition files to the IA storage class. Valid values: AFTER_1_DAY, AFTER_7_DAYS, AFTER_14_DAYS, AFTER_30_DAYS, AFTER_60_DAYS, or AFTER_90_DAYS.
	// +kubebuilder:validation:Optional
	TransitionToIa *string `json:"transitionToIa,omitempty" tf:"transition_to_ia,omitempty"`

	// Describes the policy used to transition a file from infequent access storage to primary storage. Valid values: AFTER_1_ACCESS.
	// +kubebuilder:validation:Optional
	TransitionToPrimaryStorageClass *string `json:"transitionToPrimaryStorageClass,omitempty" tf:"transition_to_primary_storage_class,omitempty"`
}

type SizeInBytesInitParameters struct {
}

type SizeInBytesObservation struct {

	// The latest known metered size (in bytes) of data stored in the file system.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`

	// The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
	ValueInIa *float64 `json:"valueInIa,omitempty" tf:"value_in_ia,omitempty"`

	// The latest known metered size (in bytes) of data stored in the Standard storage class.
	ValueInStandard *float64 `json:"valueInStandard,omitempty" tf:"value_in_standard,omitempty"`
}

type SizeInBytesParameters struct {
}

// FileSystemSpec defines the desired state of FileSystem
type FileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileSystemParameters `json:"forProvider"`
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
	InitProvider FileSystemInitParameters `json:"initProvider,omitempty"`
}

// FileSystemStatus defines the observed state of FileSystem.
type FileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystem is the Schema for the FileSystems API. Provides an Elastic File System (EFS) File System resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemSpec   `json:"spec"`
	Status            FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemList contains a list of FileSystems
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

// Repository type metadata.
var (
	FileSystem_Kind             = "FileSystem"
	FileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FileSystem_Kind}.String()
	FileSystem_KindAPIVersion   = FileSystem_Kind + "." + CRDGroupVersion.String()
	FileSystem_GroupVersionKind = CRDGroupVersion.WithKind(FileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}
