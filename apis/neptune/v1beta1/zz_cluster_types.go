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

type ClusterInitParameters struct {

	// Specifies whether upgrades between different major versions are allowed. You must set it to true when providing an engine_version parameter that uses a different major version than the DB cluster's current version. Default is false.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports audit.
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty" tf:"enable_cloudwatch_logs_exports,omitempty"`

	// The name of the database engine to be used for this Neptune cluster. Defaults to neptune.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_neptune_global_cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	NeptuneInstanceParameterGroupName *string `json:"neptuneInstanceParameterGroupName,omitempty" tf:"neptune_instance_parameter_group_name,omitempty"`

	// The port on which the Neptune accepts connections. Default is 8182.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// If set, create the Neptune cluster as a serverless one. See Serverless for example block attributes.
	ServerlessV2ScalingConfiguration []ServerlessV2ScalingConfigurationInitParameters `json:"serverlessV2ScalingConfiguration,omitempty" tf:"serverless_v2_scaling_configuration,omitempty"`

	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether the Neptune cluster is encrypted. The default is false if not specified.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterObservation struct {

	// Specifies whether upgrades between different major versions are allowed. You must set it to true when providing an engine_version parameter that uses a different major version than the DB cluster's current version. Default is false.
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// The Neptune Cluster Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// – List of Neptune Instances that are a part of this cluster
	ClusterMembers []*string `json:"clusterMembers,omitempty" tf:"cluster_members,omitempty"`

	// The Neptune Cluster Resource ID
	ClusterResourceID *string `json:"clusterResourceId,omitempty" tf:"cluster_resource_id,omitempty"`

	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports audit.
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty" tf:"enable_cloudwatch_logs_exports,omitempty"`

	// The DNS address of the Neptune instance
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The name of the database engine to be used for this Neptune cluster. Defaults to neptune.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_neptune_global_cluster.
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// The Route53 Hosted Zone ID of the endpoint
	HostedZoneID *string `json:"hostedZoneId,omitempty" tf:"hosted_zone_id,omitempty"`

	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The Neptune Cluster Identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_arn, storage_encrypted needs to be set to true.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// A cluster parameter group to associate with the cluster.
	NeptuneClusterParameterGroupName *string `json:"neptuneClusterParameterGroupName,omitempty" tf:"neptune_cluster_parameter_group_name,omitempty"`

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	NeptuneInstanceParameterGroupName *string `json:"neptuneInstanceParameterGroupName,omitempty" tf:"neptune_instance_parameter_group_name,omitempty"`

	// A Neptune subnet group to associate with this Neptune instance.
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// The port on which the Neptune accepts connections. Default is 8182.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
	ReaderEndpoint *string `json:"readerEndpoint,omitempty" tf:"reader_endpoint,omitempty"`

	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`

	// If set, create the Neptune cluster as a serverless one. See Serverless for example block attributes.
	ServerlessV2ScalingConfiguration []ServerlessV2ScalingConfigurationObservation `json:"serverlessV2ScalingConfiguration,omitempty" tf:"serverless_v2_scaling_configuration,omitempty"`

	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Specifies whether the Neptune cluster is encrypted. The default is false if not specified.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// List of VPC security groups to associate with the Cluster
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ClusterParameters struct {

	// Specifies whether upgrades between different major versions are allowed. You must set it to true when providing an engine_version parameter that uses a different major version than the DB cluster's current version. Default is false.
	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is false.
	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// The days to retain backups for. Default 1
	// +kubebuilder:validation:Optional
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period,omitempty"`

	// If set to true, tags are copied to any snapshot of the DB cluster that is created.
	// +kubebuilder:validation:Optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot,omitempty"`

	// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports audit.
	// +kubebuilder:validation:Optional
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty" tf:"enable_cloudwatch_logs_exports,omitempty"`

	// The name of the database engine to be used for this Neptune cluster. Defaults to neptune.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The database engine version.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// The global cluster identifier specified on aws_neptune_global_cluster.
	// +kubebuilder:validation:Optional
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty" tf:"global_cluster_identifier,omitempty"`

	// Specifies whether or not mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
	// +kubebuilder:validation:Optional
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty" tf:"iam_database_authentication_enabled,omitempty"`

	// References to Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleRefs []v1.Reference `json:"iamRoleRefs,omitempty" tf:"-"`

	// Selector for a list of Role in iam to populate iamRoles.
	// +kubebuilder:validation:Optional
	IAMRoleSelector *v1.Selector `json:"iamRoleSelector,omitempty" tf:"-"`

	// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:refFieldName=IAMRoleRefs
	// +crossplane:generate:reference:selectorFieldName=IAMRoleSelector
	// +kubebuilder:validation:Optional
	IAMRoles []*string `json:"iamRoles,omitempty" tf:"iam_roles,omitempty"`

	// The ARN for the KMS encryption key. When specifying kms_key_arn, storage_encrypted needs to be set to true.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// A cluster parameter group to associate with the cluster.
	// +crossplane:generate:reference:type=ClusterParameterGroup
	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupName *string `json:"neptuneClusterParameterGroupName,omitempty" tf:"neptune_cluster_parameter_group_name,omitempty"`

	// Reference to a ClusterParameterGroup to populate neptuneClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupNameRef *v1.Reference `json:"neptuneClusterParameterGroupNameRef,omitempty" tf:"-"`

	// Selector for a ClusterParameterGroup to populate neptuneClusterParameterGroupName.
	// +kubebuilder:validation:Optional
	NeptuneClusterParameterGroupNameSelector *v1.Selector `json:"neptuneClusterParameterGroupNameSelector,omitempty" tf:"-"`

	// The name of the DB parameter group to apply to all instances of the DB cluster.
	// +kubebuilder:validation:Optional
	NeptuneInstanceParameterGroupName *string `json:"neptuneInstanceParameterGroupName,omitempty" tf:"neptune_instance_parameter_group_name,omitempty"`

	// A Neptune subnet group to associate with this Neptune instance.
	// +crossplane:generate:reference:type=SubnetGroup
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupName *string `json:"neptuneSubnetGroupName,omitempty" tf:"neptune_subnet_group_name,omitempty"`

	// Reference to a SubnetGroup to populate neptuneSubnetGroupName.
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameRef *v1.Reference `json:"neptuneSubnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup to populate neptuneSubnetGroupName.
	// +kubebuilder:validation:Optional
	NeptuneSubnetGroupNameSelector *v1.Selector `json:"neptuneSubnetGroupNameSelector,omitempty" tf:"-"`

	// The port on which the Neptune accepts connections. Default is 8182.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per regionE.g., 04:00-09:00
	// +kubebuilder:validation:Optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window,omitempty"`

	// The weekly time range during which system maintenance can occur, in (UTC) e.g., wed:04:00-wed:04:30
	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifier *string `json:"replicationSourceIdentifier,omitempty" tf:"replication_source_identifier,omitempty"`

	// Reference to a Cluster to populate replicationSourceIdentifier.
	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifierRef *v1.Reference `json:"replicationSourceIdentifierRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate replicationSourceIdentifier.
	// +kubebuilder:validation:Optional
	ReplicationSourceIdentifierSelector *v1.Selector `json:"replicationSourceIdentifierSelector,omitempty" tf:"-"`

	// If set, create the Neptune cluster as a serverless one. See Serverless for example block attributes.
	// +kubebuilder:validation:Optional
	ServerlessV2ScalingConfiguration []ServerlessV2ScalingConfigurationParameters `json:"serverlessV2ScalingConfiguration,omitempty" tf:"serverless_v2_scaling_configuration,omitempty"`

	// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from final_snapshot_identifier. Default is false.
	// +kubebuilder:validation:Optional
	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot,omitempty"`

	// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
	// +crossplane:generate:reference:type=ClusterSnapshot
	// +kubebuilder:validation:Optional
	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier,omitempty"`

	// Reference to a ClusterSnapshot to populate snapshotIdentifier.
	// +kubebuilder:validation:Optional
	SnapshotIdentifierRef *v1.Reference `json:"snapshotIdentifierRef,omitempty" tf:"-"`

	// Selector for a ClusterSnapshot to populate snapshotIdentifier.
	// +kubebuilder:validation:Optional
	SnapshotIdentifierSelector *v1.Selector `json:"snapshotIdentifierSelector,omitempty" tf:"-"`

	// Specifies whether the Neptune cluster is encrypted. The default is false if not specified.
	// +kubebuilder:validation:Optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate vpcSecurityGroupIds.
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIDSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// List of VPC security groups to associate with the Cluster
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

type ServerlessV2ScalingConfigurationInitParameters struct {

	// : (default: 128) The maximum Neptune Capacity Units (NCUs) for this cluster. Must be lower or equal than 128. See AWS Documentation for more details.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// : (default: 2.5) The minimum Neptune Capacity Units (NCUs) for this cluster. Must be greater or equal than 1. See AWS Documentation for more details.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type ServerlessV2ScalingConfigurationObservation struct {

	// : (default: 128) The maximum Neptune Capacity Units (NCUs) for this cluster. Must be lower or equal than 128. See AWS Documentation for more details.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// : (default: 2.5) The minimum Neptune Capacity Units (NCUs) for this cluster. Must be greater or equal than 1. See AWS Documentation for more details.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type ServerlessV2ScalingConfigurationParameters struct {

	// : (default: 128) The maximum Neptune Capacity Units (NCUs) for this cluster. Must be lower or equal than 128. See AWS Documentation for more details.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// : (default: 2.5) The minimum Neptune Capacity Units (NCUs) for this cluster. Must be greater or equal than 1. See AWS Documentation for more details.
	// +kubebuilder:validation:Optional
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Provides an Neptune Cluster Resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
