// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get permissions for a principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
//
// > **NOTE:** This data source deals with explicitly granted permissions. Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).
//
// ## Example Usage
// ### Permissions For A Glue Catalog Database
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lakeformation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lakeformation.LookupPermissions(ctx, &lakeformation.LookupPermissionsArgs{
// 			Principal: aws_iam_role.Workflow_role.Arn,
// 			Database: lakeformation.GetPermissionsDatabase{
// 				Name:      aws_glue_catalog_database.Test.Name,
// 				CatalogId: "110376042874",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPermissions(ctx *pulumi.Context, args *LookupPermissionsArgs, opts ...pulumi.InvokeOption) (*LookupPermissionsResult, error) {
	var rv LookupPermissionsResult
	err := ctx.Invoke("aws:lakeformation/getPermissions:getPermissions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPermissions.
type LookupPermissionsArgs struct {
	// Identifier for the Data Catalog. By default, it is the account ID of the caller.
	CatalogId *string `pulumi:"catalogId"`
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource *bool `pulumi:"catalogResource"`
	// Configuration block for a data location resource. Detailed below.
	DataLocation *GetPermissionsDataLocation `pulumi:"dataLocation"`
	// Configuration block for a database resource. Detailed below.
	Database *GetPermissionsDatabase `pulumi:"database"`
	// Principal to be granted the permissions on the resource. Supported principals are IAM users or IAM roles.
	Principal string `pulumi:"principal"`
	// Configuration block for a table resource. Detailed below.
	Table *GetPermissionsTable `pulumi:"table"`
	// Configuration block for a table with columns resource. Detailed below.
	TableWithColumns *GetPermissionsTableWithColumns `pulumi:"tableWithColumns"`
}

// A collection of values returned by getPermissions.
type LookupPermissionsResult struct {
	CatalogId       *string                    `pulumi:"catalogId"`
	CatalogResource *bool                      `pulumi:"catalogResource"`
	DataLocation    GetPermissionsDataLocation `pulumi:"dataLocation"`
	Database        GetPermissionsDatabase     `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of permissions granted to the principal. For details on permissions, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions []string `pulumi:"permissions"`
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions []string                       `pulumi:"permissionsWithGrantOptions"`
	Principal                   string                         `pulumi:"principal"`
	Table                       GetPermissionsTable            `pulumi:"table"`
	TableWithColumns            GetPermissionsTableWithColumns `pulumi:"tableWithColumns"`
}
