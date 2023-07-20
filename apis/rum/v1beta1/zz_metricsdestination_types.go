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

type MetricsDestinationInitParameters struct {

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rum/v1beta1.AppMonitor
	AppMonitorName *string `json:"appMonitorName,omitempty" tf:"app_monitor_name,omitempty"`

	AppMonitorNameRef *v1.Reference `json:"appMonitorNameRef,omitempty" tf:"-"`

	AppMonitorNameSelector *v1.Selector `json:"appMonitorNameSelector,omitempty" tf:"-"`

	// Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type MetricsDestinationObservation struct {

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	AppMonitorName *string `json:"appMonitorName,omitempty" tf:"app_monitor_name,omitempty"`

	// Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MetricsDestinationParameters struct {

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rum/v1beta1.AppMonitor
	AppMonitorName *string `json:"appMonitorName,omitempty" tf:"app_monitor_name,omitempty"`

	// Reference to a AppMonitor in rum to populate appMonitorName.
	// +kubebuilder:validation:Optional
	AppMonitorNameRef *v1.Reference `json:"appMonitorNameRef,omitempty" tf:"-"`

	// Selector for a AppMonitor in rum to populate appMonitorName.
	// +kubebuilder:validation:Optional
	AppMonitorNameSelector *v1.Selector `json:"appMonitorNameSelector,omitempty" tf:"-"`

	// Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`

	// This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// MetricsDestinationSpec defines the desired state of MetricsDestination
type MetricsDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricsDestinationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider MetricsDestinationInitParameters `json:"initProvider,omitempty"`
}

// MetricsDestinationStatus defines the observed state of MetricsDestination.
type MetricsDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricsDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MetricsDestination is the Schema for the MetricsDestinations API. Provides a CloudWatch RUM Metrics Destination resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MetricsDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || has(self.initProvider.destination)",message="%!s(MISSING) is a required parameter"
	Spec   MetricsDestinationSpec   `json:"spec"`
	Status MetricsDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricsDestinationList contains a list of MetricsDestinations
type MetricsDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricsDestination `json:"items"`
}

// Repository type metadata.
var (
	MetricsDestination_Kind             = "MetricsDestination"
	MetricsDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricsDestination_Kind}.String()
	MetricsDestination_KindAPIVersion   = MetricsDestination_Kind + "." + CRDGroupVersion.String()
	MetricsDestination_GroupVersionKind = CRDGroupVersion.WithKind(MetricsDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricsDestination{}, &MetricsDestinationList{})
}
