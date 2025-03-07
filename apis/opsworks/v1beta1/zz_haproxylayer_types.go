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

type HAProxyLayerCloudwatchConfigurationInitParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	LogStreams []HAProxyLayerCloudwatchConfigurationLogStreamsInitParameters `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type HAProxyLayerCloudwatchConfigurationLogStreamsInitParameters struct {
	BatchCount *float64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	BufferDuration *float64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	File *string `json:"file,omitempty" tf:"file,omitempty"`

	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// A human-readable name for the layer.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type HAProxyLayerCloudwatchConfigurationLogStreamsObservation struct {
	BatchCount *float64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	BufferDuration *float64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	File *string `json:"file,omitempty" tf:"file,omitempty"`

	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// A human-readable name for the layer.
	LogGroupName *string `json:"logGroupName,omitempty" tf:"log_group_name,omitempty"`

	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type HAProxyLayerCloudwatchConfigurationLogStreamsParameters struct {

	// +kubebuilder:validation:Optional
	BatchCount *float64 `json:"batchCount,omitempty" tf:"batch_count,omitempty"`

	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// +kubebuilder:validation:Optional
	BufferDuration *float64 `json:"bufferDuration,omitempty" tf:"buffer_duration,omitempty"`

	// +kubebuilder:validation:Optional
	DatetimeFormat *string `json:"datetimeFormat,omitempty" tf:"datetime_format,omitempty"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	File *string `json:"file" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	FileFingerprintLines *string `json:"fileFingerprintLines,omitempty" tf:"file_fingerprint_lines,omitempty"`

	// +kubebuilder:validation:Optional
	InitialPosition *string `json:"initialPosition,omitempty" tf:"initial_position,omitempty"`

	// A human-readable name for the layer.
	// +kubebuilder:validation:Optional
	LogGroupName *string `json:"logGroupName" tf:"log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	MultilineStartPattern *string `json:"multilineStartPattern,omitempty" tf:"multiline_start_pattern,omitempty"`

	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type HAProxyLayerCloudwatchConfigurationObservation struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	LogStreams []HAProxyLayerCloudwatchConfigurationLogStreamsObservation `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type HAProxyLayerCloudwatchConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogStreams []HAProxyLayerCloudwatchConfigurationLogStreamsParameters `json:"logStreams,omitempty" tf:"log_streams,omitempty"`
}

type HAProxyLayerEBSVolumeInitParameters struct {
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	MountPoint *string `json:"mountPoint,omitempty" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	NumberOfDisks *float64 `json:"numberOfDisks,omitempty" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HAProxyLayerEBSVolumeObservation struct {
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	MountPoint *string `json:"mountPoint,omitempty" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	NumberOfDisks *float64 `json:"numberOfDisks,omitempty" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HAProxyLayerEBSVolumeParameters struct {

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// For PIOPS volumes, the IOPS per disk.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The path to mount the EBS volume on the layer's instances.
	// +kubebuilder:validation:Optional
	MountPoint *string `json:"mountPoint" tf:"mount_point,omitempty"`

	// The number of disks to use for the EBS volume.
	// +kubebuilder:validation:Optional
	NumberOfDisks *float64 `json:"numberOfDisks" tf:"number_of_disks,omitempty"`

	// The RAID level to use for the volume.
	// +kubebuilder:validation:Optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// The size of the volume in gigabytes.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size" tf:"size,omitempty"`

	// The type of volume to create. This may be standard (the default), io1 or gp2.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HAProxyLayerInitParameters struct {

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	CloudwatchConfiguration []HAProxyLayerCloudwatchConfigurationInitParameters `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// ebs_volume blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EBSVolume []HAProxyLayerEBSVolumeInitParameters `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `json:"healthcheckMethod,omitempty" tf:"healthcheck_method,omitempty"`

	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckURL *string `json:"healthcheckUrl,omitempty" tf:"healthcheck_url,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	LoadBasedAutoScaling []HAProxyLayerLoadBasedAutoScalingInitParameters `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to enable HAProxy stats.
	StatsEnabled *bool `json:"statsEnabled,omitempty" tf:"stats_enabled,omitempty"`

	// The password to use for HAProxy stats.
	StatsPassword *string `json:"statsPassword,omitempty" tf:"stats_password,omitempty"`

	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsURL *string `json:"statsUrl,omitempty" tf:"stats_url,omitempty"`

	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `json:"statsUser,omitempty" tf:"stats_user,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingDownscalingInitParameters struct {
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingDownscalingObservation struct {
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingDownscalingParameters struct {

	// +kubebuilder:validation:Optional
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// +kubebuilder:validation:Optional
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingInitParameters struct {
	Downscaling []HAProxyLayerLoadBasedAutoScalingDownscalingInitParameters `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	Upscaling []HAProxyLayerLoadBasedAutoScalingUpscalingInitParameters `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingObservation struct {
	Downscaling []HAProxyLayerLoadBasedAutoScalingDownscalingObservation `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	Upscaling []HAProxyLayerLoadBasedAutoScalingUpscalingObservation `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingParameters struct {

	// +kubebuilder:validation:Optional
	Downscaling []HAProxyLayerLoadBasedAutoScalingDownscalingParameters `json:"downscaling,omitempty" tf:"downscaling,omitempty"`

	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// +kubebuilder:validation:Optional
	Upscaling []HAProxyLayerLoadBasedAutoScalingUpscalingParameters `json:"upscaling,omitempty" tf:"upscaling,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingUpscalingInitParameters struct {
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingUpscalingObservation struct {
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerLoadBasedAutoScalingUpscalingParameters struct {

	// +kubebuilder:validation:Optional
	Alarms []*string `json:"alarms,omitempty" tf:"alarms,omitempty"`

	// +kubebuilder:validation:Optional
	CPUThreshold *float64 `json:"cpuThreshold,omitempty" tf:"cpu_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreMetricsTime *float64 `json:"ignoreMetricsTime,omitempty" tf:"ignore_metrics_time,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	LoadThreshold *float64 `json:"loadThreshold,omitempty" tf:"load_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryThreshold *float64 `json:"memoryThreshold,omitempty" tf:"memory_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ThresholdsWaitTime *float64 `json:"thresholdsWaitTime,omitempty" tf:"thresholds_wait_time,omitempty"`
}

type HAProxyLayerObservation struct {

	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	CloudwatchConfiguration []HAProxyLayerCloudwatchConfigurationObservation `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []*string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids,omitempty"`

	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// ebs_volume blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EBSVolume []HAProxyLayerEBSVolumeObservation `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `json:"healthcheckMethod,omitempty" tf:"healthcheck_method,omitempty"`

	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckURL *string `json:"healthcheckUrl,omitempty" tf:"healthcheck_url,omitempty"`

	// The id of the layer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	LoadBasedAutoScaling []HAProxyLayerLoadBasedAutoScalingObservation `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the stack the layer will belong to.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Whether to enable HAProxy stats.
	StatsEnabled *bool `json:"statsEnabled,omitempty" tf:"stats_enabled,omitempty"`

	// The password to use for HAProxy stats.
	StatsPassword *string `json:"statsPassword,omitempty" tf:"stats_password,omitempty"`

	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsURL *string `json:"statsUrl,omitempty" tf:"stats_url,omitempty"`

	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `json:"statsUser,omitempty" tf:"stats_user,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Whether to use EBS-optimized instances.
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

type HAProxyLayerParameters struct {

	// Whether to automatically assign an elastic IP address to the layer's instances.
	// +kubebuilder:validation:Optional
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	// +kubebuilder:validation:Optional
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// Whether to enable auto-healing for the layer.
	// +kubebuilder:validation:Optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchConfiguration []HAProxyLayerCloudwatchConfigurationParameters `json:"cloudwatchConfiguration,omitempty" tf:"cloudwatch_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// The ARN of an IAM profile that will be used for the layer's instances.
	// +kubebuilder:validation:Optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// Custom JSON attributes to apply to the layer.
	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// References to SecurityGroup in ec2 to populate customSecurityGroupIds.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDRefs []v1.Reference `json:"customSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate customSecurityGroupIds.
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIDSelector *v1.Selector `json:"customSecurityGroupIdSelector,omitempty" tf:"-"`

	// Ids for a set of security groups to apply to the layer's instances.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=CustomSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=CustomSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIds []*string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// Whether to enable Elastic Load Balancing connection draining.
	// +kubebuilder:validation:Optional
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// ebs_volume blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	// +kubebuilder:validation:Optional
	EBSVolume []HAProxyLayerEBSVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// Name of an Elastic Load Balancer to attach to this layer
	// +kubebuilder:validation:Optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	// +kubebuilder:validation:Optional
	HealthcheckMethod *string `json:"healthcheckMethod,omitempty" tf:"healthcheck_method,omitempty"`

	// URL path to use for instance healthchecks. Defaults to "/".
	// +kubebuilder:validation:Optional
	HealthcheckURL *string `json:"healthcheckUrl,omitempty" tf:"healthcheck_url,omitempty"`

	// Whether to install OS and package updates on each instance when it boots.
	// +kubebuilder:validation:Optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	// +kubebuilder:validation:Optional
	InstanceShutdownTimeout *float64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBasedAutoScaling []HAProxyLayerLoadBasedAutoScalingParameters `json:"loadBasedAutoScaling,omitempty" tf:"load_based_auto_scaling,omitempty"`

	// A human-readable name for the layer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the stack the layer will belong to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`

	// Whether to enable HAProxy stats.
	// +kubebuilder:validation:Optional
	StatsEnabled *bool `json:"statsEnabled,omitempty" tf:"stats_enabled,omitempty"`

	// The password to use for HAProxy stats.
	// +kubebuilder:validation:Optional
	StatsPassword *string `json:"statsPassword,omitempty" tf:"stats_password,omitempty"`

	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	// +kubebuilder:validation:Optional
	StatsURL *string `json:"statsUrl,omitempty" tf:"stats_url,omitempty"`

	// The username for HAProxy stats. Defaults to "opsworks".
	// +kubebuilder:validation:Optional
	StatsUser *string `json:"statsUser,omitempty" tf:"stats_user,omitempty"`

	// Names of a set of system packages to install on the layer's instances.
	// +kubebuilder:validation:Optional
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether to use EBS-optimized instances.
	// +kubebuilder:validation:Optional
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

// HAProxyLayerSpec defines the desired state of HAProxyLayer
type HAProxyLayerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HAProxyLayerParameters `json:"forProvider"`
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
	InitProvider HAProxyLayerInitParameters `json:"initProvider,omitempty"`
}

// HAProxyLayerStatus defines the observed state of HAProxyLayer.
type HAProxyLayerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HAProxyLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HAProxyLayer is the Schema for the HAProxyLayers API. Provides an OpsWorks HAProxy layer resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HAProxyLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statsPassword) || (has(self.initProvider) && has(self.initProvider.statsPassword))",message="spec.forProvider.statsPassword is a required parameter"
	Spec   HAProxyLayerSpec   `json:"spec"`
	Status HAProxyLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HAProxyLayerList contains a list of HAProxyLayers
type HAProxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HAProxyLayer `json:"items"`
}

// Repository type metadata.
var (
	HAProxyLayer_Kind             = "HAProxyLayer"
	HAProxyLayer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HAProxyLayer_Kind}.String()
	HAProxyLayer_KindAPIVersion   = HAProxyLayer_Kind + "." + CRDGroupVersion.String()
	HAProxyLayer_GroupVersionKind = CRDGroupVersion.WithKind(HAProxyLayer_Kind)
)

func init() {
	SchemeBuilder.Register(&HAProxyLayer{}, &HAProxyLayerList{})
}
