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

type ProviderInitParameters struct {

	// Virtual node associated with a virtual service.
	VirtualNode []ProviderVirtualNodeInitParameters `json:"virtualNode,omitempty" tf:"virtual_node,omitempty"`

	// Virtual router associated with a virtual service.
	VirtualRouter []ProviderVirtualRouterInitParameters `json:"virtualRouter,omitempty" tf:"virtual_router,omitempty"`
}

type ProviderObservation struct {

	// Virtual node associated with a virtual service.
	VirtualNode []ProviderVirtualNodeObservation `json:"virtualNode,omitempty" tf:"virtual_node,omitempty"`

	// Virtual router associated with a virtual service.
	VirtualRouter []ProviderVirtualRouterObservation `json:"virtualRouter,omitempty" tf:"virtual_router,omitempty"`
}

type ProviderParameters struct {

	// Virtual node associated with a virtual service.
	// +kubebuilder:validation:Optional
	VirtualNode []ProviderVirtualNodeParameters `json:"virtualNode,omitempty" tf:"virtual_node,omitempty"`

	// Virtual router associated with a virtual service.
	// +kubebuilder:validation:Optional
	VirtualRouter []ProviderVirtualRouterParameters `json:"virtualRouter,omitempty" tf:"virtual_router,omitempty"`
}

type ProviderVirtualNodeInitParameters struct {
}

type ProviderVirtualNodeObservation struct {

	// Name of the virtual node that is acting as a service provider. Must be between 1 and 255 characters in length.
	VirtualNodeName *string `json:"virtualNodeName,omitempty" tf:"virtual_node_name,omitempty"`
}

type ProviderVirtualNodeParameters struct {

	// Name of the virtual node that is acting as a service provider. Must be between 1 and 255 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appmesh/v1beta1.VirtualNode
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	VirtualNodeName *string `json:"virtualNodeName,omitempty" tf:"virtual_node_name,omitempty"`

	// Reference to a VirtualNode in appmesh to populate virtualNodeName.
	// +kubebuilder:validation:Optional
	VirtualNodeNameRef *v1.Reference `json:"virtualNodeNameRef,omitempty" tf:"-"`

	// Selector for a VirtualNode in appmesh to populate virtualNodeName.
	// +kubebuilder:validation:Optional
	VirtualNodeNameSelector *v1.Selector `json:"virtualNodeNameSelector,omitempty" tf:"-"`
}

type ProviderVirtualRouterInitParameters struct {
}

type ProviderVirtualRouterObservation struct {

	// Name of the virtual router that is acting as a service provider. Must be between 1 and 255 characters in length.
	VirtualRouterName *string `json:"virtualRouterName,omitempty" tf:"virtual_router_name,omitempty"`
}

type ProviderVirtualRouterParameters struct {

	// Name of the virtual router that is acting as a service provider. Must be between 1 and 255 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appmesh/v1beta1.VirtualRouter
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	VirtualRouterName *string `json:"virtualRouterName,omitempty" tf:"virtual_router_name,omitempty"`

	// Reference to a VirtualRouter in appmesh to populate virtualRouterName.
	// +kubebuilder:validation:Optional
	VirtualRouterNameRef *v1.Reference `json:"virtualRouterNameRef,omitempty" tf:"-"`

	// Selector for a VirtualRouter in appmesh to populate virtualRouterName.
	// +kubebuilder:validation:Optional
	VirtualRouterNameSelector *v1.Selector `json:"virtualRouterNameSelector,omitempty" tf:"-"`
}

type VirtualServiceInitParameters_2 struct {

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Virtual service specification to apply.
	Spec []VirtualServiceSpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualServiceObservation_2 struct {

	// ARN of the virtual service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation date of the virtual service.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// ID of the virtual service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update date of the virtual service.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName *string `json:"meshName,omitempty" tf:"mesh_name,omitempty"`

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resource owner's AWS account ID.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// Virtual service specification to apply.
	Spec []VirtualServiceSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualServiceParameters_2 struct {

	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appmesh/v1beta1.Mesh
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MeshName *string `json:"meshName,omitempty" tf:"mesh_name,omitempty"`

	// Reference to a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameRef *v1.Reference `json:"meshNameRef,omitempty" tf:"-"`

	// Selector for a Mesh in appmesh to populate meshName.
	// +kubebuilder:validation:Optional
	MeshNameSelector *v1.Selector `json:"meshNameSelector,omitempty" tf:"-"`

	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Virtual service specification to apply.
	// +kubebuilder:validation:Optional
	Spec []VirtualServiceSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualServiceSpecInitParameters struct {

	// App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
	Provider []ProviderInitParameters `json:"provider,omitempty" tf:"provider,omitempty"`
}

type VirtualServiceSpecObservation struct {

	// App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
	Provider []ProviderObservation `json:"provider,omitempty" tf:"provider,omitempty"`
}

type VirtualServiceSpecParameters struct {

	// App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
	// +kubebuilder:validation:Optional
	Provider []ProviderParameters `json:"provider,omitempty" tf:"provider,omitempty"`
}

// VirtualServiceSpec defines the desired state of VirtualService
type VirtualServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualServiceParameters_2 `json:"forProvider"`
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
	InitProvider VirtualServiceInitParameters_2 `json:"initProvider,omitempty"`
}

// VirtualServiceStatus defines the observed state of VirtualService.
type VirtualServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualServiceObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualService is the Schema for the VirtualServices API. Provides an AWS App Mesh virtual service resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   VirtualServiceSpec   `json:"spec"`
	Status VirtualServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualServiceList contains a list of VirtualServices
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualService `json:"items"`
}

// Repository type metadata.
var (
	VirtualService_Kind             = "VirtualService"
	VirtualService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualService_Kind}.String()
	VirtualService_KindAPIVersion   = VirtualService_Kind + "." + CRDGroupVersion.String()
	VirtualService_GroupVersionKind = CRDGroupVersion.WithKind(VirtualService_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualService{}, &VirtualServiceList{})
}
