// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ec2.getPrefixList` provides details about a specific prefix list (PL)
// in the current region.
//
// This can be used both to validate a prefix list given in a variable
// and to obtain the CIDR blocks (IP address ranges) for the associated
// AWS service. The latter may be useful e.g. for adding network ACL
// rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		privateS3VpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "privateS3VpcEndpoint", &ec2.VpcEndpointArgs{
// 			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
// 			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := ec2.NewNetworkAcl(ctx, "bar", &ec2.NetworkAclArgs{
// 			VpcId: pulumi.Any(aws_vpc.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewNetworkAclRule(ctx, "privateS3NetworkAclRule", &ec2.NetworkAclRuleArgs{
// 			NetworkAclId: bar.ID(),
// 			RuleNumber:   pulumi.Int(200),
// 			Egress:       pulumi.Bool(false),
// 			Protocol:     pulumi.String("tcp"),
// 			RuleAction:   pulumi.String("allow"),
// 			CidrBlock: privateS3PrefixList.ApplyT(func(privateS3PrefixList ec2.GetPrefixListResult) (string, error) {
// 				return privateS3PrefixList.CidrBlocks[0], nil
// 			}).(pulumi.StringOutput),
// 			FromPort: pulumi.Int(443),
// 			ToPort:   pulumi.Int(443),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.GetPrefixList(ctx, &ec2.GetPrefixListArgs{
// 			Filters: []ec2.GetPrefixListFilter{
// 				ec2.GetPrefixListFilter{
// 					Name: "prefix-list-id",
// 					Values: []string{
// 						"pl-68a54001",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.getPrefixList has been deprecated in favor of aws.ec2.getPrefixList
func GetPrefixList(ctx *pulumi.Context, args *GetPrefixListArgs, opts ...pulumi.InvokeOption) (*GetPrefixListResult, error) {
	var rv GetPrefixListResult
	err := ctx.Invoke("aws:index/getPrefixList:getPrefixList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixList.
type GetPrefixListArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetPrefixListFilter `pulumi:"filters"`
	// The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
	Name *string `pulumi:"name"`
	// The ID of the prefix list to select.
	PrefixListId *string `pulumi:"prefixListId"`
}

// A collection of values returned by getPrefixList.
type GetPrefixListResult struct {
	// The list of CIDR blocks for the AWS service associated with the prefix list.
	CidrBlocks []string              `pulumi:"cidrBlocks"`
	Filters    []GetPrefixListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the selected prefix list.
	Name         string  `pulumi:"name"`
	PrefixListId *string `pulumi:"prefixListId"`
}
