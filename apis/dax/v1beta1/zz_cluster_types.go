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

type ClusterObservation struct {

	// The ARN of the DAX cluster
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The DNS name of the DAX cluster without the port appended
	ClusterAddress *string `json:"clusterAddress,omitempty" tf:"cluster_address,omitempty"`

	// The configuration endpoint for this DAX cluster,
	// consisting of a DNS name and a port number
	ConfigurationEndpoint *string `json:"configurationEndpoint,omitempty" tf:"configuration_endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of node objects including id, address, port and
	// availability_zone. Referenceable e.g., as
	// ${aws_dax_cluster.test.nodes.0.address}
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// The port used by the configuration endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// List of Availability Zones in which the
	// nodes will be created
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// –  The type of encryption the
	// cluster's endpoint should support. Valid values are: NONE and TLS.
	// Default value is NONE.
	// +kubebuilder:validation:Optional
	ClusterEndpointEncryptionType *string `json:"clusterEndpointEncryptionType,omitempty" tf:"cluster_endpoint_encryption_type,omitempty"`

	// –  Description for the cluster
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A valid Amazon Resource Name (ARN) that identifies
	// an IAM role. At runtime, DAX will assume this role and use the role's
	// permissions to access DynamoDB on your behalf
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	IAMRoleArn *string `json:"iamRoleArn,omitempty" tf:"iam_role_arn,omitempty"`

	// Reference to a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnRef *v1.Reference `json:"iamRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate iamRoleArn.
	// +kubebuilder:validation:Optional
	IAMRoleArnSelector *v1.Selector `json:"iamRoleArnSelector,omitempty" tf:"-"`

	// ddd:hh24:mi
	// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
	// sun:05:00-sun:09:00
	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// –  The compute and memory capacity of the nodes. See
	// Nodes for supported node types
	// +kubebuilder:validation:Required
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// east-1:012345678999:my_sns_topic
	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// –  Name of the parameter group to associate
	// with this DAX cluster
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// node cluster, without any read
	// replicas
	// +kubebuilder:validation:Required
	ReplicationFactor *float64 `json:"replicationFactor" tf:"replication_factor,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// –  One or more VPC security groups associated
	// with the cluster
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Encrypt at rest options
	// +kubebuilder:validation:Optional
	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// –  Name of the subnet group to be used for the
	// cluster
	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NodesObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The port used by the configuration endpoint
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type NodesParameters struct {
}

type ServerSideEncryptionObservation struct {
}

type ServerSideEncryptionParameters struct {

	// Whether to enable encryption at rest. Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Provides an DAX Cluster resource.
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
