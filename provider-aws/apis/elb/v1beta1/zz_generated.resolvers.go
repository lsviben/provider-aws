/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ELB),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ELBRef,
		Selector:     mg.Spec.ForProvider.ELBSelector,
		To: reference.To{
			List:    &ELBList{},
			Managed: &ELB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ELB")
	}
	mg.Spec.ForProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ELBRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &v1beta1.InstanceList{},
			Managed: &v1beta1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ELB.
func (mg *ELB) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Instances),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.InstancesRefs,
		Selector:      mg.Spec.ForProvider.InstancesSelector,
		To: reference.To{
			List:    &v1beta1.InstanceList{},
			Managed: &v1beta1.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instances")
	}
	mg.Spec.ForProvider.Instances = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.InstancesRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetsRefs,
		Selector:      mg.Spec.ForProvider.SubnetsSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetsRefs = mrsp.ResolvedReferences

	return nil
}
