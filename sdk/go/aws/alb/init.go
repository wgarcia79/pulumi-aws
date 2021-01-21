// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

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
	case "aws:alb/listener:Listener":
		r, err = NewListener(ctx, name, nil, pulumi.URN_(urn))
	case "aws:alb/listenerCertificate:ListenerCertificate":
		r, err = NewListenerCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "aws:alb/listenerRule:ListenerRule":
		r, err = NewListenerRule(ctx, name, nil, pulumi.URN_(urn))
	case "aws:alb/loadBalancer:LoadBalancer":
		r, err = NewLoadBalancer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:alb/targetGroup:TargetGroup":
		r, err = NewTargetGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:alb/targetGroupAttachment:TargetGroupAttachment":
		r, err = NewTargetGroupAttachment(ctx, name, nil, pulumi.URN_(urn))
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
		"alb/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"alb/listenerCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"alb/listenerRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"alb/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"alb/targetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"alb/targetGroupAttachment",
		&module{version},
	)
}
