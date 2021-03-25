// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Patch Baseline data source. Useful if you wish to reuse the default baselines provided.
//
// ## Example Usage
//
// To retrieve a baseline provided by AWS:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "AWS-"
// 		opt1 := "CENTOS"
// 		_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
// 			NamePrefix:      &opt0,
// 			OperatingSystem: &opt1,
// 			Owner:           "AWS",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To retrieve a baseline on your account:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		opt1 := "MyCustomBaseline"
// 		opt2 := "WINDOWS"
// 		_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
// 			DefaultBaseline: &opt0,
// 			NamePrefix:      &opt1,
// 			OperatingSystem: &opt2,
// 			Owner:           "Self",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPatchBaseline(ctx *pulumi.Context, args *LookupPatchBaselineArgs, opts ...pulumi.InvokeOption) (*LookupPatchBaselineResult, error) {
	var rv LookupPatchBaselineResult
	err := ctx.Invoke("aws:ssm/getPatchBaseline:getPatchBaseline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPatchBaseline.
type LookupPatchBaselineArgs struct {
	// Filters the results against the baselines defaultBaseline field.
	DefaultBaseline *bool `pulumi:"defaultBaseline"`
	// Filter results by the baseline name prefix.
	NamePrefix *string `pulumi:"namePrefix"`
	// The specified OS for the baseline.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// The owner of the baseline. Valid values: `All`, `AWS`, `Self` (the current account).
	Owner string `pulumi:"owner"`
}

// A collection of values returned by getPatchBaseline.
type LookupPatchBaselineResult struct {
	DefaultBaseline *bool `pulumi:"defaultBaseline"`
	// The description of the baseline.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the baseline.
	Name            string  `pulumi:"name"`
	NamePrefix      *string `pulumi:"namePrefix"`
	OperatingSystem *string `pulumi:"operatingSystem"`
	Owner           string  `pulumi:"owner"`
}
