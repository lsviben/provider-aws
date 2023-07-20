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

type DetailsInitParameters struct {

	// The name of the Multi-Region Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point. You can enable the configuration options in any combination. See Public Access Block Configuration below for more details.
	PublicAccessBlock []PublicAccessBlockInitParameters `json:"publicAccessBlock,omitempty" tf:"public_access_block,omitempty"`

	// The Region configuration block to specify the bucket associated with the Multi-Region Access Point. See Region Configuration below for more details.
	Region []RegionInitParameters `json:"region,omitempty" tf:"region,omitempty"`
}

type DetailsObservation struct {

	// The name of the Multi-Region Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point. You can enable the configuration options in any combination. See Public Access Block Configuration below for more details.
	PublicAccessBlock []PublicAccessBlockObservation `json:"publicAccessBlock,omitempty" tf:"public_access_block,omitempty"`

	// The Region configuration block to specify the bucket associated with the Multi-Region Access Point. See Region Configuration below for more details.
	Region []RegionObservation `json:"region,omitempty" tf:"region,omitempty"`
}

type DetailsParameters struct {

	// The name of the Multi-Region Access Point.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Configuration block to manage the PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point. You can enable the configuration options in any combination. See Public Access Block Configuration below for more details.
	PublicAccessBlock []PublicAccessBlockParameters `json:"publicAccessBlock,omitempty" tf:"public_access_block,omitempty"`

	// The Region configuration block to specify the bucket associated with the Multi-Region Access Point. See Region Configuration below for more details.
	Region []RegionParameters `json:"region,omitempty" tf:"region,omitempty"`
}

type MultiRegionAccessPointInitParameters struct {

	// The AWS account ID for the owner of the buckets for which you want to create a Multi-Region Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the Multi-Region Access Point. See Details Configuration Block below for more details
	Details []DetailsInitParameters `json:"details,omitempty" tf:"details,omitempty"`

	// The Region configuration block to specify the bucket associated with the Multi-Region Access Point. See Region Configuration below for more details.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type MultiRegionAccessPointObservation struct {

	// The AWS account ID for the owner of the buckets for which you want to create a Multi-Region Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The alias for the Multi-Region Access Point.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Amazon Resource Name (ARN) of the Multi-Region Access Point.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A configuration block containing details about the Multi-Region Access Point. See Details Configuration Block below for more details
	Details []DetailsObservation `json:"details,omitempty" tf:"details,omitempty"`

	// The DNS domain name of the S3 Multi-Region Access Point in the format alias.accesspoint.s3-global.amazonaws.com. For more information, see the documentation on Multi-Region Access Point Requests.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The AWS account ID and access point name separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current status of the Multi-Region Access Point. One of: READY, INCONSISTENT_ACROSS_REGIONS, CREATING, PARTIALLY_CREATED, PARTIALLY_DELETED, DELETING.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type MultiRegionAccessPointParameters struct {

	// The AWS account ID for the owner of the buckets for which you want to create a Multi-Region Access Point.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A configuration block containing details about the Multi-Region Access Point. See Details Configuration Block below for more details
	Details []DetailsParameters `json:"details,omitempty" tf:"details,omitempty"`

	// The Region configuration block to specify the bucket associated with the Multi-Region Access Point. See Region Configuration below for more details.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type PublicAccessBlockInitParameters struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type PublicAccessBlockObservation struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type PublicAccessBlockParameters struct {

	// Whether Amazon S3 should block public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect existing policies or ACLs. When set to true causes the following behavior:
	BlockPublicAcls *bool `json:"blockPublicAcls,omitempty" tf:"block_public_acls,omitempty"`

	// Whether Amazon S3 should block public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect existing bucket policies. When set to true causes Amazon S3 to:
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// Whether Amazon S3 should ignore public ACLs for buckets in this account. Defaults to true. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to true causes Amazon S3 to:
	IgnorePublicAcls *bool `json:"ignorePublicAcls,omitempty" tf:"ignore_public_acls,omitempty"`

	// Whether Amazon S3 should restrict public bucket policies for buckets in this account. Defaults to true. Enabling this setting does not affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked. When set to true:
	RestrictPublicBuckets *bool `json:"restrictPublicBuckets,omitempty" tf:"restrict_public_buckets,omitempty"`
}

type RegionInitParameters struct {

	// The name of the associated bucket for the Region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`
}

type RegionObservation struct {

	// The name of the associated bucket for the Region.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`
}

type RegionParameters struct {

	// The name of the associated bucket for the Region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`
}

// MultiRegionAccessPointSpec defines the desired state of MultiRegionAccessPoint
type MultiRegionAccessPointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MultiRegionAccessPointParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider MultiRegionAccessPointInitParameters `json:"initProvider,omitempty"`
}

// MultiRegionAccessPointStatus defines the observed state of MultiRegionAccessPoint.
type MultiRegionAccessPointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MultiRegionAccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MultiRegionAccessPoint is the Schema for the MultiRegionAccessPoints API. Provides a resource to manage an S3 Multi-Region Access Point associated with specified buckets.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MultiRegionAccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.details) || has(self.initProvider.details)",message="%!s(MISSING) is a required parameter"
	Spec   MultiRegionAccessPointSpec   `json:"spec"`
	Status MultiRegionAccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MultiRegionAccessPointList contains a list of MultiRegionAccessPoints
type MultiRegionAccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultiRegionAccessPoint `json:"items"`
}

// Repository type metadata.
var (
	MultiRegionAccessPoint_Kind             = "MultiRegionAccessPoint"
	MultiRegionAccessPoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MultiRegionAccessPoint_Kind}.String()
	MultiRegionAccessPoint_KindAPIVersion   = MultiRegionAccessPoint_Kind + "." + CRDGroupVersion.String()
	MultiRegionAccessPoint_GroupVersionKind = CRDGroupVersion.WithKind(MultiRegionAccessPoint_Kind)
)

func init() {
	SchemeBuilder.Register(&MultiRegionAccessPoint{}, &MultiRegionAccessPointList{})
}
