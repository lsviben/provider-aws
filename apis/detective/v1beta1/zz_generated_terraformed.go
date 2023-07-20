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

// GetTerraformResourceType returns Terraform resource type for this Graph
func (mg *Graph) GetTerraformResourceType() string {
	return "aws_detective_graph"
}

// GetConnectionDetailsMapping for this Graph
func (tr *Graph) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Graph
func (tr *Graph) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Graph
func (tr *Graph) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Graph
func (tr *Graph) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Graph
func (tr *Graph) GetParameters() (map[string]any, error) {
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

// SetParameters for this Graph
func (tr *Graph) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Graph using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Graph) LateInitialize(attrs []byte) (bool, error) {
	params := &GraphParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Graph) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Graph
func (tr *Graph) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this InvitationAccepter
func (mg *InvitationAccepter) GetTerraformResourceType() string {
	return "aws_detective_invitation_accepter"
}

// GetConnectionDetailsMapping for this InvitationAccepter
func (tr *InvitationAccepter) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this InvitationAccepter
func (tr *InvitationAccepter) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this InvitationAccepter
func (tr *InvitationAccepter) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this InvitationAccepter
func (tr *InvitationAccepter) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this InvitationAccepter
func (tr *InvitationAccepter) GetParameters() (map[string]any, error) {
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

// SetParameters for this InvitationAccepter
func (tr *InvitationAccepter) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this InvitationAccepter using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *InvitationAccepter) LateInitialize(attrs []byte) (bool, error) {
	params := &InvitationAccepterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *InvitationAccepter) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this InvitationAccepter
func (tr *InvitationAccepter) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this Member
func (mg *Member) GetTerraformResourceType() string {
	return "aws_detective_member"
}

// GetConnectionDetailsMapping for this Member
func (tr *Member) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Member
func (tr *Member) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Member
func (tr *Member) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Member
func (tr *Member) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Member
func (tr *Member) GetParameters() (map[string]any, error) {
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

// SetParameters for this Member
func (tr *Member) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Member using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Member) LateInitialize(attrs []byte) (bool, error) {
	params := &MemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Member) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Member
func (tr *Member) GetIgnorableFields() ([]string, error) {
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
