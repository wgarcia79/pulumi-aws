// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EC2 placement group. Read more about placement groups
// in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
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
// 		_, err := ec2.NewPlacementGroup(ctx, "web", &ec2.PlacementGroupArgs{
// 			Strategy: pulumi.String("cluster"),
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
// Placement groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/placementGroup:PlacementGroup prod_pg production-placement-group
// ```
type PlacementGroup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the placement group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the placement group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the placement group.
	PlacementGroupId pulumi.StringOutput `pulumi:"placementGroupId"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	var resource PlacementGroup
	err := ctx.RegisterResource("aws:ec2/placementGroup:PlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementGroupState, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	var resource PlacementGroup
	err := ctx.ReadResource("aws:ec2/placementGroup:PlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementGroup resources.
type placementGroupState struct {
	// Amazon Resource Name (ARN) of the placement group.
	Arn *string `pulumi:"arn"`
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The ID of the placement group.
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy *string `pulumi:"strategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PlacementGroupState struct {
	// Amazon Resource Name (ARN) of the placement group.
	Arn pulumi.StringPtrInput
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The ID of the placement group.
	PlacementGroupId pulumi.StringPtrInput
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy string `pulumi:"strategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

type PlacementGroupInput interface {
	pulumi.Input

	ToPlacementGroupOutput() PlacementGroupOutput
	ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput
}

func (*PlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (i *PlacementGroup) ToPlacementGroupOutput() PlacementGroupOutput {
	return i.ToPlacementGroupOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupOutput)
}

func (i *PlacementGroup) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return i.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupPtrOutput)
}

type PlacementGroupPtrInput interface {
	pulumi.Input

	ToPlacementGroupPtrOutput() PlacementGroupPtrOutput
	ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput
}

type placementGroupPtrType PlacementGroupArgs

func (*placementGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil))
}

func (i *placementGroupPtrType) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return i.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (i *placementGroupPtrType) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupPtrOutput)
}

// PlacementGroupArrayInput is an input type that accepts PlacementGroupArray and PlacementGroupArrayOutput values.
// You can construct a concrete instance of `PlacementGroupArrayInput` via:
//
//          PlacementGroupArray{ PlacementGroupArgs{...} }
type PlacementGroupArrayInput interface {
	pulumi.Input

	ToPlacementGroupArrayOutput() PlacementGroupArrayOutput
	ToPlacementGroupArrayOutputWithContext(context.Context) PlacementGroupArrayOutput
}

type PlacementGroupArray []PlacementGroupInput

func (PlacementGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlacementGroup)(nil)).Elem()
}

func (i PlacementGroupArray) ToPlacementGroupArrayOutput() PlacementGroupArrayOutput {
	return i.ToPlacementGroupArrayOutputWithContext(context.Background())
}

func (i PlacementGroupArray) ToPlacementGroupArrayOutputWithContext(ctx context.Context) PlacementGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupArrayOutput)
}

// PlacementGroupMapInput is an input type that accepts PlacementGroupMap and PlacementGroupMapOutput values.
// You can construct a concrete instance of `PlacementGroupMapInput` via:
//
//          PlacementGroupMap{ "key": PlacementGroupArgs{...} }
type PlacementGroupMapInput interface {
	pulumi.Input

	ToPlacementGroupMapOutput() PlacementGroupMapOutput
	ToPlacementGroupMapOutputWithContext(context.Context) PlacementGroupMapOutput
}

type PlacementGroupMap map[string]PlacementGroupInput

func (PlacementGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlacementGroup)(nil)).Elem()
}

func (i PlacementGroupMap) ToPlacementGroupMapOutput() PlacementGroupMapOutput {
	return i.ToPlacementGroupMapOutputWithContext(context.Background())
}

func (i PlacementGroupMap) ToPlacementGroupMapOutputWithContext(ctx context.Context) PlacementGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupMapOutput)
}

type PlacementGroupOutput struct {
	*pulumi.OutputState
}

func (PlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (o PlacementGroupOutput) ToPlacementGroupOutput() PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return o.ToPlacementGroupPtrOutputWithContext(context.Background())
}

func (o PlacementGroupOutput) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return o.ApplyT(func(v PlacementGroup) *PlacementGroup {
		return &v
	}).(PlacementGroupPtrOutput)
}

type PlacementGroupPtrOutput struct {
	*pulumi.OutputState
}

func (PlacementGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementGroup)(nil))
}

func (o PlacementGroupPtrOutput) ToPlacementGroupPtrOutput() PlacementGroupPtrOutput {
	return o
}

func (o PlacementGroupPtrOutput) ToPlacementGroupPtrOutputWithContext(ctx context.Context) PlacementGroupPtrOutput {
	return o
}

type PlacementGroupArrayOutput struct{ *pulumi.OutputState }

func (PlacementGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlacementGroup)(nil))
}

func (o PlacementGroupArrayOutput) ToPlacementGroupArrayOutput() PlacementGroupArrayOutput {
	return o
}

func (o PlacementGroupArrayOutput) ToPlacementGroupArrayOutputWithContext(ctx context.Context) PlacementGroupArrayOutput {
	return o
}

func (o PlacementGroupArrayOutput) Index(i pulumi.IntInput) PlacementGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlacementGroup {
		return vs[0].([]PlacementGroup)[vs[1].(int)]
	}).(PlacementGroupOutput)
}

type PlacementGroupMapOutput struct{ *pulumi.OutputState }

func (PlacementGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PlacementGroup)(nil))
}

func (o PlacementGroupMapOutput) ToPlacementGroupMapOutput() PlacementGroupMapOutput {
	return o
}

func (o PlacementGroupMapOutput) ToPlacementGroupMapOutputWithContext(ctx context.Context) PlacementGroupMapOutput {
	return o
}

func (o PlacementGroupMapOutput) MapIndex(k pulumi.StringInput) PlacementGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PlacementGroup {
		return vs[0].(map[string]PlacementGroup)[vs[1].(string)]
	}).(PlacementGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(PlacementGroupOutput{})
	pulumi.RegisterOutputType(PlacementGroupPtrOutput{})
	pulumi.RegisterOutputType(PlacementGroupArrayOutput{})
	pulumi.RegisterOutputType(PlacementGroupMapOutput{})
}
