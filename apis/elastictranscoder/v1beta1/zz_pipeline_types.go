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

type ContentConfigInitParameters struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ContentConfigObservation struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ContentConfigParameters struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ContentConfigPermissionsInitParameters struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ContentConfigPermissionsObservation struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ContentConfigPermissionsParameters struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type NotificationsInitParameters struct {

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
	Completed *string `json:"completed,omitempty" tf:"completed,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
	Progressing *string `json:"progressing,omitempty" tf:"progressing,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.
	Warning *string `json:"warning,omitempty" tf:"warning,omitempty"`
}

type NotificationsObservation struct {

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
	Completed *string `json:"completed,omitempty" tf:"completed,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
	Progressing *string `json:"progressing,omitempty" tf:"progressing,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.
	Warning *string `json:"warning,omitempty" tf:"warning,omitempty"`
}

type NotificationsParameters struct {

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder has finished processing a job in this pipeline.
	Completed *string `json:"completed,omitempty" tf:"completed,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters an error condition while processing a job in this pipeline.
	Error *string `json:"error,omitempty" tf:"error,omitempty"`

	// The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify when Elastic Transcoder has started to process a job in this pipeline.
	Progressing *string `json:"progressing,omitempty" tf:"progressing,omitempty"`

	// The topic ARN for the Amazon SNS topic that you want to notify when Elastic Transcoder encounters a warning condition while processing a job in this pipeline.
	Warning *string `json:"warning,omitempty" tf:"warning,omitempty"`
}

type PipelineInitParameters struct {

	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`

	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig []ContentConfigInitParameters `json:"contentConfig,omitempty" tf:"content_config,omitempty"`

	// The permissions for the content_config object. (documented below)
	ContentConfigPermissions []ContentConfigPermissionsInitParameters `json:"contentConfigPermissions,omitempty" tf:"content_config_permissions,omitempty"`

	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InputBucket *string `json:"inputBucket,omitempty" tf:"input_bucket,omitempty"`

	InputBucketRef *v1.Reference `json:"inputBucketRef,omitempty" tf:"-"`

	InputBucketSelector *v1.Selector `json:"inputBucketSelector,omitempty" tf:"-"`

	// The name of the pipeline. Maximum 40 characters
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications []NotificationsInitParameters `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket *string `json:"outputBucket,omitempty" tf:"output_bucket,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig []ThumbnailConfigInitParameters `json:"thumbnailConfig,omitempty" tf:"thumbnail_config,omitempty"`

	// The permissions for the thumbnail_config object. (documented below)
	ThumbnailConfigPermissions []ThumbnailConfigPermissionsInitParameters `json:"thumbnailConfigPermissions,omitempty" tf:"thumbnail_config_permissions,omitempty"`
}

type PipelineObservation struct {

	// The ARN of the Elastictranscoder pipeline.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`

	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig []ContentConfigObservation `json:"contentConfig,omitempty" tf:"content_config,omitempty"`

	// The permissions for the content_config object. (documented below)
	ContentConfigPermissions []ContentConfigPermissionsObservation `json:"contentConfigPermissions,omitempty" tf:"content_config_permissions,omitempty"`

	// The ID of the Elastictranscoder pipeline.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket *string `json:"inputBucket,omitempty" tf:"input_bucket,omitempty"`

	// The name of the pipeline. Maximum 40 characters
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications []NotificationsObservation `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket *string `json:"outputBucket,omitempty" tf:"output_bucket,omitempty"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig []ThumbnailConfigObservation `json:"thumbnailConfig,omitempty" tf:"thumbnail_config,omitempty"`

	// The permissions for the thumbnail_config object. (documented below)
	ThumbnailConfigPermissions []ThumbnailConfigPermissionsObservation `json:"thumbnailConfigPermissions,omitempty" tf:"thumbnail_config_permissions,omitempty"`
}

type PipelineParameters struct {

	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKMSKeyArn *string `json:"awsKmsKeyArn,omitempty" tf:"aws_kms_key_arn,omitempty"`

	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig []ContentConfigParameters `json:"contentConfig,omitempty" tf:"content_config,omitempty"`

	// The permissions for the content_config object. (documented below)
	ContentConfigPermissions []ContentConfigPermissionsParameters `json:"contentConfigPermissions,omitempty" tf:"content_config_permissions,omitempty"`

	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	InputBucket *string `json:"inputBucket,omitempty" tf:"input_bucket,omitempty"`

	// Reference to a Bucket in s3 to populate inputBucket.
	// +kubebuilder:validation:Optional
	InputBucketRef *v1.Reference `json:"inputBucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate inputBucket.
	// +kubebuilder:validation:Optional
	InputBucketSelector *v1.Selector `json:"inputBucketSelector,omitempty" tf:"-"`

	// The name of the pipeline. Maximum 40 characters
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications []NotificationsParameters `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket *string `json:"outputBucket,omitempty" tf:"output_bucket,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig []ThumbnailConfigParameters `json:"thumbnailConfig,omitempty" tf:"thumbnail_config,omitempty"`

	// The permissions for the thumbnail_config object. (documented below)
	ThumbnailConfigPermissions []ThumbnailConfigPermissionsParameters `json:"thumbnailConfigPermissions,omitempty" tf:"thumbnail_config_permissions,omitempty"`
}

type ThumbnailConfigInitParameters struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ThumbnailConfigObservation struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ThumbnailConfigParameters struct {

	// The Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The Amazon S3 storage class, Standard or ReducedRedundancy, that you want Elastic Transcoder to assign to the files and playlists that it stores in your Amazon S3 bucket.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type ThumbnailConfigPermissionsInitParameters struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ThumbnailConfigPermissionsObservation struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

type ThumbnailConfigPermissionsParameters struct {

	// The permission that you want to give to the AWS user that you specified in content_config_permissions.grantee. Valid values are Read, ReadAcp, WriteAcp or FullControl.
	Access []*string `json:"access,omitempty" tf:"access,omitempty"`

	// The AWS user or group that you want to have access to transcoded files and playlists.
	Grantee *string `json:"grantee,omitempty" tf:"grantee,omitempty"`

	// Specify the type of value that appears in the content_config_permissions.grantee object. Valid values are Canonical, Email or Group.
	GranteeType *string `json:"granteeType,omitempty" tf:"grantee_type,omitempty"`
}

// PipelineSpec defines the desired state of Pipeline
type PipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PipelineParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider PipelineInitParameters `json:"initProvider,omitempty"`
}

// PipelineStatus defines the observed state of Pipeline.
type PipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pipeline is the Schema for the Pipelines API. Provides an Elastic Transcoder pipeline resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PipelineSpec   `json:"spec"`
	Status            PipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineList contains a list of Pipelines
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

// Repository type metadata.
var (
	Pipeline_Kind             = "Pipeline"
	Pipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pipeline_Kind}.String()
	Pipeline_KindAPIVersion   = Pipeline_Kind + "." + CRDGroupVersion.String()
	Pipeline_GroupVersionKind = CRDGroupVersion.WithKind(Pipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
