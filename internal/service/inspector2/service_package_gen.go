// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package inspector2

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	inspector2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDelegatedAdminAccount,
			TypeName: "aws_inspector2_delegated_admin_account",
			Name:     "Delegated Admin Account",
		},
		{
			Factory:  ResourceEnabler,
			TypeName: "aws_inspector2_enabler",
		},
		{
			Factory:  ResourceMemberAssociation,
			TypeName: "aws_inspector2_member_association",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_inspector2_organization_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Inspector2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*inspector2_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return inspector2_sdkv2.NewFromConfig(cfg,
		inspector2_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
