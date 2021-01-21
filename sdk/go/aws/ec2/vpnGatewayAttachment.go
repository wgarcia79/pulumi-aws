// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Virtual Private Gateway attachment resource, allowing for an existing
// hardware VPN gateway to be attached and/or detached from a VPC.
//
// > **Note:** The `ec2.VpnGateway`
// resource can also automatically attach the Virtual Private Gateway it creates
// to an existing VPC by setting the `vpcId` attribute accordingly.
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
// 		network, err := ec2.NewVpc(ctx, "network", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vpn, err := ec2.NewVpnGateway(ctx, "vpn", &ec2.VpnGatewayArgs{
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("example-vpn-gateway"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewVpnGatewayAttachment(ctx, "vpnAttachment", &ec2.VpnGatewayAttachmentArgs{
// 			VpcId:        network.ID(),
// 			VpnGatewayId: vpn.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// See [Virtual Private Cloud](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Introduction.html)
// and [Virtual Private Gateway](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_VPN.html) user
// guides for more information.
//
// ## Import
//
// This resource does not support importing.
type VpnGatewayAttachment struct {
	pulumi.CustomResourceState

	// The ID of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewVpnGatewayAttachment registers a new resource with the given unique name, arguments, and options.
func NewVpnGatewayAttachment(ctx *pulumi.Context,
	name string, args *VpnGatewayAttachmentArgs, opts ...pulumi.ResourceOption) (*VpnGatewayAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.VpnGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'VpnGatewayId'")
	}
	var resource VpnGatewayAttachment
	err := ctx.RegisterResource("aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGatewayAttachment gets an existing VpnGatewayAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGatewayAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayAttachmentState, opts ...pulumi.ResourceOption) (*VpnGatewayAttachment, error) {
	var resource VpnGatewayAttachment
	err := ctx.ReadResource("aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGatewayAttachment resources.
type vpnGatewayAttachmentState struct {
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// The ID of the Virtual Private Gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type VpnGatewayAttachmentState struct {
	// The ID of the VPC.
	VpcId pulumi.StringPtrInput
	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringPtrInput
}

func (VpnGatewayAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayAttachmentState)(nil)).Elem()
}

type vpnGatewayAttachmentArgs struct {
	// The ID of the VPC.
	VpcId string `pulumi:"vpcId"`
	// The ID of the Virtual Private Gateway.
	VpnGatewayId string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a VpnGatewayAttachment resource.
type VpnGatewayAttachmentArgs struct {
	// The ID of the VPC.
	VpcId pulumi.StringInput
	// The ID of the Virtual Private Gateway.
	VpnGatewayId pulumi.StringInput
}

func (VpnGatewayAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayAttachmentArgs)(nil)).Elem()
}

type VpnGatewayAttachmentInput interface {
	pulumi.Input

	ToVpnGatewayAttachmentOutput() VpnGatewayAttachmentOutput
	ToVpnGatewayAttachmentOutputWithContext(ctx context.Context) VpnGatewayAttachmentOutput
}

func (*VpnGatewayAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayAttachment)(nil))
}

func (i *VpnGatewayAttachment) ToVpnGatewayAttachmentOutput() VpnGatewayAttachmentOutput {
	return i.ToVpnGatewayAttachmentOutputWithContext(context.Background())
}

func (i *VpnGatewayAttachment) ToVpnGatewayAttachmentOutputWithContext(ctx context.Context) VpnGatewayAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayAttachmentOutput)
}

func (i *VpnGatewayAttachment) ToVpnGatewayAttachmentPtrOutput() VpnGatewayAttachmentPtrOutput {
	return i.ToVpnGatewayAttachmentPtrOutputWithContext(context.Background())
}

func (i *VpnGatewayAttachment) ToVpnGatewayAttachmentPtrOutputWithContext(ctx context.Context) VpnGatewayAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayAttachmentPtrOutput)
}

type VpnGatewayAttachmentPtrInput interface {
	pulumi.Input

	ToVpnGatewayAttachmentPtrOutput() VpnGatewayAttachmentPtrOutput
	ToVpnGatewayAttachmentPtrOutputWithContext(ctx context.Context) VpnGatewayAttachmentPtrOutput
}

type vpnGatewayAttachmentPtrType VpnGatewayAttachmentArgs

func (*vpnGatewayAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGatewayAttachment)(nil))
}

func (i *vpnGatewayAttachmentPtrType) ToVpnGatewayAttachmentPtrOutput() VpnGatewayAttachmentPtrOutput {
	return i.ToVpnGatewayAttachmentPtrOutputWithContext(context.Background())
}

func (i *vpnGatewayAttachmentPtrType) ToVpnGatewayAttachmentPtrOutputWithContext(ctx context.Context) VpnGatewayAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayAttachmentPtrOutput)
}

// VpnGatewayAttachmentArrayInput is an input type that accepts VpnGatewayAttachmentArray and VpnGatewayAttachmentArrayOutput values.
// You can construct a concrete instance of `VpnGatewayAttachmentArrayInput` via:
//
//          VpnGatewayAttachmentArray{ VpnGatewayAttachmentArgs{...} }
type VpnGatewayAttachmentArrayInput interface {
	pulumi.Input

	ToVpnGatewayAttachmentArrayOutput() VpnGatewayAttachmentArrayOutput
	ToVpnGatewayAttachmentArrayOutputWithContext(context.Context) VpnGatewayAttachmentArrayOutput
}

type VpnGatewayAttachmentArray []VpnGatewayAttachmentInput

func (VpnGatewayAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpnGatewayAttachment)(nil))
}

func (i VpnGatewayAttachmentArray) ToVpnGatewayAttachmentArrayOutput() VpnGatewayAttachmentArrayOutput {
	return i.ToVpnGatewayAttachmentArrayOutputWithContext(context.Background())
}

func (i VpnGatewayAttachmentArray) ToVpnGatewayAttachmentArrayOutputWithContext(ctx context.Context) VpnGatewayAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayAttachmentArrayOutput)
}

// VpnGatewayAttachmentMapInput is an input type that accepts VpnGatewayAttachmentMap and VpnGatewayAttachmentMapOutput values.
// You can construct a concrete instance of `VpnGatewayAttachmentMapInput` via:
//
//          VpnGatewayAttachmentMap{ "key": VpnGatewayAttachmentArgs{...} }
type VpnGatewayAttachmentMapInput interface {
	pulumi.Input

	ToVpnGatewayAttachmentMapOutput() VpnGatewayAttachmentMapOutput
	ToVpnGatewayAttachmentMapOutputWithContext(context.Context) VpnGatewayAttachmentMapOutput
}

type VpnGatewayAttachmentMap map[string]VpnGatewayAttachmentInput

func (VpnGatewayAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpnGatewayAttachment)(nil))
}

func (i VpnGatewayAttachmentMap) ToVpnGatewayAttachmentMapOutput() VpnGatewayAttachmentMapOutput {
	return i.ToVpnGatewayAttachmentMapOutputWithContext(context.Background())
}

func (i VpnGatewayAttachmentMap) ToVpnGatewayAttachmentMapOutputWithContext(ctx context.Context) VpnGatewayAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayAttachmentMapOutput)
}

type VpnGatewayAttachmentOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayAttachment)(nil))
}

func (o VpnGatewayAttachmentOutput) ToVpnGatewayAttachmentOutput() VpnGatewayAttachmentOutput {
	return o
}

func (o VpnGatewayAttachmentOutput) ToVpnGatewayAttachmentOutputWithContext(ctx context.Context) VpnGatewayAttachmentOutput {
	return o
}

func (o VpnGatewayAttachmentOutput) ToVpnGatewayAttachmentPtrOutput() VpnGatewayAttachmentPtrOutput {
	return o.ToVpnGatewayAttachmentPtrOutputWithContext(context.Background())
}

func (o VpnGatewayAttachmentOutput) ToVpnGatewayAttachmentPtrOutputWithContext(ctx context.Context) VpnGatewayAttachmentPtrOutput {
	return o.ApplyT(func(v VpnGatewayAttachment) *VpnGatewayAttachment {
		return &v
	}).(VpnGatewayAttachmentPtrOutput)
}

type VpnGatewayAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGatewayAttachment)(nil))
}

func (o VpnGatewayAttachmentPtrOutput) ToVpnGatewayAttachmentPtrOutput() VpnGatewayAttachmentPtrOutput {
	return o
}

func (o VpnGatewayAttachmentPtrOutput) ToVpnGatewayAttachmentPtrOutputWithContext(ctx context.Context) VpnGatewayAttachmentPtrOutput {
	return o
}

type VpnGatewayAttachmentArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayAttachment)(nil))
}

func (o VpnGatewayAttachmentArrayOutput) ToVpnGatewayAttachmentArrayOutput() VpnGatewayAttachmentArrayOutput {
	return o
}

func (o VpnGatewayAttachmentArrayOutput) ToVpnGatewayAttachmentArrayOutputWithContext(ctx context.Context) VpnGatewayAttachmentArrayOutput {
	return o
}

func (o VpnGatewayAttachmentArrayOutput) Index(i pulumi.IntInput) VpnGatewayAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayAttachment {
		return vs[0].([]VpnGatewayAttachment)[vs[1].(int)]
	}).(VpnGatewayAttachmentOutput)
}

type VpnGatewayAttachmentMapOutput struct{ *pulumi.OutputState }

func (VpnGatewayAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnGatewayAttachment)(nil))
}

func (o VpnGatewayAttachmentMapOutput) ToVpnGatewayAttachmentMapOutput() VpnGatewayAttachmentMapOutput {
	return o
}

func (o VpnGatewayAttachmentMapOutput) ToVpnGatewayAttachmentMapOutputWithContext(ctx context.Context) VpnGatewayAttachmentMapOutput {
	return o
}

func (o VpnGatewayAttachmentMapOutput) MapIndex(k pulumi.StringInput) VpnGatewayAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpnGatewayAttachment {
		return vs[0].(map[string]VpnGatewayAttachment)[vs[1].(string)]
	}).(VpnGatewayAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayAttachmentOutput{})
	pulumi.RegisterOutputType(VpnGatewayAttachmentPtrOutput{})
	pulumi.RegisterOutputType(VpnGatewayAttachmentArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayAttachmentMapOutput{})
}
