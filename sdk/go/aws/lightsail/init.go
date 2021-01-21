// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:lightsail/domain:Domain":
		r, err = NewDomain(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lightsail/instance:Instance":
		r, err = NewInstance(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lightsail/keyPair:KeyPair":
		r, err = NewKeyPair(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lightsail/staticIp:StaticIp":
		r, err = NewStaticIp(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lightsail/staticIpAttachment:StaticIpAttachment":
		r, err = NewStaticIpAttachment(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"lightsail/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lightsail/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lightsail/keyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lightsail/staticIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lightsail/staticIpAttachment",
		&module{version},
	)
}
