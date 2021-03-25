// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a Service Quotas Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicequotas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicequotas.GetService(ctx, &servicequotas.GetServiceArgs{
// 			ServiceName: "Amazon Virtual Private Cloud (Amazon VPC)",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("aws:servicequotas/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// Service name to lookup within Service Quotas. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Code of the service.
	ServiceCode string `pulumi:"serviceCode"`
	ServiceName string `pulumi:"serviceName"`
}
