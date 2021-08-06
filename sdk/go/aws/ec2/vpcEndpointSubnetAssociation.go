// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create an association between a VPC endpoint and a subnet.
//
// > **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
// both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
// and a single `subnetId`) and a VPC Endpoint resource with a `subnetIds`
// attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
// Association resource. Doing so will cause a conflict of associations and will overwrite the association.
//
// ## Example Usage
//
// Basic usage:
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
// 		_, err := ec2.NewVpcEndpointSubnetAssociation(ctx, "snEc2", &ec2.VpcEndpointSubnetAssociationArgs{
// 			VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Ec2.Id),
// 			SubnetId:      pulumi.Any(aws_subnet.Sn.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VpcEndpointSubnetAssociation struct {
	pulumi.CustomResourceState

	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId pulumi.StringOutput `pulumi:"vpcEndpointId"`
}

// NewVpcEndpointSubnetAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointSubnetAssociation(ctx *pulumi.Context,
	name string, args *VpcEndpointSubnetAssociationArgs, opts ...pulumi.ResourceOption) (*VpcEndpointSubnetAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'VpcEndpointId'")
	}
	var resource VpcEndpointSubnetAssociation
	err := ctx.RegisterResource("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointSubnetAssociation gets an existing VpcEndpointSubnetAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointSubnetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointSubnetAssociationState, opts ...pulumi.ResourceOption) (*VpcEndpointSubnetAssociation, error) {
	var resource VpcEndpointSubnetAssociation
	err := ctx.ReadResource("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointSubnetAssociation resources.
type vpcEndpointSubnetAssociationState struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId *string `pulumi:"subnetId"`
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
}

type VpcEndpointSubnetAssociationState struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId pulumi.StringPtrInput
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId pulumi.StringPtrInput
}

func (VpcEndpointSubnetAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointSubnetAssociationState)(nil)).Elem()
}

type vpcEndpointSubnetAssociationArgs struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId string `pulumi:"subnetId"`
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId string `pulumi:"vpcEndpointId"`
}

// The set of arguments for constructing a VpcEndpointSubnetAssociation resource.
type VpcEndpointSubnetAssociationArgs struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId pulumi.StringInput
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId pulumi.StringInput
}

func (VpcEndpointSubnetAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointSubnetAssociationArgs)(nil)).Elem()
}

type VpcEndpointSubnetAssociationInput interface {
	pulumi.Input

	ToVpcEndpointSubnetAssociationOutput() VpcEndpointSubnetAssociationOutput
	ToVpcEndpointSubnetAssociationOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationOutput
}

func (*VpcEndpointSubnetAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointSubnetAssociation)(nil))
}

func (i *VpcEndpointSubnetAssociation) ToVpcEndpointSubnetAssociationOutput() VpcEndpointSubnetAssociationOutput {
	return i.ToVpcEndpointSubnetAssociationOutputWithContext(context.Background())
}

func (i *VpcEndpointSubnetAssociation) ToVpcEndpointSubnetAssociationOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointSubnetAssociationOutput)
}

func (i *VpcEndpointSubnetAssociation) ToVpcEndpointSubnetAssociationPtrOutput() VpcEndpointSubnetAssociationPtrOutput {
	return i.ToVpcEndpointSubnetAssociationPtrOutputWithContext(context.Background())
}

func (i *VpcEndpointSubnetAssociation) ToVpcEndpointSubnetAssociationPtrOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointSubnetAssociationPtrOutput)
}

type VpcEndpointSubnetAssociationPtrInput interface {
	pulumi.Input

	ToVpcEndpointSubnetAssociationPtrOutput() VpcEndpointSubnetAssociationPtrOutput
	ToVpcEndpointSubnetAssociationPtrOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationPtrOutput
}

type vpcEndpointSubnetAssociationPtrType VpcEndpointSubnetAssociationArgs

func (*vpcEndpointSubnetAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointSubnetAssociation)(nil))
}

func (i *vpcEndpointSubnetAssociationPtrType) ToVpcEndpointSubnetAssociationPtrOutput() VpcEndpointSubnetAssociationPtrOutput {
	return i.ToVpcEndpointSubnetAssociationPtrOutputWithContext(context.Background())
}

func (i *vpcEndpointSubnetAssociationPtrType) ToVpcEndpointSubnetAssociationPtrOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointSubnetAssociationPtrOutput)
}

// VpcEndpointSubnetAssociationArrayInput is an input type that accepts VpcEndpointSubnetAssociationArray and VpcEndpointSubnetAssociationArrayOutput values.
// You can construct a concrete instance of `VpcEndpointSubnetAssociationArrayInput` via:
//
//          VpcEndpointSubnetAssociationArray{ VpcEndpointSubnetAssociationArgs{...} }
type VpcEndpointSubnetAssociationArrayInput interface {
	pulumi.Input

	ToVpcEndpointSubnetAssociationArrayOutput() VpcEndpointSubnetAssociationArrayOutput
	ToVpcEndpointSubnetAssociationArrayOutputWithContext(context.Context) VpcEndpointSubnetAssociationArrayOutput
}

type VpcEndpointSubnetAssociationArray []VpcEndpointSubnetAssociationInput

func (VpcEndpointSubnetAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointSubnetAssociation)(nil)).Elem()
}

func (i VpcEndpointSubnetAssociationArray) ToVpcEndpointSubnetAssociationArrayOutput() VpcEndpointSubnetAssociationArrayOutput {
	return i.ToVpcEndpointSubnetAssociationArrayOutputWithContext(context.Background())
}

func (i VpcEndpointSubnetAssociationArray) ToVpcEndpointSubnetAssociationArrayOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointSubnetAssociationArrayOutput)
}

// VpcEndpointSubnetAssociationMapInput is an input type that accepts VpcEndpointSubnetAssociationMap and VpcEndpointSubnetAssociationMapOutput values.
// You can construct a concrete instance of `VpcEndpointSubnetAssociationMapInput` via:
//
//          VpcEndpointSubnetAssociationMap{ "key": VpcEndpointSubnetAssociationArgs{...} }
type VpcEndpointSubnetAssociationMapInput interface {
	pulumi.Input

	ToVpcEndpointSubnetAssociationMapOutput() VpcEndpointSubnetAssociationMapOutput
	ToVpcEndpointSubnetAssociationMapOutputWithContext(context.Context) VpcEndpointSubnetAssociationMapOutput
}

type VpcEndpointSubnetAssociationMap map[string]VpcEndpointSubnetAssociationInput

func (VpcEndpointSubnetAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointSubnetAssociation)(nil)).Elem()
}

func (i VpcEndpointSubnetAssociationMap) ToVpcEndpointSubnetAssociationMapOutput() VpcEndpointSubnetAssociationMapOutput {
	return i.ToVpcEndpointSubnetAssociationMapOutputWithContext(context.Background())
}

func (i VpcEndpointSubnetAssociationMap) ToVpcEndpointSubnetAssociationMapOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointSubnetAssociationMapOutput)
}

type VpcEndpointSubnetAssociationOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointSubnetAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointSubnetAssociation)(nil))
}

func (o VpcEndpointSubnetAssociationOutput) ToVpcEndpointSubnetAssociationOutput() VpcEndpointSubnetAssociationOutput {
	return o
}

func (o VpcEndpointSubnetAssociationOutput) ToVpcEndpointSubnetAssociationOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationOutput {
	return o
}

func (o VpcEndpointSubnetAssociationOutput) ToVpcEndpointSubnetAssociationPtrOutput() VpcEndpointSubnetAssociationPtrOutput {
	return o.ToVpcEndpointSubnetAssociationPtrOutputWithContext(context.Background())
}

func (o VpcEndpointSubnetAssociationOutput) ToVpcEndpointSubnetAssociationPtrOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationPtrOutput {
	return o.ApplyT(func(v VpcEndpointSubnetAssociation) *VpcEndpointSubnetAssociation {
		return &v
	}).(VpcEndpointSubnetAssociationPtrOutput)
}

type VpcEndpointSubnetAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (VpcEndpointSubnetAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointSubnetAssociation)(nil))
}

func (o VpcEndpointSubnetAssociationPtrOutput) ToVpcEndpointSubnetAssociationPtrOutput() VpcEndpointSubnetAssociationPtrOutput {
	return o
}

func (o VpcEndpointSubnetAssociationPtrOutput) ToVpcEndpointSubnetAssociationPtrOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationPtrOutput {
	return o
}

type VpcEndpointSubnetAssociationArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointSubnetAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcEndpointSubnetAssociation)(nil))
}

func (o VpcEndpointSubnetAssociationArrayOutput) ToVpcEndpointSubnetAssociationArrayOutput() VpcEndpointSubnetAssociationArrayOutput {
	return o
}

func (o VpcEndpointSubnetAssociationArrayOutput) ToVpcEndpointSubnetAssociationArrayOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationArrayOutput {
	return o
}

func (o VpcEndpointSubnetAssociationArrayOutput) Index(i pulumi.IntInput) VpcEndpointSubnetAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcEndpointSubnetAssociation {
		return vs[0].([]VpcEndpointSubnetAssociation)[vs[1].(int)]
	}).(VpcEndpointSubnetAssociationOutput)
}

type VpcEndpointSubnetAssociationMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointSubnetAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcEndpointSubnetAssociation)(nil))
}

func (o VpcEndpointSubnetAssociationMapOutput) ToVpcEndpointSubnetAssociationMapOutput() VpcEndpointSubnetAssociationMapOutput {
	return o
}

func (o VpcEndpointSubnetAssociationMapOutput) ToVpcEndpointSubnetAssociationMapOutputWithContext(ctx context.Context) VpcEndpointSubnetAssociationMapOutput {
	return o
}

func (o VpcEndpointSubnetAssociationMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointSubnetAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcEndpointSubnetAssociation {
		return vs[0].(map[string]VpcEndpointSubnetAssociation)[vs[1].(string)]
	}).(VpcEndpointSubnetAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointSubnetAssociationOutput{})
	pulumi.RegisterOutputType(VpcEndpointSubnetAssociationPtrOutput{})
	pulumi.RegisterOutputType(VpcEndpointSubnetAssociationArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointSubnetAssociationMapOutput{})
}
