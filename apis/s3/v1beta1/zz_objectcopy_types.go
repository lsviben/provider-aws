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

type ObjectCopyGrantInitParameters struct {

	// Email address of the grantee. Used only when type is AmazonCustomerByEmail.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Canonical user ID of the grantee. Used only when type is CanonicalUser.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of permissions to grant to grantee. Valid values are READ, READ_ACP, WRITE_ACP, FULL_CONTROL.
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// - Type of grantee. Valid values are CanonicalUser, Group, and AmazonCustomerByEmail.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URI of the grantee group. Used only when type is Group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ObjectCopyGrantObservation struct {

	// Email address of the grantee. Used only when type is AmazonCustomerByEmail.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Canonical user ID of the grantee. Used only when type is CanonicalUser.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of permissions to grant to grantee. Valid values are READ, READ_ACP, WRITE_ACP, FULL_CONTROL.
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// - Type of grantee. Valid values are CanonicalUser, Group, and AmazonCustomerByEmail.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URI of the grantee group. Used only when type is Group.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ObjectCopyGrantParameters struct {

	// Email address of the grantee. Used only when type is AmazonCustomerByEmail.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Canonical user ID of the grantee. Used only when type is CanonicalUser.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of permissions to grant to grantee. Valid values are READ, READ_ACP, WRITE_ACP, FULL_CONTROL.
	// +kubebuilder:validation:Optional
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// - Type of grantee. Valid values are CanonicalUser, Group, and AmazonCustomerByEmail.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// URI of the grantee group. Used only when type is Group.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ObjectCopyInitParameters struct {

	// Canned ACL to apply. Defaults to private. Valid values are private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, and bucket-owner-full-control. Conflicts with grant.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Name of the bucket to put the file in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopyIfMatch *string `json:"copyIfMatch,omitempty" tf:"copy_if_match,omitempty"`

	// Copies the object if it has been modified since the specified time, in RFC3339 format.
	CopyIfModifiedSince *string `json:"copyIfModifiedSince,omitempty" tf:"copy_if_modified_since,omitempty"`

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopyIfNoneMatch *string `json:"copyIfNoneMatch,omitempty" tf:"copy_if_none_match,omitempty"`

	// Copies the object if it hasn't been modified since the specified time, in RFC3339 format.
	CopyIfUnmodifiedSince *string `json:"copyIfUnmodifiedSince,omitempty" tf:"copy_if_unmodified_since,omitempty"`

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	CustomerAlgorithm *string `json:"customerAlgorithm,omitempty" tf:"customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	CustomerKeyMd5 *string `json:"customerKeyMd5,omitempty" tf:"customer_key_md5,omitempty"`

	// Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedSourceBucketOwner *string `json:"expectedSourceBucketOwner,omitempty" tf:"expected_source_bucket_owner,omitempty"`

	// Date and time at which the object is no longer cacheable, in RFC3339 format.
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Configuration block for header grants. Documented below. Conflicts with acl.
	Grant []ObjectCopyGrantInitParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// Name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are COPY and REPLACE.
	MetadataDirective *string `json:"metadataDirective,omitempty" tf:"metadata_directive,omitempty"`

	// The legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is requester.
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// Specifies server-side encryption of the object in S3. Valid values are AES256 and aws:kms.
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (/). For example, testbucket/test1.json. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format arn:aws:s3:<Region>:<account-id>:accesspoint/<access-point-name>/object/<key>. For example, arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the algorithm to use when decrypting the source object (for example, AES256).
	SourceCustomerAlgorithm *string `json:"sourceCustomerAlgorithm,omitempty" tf:"source_customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	SourceCustomerKeyMd5 *string `json:"sourceCustomerKeyMd5,omitempty" tf:"source_customer_key_md5,omitempty"`

	// Specifies the desired storage class for the object. Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are COPY and REPLACE.
	TaggingDirective *string `json:"taggingDirective,omitempty" tf:"tagging_directive,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectCopyObservation struct {

	// Canned ACL to apply. Defaults to private. Valid values are private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, and bucket-owner-full-control. Conflicts with grant.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Name of the bucket to put the file in.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopyIfMatch *string `json:"copyIfMatch,omitempty" tf:"copy_if_match,omitempty"`

	// Copies the object if it has been modified since the specified time, in RFC3339 format.
	CopyIfModifiedSince *string `json:"copyIfModifiedSince,omitempty" tf:"copy_if_modified_since,omitempty"`

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	CopyIfNoneMatch *string `json:"copyIfNoneMatch,omitempty" tf:"copy_if_none_match,omitempty"`

	// Copies the object if it hasn't been modified since the specified time, in RFC3339 format.
	CopyIfUnmodifiedSince *string `json:"copyIfUnmodifiedSince,omitempty" tf:"copy_if_unmodified_since,omitempty"`

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	CustomerAlgorithm *string `json:"customerAlgorithm,omitempty" tf:"customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	CustomerKeyMd5 *string `json:"customerKeyMd5,omitempty" tf:"customer_key_md5,omitempty"`

	// ETag generated for the object (an MD5 sum of the object content). For plaintext objects or objects encrypted with an AWS-managed key, the hash is an MD5 digest of the object data. For objects encrypted with a KMS key or objects created by either the Multipart Upload or Part Copy operation, the hash is not an MD5 digest, regardless of the method of encryption. More information on possible values can be found on Common Response Headers.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedSourceBucketOwner *string `json:"expectedSourceBucketOwner,omitempty" tf:"expected_source_bucket_owner,omitempty"`

	// If the object expiration is configured, this attribute will be set.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Date and time at which the object is no longer cacheable, in RFC3339 format.
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Configuration block for header grants. Documented below. Conflicts with acl.
	Grant []ObjectCopyGrantObservation `json:"grant,omitempty" tf:"grant,omitempty"`

	// Canonical user ID of the grantee. Used only when type is CanonicalUser.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the object once it is in the bucket.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Returns the date that the object was last modified, in RFC3339 format.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are COPY and REPLACE.
	MetadataDirective *string `json:"metadataDirective,omitempty" tf:"metadata_directive,omitempty"`

	// The legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// If present, indicates that the requester was successfully charged for the request.
	RequestCharged *bool `json:"requestCharged,omitempty" tf:"request_charged,omitempty"`

	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is requester.
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// Specifies server-side encryption of the object in S3. Valid values are AES256 and aws:kms.
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (/). For example, testbucket/test1.json. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format arn:aws:s3:<Region>:<account-id>:accesspoint/<access-point-name>/object/<key>. For example, arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the algorithm to use when decrypting the source object (for example, AES256).
	SourceCustomerAlgorithm *string `json:"sourceCustomerAlgorithm,omitempty" tf:"source_customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	SourceCustomerKeyMd5 *string `json:"sourceCustomerKeyMd5,omitempty" tf:"source_customer_key_md5,omitempty"`

	// Version of the copied object in the source bucket.
	SourceVersionID *string `json:"sourceVersionId,omitempty" tf:"source_version_id,omitempty"`

	// Specifies the desired storage class for the object. Defaults to STANDARD.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are COPY and REPLACE.
	TaggingDirective *string `json:"taggingDirective,omitempty" tf:"tagging_directive,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Version ID of the newly created copy.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`

	// Specifies a target URL for website redirect.
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectCopyParameters struct {

	// Canned ACL to apply. Defaults to private. Valid values are private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, and bucket-owner-full-control. Conflicts with grant.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// Name of the bucket to put the file in.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control for further details.
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Specifies presentational information for the object. Read w3c content_disposition for further information.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read w3c content encoding for further information.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Language the content is in e.g., en-US or en-GB.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	// +kubebuilder:validation:Optional
	CopyIfMatch *string `json:"copyIfMatch,omitempty" tf:"copy_if_match,omitempty"`

	// Copies the object if it has been modified since the specified time, in RFC3339 format.
	// +kubebuilder:validation:Optional
	CopyIfModifiedSince *string `json:"copyIfModifiedSince,omitempty" tf:"copy_if_modified_since,omitempty"`

	// Copies the object if its entity tag (ETag) is different than the specified ETag.
	// +kubebuilder:validation:Optional
	CopyIfNoneMatch *string `json:"copyIfNoneMatch,omitempty" tf:"copy_if_none_match,omitempty"`

	// Copies the object if it hasn't been modified since the specified time, in RFC3339 format.
	// +kubebuilder:validation:Optional
	CopyIfUnmodifiedSince *string `json:"copyIfUnmodifiedSince,omitempty" tf:"copy_if_unmodified_since,omitempty"`

	// Specifies the algorithm to use to when encrypting the object (for example, AES256).
	// +kubebuilder:validation:Optional
	CustomerAlgorithm *string `json:"customerAlgorithm,omitempty" tf:"customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	// +kubebuilder:validation:Optional
	CustomerKeyMd5 *string `json:"customerKeyMd5,omitempty" tf:"customer_key_md5,omitempty"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting data. This value is used to store the object and then it is discarded; Amazon S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm header.
	// +kubebuilder:validation:Optional
	CustomerKeySecretRef *v1.SecretKeySelector `json:"customerKeySecretRef,omitempty" tf:"-"`

	// Account id of the expected destination bucket owner. If the destination bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Account id of the expected source bucket owner. If the source bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	// +kubebuilder:validation:Optional
	ExpectedSourceBucketOwner *string `json:"expectedSourceBucketOwner,omitempty" tf:"expected_source_bucket_owner,omitempty"`

	// Date and time at which the object is no longer cacheable, in RFC3339 format.
	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// Allow the object to be deleted by removing any legal hold on any object version. Default is false. This value should be set to true only if the bucket has S3 object lock enabled.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Configuration block for header grants. Documented below. Conflicts with acl.
	// +kubebuilder:validation:Optional
	Grant []ObjectCopyGrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// Specifies the AWS KMS Encryption Context to use for object encryption. The value is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
	// +kubebuilder:validation:Optional
	KMSEncryptionContextSecretRef *v1.SecretKeySelector `json:"kmsEncryptionContextSecretRef,omitempty" tf:"-"`

	// Specifies the AWS KMS Key ARN to use for object encryption. This value is a fully qualified ARN of the KMS Key. If using aws_kms_key, use the exported arn attribute: kms_key_id = aws_kms_key.foo.arn
	// +kubebuilder:validation:Optional
	KMSKeyIDSecretRef *v1.SecretKeySelector `json:"kmsKeyIdSecretRef,omitempty" tf:"-"`

	// Name of the object once it is in the bucket.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of keys/values to provision metadata (will be automatically prefixed by x-amz-meta-, note that only lowercase label are currently supported by the AWS Go API).
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are COPY and REPLACE.
	// +kubebuilder:validation:Optional
	MetadataDirective *string `json:"metadataDirective,omitempty" tf:"metadata_directive,omitempty"`

	// The legal hold status that you want to apply to the specified object. Valid values are ON and OFF.
	// +kubebuilder:validation:Optional
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// Object lock retention mode that you want to apply to this object. Valid values are GOVERNANCE and COMPLIANCE.
	// +kubebuilder:validation:Optional
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// Date and time, in RFC3339 format, when this object's object lock will expire.
	// +kubebuilder:validation:Optional
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see Downloading Objects in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html) in the Amazon S3 Developer Guide. If included, the only valid value is requester.
	// +kubebuilder:validation:Optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// Specifies server-side encryption of the object in S3. Valid values are AES256 and aws:kms.
	// +kubebuilder:validation:Optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (/). For example, testbucket/test1.json. For objects accessed through access points, specify the ARN of the object as accessed through the access point, in the format arn:aws:s3:<Region>:<account-id>:accesspoint/<access-point-name>/object/<key>. For example, arn:aws:s3:us-west-2:9999912999:accesspoint/my-access-point/object/testbucket/test1.json.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the algorithm to use when decrypting the source object (for example, AES256).
	// +kubebuilder:validation:Optional
	SourceCustomerAlgorithm *string `json:"sourceCustomerAlgorithm,omitempty" tf:"source_customer_algorithm,omitempty"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. Amazon S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error.
	// +kubebuilder:validation:Optional
	SourceCustomerKeyMd5 *string `json:"sourceCustomerKeyMd5,omitempty" tf:"source_customer_key_md5,omitempty"`

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt the source object. The encryption key provided in this header must be one that was used when the source object was created.
	// +kubebuilder:validation:Optional
	SourceCustomerKeySecretRef *v1.SecretKeySelector `json:"sourceCustomerKeySecretRef,omitempty" tf:"-"`

	// Specifies the desired storage class for the object. Defaults to STANDARD.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are COPY and REPLACE.
	// +kubebuilder:validation:Optional
	TaggingDirective *string `json:"taggingDirective,omitempty" tf:"tagging_directive,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a target URL for website redirect.
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// ObjectCopySpec defines the desired state of ObjectCopy
type ObjectCopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectCopyParameters `json:"forProvider"`
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
	InitProvider ObjectCopyInitParameters `json:"initProvider,omitempty"`
}

// ObjectCopyStatus defines the observed state of ObjectCopy.
type ObjectCopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectCopy is the Schema for the ObjectCopys API. Provides a resource for copying an S3 object.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ObjectCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   ObjectCopySpec   `json:"spec"`
	Status ObjectCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectCopyList contains a list of ObjectCopys
type ObjectCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectCopy `json:"items"`
}

// Repository type metadata.
var (
	ObjectCopy_Kind             = "ObjectCopy"
	ObjectCopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectCopy_Kind}.String()
	ObjectCopy_KindAPIVersion   = ObjectCopy_Kind + "." + CRDGroupVersion.String()
	ObjectCopy_GroupVersionKind = CRDGroupVersion.WithKind(ObjectCopy_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectCopy{}, &ObjectCopyList{})
}
