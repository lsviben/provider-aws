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

type ActionInitParameters struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// The region in which to run the action.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ActionObservation struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// The region in which to run the action.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ActionParameters struct {

	// A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are Approval, Build, Deploy, Invoke, Source and Test.
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// A map of the action declaration's configuration. Configurations options for action types and providers can be found in the Pipeline Structure Reference and Action Structure Reference documentation.
	Configuration map[string]*string `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of artifact names to be worked on.
	InputArtifacts []*string `json:"inputArtifacts,omitempty" tf:"input_artifacts,omitempty"`

	// The action declaration's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace all output variables will be accessed from.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// A list of artifact names to output. Output artifact names must be unique within a pipeline.
	OutputArtifacts []*string `json:"outputArtifacts,omitempty" tf:"output_artifacts,omitempty"`

	// The creator of the action being called. Possible values are AWS, Custom and ThirdParty.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The provider of the service being called by the action. Valid providers are determined by the action category. Provider names are listed in the Action Structure Reference documentation.
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// The region in which to run the action.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The order in which actions are run.
	RunOrder *float64 `json:"runOrder,omitempty" tf:"run_order,omitempty"`

	// A string that identifies the action type.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ArtifactStoreInitParameters struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	EncryptionKey []EncryptionKeyInitParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only S3 is supported.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ArtifactStoreObservation struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	EncryptionKey []EncryptionKeyObservation `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only S3 is supported.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ArtifactStoreParameters struct {

	// The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An encryption_key block is documented below.
	EncryptionKey []EncryptionKeyParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// The location where AWS CodePipeline stores artifacts for a pipeline; currently only S3 is supported.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Bucket in s3 to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of the artifact store, such as Amazon S3
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CodepipelineInitParameters struct {

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStore []ArtifactStoreInitParameters `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// The region in which to run the action.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	Stage []StageInitParameters `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CodepipelineObservation struct {

	// The codepipeline ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStore []ArtifactStoreObservation `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// The codepipeline ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	Stage []StageObservation `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CodepipelineParameters struct {

	// One or more artifact_store blocks. Artifact stores are documented below.
	ArtifactStore []ArtifactStoreParameters `json:"artifactStore,omitempty" tf:"artifact_store,omitempty"`

	// The region in which to run the action.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// (Minimum of at least two stage blocks is required) A stage block. Stages are documented below.
	Stage []StageParameters `json:"stage,omitempty" tf:"stage,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EncryptionKeyInitParameters struct {

	// The KMS key ARN or ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionKeyObservation struct {

	// The KMS key ARN or ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionKeyParameters struct {

	// The KMS key ARN or ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The type of key; currently only KMS is supported
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StageInitParameters struct {

	// The action(s) to include in the stage. Defined as an action block below
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the stage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StageObservation struct {

	// The action(s) to include in the stage. Defined as an action block below
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the stage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StageParameters struct {

	// The action(s) to include in the stage. Defined as an action block below
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The name of the stage.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// CodepipelineSpec defines the desired state of Codepipeline
type CodepipelineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CodepipelineParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CodepipelineInitParameters `json:"initProvider,omitempty"`
}

// CodepipelineStatus defines the observed state of Codepipeline.
type CodepipelineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CodepipelineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Codepipeline is the Schema for the Codepipelines API. Provides a CodePipeline
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Codepipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.artifactStore) || has(self.initProvider.artifactStore)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stage) || has(self.initProvider.stage)",message="%!s(MISSING) is a required parameter"
	Spec   CodepipelineSpec   `json:"spec"`
	Status CodepipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodepipelineList contains a list of Codepipelines
type CodepipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Codepipeline `json:"items"`
}

// Repository type metadata.
var (
	Codepipeline_Kind             = "Codepipeline"
	Codepipeline_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Codepipeline_Kind}.String()
	Codepipeline_KindAPIVersion   = Codepipeline_Kind + "." + CRDGroupVersion.String()
	Codepipeline_GroupVersionKind = CRDGroupVersion.WithKind(Codepipeline_Kind)
)

func init() {
	SchemeBuilder.Register(&Codepipeline{}, &CodepipelineList{})
}
