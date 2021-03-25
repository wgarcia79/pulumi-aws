// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

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
	case "aws:gamelift/alias:Alias":
		r, err = NewAlias(ctx, name, nil, pulumi.URN_(urn))
	case "aws:gamelift/build:Build":
		r, err = NewBuild(ctx, name, nil, pulumi.URN_(urn))
	case "aws:gamelift/fleet:Fleet":
		r, err = NewFleet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:gamelift/gameSessionQueue:GameSessionQueue":
		r, err = NewGameSessionQueue(ctx, name, nil, pulumi.URN_(urn))
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
		"gamelift/alias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"gamelift/build",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"gamelift/fleet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"gamelift/gameSessionQueue",
		&module{version},
	)
}
