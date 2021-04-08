// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

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
	case "aws:redshift/cluster:Cluster":
		r = &Cluster{}
	case "aws:redshift/eventSubscription:EventSubscription":
		r = &EventSubscription{}
	case "aws:redshift/parameterGroup:ParameterGroup":
		r = &ParameterGroup{}
	case "aws:redshift/securityGroup:SecurityGroup":
		r = &SecurityGroup{}
	case "aws:redshift/snapshotCopyGrant:SnapshotCopyGrant":
		r = &SnapshotCopyGrant{}
	case "aws:redshift/snapshotSchedule:SnapshotSchedule":
		r = &SnapshotSchedule{}
	case "aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation":
		r = &SnapshotScheduleAssociation{}
	case "aws:redshift/subnetGroup:SubnetGroup":
		r = &SubnetGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/eventSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/parameterGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/securityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/snapshotCopyGrant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/snapshotSchedule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/snapshotScheduleAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"redshift/subnetGroup",
		&module{version},
	)
}
