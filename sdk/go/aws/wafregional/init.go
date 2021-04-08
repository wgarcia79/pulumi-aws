// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

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
	case "aws:wafregional/byteMatchSet:ByteMatchSet":
		r = &ByteMatchSet{}
	case "aws:wafregional/geoMatchSet:GeoMatchSet":
		r = &GeoMatchSet{}
	case "aws:wafregional/ipSet:IpSet":
		r = &IpSet{}
	case "aws:wafregional/rateBasedRule:RateBasedRule":
		r = &RateBasedRule{}
	case "aws:wafregional/regexMatchSet:RegexMatchSet":
		r = &RegexMatchSet{}
	case "aws:wafregional/regexPatternSet:RegexPatternSet":
		r = &RegexPatternSet{}
	case "aws:wafregional/rule:Rule":
		r = &Rule{}
	case "aws:wafregional/ruleGroup:RuleGroup":
		r = &RuleGroup{}
	case "aws:wafregional/sizeConstraintSet:SizeConstraintSet":
		r = &SizeConstraintSet{}
	case "aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet":
		r = &SqlInjectionMatchSet{}
	case "aws:wafregional/webAcl:WebAcl":
		r = &WebAcl{}
	case "aws:wafregional/webAclAssociation:WebAclAssociation":
		r = &WebAclAssociation{}
	case "aws:wafregional/xssMatchSet:XssMatchSet":
		r = &XssMatchSet{}
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
		"wafregional/byteMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/geoMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/ipSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/rateBasedRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/regexMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/regexPatternSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/ruleGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/sizeConstraintSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/sqlInjectionMatchSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/webAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/webAclAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafregional/xssMatchSet",
		&module{version},
	)
}
