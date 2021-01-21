// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ElastiCache parameter group resource.
//
// > **NOTE:** Attempting to remove the `reserved-memory` parameter when `family` is set to `redis2.6` or `redis2.8` may show a perpetual difference in this provider due to an Elasticache API limitation. Leave that parameter configured with any value to workaround the issue.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.NewParameterGroup(ctx, "_default", &elasticache.ParameterGroupArgs{
// 			Family: pulumi.String("redis2.8"),
// 			Parameters: elasticache.ParameterGroupParameterArray{
// 				&elasticache.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("activerehashing"),
// 					Value: pulumi.String("yes"),
// 				},
// 				&elasticache.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("min-slaves-to-write"),
// 					Value: pulumi.String("2"),
// 				},
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
// ElastiCache Parameter Groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:elasticache/parameterGroup:ParameterGroup default redis-params
// ```
type ParameterGroup struct {
	pulumi.CustomResourceState

	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The family of the ElastiCache parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the ElastiCache parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of ElastiCache parameters to apply.
	Parameters ParameterGroupParameterArrayOutput `pulumi:"parameters"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:elasticache/parameterGroup:ParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterGroupState, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	var resource ParameterGroup
	err := ctx.ReadResource("aws:elasticache/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the ElastiCache parameter group.
	Family *string `pulumi:"family"`
	// The name of the ElastiCache parameter.
	Name *string `pulumi:"name"`
	// A list of ElastiCache parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
}

type ParameterGroupState struct {
	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the ElastiCache parameter group.
	Family pulumi.StringPtrInput
	// The name of the ElastiCache parameter.
	Name pulumi.StringPtrInput
	// A list of ElastiCache parameters to apply.
	Parameters ParameterGroupParameterArrayInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the ElastiCache parameter group.
	Family string `pulumi:"family"`
	// The name of the ElastiCache parameter.
	Name *string `pulumi:"name"`
	// A list of ElastiCache parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// The description of the ElastiCache parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the ElastiCache parameter group.
	Family pulumi.StringInput
	// The name of the ElastiCache parameter.
	Name pulumi.StringPtrInput
	// A list of ElastiCache parameters to apply.
	Parameters ParameterGroupParameterArrayInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}

type ParameterGroupInput interface {
	pulumi.Input

	ToParameterGroupOutput() ParameterGroupOutput
	ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput
}

func (*ParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroup)(nil))
}

func (i *ParameterGroup) ToParameterGroupOutput() ParameterGroupOutput {
	return i.ToParameterGroupOutputWithContext(context.Background())
}

func (i *ParameterGroup) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupOutput)
}

func (i *ParameterGroup) ToParameterGroupPtrOutput() ParameterGroupPtrOutput {
	return i.ToParameterGroupPtrOutputWithContext(context.Background())
}

func (i *ParameterGroup) ToParameterGroupPtrOutputWithContext(ctx context.Context) ParameterGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupPtrOutput)
}

type ParameterGroupPtrInput interface {
	pulumi.Input

	ToParameterGroupPtrOutput() ParameterGroupPtrOutput
	ToParameterGroupPtrOutputWithContext(ctx context.Context) ParameterGroupPtrOutput
}

type parameterGroupPtrType ParameterGroupArgs

func (*parameterGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil))
}

func (i *parameterGroupPtrType) ToParameterGroupPtrOutput() ParameterGroupPtrOutput {
	return i.ToParameterGroupPtrOutputWithContext(context.Background())
}

func (i *parameterGroupPtrType) ToParameterGroupPtrOutputWithContext(ctx context.Context) ParameterGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupPtrOutput)
}

// ParameterGroupArrayInput is an input type that accepts ParameterGroupArray and ParameterGroupArrayOutput values.
// You can construct a concrete instance of `ParameterGroupArrayInput` via:
//
//          ParameterGroupArray{ ParameterGroupArgs{...} }
type ParameterGroupArrayInput interface {
	pulumi.Input

	ToParameterGroupArrayOutput() ParameterGroupArrayOutput
	ToParameterGroupArrayOutputWithContext(context.Context) ParameterGroupArrayOutput
}

type ParameterGroupArray []ParameterGroupInput

func (ParameterGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ParameterGroup)(nil))
}

func (i ParameterGroupArray) ToParameterGroupArrayOutput() ParameterGroupArrayOutput {
	return i.ToParameterGroupArrayOutputWithContext(context.Background())
}

func (i ParameterGroupArray) ToParameterGroupArrayOutputWithContext(ctx context.Context) ParameterGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupArrayOutput)
}

// ParameterGroupMapInput is an input type that accepts ParameterGroupMap and ParameterGroupMapOutput values.
// You can construct a concrete instance of `ParameterGroupMapInput` via:
//
//          ParameterGroupMap{ "key": ParameterGroupArgs{...} }
type ParameterGroupMapInput interface {
	pulumi.Input

	ToParameterGroupMapOutput() ParameterGroupMapOutput
	ToParameterGroupMapOutputWithContext(context.Context) ParameterGroupMapOutput
}

type ParameterGroupMap map[string]ParameterGroupInput

func (ParameterGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ParameterGroup)(nil))
}

func (i ParameterGroupMap) ToParameterGroupMapOutput() ParameterGroupMapOutput {
	return i.ToParameterGroupMapOutputWithContext(context.Background())
}

func (i ParameterGroupMap) ToParameterGroupMapOutputWithContext(ctx context.Context) ParameterGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupMapOutput)
}

type ParameterGroupOutput struct {
	*pulumi.OutputState
}

func (ParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroup)(nil))
}

func (o ParameterGroupOutput) ToParameterGroupOutput() ParameterGroupOutput {
	return o
}

func (o ParameterGroupOutput) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return o
}

func (o ParameterGroupOutput) ToParameterGroupPtrOutput() ParameterGroupPtrOutput {
	return o.ToParameterGroupPtrOutputWithContext(context.Background())
}

func (o ParameterGroupOutput) ToParameterGroupPtrOutputWithContext(ctx context.Context) ParameterGroupPtrOutput {
	return o.ApplyT(func(v ParameterGroup) *ParameterGroup {
		return &v
	}).(ParameterGroupPtrOutput)
}

type ParameterGroupPtrOutput struct {
	*pulumi.OutputState
}

func (ParameterGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterGroup)(nil))
}

func (o ParameterGroupPtrOutput) ToParameterGroupPtrOutput() ParameterGroupPtrOutput {
	return o
}

func (o ParameterGroupPtrOutput) ToParameterGroupPtrOutputWithContext(ctx context.Context) ParameterGroupPtrOutput {
	return o
}

type ParameterGroupArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroup)(nil))
}

func (o ParameterGroupArrayOutput) ToParameterGroupArrayOutput() ParameterGroupArrayOutput {
	return o
}

func (o ParameterGroupArrayOutput) ToParameterGroupArrayOutputWithContext(ctx context.Context) ParameterGroupArrayOutput {
	return o
}

func (o ParameterGroupArrayOutput) Index(i pulumi.IntInput) ParameterGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroup {
		return vs[0].([]ParameterGroup)[vs[1].(int)]
	}).(ParameterGroupOutput)
}

type ParameterGroupMapOutput struct{ *pulumi.OutputState }

func (ParameterGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterGroup)(nil))
}

func (o ParameterGroupMapOutput) ToParameterGroupMapOutput() ParameterGroupMapOutput {
	return o
}

func (o ParameterGroupMapOutput) ToParameterGroupMapOutputWithContext(ctx context.Context) ParameterGroupMapOutput {
	return o
}

func (o ParameterGroupMapOutput) MapIndex(k pulumi.StringInput) ParameterGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterGroup {
		return vs[0].(map[string]ParameterGroup)[vs[1].(string)]
	}).(ParameterGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(ParameterGroupOutput{})
	pulumi.RegisterOutputType(ParameterGroupPtrOutput{})
	pulumi.RegisterOutputType(ParameterGroupArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupMapOutput{})
}
