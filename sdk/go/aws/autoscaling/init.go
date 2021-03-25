// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

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
	case "aws:autoscaling/attachment:Attachment":
		r, err = NewAttachment(ctx, name, nil, pulumi.URN_(urn))
	case "aws:autoscaling/group:Group":
		r, err = NewGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:autoscaling/lifecycleHook:LifecycleHook":
		r, err = NewLifecycleHook(ctx, name, nil, pulumi.URN_(urn))
	case "aws:autoscaling/notification:Notification":
		r, err = NewNotification(ctx, name, nil, pulumi.URN_(urn))
	case "aws:autoscaling/policy:Policy":
		r, err = NewPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "aws:autoscaling/schedule:Schedule":
		r, err = NewSchedule(ctx, name, nil, pulumi.URN_(urn))
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
		"autoscaling/attachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"autoscaling/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"autoscaling/lifecycleHook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"autoscaling/notification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"autoscaling/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"autoscaling/schedule",
		&module{version},
	)
}
