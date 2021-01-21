// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

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
	case "aws:codeartifact/domain:Domain":
		r, err = NewDomain(ctx, name, nil, pulumi.URN_(urn))
	case "aws:codeartifact/domainPermissions:DomainPermissions":
		r, err = NewDomainPermissions(ctx, name, nil, pulumi.URN_(urn))
	case "aws:codeartifact/repository:Repository":
		r, err = NewRepository(ctx, name, nil, pulumi.URN_(urn))
	case "aws:codeartifact/repositoryPermissionsPolicy:RepositoryPermissionsPolicy":
		r, err = NewRepositoryPermissionsPolicy(ctx, name, nil, pulumi.URN_(urn))
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
		"codeartifact/domain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codeartifact/domainPermissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codeartifact/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codeartifact/repositoryPermissionsPolicy",
		&module{version},
	)
}
