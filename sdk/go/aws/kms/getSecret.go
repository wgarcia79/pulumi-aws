// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// !> **WARNING:** This data source was removed in version 2.0.0 of the AWS Provider. You can migrate existing configurations to the [`kms.getSecrets` data source](https://www.terraform.io/docs/providers/aws/d/kms_secrets.html) following instructions available in the [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/aws/guides/version-2-upgrade.html#data-source-aws_kms_secret).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_secret.html.markdown.
func GetSecret(ctx *pulumi.Context, args *GetSecretArgs, opts ...pulumi.InvokeOption) (*GetSecretResult, error) {
	var rv GetSecretResult
	err := ctx.Invoke("aws:kms/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type GetSecretArgs struct {
	Secrets []GetSecretSecret `pulumi:"secrets"`
}


// A collection of values returned by getSecret.
type GetSecretResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Secrets []GetSecretSecret `pulumi:"secrets"`
}

