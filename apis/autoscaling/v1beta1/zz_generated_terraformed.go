/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/imdario/mergo"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Attachment
func (mg *Attachment) GetTerraformResourceType() string {
	return "aws_autoscaling_attachment"
}

// GetConnectionDetailsMapping for this Attachment
func (tr *Attachment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Attachment
func (tr *Attachment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Attachment
func (tr *Attachment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Attachment
func (tr *Attachment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Attachment
func (tr *Attachment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this Attachment
func (tr *Attachment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Attachment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Attachment) LateInitialize(attrs []byte) (bool, error) {
	params := &AttachmentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Attachment) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Attachment
func (tr *Attachment) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this AutoscalingGroup
func (mg *AutoscalingGroup) GetTerraformResourceType() string {
	return "aws_autoscaling_group"
}

// GetConnectionDetailsMapping for this AutoscalingGroup
func (tr *AutoscalingGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AutoscalingGroup
func (tr *AutoscalingGroup) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AutoscalingGroup
func (tr *AutoscalingGroup) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AutoscalingGroup
func (tr *AutoscalingGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AutoscalingGroup
func (tr *AutoscalingGroup) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this AutoscalingGroup
func (tr *AutoscalingGroup) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AutoscalingGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AutoscalingGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &AutoscalingGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("AvailabilityZones"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AutoscalingGroup) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this AutoscalingGroup
func (tr *AutoscalingGroup) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this GroupTag
func (mg *GroupTag) GetTerraformResourceType() string {
	return "aws_autoscaling_group_tag"
}

// GetConnectionDetailsMapping for this GroupTag
func (tr *GroupTag) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this GroupTag
func (tr *GroupTag) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this GroupTag
func (tr *GroupTag) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this GroupTag
func (tr *GroupTag) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this GroupTag
func (tr *GroupTag) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this GroupTag
func (tr *GroupTag) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this GroupTag using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *GroupTag) LateInitialize(attrs []byte) (bool, error) {
	params := &GroupTagParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *GroupTag) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this GroupTag
func (tr *GroupTag) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this LifecycleHook
func (mg *LifecycleHook) GetTerraformResourceType() string {
	return "aws_autoscaling_lifecycle_hook"
}

// GetConnectionDetailsMapping for this LifecycleHook
func (tr *LifecycleHook) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LifecycleHook
func (tr *LifecycleHook) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LifecycleHook
func (tr *LifecycleHook) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LifecycleHook
func (tr *LifecycleHook) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LifecycleHook
func (tr *LifecycleHook) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this LifecycleHook
func (tr *LifecycleHook) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LifecycleHook using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LifecycleHook) LateInitialize(attrs []byte) (bool, error) {
	params := &LifecycleHookParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LifecycleHook) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this LifecycleHook
func (tr *LifecycleHook) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this Notification
func (mg *Notification) GetTerraformResourceType() string {
	return "aws_autoscaling_notification"
}

// GetConnectionDetailsMapping for this Notification
func (tr *Notification) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Notification
func (tr *Notification) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Notification
func (tr *Notification) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Notification
func (tr *Notification) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Notification
func (tr *Notification) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this Notification
func (tr *Notification) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Notification using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Notification) LateInitialize(attrs []byte) (bool, error) {
	params := &NotificationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Notification) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Notification
func (tr *Notification) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this Policy
func (mg *Policy) GetTerraformResourceType() string {
	return "aws_autoscaling_policy"
}

// GetConnectionDetailsMapping for this Policy
func (tr *Policy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Policy
func (tr *Policy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Policy
func (tr *Policy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Policy
func (tr *Policy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Policy
func (tr *Policy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this Policy
func (tr *Policy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Policy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Policy) LateInitialize(attrs []byte) (bool, error) {
	params := &PolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Policy) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Policy
func (tr *Policy) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this Schedule
func (mg *Schedule) GetTerraformResourceType() string {
	return "aws_autoscaling_schedule"
}

// GetConnectionDetailsMapping for this Schedule
func (tr *Schedule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Schedule
func (tr *Schedule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Schedule
func (tr *Schedule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Schedule
func (tr *Schedule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Schedule
func (tr *Schedule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this Schedule
func (tr *Schedule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Schedule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Schedule) LateInitialize(attrs []byte) (bool, error) {
	params := &ScheduleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Schedule) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Schedule
func (tr *Schedule) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}

// GetTerraformResourceType returns Terraform resource type for this LaunchConfiguration
func (mg *LaunchConfiguration) GetTerraformResourceType() string {
	return "aws_launch_configuration"
}

// GetConnectionDetailsMapping for this LaunchConfiguration
func (tr *LaunchConfiguration) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LaunchConfiguration
func (tr *LaunchConfiguration) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LaunchConfiguration
func (tr *LaunchConfiguration) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LaunchConfiguration
func (tr *LaunchConfiguration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LaunchConfiguration
func (tr *LaunchConfiguration) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return base, mergo.Merge(&base, initBase, mergo.WithSliceDeepCopy)
}

// SetParameters for this LaunchConfiguration
func (tr *LaunchConfiguration) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LaunchConfiguration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LaunchConfiguration) LateInitialize(attrs []byte) (bool, error) {
	params := &LaunchConfigurationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LaunchConfiguration) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this LaunchConfiguration
func (tr *LaunchConfiguration) GetIgnorableFields() ([]string, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	err = json.TFParser.Unmarshal(p, &base)
	if err != nil {
		return nil, err
	}

	i, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	initBase := map[string]any{}
	err = json.TFParser.Unmarshal(i, &initBase)
	if err != nil {
		return nil, err
	}

	return resource.GetIgnoredFields(base, initBase), nil
}
