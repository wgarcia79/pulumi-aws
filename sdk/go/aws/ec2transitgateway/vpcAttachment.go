// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EC2 Transit Gateway VPC Attachment. For examples of custom route table association and propagation, see the EC2 Transit Gateway Networking Examples Guide.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.NewVpcAttachment(ctx, "example", &ec2transitgateway.VpcAttachmentArgs{
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example.Id),
// 			},
// 			TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
// 			VpcId:            pulumi.Any(aws_vpc.Example.Id),
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
// `aws_ec2_transit_gateway_vpc_attachment` can be imported by using the EC2 Transit Gateway Attachment identifier, e.g.,
//
// ```sh
//  $ pulumi import aws:ec2transitgateway/vpcAttachment:VpcAttachment example tgw-attach-12345678
// ```
type VpcAttachment struct {
	pulumi.CustomResourceState

	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport pulumi.StringPtrOutput `pulumi:"applianceModeSupport"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrOutput `pulumi:"dnsSupport"`
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support pulumi.StringPtrOutput `pulumi:"ipv6Support"`
	// Identifiers of EC2 Subnets.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	Tags      pulumi.StringMapOutput   `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrOutput `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrOutput `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
	// Identifier of EC2 VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Identifier of the AWS account that owns the EC2 VPC.
	VpcOwnerId pulumi.StringOutput `pulumi:"vpcOwnerId"`
}

// NewVpcAttachment registers a new resource with the given unique name, arguments, and options.
func NewVpcAttachment(ctx *pulumi.Context,
	name string, args *VpcAttachmentArgs, opts ...pulumi.ResourceOption) (*VpcAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VpcAttachment
	err := ctx.RegisterResource("aws:ec2transitgateway/vpcAttachment:VpcAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcAttachment gets an existing VpcAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcAttachmentState, opts ...pulumi.ResourceOption) (*VpcAttachment, error) {
	var resource VpcAttachment
	err := ctx.ReadResource("aws:ec2transitgateway/vpcAttachment:VpcAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcAttachment resources.
type vpcAttachmentState struct {
	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport *string `pulumi:"applianceModeSupport"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport *string `pulumi:"dnsSupport"`
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support *string `pulumi:"ipv6Support"`
	// Identifiers of EC2 Subnets.
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation *bool `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation *bool `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of EC2 VPC.
	VpcId *string `pulumi:"vpcId"`
	// Identifier of the AWS account that owns the EC2 VPC.
	VpcOwnerId *string `pulumi:"vpcOwnerId"`
}

type VpcAttachmentState struct {
	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport pulumi.StringPtrInput
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrInput
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support pulumi.StringPtrInput
	// Identifiers of EC2 Subnets.
	SubnetIds pulumi.StringArrayInput
	Tags      pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrInput
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// Identifier of EC2 VPC.
	VpcId pulumi.StringPtrInput
	// Identifier of the AWS account that owns the EC2 VPC.
	VpcOwnerId pulumi.StringPtrInput
}

func (VpcAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAttachmentState)(nil)).Elem()
}

type vpcAttachmentArgs struct {
	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport *string `pulumi:"applianceModeSupport"`
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport *string `pulumi:"dnsSupport"`
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support *string `pulumi:"ipv6Support"`
	// Identifiers of EC2 Subnets.
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation *bool `pulumi:"transitGatewayDefaultRouteTableAssociation"`
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation *bool `pulumi:"transitGatewayDefaultRouteTablePropagation"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
	// Identifier of EC2 VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcAttachment resource.
type VpcAttachmentArgs struct {
	// Whether Appliance Mode support is enabled. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. Valid values: `disable`, `enable`. Default value: `disable`.
	ApplianceModeSupport pulumi.StringPtrInput
	// Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
	DnsSupport pulumi.StringPtrInput
	// Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
	Ipv6Support pulumi.StringPtrInput
	// Identifiers of EC2 Subnets.
	SubnetIds pulumi.StringArrayInput
	Tags      pulumi.StringMapInput
	// Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTableAssociation pulumi.BoolPtrInput
	// Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
	TransitGatewayDefaultRouteTablePropagation pulumi.BoolPtrInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput
	// Identifier of EC2 VPC.
	VpcId pulumi.StringInput
}

func (VpcAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAttachmentArgs)(nil)).Elem()
}

type VpcAttachmentInput interface {
	pulumi.Input

	ToVpcAttachmentOutput() VpcAttachmentOutput
	ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput
}

func (*VpcAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcAttachment)(nil))
}

func (i *VpcAttachment) ToVpcAttachmentOutput() VpcAttachmentOutput {
	return i.ToVpcAttachmentOutputWithContext(context.Background())
}

func (i *VpcAttachment) ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentOutput)
}

func (i *VpcAttachment) ToVpcAttachmentPtrOutput() VpcAttachmentPtrOutput {
	return i.ToVpcAttachmentPtrOutputWithContext(context.Background())
}

func (i *VpcAttachment) ToVpcAttachmentPtrOutputWithContext(ctx context.Context) VpcAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentPtrOutput)
}

type VpcAttachmentPtrInput interface {
	pulumi.Input

	ToVpcAttachmentPtrOutput() VpcAttachmentPtrOutput
	ToVpcAttachmentPtrOutputWithContext(ctx context.Context) VpcAttachmentPtrOutput
}

type vpcAttachmentPtrType VpcAttachmentArgs

func (*vpcAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAttachment)(nil))
}

func (i *vpcAttachmentPtrType) ToVpcAttachmentPtrOutput() VpcAttachmentPtrOutput {
	return i.ToVpcAttachmentPtrOutputWithContext(context.Background())
}

func (i *vpcAttachmentPtrType) ToVpcAttachmentPtrOutputWithContext(ctx context.Context) VpcAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentPtrOutput)
}

// VpcAttachmentArrayInput is an input type that accepts VpcAttachmentArray and VpcAttachmentArrayOutput values.
// You can construct a concrete instance of `VpcAttachmentArrayInput` via:
//
//          VpcAttachmentArray{ VpcAttachmentArgs{...} }
type VpcAttachmentArrayInput interface {
	pulumi.Input

	ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput
	ToVpcAttachmentArrayOutputWithContext(context.Context) VpcAttachmentArrayOutput
}

type VpcAttachmentArray []VpcAttachmentInput

func (VpcAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcAttachment)(nil)).Elem()
}

func (i VpcAttachmentArray) ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput {
	return i.ToVpcAttachmentArrayOutputWithContext(context.Background())
}

func (i VpcAttachmentArray) ToVpcAttachmentArrayOutputWithContext(ctx context.Context) VpcAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentArrayOutput)
}

// VpcAttachmentMapInput is an input type that accepts VpcAttachmentMap and VpcAttachmentMapOutput values.
// You can construct a concrete instance of `VpcAttachmentMapInput` via:
//
//          VpcAttachmentMap{ "key": VpcAttachmentArgs{...} }
type VpcAttachmentMapInput interface {
	pulumi.Input

	ToVpcAttachmentMapOutput() VpcAttachmentMapOutput
	ToVpcAttachmentMapOutputWithContext(context.Context) VpcAttachmentMapOutput
}

type VpcAttachmentMap map[string]VpcAttachmentInput

func (VpcAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcAttachment)(nil)).Elem()
}

func (i VpcAttachmentMap) ToVpcAttachmentMapOutput() VpcAttachmentMapOutput {
	return i.ToVpcAttachmentMapOutputWithContext(context.Background())
}

func (i VpcAttachmentMap) ToVpcAttachmentMapOutputWithContext(ctx context.Context) VpcAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentMapOutput)
}

type VpcAttachmentOutput struct{ *pulumi.OutputState }

func (VpcAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcAttachment)(nil))
}

func (o VpcAttachmentOutput) ToVpcAttachmentOutput() VpcAttachmentOutput {
	return o
}

func (o VpcAttachmentOutput) ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput {
	return o
}

func (o VpcAttachmentOutput) ToVpcAttachmentPtrOutput() VpcAttachmentPtrOutput {
	return o.ToVpcAttachmentPtrOutputWithContext(context.Background())
}

func (o VpcAttachmentOutput) ToVpcAttachmentPtrOutputWithContext(ctx context.Context) VpcAttachmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpcAttachment) *VpcAttachment {
		return &v
	}).(VpcAttachmentPtrOutput)
}

type VpcAttachmentPtrOutput struct{ *pulumi.OutputState }

func (VpcAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAttachment)(nil))
}

func (o VpcAttachmentPtrOutput) ToVpcAttachmentPtrOutput() VpcAttachmentPtrOutput {
	return o
}

func (o VpcAttachmentPtrOutput) ToVpcAttachmentPtrOutputWithContext(ctx context.Context) VpcAttachmentPtrOutput {
	return o
}

func (o VpcAttachmentPtrOutput) Elem() VpcAttachmentOutput {
	return o.ApplyT(func(v *VpcAttachment) VpcAttachment {
		if v != nil {
			return *v
		}
		var ret VpcAttachment
		return ret
	}).(VpcAttachmentOutput)
}

type VpcAttachmentArrayOutput struct{ *pulumi.OutputState }

func (VpcAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcAttachment)(nil))
}

func (o VpcAttachmentArrayOutput) ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput {
	return o
}

func (o VpcAttachmentArrayOutput) ToVpcAttachmentArrayOutputWithContext(ctx context.Context) VpcAttachmentArrayOutput {
	return o
}

func (o VpcAttachmentArrayOutput) Index(i pulumi.IntInput) VpcAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcAttachment {
		return vs[0].([]VpcAttachment)[vs[1].(int)]
	}).(VpcAttachmentOutput)
}

type VpcAttachmentMapOutput struct{ *pulumi.OutputState }

func (VpcAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcAttachment)(nil))
}

func (o VpcAttachmentMapOutput) ToVpcAttachmentMapOutput() VpcAttachmentMapOutput {
	return o
}

func (o VpcAttachmentMapOutput) ToVpcAttachmentMapOutputWithContext(ctx context.Context) VpcAttachmentMapOutput {
	return o
}

func (o VpcAttachmentMapOutput) MapIndex(k pulumi.StringInput) VpcAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcAttachment {
		return vs[0].(map[string]VpcAttachment)[vs[1].(string)]
	}).(VpcAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentInput)(nil)).Elem(), &VpcAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentPtrInput)(nil)).Elem(), &VpcAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentArrayInput)(nil)).Elem(), VpcAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentMapInput)(nil)).Elem(), VpcAttachmentMap{})
	pulumi.RegisterOutputType(VpcAttachmentOutput{})
	pulumi.RegisterOutputType(VpcAttachmentPtrOutput{})
	pulumi.RegisterOutputType(VpcAttachmentArrayOutput{})
	pulumi.RegisterOutputType(VpcAttachmentMapOutput{})
}
