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

type DeploymentObservation struct {

	// The creation date of the deployment
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// The execution ARN to be used in lambda_permission's source_arn
	// when allowing API Gateway to invoke a Lambda function,
	// e.g., arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod
	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	// The ID of the deployment
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL to invoke the API pointing to the stage,
	// e.g., https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod
	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`
}

type DeploymentParameters struct {

	// Description of the deployment
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// REST API identifier.
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

	// Description to set on the stage managed by the stage_name argument.
	// +kubebuilder:validation:Optional
	StageDescription *string `json:"stageDescription,omitempty" tf:"stage_description,omitempty"`

	// Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the aws_api_gateway_stage resource instead to manage stages.
	// +kubebuilder:validation:Optional
	StageName *string `json:"stageName,omitempty" tf:"stage_name,omitempty"`

	// argument or explicit resource references using the resource . The triggers argument should be preferred over depends_on, since depends_on can only capture dependency ordering and will not cause the resource to recreate (redeploy the REST API) with upstream configuration changes.
	// +kubebuilder:validation:Optional
	Triggers map[string]*string `json:"triggers,omitempty" tf:"triggers,omitempty"`

	// Map to set on the stage managed by the stage_name argument.
	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Deployment is the Schema for the Deployments API. Manages an API Gateway REST Deployment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
