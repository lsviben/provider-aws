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

type ConfigurationInitParameters_2 struct {

	// Authentication strategy associated with the configuration. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Broker configuration in XML format. See official docs for supported parameters and format of the XML.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of the configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationObservation_2 struct {

	// ARN of the configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Authentication strategy associated with the configuration. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Broker configuration in XML format. See official docs for supported parameters and format of the XML.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Description of the configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Unique ID that Amazon MQ generates for the configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Latest revision of the configuration.
	LatestRevision *float64 `json:"latestRevision,omitempty" tf:"latest_revision,omitempty"`

	// Name of the configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConfigurationParameters_2 struct {

	// Authentication strategy associated with the configuration. Valid values are simple and ldap. ldap is not supported for engine_type RabbitMQ.
	// +kubebuilder:validation:Optional
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy,omitempty"`

	// Broker configuration in XML format. See official docs for supported parameters and format of the XML.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Description of the configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Type of broker engine. Valid values are ActiveMQ and RabbitMQ.
	// +kubebuilder:validation:Optional
	EngineType *string `json:"engineType,omitempty" tf:"engine_type,omitempty"`

	// Version of the broker engine.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of the configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConfigurationSpec defines the desired state of Configuration
type ConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationParameters_2 `json:"forProvider"`
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
	InitProvider ConfigurationInitParameters_2 `json:"initProvider,omitempty"`
}

// ConfigurationStatus defines the observed state of Configuration.
type ConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Configuration is the Schema for the Configurations API. Provides an MQ configuration Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.data) || (has(self.initProvider) && has(self.initProvider.data))",message="spec.forProvider.data is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineType) || (has(self.initProvider) && has(self.initProvider.engineType))",message="spec.forProvider.engineType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engineVersion) || (has(self.initProvider) && has(self.initProvider.engineVersion))",message="spec.forProvider.engineVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ConfigurationSpec   `json:"spec"`
	Status ConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList contains a list of Configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Repository type metadata.
var (
	Configuration_Kind             = "Configuration"
	Configuration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Configuration_Kind}.String()
	Configuration_KindAPIVersion   = Configuration_Kind + "." + CRDGroupVersion.String()
	Configuration_GroupVersionKind = CRDGroupVersion.WithKind(Configuration_Kind)
)

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
