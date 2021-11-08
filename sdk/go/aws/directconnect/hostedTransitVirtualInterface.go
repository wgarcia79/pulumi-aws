// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Direct Connect hosted transit virtual interface resource.
// This resource represents the allocator's side of the hosted virtual interface.
// A hosted virtual interface is a virtual interface that is owned by another AWS account.
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
// 		_, err := directconnect.NewHostedTransitVirtualInterface(ctx, "example", &directconnect.HostedTransitVirtualInterfaceArgs{
// 			ConnectionId:  pulumi.Any(aws_dx_connection.Example.Id),
// 			Vlan:          pulumi.Int(4094),
// 			AddressFamily: pulumi.String("ipv4"),
// 			BgpAsn:        pulumi.Int(65352),
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
// Direct Connect hosted transit virtual interfaces can be imported using the `vif id`, e.g.,
//
// ```sh
//  $ pulumi import aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface test dxvif-33cc44dd
// ```
type HostedTransitVirtualInterface struct {
	pulumi.CustomResourceState

	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringOutput `pulumi:"amazonAddress"`
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringOutput `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringOutput `pulumi:"customerAddress"`
	// Indicates whether jumbo frames (8500 MTU) are supported.
	JumboFrameCapable pulumi.BoolOutput `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewHostedTransitVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewHostedTransitVirtualInterface(ctx *pulumi.Context,
	name string, args *HostedTransitVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.BgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'BgpAsn'")
	}
	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.OwnerAccountId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerAccountId'")
	}
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	var resource HostedTransitVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedTransitVirtualInterface gets an existing HostedTransitVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedTransitVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedTransitVirtualInterfaceState, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterface, error) {
	var resource HostedTransitVirtualInterface
	err := ctx.ReadResource("aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedTransitVirtualInterface resources.
type hostedTransitVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily *string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn *int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId *string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// Indicates whether jumbo frames (8500 MTU) are supported.
	JumboFrameCapable *bool `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan *int `pulumi:"vlan"`
}

type HostedTransitVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringPtrInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	AmazonSideAsn pulumi.StringPtrInput
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntPtrInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringPtrInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// Indicates whether jumbo frames (8500 MTU) are supported.
	JumboFrameCapable pulumi.BoolPtrInput
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringPtrInput
	// The VLAN ID.
	Vlan pulumi.IntPtrInput
}

func (HostedTransitVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceState)(nil)).Elem()
}

type hostedTransitVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId string `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan int `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedTransitVirtualInterface resource.
type HostedTransitVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringInput
	// The VLAN ID.
	Vlan pulumi.IntInput
}

func (HostedTransitVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceArgs)(nil)).Elem()
}

type HostedTransitVirtualInterfaceInput interface {
	pulumi.Input

	ToHostedTransitVirtualInterfaceOutput() HostedTransitVirtualInterfaceOutput
	ToHostedTransitVirtualInterfaceOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceOutput
}

func (*HostedTransitVirtualInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedTransitVirtualInterface)(nil))
}

func (i *HostedTransitVirtualInterface) ToHostedTransitVirtualInterfaceOutput() HostedTransitVirtualInterfaceOutput {
	return i.ToHostedTransitVirtualInterfaceOutputWithContext(context.Background())
}

func (i *HostedTransitVirtualInterface) ToHostedTransitVirtualInterfaceOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedTransitVirtualInterfaceOutput)
}

func (i *HostedTransitVirtualInterface) ToHostedTransitVirtualInterfacePtrOutput() HostedTransitVirtualInterfacePtrOutput {
	return i.ToHostedTransitVirtualInterfacePtrOutputWithContext(context.Background())
}

func (i *HostedTransitVirtualInterface) ToHostedTransitVirtualInterfacePtrOutputWithContext(ctx context.Context) HostedTransitVirtualInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedTransitVirtualInterfacePtrOutput)
}

type HostedTransitVirtualInterfacePtrInput interface {
	pulumi.Input

	ToHostedTransitVirtualInterfacePtrOutput() HostedTransitVirtualInterfacePtrOutput
	ToHostedTransitVirtualInterfacePtrOutputWithContext(ctx context.Context) HostedTransitVirtualInterfacePtrOutput
}

type hostedTransitVirtualInterfacePtrType HostedTransitVirtualInterfaceArgs

func (*hostedTransitVirtualInterfacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedTransitVirtualInterface)(nil))
}

func (i *hostedTransitVirtualInterfacePtrType) ToHostedTransitVirtualInterfacePtrOutput() HostedTransitVirtualInterfacePtrOutput {
	return i.ToHostedTransitVirtualInterfacePtrOutputWithContext(context.Background())
}

func (i *hostedTransitVirtualInterfacePtrType) ToHostedTransitVirtualInterfacePtrOutputWithContext(ctx context.Context) HostedTransitVirtualInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedTransitVirtualInterfacePtrOutput)
}

// HostedTransitVirtualInterfaceArrayInput is an input type that accepts HostedTransitVirtualInterfaceArray and HostedTransitVirtualInterfaceArrayOutput values.
// You can construct a concrete instance of `HostedTransitVirtualInterfaceArrayInput` via:
//
//          HostedTransitVirtualInterfaceArray{ HostedTransitVirtualInterfaceArgs{...} }
type HostedTransitVirtualInterfaceArrayInput interface {
	pulumi.Input

	ToHostedTransitVirtualInterfaceArrayOutput() HostedTransitVirtualInterfaceArrayOutput
	ToHostedTransitVirtualInterfaceArrayOutputWithContext(context.Context) HostedTransitVirtualInterfaceArrayOutput
}

type HostedTransitVirtualInterfaceArray []HostedTransitVirtualInterfaceInput

func (HostedTransitVirtualInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedTransitVirtualInterface)(nil)).Elem()
}

func (i HostedTransitVirtualInterfaceArray) ToHostedTransitVirtualInterfaceArrayOutput() HostedTransitVirtualInterfaceArrayOutput {
	return i.ToHostedTransitVirtualInterfaceArrayOutputWithContext(context.Background())
}

func (i HostedTransitVirtualInterfaceArray) ToHostedTransitVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedTransitVirtualInterfaceArrayOutput)
}

// HostedTransitVirtualInterfaceMapInput is an input type that accepts HostedTransitVirtualInterfaceMap and HostedTransitVirtualInterfaceMapOutput values.
// You can construct a concrete instance of `HostedTransitVirtualInterfaceMapInput` via:
//
//          HostedTransitVirtualInterfaceMap{ "key": HostedTransitVirtualInterfaceArgs{...} }
type HostedTransitVirtualInterfaceMapInput interface {
	pulumi.Input

	ToHostedTransitVirtualInterfaceMapOutput() HostedTransitVirtualInterfaceMapOutput
	ToHostedTransitVirtualInterfaceMapOutputWithContext(context.Context) HostedTransitVirtualInterfaceMapOutput
}

type HostedTransitVirtualInterfaceMap map[string]HostedTransitVirtualInterfaceInput

func (HostedTransitVirtualInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedTransitVirtualInterface)(nil)).Elem()
}

func (i HostedTransitVirtualInterfaceMap) ToHostedTransitVirtualInterfaceMapOutput() HostedTransitVirtualInterfaceMapOutput {
	return i.ToHostedTransitVirtualInterfaceMapOutputWithContext(context.Background())
}

func (i HostedTransitVirtualInterfaceMap) ToHostedTransitVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedTransitVirtualInterfaceMapOutput)
}

type HostedTransitVirtualInterfaceOutput struct{ *pulumi.OutputState }

func (HostedTransitVirtualInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostedTransitVirtualInterface)(nil))
}

func (o HostedTransitVirtualInterfaceOutput) ToHostedTransitVirtualInterfaceOutput() HostedTransitVirtualInterfaceOutput {
	return o
}

func (o HostedTransitVirtualInterfaceOutput) ToHostedTransitVirtualInterfaceOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceOutput {
	return o
}

func (o HostedTransitVirtualInterfaceOutput) ToHostedTransitVirtualInterfacePtrOutput() HostedTransitVirtualInterfacePtrOutput {
	return o.ToHostedTransitVirtualInterfacePtrOutputWithContext(context.Background())
}

func (o HostedTransitVirtualInterfaceOutput) ToHostedTransitVirtualInterfacePtrOutputWithContext(ctx context.Context) HostedTransitVirtualInterfacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostedTransitVirtualInterface) *HostedTransitVirtualInterface {
		return &v
	}).(HostedTransitVirtualInterfacePtrOutput)
}

type HostedTransitVirtualInterfacePtrOutput struct{ *pulumi.OutputState }

func (HostedTransitVirtualInterfacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedTransitVirtualInterface)(nil))
}

func (o HostedTransitVirtualInterfacePtrOutput) ToHostedTransitVirtualInterfacePtrOutput() HostedTransitVirtualInterfacePtrOutput {
	return o
}

func (o HostedTransitVirtualInterfacePtrOutput) ToHostedTransitVirtualInterfacePtrOutputWithContext(ctx context.Context) HostedTransitVirtualInterfacePtrOutput {
	return o
}

func (o HostedTransitVirtualInterfacePtrOutput) Elem() HostedTransitVirtualInterfaceOutput {
	return o.ApplyT(func(v *HostedTransitVirtualInterface) HostedTransitVirtualInterface {
		if v != nil {
			return *v
		}
		var ret HostedTransitVirtualInterface
		return ret
	}).(HostedTransitVirtualInterfaceOutput)
}

type HostedTransitVirtualInterfaceArrayOutput struct{ *pulumi.OutputState }

func (HostedTransitVirtualInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostedTransitVirtualInterface)(nil))
}

func (o HostedTransitVirtualInterfaceArrayOutput) ToHostedTransitVirtualInterfaceArrayOutput() HostedTransitVirtualInterfaceArrayOutput {
	return o
}

func (o HostedTransitVirtualInterfaceArrayOutput) ToHostedTransitVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceArrayOutput {
	return o
}

func (o HostedTransitVirtualInterfaceArrayOutput) Index(i pulumi.IntInput) HostedTransitVirtualInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostedTransitVirtualInterface {
		return vs[0].([]HostedTransitVirtualInterface)[vs[1].(int)]
	}).(HostedTransitVirtualInterfaceOutput)
}

type HostedTransitVirtualInterfaceMapOutput struct{ *pulumi.OutputState }

func (HostedTransitVirtualInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostedTransitVirtualInterface)(nil))
}

func (o HostedTransitVirtualInterfaceMapOutput) ToHostedTransitVirtualInterfaceMapOutput() HostedTransitVirtualInterfaceMapOutput {
	return o
}

func (o HostedTransitVirtualInterfaceMapOutput) ToHostedTransitVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedTransitVirtualInterfaceMapOutput {
	return o
}

func (o HostedTransitVirtualInterfaceMapOutput) MapIndex(k pulumi.StringInput) HostedTransitVirtualInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostedTransitVirtualInterface {
		return vs[0].(map[string]HostedTransitVirtualInterface)[vs[1].(string)]
	}).(HostedTransitVirtualInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostedTransitVirtualInterfaceInput)(nil)).Elem(), &HostedTransitVirtualInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedTransitVirtualInterfacePtrInput)(nil)).Elem(), &HostedTransitVirtualInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedTransitVirtualInterfaceArrayInput)(nil)).Elem(), HostedTransitVirtualInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedTransitVirtualInterfaceMapInput)(nil)).Elem(), HostedTransitVirtualInterfaceMap{})
	pulumi.RegisterOutputType(HostedTransitVirtualInterfaceOutput{})
	pulumi.RegisterOutputType(HostedTransitVirtualInterfacePtrOutput{})
	pulumi.RegisterOutputType(HostedTransitVirtualInterfaceArrayOutput{})
	pulumi.RegisterOutputType(HostedTransitVirtualInterfaceMapOutput{})
}
