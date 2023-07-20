/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SMSChannel.
func (mg *SMSChannel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApplicationID),
		Extract:      resource.ExtractParamPath("application_id", true),
		Reference:    mg.Spec.ForProvider.ApplicationIDRef,
		Selector:     mg.Spec.ForProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &AppList{},
			Managed: &App{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApplicationID")
	}
	mg.Spec.ForProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApplicationIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ApplicationID),
		Extract:      resource.ExtractParamPath("application_id", true),
		Reference:    mg.Spec.InitProvider.ApplicationIDRef,
		Selector:     mg.Spec.InitProvider.ApplicationIDSelector,
		To: reference.To{
			List:    &AppList{},
			Managed: &App{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ApplicationID")
	}
	mg.Spec.InitProvider.ApplicationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ApplicationIDRef = rsp.ResolvedReference

	return nil
}
