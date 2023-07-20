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

type RoleAssociationInitParameters struct {

	// The AWS SSO group ids to be assigned the role given in role.
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The grafana role. Valid values can be found here.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The AWS SSO user ids to be assigned the role given in role.
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`

	// The workspace id.
	// +crossplane:generate:reference:type=Workspace
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

type RoleAssociationObservation struct {

	// The AWS SSO group ids to be assigned the role given in role.
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The grafana role. Valid values can be found here.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The AWS SSO user ids to be assigned the role given in role.
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`

	// The workspace id.
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type RoleAssociationParameters struct {

	// The AWS SSO group ids to be assigned the role given in role.
	GroupIds []*string `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// The grafana role. Valid values can be found here.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The AWS SSO user ids to be assigned the role given in role.
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`

	// The workspace id.
	// +crossplane:generate:reference:type=Workspace
	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`

	// Reference to a Workspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDRef *v1.Reference `json:"workspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace to populate workspaceId.
	// +kubebuilder:validation:Optional
	WorkspaceIDSelector *v1.Selector `json:"workspaceIdSelector,omitempty" tf:"-"`
}

// RoleAssociationSpec defines the desired state of RoleAssociation
type RoleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider RoleAssociationInitParameters `json:"initProvider,omitempty"`
}

// RoleAssociationStatus defines the observed state of RoleAssociation.
type RoleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssociation is the Schema for the RoleAssociations API. Provides an Amazon Managed Grafana workspace role association resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RoleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="%!s(MISSING) is a required parameter"
	Spec   RoleAssociationSpec   `json:"spec"`
	Status RoleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssociationList contains a list of RoleAssociations
type RoleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssociation `json:"items"`
}

// Repository type metadata.
var (
	RoleAssociation_Kind             = "RoleAssociation"
	RoleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAssociation_Kind}.String()
	RoleAssociation_KindAPIVersion   = RoleAssociation_Kind + "." + CRDGroupVersion.String()
	RoleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(RoleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAssociation{}, &RoleAssociationList{})
}
