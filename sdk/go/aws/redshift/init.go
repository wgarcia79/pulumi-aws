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
		r, err = NewCluster(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/eventSubscription:EventSubscription":
		r, err = NewEventSubscription(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/parameterGroup:ParameterGroup":
		r, err = NewParameterGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/securityGroup:SecurityGroup":
		r, err = NewSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/snapshotCopyGrant:SnapshotCopyGrant":
		r, err = NewSnapshotCopyGrant(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/snapshotSchedule:SnapshotSchedule":
		r, err = NewSnapshotSchedule(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/snapshotScheduleAssociation:SnapshotScheduleAssociation":
		r, err = NewSnapshotScheduleAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:redshift/subnetGroup:SubnetGroup":
		r, err = NewSubnetGroup(ctx, name, nil, pulumi.URN_(urn))
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
