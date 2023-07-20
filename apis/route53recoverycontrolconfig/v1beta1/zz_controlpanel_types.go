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

type ControlPanelInitParameters struct {

	// ARN of the cluster in which this control panel will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	ClusterArnRef *v1.Reference `json:"clusterArnRef,omitempty" tf:"-"`

	ClusterArnSelector *v1.Selector `json:"clusterArnSelector,omitempty" tf:"-"`

	// Name describing the control panel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

type ControlPanelObservation struct {

	// ARN of the control panel.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the cluster in which this control panel will reside.
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Whether a control panel is default.
	DefaultControlPanel *bool `json:"defaultControlPanel,omitempty" tf:"default_control_panel,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name describing the control panel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number routing controls in a control panel.
	RoutingControlCount *float64 `json:"routingControlCount,omitempty" tf:"routing_control_count,omitempty"`

	// Status of control panel: PENDING when it is being created/updated, PENDING_DELETION when it is being deleted, and DEPLOYED otherwise.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ControlPanelParameters struct {

	// ARN of the cluster in which this control panel will reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53recoverycontrolconfig/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	ClusterArn *string `json:"clusterArn,omitempty" tf:"cluster_arn,omitempty"`

	// Reference to a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnRef *v1.Reference `json:"clusterArnRef,omitempty" tf:"-"`

	// Selector for a Cluster in route53recoverycontrolconfig to populate clusterArn.
	// +kubebuilder:validation:Optional
	ClusterArnSelector *v1.Selector `json:"clusterArnSelector,omitempty" tf:"-"`

	// Name describing the control panel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`
}

// ControlPanelSpec defines the desired state of ControlPanel
type ControlPanelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ControlPanelParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ControlPanelInitParameters `json:"initProvider,omitempty"`
}

// ControlPanelStatus defines the observed state of ControlPanel.
type ControlPanelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ControlPanelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPanel is the Schema for the ControlPanels API. Provides an AWS Route 53 Recovery Control Config Control Panel
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ControlPanel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	Spec   ControlPanelSpec   `json:"spec"`
	Status ControlPanelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPanelList contains a list of ControlPanels
type ControlPanelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPanel `json:"items"`
}

// Repository type metadata.
var (
	ControlPanel_Kind             = "ControlPanel"
	ControlPanel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ControlPanel_Kind}.String()
	ControlPanel_KindAPIVersion   = ControlPanel_Kind + "." + CRDGroupVersion.String()
	ControlPanel_GroupVersionKind = CRDGroupVersion.WithKind(ControlPanel_Kind)
)

func init() {
	SchemeBuilder.Register(&ControlPanel{}, &ControlPanelList{})
}
