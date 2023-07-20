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

type DNSTargetResourceInitParameters struct {

	// DNS Name that acts as the ingress point to a portion of application.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Hosted Zone ARN that contains the DNS record with the provided name of target resource.
	HostedZoneArn *string `json:"hostedZoneArn,omitempty" tf:"hosted_zone_arn,omitempty"`

	// Route53 record set id to uniquely identify a record given a domain_name and a record_type.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`

	// Type of DNS Record of target resource.
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// Target resource the R53 record specified with the above params points to.
	TargetResource []TargetResourceInitParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`
}

type DNSTargetResourceObservation struct {

	// DNS Name that acts as the ingress point to a portion of application.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Hosted Zone ARN that contains the DNS record with the provided name of target resource.
	HostedZoneArn *string `json:"hostedZoneArn,omitempty" tf:"hosted_zone_arn,omitempty"`

	// Route53 record set id to uniquely identify a record given a domain_name and a record_type.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`

	// Type of DNS Record of target resource.
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// Target resource the R53 record specified with the above params points to.
	TargetResource []TargetResourceObservation `json:"targetResource,omitempty" tf:"target_resource,omitempty"`
}

type DNSTargetResourceParameters struct {

	// DNS Name that acts as the ingress point to a portion of application.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Hosted Zone ARN that contains the DNS record with the provided name of target resource.
	HostedZoneArn *string `json:"hostedZoneArn,omitempty" tf:"hosted_zone_arn,omitempty"`

	// Route53 record set id to uniquely identify a record given a domain_name and a record_type.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`

	// Type of DNS Record of target resource.
	RecordType *string `json:"recordType,omitempty" tf:"record_type,omitempty"`

	// Target resource the R53 record specified with the above params points to.
	TargetResource []TargetResourceParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`
}

type NlbResourceInitParameters struct {

	// NLB resource ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type NlbResourceObservation struct {

	// NLB resource ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type NlbResourceParameters struct {

	// NLB resource ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type R53ResourceInitParameters struct {

	// Domain name that is targeted.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource record set ID that is targeted.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`
}

type R53ResourceObservation struct {

	// Domain name that is targeted.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource record set ID that is targeted.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`
}

type R53ResourceParameters struct {

	// Domain name that is targeted.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Resource record set ID that is targeted.
	RecordSetID *string `json:"recordSetId,omitempty" tf:"record_set_id,omitempty"`
}

type ResourceSetInitParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Type of the resources in the resource set.
	ResourceSetType *string `json:"resourceSetType,omitempty" tf:"resource_set_type,omitempty"`

	// List of resources to add to this resource set. See below.
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourceSetObservation struct {

	// ARN of the resource set
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the resources in the resource set.
	ResourceSetType *string `json:"resourceSetType,omitempty" tf:"resource_set_type,omitempty"`

	// List of resources to add to this resource set. See below.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResourceSetParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Type of the resources in the resource set.
	ResourceSetType *string `json:"resourceSetType,omitempty" tf:"resource_set_type,omitempty"`

	// List of resources to add to this resource set. See below.
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ResourcesInitParameters struct {

	// Component for DNS/Routing Control Readiness Checks.
	DNSTargetResource []DNSTargetResourceInitParameters `json:"dnsTargetResource,omitempty" tf:"dns_target_resource,omitempty"`

	// Recovery group ARN or cell ARN that contains this resource set.
	ReadinessScopes []*string `json:"readinessScopes,omitempty" tf:"readiness_scopes,omitempty"`

	// ARN of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatch/v1beta1.MetricAlarm
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type ResourcesObservation struct {

	// Unique identified for DNS Target Resources, use for readiness checks.
	ComponentID *string `json:"componentId,omitempty" tf:"component_id,omitempty"`

	// Component for DNS/Routing Control Readiness Checks.
	DNSTargetResource []DNSTargetResourceObservation `json:"dnsTargetResource,omitempty" tf:"dns_target_resource,omitempty"`

	// Recovery group ARN or cell ARN that contains this resource set.
	ReadinessScopes []*string `json:"readinessScopes,omitempty" tf:"readiness_scopes,omitempty"`

	// ARN of the resource.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type ResourcesParameters struct {

	// Component for DNS/Routing Control Readiness Checks.
	DNSTargetResource []DNSTargetResourceParameters `json:"dnsTargetResource,omitempty" tf:"dns_target_resource,omitempty"`

	// Recovery group ARN or cell ARN that contains this resource set.
	ReadinessScopes []*string `json:"readinessScopes,omitempty" tf:"readiness_scopes,omitempty"`

	// ARN of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatch/v1beta1.MetricAlarm
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a MetricAlarm in cloudwatch to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a MetricAlarm in cloudwatch to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type TargetResourceInitParameters struct {

	// NLB resource a DNS Target Resource points to. Required if r53_resource is not set.
	NlbResource []NlbResourceInitParameters `json:"nlbResource,omitempty" tf:"nlb_resource,omitempty"`

	// Route53 resource a DNS Target Resource record points to.
	R53Resource []R53ResourceInitParameters `json:"r53Resource,omitempty" tf:"r53_resource,omitempty"`
}

type TargetResourceObservation struct {

	// NLB resource a DNS Target Resource points to. Required if r53_resource is not set.
	NlbResource []NlbResourceObservation `json:"nlbResource,omitempty" tf:"nlb_resource,omitempty"`

	// Route53 resource a DNS Target Resource record points to.
	R53Resource []R53ResourceObservation `json:"r53Resource,omitempty" tf:"r53_resource,omitempty"`
}

type TargetResourceParameters struct {

	// NLB resource a DNS Target Resource points to. Required if r53_resource is not set.
	NlbResource []NlbResourceParameters `json:"nlbResource,omitempty" tf:"nlb_resource,omitempty"`

	// Route53 resource a DNS Target Resource record points to.
	R53Resource []R53ResourceParameters `json:"r53Resource,omitempty" tf:"r53_resource,omitempty"`
}

// ResourceSetSpec defines the desired state of ResourceSet
type ResourceSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceSetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ResourceSetInitParameters `json:"initProvider,omitempty"`
}

// ResourceSetStatus defines the observed state of ResourceSet.
type ResourceSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSet is the Schema for the ResourceSets API. Provides an AWS Route 53 Recovery Readiness Resource Set
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceSetType) || has(self.initProvider.resourceSetType)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resources) || has(self.initProvider.resources)",message="%!s(MISSING) is a required parameter"
	Spec   ResourceSetSpec   `json:"spec"`
	Status ResourceSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceSetList contains a list of ResourceSets
type ResourceSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceSet `json:"items"`
}

// Repository type metadata.
var (
	ResourceSet_Kind             = "ResourceSet"
	ResourceSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceSet_Kind}.String()
	ResourceSet_KindAPIVersion   = ResourceSet_Kind + "." + CRDGroupVersion.String()
	ResourceSet_GroupVersionKind = CRDGroupVersion.WithKind(ResourceSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceSet{}, &ResourceSetList{})
}
