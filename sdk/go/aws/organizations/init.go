// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

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
	case "aws:organizations/account:Account":
		r, err = NewAccount(ctx, name, nil, pulumi.URN_(urn))
	case "aws:organizations/organization:Organization":
		r, err = NewOrganization(ctx, name, nil, pulumi.URN_(urn))
	case "aws:organizations/organizationalUnit:OrganizationalUnit":
		r, err = NewOrganizationalUnit(ctx, name, nil, pulumi.URN_(urn))
	case "aws:organizations/policy:Policy":
		r, err = NewPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:organizations/policyAttachment:PolicyAttachment":
		r, err = NewPolicyAttachment(ctx, name, nil, pulumi.URN_(urn))
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
		"organizations/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"organizations/organization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"organizations/organizationalUnit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"organizations/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"organizations/policyAttachment",
		&module{version},
	)
}
