// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

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
	case "aws:appsync/apiKey:ApiKey":
		r, err = NewApiKey(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appsync/dataSource:DataSource":
		r, err = NewDataSource(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appsync/function:Function":
		r, err = NewFunction(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appsync/graphQLApi:GraphQLApi":
		r, err = NewGraphQLApi(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appsync/resolver:Resolver":
		r, err = NewResolver(ctx, name, nil, pulumi.URN_(urn))
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
		"appsync/apiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appsync/dataSource",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appsync/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appsync/graphQLApi",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appsync/resolver",
		&module{version},
	)
}
