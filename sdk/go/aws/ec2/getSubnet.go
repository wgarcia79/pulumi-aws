// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.Subnet` provides details about a specific VPC subnet.
//
// This resource can prove useful when a module accepts a subnet ID as an input variable and needs to, for example, determine the ID of the VPC that the subnet belongs to.
//
// ## Example Usage
// ### Filter Example
//
// If you want to match against tag `Name`, use:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
// 			Filters: []ec2.GetSubnetFilter{
// 				ec2.GetSubnetFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"yakdriver",
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
func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("aws:ec2/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	// Availability zone where the subnet must reside.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// ID of the Availability Zone for the subnet.
	AvailabilityZoneId *string `pulumi:"availabilityZoneId"`
	// CIDR block of the desired subnet.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Whether the desired subnet must be the default subnet for its associated availability zone.
	DefaultForAz *bool `pulumi:"defaultForAz"`
	// Configuration block. Detailed below.
	Filters []GetSubnetFilter `pulumi:"filters"`
	// ID of the specific subnet to retrieve.
	Id *string `pulumi:"id"`
	// IPv6 CIDR block of the desired subnet.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// State that the desired subnet must have.
	State *string `pulumi:"state"`
	// Map of tags, each pair of which must exactly match a pair on the desired subnet.
	Tags map[string]string `pulumi:"tags"`
	// ID of the VPC that the desired subnet belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	// ARN of the subnet.
	Arn string `pulumi:"arn"`
	// Whether an IPv6 address is assigned on creation.
	AssignIpv6AddressOnCreation bool   `pulumi:"assignIpv6AddressOnCreation"`
	AvailabilityZone            string `pulumi:"availabilityZone"`
	AvailabilityZoneId          string `pulumi:"availabilityZoneId"`
	// Available IP addresses of the subnet.
	AvailableIpAddressCount int    `pulumi:"availableIpAddressCount"`
	CidrBlock               string `pulumi:"cidrBlock"`
	// Identifier of customer owned IPv4 address pool.
	CustomerOwnedIpv4Pool string            `pulumi:"customerOwnedIpv4Pool"`
	DefaultForAz          bool              `pulumi:"defaultForAz"`
	Filters               []GetSubnetFilter `pulumi:"filters"`
	Id                    string            `pulumi:"id"`
	Ipv6CidrBlock         string            `pulumi:"ipv6CidrBlock"`
	// Association ID of the IPv6 CIDR block.
	Ipv6CidrBlockAssociationId string `pulumi:"ipv6CidrBlockAssociationId"`
	// Whether customer owned IP addresses are assigned on network interface creation.
	MapCustomerOwnedIpOnLaunch bool `pulumi:"mapCustomerOwnedIpOnLaunch"`
	// Whether public IP addresses are assigned on instance launch.
	MapPublicIpOnLaunch bool `pulumi:"mapPublicIpOnLaunch"`
	// ARN of the Outpost.
	OutpostArn string `pulumi:"outpostArn"`
	// ID of the AWS account that owns the subnet.
	OwnerId string            `pulumi:"ownerId"`
	State   string            `pulumi:"state"`
	Tags    map[string]string `pulumi:"tags"`
	VpcId   string            `pulumi:"vpcId"`
}
