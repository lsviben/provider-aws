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

type QueuePolicyInitParameters struct {

	// The JSON policy for the SQS queue.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type QueuePolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The JSON policy for the SQS queue.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The URL of the SQS Queue to which to attach the policy
	QueueURL *string `json:"queueUrl,omitempty" tf:"queue_url,omitempty"`
}

type QueuePolicyParameters struct {

	// The JSON policy for the SQS queue.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The URL of the SQS Queue to which to attach the policy
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sqs/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.TerraformID()
	// +kubebuilder:validation:Optional
	QueueURL *string `json:"queueUrl,omitempty" tf:"queue_url,omitempty"`

	// Reference to a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLRef *v1.Reference `json:"queueUrlRef,omitempty" tf:"-"`

	// Selector for a Queue in sqs to populate queueUrl.
	// +kubebuilder:validation:Optional
	QueueURLSelector *v1.Selector `json:"queueUrlSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// QueuePolicySpec defines the desired state of QueuePolicy
type QueuePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueuePolicyParameters `json:"forProvider"`
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
	InitProvider QueuePolicyInitParameters `json:"initProvider,omitempty"`
}

// QueuePolicyStatus defines the observed state of QueuePolicy.
type QueuePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueuePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueuePolicy is the Schema for the QueuePolicys API. Provides a SQS Queue Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type QueuePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   QueuePolicySpec   `json:"spec"`
	Status QueuePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueuePolicyList contains a list of QueuePolicys
type QueuePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueuePolicy `json:"items"`
}

// Repository type metadata.
var (
	QueuePolicy_Kind             = "QueuePolicy"
	QueuePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueuePolicy_Kind}.String()
	QueuePolicy_KindAPIVersion   = QueuePolicy_Kind + "." + CRDGroupVersion.String()
	QueuePolicy_GroupVersionKind = CRDGroupVersion.WithKind(QueuePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&QueuePolicy{}, &QueuePolicyList{})
}
