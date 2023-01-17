// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_s3_account_public_access_block":      dataSourceAccountPublicAccessBlock,
		"aws_s3control_multi_region_access_point": dataSourceMultiRegionAccessPoint,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_s3_access_point":                             resourceAccessPoint,
		"aws_s3_account_public_access_block":              resourceAccountPublicAccessBlock,
		"aws_s3control_access_point_policy":               resourceAccessPointPolicy,
		"aws_s3control_bucket":                            resourceBucket,
		"aws_s3control_bucket_lifecycle_configuration":    resourceBucketLifecycleConfiguration,
		"aws_s3control_bucket_policy":                     resourceBucketPolicy,
		"aws_s3control_multi_region_access_point":         resourceMultiRegionAccessPoint,
		"aws_s3control_multi_region_access_point_policy":  resourceMultiRegionAccessPointPolicy,
		"aws_s3control_object_lambda_access_point":        resourceObjectLambdaAccessPoint,
		"aws_s3control_object_lambda_access_point_policy": resourceObjectLambdaAccessPointPolicy,
		"aws_s3control_storage_lens_configuration":        resourceStorageLensConfiguration,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return "s3control"
}

var (
	_sp                                = &servicePackage{}
	ServicePackage intf.ServicePackage = _sp
)
