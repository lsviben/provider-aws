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

type CloudwatchEncryptionInitParameters struct {

	// Encryption mode to use for CloudWatch data. Valid values: DISABLED, SSE-KMS. Default value: DISABLED.
	CloudwatchEncryptionMode *string `json:"cloudwatchEncryptionMode,omitempty" tf:"cloudwatch_encryption_mode,omitempty"`
}

type CloudwatchEncryptionObservation struct {

	// Encryption mode to use for CloudWatch data. Valid values: DISABLED, SSE-KMS. Default value: DISABLED.
	CloudwatchEncryptionMode *string `json:"cloudwatchEncryptionMode,omitempty" tf:"cloudwatch_encryption_mode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type CloudwatchEncryptionParameters struct {

	// Encryption mode to use for CloudWatch data. Valid values: DISABLED, SSE-KMS. Default value: DISABLED.
	// +kubebuilder:validation:Optional
	CloudwatchEncryptionMode *string `json:"cloudwatchEncryptionMode,omitempty" tf:"cloudwatch_encryption_mode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`
}

type EncryptionConfigurationInitParameters struct {
	CloudwatchEncryption []CloudwatchEncryptionInitParameters `json:"cloudwatchEncryption,omitempty" tf:"cloudwatch_encryption,omitempty"`

	JobBookmarksEncryption []JobBookmarksEncryptionInitParameters `json:"jobBookmarksEncryption,omitempty" tf:"job_bookmarks_encryption,omitempty"`

	// A s3_encryption  block as described below, which contains encryption configuration for S3 data.
	S3Encryption []S3EncryptionInitParameters `json:"s3Encryption,omitempty" tf:"s3_encryption,omitempty"`
}

type EncryptionConfigurationObservation struct {
	CloudwatchEncryption []CloudwatchEncryptionObservation `json:"cloudwatchEncryption,omitempty" tf:"cloudwatch_encryption,omitempty"`

	JobBookmarksEncryption []JobBookmarksEncryptionObservation `json:"jobBookmarksEncryption,omitempty" tf:"job_bookmarks_encryption,omitempty"`

	// A s3_encryption  block as described below, which contains encryption configuration for S3 data.
	S3Encryption []S3EncryptionObservation `json:"s3Encryption,omitempty" tf:"s3_encryption,omitempty"`
}

type EncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchEncryption []CloudwatchEncryptionParameters `json:"cloudwatchEncryption" tf:"cloudwatch_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	JobBookmarksEncryption []JobBookmarksEncryptionParameters `json:"jobBookmarksEncryption" tf:"job_bookmarks_encryption,omitempty"`

	// A s3_encryption  block as described below, which contains encryption configuration for S3 data.
	// +kubebuilder:validation:Optional
	S3Encryption []S3EncryptionParameters `json:"s3Encryption" tf:"s3_encryption,omitempty"`
}

type JobBookmarksEncryptionInitParameters struct {

	// Encryption mode to use for job bookmarks data. Valid values: CSE-KMS, DISABLED. Default value: DISABLED.
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode,omitempty" tf:"job_bookmarks_encryption_mode,omitempty"`
}

type JobBookmarksEncryptionObservation struct {

	// Encryption mode to use for job bookmarks data. Valid values: CSE-KMS, DISABLED. Default value: DISABLED.
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode,omitempty" tf:"job_bookmarks_encryption_mode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type JobBookmarksEncryptionParameters struct {

	// Encryption mode to use for job bookmarks data. Valid values: CSE-KMS, DISABLED. Default value: DISABLED.
	// +kubebuilder:validation:Optional
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode,omitempty" tf:"job_bookmarks_encryption_mode,omitempty"`

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`
}

type S3EncryptionInitParameters struct {

	// Encryption mode to use for S3 data. Valid values: DISABLED, SSE-KMS, SSE-S3. Default value: DISABLED.
	S3EncryptionMode *string `json:"s3EncryptionMode,omitempty" tf:"s3_encryption_mode,omitempty"`
}

type S3EncryptionObservation struct {

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Encryption mode to use for S3 data. Valid values: DISABLED, SSE-KMS, SSE-S3. Default value: DISABLED.
	S3EncryptionMode *string `json:"s3EncryptionMode,omitempty" tf:"s3_encryption_mode,omitempty"`
}

type S3EncryptionParameters struct {

	// Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Encryption mode to use for S3 data. Valid values: DISABLED, SSE-KMS, SSE-S3. Default value: DISABLED.
	// +kubebuilder:validation:Optional
	S3EncryptionMode *string `json:"s3EncryptionMode,omitempty" tf:"s3_encryption_mode,omitempty"`
}

type SecurityConfigurationInitParameters struct {

	// –  Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration []EncryptionConfigurationInitParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`
}

type SecurityConfigurationObservation struct {

	// –  Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration []EncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Glue security configuration name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityConfigurationParameters struct {

	// –  Configuration block containing encryption configuration. Detailed below.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SecurityConfigurationSpec defines the desired state of SecurityConfiguration
type SecurityConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityConfigurationParameters `json:"forProvider"`
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
	InitProvider SecurityConfigurationInitParameters `json:"initProvider,omitempty"`
}

// SecurityConfigurationStatus defines the observed state of SecurityConfiguration.
type SecurityConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfiguration is the Schema for the SecurityConfigurations API. Manages a Glue Security Configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.encryptionConfiguration) || (has(self.initProvider) && has(self.initProvider.encryptionConfiguration))",message="spec.forProvider.encryptionConfiguration is a required parameter"
	Spec   SecurityConfigurationSpec   `json:"spec"`
	Status SecurityConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfigurationList contains a list of SecurityConfigurations
type SecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityConfiguration `json:"items"`
}

// Repository type metadata.
var (
	SecurityConfiguration_Kind             = "SecurityConfiguration"
	SecurityConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityConfiguration_Kind}.String()
	SecurityConfiguration_KindAPIVersion   = SecurityConfiguration_Kind + "." + CRDGroupVersion.String()
	SecurityConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(SecurityConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityConfiguration{}, &SecurityConfigurationList{})
}
