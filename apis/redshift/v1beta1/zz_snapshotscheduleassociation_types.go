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

type SnapshotScheduleAssociationInitParameters struct {

	// The cluster identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The snapshot schedule identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SnapshotSchedule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`

	ScheduleIdentifierRef *v1.Reference `json:"scheduleIdentifierRef,omitempty" tf:"-"`

	ScheduleIdentifierSelector *v1.Selector `json:"scheduleIdentifierSelector,omitempty" tf:"-"`
}

type SnapshotScheduleAssociationObservation struct {

	// The cluster identifier.
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The snapshot schedule identifier.
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`
}

type SnapshotScheduleAssociationParameters struct {

	// The cluster identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier,omitempty"`

	// Reference to a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierRef *v1.Reference `json:"clusterIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster in redshift to populate clusterIdentifier.
	// +kubebuilder:validation:Optional
	ClusterIdentifierSelector *v1.Selector `json:"clusterIdentifierSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The snapshot schedule identifier.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/redshift/v1beta1.SnapshotSchedule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ScheduleIdentifier *string `json:"scheduleIdentifier,omitempty" tf:"schedule_identifier,omitempty"`

	// Reference to a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierRef *v1.Reference `json:"scheduleIdentifierRef,omitempty" tf:"-"`

	// Selector for a SnapshotSchedule in redshift to populate scheduleIdentifier.
	// +kubebuilder:validation:Optional
	ScheduleIdentifierSelector *v1.Selector `json:"scheduleIdentifierSelector,omitempty" tf:"-"`
}

// SnapshotScheduleAssociationSpec defines the desired state of SnapshotScheduleAssociation
type SnapshotScheduleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotScheduleAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SnapshotScheduleAssociationInitParameters `json:"initProvider,omitempty"`
}

// SnapshotScheduleAssociationStatus defines the observed state of SnapshotScheduleAssociation.
type SnapshotScheduleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotScheduleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleAssociation is the Schema for the SnapshotScheduleAssociations API. Provides an Association Redshift Cluster and Snapshot Schedule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SnapshotScheduleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotScheduleAssociationSpec   `json:"spec"`
	Status            SnapshotScheduleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotScheduleAssociationList contains a list of SnapshotScheduleAssociations
type SnapshotScheduleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotScheduleAssociation `json:"items"`
}

// Repository type metadata.
var (
	SnapshotScheduleAssociation_Kind             = "SnapshotScheduleAssociation"
	SnapshotScheduleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SnapshotScheduleAssociation_Kind}.String()
	SnapshotScheduleAssociation_KindAPIVersion   = SnapshotScheduleAssociation_Kind + "." + CRDGroupVersion.String()
	SnapshotScheduleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SnapshotScheduleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SnapshotScheduleAssociation{}, &SnapshotScheduleAssociationList{})
}
