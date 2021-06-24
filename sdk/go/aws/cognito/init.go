// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

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
	case "aws:cognito/identityPool:IdentityPool":
		r = &IdentityPool{}
	case "aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment":
		r = &IdentityPoolRoleAttachment{}
	case "aws:cognito/identityProvider:IdentityProvider":
		r = &IdentityProvider{}
	case "aws:cognito/resourceServer:ResourceServer":
		r = &ResourceServer{}
	case "aws:cognito/userGroup:UserGroup":
		r = &UserGroup{}
	case "aws:cognito/userPool:UserPool":
		r = &UserPool{}
	case "aws:cognito/userPoolClient:UserPoolClient":
		r = &UserPoolClient{}
	case "aws:cognito/userPoolDomain:UserPoolDomain":
		r = &UserPoolDomain{}
	case "aws:cognito/userPoolUICustomization:UserPoolUICustomization":
		r = &UserPoolUICustomization{}
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
		"cognito/identityPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/identityPoolRoleAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/identityProvider",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/resourceServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cognito/userPoolUICustomization",
		&module{version},
	)
}
