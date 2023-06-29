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

type BatchTargetObservation struct {

	// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
	ArraySize *float64 `json:"arraySize,omitempty" tf:"array_size,omitempty"`

	// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
	JobAttempts *float64 `json:"jobAttempts,omitempty" tf:"job_attempts,omitempty"`

	// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
	JobDefinition *string `json:"jobDefinition,omitempty" tf:"job_definition,omitempty"`

	// The name to use for this execution of the job, if the target is an AWS Batch job.
	JobName *string `json:"jobName,omitempty" tf:"job_name,omitempty"`
}

type BatchTargetParameters struct {

	// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
	// +kubebuilder:validation:Optional
	ArraySize *float64 `json:"arraySize,omitempty" tf:"array_size,omitempty"`

	// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
	// +kubebuilder:validation:Optional
	JobAttempts *float64 `json:"jobAttempts,omitempty" tf:"job_attempts,omitempty"`

	// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
	// +kubebuilder:validation:Required
	JobDefinition *string `json:"jobDefinition" tf:"job_definition,omitempty"`

	// The name to use for this execution of the job, if the target is an AWS Batch job.
	// +kubebuilder:validation:Required
	JobName *string `json:"jobName" tf:"job_name,omitempty"`
}

type CapacityProviderStrategyObservation struct {

	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to 0.
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Short name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider,omitempty" tf:"capacity_provider,omitempty"`

	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type CapacityProviderStrategyParameters struct {

	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined. Defaults to 0.
	// +kubebuilder:validation:Optional
	Base *float64 `json:"base,omitempty" tf:"base,omitempty"`

	// Short name of the capacity provider.
	// +kubebuilder:validation:Required
	CapacityProvider *string `json:"capacityProvider" tf:"capacity_provider,omitempty"`

	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The weight value is taken into consideration after the base value, if defined, is satisfied.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type DeadLetterConfigObservation struct {

	// - ARN of the SQS queue specified as the target for the dead-letter queue.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type DeadLetterConfigParameters struct {

	// - ARN of the SQS queue specified as the target for the dead-letter queue.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type EcsTargetObservation struct {

	// The capacity provider strategy to use for the task. If a capacity_provider_strategy specified, the launch_type parameter must be omitted. If no capacity_provider_strategy or launch_type is specified, the default capacity provider strategy for the cluster is used. Can be one or more. See below.
	CapacityProviderStrategy []CapacityProviderStrategyObservation `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the task.
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags,omitempty"`

	// Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task.
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command,omitempty"`

	// Specifies an ECS task group for the task. The maximum length is 255 characters.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values include: EC2, EXTERNAL, or FARGATE.
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type,omitempty"`

	// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launch_type is FARGATE because the awsvpc mode is required for Fargate tasks.
	NetworkConfiguration []NetworkConfigurationObservation `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime). See Below.
	PlacementConstraint []PlacementConstraintObservation `json:"placementConstraint,omitempty" tf:"placement_constraint,omitempty"`

	// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see AWS Fargate Platform Versions.
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	// Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. The only valid value is: TASK_DEFINITION.
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// A map of tags to assign to ecs resources.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The number of tasks to create based on the TaskDefinition. Defaults to 1.
	TaskCount *float64 `json:"taskCount,omitempty" tf:"task_count,omitempty"`

	// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
	TaskDefinitionArn *string `json:"taskDefinitionArn,omitempty" tf:"task_definition_arn,omitempty"`
}

type EcsTargetParameters struct {

	// The capacity provider strategy to use for the task. If a capacity_provider_strategy specified, the launch_type parameter must be omitted. If no capacity_provider_strategy or launch_type is specified, the default capacity provider strategy for the cluster is used. Can be one or more. See below.
	// +kubebuilder:validation:Optional
	CapacityProviderStrategy []CapacityProviderStrategyParameters `json:"capacityProviderStrategy,omitempty" tf:"capacity_provider_strategy,omitempty"`

	// Specifies whether to enable Amazon ECS managed tags for the task.
	// +kubebuilder:validation:Optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags,omitempty"`

	// Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task.
	// +kubebuilder:validation:Optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command,omitempty"`

	// Specifies an ECS task group for the task. The maximum length is 255 characters.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task. Valid values include: EC2, EXTERNAL, or FARGATE.
	// +kubebuilder:validation:Optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type,omitempty"`

	// Use this if the ECS task uses the awsvpc network mode. This specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. Required if launch_type is FARGATE because the awsvpc mode is required for Fargate tasks.
	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime). See Below.
	// +kubebuilder:validation:Optional
	PlacementConstraint []PlacementConstraintParameters `json:"placementConstraint,omitempty" tf:"placement_constraint,omitempty"`

	// Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as 1.1.0. This is used only if LaunchType is FARGATE. For more information about valid platform versions, see AWS Fargate Platform Versions.
	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`

	// Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. The only valid value is: TASK_DEFINITION.
	// +kubebuilder:validation:Optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// A map of tags to assign to ecs resources.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The number of tasks to create based on the TaskDefinition. Defaults to 1.
	// +kubebuilder:validation:Optional
	TaskCount *float64 `json:"taskCount,omitempty" tf:"task_count,omitempty"`

	// The ARN of the task definition to use if the event target is an Amazon ECS cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecs/v1beta1.TaskDefinition
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TaskDefinitionArn *string `json:"taskDefinitionArn,omitempty" tf:"task_definition_arn,omitempty"`

	// Reference to a TaskDefinition in ecs to populate taskDefinitionArn.
	// +kubebuilder:validation:Optional
	TaskDefinitionArnRef *v1.Reference `json:"taskDefinitionArnRef,omitempty" tf:"-"`

	// Selector for a TaskDefinition in ecs to populate taskDefinitionArn.
	// +kubebuilder:validation:Optional
	TaskDefinitionArnSelector *v1.Selector `json:"taskDefinitionArnSelector,omitempty" tf:"-"`
}

type HTTPTargetObservation struct {

	// Enables you to specify HTTP headers to add to the request.
	HeaderParameters map[string]*string `json:"headerParameters,omitempty" tf:"header_parameters,omitempty"`

	// The list of values that correspond sequentially to any path variables in your endpoint ARN (for example arn:aws:execute-api:us-east-1:123456:myapi/*/POST/pets/*).
	PathParameterValues []*string `json:"pathParameterValues,omitempty" tf:"path_parameter_values,omitempty"`

	// Represents keys/values of query string parameters that are appended to the invoked endpoint.
	QueryStringParameters map[string]*string `json:"queryStringParameters,omitempty" tf:"query_string_parameters,omitempty"`
}

type HTTPTargetParameters struct {

	// Enables you to specify HTTP headers to add to the request.
	// +kubebuilder:validation:Optional
	HeaderParameters map[string]*string `json:"headerParameters,omitempty" tf:"header_parameters,omitempty"`

	// The list of values that correspond sequentially to any path variables in your endpoint ARN (for example arn:aws:execute-api:us-east-1:123456:myapi/*/POST/pets/*).
	// +kubebuilder:validation:Optional
	PathParameterValues []*string `json:"pathParameterValues,omitempty" tf:"path_parameter_values,omitempty"`

	// Represents keys/values of query string parameters that are appended to the invoked endpoint.
	// +kubebuilder:validation:Optional
	QueryStringParameters map[string]*string `json:"queryStringParameters,omitempty" tf:"query_string_parameters,omitempty"`
}

type InputTransformerObservation struct {

	// Key value pairs specified in the form of JSONPath (for example, time = $.time)
	InputPaths map[string]*string `json:"inputPaths,omitempty" tf:"input_paths,omitempty"`

	// Template to customize data sent to the target. Must be valid JSON. To send a string value, the string value must include double quotes.g., "\"Your string goes here.\\nA new line.\""
	InputTemplate *string `json:"inputTemplate,omitempty" tf:"input_template,omitempty"`
}

type InputTransformerParameters struct {

	// Key value pairs specified in the form of JSONPath (for example, time = $.time)
	// +kubebuilder:validation:Optional
	InputPaths map[string]*string `json:"inputPaths,omitempty" tf:"input_paths,omitempty"`

	// Template to customize data sent to the target. Must be valid JSON. To send a string value, the string value must include double quotes.g., "\"Your string goes here.\\nA new line.\""
	// +kubebuilder:validation:Required
	InputTemplate *string `json:"inputTemplate" tf:"input_template,omitempty"`
}

type KinesisTargetObservation struct {

	// The JSON path to be extracted from the event and used as the partition key.
	PartitionKeyPath *string `json:"partitionKeyPath,omitempty" tf:"partition_key_path,omitempty"`
}

type KinesisTargetParameters struct {

	// The JSON path to be extracted from the event and used as the partition key.
	// +kubebuilder:validation:Optional
	PartitionKeyPath *string `json:"partitionKeyPath,omitempty" tf:"partition_key_path,omitempty"`
}

type NetworkConfigurationObservation struct {

	// Assign a public IP address to the ENI (Fargate launch type only). Valid values are true or false. Defaults to false.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The subnets associated with the task or service.
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

type NetworkConfigurationParameters struct {

	// Assign a public IP address to the ENI (Fargate launch type only). Valid values are true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The subnets associated with the task or service.
	// +kubebuilder:validation:Required
	Subnets []*string `json:"subnets" tf:"subnets,omitempty"`
}

type PlacementConstraintObservation struct {

	// Cluster Query Language expression to apply to the constraint. Does not need to be specified for the distinctInstance type. For more information, see Cluster Query Language in the Amazon EC2 Container Service Developer Guide.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Type of constraint. The only valid values at this time are memberOf and distinctInstance.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PlacementConstraintParameters struct {

	// Cluster Query Language expression to apply to the constraint. Does not need to be specified for the distinctInstance type. For more information, see Cluster Query Language in the Amazon EC2 Container Service Developer Guide.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// Type of constraint. The only valid values at this time are memberOf and distinctInstance.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedshiftTargetObservation struct {

	// The database user name.
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The name of the database.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// The SQL statement text to run.
	SQL *string `json:"sql,omitempty" tf:"sql,omitempty"`

	// The name or ARN of the secret that enables access to the database.
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// The name of the SQL statement.
	StatementName *string `json:"statementName,omitempty" tf:"statement_name,omitempty"`

	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent *bool `json:"withEvent,omitempty" tf:"with_event,omitempty"`
}

type RedshiftTargetParameters struct {

	// The database user name.
	// +kubebuilder:validation:Optional
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The name of the database.
	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// The SQL statement text to run.
	// +kubebuilder:validation:Optional
	SQL *string `json:"sql,omitempty" tf:"sql,omitempty"`

	// The name or ARN of the secret that enables access to the database.
	// +kubebuilder:validation:Optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn,omitempty"`

	// The name of the SQL statement.
	// +kubebuilder:validation:Optional
	StatementName *string `json:"statementName,omitempty" tf:"statement_name,omitempty"`

	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	// +kubebuilder:validation:Optional
	WithEvent *bool `json:"withEvent,omitempty" tf:"with_event,omitempty"`
}

type RetryPolicyObservation struct {

	// The age in seconds to continue to make retry attempts.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// maximum number of retry attempts to make before the request fails
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`
}

type RetryPolicyParameters struct {

	// The age in seconds to continue to make retry attempts.
	// +kubebuilder:validation:Optional
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// maximum number of retry attempts to make before the request fails
	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`
}

type RunCommandTargetsObservation struct {

	// Can be either tag:tag-key or InstanceIds.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// If Key is tag:tag-key, Values is a list of tag values. If Key is InstanceIds, Values is a list of Amazon EC2 instance IDs.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type RunCommandTargetsParameters struct {

	// Can be either tag:tag-key or InstanceIds.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// If Key is tag:tag-key, Values is a list of tag values. If Key is InstanceIds, Values is a list of Amazon EC2 instance IDs.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SqsTargetObservation struct {

	// The FIFO message group ID to use as the target.
	MessageGroupID *string `json:"messageGroupId,omitempty" tf:"message_group_id,omitempty"`
}

type SqsTargetParameters struct {

	// The FIFO message group ID to use as the target.
	// +kubebuilder:validation:Optional
	MessageGroupID *string `json:"messageGroupId,omitempty" tf:"message_group_id,omitempty"`
}

type TargetObservation struct {

	// The Amazon Resource Name (ARN) of the target.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget []BatchTargetObservation `json:"batchTarget,omitempty" tf:"batch_target,omitempty"`

	// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
	DeadLetterConfig []DeadLetterConfigObservation `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget []EcsTargetObservation `json:"ecsTarget,omitempty" tf:"ecs_target,omitempty"`

	// The event bus to associate with the rule. If you omit this, the default event bus is used.
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
	HTTPTarget []HTTPTargetObservation `json:"httpTarget,omitempty" tf:"http_target,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Valid JSON text passed to the target. Conflicts with input_path and input_transformer.
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// The value of the JSONPath that is used for extracting part of the matched event when passing it to the target. Conflicts with input and input_transformer.
	InputPath *string `json:"inputPath,omitempty" tf:"input_path,omitempty"`

	// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with input and input_path.
	InputTransformer []InputTransformerObservation `json:"inputTransformer,omitempty" tf:"input_transformer,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget []KinesisTargetObservation `json:"kinesisTarget,omitempty" tf:"kinesis_target,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
	RedshiftTarget []RedshiftTargetObservation `json:"redshiftTarget,omitempty" tf:"redshift_target,omitempty"`

	// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
	RetryPolicy []RetryPolicyObservation `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if ecs_target is used or target in arn is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The name of the rule you want to add targets to.
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`

	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []RunCommandTargetsObservation `json:"runCommandTargets,omitempty" tf:"run_command_targets,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget []SqsTargetObservation `json:"sqsTarget,omitempty" tf:"sqs_target,omitempty"`

	// The unique target assignment ID. If missing, will generate a random, unique id.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`
}

type TargetParameters struct {

	// The Amazon Resource Name (ARN) of the target.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	BatchTarget []BatchTargetParameters `json:"batchTarget,omitempty" tf:"batch_target,omitempty"`

	// Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config,omitempty"`

	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	EcsTarget []EcsTargetParameters `json:"ecsTarget,omitempty" tf:"ecs_target,omitempty"`

	// The event bus to associate with the rule. If you omit this, the default event bus is used.
	// +crossplane:generate:reference:type=Bus
	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name,omitempty"`

	// Reference to a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameRef *v1.Reference `json:"eventBusNameRef,omitempty" tf:"-"`

	// Selector for a Bus to populate eventBusName.
	// +kubebuilder:validation:Optional
	EventBusNameSelector *v1.Selector `json:"eventBusNameSelector,omitempty" tf:"-"`

	// Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
	// +kubebuilder:validation:Optional
	HTTPTarget []HTTPTargetParameters `json:"httpTarget,omitempty" tf:"http_target,omitempty"`

	// Valid JSON text passed to the target. Conflicts with input_path and input_transformer.
	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// The value of the JSONPath that is used for extracting part of the matched event when passing it to the target. Conflicts with input and input_transformer.
	// +kubebuilder:validation:Optional
	InputPath *string `json:"inputPath,omitempty" tf:"input_path,omitempty"`

	// Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with input and input_path.
	// +kubebuilder:validation:Optional
	InputTransformer []InputTransformerParameters `json:"inputTransformer,omitempty" tf:"input_transformer,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	KinesisTarget []KinesisTargetParameters `json:"kinesisTarget,omitempty" tf:"kinesis_target,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	RedshiftTarget []RedshiftTargetParameters `json:"redshiftTarget,omitempty" tf:"redshift_target,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	RetryPolicy []RetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if ecs_target is used or target in arn is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// The name of the rule you want to add targets to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudwatchevents/v1beta1.Rule
	// +kubebuilder:validation:Optional
	Rule *string `json:"rule,omitempty" tf:"rule,omitempty"`

	// Reference to a Rule in cloudwatchevents to populate rule.
	// +kubebuilder:validation:Optional
	RuleRef *v1.Reference `json:"ruleRef,omitempty" tf:"-"`

	// Selector for a Rule in cloudwatchevents to populate rule.
	// +kubebuilder:validation:Optional
	RuleSelector *v1.Selector `json:"ruleSelector,omitempty" tf:"-"`

	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	// +kubebuilder:validation:Optional
	RunCommandTargets []RunCommandTargetsParameters `json:"runCommandTargets,omitempty" tf:"run_command_targets,omitempty"`

	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	// +kubebuilder:validation:Optional
	SqsTarget []SqsTargetParameters `json:"sqsTarget,omitempty" tf:"sqs_target,omitempty"`

	// The unique target assignment ID. If missing, will generate a random, unique id.
	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`
}

// TargetSpec defines the desired state of Target
type TargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetParameters `json:"forProvider"`
}

// TargetStatus defines the observed state of Target.
type TargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Target is the Schema for the Targets API. Provides an EventBridge Target resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.arn)",message="arn is a required parameter"
	Spec   TargetSpec   `json:"spec"`
	Status TargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetList contains a list of Targets
type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target `json:"items"`
}

// Repository type metadata.
var (
	Target_Kind             = "Target"
	Target_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Target_Kind}.String()
	Target_KindAPIVersion   = Target_Kind + "." + CRDGroupVersion.String()
	Target_GroupVersionKind = CRDGroupVersion.WithKind(Target_Kind)
)

func init() {
	SchemeBuilder.Register(&Target{}, &TargetList{})
}
