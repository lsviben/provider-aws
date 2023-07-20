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

// GetTerraformResourceType returns Terraform resource type for this ApprovalRuleTemplate
func (mg *ApprovalRuleTemplate) GetTerraformResourceType() string {
	return "aws_codecommit_approval_rule_template"
}

// GetConnectionDetailsMapping for this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) GetParameters() (map[string]any, error) {
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

// SetParameters for this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ApprovalRuleTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ApprovalRuleTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &ApprovalRuleTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ApprovalRuleTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this ApprovalRuleTemplate
func (tr *ApprovalRuleTemplate) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this ApprovalRuleTemplateAssociation
func (mg *ApprovalRuleTemplateAssociation) GetTerraformResourceType() string {
	return "aws_codecommit_approval_rule_template_association"
}

// GetConnectionDetailsMapping for this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) GetParameters() (map[string]any, error) {
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

// SetParameters for this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ApprovalRuleTemplateAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ApprovalRuleTemplateAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &ApprovalRuleTemplateAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ApprovalRuleTemplateAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this ApprovalRuleTemplateAssociation
func (tr *ApprovalRuleTemplateAssociation) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this Repository
func (mg *Repository) GetTerraformResourceType() string {
	return "aws_codecommit_repository"
}

// GetConnectionDetailsMapping for this Repository
func (tr *Repository) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Repository
func (tr *Repository) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Repository
func (tr *Repository) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Repository
func (tr *Repository) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Repository
func (tr *Repository) GetParameters() (map[string]any, error) {
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

// SetParameters for this Repository
func (tr *Repository) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Repository using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Repository) LateInitialize(attrs []byte) (bool, error) {
	params := &RepositoryParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Repository) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Repository
func (tr *Repository) GetIgnorableFields() ([]string, error) {
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

// GetTerraformResourceType returns Terraform resource type for this Trigger
func (mg *Trigger) GetTerraformResourceType() string {
	return "aws_codecommit_trigger"
}

// GetConnectionDetailsMapping for this Trigger
func (tr *Trigger) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Trigger
func (tr *Trigger) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Trigger
func (tr *Trigger) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Trigger
func (tr *Trigger) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Trigger
func (tr *Trigger) GetParameters() (map[string]any, error) {
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

// SetParameters for this Trigger
func (tr *Trigger) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Trigger using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Trigger) LateInitialize(attrs []byte) (bool, error) {
	params := &TriggerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Trigger) GetTerraformSchemaVersion() int {
	return 0
}

// GetIgnorableFields of this Trigger
func (tr *Trigger) GetIgnorableFields() ([]string, error) {
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
