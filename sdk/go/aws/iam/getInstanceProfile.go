// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM instance profile. By using this data source, you can reference IAM
// instance profile properties without having to hard code ARNs as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_instance_profile.html.markdown.
func LookupInstanceProfile(ctx *pulumi.Context, args *LookupInstanceProfileArgs, opts ...pulumi.InvokeOption) (*LookupInstanceProfileResult, error) {
	var rv LookupInstanceProfileResult
	err := ctx.Invoke("aws:iam/getInstanceProfile:getInstanceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceProfile.
type LookupInstanceProfileArgs struct {
	// The friendly IAM instance profile name to match.
	Name string `pulumi:"name"`
}


// A collection of values returned by getInstanceProfile.
type LookupInstanceProfileResult struct {
	// The Amazon Resource Name (ARN) specifying the instance profile.
	Arn string `pulumi:"arn"`
	// The string representation of the date the instance profile
	// was created.
	CreateDate string `pulumi:"createDate"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The path to the instance profile.
	Path string `pulumi:"path"`
	// The role arn associated with this instance profile.
	RoleArn string `pulumi:"roleArn"`
	// The role id associated with this instance profile.
	RoleId string `pulumi:"roleId"`
	// The role name associated with this instance profile.
	RoleName string `pulumi:"roleName"`
}

