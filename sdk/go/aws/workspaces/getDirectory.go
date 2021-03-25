// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about an AWS WorkSpaces directory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/workspaces"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := workspaces.LookupDirectory(ctx, &workspaces.LookupDirectoryArgs{
// 			DirectoryId: "d-9067783251",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDirectory(ctx *pulumi.Context, args *LookupDirectoryArgs, opts ...pulumi.InvokeOption) (*LookupDirectoryResult, error) {
	var rv LookupDirectoryResult
	err := ctx.Invoke("aws:workspaces/getDirectory:getDirectory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDirectory.
type LookupDirectoryArgs struct {
	// The directory identifier for registration in WorkSpaces service.
	DirectoryId string `pulumi:"directoryId"`
	// A map of tags assigned to the WorkSpaces directory.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDirectory.
type LookupDirectoryResult struct {
	// The directory alias.
	Alias string `pulumi:"alias"`
	// The user name for the service account.
	CustomerUserName string `pulumi:"customerUserName"`
	DirectoryId      string `pulumi:"directoryId"`
	// The name of the directory.
	DirectoryName string `pulumi:"directoryName"`
	// The directory type.
	DirectoryType string `pulumi:"directoryType"`
	// The IP addresses of the DNS servers for the directory.
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces to make calls to other services, such as Amazon EC2, on your behalf.
	IamRoleId string `pulumi:"iamRoleId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The identifiers of the IP access control groups associated with the directory.
	IpGroupIds []string `pulumi:"ipGroupIds"`
	// The registration code for the directory. This is the code that users enter in their Amazon WorkSpaces client application to connect to the directory.
	RegistrationCode string `pulumi:"registrationCode"`
	// The permissions to enable or disable self-service capabilities.
	SelfServicePermissions []GetDirectorySelfServicePermission `pulumi:"selfServicePermissions"`
	// The identifiers of the subnets where the directory resides.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags assigned to the WorkSpaces directory.
	Tags map[string]string `pulumi:"tags"`
	// (Optional) Specifies which devices and operating systems users can use to access their WorkSpaces. Defined below.
	WorkspaceAccessProperties []GetDirectoryWorkspaceAccessProperty `pulumi:"workspaceAccessProperties"`
	// The default properties that are used for creating WorkSpaces. Defined below.
	WorkspaceCreationProperties GetDirectoryWorkspaceCreationProperties `pulumi:"workspaceCreationProperties"`
	// The identifier of the security group that is assigned to new WorkSpaces. Defined below.
	WorkspaceSecurityGroupId string `pulumi:"workspaceSecurityGroupId"`
}
