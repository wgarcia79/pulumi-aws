// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// !> **WARNING:** This data source was removed in version 2.0.0 of the Terraform AWS Provider. You can migrate existing configurations to the [`aws_kms_secrets` data source](https://www.terraform.io/docs/providers/aws/d/kms_secrets.html) following instructions available in the [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/aws/guides/version-2-upgrade.html#data-source-aws_kms_secret).
func LookupSecret(ctx *pulumi.Context, args *GetSecretArgs) (*GetSecretResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["secrets"] = args.Secrets
	}
	outputs, err := ctx.Invoke("aws:kms/getSecret:getSecret", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSecretResult{
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSecret.
type GetSecretArgs struct {
	Secrets interface{}
}

// A collection of values returned by getSecret.
type GetSecretResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
