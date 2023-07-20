/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/secretsmanager/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DataSource.
func (mg *DataSource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketNameRef,
				Selector:     mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketNameSelector,
				To: reference.To{
					List:    &v1beta1.BucketList{},
					Managed: &v1beta1.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketName")
			}
			mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Configuration[i3].S3Configuration[i4].BucketNameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsRef,
						Selector:     mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsSelector,
						To: reference.To{
							List:    &v1beta11.SecretList{},
							Managed: &v1beta11.Secret{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials")
					}
					mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsRef,
					Selector:     mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsSelector,
					To: reference.To{
						List:    &v1beta11.SecretList{},
						Managed: &v1beta11.Secret{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials")
				}
				mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IndexIDRef,
		Selector:     mg.Spec.ForProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IndexID")
	}
	mg.Spec.ForProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Configuration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketName),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketNameRef,
				Selector:     mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketNameSelector,
				To: reference.To{
					List:    &v1beta1.BucketList{},
					Managed: &v1beta1.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketName")
			}
			mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Configuration[i3].S3Configuration[i4].BucketNameRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsRef,
						Selector:     mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsSelector,
						To: reference.To{
							List:    &v1beta11.SecretList{},
							Managed: &v1beta11.Secret{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials")
					}
					mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].Credentials = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].AuthenticationConfiguration[i5].BasicAuthentication[i6].CredentialsRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsRef,
					Selector:     mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsSelector,
					To: reference.To{
						List:    &v1beta11.SecretList{},
						Managed: &v1beta11.Secret{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials")
				}
				mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].Credentials = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Configuration[i3].WebCrawlerConfiguration[i4].ProxyConfiguration[i5].CredentialsRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.IndexIDRef,
		Selector:     mg.Spec.InitProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IndexID")
	}
	mg.Spec.InitProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Experience.
func (mg *Experience) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IndexIDRef,
		Selector:     mg.Spec.ForProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IndexID")
	}
	mg.Spec.ForProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.IndexIDRef,
		Selector:     mg.Spec.InitProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IndexID")
	}
	mg.Spec.InitProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Index.
func (mg *Index) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuerySuggestionsBlockList.
func (mg *QuerySuggestionsBlockList) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IndexIDRef,
		Selector:     mg.Spec.ForProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IndexID")
	}
	mg.Spec.ForProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceS3Path[i3].Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SourceS3Path[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.SourceS3Path[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SourceS3Path[i3].Bucket")
		}
		mg.Spec.ForProvider.SourceS3Path[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SourceS3Path[i3].BucketRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.IndexIDRef,
		Selector:     mg.Spec.InitProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IndexID")
	}
	mg.Spec.InitProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceS3Path[i3].Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SourceS3Path[i3].BucketRef,
			Selector:     mg.Spec.InitProvider.SourceS3Path[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SourceS3Path[i3].Bucket")
		}
		mg.Spec.InitProvider.SourceS3Path[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SourceS3Path[i3].BucketRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Thesaurus.
func (mg *Thesaurus) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IndexIDRef,
		Selector:     mg.Spec.ForProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IndexID")
	}
	mg.Spec.ForProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceS3Path[i3].Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SourceS3Path[i3].BucketRef,
			Selector:     mg.Spec.ForProvider.SourceS3Path[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SourceS3Path[i3].Bucket")
		}
		mg.Spec.ForProvider.SourceS3Path[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SourceS3Path[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceS3Path[i3].Key),
			Extract:      resource.ExtractParamPath("key", false),
			Reference:    mg.Spec.ForProvider.SourceS3Path[i3].KeyRef,
			Selector:     mg.Spec.ForProvider.SourceS3Path[i3].KeySelector,
			To: reference.To{
				List:    &v1beta1.ObjectList{},
				Managed: &v1beta1.Object{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SourceS3Path[i3].Key")
		}
		mg.Spec.ForProvider.SourceS3Path[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SourceS3Path[i3].KeyRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IndexID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.IndexIDRef,
		Selector:     mg.Spec.InitProvider.IndexIDSelector,
		To: reference.To{
			List:    &IndexList{},
			Managed: &Index{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IndexID")
	}
	mg.Spec.InitProvider.IndexID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IndexIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta12.RoleList{},
			Managed: &v1beta12.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceS3Path[i3].Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SourceS3Path[i3].BucketRef,
			Selector:     mg.Spec.InitProvider.SourceS3Path[i3].BucketSelector,
			To: reference.To{
				List:    &v1beta1.BucketList{},
				Managed: &v1beta1.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SourceS3Path[i3].Bucket")
		}
		mg.Spec.InitProvider.SourceS3Path[i3].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SourceS3Path[i3].BucketRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.SourceS3Path); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SourceS3Path[i3].Key),
			Extract:      resource.ExtractParamPath("key", false),
			Reference:    mg.Spec.InitProvider.SourceS3Path[i3].KeyRef,
			Selector:     mg.Spec.InitProvider.SourceS3Path[i3].KeySelector,
			To: reference.To{
				List:    &v1beta1.ObjectList{},
				Managed: &v1beta1.Object{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SourceS3Path[i3].Key")
		}
		mg.Spec.InitProvider.SourceS3Path[i3].Key = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SourceS3Path[i3].KeyRef = rsp.ResolvedReference

	}

	return nil
}
