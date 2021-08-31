// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// `aws_ecs_tag` can be imported by using the ECS resource identifier and key, separated by a comma (`,`), e.g.
//
// ```sh
//  $ pulumi import aws:ecs/tag:Tag example arn:aws:ecs:us-east-1:123456789012:cluster/example,Name
// ```
type Tag struct {
	pulumi.CustomResourceState

	// Tag name.
	Key pulumi.StringOutput `pulumi:"key"`
	// Amazon Resource Name (ARN) of the ECS resource to tag.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// Tag value.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Tag
	err := ctx.RegisterResource("aws:ecs/tag:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("aws:ecs/tag:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
	// Tag name.
	Key *string `pulumi:"key"`
	// Amazon Resource Name (ARN) of the ECS resource to tag.
	ResourceArn *string `pulumi:"resourceArn"`
	// Tag value.
	Value *string `pulumi:"value"`
}

type TagState struct {
	// Tag name.
	Key pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the ECS resource to tag.
	ResourceArn pulumi.StringPtrInput
	// Tag value.
	Value pulumi.StringPtrInput
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// Tag name.
	Key string `pulumi:"key"`
	// Amazon Resource Name (ARN) of the ECS resource to tag.
	ResourceArn string `pulumi:"resourceArn"`
	// Tag value.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// Tag name.
	Key pulumi.StringInput
	// Amazon Resource Name (ARN) of the ECS resource to tag.
	ResourceArn pulumi.StringInput
	// Tag value.
	Value pulumi.StringInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

func (*Tag) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil))
}

func (i *Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i *Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

func (i *Tag) ToTagPtrOutput() TagPtrOutput {
	return i.ToTagPtrOutputWithContext(context.Background())
}

func (i *Tag) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPtrOutput)
}

type TagPtrInput interface {
	pulumi.Input

	ToTagPtrOutput() TagPtrOutput
	ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput
}

type tagPtrType TagArgs

func (*tagPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil))
}

func (i *tagPtrType) ToTagPtrOutput() TagPtrOutput {
	return i.ToTagPtrOutputWithContext(context.Background())
}

func (i *tagPtrType) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagPtrOutput)
}

// TagArrayInput is an input type that accepts TagArray and TagArrayOutput values.
// You can construct a concrete instance of `TagArrayInput` via:
//
//          TagArray{ TagArgs{...} }
type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

// TagMapInput is an input type that accepts TagMap and TagMapOutput values.
// You can construct a concrete instance of `TagMapInput` via:
//
//          TagMap{ "key": TagArgs{...} }
type TagMapInput interface {
	pulumi.Input

	ToTagMapOutput() TagMapOutput
	ToTagMapOutputWithContext(context.Context) TagMapOutput
}

type TagMap map[string]TagInput

func (TagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tag)(nil)).Elem()
}

func (i TagMap) ToTagMapOutput() TagMapOutput {
	return i.ToTagMapOutputWithContext(context.Background())
}

func (i TagMap) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagMapOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil))
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) ToTagPtrOutput() TagPtrOutput {
	return o.ToTagPtrOutputWithContext(context.Background())
}

func (o TagOutput) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Tag) *Tag {
		return &v
	}).(TagPtrOutput)
}

type TagPtrOutput struct{ *pulumi.OutputState }

func (TagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil))
}

func (o TagPtrOutput) ToTagPtrOutput() TagPtrOutput {
	return o
}

func (o TagPtrOutput) ToTagPtrOutputWithContext(ctx context.Context) TagPtrOutput {
	return o
}

func (o TagPtrOutput) Elem() TagOutput {
	return o.ApplyT(func(v *Tag) Tag {
		if v != nil {
			return *v
		}
		var ret Tag
		return ret
	}).(TagOutput)
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil))
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

func (o TagArrayOutput) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return o
}

func (o TagArrayOutput) Index(i pulumi.IntInput) TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Tag {
		return vs[0].([]Tag)[vs[1].(int)]
	}).(TagOutput)
}

type TagMapOutput struct{ *pulumi.OutputState }

func (TagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Tag)(nil))
}

func (o TagMapOutput) ToTagMapOutput() TagMapOutput {
	return o
}

func (o TagMapOutput) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return o
}

func (o TagMapOutput) MapIndex(k pulumi.StringInput) TagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Tag {
		return vs[0].(map[string]Tag)[vs[1].(string)]
	}).(TagOutput)
}

func init() {
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagPtrOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagMapOutput{})
}
