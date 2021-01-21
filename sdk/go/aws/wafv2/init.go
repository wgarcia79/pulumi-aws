// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

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
	case "aws:wafv2/ipSet:IpSet":
		r, err = NewIpSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafv2/regexPatternSet:RegexPatternSet":
		r, err = NewRegexPatternSet(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafv2/ruleGroup:RuleGroup":
		r, err = NewRuleGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafv2/webAcl:WebAcl":
		r, err = NewWebAcl(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafv2/webAclAssociation:WebAclAssociation":
		r, err = NewWebAclAssociation(ctx, name, nil, pulumi.URN_(urn))
	case "aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration":
		r, err = NewWebAclLoggingConfiguration(ctx, name, nil, pulumi.URN_(urn))
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
		"wafv2/ipSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafv2/regexPatternSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafv2/ruleGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafv2/webAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafv2/webAclAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"wafv2/webAclLoggingConfiguration",
		&module{version},
	)
}
