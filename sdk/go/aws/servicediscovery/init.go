// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

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
	case "aws:servicediscovery/httpNamespace:HttpNamespace":
		r = &HttpNamespace{}
	case "aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace":
		r = &PrivateDnsNamespace{}
	case "aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace":
		r = &PublicDnsNamespace{}
	case "aws:servicediscovery/service:Service":
		r = &Service{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"servicediscovery/httpNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicediscovery/privateDnsNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicediscovery/publicDnsNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"servicediscovery/service",
		&module{version},
	)
}
