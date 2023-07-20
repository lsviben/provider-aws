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

type AppImageConfigInitParameters struct {

	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig []KernelGatewayImageConfigInitParameters `json:"kernelGatewayImageConfig,omitempty" tf:"kernel_gateway_image_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppImageConfigObservation struct {

	// The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the App Image Config.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig []KernelGatewayImageConfigObservation `json:"kernelGatewayImageConfig,omitempty" tf:"kernel_gateway_image_config,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AppImageConfigParameters struct {

	// The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
	KernelGatewayImageConfig []KernelGatewayImageConfigParameters `json:"kernelGatewayImageConfig,omitempty" tf:"kernel_gateway_image_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FileSystemConfigInitParameters struct {

	// The default POSIX group ID (GID). If not specified, defaults to 100. Valid values are 0 and 100.
	DefaultGID *float64 `json:"defaultGid,omitempty" tf:"default_gid,omitempty"`

	// The default POSIX user ID (UID). If not specified, defaults to 1000. Valid values are 0 and 1000.
	DefaultUID *float64 `json:"defaultUid,omitempty" tf:"default_uid,omitempty"`

	// The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`
}

type FileSystemConfigObservation struct {

	// The default POSIX group ID (GID). If not specified, defaults to 100. Valid values are 0 and 100.
	DefaultGID *float64 `json:"defaultGid,omitempty" tf:"default_gid,omitempty"`

	// The default POSIX user ID (UID). If not specified, defaults to 1000. Valid values are 0 and 1000.
	DefaultUID *float64 `json:"defaultUid,omitempty" tf:"default_uid,omitempty"`

	// The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`
}

type FileSystemConfigParameters struct {

	// The default POSIX group ID (GID). If not specified, defaults to 100. Valid values are 0 and 100.
	DefaultGID *float64 `json:"defaultGid,omitempty" tf:"default_gid,omitempty"`

	// The default POSIX user ID (UID). If not specified, defaults to 1000. Valid values are 0 and 1000.
	DefaultUID *float64 `json:"defaultUid,omitempty" tf:"default_uid,omitempty"`

	// The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`
}

type KernelGatewayImageConfigInitParameters struct {

	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig []FileSystemConfigInitParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// The default branch for the Git repository. See Kernel Spec details below.
	KernelSpec []KernelSpecInitParameters `json:"kernelSpec,omitempty" tf:"kernel_spec,omitempty"`
}

type KernelGatewayImageConfigObservation struct {

	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig []FileSystemConfigObservation `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// The default branch for the Git repository. See Kernel Spec details below.
	KernelSpec []KernelSpecObservation `json:"kernelSpec,omitempty" tf:"kernel_spec,omitempty"`
}

type KernelGatewayImageConfigParameters struct {

	// The URL where the Git repository is located. See File System Config details below.
	FileSystemConfig []FileSystemConfigParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config,omitempty"`

	// The default branch for the Git repository. See Kernel Spec details below.
	KernelSpec []KernelSpecParameters `json:"kernelSpec,omitempty" tf:"kernel_spec,omitempty"`
}

type KernelSpecInitParameters struct {

	// The display name of the kernel.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the kernel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type KernelSpecObservation struct {

	// The display name of the kernel.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the kernel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type KernelSpecParameters struct {

	// The display name of the kernel.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the kernel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// AppImageConfigSpec defines the desired state of AppImageConfig
type AppImageConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppImageConfigParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AppImageConfigInitParameters `json:"initProvider,omitempty"`
}

// AppImageConfigStatus defines the observed state of AppImageConfig.
type AppImageConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppImageConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppImageConfig is the Schema for the AppImageConfigs API. Provides a SageMaker App Image Config resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AppImageConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppImageConfigSpec   `json:"spec"`
	Status            AppImageConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppImageConfigList contains a list of AppImageConfigs
type AppImageConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppImageConfig `json:"items"`
}

// Repository type metadata.
var (
	AppImageConfig_Kind             = "AppImageConfig"
	AppImageConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppImageConfig_Kind}.String()
	AppImageConfig_KindAPIVersion   = AppImageConfig_Kind + "." + CRDGroupVersion.String()
	AppImageConfig_GroupVersionKind = CRDGroupVersion.WithKind(AppImageConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&AppImageConfig{}, &AppImageConfigList{})
}
