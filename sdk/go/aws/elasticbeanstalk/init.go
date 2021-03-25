// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

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
	case "aws:elasticbeanstalk/application:Application":
		r, err = NewApplication(ctx, name, nil, pulumi.URN_(urn))
	case "aws:elasticbeanstalk/applicationVersion:ApplicationVersion":
		r, err = NewApplicationVersion(ctx, name, nil, pulumi.URN_(urn))
	case "aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate":
		r, err = NewConfigurationTemplate(ctx, name, nil, pulumi.URN_(urn))
	case "aws:elasticbeanstalk/environment:Environment":
		r, err = NewEnvironment(ctx, name, nil, pulumi.URN_(urn))
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
		"elasticbeanstalk/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticbeanstalk/applicationVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticbeanstalk/configurationTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticbeanstalk/environment",
		&module{version},
	)
}
