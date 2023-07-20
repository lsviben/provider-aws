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

type AttachmentInitParameters struct {
}

type AttachmentObservation struct {

	// The identifier of the firewall endpoint that AWS Network Firewall has instantiated in the subnet. You use this to identify the firewall endpoint in the VPC route tables, when you redirect the VPC traffic through the endpoint.
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// The unique identifier for the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type AttachmentParameters struct {
}

type EncryptionConfigurationInitParameters struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionConfigurationObservation struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EncryptionConfigurationParameters struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FirewallInitParameters struct {

	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to false.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// A friendly description of the firewall.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []EncryptionConfigurationInitParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkfirewall/v1beta1.FirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	FirewallPolicyArn *string `json:"firewallPolicyArn,omitempty" tf:"firewall_policy_arn,omitempty"`

	FirewallPolicyArnRef *v1.Reference `json:"firewallPolicyArnRef,omitempty" tf:"-"`

	FirewallPolicyArnSelector *v1.Selector `json:"firewallPolicyArnSelector,omitempty" tf:"-"`

	// (Option) A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to false.
	FirewallPolicyChangeProtection *bool `json:"firewallPolicyChangeProtection,omitempty" tf:"firewall_policy_change_protection,omitempty"`

	// A friendly name of the firewall.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to false.
	SubnetChangeProtection *bool `json:"subnetChangeProtection,omitempty" tf:"subnet_change_protection,omitempty"`

	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMapping []SubnetMappingInitParameters `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type FirewallObservation struct {

	// The Amazon Resource Name (ARN) that identifies the firewall.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to false.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// A friendly description of the firewall.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []EncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn *string `json:"firewallPolicyArn,omitempty" tf:"firewall_policy_arn,omitempty"`

	// (Option) A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to false.
	FirewallPolicyChangeProtection *bool `json:"firewallPolicyChangeProtection,omitempty" tf:"firewall_policy_change_protection,omitempty"`

	// Nested list of information about the current status of the firewall.
	FirewallStatus []FirewallStatusObservation `json:"firewallStatus,omitempty" tf:"firewall_status,omitempty"`

	// The Amazon Resource Name (ARN) that identifies the firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A friendly name of the firewall.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to false.
	SubnetChangeProtection *bool `json:"subnetChangeProtection,omitempty" tf:"subnet_change_protection,omitempty"`

	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMapping []SubnetMappingObservation `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A string token used when updating a firewall.
	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`

	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type FirewallParameters struct {

	// A boolean flag indicating whether it is possible to delete the firewall. Defaults to false.
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection,omitempty"`

	// A friendly description of the firewall.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkfirewall/v1beta1.FirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	FirewallPolicyArn *string `json:"firewallPolicyArn,omitempty" tf:"firewall_policy_arn,omitempty"`

	// Reference to a FirewallPolicy in networkfirewall to populate firewallPolicyArn.
	// +kubebuilder:validation:Optional
	FirewallPolicyArnRef *v1.Reference `json:"firewallPolicyArnRef,omitempty" tf:"-"`

	// Selector for a FirewallPolicy in networkfirewall to populate firewallPolicyArn.
	// +kubebuilder:validation:Optional
	FirewallPolicyArnSelector *v1.Selector `json:"firewallPolicyArnSelector,omitempty" tf:"-"`

	// (Option) A boolean flag indicating whether it is possible to change the associated firewall policy. Defaults to false.
	FirewallPolicyChangeProtection *bool `json:"firewallPolicyChangeProtection,omitempty" tf:"firewall_policy_change_protection,omitempty"`

	// A friendly name of the firewall.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// A boolean flag indicating whether it is possible to change the associated subnet(s). Defaults to false.
	SubnetChangeProtection *bool `json:"subnetChangeProtection,omitempty" tf:"subnet_change_protection,omitempty"`

	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMapping []SubnetMappingParameters `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type FirewallStatusInitParameters struct {
}

type FirewallStatusObservation struct {

	// Set of subnets configured for use by the firewall.
	SyncStates []SyncStatesObservation `json:"syncStates,omitempty" tf:"sync_states,omitempty"`
}

type FirewallStatusParameters struct {
}

type SubnetMappingInitParameters struct {

	// The subnet's IP address type. Valida values: "DUALSTACK", "IPV4".
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The unique identifier for the subnet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type SubnetMappingObservation struct {

	// The subnet's IP address type. Valida values: "DUALSTACK", "IPV4".
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The unique identifier for the subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetMappingParameters struct {

	// The subnet's IP address type. Valida values: "DUALSTACK", "IPV4".
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The unique identifier for the subnet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in ec2 to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type SyncStatesInitParameters struct {
}

type SyncStatesObservation struct {

	// Nested list describing the attachment status of the firewall's association with a single VPC subnet.
	Attachment []AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	// The Availability Zone where the subnet is configured.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
}

type SyncStatesParameters struct {
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API. Provides an AWS Network Firewall Firewall resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetMapping) || has(self.initProvider.subnetMapping)",message="%!s(MISSING) is a required parameter"
	Spec   FirewallSpec   `json:"spec"`
	Status FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
