// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route 53 Resolver endpoint resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53.NewResolverEndpoint(ctx, "foo", &route53.ResolverEndpointArgs{
// 			Direction: pulumi.String("INBOUND"),
// 			SecurityGroupIds: pulumi.StringArray{
// 				pulumi.Any(aws_security_group.Sg1.Id),
// 				pulumi.Any(aws_security_group.Sg2.Id),
// 			},
// 			IpAddresses: route53.ResolverEndpointIpAddressArray{
// 				&route53.ResolverEndpointIpAddressArgs{
// 					SubnetId: pulumi.Any(aws_subnet.Sn1.Id),
// 				},
// 				&route53.ResolverEndpointIpAddressArgs{
// 					SubnetId: pulumi.Any(aws_subnet.Sn2.Id),
// 					Ip:       pulumi.String("10.0.64.4"),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Prod"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
//  Route 53 Resolver endpoints can be imported using the Route 53 Resolver endpoint ID, e.g.
//
// ```sh
//  $ pulumi import aws:route53/resolverEndpoint:ResolverEndpoint foo rslvr-in-abcdef01234567890
// ```
type ResolverEndpoint struct {
	pulumi.CustomResourceState

	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringOutput `pulumi:"hostVpcId"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayOutput `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.IpAddresses == nil {
		return nil, errors.New("invalid value for required argument 'IpAddresses'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	var resource ResolverEndpoint
	err := ctx.RegisterResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverEndpointState, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	var resource ResolverEndpoint
	err := ctx.ReadResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type resolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn *string `pulumi:"arn"`
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction *string `pulumi:"direction"`
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId *string `pulumi:"hostVpcId"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses []ResolverEndpointIpAddress `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name *string `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ResolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringPtrInput
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringPtrInput
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringPtrInput
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayInput
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringPtrInput
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ResolverEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointState)(nil)).Elem()
}

type resolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction string `pulumi:"direction"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses []ResolverEndpointIpAddress `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name *string `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringInput
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayInput
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringPtrInput
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ResolverEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointArgs)(nil)).Elem()
}

type ResolverEndpointInput interface {
	pulumi.Input

	ToResolverEndpointOutput() ResolverEndpointOutput
	ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput
}

func (*ResolverEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverEndpoint)(nil))
}

func (i *ResolverEndpoint) ToResolverEndpointOutput() ResolverEndpointOutput {
	return i.ToResolverEndpointOutputWithContext(context.Background())
}

func (i *ResolverEndpoint) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointOutput)
}

func (i *ResolverEndpoint) ToResolverEndpointPtrOutput() ResolverEndpointPtrOutput {
	return i.ToResolverEndpointPtrOutputWithContext(context.Background())
}

func (i *ResolverEndpoint) ToResolverEndpointPtrOutputWithContext(ctx context.Context) ResolverEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointPtrOutput)
}

type ResolverEndpointPtrInput interface {
	pulumi.Input

	ToResolverEndpointPtrOutput() ResolverEndpointPtrOutput
	ToResolverEndpointPtrOutputWithContext(ctx context.Context) ResolverEndpointPtrOutput
}

type resolverEndpointPtrType ResolverEndpointArgs

func (*resolverEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil))
}

func (i *resolverEndpointPtrType) ToResolverEndpointPtrOutput() ResolverEndpointPtrOutput {
	return i.ToResolverEndpointPtrOutputWithContext(context.Background())
}

func (i *resolverEndpointPtrType) ToResolverEndpointPtrOutputWithContext(ctx context.Context) ResolverEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointPtrOutput)
}

// ResolverEndpointArrayInput is an input type that accepts ResolverEndpointArray and ResolverEndpointArrayOutput values.
// You can construct a concrete instance of `ResolverEndpointArrayInput` via:
//
//          ResolverEndpointArray{ ResolverEndpointArgs{...} }
type ResolverEndpointArrayInput interface {
	pulumi.Input

	ToResolverEndpointArrayOutput() ResolverEndpointArrayOutput
	ToResolverEndpointArrayOutputWithContext(context.Context) ResolverEndpointArrayOutput
}

type ResolverEndpointArray []ResolverEndpointInput

func (ResolverEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ResolverEndpoint)(nil))
}

func (i ResolverEndpointArray) ToResolverEndpointArrayOutput() ResolverEndpointArrayOutput {
	return i.ToResolverEndpointArrayOutputWithContext(context.Background())
}

func (i ResolverEndpointArray) ToResolverEndpointArrayOutputWithContext(ctx context.Context) ResolverEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointArrayOutput)
}

// ResolverEndpointMapInput is an input type that accepts ResolverEndpointMap and ResolverEndpointMapOutput values.
// You can construct a concrete instance of `ResolverEndpointMapInput` via:
//
//          ResolverEndpointMap{ "key": ResolverEndpointArgs{...} }
type ResolverEndpointMapInput interface {
	pulumi.Input

	ToResolverEndpointMapOutput() ResolverEndpointMapOutput
	ToResolverEndpointMapOutputWithContext(context.Context) ResolverEndpointMapOutput
}

type ResolverEndpointMap map[string]ResolverEndpointInput

func (ResolverEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ResolverEndpoint)(nil))
}

func (i ResolverEndpointMap) ToResolverEndpointMapOutput() ResolverEndpointMapOutput {
	return i.ToResolverEndpointMapOutputWithContext(context.Background())
}

func (i ResolverEndpointMap) ToResolverEndpointMapOutputWithContext(ctx context.Context) ResolverEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverEndpointMapOutput)
}

type ResolverEndpointOutput struct {
	*pulumi.OutputState
}

func (ResolverEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverEndpoint)(nil))
}

func (o ResolverEndpointOutput) ToResolverEndpointOutput() ResolverEndpointOutput {
	return o
}

func (o ResolverEndpointOutput) ToResolverEndpointOutputWithContext(ctx context.Context) ResolverEndpointOutput {
	return o
}

func (o ResolverEndpointOutput) ToResolverEndpointPtrOutput() ResolverEndpointPtrOutput {
	return o.ToResolverEndpointPtrOutputWithContext(context.Background())
}

func (o ResolverEndpointOutput) ToResolverEndpointPtrOutputWithContext(ctx context.Context) ResolverEndpointPtrOutput {
	return o.ApplyT(func(v ResolverEndpoint) *ResolverEndpoint {
		return &v
	}).(ResolverEndpointPtrOutput)
}

type ResolverEndpointPtrOutput struct {
	*pulumi.OutputState
}

func (ResolverEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverEndpoint)(nil))
}

func (o ResolverEndpointPtrOutput) ToResolverEndpointPtrOutput() ResolverEndpointPtrOutput {
	return o
}

func (o ResolverEndpointPtrOutput) ToResolverEndpointPtrOutputWithContext(ctx context.Context) ResolverEndpointPtrOutput {
	return o
}

type ResolverEndpointArrayOutput struct{ *pulumi.OutputState }

func (ResolverEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResolverEndpoint)(nil))
}

func (o ResolverEndpointArrayOutput) ToResolverEndpointArrayOutput() ResolverEndpointArrayOutput {
	return o
}

func (o ResolverEndpointArrayOutput) ToResolverEndpointArrayOutputWithContext(ctx context.Context) ResolverEndpointArrayOutput {
	return o
}

func (o ResolverEndpointArrayOutput) Index(i pulumi.IntInput) ResolverEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResolverEndpoint {
		return vs[0].([]ResolverEndpoint)[vs[1].(int)]
	}).(ResolverEndpointOutput)
}

type ResolverEndpointMapOutput struct{ *pulumi.OutputState }

func (ResolverEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResolverEndpoint)(nil))
}

func (o ResolverEndpointMapOutput) ToResolverEndpointMapOutput() ResolverEndpointMapOutput {
	return o
}

func (o ResolverEndpointMapOutput) ToResolverEndpointMapOutputWithContext(ctx context.Context) ResolverEndpointMapOutput {
	return o
}

func (o ResolverEndpointMapOutput) MapIndex(k pulumi.StringInput) ResolverEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResolverEndpoint {
		return vs[0].(map[string]ResolverEndpoint)[vs[1].(string)]
	}).(ResolverEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(ResolverEndpointOutput{})
	pulumi.RegisterOutputType(ResolverEndpointPtrOutput{})
	pulumi.RegisterOutputType(ResolverEndpointArrayOutput{})
	pulumi.RegisterOutputType(ResolverEndpointMapOutput{})
}
