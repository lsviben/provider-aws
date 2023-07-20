/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessPoint.
func (mg *AccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FileSystemIDRef,
		Selector:     mg.Spec.InitProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemID")
	}
	mg.Spec.InitProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FileSystemIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicy.
func (mg *BackupPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FileSystemIDRef,
		Selector:     mg.Spec.InitProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemID")
	}
	mg.Spec.InitProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FileSystemIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FileSystem.
func (mg *FileSystem) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.KMSKeyIDRef,
		Selector:     mg.Spec.InitProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta1.KeyList{},
			Managed: &v1beta1.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyID")
	}
	mg.Spec.InitProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FileSystemPolicy.
func (mg *FileSystemPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FileSystemIDRef,
		Selector:     mg.Spec.InitProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemID")
	}
	mg.Spec.InitProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FileSystemIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MountTarget.
func (mg *MountTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FileSystemIDRef,
		Selector:     mg.Spec.InitProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FileSystemID")
	}
	mg.Spec.InitProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FileSystemIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.InitProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroups")
	}
	mg.Spec.InitProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SubnetIDRef,
		Selector:     mg.Spec.InitProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SubnetID")
	}
	mg.Spec.InitProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ReplicationConfiguration.
func (mg *ReplicationConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceFileSystemID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SourceFileSystemIDRef,
		Selector:     mg.Spec.ForProvider.SourceFileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceFileSystemID")
	}
	mg.Spec.ForProvider.SourceFileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceFileSystemIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceFileSystemID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SourceFileSystemIDRef,
		Selector:     mg.Spec.InitProvider.SourceFileSystemIDSelector,
		To: reference.To{
			List:    &FileSystemList{},
			Managed: &FileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SourceFileSystemID")
	}
	mg.Spec.InitProvider.SourceFileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SourceFileSystemIDRef = rsp.ResolvedReference

	return nil
}
