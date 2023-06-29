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

type HTTPNamespaceObservation struct {

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description that you specify for the namespace when you create it.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of an HTTP namespace.
	HTTPName *string `json:"httpName,omitempty" tf:"http_name,omitempty"`

	// The ID of a namespace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the http namespace.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HTTPNamespaceParameters struct {

	// The description that you specify for the namespace when you create it.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the http namespace.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HTTPNamespaceSpec defines the desired state of HTTPNamespace
type HTTPNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPNamespaceParameters `json:"forProvider"`
}

// HTTPNamespaceStatus defines the observed state of HTTPNamespace.
type HTTPNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPNamespace is the Schema for the HTTPNamespaces API. Provides a Service Discovery HTTP Namespace resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HTTPNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   HTTPNamespaceSpec   `json:"spec"`
	Status HTTPNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPNamespaceList contains a list of HTTPNamespaces
type HTTPNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPNamespace `json:"items"`
}

// Repository type metadata.
var (
	HTTPNamespace_Kind             = "HTTPNamespace"
	HTTPNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPNamespace_Kind}.String()
	HTTPNamespace_KindAPIVersion   = HTTPNamespace_Kind + "." + CRDGroupVersion.String()
	HTTPNamespace_GroupVersionKind = CRDGroupVersion.WithKind(HTTPNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPNamespace{}, &HTTPNamespaceList{})
}
