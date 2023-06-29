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

type ScriptObservation struct {

	// GameLift Script ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// GameLift Script ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the script
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Information indicating where your game script files are stored. See below.
	StorageLocation []ScriptStorageLocationObservation `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Version that is associated with this script.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`
}

type ScriptParameters struct {

	// Name of the script
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Information indicating where your game script files are stored. See below.
	// +kubebuilder:validation:Optional
	StorageLocation []ScriptStorageLocationParameters `json:"storageLocation,omitempty" tf:"storage_location,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Version that is associated with this script.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// A data object containing your Realtime scripts and dependencies as a zip  file. The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	// +kubebuilder:validation:Optional
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`
}

type ScriptStorageLocationObservation struct {

	// Name of your S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Name of the zip file containing your script files.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type ScriptStorageLocationParameters struct {

	// Name of your S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Name of the zip file containing your script files.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Object
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("key",false)
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Reference to a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeyRef *v1.Reference `json:"keyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate key.
	// +kubebuilder:validation:Optional
	KeySelector *v1.Selector `json:"keySelector,omitempty" tf:"-"`

	// A specific version of the file. If not set, the latest version of the file is retrieved.
	// +kubebuilder:validation:Optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`

	// ARN of the access role that allows Amazon GameLift to access your S3 bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`
}

// ScriptSpec defines the desired state of Script
type ScriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScriptParameters `json:"forProvider"`
}

// ScriptStatus defines the observed state of Script.
type ScriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Script is the Schema for the Scripts API. Provides a GameLift Script resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Script struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ScriptSpec   `json:"spec"`
	Status ScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScriptList contains a list of Scripts
type ScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Script `json:"items"`
}

// Repository type metadata.
var (
	Script_Kind             = "Script"
	Script_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Script_Kind}.String()
	Script_KindAPIVersion   = Script_Kind + "." + CRDGroupVersion.String()
	Script_GroupVersionKind = CRDGroupVersion.WithKind(Script_Kind)
)

func init() {
	SchemeBuilder.Register(&Script{}, &ScriptList{})
}
