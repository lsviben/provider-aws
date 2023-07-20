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

type ActivityInitParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ActivityObservation struct {

	// The date the activity was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The Amazon Resource Name (ARN) that identifies the created activity.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ActivityParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ActivitySpec defines the desired state of Activity
type ActivitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActivityParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ActivityInitParameters `json:"initProvider,omitempty"`
}

// ActivityStatus defines the observed state of Activity.
type ActivityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActivityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Activity is the Schema for the Activitys API. Provides a Step Function Activity resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Activity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActivitySpec   `json:"spec"`
	Status            ActivityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActivityList contains a list of Activitys
type ActivityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Activity `json:"items"`
}

// Repository type metadata.
var (
	Activity_Kind             = "Activity"
	Activity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Activity_Kind}.String()
	Activity_KindAPIVersion   = Activity_Kind + "." + CRDGroupVersion.String()
	Activity_GroupVersionKind = CRDGroupVersion.WithKind(Activity_Kind)
)

func init() {
	SchemeBuilder.Register(&Activity{}, &ActivityList{})
}
