// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package directconnect

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	directconnect_sdkv2 "github.com/aws/aws-sdk-go-v2/service/directconnect"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceConnection,
			TypeName: "aws_dx_connection",
		},
		{
			Factory:  DataSourceGateway,
			TypeName: "aws_dx_gateway",
		},
		{
			Factory:  DataSourceLocation,
			TypeName: "aws_dx_location",
		},
		{
			Factory:  DataSourceLocations,
			TypeName: "aws_dx_locations",
		},
		{
			Factory:  DataSourceRouterConfiguration,
			TypeName: "aws_dx_router_configuration",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBGPPeer,
			TypeName: "aws_dx_bgp_peer",
			Name:     "BGP Peer",
		},
		{
			Factory:  resourceConnection,
			TypeName: "aws_dx_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnectionAssociation,
			TypeName: "aws_dx_connection_association",
			Name:     "Connection LAG Association",
		},
		{
			Factory:  ResourceConnectionConfirmation,
			TypeName: "aws_dx_connection_confirmation",
		},
		{
			Factory:  ResourceGateway,
			TypeName: "aws_dx_gateway",
		},
		{
			Factory:  ResourceGatewayAssociation,
			TypeName: "aws_dx_gateway_association",
		},
		{
			Factory:  ResourceGatewayAssociationProposal,
			TypeName: "aws_dx_gateway_association_proposal",
		},
		{
			Factory:  ResourceHostedConnection,
			TypeName: "aws_dx_hosted_connection",
		},
		{
			Factory:  ResourceHostedPrivateVirtualInterface,
			TypeName: "aws_dx_hosted_private_virtual_interface",
		},
		{
			Factory:  ResourceHostedPrivateVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_private_virtual_interface_accepter",
			Name:     "Hosted Private Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceHostedPublicVirtualInterface,
			TypeName: "aws_dx_hosted_public_virtual_interface",
		},
		{
			Factory:  ResourceHostedPublicVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_public_virtual_interface_accepter",
			Name:     "Hosted Public Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceHostedTransitVirtualInterface,
			TypeName: "aws_dx_hosted_transit_virtual_interface",
		},
		{
			Factory:  ResourceHostedTransitVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_transit_virtual_interface_accepter",
			Name:     "Hosted Transit Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceLag,
			TypeName: "aws_dx_lag",
			Name:     "LAG",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMacSecKeyAssociation,
			TypeName: "aws_dx_macsec_key_association",
		},
		{
			Factory:  ResourcePrivateVirtualInterface,
			TypeName: "aws_dx_private_virtual_interface",
			Name:     "Private Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourcePublicVirtualInterface,
			TypeName: "aws_dx_public_virtual_interface",
			Name:     "Public Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceTransitVirtualInterface,
			TypeName: "aws_dx_transit_virtual_interface",
			Name:     "Transit Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DirectConnect
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*directconnect_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return directconnect_sdkv2.NewFromConfig(cfg,
		directconnect_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
