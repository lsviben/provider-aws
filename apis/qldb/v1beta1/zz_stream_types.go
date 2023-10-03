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

type KinesisConfigurationInitParameters struct {

	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call. Default: true.
	AggregationEnabled *bool `json:"aggregationEnabled,omitempty" tf:"aggregation_enabled,omitempty"`
}

type KinesisConfigurationObservation struct {

	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call. Default: true.
	AggregationEnabled *bool `json:"aggregationEnabled,omitempty" tf:"aggregation_enabled,omitempty"`

	// The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type KinesisConfigurationParameters struct {

	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call. Default: true.
	// +kubebuilder:validation:Optional
	AggregationEnabled *bool `json:"aggregationEnabled,omitempty" tf:"aggregation_enabled,omitempty"`

	// The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

type StreamInitParameters struct {

	// The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".
	ExclusiveEndTime *string `json:"exclusiveEndTime,omitempty" tf:"exclusive_end_time,omitempty"`

	// The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".  This cannot be in the future and must be before exclusive_end_time.  If you provide a value that is before the ledger's CreationDateTime, QLDB effectively defaults it to the ledger's CreationDateTime.
	InclusiveStartTime *string `json:"inclusiveStartTime,omitempty" tf:"inclusive_start_time,omitempty"`

	// The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
	KinesisConfiguration []KinesisConfigurationInitParameters `json:"kinesisConfiguration,omitempty" tf:"kinesis_configuration,omitempty"`

	// The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the Amazon QLDB Developer Guide.
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StreamObservation struct {

	// The ARN of the QLDB Stream.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".
	ExclusiveEndTime *string `json:"exclusiveEndTime,omitempty" tf:"exclusive_end_time,omitempty"`

	// The ID of the QLDB Stream.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".  This cannot be in the future and must be before exclusive_end_time.  If you provide a value that is before the ledger's CreationDateTime, QLDB effectively defaults it to the ledger's CreationDateTime.
	InclusiveStartTime *string `json:"inclusiveStartTime,omitempty" tf:"inclusive_start_time,omitempty"`

	// The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
	KinesisConfiguration []KinesisConfigurationObservation `json:"kinesisConfiguration,omitempty" tf:"kinesis_configuration,omitempty"`

	// The name of the QLDB ledger.
	LedgerName *string `json:"ledgerName,omitempty" tf:"ledger_name,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the Amazon QLDB Developer Guide.
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StreamParameters struct {

	// The exclusive date and time that specifies when the stream ends. If you don't define this parameter, the stream runs indefinitely until you cancel it. It must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".
	// +kubebuilder:validation:Optional
	ExclusiveEndTime *string `json:"exclusiveEndTime,omitempty" tf:"exclusive_end_time,omitempty"`

	// The inclusive start date and time from which to start streaming journal data. This parameter must be in ISO 8601 date and time format and in Universal Coordinated Time (UTC). For example: "2019-06-13T21:36:34Z".  This cannot be in the future and must be before exclusive_end_time.  If you provide a value that is before the ledger's CreationDateTime, QLDB effectively defaults it to the ledger's CreationDateTime.
	// +kubebuilder:validation:Optional
	InclusiveStartTime *string `json:"inclusiveStartTime,omitempty" tf:"inclusive_start_time,omitempty"`

	// The configuration settings of the Kinesis Data Streams destination for your stream request. Documented below.
	// +kubebuilder:validation:Optional
	KinesisConfiguration []KinesisConfigurationParameters `json:"kinesisConfiguration,omitempty" tf:"kinesis_configuration,omitempty"`

	// The name of the QLDB ledger.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/qldb/v1beta1.Ledger
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	LedgerName *string `json:"ledgerName,omitempty" tf:"ledger_name,omitempty"`

	// Reference to a Ledger in qldb to populate ledgerName.
	// +kubebuilder:validation:Optional
	LedgerNameRef *v1.Reference `json:"ledgerNameRef,omitempty" tf:"-"`

	// Selector for a Ledger in qldb to populate ledgerName.
	// +kubebuilder:validation:Optional
	LedgerNameSelector *v1.Selector `json:"ledgerNameSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.
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

	// The name that you want to assign to the QLDB journal stream. User-defined names can help identify and indicate the purpose of a stream.  Your stream name must be unique among other active streams for a given ledger. Stream names have the same naming constraints as ledger names, as defined in the Amazon QLDB Developer Guide.
	// +kubebuilder:validation:Optional
	StreamName *string `json:"streamName,omitempty" tf:"stream_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StreamSpec defines the desired state of Stream
type StreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamParameters `json:"forProvider"`
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
	InitProvider StreamInitParameters `json:"initProvider,omitempty"`
}

// StreamStatus defines the observed state of Stream.
type StreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stream is the Schema for the Streams API. Provides a QLDB Stream resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.inclusiveStartTime) || (has(self.initProvider) && has(self.initProvider.inclusiveStartTime))",message="spec.forProvider.inclusiveStartTime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kinesisConfiguration) || (has(self.initProvider) && has(self.initProvider.kinesisConfiguration))",message="spec.forProvider.kinesisConfiguration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.streamName) || (has(self.initProvider) && has(self.initProvider.streamName))",message="spec.forProvider.streamName is a required parameter"
	Spec   StreamSpec   `json:"spec"`
	Status StreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamList contains a list of Streams
type StreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stream `json:"items"`
}

// Repository type metadata.
var (
	Stream_Kind             = "Stream"
	Stream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stream_Kind}.String()
	Stream_KindAPIVersion   = Stream_Kind + "." + CRDGroupVersion.String()
	Stream_GroupVersionKind = CRDGroupVersion.WithKind(Stream_Kind)
)

func init() {
	SchemeBuilder.Register(&Stream{}, &StreamList{})
}
