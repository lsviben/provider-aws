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

type DeploymentConfigInitParameters struct {

	// The compute platform can be Server, Lambda, or ECS. Default is Server.
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// A minimum_healthy_hosts block. Required for Server compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts []MinimumHealthyHostsInitParameters `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts,omitempty"`

	// A traffic_routing_config block. Traffic Routing Config is documented below.
	TrafficRoutingConfig []TrafficRoutingConfigInitParameters `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config,omitempty"`
}

type DeploymentConfigObservation struct {

	// The compute platform can be Server, Lambda, or ECS. Default is Server.
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// The AWS Assigned deployment config id
	DeploymentConfigID *string `json:"deploymentConfigId,omitempty" tf:"deployment_config_id,omitempty"`

	// The deployment group's config name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A minimum_healthy_hosts block. Required for Server compute platform. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts []MinimumHealthyHostsObservation `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts,omitempty"`

	// A traffic_routing_config block. Traffic Routing Config is documented below.
	TrafficRoutingConfig []TrafficRoutingConfigObservation `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config,omitempty"`
}

type DeploymentConfigParameters struct {

	// The compute platform can be Server, Lambda, or ECS. Default is Server.
	// +kubebuilder:validation:Optional
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`

	// A minimum_healthy_hosts block. Required for Server compute platform. Minimum Healthy Hosts are documented below.
	// +kubebuilder:validation:Optional
	MinimumHealthyHosts []MinimumHealthyHostsParameters `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A traffic_routing_config block. Traffic Routing Config is documented below.
	// +kubebuilder:validation:Optional
	TrafficRoutingConfig []TrafficRoutingConfigParameters `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config,omitempty"`
}

type MinimumHealthyHostsInitParameters struct {

	// The type can either be FLEET_PERCENT or HOST_COUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value when the type is FLEET_PERCENT represents the minimum number of healthy instances as
	// a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
	// deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	// When the type is HOST_COUNT, the value represents the minimum number of healthy instances as an absolute value.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type MinimumHealthyHostsObservation struct {

	// The type can either be FLEET_PERCENT or HOST_COUNT.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value when the type is FLEET_PERCENT represents the minimum number of healthy instances as
	// a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
	// deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	// When the type is HOST_COUNT, the value represents the minimum number of healthy instances as an absolute value.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type MinimumHealthyHostsParameters struct {

	// The type can either be FLEET_PERCENT or HOST_COUNT.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value when the type is FLEET_PERCENT represents the minimum number of healthy instances as
	// a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
	// deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
	// When the type is HOST_COUNT, the value represents the minimum number of healthy instances as an absolute value.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type TimeBasedCanaryInitParameters struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TimeBasedCanaryObservation struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TimeBasedCanaryParameters struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	// +kubebuilder:validation:Optional
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TimeBasedLinearInitParameters struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TimeBasedLinearObservation struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TimeBasedLinearParameters struct {

	// The number of minutes between the first and second traffic shifts of a TimeBasedCanary deployment.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The percentage of traffic to shift in the first increment of a TimeBasedCanary deployment.
	// +kubebuilder:validation:Optional
	Percentage *float64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type TrafficRoutingConfigInitParameters struct {

	// The time based canary configuration information. If type is TimeBasedLinear, use time_based_linear instead.
	TimeBasedCanary []TimeBasedCanaryInitParameters `json:"timeBasedCanary,omitempty" tf:"time_based_canary,omitempty"`

	// The time based linear configuration information. If type is TimeBasedCanary, use time_based_canary instead.
	TimeBasedLinear []TimeBasedLinearInitParameters `json:"timeBasedLinear,omitempty" tf:"time_based_linear,omitempty"`

	// Type of traffic routing config. One of TimeBasedCanary, TimeBasedLinear, AllAtOnce.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrafficRoutingConfigObservation struct {

	// The time based canary configuration information. If type is TimeBasedLinear, use time_based_linear instead.
	TimeBasedCanary []TimeBasedCanaryObservation `json:"timeBasedCanary,omitempty" tf:"time_based_canary,omitempty"`

	// The time based linear configuration information. If type is TimeBasedCanary, use time_based_canary instead.
	TimeBasedLinear []TimeBasedLinearObservation `json:"timeBasedLinear,omitempty" tf:"time_based_linear,omitempty"`

	// Type of traffic routing config. One of TimeBasedCanary, TimeBasedLinear, AllAtOnce.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TrafficRoutingConfigParameters struct {

	// The time based canary configuration information. If type is TimeBasedLinear, use time_based_linear instead.
	// +kubebuilder:validation:Optional
	TimeBasedCanary []TimeBasedCanaryParameters `json:"timeBasedCanary,omitempty" tf:"time_based_canary,omitempty"`

	// The time based linear configuration information. If type is TimeBasedCanary, use time_based_canary instead.
	// +kubebuilder:validation:Optional
	TimeBasedLinear []TimeBasedLinearParameters `json:"timeBasedLinear,omitempty" tf:"time_based_linear,omitempty"`

	// Type of traffic routing config. One of TimeBasedCanary, TimeBasedLinear, AllAtOnce.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DeploymentConfigSpec defines the desired state of DeploymentConfig
type DeploymentConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentConfigParameters `json:"forProvider"`
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
	InitProvider DeploymentConfigInitParameters `json:"initProvider,omitempty"`
}

// DeploymentConfigStatus defines the observed state of DeploymentConfig.
type DeploymentConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentConfig is the Schema for the DeploymentConfigs API. Provides a CodeDeploy deployment config.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DeploymentConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentConfigSpec   `json:"spec"`
	Status            DeploymentConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentConfigList contains a list of DeploymentConfigs
type DeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentConfig `json:"items"`
}

// Repository type metadata.
var (
	DeploymentConfig_Kind             = "DeploymentConfig"
	DeploymentConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeploymentConfig_Kind}.String()
	DeploymentConfig_KindAPIVersion   = DeploymentConfig_Kind + "." + CRDGroupVersion.String()
	DeploymentConfig_GroupVersionKind = CRDGroupVersion.WithKind(DeploymentConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&DeploymentConfig{}, &DeploymentConfigList{})
}
