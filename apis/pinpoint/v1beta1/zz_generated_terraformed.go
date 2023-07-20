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

// GetTerraformResourceType returns Terraform resource type for this App
func (mg *App) GetTerraformResourceType() string {
	return "aws_pinpoint_app"
}

// GetConnectionDetailsMapping for this App
func (tr *App) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this App
func (tr *App) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this App
func (tr *App) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this App
func (tr *App) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this App
func (tr *App) GetParameters() (map[string]any, error) {
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

// SetParameters for this App
func (tr *App) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this App using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *App) LateInitialize(attrs []byte) (bool, error) {
	params := &AppParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *App) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this App
func (tr *App) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this SMSChannel
func (mg *SMSChannel) GetTerraformResourceType() string {
	return "aws_pinpoint_sms_channel"
}

// GetConnectionDetailsMapping for this SMSChannel
func (tr *SMSChannel) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SMSChannel
func (tr *SMSChannel) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SMSChannel
func (tr *SMSChannel) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SMSChannel
func (tr *SMSChannel) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SMSChannel
func (tr *SMSChannel) GetParameters() (map[string]any, error) {
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

// SetParameters for this SMSChannel
func (tr *SMSChannel) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SMSChannel using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SMSChannel) LateInitialize(attrs []byte) (bool, error) {
	params := &SMSChannelParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SMSChannel) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this SMSChannel
func (tr *SMSChannel) GetIgnorableFields() ([]string, error) {
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
