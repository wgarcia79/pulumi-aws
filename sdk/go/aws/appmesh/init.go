// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:appmesh/gatewayRoute:GatewayRoute":
		r = &GatewayRoute{}
	case "aws:appmesh/mesh:Mesh":
		r = &Mesh{}
	case "aws:appmesh/route:Route":
		r = &Route{}
	case "aws:appmesh/virtualGateway:VirtualGateway":
		r = &VirtualGateway{}
	case "aws:appmesh/virtualNode:VirtualNode":
		r = &VirtualNode{}
	case "aws:appmesh/virtualRouter:VirtualRouter":
		r = &VirtualRouter{}
	case "aws:appmesh/virtualService:VirtualService":
		r = &VirtualService{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/gatewayRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/mesh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualNode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualRouter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualService",
		&module{version},
	)
}
