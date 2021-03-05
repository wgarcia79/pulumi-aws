// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

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
	case "aws:route53/delegationSet:DelegationSet":
		r, err = NewDelegationSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/healthCheck:HealthCheck":
		r, err = NewHealthCheck(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/hostedZoneDnsSec:HostedZoneDnsSec":
		r, err = NewHostedZoneDnsSec(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/keySigningKey:KeySigningKey":
		r, err = NewKeySigningKey(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/queryLog:QueryLog":
		r, err = NewQueryLog(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/record:Record":
		r, err = NewRecord(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverDnsSecConfig:ResolverDnsSecConfig":
		r, err = NewResolverDnsSecConfig(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverEndpoint:ResolverEndpoint":
		r, err = NewResolverEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig":
		r, err = NewResolverQueryLogConfig(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation":
		r, err = NewResolverQueryLogConfigAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverRule:ResolverRule":
		r, err = NewResolverRule(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/resolverRuleAssociation:ResolverRuleAssociation":
		r, err = NewResolverRuleAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization":
		r, err = NewVpcAssociationAuthorization(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/zone:Zone":
		r, err = NewZone(ctx, name, nil, pulumi.URN_(urn))
	case "aws:route53/zoneAssociation:ZoneAssociation":
		r, err = NewZoneAssociation(ctx, name, nil, pulumi.URN_(urn))
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
		"route53/delegationSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/healthCheck",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/hostedZoneDnsSec",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/keySigningKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/queryLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/record",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverDnsSecConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverQueryLogConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverQueryLogConfigAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/resolverRuleAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/vpcAssociationAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"route53/zoneAssociation",
		&module{version},
	)
}
