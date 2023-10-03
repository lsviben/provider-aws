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

type InputSecurityGroupInitParameters struct {

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whitelist rules. See Whitelist Rules for more details.
	WhitelistRules []WhitelistRulesInitParameters `json:"whitelistRules,omitempty" tf:"whitelist_rules,omitempty"`
}

type InputSecurityGroupObservation struct {

	// ARN of the InputSecurityGroup.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// InputSecurityGroup Id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of inputs currently using this InputSecurityGroup.
	Inputs []*string `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Whitelist rules. See Whitelist Rules for more details.
	WhitelistRules []WhitelistRulesObservation `json:"whitelistRules,omitempty" tf:"whitelist_rules,omitempty"`
}

type InputSecurityGroupParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whitelist rules. See Whitelist Rules for more details.
	// +kubebuilder:validation:Optional
	WhitelistRules []WhitelistRulesParameters `json:"whitelistRules,omitempty" tf:"whitelist_rules,omitempty"`
}

type WhitelistRulesInitParameters struct {

	// The IPv4 CIDR that's whitelisted.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`
}

type WhitelistRulesObservation struct {

	// The IPv4 CIDR that's whitelisted.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`
}

type WhitelistRulesParameters struct {

	// The IPv4 CIDR that's whitelisted.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`
}

// InputSecurityGroupSpec defines the desired state of InputSecurityGroup
type InputSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InputSecurityGroupParameters `json:"forProvider"`
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
	InitProvider InputSecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// InputSecurityGroupStatus defines the observed state of InputSecurityGroup.
type InputSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InputSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InputSecurityGroup is the Schema for the InputSecurityGroups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InputSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.whitelistRules) || (has(self.initProvider) && has(self.initProvider.whitelistRules))",message="spec.forProvider.whitelistRules is a required parameter"
	Spec   InputSecurityGroupSpec   `json:"spec"`
	Status InputSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InputSecurityGroupList contains a list of InputSecurityGroups
type InputSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InputSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	InputSecurityGroup_Kind             = "InputSecurityGroup"
	InputSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InputSecurityGroup_Kind}.String()
	InputSecurityGroup_KindAPIVersion   = InputSecurityGroup_Kind + "." + CRDGroupVersion.String()
	InputSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(InputSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&InputSecurityGroup{}, &InputSecurityGroupList{})
}
