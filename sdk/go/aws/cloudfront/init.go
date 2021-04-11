// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

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
	case "aws:cloudfront/cachePolicy:CachePolicy":
		r, err = NewCachePolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/distribution:Distribution":
		r, err = NewDistribution(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/keyGroup:KeyGroup":
		r, err = NewKeyGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/originAccessIdentity:OriginAccessIdentity":
		r, err = NewOriginAccessIdentity(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/originRequestPolicy:OriginRequestPolicy":
		r, err = NewOriginRequestPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/publicKey:PublicKey":
		r, err = NewPublicKey(ctx, name, nil, pulumi.URN_(urn))
	case "aws:cloudfront/realtimeLogConfig:RealtimeLogConfig":
		r, err = NewRealtimeLogConfig(ctx, name, nil, pulumi.URN_(urn))
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
		"cloudfront/cachePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/distribution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/keyGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/originAccessIdentity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/originRequestPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/publicKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cloudfront/realtimeLogConfig",
		&module{version},
	)
}
