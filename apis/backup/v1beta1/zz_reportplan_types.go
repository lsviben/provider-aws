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

type ReportDeliveryChannelInitParameters struct {

	// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// The unique name of the S3 bucket that receives your reports.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type ReportDeliveryChannelObservation struct {

	// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// The unique name of the S3 bucket that receives your reports.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type ReportDeliveryChannelParameters struct {

	// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
	Formats []*string `json:"formats,omitempty" tf:"formats,omitempty"`

	// The unique name of the S3 bucket that receives your reports.
	S3BucketName *string `json:"s3BucketName,omitempty" tf:"s3_bucket_name,omitempty"`

	// The prefix for where Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`
}

type ReportPlanInitParameters struct {

	// The description of the report plan with a maximum of 1,024 characters
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// An object that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. Detailed below.
	ReportDeliveryChannel []ReportDeliveryChannelInitParameters `json:"reportDeliveryChannel,omitempty" tf:"report_delivery_channel,omitempty"`

	// An object that identifies the report template for the report. Reports are built using a report template. Detailed below.
	ReportSetting []ReportSettingInitParameters `json:"reportSetting,omitempty" tf:"report_setting,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReportPlanObservation struct {

	// The ARN of the backup report plan.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that a report plan is created, in Unix format and Coordinated Universal Time (UTC).
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The deployment status of a report plan. The statuses are: CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED.
	DeploymentStatus *string `json:"deploymentStatus,omitempty" tf:"deployment_status,omitempty"`

	// The description of the report plan with a maximum of 1,024 characters
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the backup report plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// An object that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. Detailed below.
	ReportDeliveryChannel []ReportDeliveryChannelObservation `json:"reportDeliveryChannel,omitempty" tf:"report_delivery_channel,omitempty"`

	// An object that identifies the report template for the report. Reports are built using a report template. Detailed below.
	ReportSetting []ReportSettingObservation `json:"reportSetting,omitempty" tf:"report_setting,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReportPlanParameters struct {

	// The description of the report plan with a maximum of 1,024 characters
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	Region *string `json:"region,omitempty" tf:"-"`

	// An object that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. Detailed below.
	ReportDeliveryChannel []ReportDeliveryChannelParameters `json:"reportDeliveryChannel,omitempty" tf:"report_delivery_channel,omitempty"`

	// An object that identifies the report template for the report. Reports are built using a report template. Detailed below.
	ReportSetting []ReportSettingParameters `json:"reportSetting,omitempty" tf:"report_setting,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReportSettingInitParameters struct {

	// Specifies the Amazon Resource Names (ARNs) of the frameworks a report covers.
	FrameworkArns []*string `json:"frameworkArns,omitempty" tf:"framework_arns,omitempty"`

	// Specifies the number of frameworks a report covers.
	NumberOfFrameworks *float64 `json:"numberOfFrameworks,omitempty" tf:"number_of_frameworks,omitempty"`

	// Identifies the report template for the report. Reports are built using a report template. The report templates are: RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT.
	ReportTemplate *string `json:"reportTemplate,omitempty" tf:"report_template,omitempty"`
}

type ReportSettingObservation struct {

	// Specifies the Amazon Resource Names (ARNs) of the frameworks a report covers.
	FrameworkArns []*string `json:"frameworkArns,omitempty" tf:"framework_arns,omitempty"`

	// Specifies the number of frameworks a report covers.
	NumberOfFrameworks *float64 `json:"numberOfFrameworks,omitempty" tf:"number_of_frameworks,omitempty"`

	// Identifies the report template for the report. Reports are built using a report template. The report templates are: RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT.
	ReportTemplate *string `json:"reportTemplate,omitempty" tf:"report_template,omitempty"`
}

type ReportSettingParameters struct {

	// Specifies the Amazon Resource Names (ARNs) of the frameworks a report covers.
	FrameworkArns []*string `json:"frameworkArns,omitempty" tf:"framework_arns,omitempty"`

	// Specifies the number of frameworks a report covers.
	NumberOfFrameworks *float64 `json:"numberOfFrameworks,omitempty" tf:"number_of_frameworks,omitempty"`

	// Identifies the report template for the report. Reports are built using a report template. The report templates are: RESOURCE_COMPLIANCE_REPORT | CONTROL_COMPLIANCE_REPORT | BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT.
	ReportTemplate *string `json:"reportTemplate,omitempty" tf:"report_template,omitempty"`
}

// ReportPlanSpec defines the desired state of ReportPlan
type ReportPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReportPlanParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ReportPlanInitParameters `json:"initProvider,omitempty"`
}

// ReportPlanStatus defines the observed state of ReportPlan.
type ReportPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReportPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReportPlan is the Schema for the ReportPlans API. Provides an AWS Backup Report Plan resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReportPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reportDeliveryChannel) || has(self.initProvider.reportDeliveryChannel)",message="%!s(MISSING) is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reportSetting) || has(self.initProvider.reportSetting)",message="%!s(MISSING) is a required parameter"
	Spec   ReportPlanSpec   `json:"spec"`
	Status ReportPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReportPlanList contains a list of ReportPlans
type ReportPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReportPlan `json:"items"`
}

// Repository type metadata.
var (
	ReportPlan_Kind             = "ReportPlan"
	ReportPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReportPlan_Kind}.String()
	ReportPlan_KindAPIVersion   = ReportPlan_Kind + "." + CRDGroupVersion.String()
	ReportPlan_GroupVersionKind = CRDGroupVersion.WithKind(ReportPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&ReportPlan{}, &ReportPlanList{})
}
