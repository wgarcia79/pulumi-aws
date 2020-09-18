// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about DocumentDB orderable DB instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/docdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "3.6.0"
// 		opt1 := "na"
// 		_, err := docdb.GetOrderableDbInstance(ctx, &docdb.GetOrderableDbInstanceArgs{
// 			Engine:        "docdb",
// 			EngineVersion: &opt0,
// 			LicenseModel:  &opt1,
// 			PreferredInstanceClasses: []string{
// 				"db.r5.large",
// 				"db.r4.large",
// 				"db.t3.medium",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrderableDbInstance(ctx *pulumi.Context, args *GetOrderableDbInstanceArgs, opts ...pulumi.InvokeOption) (*GetOrderableDbInstanceResult, error) {
	var rv GetOrderableDbInstanceResult
	err := ctx.Invoke("aws:docdb/getOrderableDbInstance:getOrderableDbInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderableDbInstance.
type GetOrderableDbInstanceArgs struct {
	// DB engine. Engine values include `docdb`.
	Engine string `pulumi:"engine"`
	// Version of the DB engine. For example, `3.6.0`.
	EngineVersion *string `pulumi:"engineVersion"`
	// DB instance class. Examples of classes are `db.r5.12xlarge`, `db.r5.24xlarge`, `db.r5.2xlarge`, `db.r5.4xlarge`, `db.r5.large`, `db.r5.xlarge`, and `db.t3.medium`.
	InstanceClass *string `pulumi:"instanceClass"`
	// License model. Examples of license models are `general-public-license`, `na`, `bring-your-own-license`, and `amazon-license`.
	LicenseModel *string `pulumi:"licenseModel"`
	// Ordered list of preferred DocumentDB DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
	PreferredInstanceClasses []string `pulumi:"preferredInstanceClasses"`
	// Boolean that indicates whether to show only VPC or non-VPC offerings.
	Vpc *bool `pulumi:"vpc"`
}

// A collection of values returned by getOrderableDbInstance.
type GetOrderableDbInstanceResult struct {
	// Availability zones where the instance is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	Engine            string   `pulumi:"engine"`
	EngineVersion     string   `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string   `pulumi:"id"`
	InstanceClass            string   `pulumi:"instanceClass"`
	LicenseModel             string   `pulumi:"licenseModel"`
	PreferredInstanceClasses []string `pulumi:"preferredInstanceClasses"`
	Vpc                      bool     `pulumi:"vpc"`
}
