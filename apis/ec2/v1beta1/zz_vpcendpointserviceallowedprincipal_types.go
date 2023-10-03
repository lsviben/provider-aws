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

type VPCEndpointServiceAllowedPrincipalInitParameters struct {

	// The ARN of the principal to allow permissions.
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`
}

type VPCEndpointServiceAllowedPrincipalObservation struct {

	// The ID of the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the principal to allow permissions.
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// The ID of the VPC endpoint service to allow permission.
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`
}

type VPCEndpointServiceAllowedPrincipalParameters struct {

	// The ARN of the principal to allow permissions.
	// +kubebuilder:validation:Optional
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the VPC endpoint service to allow permission.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCEndpointService
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCEndpointServiceID *string `json:"vpcEndpointServiceId,omitempty" tf:"vpc_endpoint_service_id,omitempty"`

	// Reference to a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDRef *v1.Reference `json:"vpcEndpointServiceIdRef,omitempty" tf:"-"`

	// Selector for a VPCEndpointService in ec2 to populate vpcEndpointServiceId.
	// +kubebuilder:validation:Optional
	VPCEndpointServiceIDSelector *v1.Selector `json:"vpcEndpointServiceIdSelector,omitempty" tf:"-"`
}

// VPCEndpointServiceAllowedPrincipalSpec defines the desired state of VPCEndpointServiceAllowedPrincipal
type VPCEndpointServiceAllowedPrincipalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCEndpointServiceAllowedPrincipalParameters `json:"forProvider"`
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
	InitProvider VPCEndpointServiceAllowedPrincipalInitParameters `json:"initProvider,omitempty"`
}

// VPCEndpointServiceAllowedPrincipalStatus defines the observed state of VPCEndpointServiceAllowedPrincipal.
type VPCEndpointServiceAllowedPrincipalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCEndpointServiceAllowedPrincipalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceAllowedPrincipal is the Schema for the VPCEndpointServiceAllowedPrincipals API. Provides a resource to allow a principal to discover a VPC endpoint service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCEndpointServiceAllowedPrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalArn) || (has(self.initProvider) && has(self.initProvider.principalArn))",message="spec.forProvider.principalArn is a required parameter"
	Spec   VPCEndpointServiceAllowedPrincipalSpec   `json:"spec"`
	Status VPCEndpointServiceAllowedPrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCEndpointServiceAllowedPrincipalList contains a list of VPCEndpointServiceAllowedPrincipals
type VPCEndpointServiceAllowedPrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointServiceAllowedPrincipal `json:"items"`
}

// Repository type metadata.
var (
	VPCEndpointServiceAllowedPrincipal_Kind             = "VPCEndpointServiceAllowedPrincipal"
	VPCEndpointServiceAllowedPrincipal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCEndpointServiceAllowedPrincipal_Kind}.String()
	VPCEndpointServiceAllowedPrincipal_KindAPIVersion   = VPCEndpointServiceAllowedPrincipal_Kind + "." + CRDGroupVersion.String()
	VPCEndpointServiceAllowedPrincipal_GroupVersionKind = CRDGroupVersion.WithKind(VPCEndpointServiceAllowedPrincipal_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCEndpointServiceAllowedPrincipal{}, &VPCEndpointServiceAllowedPrincipalList{})
}
