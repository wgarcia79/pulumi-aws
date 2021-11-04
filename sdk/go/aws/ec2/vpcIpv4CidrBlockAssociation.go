// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to associate additional IPv4 CIDR blocks with a VPC.
//
// When a VPC is created, a primary IPv4 CIDR block for the VPC must be specified.
// The `ec2.VpcIpv4CidrBlockAssociation` resource allows further IPv4 CIDR blocks to be added to the VPC.
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
// 		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewVpcIpv4CidrBlockAssociation(ctx, "secondaryCidr", &ec2.VpcIpv4CidrBlockAssociationArgs{
// 			VpcId:     main.ID(),
// 			CidrBlock: pulumi.String("172.2.0.0/16"),
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
// `aws_vpc_ipv4_cidr_block_association` can be imported by using the VPC CIDR Association ID, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation example vpc-cidr-assoc-xxxxxxxx
// ```
type VpcIpv4CidrBlockAssociation struct {
	pulumi.CustomResourceState

	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcIpv4CidrBlockAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcIpv4CidrBlockAssociation(ctx *pulumi.Context,
	name string, args *VpcIpv4CidrBlockAssociationArgs, opts ...pulumi.ResourceOption) (*VpcIpv4CidrBlockAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VpcIpv4CidrBlockAssociation
	err := ctx.RegisterResource("aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpv4CidrBlockAssociation gets an existing VpcIpv4CidrBlockAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpv4CidrBlockAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpv4CidrBlockAssociationState, opts ...pulumi.ResourceOption) (*VpcIpv4CidrBlockAssociation, error) {
	var resource VpcIpv4CidrBlockAssociation
	err := ctx.ReadResource("aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpv4CidrBlockAssociation resources.
type vpcIpv4CidrBlockAssociationState struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId *string `pulumi:"vpcId"`
}

type VpcIpv4CidrBlockAssociationState struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringPtrInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringPtrInput
}

func (VpcIpv4CidrBlockAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv4CidrBlockAssociationState)(nil)).Elem()
}

type vpcIpv4CidrBlockAssociationArgs struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock string `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcIpv4CidrBlockAssociation resource.
type VpcIpv4CidrBlockAssociationArgs struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringInput
}

func (VpcIpv4CidrBlockAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv4CidrBlockAssociationArgs)(nil)).Elem()
}

type VpcIpv4CidrBlockAssociationInput interface {
	pulumi.Input

	ToVpcIpv4CidrBlockAssociationOutput() VpcIpv4CidrBlockAssociationOutput
	ToVpcIpv4CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationOutput
}

func (*VpcIpv4CidrBlockAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcIpv4CidrBlockAssociation)(nil))
}

func (i *VpcIpv4CidrBlockAssociation) ToVpcIpv4CidrBlockAssociationOutput() VpcIpv4CidrBlockAssociationOutput {
	return i.ToVpcIpv4CidrBlockAssociationOutputWithContext(context.Background())
}

func (i *VpcIpv4CidrBlockAssociation) ToVpcIpv4CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv4CidrBlockAssociationOutput)
}

func (i *VpcIpv4CidrBlockAssociation) ToVpcIpv4CidrBlockAssociationPtrOutput() VpcIpv4CidrBlockAssociationPtrOutput {
	return i.ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(context.Background())
}

func (i *VpcIpv4CidrBlockAssociation) ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv4CidrBlockAssociationPtrOutput)
}

type VpcIpv4CidrBlockAssociationPtrInput interface {
	pulumi.Input

	ToVpcIpv4CidrBlockAssociationPtrOutput() VpcIpv4CidrBlockAssociationPtrOutput
	ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationPtrOutput
}

type vpcIpv4CidrBlockAssociationPtrType VpcIpv4CidrBlockAssociationArgs

func (*vpcIpv4CidrBlockAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpv4CidrBlockAssociation)(nil))
}

func (i *vpcIpv4CidrBlockAssociationPtrType) ToVpcIpv4CidrBlockAssociationPtrOutput() VpcIpv4CidrBlockAssociationPtrOutput {
	return i.ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(context.Background())
}

func (i *vpcIpv4CidrBlockAssociationPtrType) ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv4CidrBlockAssociationPtrOutput)
}

// VpcIpv4CidrBlockAssociationArrayInput is an input type that accepts VpcIpv4CidrBlockAssociationArray and VpcIpv4CidrBlockAssociationArrayOutput values.
// You can construct a concrete instance of `VpcIpv4CidrBlockAssociationArrayInput` via:
//
//          VpcIpv4CidrBlockAssociationArray{ VpcIpv4CidrBlockAssociationArgs{...} }
type VpcIpv4CidrBlockAssociationArrayInput interface {
	pulumi.Input

	ToVpcIpv4CidrBlockAssociationArrayOutput() VpcIpv4CidrBlockAssociationArrayOutput
	ToVpcIpv4CidrBlockAssociationArrayOutputWithContext(context.Context) VpcIpv4CidrBlockAssociationArrayOutput
}

type VpcIpv4CidrBlockAssociationArray []VpcIpv4CidrBlockAssociationInput

func (VpcIpv4CidrBlockAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpv4CidrBlockAssociation)(nil)).Elem()
}

func (i VpcIpv4CidrBlockAssociationArray) ToVpcIpv4CidrBlockAssociationArrayOutput() VpcIpv4CidrBlockAssociationArrayOutput {
	return i.ToVpcIpv4CidrBlockAssociationArrayOutputWithContext(context.Background())
}

func (i VpcIpv4CidrBlockAssociationArray) ToVpcIpv4CidrBlockAssociationArrayOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv4CidrBlockAssociationArrayOutput)
}

// VpcIpv4CidrBlockAssociationMapInput is an input type that accepts VpcIpv4CidrBlockAssociationMap and VpcIpv4CidrBlockAssociationMapOutput values.
// You can construct a concrete instance of `VpcIpv4CidrBlockAssociationMapInput` via:
//
//          VpcIpv4CidrBlockAssociationMap{ "key": VpcIpv4CidrBlockAssociationArgs{...} }
type VpcIpv4CidrBlockAssociationMapInput interface {
	pulumi.Input

	ToVpcIpv4CidrBlockAssociationMapOutput() VpcIpv4CidrBlockAssociationMapOutput
	ToVpcIpv4CidrBlockAssociationMapOutputWithContext(context.Context) VpcIpv4CidrBlockAssociationMapOutput
}

type VpcIpv4CidrBlockAssociationMap map[string]VpcIpv4CidrBlockAssociationInput

func (VpcIpv4CidrBlockAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpv4CidrBlockAssociation)(nil)).Elem()
}

func (i VpcIpv4CidrBlockAssociationMap) ToVpcIpv4CidrBlockAssociationMapOutput() VpcIpv4CidrBlockAssociationMapOutput {
	return i.ToVpcIpv4CidrBlockAssociationMapOutputWithContext(context.Background())
}

func (i VpcIpv4CidrBlockAssociationMap) ToVpcIpv4CidrBlockAssociationMapOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpv4CidrBlockAssociationMapOutput)
}

type VpcIpv4CidrBlockAssociationOutput struct{ *pulumi.OutputState }

func (VpcIpv4CidrBlockAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcIpv4CidrBlockAssociation)(nil))
}

func (o VpcIpv4CidrBlockAssociationOutput) ToVpcIpv4CidrBlockAssociationOutput() VpcIpv4CidrBlockAssociationOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationOutput) ToVpcIpv4CidrBlockAssociationOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationOutput) ToVpcIpv4CidrBlockAssociationPtrOutput() VpcIpv4CidrBlockAssociationPtrOutput {
	return o.ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(context.Background())
}

func (o VpcIpv4CidrBlockAssociationOutput) ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpcIpv4CidrBlockAssociation) *VpcIpv4CidrBlockAssociation {
		return &v
	}).(VpcIpv4CidrBlockAssociationPtrOutput)
}

type VpcIpv4CidrBlockAssociationPtrOutput struct{ *pulumi.OutputState }

func (VpcIpv4CidrBlockAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpv4CidrBlockAssociation)(nil))
}

func (o VpcIpv4CidrBlockAssociationPtrOutput) ToVpcIpv4CidrBlockAssociationPtrOutput() VpcIpv4CidrBlockAssociationPtrOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationPtrOutput) ToVpcIpv4CidrBlockAssociationPtrOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationPtrOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationPtrOutput) Elem() VpcIpv4CidrBlockAssociationOutput {
	return o.ApplyT(func(v *VpcIpv4CidrBlockAssociation) VpcIpv4CidrBlockAssociation {
		if v != nil {
			return *v
		}
		var ret VpcIpv4CidrBlockAssociation
		return ret
	}).(VpcIpv4CidrBlockAssociationOutput)
}

type VpcIpv4CidrBlockAssociationArrayOutput struct{ *pulumi.OutputState }

func (VpcIpv4CidrBlockAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcIpv4CidrBlockAssociation)(nil))
}

func (o VpcIpv4CidrBlockAssociationArrayOutput) ToVpcIpv4CidrBlockAssociationArrayOutput() VpcIpv4CidrBlockAssociationArrayOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationArrayOutput) ToVpcIpv4CidrBlockAssociationArrayOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationArrayOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationArrayOutput) Index(i pulumi.IntInput) VpcIpv4CidrBlockAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcIpv4CidrBlockAssociation {
		return vs[0].([]VpcIpv4CidrBlockAssociation)[vs[1].(int)]
	}).(VpcIpv4CidrBlockAssociationOutput)
}

type VpcIpv4CidrBlockAssociationMapOutput struct{ *pulumi.OutputState }

func (VpcIpv4CidrBlockAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcIpv4CidrBlockAssociation)(nil))
}

func (o VpcIpv4CidrBlockAssociationMapOutput) ToVpcIpv4CidrBlockAssociationMapOutput() VpcIpv4CidrBlockAssociationMapOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationMapOutput) ToVpcIpv4CidrBlockAssociationMapOutputWithContext(ctx context.Context) VpcIpv4CidrBlockAssociationMapOutput {
	return o
}

func (o VpcIpv4CidrBlockAssociationMapOutput) MapIndex(k pulumi.StringInput) VpcIpv4CidrBlockAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcIpv4CidrBlockAssociation {
		return vs[0].(map[string]VpcIpv4CidrBlockAssociation)[vs[1].(string)]
	}).(VpcIpv4CidrBlockAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv4CidrBlockAssociationInput)(nil)).Elem(), &VpcIpv4CidrBlockAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv4CidrBlockAssociationPtrInput)(nil)).Elem(), &VpcIpv4CidrBlockAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv4CidrBlockAssociationArrayInput)(nil)).Elem(), VpcIpv4CidrBlockAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpv4CidrBlockAssociationMapInput)(nil)).Elem(), VpcIpv4CidrBlockAssociationMap{})
	pulumi.RegisterOutputType(VpcIpv4CidrBlockAssociationOutput{})
	pulumi.RegisterOutputType(VpcIpv4CidrBlockAssociationPtrOutput{})
	pulumi.RegisterOutputType(VpcIpv4CidrBlockAssociationArrayOutput{})
	pulumi.RegisterOutputType(VpcIpv4CidrBlockAssociationMapOutput{})
}
