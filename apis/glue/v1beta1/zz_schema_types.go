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

type SchemaInitParameters struct {

	// The compatibility mode of the schema. Values values are: NONE, DISABLED, BACKWARD, BACKWARD_ALL, FORWARD, FORWARD_ALL, FULL, and FULL_ALL.
	Compatibility *string `json:"compatibility,omitempty" tf:"compatibility,omitempty"`

	// The data format of the schema definition. Valid values are AVRO, JSON and PROTOBUF.
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// –  A description of the schema.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The schema definition using the data_format setting for schema_name.
	SchemaDefinition *string `json:"schemaDefinition,omitempty" tf:"schema_definition,omitempty"`

	// –  The Name of the schema.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SchemaObservation struct {

	// Amazon Resource Name (ARN) of the schema.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The compatibility mode of the schema. Values values are: NONE, DISABLED, BACKWARD, BACKWARD_ALL, FORWARD, FORWARD_ALL, FULL, and FULL_ALL.
	Compatibility *string `json:"compatibility,omitempty" tf:"compatibility,omitempty"`

	// The data format of the schema definition. Valid values are AVRO, JSON and PROTOBUF.
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// –  A description of the schema.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Amazon Resource Name (ARN) of the schema.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion *float64 `json:"latestSchemaVersion,omitempty" tf:"latest_schema_version,omitempty"`

	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion *float64 `json:"nextSchemaVersion,omitempty" tf:"next_schema_version,omitempty"`

	// The ARN of the Glue Registry to create the schema in.
	RegistryArn *string `json:"registryArn,omitempty" tf:"registry_arn,omitempty"`

	// The name of the Glue Registry.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// The version number of the checkpoint (the last time the compatibility mode was changed).
	SchemaCheckpoint *float64 `json:"schemaCheckpoint,omitempty" tf:"schema_checkpoint,omitempty"`

	// The schema definition using the data_format setting for schema_name.
	SchemaDefinition *string `json:"schemaDefinition,omitempty" tf:"schema_definition,omitempty"`

	// –  The Name of the schema.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SchemaParameters struct {

	// The compatibility mode of the schema. Values values are: NONE, DISABLED, BACKWARD, BACKWARD_ALL, FORWARD, FORWARD_ALL, FULL, and FULL_ALL.
	// +kubebuilder:validation:Optional
	Compatibility *string `json:"compatibility,omitempty" tf:"compatibility,omitempty"`

	// The data format of the schema definition. Valid values are AVRO, JSON and PROTOBUF.
	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// –  A description of the schema.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the Glue Registry to create the schema in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.Registry
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RegistryArn *string `json:"registryArn,omitempty" tf:"registry_arn,omitempty"`

	// Reference to a Registry in glue to populate registryArn.
	// +kubebuilder:validation:Optional
	RegistryArnRef *v1.Reference `json:"registryArnRef,omitempty" tf:"-"`

	// Selector for a Registry in glue to populate registryArn.
	// +kubebuilder:validation:Optional
	RegistryArnSelector *v1.Selector `json:"registryArnSelector,omitempty" tf:"-"`

	// The schema definition using the data_format setting for schema_name.
	// +kubebuilder:validation:Optional
	SchemaDefinition *string `json:"schemaDefinition,omitempty" tf:"schema_definition,omitempty"`

	// –  The Name of the schema.
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SchemaSpec defines the desired state of Schema
type SchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchemaParameters `json:"forProvider"`
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
	InitProvider SchemaInitParameters `json:"initProvider,omitempty"`
}

// SchemaStatus defines the observed state of Schema.
type SchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is the Schema for the Schemas API. Provides a Glue Schema resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Schema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.compatibility) || (has(self.initProvider) && has(self.initProvider.compatibility))",message="spec.forProvider.compatibility is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataFormat) || (has(self.initProvider) && has(self.initProvider.dataFormat))",message="spec.forProvider.dataFormat is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemaDefinition) || (has(self.initProvider) && has(self.initProvider.schemaDefinition))",message="spec.forProvider.schemaDefinition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schemaName) || (has(self.initProvider) && has(self.initProvider.schemaName))",message="spec.forProvider.schemaName is a required parameter"
	Spec   SchemaSpec   `json:"spec"`
	Status SchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaList contains a list of Schemas
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schema `json:"items"`
}

// Repository type metadata.
var (
	Schema_Kind             = "Schema"
	Schema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schema_Kind}.String()
	Schema_KindAPIVersion   = Schema_Kind + "." + CRDGroupVersion.String()
	Schema_GroupVersionKind = CRDGroupVersion.WithKind(Schema_Kind)
)

func init() {
	SchemeBuilder.Register(&Schema{}, &SchemaList{})
}
