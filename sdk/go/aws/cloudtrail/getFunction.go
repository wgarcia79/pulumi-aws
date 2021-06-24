// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about a CloudFront Function.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		functionName := cfg.Require("functionName")
// 		_, err := cloudtrail.GetFunction(ctx, &cloudtrail.GetFunctionArgs{
// 			Name: functionName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetFunction(ctx *pulumi.Context, args *GetFunctionArgs, opts ...pulumi.InvokeOption) (*GetFunctionResult, error) {
	var rv GetFunctionResult
	err := ctx.Invoke("aws:cloudtrail/getFunction:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunction.
type GetFunctionArgs struct {
	// Name of the CloudFront function.
	Name string `pulumi:"name"`
	// The function’s stage, either `DEVELOPMENT` or `LIVE`.
	Stage string `pulumi:"stage"`
}

// A collection of values returned by getFunction.
type GetFunctionResult struct {
	// Amazon Resource Name (ARN) identifying your CloudFront Function.
	Arn string `pulumi:"arn"`
	// Source code of the function
	Code string `pulumi:"code"`
	// Comment.
	Comment string `pulumi:"comment"`
	// ETag hash of the function
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// When this resource was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	Name             string `pulumi:"name"`
	// Identifier of the function's runtime.
	Runtime string `pulumi:"runtime"`
	Stage   string `pulumi:"stage"`
	// Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
	Status string `pulumi:"status"`
}
