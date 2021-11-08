// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Catalog Tag Option.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicecatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicecatalog.NewTagOption(ctx, "example", &servicecatalog.TagOptionArgs{
// 			Key:   pulumi.String("nyckel"),
// 			Value: pulumi.String("värde"),
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
// `aws_servicecatalog_tag_option` can be imported using the tag option ID, e.g.,
//
// ```sh
//  $ pulumi import aws:servicecatalog/tagOption:TagOption example tag-pjtvagohlyo3m
// ```
type TagOption struct {
	pulumi.CustomResourceState

	// Whether tag option is active. Default is `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Tag option key.
	Key   pulumi.StringOutput `pulumi:"key"`
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Tag option value.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewTagOption registers a new resource with the given unique name, arguments, and options.
func NewTagOption(ctx *pulumi.Context,
	name string, args *TagOptionArgs, opts ...pulumi.ResourceOption) (*TagOption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource TagOption
	err := ctx.RegisterResource("aws:servicecatalog/tagOption:TagOption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagOption gets an existing TagOption resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagOption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagOptionState, opts ...pulumi.ResourceOption) (*TagOption, error) {
	var resource TagOption
	err := ctx.ReadResource("aws:servicecatalog/tagOption:TagOption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagOption resources.
type tagOptionState struct {
	// Whether tag option is active. Default is `true`.
	Active *bool `pulumi:"active"`
	// Tag option key.
	Key   *string `pulumi:"key"`
	Owner *string `pulumi:"owner"`
	// Tag option value.
	Value *string `pulumi:"value"`
}

type TagOptionState struct {
	// Whether tag option is active. Default is `true`.
	Active pulumi.BoolPtrInput
	// Tag option key.
	Key   pulumi.StringPtrInput
	Owner pulumi.StringPtrInput
	// Tag option value.
	Value pulumi.StringPtrInput
}

func (TagOptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionState)(nil)).Elem()
}

type tagOptionArgs struct {
	// Whether tag option is active. Default is `true`.
	Active *bool `pulumi:"active"`
	// Tag option key.
	Key string `pulumi:"key"`
	// Tag option value.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a TagOption resource.
type TagOptionArgs struct {
	// Whether tag option is active. Default is `true`.
	Active pulumi.BoolPtrInput
	// Tag option key.
	Key pulumi.StringInput
	// Tag option value.
	Value pulumi.StringInput
}

func (TagOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionArgs)(nil)).Elem()
}

type TagOptionInput interface {
	pulumi.Input

	ToTagOptionOutput() TagOptionOutput
	ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput
}

func (*TagOption) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOption)(nil))
}

func (i *TagOption) ToTagOptionOutput() TagOptionOutput {
	return i.ToTagOptionOutputWithContext(context.Background())
}

func (i *TagOption) ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionOutput)
}

func (i *TagOption) ToTagOptionPtrOutput() TagOptionPtrOutput {
	return i.ToTagOptionPtrOutputWithContext(context.Background())
}

func (i *TagOption) ToTagOptionPtrOutputWithContext(ctx context.Context) TagOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionPtrOutput)
}

type TagOptionPtrInput interface {
	pulumi.Input

	ToTagOptionPtrOutput() TagOptionPtrOutput
	ToTagOptionPtrOutputWithContext(ctx context.Context) TagOptionPtrOutput
}

type tagOptionPtrType TagOptionArgs

func (*tagOptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOption)(nil))
}

func (i *tagOptionPtrType) ToTagOptionPtrOutput() TagOptionPtrOutput {
	return i.ToTagOptionPtrOutputWithContext(context.Background())
}

func (i *tagOptionPtrType) ToTagOptionPtrOutputWithContext(ctx context.Context) TagOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionPtrOutput)
}

// TagOptionArrayInput is an input type that accepts TagOptionArray and TagOptionArrayOutput values.
// You can construct a concrete instance of `TagOptionArrayInput` via:
//
//          TagOptionArray{ TagOptionArgs{...} }
type TagOptionArrayInput interface {
	pulumi.Input

	ToTagOptionArrayOutput() TagOptionArrayOutput
	ToTagOptionArrayOutputWithContext(context.Context) TagOptionArrayOutput
}

type TagOptionArray []TagOptionInput

func (TagOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TagOption)(nil)).Elem()
}

func (i TagOptionArray) ToTagOptionArrayOutput() TagOptionArrayOutput {
	return i.ToTagOptionArrayOutputWithContext(context.Background())
}

func (i TagOptionArray) ToTagOptionArrayOutputWithContext(ctx context.Context) TagOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionArrayOutput)
}

// TagOptionMapInput is an input type that accepts TagOptionMap and TagOptionMapOutput values.
// You can construct a concrete instance of `TagOptionMapInput` via:
//
//          TagOptionMap{ "key": TagOptionArgs{...} }
type TagOptionMapInput interface {
	pulumi.Input

	ToTagOptionMapOutput() TagOptionMapOutput
	ToTagOptionMapOutputWithContext(context.Context) TagOptionMapOutput
}

type TagOptionMap map[string]TagOptionInput

func (TagOptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TagOption)(nil)).Elem()
}

func (i TagOptionMap) ToTagOptionMapOutput() TagOptionMapOutput {
	return i.ToTagOptionMapOutputWithContext(context.Background())
}

func (i TagOptionMap) ToTagOptionMapOutputWithContext(ctx context.Context) TagOptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionMapOutput)
}

type TagOptionOutput struct{ *pulumi.OutputState }

func (TagOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOption)(nil))
}

func (o TagOptionOutput) ToTagOptionOutput() TagOptionOutput {
	return o
}

func (o TagOptionOutput) ToTagOptionOutputWithContext(ctx context.Context) TagOptionOutput {
	return o
}

func (o TagOptionOutput) ToTagOptionPtrOutput() TagOptionPtrOutput {
	return o.ToTagOptionPtrOutputWithContext(context.Background())
}

func (o TagOptionOutput) ToTagOptionPtrOutputWithContext(ctx context.Context) TagOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagOption) *TagOption {
		return &v
	}).(TagOptionPtrOutput)
}

type TagOptionPtrOutput struct{ *pulumi.OutputState }

func (TagOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOption)(nil))
}

func (o TagOptionPtrOutput) ToTagOptionPtrOutput() TagOptionPtrOutput {
	return o
}

func (o TagOptionPtrOutput) ToTagOptionPtrOutputWithContext(ctx context.Context) TagOptionPtrOutput {
	return o
}

func (o TagOptionPtrOutput) Elem() TagOptionOutput {
	return o.ApplyT(func(v *TagOption) TagOption {
		if v != nil {
			return *v
		}
		var ret TagOption
		return ret
	}).(TagOptionOutput)
}

type TagOptionArrayOutput struct{ *pulumi.OutputState }

func (TagOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagOption)(nil))
}

func (o TagOptionArrayOutput) ToTagOptionArrayOutput() TagOptionArrayOutput {
	return o
}

func (o TagOptionArrayOutput) ToTagOptionArrayOutputWithContext(ctx context.Context) TagOptionArrayOutput {
	return o
}

func (o TagOptionArrayOutput) Index(i pulumi.IntInput) TagOptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagOption {
		return vs[0].([]TagOption)[vs[1].(int)]
	}).(TagOptionOutput)
}

type TagOptionMapOutput struct{ *pulumi.OutputState }

func (TagOptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagOption)(nil))
}

func (o TagOptionMapOutput) ToTagOptionMapOutput() TagOptionMapOutput {
	return o
}

func (o TagOptionMapOutput) ToTagOptionMapOutputWithContext(ctx context.Context) TagOptionMapOutput {
	return o
}

func (o TagOptionMapOutput) MapIndex(k pulumi.StringInput) TagOptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagOption {
		return vs[0].(map[string]TagOption)[vs[1].(string)]
	}).(TagOptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagOptionInput)(nil)).Elem(), &TagOption{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagOptionPtrInput)(nil)).Elem(), &TagOption{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagOptionArrayInput)(nil)).Elem(), TagOptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagOptionMapInput)(nil)).Elem(), TagOptionMap{})
	pulumi.RegisterOutputType(TagOptionOutput{})
	pulumi.RegisterOutputType(TagOptionPtrOutput{})
	pulumi.RegisterOutputType(TagOptionArrayOutput{})
	pulumi.RegisterOutputType(TagOptionMapOutput{})
}
