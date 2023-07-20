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

// GetTerraformResourceType returns Terraform resource type for this DeliveryStream
func (mg *DeliveryStream) GetTerraformResourceType() string {
	return "aws_kinesis_firehose_delivery_stream"
}

// GetConnectionDetailsMapping for this DeliveryStream
func (tr *DeliveryStream) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"http_endpoint_configuration[*].access_key": "spec.forProvider.httpEndpointConfiguration[*].accessKeySecretRef", "redshift_configuration[*].password": "spec.forProvider.redshiftConfiguration[*].passwordSecretRef"}
}

// GetObservation of this DeliveryStream
func (tr *DeliveryStream) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DeliveryStream
func (tr *DeliveryStream) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DeliveryStream
func (tr *DeliveryStream) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DeliveryStream
func (tr *DeliveryStream) GetParameters() (map[string]any, error) {
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

// SetParameters for this DeliveryStream
func (tr *DeliveryStream) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DeliveryStream using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DeliveryStream) LateInitialize(attrs []byte) (bool, error) {
	params := &DeliveryStreamParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("ServerSideEncryption"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DeliveryStream) GetTerraformSchemaVersion() int {
	return 1
}

// GetIgnorableFields of this DeliveryStream
func (tr *DeliveryStream) GetIgnorableFields() ([]string, error) {
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
