// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

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
	case "aws:datasync/agent:Agent":
		r, err = NewAgent(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/efsLocation:EfsLocation":
		r, err = NewEfsLocation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/locationFsxWindows:LocationFsxWindows":
		r, err = NewLocationFsxWindows(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/locationSmb:LocationSmb":
		r, err = NewLocationSmb(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/nfsLocation:NfsLocation":
		r, err = NewNfsLocation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/s3Location:S3Location":
		r, err = NewS3Location(ctx, name, nil, pulumi.URN_(urn))
	case "aws:datasync/task:Task":
		r, err = NewTask(ctx, name, nil, pulumi.URN_(urn))
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
		"datasync/agent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/efsLocation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/locationFsxWindows",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/locationSmb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/nfsLocation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/s3Location",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"datasync/task",
		&module{version},
	)
}
