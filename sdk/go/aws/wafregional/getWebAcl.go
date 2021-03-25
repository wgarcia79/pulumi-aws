// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `wafregional.WebAcl` Retrieves a WAF Regional Web ACL Resource Id.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafregional"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafregional.LookupWebAcl(ctx, &wafregional.LookupWebAclArgs{
// 			Name: "tfWAFRegionalWebACL",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupWebAcl(ctx *pulumi.Context, args *LookupWebAclArgs, opts ...pulumi.InvokeOption) (*LookupWebAclResult, error) {
	var rv LookupWebAclResult
	err := ctx.Invoke("aws:wafregional/getWebAcl:getWebAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebAcl.
type LookupWebAclArgs struct {
	// The name of the WAF Regional Web ACL.
	Name string `pulumi:"name"`
}

// A collection of values returned by getWebAcl.
type LookupWebAclResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
