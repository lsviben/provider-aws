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

type ConnectionInitParameters struct {

	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with provider_type
	HostArn *string `json:"hostArn,omitempty" tf:"host_arn,omitempty"`

	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing name will create a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the external provider where your third-party code repository is configured. Valid values are Bitbucket, GitHub or GitHubEnterpriseServer. Changing provider_type will create a new resource. Conflicts with host_arn
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConnectionObservation struct {

	// The codestar connection ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The codestar connection status. Possible values are PENDING, AVAILABLE and ERROR.
	ConnectionStatus *string `json:"connectionStatus,omitempty" tf:"connection_status,omitempty"`

	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with provider_type
	HostArn *string `json:"hostArn,omitempty" tf:"host_arn,omitempty"`

	// The codestar connection ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing name will create a new resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the external provider where your third-party code repository is configured. Valid values are Bitbucket, GitHub or GitHubEnterpriseServer. Changing provider_type will create a new resource. Conflicts with host_arn
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConnectionParameters struct {

	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with provider_type
	// +kubebuilder:validation:Optional
	HostArn *string `json:"hostArn,omitempty" tf:"host_arn,omitempty"`

	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing name will create a new resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the external provider where your third-party code repository is configured. Valid values are Bitbucket, GitHub or GitHubEnterpriseServer. Changing provider_type will create a new resource. Conflicts with host_arn
	// +kubebuilder:validation:Optional
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Provides a CodeStar Connection
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ConnectionSpec   `json:"spec"`
	Status ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}
