// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

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
	case "aws:glue/catalogDatabase:CatalogDatabase":
		r, err = NewCatalogDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/catalogTable:CatalogTable":
		r, err = NewCatalogTable(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/classifier:Classifier":
		r, err = NewClassifier(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/connection:Connection":
		r, err = NewConnection(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/crawler:Crawler":
		r, err = NewCrawler(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings":
		r, err = NewDataCatalogEncryptionSettings(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/devEndpoint:DevEndpoint":
		r, err = NewDevEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/job:Job":
		r, err = NewJob(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/mLTransform:MLTransform":
		r, err = NewMLTransform(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/partition:Partition":
		r, err = NewPartition(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/registry:Registry":
		r, err = NewRegistry(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/resourcePolicy:ResourcePolicy":
		r, err = NewResourcePolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/schema:Schema":
		r, err = NewSchema(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/securityConfiguration:SecurityConfiguration":
		r, err = NewSecurityConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/trigger:Trigger":
		r, err = NewTrigger(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/userDefinedFunction:UserDefinedFunction":
		r, err = NewUserDefinedFunction(ctx, name, nil, pulumi.URN_(urn))
	case "aws:glue/workflow:Workflow":
		r, err = NewWorkflow(ctx, name, nil, pulumi.URN_(urn))
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
		"glue/catalogDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/catalogTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/classifier",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/crawler",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/dataCatalogEncryptionSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/devEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/job",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/mLTransform",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/partition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/registry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/resourcePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/schema",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/securityConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/trigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/userDefinedFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"glue/workflow",
		&module{version},
	)
}
