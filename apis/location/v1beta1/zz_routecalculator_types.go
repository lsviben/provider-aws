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

type RouteCalculatorInitParameters struct {

	// Specifies the data provider of traffic and road network data.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// The optional description for the route calculator resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteCalculatorObservation struct {

	// The Amazon Resource Name (ARN) for the Route calculator resource. Use the ARN when you specify a resource across AWS.
	CalculatorArn *string `json:"calculatorArn,omitempty" tf:"calculator_arn,omitempty"`

	// The timestamp for when the route calculator resource was created in ISO 8601 format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Specifies the data provider of traffic and road network data.
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// The optional description for the route calculator resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The timestamp for when the route calculator resource was last update in ISO 8601.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type RouteCalculatorParameters struct {

	// Specifies the data provider of traffic and road network data.
	// +kubebuilder:validation:Optional
	DataSource *string `json:"dataSource,omitempty" tf:"data_source,omitempty"`

	// The optional description for the route calculator resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RouteCalculatorSpec defines the desired state of RouteCalculator
type RouteCalculatorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteCalculatorParameters `json:"forProvider"`
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
	InitProvider RouteCalculatorInitParameters `json:"initProvider,omitempty"`
}

// RouteCalculatorStatus defines the observed state of RouteCalculator.
type RouteCalculatorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteCalculatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteCalculator is the Schema for the RouteCalculators API. Provides a Location Service Route Calculator.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RouteCalculator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataSource) || (has(self.initProvider) && has(self.initProvider.dataSource))",message="spec.forProvider.dataSource is a required parameter"
	Spec   RouteCalculatorSpec   `json:"spec"`
	Status RouteCalculatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteCalculatorList contains a list of RouteCalculators
type RouteCalculatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteCalculator `json:"items"`
}

// Repository type metadata.
var (
	RouteCalculator_Kind             = "RouteCalculator"
	RouteCalculator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteCalculator_Kind}.String()
	RouteCalculator_KindAPIVersion   = RouteCalculator_Kind + "." + CRDGroupVersion.String()
	RouteCalculator_GroupVersionKind = CRDGroupVersion.WithKind(RouteCalculator_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteCalculator{}, &RouteCalculatorList{})
}
