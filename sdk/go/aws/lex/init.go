// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lex

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
	case "aws:lex/bot:Bot":
		r, err = NewBot(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lex/botAlias:BotAlias":
		r, err = NewBotAlias(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lex/intent:Intent":
		r, err = NewIntent(ctx, name, nil, pulumi.URN_(urn))
	case "aws:lex/slotType:SlotType":
		r, err = NewSlotType(ctx, name, nil, pulumi.URN_(urn))
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
		"lex/bot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lex/botAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lex/intent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lex/slotType",
		&module{version},
	)
}
