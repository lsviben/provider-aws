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

type CertificateInitParameters struct {

	// Boolean flag to indicate if the certificate should be active
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The certificate signing request. Review
	// CreateCertificateFromCsr
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review CreateKeysAndCertificate
	// for more information on generating keys and a certificate.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`
}

type CertificateObservation struct {

	// Boolean flag to indicate if the certificate should be active
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The ARN of the created certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The certificate signing request. Review
	// CreateCertificateFromCsr
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review CreateKeysAndCertificate
	// for more information on generating keys and a certificate.
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// The internal ID assigned to this certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CertificateParameters struct {

	// Boolean flag to indicate if the certificate should be active
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// The CA certificate for the certificate to be registered. If this is set, the CA needs to be registered with AWS IoT beforehand.
	// +kubebuilder:validation:Optional
	CAPemSecretRef *v1.SecretKeySelector `json:"caPemSecretRef,omitempty" tf:"-"`

	// The certificate to be registered. If ca_pem is unspecified, review
	// RegisterCertificateWithoutCA.
	// If ca_pem is specified, review
	// RegisterCertificate
	// for more information on registering a certificate.
	// +kubebuilder:validation:Optional
	CertificatePemSecretRef *v1.SecretKeySelector `json:"certificatePemSecretRef,omitempty" tf:"-"`

	// The certificate signing request. Review
	// CreateCertificateFromCsr
	// for more information on generating a certificate from a certificate signing request (CSR).
	// If none is specified both the certificate and keys will be generated, review CreateKeysAndCertificate
	// for more information on generating keys and a certificate.
	// +kubebuilder:validation:Optional
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Creates and manages an AWS IoT certificate.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.active) || (has(self.initProvider) && has(self.initProvider.active))",message="spec.forProvider.active is a required parameter"
	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
