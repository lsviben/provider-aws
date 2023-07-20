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

type IntegrationInitParameters struct {

	// List of cache key parameters for the integration.
	CacheKeyParameters []*string `json:"cacheKeyParameters,omitempty" tf:"cache_key_parameters,omitempty"`

	// Integration's cache namespace.
	CacheNamespace *string `json:"cacheNamespace,omitempty" tf:"cache_namespace,omitempty"`

	// ID of the VpcLink used for the integration. Required if connection_type is VPC_LINK
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.VPCLink
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// Integration input's connectionType. Valid values are INTERNET (default for connections through the public routable internet), and VPC_LINK (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle request payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`

	// Credentials required for the integration. For AWS integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::\*:user/\*.
	Credentials *string `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// HTTP method (GET, POST, PUT, DELETE, HEAD, OPTION, ANY)
	// when calling the associated resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("http_method",false)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// Integration HTTP method
	// (GET, POST, PUT, DELETE, HEAD, OPTIONs, ANY, PATCH) specifying how API Gateway will interact with the back end.
	// Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// Not all methods are compatible with all AWS integrations.
	// e.g., Lambda function can only be invoked via POST.
	IntegrationHTTPMethod *string `json:"integrationHttpMethod,omitempty" tf:"integration_http_method,omitempty"`

	// Integration passthrough behavior (WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER).  Required if request_templates is used.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Map of request query string parameters and headers that should be passed to the backend responder.
	// For example: request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of the integration's request templates.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// API resource ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// ID of the associated REST API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// TLS configuration. See below.
	TLSConfig []TLSConfigInitParameters `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`

	// Integration input's type. Valid values are HTTP (for HTTP backends), MOCK (not calling any real backend), AWS (for AWS services), AWS_PROXY (for Lambda proxy integration) and HTTP_PROXY (for HTTP proxy integration). An HTTP or HTTP_PROXY integration with a connection_type of VPC_LINK is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Input's URI. Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}. region, subdomain and service are used to determine the right endpoint.
	// e.g., arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("invoke_arn",true)
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	URIRef *v1.Reference `json:"uriRef,omitempty" tf:"-"`

	URISelector *v1.Selector `json:"uriSelector,omitempty" tf:"-"`
}

type IntegrationObservation struct {

	// List of cache key parameters for the integration.
	CacheKeyParameters []*string `json:"cacheKeyParameters,omitempty" tf:"cache_key_parameters,omitempty"`

	// Integration's cache namespace.
	CacheNamespace *string `json:"cacheNamespace,omitempty" tf:"cache_namespace,omitempty"`

	// ID of the VpcLink used for the integration. Required if connection_type is VPC_LINK
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Integration input's connectionType. Valid values are INTERNET (default for connections through the public routable internet), and VPC_LINK (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle request payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`

	// Credentials required for the integration. For AWS integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::\*:user/\*.
	Credentials *string `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// HTTP method (GET, POST, PUT, DELETE, HEAD, OPTION, ANY)
	// when calling the associated resource.
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Integration HTTP method
	// (GET, POST, PUT, DELETE, HEAD, OPTIONs, ANY, PATCH) specifying how API Gateway will interact with the back end.
	// Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// Not all methods are compatible with all AWS integrations.
	// e.g., Lambda function can only be invoked via POST.
	IntegrationHTTPMethod *string `json:"integrationHttpMethod,omitempty" tf:"integration_http_method,omitempty"`

	// Integration passthrough behavior (WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER).  Required if request_templates is used.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// Map of request query string parameters and headers that should be passed to the backend responder.
	// For example: request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of the integration's request templates.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// API resource ID.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// ID of the associated REST API.
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// TLS configuration. See below.
	TLSConfig []TLSConfigObservation `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`

	// Integration input's type. Valid values are HTTP (for HTTP backends), MOCK (not calling any real backend), AWS (for AWS services), AWS_PROXY (for Lambda proxy integration) and HTTP_PROXY (for HTTP proxy integration). An HTTP or HTTP_PROXY integration with a connection_type of VPC_LINK is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Input's URI. Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}. region, subdomain and service are used to determine the right endpoint.
	// e.g., arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type IntegrationParameters struct {

	// List of cache key parameters for the integration.
	CacheKeyParameters []*string `json:"cacheKeyParameters,omitempty" tf:"cache_key_parameters,omitempty"`

	// Integration's cache namespace.
	CacheNamespace *string `json:"cacheNamespace,omitempty" tf:"cache_namespace,omitempty"`

	// ID of the VpcLink used for the integration. Required if connection_type is VPC_LINK
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.VPCLink
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// Reference to a VPCLink in apigateway to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// Selector for a VPCLink in apigateway to populate connectionId.
	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// Integration input's connectionType. Valid values are INTERNET (default for connections through the public routable internet), and VPC_LINK (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// How to handle request payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`

	// Credentials required for the integration. For AWS integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::\*:user/\*.
	Credentials *string `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// HTTP method (GET, POST, PUT, DELETE, HEAD, OPTION, ANY)
	// when calling the associated resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Method
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("http_method",false)
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// Reference to a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodRef *v1.Reference `json:"httpMethodRef,omitempty" tf:"-"`

	// Selector for a Method in apigateway to populate httpMethod.
	// +kubebuilder:validation:Optional
	HTTPMethodSelector *v1.Selector `json:"httpMethodSelector,omitempty" tf:"-"`

	// Integration HTTP method
	// (GET, POST, PUT, DELETE, HEAD, OPTIONs, ANY, PATCH) specifying how API Gateway will interact with the back end.
	// Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// Not all methods are compatible with all AWS integrations.
	// e.g., Lambda function can only be invoked via POST.
	IntegrationHTTPMethod *string `json:"integrationHttpMethod,omitempty" tf:"integration_http_method,omitempty"`

	// Integration passthrough behavior (WHEN_NO_MATCH, WHEN_NO_TEMPLATES, NEVER).  Required if request_templates is used.
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// Map of request query string parameters and headers that should be passed to the backend responder.
	// For example: request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// Map of the integration's request templates.
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// API resource ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.Resource
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Resource in apigateway to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// ID of the associated REST API.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/apigateway/v1beta1.RestAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	RestAPIID *string `json:"restApiId,omitempty" tf:"rest_api_id,omitempty"`

	// Reference to a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDRef *v1.Reference `json:"restApiIdRef,omitempty" tf:"-"`

	// Selector for a RestAPI in apigateway to populate restApiId.
	// +kubebuilder:validation:Optional
	RestAPIIDSelector *v1.Selector `json:"restApiIdSelector,omitempty" tf:"-"`

	// TLS configuration. See below.
	TLSConfig []TLSConfigParameters `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds.
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`

	// Integration input's type. Valid values are HTTP (for HTTP backends), MOCK (not calling any real backend), AWS (for AWS services), AWS_PROXY (for Lambda proxy integration) and HTTP_PROXY (for HTTP proxy integration). An HTTP or HTTP_PROXY integration with a connection_type of VPC_LINK is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Input's URI. Required if type is AWS, AWS_PROXY, HTTP or HTTP_PROXY.
	// For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}. region, subdomain and service are used to determine the right endpoint.
	// e.g., arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("invoke_arn",true)
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// Reference to a Function in lambda to populate uri.
	// +kubebuilder:validation:Optional
	URIRef *v1.Reference `json:"uriRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate uri.
	// +kubebuilder:validation:Optional
	URISelector *v1.Selector `json:"uriSelector,omitempty" tf:"-"`
}

type TLSConfigInitParameters struct {

	// Whether or not API Gateway skips verification that the certificate for an integration endpoint is issued by a supported certificate authority. This isn’t recommended, but it enables you to use certificates that are signed by private certificate authorities, or certificates that are self-signed. If enabled, API Gateway still performs basic certificate validation, which includes checking the certificate's expiration date, hostname, and presence of a root certificate authority. Supported only for HTTP and HTTP_PROXY integrations.
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty" tf:"insecure_skip_verification,omitempty"`
}

type TLSConfigObservation struct {

	// Whether or not API Gateway skips verification that the certificate for an integration endpoint is issued by a supported certificate authority. This isn’t recommended, but it enables you to use certificates that are signed by private certificate authorities, or certificates that are self-signed. If enabled, API Gateway still performs basic certificate validation, which includes checking the certificate's expiration date, hostname, and presence of a root certificate authority. Supported only for HTTP and HTTP_PROXY integrations.
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty" tf:"insecure_skip_verification,omitempty"`
}

type TLSConfigParameters struct {

	// Whether or not API Gateway skips verification that the certificate for an integration endpoint is issued by a supported certificate authority. This isn’t recommended, but it enables you to use certificates that are signed by private certificate authorities, or certificates that are self-signed. If enabled, API Gateway still performs basic certificate validation, which includes checking the certificate's expiration date, hostname, and presence of a root certificate authority. Supported only for HTTP and HTTP_PROXY integrations.
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty" tf:"insecure_skip_verification,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider IntegrationInitParameters `json:"initProvider,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Integration is the Schema for the Integrations API. Provides an HTTP Method Integration for an API Gateway Integration.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="%!s(MISSING) is a required parameter"
	Spec   IntegrationSpec   `json:"spec"`
	Status IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
