// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Resource Access Manager (RAM) Resource Share. To associate principals with the share, see the `ram.PrincipalAssociation` resource. To associate resources with the share, see the `ram.ResourceAssociation` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ram"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ram.NewResourceShare(ctx, "example", &ram.ResourceShareArgs{
// 			AllowExternalPrincipals: pulumi.Bool(true),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Production"),
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
// Resource shares can be imported using the `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:ram/resourceShare:ResourceShare example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12
// ```
type ResourceShare struct {
	pulumi.CustomResourceState

	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolPtrOutput `pulumi:"allowExternalPrincipals"`
	// The Amazon Resource Name (ARN) of the resource share.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the resource share.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource share. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewResourceShare registers a new resource with the given unique name, arguments, and options.
func NewResourceShare(ctx *pulumi.Context,
	name string, args *ResourceShareArgs, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	if args == nil {
		args = &ResourceShareArgs{}
	}

	var resource ResourceShare
	err := ctx.RegisterResource("aws:ram/resourceShare:ResourceShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceShare gets an existing ResourceShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceShareState, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	var resource ResourceShare
	err := ctx.ReadResource("aws:ram/resourceShare:ResourceShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceShare resources.
type resourceShareState struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals *bool `pulumi:"allowExternalPrincipals"`
	// The Amazon Resource Name (ARN) of the resource share.
	Arn *string `pulumi:"arn"`
	// The name of the resource share.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource share. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ResourceShareState struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the resource share.
	Arn pulumi.StringPtrInput
	// The name of the resource share.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource share. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ResourceShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareState)(nil)).Elem()
}

type resourceShareArgs struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals *bool `pulumi:"allowExternalPrincipals"`
	// The name of the resource share.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource share. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceShare resource.
type ResourceShareArgs struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolPtrInput
	// The name of the resource share.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource share. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ResourceShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareArgs)(nil)).Elem()
}

type ResourceShareInput interface {
	pulumi.Input

	ToResourceShareOutput() ResourceShareOutput
	ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput
}

func (*ResourceShare) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceShare)(nil))
}

func (i *ResourceShare) ToResourceShareOutput() ResourceShareOutput {
	return i.ToResourceShareOutputWithContext(context.Background())
}

func (i *ResourceShare) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareOutput)
}

func (i *ResourceShare) ToResourceSharePtrOutput() ResourceSharePtrOutput {
	return i.ToResourceSharePtrOutputWithContext(context.Background())
}

func (i *ResourceShare) ToResourceSharePtrOutputWithContext(ctx context.Context) ResourceSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSharePtrOutput)
}

type ResourceSharePtrInput interface {
	pulumi.Input

	ToResourceSharePtrOutput() ResourceSharePtrOutput
	ToResourceSharePtrOutputWithContext(ctx context.Context) ResourceSharePtrOutput
}

type resourceSharePtrType ResourceShareArgs

func (*resourceSharePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceShare)(nil))
}

func (i *resourceSharePtrType) ToResourceSharePtrOutput() ResourceSharePtrOutput {
	return i.ToResourceSharePtrOutputWithContext(context.Background())
}

func (i *resourceSharePtrType) ToResourceSharePtrOutputWithContext(ctx context.Context) ResourceSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSharePtrOutput)
}

// ResourceShareArrayInput is an input type that accepts ResourceShareArray and ResourceShareArrayOutput values.
// You can construct a concrete instance of `ResourceShareArrayInput` via:
//
//          ResourceShareArray{ ResourceShareArgs{...} }
type ResourceShareArrayInput interface {
	pulumi.Input

	ToResourceShareArrayOutput() ResourceShareArrayOutput
	ToResourceShareArrayOutputWithContext(context.Context) ResourceShareArrayOutput
}

type ResourceShareArray []ResourceShareInput

func (ResourceShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceShare)(nil)).Elem()
}

func (i ResourceShareArray) ToResourceShareArrayOutput() ResourceShareArrayOutput {
	return i.ToResourceShareArrayOutputWithContext(context.Background())
}

func (i ResourceShareArray) ToResourceShareArrayOutputWithContext(ctx context.Context) ResourceShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareArrayOutput)
}

// ResourceShareMapInput is an input type that accepts ResourceShareMap and ResourceShareMapOutput values.
// You can construct a concrete instance of `ResourceShareMapInput` via:
//
//          ResourceShareMap{ "key": ResourceShareArgs{...} }
type ResourceShareMapInput interface {
	pulumi.Input

	ToResourceShareMapOutput() ResourceShareMapOutput
	ToResourceShareMapOutputWithContext(context.Context) ResourceShareMapOutput
}

type ResourceShareMap map[string]ResourceShareInput

func (ResourceShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceShare)(nil)).Elem()
}

func (i ResourceShareMap) ToResourceShareMapOutput() ResourceShareMapOutput {
	return i.ToResourceShareMapOutputWithContext(context.Background())
}

func (i ResourceShareMap) ToResourceShareMapOutputWithContext(ctx context.Context) ResourceShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareMapOutput)
}

type ResourceShareOutput struct{ *pulumi.OutputState }

func (ResourceShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceShare)(nil))
}

func (o ResourceShareOutput) ToResourceShareOutput() ResourceShareOutput {
	return o
}

func (o ResourceShareOutput) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return o
}

func (o ResourceShareOutput) ToResourceSharePtrOutput() ResourceSharePtrOutput {
	return o.ToResourceSharePtrOutputWithContext(context.Background())
}

func (o ResourceShareOutput) ToResourceSharePtrOutputWithContext(ctx context.Context) ResourceSharePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceShare) *ResourceShare {
		return &v
	}).(ResourceSharePtrOutput)
}

type ResourceSharePtrOutput struct{ *pulumi.OutputState }

func (ResourceSharePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceShare)(nil))
}

func (o ResourceSharePtrOutput) ToResourceSharePtrOutput() ResourceSharePtrOutput {
	return o
}

func (o ResourceSharePtrOutput) ToResourceSharePtrOutputWithContext(ctx context.Context) ResourceSharePtrOutput {
	return o
}

func (o ResourceSharePtrOutput) Elem() ResourceShareOutput {
	return o.ApplyT(func(v *ResourceShare) ResourceShare {
		if v != nil {
			return *v
		}
		var ret ResourceShare
		return ret
	}).(ResourceShareOutput)
}

type ResourceShareArrayOutput struct{ *pulumi.OutputState }

func (ResourceShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceShare)(nil))
}

func (o ResourceShareArrayOutput) ToResourceShareArrayOutput() ResourceShareArrayOutput {
	return o
}

func (o ResourceShareArrayOutput) ToResourceShareArrayOutputWithContext(ctx context.Context) ResourceShareArrayOutput {
	return o
}

func (o ResourceShareArrayOutput) Index(i pulumi.IntInput) ResourceShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceShare {
		return vs[0].([]ResourceShare)[vs[1].(int)]
	}).(ResourceShareOutput)
}

type ResourceShareMapOutput struct{ *pulumi.OutputState }

func (ResourceShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceShare)(nil))
}

func (o ResourceShareMapOutput) ToResourceShareMapOutput() ResourceShareMapOutput {
	return o
}

func (o ResourceShareMapOutput) ToResourceShareMapOutputWithContext(ctx context.Context) ResourceShareMapOutput {
	return o
}

func (o ResourceShareMapOutput) MapIndex(k pulumi.StringInput) ResourceShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceShare {
		return vs[0].(map[string]ResourceShare)[vs[1].(string)]
	}).(ResourceShareOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareInput)(nil)).Elem(), &ResourceShare{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceSharePtrInput)(nil)).Elem(), &ResourceShare{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareArrayInput)(nil)).Elem(), ResourceShareArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceShareMapInput)(nil)).Elem(), ResourceShareMap{})
	pulumi.RegisterOutputType(ResourceShareOutput{})
	pulumi.RegisterOutputType(ResourceSharePtrOutput{})
	pulumi.RegisterOutputType(ResourceShareArrayOutput{})
	pulumi.RegisterOutputType(ResourceShareMapOutput{})
}
