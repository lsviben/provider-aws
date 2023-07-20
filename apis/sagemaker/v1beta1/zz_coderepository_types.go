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

type CodeRepositoryInitParameters struct {

	// Specifies details about the repository. see Git Config details below.
	GitConfig []GitConfigInitParameters `json:"gitConfig,omitempty" tf:"git_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CodeRepositoryObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this Code Repository.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Specifies details about the repository. see Git Config details below.
	GitConfig []GitConfigObservation `json:"gitConfig,omitempty" tf:"git_config,omitempty"`

	// The name of the Code Repository.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CodeRepositoryParameters struct {

	// Specifies details about the repository. see Git Config details below.
	GitConfig []GitConfigParameters `json:"gitConfig,omitempty" tf:"git_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GitConfigInitParameters struct {

	// The default branch for the Git repository.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// The URL where the Git repository is located.
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository. The secret must have a staging label of AWSCURRENT and must be in the following format: {"username": UserName, "password": Password}
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	SecretArnRef *v1.Reference `json:"secretArnRef,omitempty" tf:"-"`

	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`
}

type GitConfigObservation struct {

	// The default branch for the Git repository.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// The URL where the Git repository is located.
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository. The secret must have a staging label of AWSCURRENT and must be in the following format: {"username": UserName, "password": Password}
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`
}

type GitConfigParameters struct {

	// The default branch for the Git repository.
	Branch *string `json:"branch,omitempty" tf:"branch,omitempty"`

	// The URL where the Git repository is located.
	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository. The secret must have a staging label of AWSCURRENT and must be in the following format: {"username": UserName, "password": Password}
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnRef *v1.Reference `json:"secretArnRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretArn.
	// +kubebuilder:validation:Optional
	SecretArnSelector *v1.Selector `json:"secretArnSelector,omitempty" tf:"-"`
}

// CodeRepositorySpec defines the desired state of CodeRepository
type CodeRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodeRepositoryParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CodeRepositoryInitParameters `json:"initProvider,omitempty"`
}

// CodeRepositoryStatus defines the observed state of CodeRepository.
type CodeRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodeRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodeRepository is the Schema for the CodeRepositorys API. Provides a SageMaker Code Repository resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CodeRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.gitConfig) || has(self.initProvider.gitConfig)",message="%!s(MISSING) is a required parameter"
	Spec   CodeRepositorySpec   `json:"spec"`
	Status CodeRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeRepositoryList contains a list of CodeRepositorys
type CodeRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeRepository `json:"items"`
}

// Repository type metadata.
var (
	CodeRepository_Kind             = "CodeRepository"
	CodeRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CodeRepository_Kind}.String()
	CodeRepository_KindAPIVersion   = CodeRepository_Kind + "." + CRDGroupVersion.String()
	CodeRepository_GroupVersionKind = CRDGroupVersion.WithKind(CodeRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&CodeRepository{}, &CodeRepositoryList{})
}
