// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a specific AWS Direct Connect location in the current AWS Region.
// These are the locations that can be specified when configuring [`directconnect.Connection`](https://www.terraform.io/docs/providers/aws/r/dx_connection.html) or [`directconnect.LinkAggregationGroup`](https://www.terraform.io/docs/providers/aws/r/dx_lag.html) resources.
//
// > **Note:** This data source is different from the [`directconnect.getLocations`](https://www.terraform.io/docs/providers/aws/d/dx_locations.html) data source which retrieves information about all the AWS Direct Connect locations in the current AWS Region.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/directconnect"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := directconnect.GetLocation(ctx, &directconnect.GetLocationArgs{
// 			LocationCode: "CS32A-24FL",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLocation(ctx *pulumi.Context, args *GetLocationArgs, opts ...pulumi.InvokeOption) (*GetLocationResult, error) {
	var rv GetLocationResult
	err := ctx.Invoke("aws:directconnect/getLocation:getLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocation.
type GetLocationArgs struct {
	// The code for the location to retrieve.
	LocationCode string `pulumi:"locationCode"`
}

// A collection of values returned by getLocation.
type GetLocationResult struct {
	// The available port speeds for the location.
	AvailablePortSpeeds []string `pulumi:"availablePortSpeeds"`
	// The names of the service providers for the location.
	AvailableProviders []string `pulumi:"availableProviders"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	LocationCode string `pulumi:"locationCode"`
	// The name of the location. This includes the name of the colocation partner and the physical site of the building.
	LocationName string `pulumi:"locationName"`
}
