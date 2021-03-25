// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get all direct child organizational units under a parent organizational unit. This only provides immediate children, not all children.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		org, err := organizations.LookupOrganization(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.GetOrganizationalUnits(ctx, &organizations.GetOrganizationalUnitsArgs{
// 			ParentId: org.Roots[0].Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganizationalUnits(ctx *pulumi.Context, args *GetOrganizationalUnitsArgs, opts ...pulumi.InvokeOption) (*GetOrganizationalUnitsResult, error) {
	var rv GetOrganizationalUnitsResult
	err := ctx.Invoke("aws:organizations/getOrganizationalUnits:getOrganizationalUnits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationalUnits.
type GetOrganizationalUnitsArgs struct {
	// The parent ID of the organizational unit.
	ParentId string `pulumi:"parentId"`
}

// A collection of values returned by getOrganizationalUnits.
type GetOrganizationalUnitsResult struct {
	// List of child organizational units, which have the following attributes:
	Childrens []GetOrganizationalUnitsChildren `pulumi:"childrens"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ParentId string `pulumi:"parentId"`
}
