// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DAX Subnet Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dax"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dax.NewSubnetGroup(ctx, "example", &dax.SubnetGroupArgs{
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example1.Id),
// 				pulumi.Any(aws_subnet.Example2.Id),
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
// DAX Subnet Group can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:dax/subnetGroup:SubnetGroup example my_dax_sg
// ```
type SubnetGroup struct {
	pulumi.CustomResourceState

	// A description of the subnet group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the subnet group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// VPC ID of the subnet group.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:dax/subnetGroup:SubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetGroupState, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	var resource SubnetGroup
	err := ctx.ReadResource("aws:dax/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	// A description of the subnet group.
	Description *string `pulumi:"description"`
	// The name of the subnet group.
	Name *string `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// VPC ID of the subnet group.
	VpcId *string `pulumi:"vpcId"`
}

type SubnetGroupState struct {
	// A description of the subnet group.
	Description pulumi.StringPtrInput
	// The name of the subnet group.
	Name pulumi.StringPtrInput
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// VPC ID of the subnet group.
	VpcId pulumi.StringPtrInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// A description of the subnet group.
	Description *string `pulumi:"description"`
	// The name of the subnet group.
	Name *string `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// A description of the subnet group.
	Description pulumi.StringPtrInput
	// The name of the subnet group.
	Name pulumi.StringPtrInput
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}

type SubnetGroupInput interface {
	pulumi.Input

	ToSubnetGroupOutput() SubnetGroupOutput
	ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput
}

func (*SubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroup)(nil))
}

func (i *SubnetGroup) ToSubnetGroupOutput() SubnetGroupOutput {
	return i.ToSubnetGroupOutputWithContext(context.Background())
}

func (i *SubnetGroup) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupOutput)
}

func (i *SubnetGroup) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return i.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *SubnetGroup) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupPtrOutput)
}

type SubnetGroupPtrInput interface {
	pulumi.Input

	ToSubnetGroupPtrOutput() SubnetGroupPtrOutput
	ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput
}

type subnetGroupPtrType SubnetGroupArgs

func (*subnetGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetGroup)(nil))
}

func (i *subnetGroupPtrType) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return i.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *subnetGroupPtrType) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupPtrOutput)
}

// SubnetGroupArrayInput is an input type that accepts SubnetGroupArray and SubnetGroupArrayOutput values.
// You can construct a concrete instance of `SubnetGroupArrayInput` via:
//
//          SubnetGroupArray{ SubnetGroupArgs{...} }
type SubnetGroupArrayInput interface {
	pulumi.Input

	ToSubnetGroupArrayOutput() SubnetGroupArrayOutput
	ToSubnetGroupArrayOutputWithContext(context.Context) SubnetGroupArrayOutput
}

type SubnetGroupArray []SubnetGroupInput

func (SubnetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetGroup)(nil)).Elem()
}

func (i SubnetGroupArray) ToSubnetGroupArrayOutput() SubnetGroupArrayOutput {
	return i.ToSubnetGroupArrayOutputWithContext(context.Background())
}

func (i SubnetGroupArray) ToSubnetGroupArrayOutputWithContext(ctx context.Context) SubnetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupArrayOutput)
}

// SubnetGroupMapInput is an input type that accepts SubnetGroupMap and SubnetGroupMapOutput values.
// You can construct a concrete instance of `SubnetGroupMapInput` via:
//
//          SubnetGroupMap{ "key": SubnetGroupArgs{...} }
type SubnetGroupMapInput interface {
	pulumi.Input

	ToSubnetGroupMapOutput() SubnetGroupMapOutput
	ToSubnetGroupMapOutputWithContext(context.Context) SubnetGroupMapOutput
}

type SubnetGroupMap map[string]SubnetGroupInput

func (SubnetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetGroup)(nil)).Elem()
}

func (i SubnetGroupMap) ToSubnetGroupMapOutput() SubnetGroupMapOutput {
	return i.ToSubnetGroupMapOutputWithContext(context.Background())
}

func (i SubnetGroupMap) ToSubnetGroupMapOutputWithContext(ctx context.Context) SubnetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupMapOutput)
}

type SubnetGroupOutput struct{ *pulumi.OutputState }

func (SubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroup)(nil))
}

func (o SubnetGroupOutput) ToSubnetGroupOutput() SubnetGroupOutput {
	return o
}

func (o SubnetGroupOutput) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return o
}

func (o SubnetGroupOutput) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return o.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (o SubnetGroupOutput) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetGroup) *SubnetGroup {
		return &v
	}).(SubnetGroupPtrOutput)
}

type SubnetGroupPtrOutput struct{ *pulumi.OutputState }

func (SubnetGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetGroup)(nil))
}

func (o SubnetGroupPtrOutput) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return o
}

func (o SubnetGroupPtrOutput) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return o
}

func (o SubnetGroupPtrOutput) Elem() SubnetGroupOutput {
	return o.ApplyT(func(v *SubnetGroup) SubnetGroup {
		if v != nil {
			return *v
		}
		var ret SubnetGroup
		return ret
	}).(SubnetGroupOutput)
}

type SubnetGroupArrayOutput struct{ *pulumi.OutputState }

func (SubnetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetGroup)(nil))
}

func (o SubnetGroupArrayOutput) ToSubnetGroupArrayOutput() SubnetGroupArrayOutput {
	return o
}

func (o SubnetGroupArrayOutput) ToSubnetGroupArrayOutputWithContext(ctx context.Context) SubnetGroupArrayOutput {
	return o
}

func (o SubnetGroupArrayOutput) Index(i pulumi.IntInput) SubnetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetGroup {
		return vs[0].([]SubnetGroup)[vs[1].(int)]
	}).(SubnetGroupOutput)
}

type SubnetGroupMapOutput struct{ *pulumi.OutputState }

func (SubnetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SubnetGroup)(nil))
}

func (o SubnetGroupMapOutput) ToSubnetGroupMapOutput() SubnetGroupMapOutput {
	return o
}

func (o SubnetGroupMapOutput) ToSubnetGroupMapOutputWithContext(ctx context.Context) SubnetGroupMapOutput {
	return o
}

func (o SubnetGroupMapOutput) MapIndex(k pulumi.StringInput) SubnetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SubnetGroup {
		return vs[0].(map[string]SubnetGroup)[vs[1].(string)]
	}).(SubnetGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetGroupOutput{})
	pulumi.RegisterOutputType(SubnetGroupPtrOutput{})
	pulumi.RegisterOutputType(SubnetGroupArrayOutput{})
	pulumi.RegisterOutputType(SubnetGroupMapOutput{})
}
