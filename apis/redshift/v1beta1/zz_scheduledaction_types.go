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

type PauseClusterInitParameters struct {

	// The identifier of the cluster to be paused.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
}

type PauseClusterObservation struct {

	// The identifier of the cluster to be paused.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
}

type PauseClusterParameters struct {

	// The identifier of the cluster to be paused.
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier" tf:"cluster_identifier,omitempty"`
}

type ResizeClusterInitParameters struct {

	// A boolean value indicating whether the resize operation is using the classic resize process. Default: false.
	Classic *bool `json:"classic,omitempty" tf:"classic,omitempty"`

	// The unique identifier for the cluster to resize.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The new cluster type for the specified cluster.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The new node type for the nodes you are adding.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The new number of nodes for the cluster.
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`
}

type ResizeClusterObservation struct {

	// A boolean value indicating whether the resize operation is using the classic resize process. Default: false.
	Classic *bool `json:"classic,omitempty" tf:"classic,omitempty"`

	// The unique identifier for the cluster to resize.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// The new cluster type for the specified cluster.
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The new node type for the nodes you are adding.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The new number of nodes for the cluster.
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`
}

type ResizeClusterParameters struct {

	// A boolean value indicating whether the resize operation is using the classic resize process. Default: false.
	// +kubebuilder:validation:Optional
	Classic *bool `json:"classic,omitempty" tf:"classic,omitempty"`

	// The unique identifier for the cluster to resize.
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier" tf:"cluster_identifier,omitempty"`

	// The new cluster type for the specified cluster.
	// +kubebuilder:validation:Optional
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// The new node type for the nodes you are adding.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The new number of nodes for the cluster.
	// +kubebuilder:validation:Optional
	NumberOfNodes *float64 `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`
}

type ResumeClusterInitParameters struct {

	// The identifier of the cluster to be resumed.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
}

type ResumeClusterObservation struct {

	// The identifier of the cluster to be resumed.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`
}

type ResumeClusterParameters struct {

	// The identifier of the cluster to be resumed.
	// +kubebuilder:validation:Optional
	ClusterIdentifier *string `json:"clusterIdentifier" tf:"cluster_identifier,omitempty"`
}

type ScheduledActionInitParameters struct {

	// The description of the scheduled action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable the scheduled action. Default is true .
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example at(2016-03-04T17:27:00) or cron(0 10 ? * MON *). See Scheduled Action for more information.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Target action. Documented below.
	TargetAction []TargetActionInitParameters `json:"targetAction,omitempty" tf:"target_action,omitempty"`
}

type ScheduledActionObservation struct {

	// The description of the scheduled action.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable the scheduled action. Default is true .
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The IAM role to assume to run the scheduled action.
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// The Redshift Scheduled Action name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example at(2016-03-04T17:27:00) or cron(0 10 ? * MON *). See Scheduled Action for more information.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Target action. Documented below.
	TargetAction []TargetActionObservation `json:"targetAction,omitempty" tf:"target_action,omitempty"`
}

type ScheduledActionParameters struct {

	// The description of the scheduled action.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to enable the scheduled action. Default is true .
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The IAM role to assume to run the scheduled action.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	IAMRole *string `json:"iamRole,omitempty" tf:"iam_role,omitempty"`

	// Reference to a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleRef *v1.Reference `json:"iamRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRole.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example at(2016-03-04T17:27:00) or cron(0 10 ? * MON *). See Scheduled Action for more information.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Target action. Documented below.
	// +kubebuilder:validation:Optional
	TargetAction []TargetActionParameters `json:"targetAction,omitempty" tf:"target_action,omitempty"`
}

type TargetActionInitParameters struct {

	// An action that runs a PauseCluster API operation. Documented below.
	PauseCluster []PauseClusterInitParameters `json:"pauseCluster,omitempty" tf:"pause_cluster,omitempty"`

	// An action that runs a ResizeCluster API operation. Documented below.
	ResizeCluster []ResizeClusterInitParameters `json:"resizeCluster,omitempty" tf:"resize_cluster,omitempty"`

	// An action that runs a ResumeCluster API operation. Documented below.
	ResumeCluster []ResumeClusterInitParameters `json:"resumeCluster,omitempty" tf:"resume_cluster,omitempty"`
}

type TargetActionObservation struct {

	// An action that runs a PauseCluster API operation. Documented below.
	PauseCluster []PauseClusterObservation `json:"pauseCluster,omitempty" tf:"pause_cluster,omitempty"`

	// An action that runs a ResizeCluster API operation. Documented below.
	ResizeCluster []ResizeClusterObservation `json:"resizeCluster,omitempty" tf:"resize_cluster,omitempty"`

	// An action that runs a ResumeCluster API operation. Documented below.
	ResumeCluster []ResumeClusterObservation `json:"resumeCluster,omitempty" tf:"resume_cluster,omitempty"`
}

type TargetActionParameters struct {

	// An action that runs a PauseCluster API operation. Documented below.
	// +kubebuilder:validation:Optional
	PauseCluster []PauseClusterParameters `json:"pauseCluster,omitempty" tf:"pause_cluster,omitempty"`

	// An action that runs a ResizeCluster API operation. Documented below.
	// +kubebuilder:validation:Optional
	ResizeCluster []ResizeClusterParameters `json:"resizeCluster,omitempty" tf:"resize_cluster,omitempty"`

	// An action that runs a ResumeCluster API operation. Documented below.
	// +kubebuilder:validation:Optional
	ResumeCluster []ResumeClusterParameters `json:"resumeCluster,omitempty" tf:"resume_cluster,omitempty"`
}

// ScheduledActionSpec defines the desired state of ScheduledAction
type ScheduledActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduledActionParameters `json:"forProvider"`
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
	InitProvider ScheduledActionInitParameters `json:"initProvider,omitempty"`
}

// ScheduledActionStatus defines the observed state of ScheduledAction.
type ScheduledActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduledActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledAction is the Schema for the ScheduledActions API. Provides a Redshift Scheduled Action resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ScheduledAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetAction) || (has(self.initProvider) && has(self.initProvider.targetAction))",message="spec.forProvider.targetAction is a required parameter"
	Spec   ScheduledActionSpec   `json:"spec"`
	Status ScheduledActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduledActionList contains a list of ScheduledActions
type ScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScheduledAction `json:"items"`
}

// Repository type metadata.
var (
	ScheduledAction_Kind             = "ScheduledAction"
	ScheduledAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScheduledAction_Kind}.String()
	ScheduledAction_KindAPIVersion   = ScheduledAction_Kind + "." + CRDGroupVersion.String()
	ScheduledAction_GroupVersionKind = CRDGroupVersion.WithKind(ScheduledAction_Kind)
)

func init() {
	SchemeBuilder.Register(&ScheduledAction{}, &ScheduledActionList{})
}
