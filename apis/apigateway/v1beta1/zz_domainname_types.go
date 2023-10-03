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

type DomainNameInitParameters struct {

	// Certificate issued for the domain name being registered, in PEM format. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Unique name to use when registering this certificate as an IAM server certificate. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name. Required if certificate_arn is not set.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Fully-qualified domain name to register.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Configuration block defining API endpoint information including type. See below.
	EndpointConfiguration []EndpointConfigurationInitParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Mutual TLS authentication configuration for the domain name. See below.
	MutualTLSAuthentication []MutualTLSAuthenticationInitParameters `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when certificate_arn is issued via an ACM Private CA or mutual_tls_authentication is configured with an ACM-imported certificate.)
	OwnershipVerificationCertificateArn *string `json:"ownershipVerificationCertificateArn,omitempty" tf:"ownership_verification_certificate_arn,omitempty"`

	// User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key.
	RegionalCertificateName *string `json:"regionalCertificateName,omitempty" tf:"regional_certificate_name,omitempty"`

	// Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are TLS_1_0 and TLS_1_2. Must be configured to perform drift detection.
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DomainNameObservation struct {

	// ARN of domain name.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with certificate_name, certificate_body, certificate_chain, certificate_private_key, regional_certificate_arn, and regional_certificate_name.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Certificate issued for the domain name being registered, in PEM format. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Unique name to use when registering this certificate as an IAM server certificate. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name. Required if certificate_arn is not set.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Upload date associated with the domain certificate.
	CertificateUploadDate *string `json:"certificateUploadDate,omitempty" tf:"certificate_upload_date,omitempty"`

	// Hostname created by Cloudfront to represent the distribution that implements this domain name mapping.
	CloudfrontDomainName *string `json:"cloudfrontDomainName,omitempty" tf:"cloudfront_domain_name,omitempty"`

	// For convenience, the hosted zone ID (Z2FDTNDATAQYW2) that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneID *string `json:"cloudfrontZoneId,omitempty" tf:"cloudfront_zone_id,omitempty"`

	// Fully-qualified domain name to register.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Configuration block defining API endpoint information including type. See below.
	EndpointConfiguration []EndpointConfigurationObservation `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Internal identifier assigned to this domain name by API Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Mutual TLS authentication configuration for the domain name. See below.
	MutualTLSAuthentication []MutualTLSAuthenticationObservation `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when certificate_arn is issued via an ACM Private CA or mutual_tls_authentication is configured with an ACM-imported certificate.)
	OwnershipVerificationCertificateArn *string `json:"ownershipVerificationCertificateArn,omitempty" tf:"ownership_verification_certificate_arn,omitempty"`

	// ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key.
	RegionalCertificateArn *string `json:"regionalCertificateArn,omitempty" tf:"regional_certificate_arn,omitempty"`

	// User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key.
	RegionalCertificateName *string `json:"regionalCertificateName,omitempty" tf:"regional_certificate_name,omitempty"`

	// Hostname for the custom domain's regional endpoint.
	RegionalDomainName *string `json:"regionalDomainName,omitempty" tf:"regional_domain_name,omitempty"`

	// Hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneID *string `json:"regionalZoneId,omitempty" tf:"regional_zone_id,omitempty"`

	// Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are TLS_1_0 and TLS_1_2. Must be configured to perform drift detection.
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DomainNameParameters struct {

	// ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with certificate_name, certificate_body, certificate_chain, certificate_private_key, regional_certificate_arn, and regional_certificate_name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acm/v1beta1.CertificateValidation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("certificate_arn",false)
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a CertificateValidation in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a CertificateValidation in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// Certificate issued for the domain name being registered, in PEM format. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Unique name to use when registering this certificate as an IAM server certificate. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name. Required if certificate_arn is not set.
	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// Private key associated with the domain certificate given in certificate_body. Only valid for EDGE endpoint configuration type. Conflicts with certificate_arn, regional_certificate_arn, and regional_certificate_name.
	// +kubebuilder:validation:Optional
	CertificatePrivateKeySecretRef *v1.SecretKeySelector `json:"certificatePrivateKeySecretRef,omitempty" tf:"-"`

	// Fully-qualified domain name to register.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Configuration block defining API endpoint information including type. See below.
	// +kubebuilder:validation:Optional
	EndpointConfiguration []EndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// Mutual TLS authentication configuration for the domain name. See below.
	// +kubebuilder:validation:Optional
	MutualTLSAuthentication []MutualTLSAuthenticationParameters `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when certificate_arn is issued via an ACM Private CA or mutual_tls_authentication is configured with an ACM-imported certificate.)
	// +kubebuilder:validation:Optional
	OwnershipVerificationCertificateArn *string `json:"ownershipVerificationCertificateArn,omitempty" tf:"ownership_verification_certificate_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acm/v1beta1.CertificateValidation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("certificate_arn",false)
	// +kubebuilder:validation:Optional
	RegionalCertificateArn *string `json:"regionalCertificateArn,omitempty" tf:"regional_certificate_arn,omitempty"`

	// Reference to a CertificateValidation in acm to populate regionalCertificateArn.
	// +kubebuilder:validation:Optional
	RegionalCertificateArnRef *v1.Reference `json:"regionalCertificateArnRef,omitempty" tf:"-"`

	// Selector for a CertificateValidation in acm to populate regionalCertificateArn.
	// +kubebuilder:validation:Optional
	RegionalCertificateArnSelector *v1.Selector `json:"regionalCertificateArnSelector,omitempty" tf:"-"`

	// User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with certificate_arn, certificate_name, certificate_body, certificate_chain, and certificate_private_key.
	// +kubebuilder:validation:Optional
	RegionalCertificateName *string `json:"regionalCertificateName,omitempty" tf:"regional_certificate_name,omitempty"`

	// Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are TLS_1_0 and TLS_1_2. Must be configured to perform drift detection.
	// +kubebuilder:validation:Optional
	SecurityPolicy *string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EndpointConfigurationInitParameters struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE or REGIONAL. If unspecified, defaults to EDGE. Must be declared as REGIONAL in non-Commercial partitions. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type EndpointConfigurationObservation struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE or REGIONAL. If unspecified, defaults to EDGE. Must be declared as REGIONAL in non-Commercial partitions. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type EndpointConfigurationParameters struct {

	// List of endpoint types. This resource currently only supports managing a single value. Valid values: EDGE or REGIONAL. If unspecified, defaults to EDGE. Must be declared as REGIONAL in non-Commercial partitions. Refer to the documentation for more information on the difference between edge-optimized and regional APIs.
	// +kubebuilder:validation:Optional
	Types []*string `json:"types" tf:"types,omitempty"`
}

type MutualTLSAuthenticationInitParameters struct {

	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	TruststoreURI *string `json:"truststoreUri,omitempty" tf:"truststore_uri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

type MutualTLSAuthenticationObservation struct {

	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	TruststoreURI *string `json:"truststoreUri,omitempty" tf:"truststore_uri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

type MutualTLSAuthenticationParameters struct {

	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	// +kubebuilder:validation:Optional
	TruststoreURI *string `json:"truststoreUri" tf:"truststore_uri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	// +kubebuilder:validation:Optional
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version,omitempty"`
}

// DomainNameSpec defines the desired state of DomainName
type DomainNameSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainNameParameters `json:"forProvider"`
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
	InitProvider DomainNameInitParameters `json:"initProvider,omitempty"`
}

// DomainNameStatus defines the observed state of DomainName.
type DomainNameStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainNameObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainName is the Schema for the DomainNames API. Registers a custom domain name for use with AWS API Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DomainName struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domainName) || (has(self.initProvider) && has(self.initProvider.domainName))",message="spec.forProvider.domainName is a required parameter"
	Spec   DomainNameSpec   `json:"spec"`
	Status DomainNameStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainNameList contains a list of DomainNames
type DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainName `json:"items"`
}

// Repository type metadata.
var (
	DomainName_Kind             = "DomainName"
	DomainName_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainName_Kind}.String()
	DomainName_KindAPIVersion   = DomainName_Kind + "." + CRDGroupVersion.String()
	DomainName_GroupVersionKind = CRDGroupVersion.WithKind(DomainName_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainName{}, &DomainNameList{})
}
