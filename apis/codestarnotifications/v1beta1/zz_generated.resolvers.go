/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/codecommit/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this NotificationRule.
func (mg *NotificationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Resource),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ResourceRef,
		Selector:     mg.Spec.ForProvider.ResourceSelector,
		To: reference.To{
			List:    &v1beta1.RepositoryList{},
			Managed: &v1beta1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Resource")
	}
	mg.Spec.ForProvider.Resource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Target); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Target[i3].Address),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Target[i3].AddressRef,
			Selector:     mg.Spec.ForProvider.Target[i3].AddressSelector,
			To: reference.To{
				List:    &v1beta11.TopicList{},
				Managed: &v1beta11.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Target[i3].Address")
		}
		mg.Spec.ForProvider.Target[i3].Address = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Target[i3].AddressRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Resource),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.ResourceRef,
		Selector:     mg.Spec.InitProvider.ResourceSelector,
		To: reference.To{
			List:    &v1beta1.RepositoryList{},
			Managed: &v1beta1.Repository{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Resource")
	}
	mg.Spec.InitProvider.Resource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Target); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Target[i3].Address),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.Target[i3].AddressRef,
			Selector:     mg.Spec.InitProvider.Target[i3].AddressSelector,
			To: reference.To{
				List:    &v1beta11.TopicList{},
				Managed: &v1beta11.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Target[i3].Address")
		}
		mg.Spec.InitProvider.Target[i3].Address = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Target[i3].AddressRef = rsp.ResolvedReference

	}

	return nil
}
