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

type AccessLogSettingsObservation struct {
}

type AccessLogSettingsParameters struct {

	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-. Automatically removes trailing :* if present.
	// +kubebuilder:validation:Required
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`

	// The formatting and values recorded in the logs.
	// For more information on configuring the log format rules visit the AWS documentation
	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`
}

type CanarySettingsObservation struct {
}

type CanarySettingsParameters struct {

	// The percent 0.0 - 100.0 of traffic to divert to the canary deployment.
	// +kubebuilder:validation:Optional
	PercentTraffic *float64 `json:"percentTraffic,omitempty" tf:"percent_traffic,omitempty"`

	// A map of overridden stage variables (including new variables) for the canary deployment.
	// +kubebuilder:validation:Optional
	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty" tf:"stage_variable_overrides,omitempty"`

	// Whether the canary deployment uses the stage cache. Defaults to false.
	// +kubebuilder:validation:Optional
	UseStageCache *bool `json:"useStageCache,omitempty" tf:"use_stage_cache,omitempty"`
}

type StageObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The execution ARN to be used in lambda_permission's source_arn
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// The ID of the stage
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL to invoke the API pointing to the stage,
	// e.g., https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod
	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ARN of the WebAcl associated with the Stage.
	WebACLArn *string `json:"webAclArn,omitempty" tf:"web_acl_arn,omitempty"`
}

type StageParameters struct {

	// Enables access logs for the API stage. See Access Log Settings below.
	// +kubebuilder:validation:Optional
	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// Specifies whether a cache cluster is enabled for the stage
	// +kubebuilder:validation:Optional
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`

	// The size of the cache cluster for the stage, if enabled. Allowed values include 0.5, 1.6, 6.1, 13.5, 28.4, 58.2, 118 and 237.
	// +kubebuilder:validation:Optional
	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`

	// Configuration settings of a canary deployment. See Canary Settings below.
	// +kubebuilder:validation:Optional
	CanarySettings []CanarySettingsParameters `json:"canarySettings,omitempty" tf:"canary_settings,omitempty"`

	// The identifier of a client certificate for the stage.
	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// The ID of the deployment that the stage points to
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Deployment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Reference to a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDRef *v1.Reference `json:"deploymentIdRef,omitempty" tf:"-"`

	// Selector for a Deployment in apigateway to populate deploymentId.
	// +kubebuilder:validation:Optional
	DeploymentIDSelector *v1.Selector `json:"deploymentIdSelector,omitempty" tf:"-"`

	// The description of the stage.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The version of the associated API documentation
	// +kubebuilder:validation:Optional
	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the associated REST API
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// The name of the stage
	// +kubebuilder:validation:Required
	StageName *string `json:"stageName" tf:"stage_name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map that defines the stage variables
	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// Whether active tracing with X-ray is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

// StageSpec defines the desired state of Stage
type StageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StageParameters `json:"forProvider"`
}

// StageStatus defines the observed state of Stage.
type StageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stage is the Schema for the Stages API. Manages an API Gateway Stage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StageSpec   `json:"spec"`
	Status            StageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stages
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stage `json:"items"`
}

// Repository type metadata.
var (
	Stage_Kind             = "Stage"
	Stage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stage_Kind}.String()
	Stage_KindAPIVersion   = Stage_Kind + "." + CRDGroupVersion.String()
	Stage_GroupVersionKind = CRDGroupVersion.WithKind(Stage_Kind)
)

func init() {
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
