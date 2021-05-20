// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a VPC NAT Gateway.
//
// ## Example Usage
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
// 		_, err := ec2.NewNatGateway(ctx, "gw", &ec2.NatGatewayArgs{
// 			AllocationId: pulumi.Any(aws_eip.Nat.Id),
// 			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Usage with tags:
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
// 		_, err := ec2.NewNatGateway(ctx, "gw", &ec2.NatGatewayArgs{
// 			AllocationId: pulumi.Any(aws_eip.Nat.Id),
// 			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("gw NAT"),
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
// NAT Gateways can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/natGateway:NatGateway private_gw nat-05dba92075d71c408
// ```
type NatGateway struct {
	pulumi.CustomResourceState

	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The private IP address of the NAT Gateway.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The public IP address of the NAT Gateway.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationId == nil {
		return nil, errors.New("invalid value for required argument 'AllocationId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource NatGateway
	err := ctx.RegisterResource("aws:ec2/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("aws:ec2/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId *string `pulumi:"allocationId"`
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The private IP address of the NAT Gateway.
	PrivateIp *string `pulumi:"privateIp"`
	// The public IP address of the NAT Gateway.
	PublicIp *string `pulumi:"publicIp"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type NatGatewayState struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringPtrInput
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId pulumi.StringPtrInput
	// The private IP address of the NAT Gateway.
	PrivateIp pulumi.StringPtrInput
	// The public IP address of the NAT Gateway.
	PublicIp pulumi.StringPtrInput
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId string `pulumi:"allocationId"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringInput
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGateway)(nil))
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

func (i *NatGateway) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return i.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPtrOutput)
}

type NatGatewayPtrInput interface {
	pulumi.Input

	ToNatGatewayPtrOutput() NatGatewayPtrOutput
	ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput
}

type natGatewayPtrType NatGatewayArgs

func (*natGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil))
}

func (i *natGatewayPtrType) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return i.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (i *natGatewayPtrType) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayPtrOutput)
}

// NatGatewayArrayInput is an input type that accepts NatGatewayArray and NatGatewayArrayOutput values.
// You can construct a concrete instance of `NatGatewayArrayInput` via:
//
//          NatGatewayArray{ NatGatewayArgs{...} }
type NatGatewayArrayInput interface {
	pulumi.Input

	ToNatGatewayArrayOutput() NatGatewayArrayOutput
	ToNatGatewayArrayOutputWithContext(context.Context) NatGatewayArrayOutput
}

type NatGatewayArray []NatGatewayInput

func (NatGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NatGateway)(nil))
}

func (i NatGatewayArray) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return i.ToNatGatewayArrayOutputWithContext(context.Background())
}

func (i NatGatewayArray) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayArrayOutput)
}

// NatGatewayMapInput is an input type that accepts NatGatewayMap and NatGatewayMapOutput values.
// You can construct a concrete instance of `NatGatewayMapInput` via:
//
//          NatGatewayMap{ "key": NatGatewayArgs{...} }
type NatGatewayMapInput interface {
	pulumi.Input

	ToNatGatewayMapOutput() NatGatewayMapOutput
	ToNatGatewayMapOutputWithContext(context.Context) NatGatewayMapOutput
}

type NatGatewayMap map[string]NatGatewayInput

func (NatGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NatGateway)(nil))
}

func (i NatGatewayMap) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return i.ToNatGatewayMapOutputWithContext(context.Background())
}

func (i NatGatewayMap) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayMapOutput)
}

type NatGatewayOutput struct {
	*pulumi.OutputState
}

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGateway)(nil))
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return o.ToNatGatewayPtrOutputWithContext(context.Background())
}

func (o NatGatewayOutput) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return o.ApplyT(func(v NatGateway) *NatGateway {
		return &v
	}).(NatGatewayPtrOutput)
}

type NatGatewayPtrOutput struct {
	*pulumi.OutputState
}

func (NatGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil))
}

func (o NatGatewayPtrOutput) ToNatGatewayPtrOutput() NatGatewayPtrOutput {
	return o
}

func (o NatGatewayPtrOutput) ToNatGatewayPtrOutputWithContext(ctx context.Context) NatGatewayPtrOutput {
	return o
}

type NatGatewayArrayOutput struct{ *pulumi.OutputState }

func (NatGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NatGateway)(nil))
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) Index(i pulumi.IntInput) NatGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NatGateway {
		return vs[0].([]NatGateway)[vs[1].(int)]
	}).(NatGatewayOutput)
}

type NatGatewayMapOutput struct{ *pulumi.OutputState }

func (NatGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NatGateway)(nil))
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) MapIndex(k pulumi.StringInput) NatGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NatGateway {
		return vs[0].(map[string]NatGateway)[vs[1].(string)]
	}).(NatGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(NatGatewayOutput{})
	pulumi.RegisterOutputType(NatGatewayPtrOutput{})
	pulumi.RegisterOutputType(NatGatewayArrayOutput{})
	pulumi.RegisterOutputType(NatGatewayMapOutput{})
}
