// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Gamelift Alias resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewAlias(ctx, "example", &gamelift.AliasArgs{
// 			Description: pulumi.String("Example Description"),
// 			RoutingStrategy: &gamelift.AliasRoutingStrategyArgs{
// 				Message: pulumi.String("Example Message"),
// 				Type:    pulumi.String("TERMINAL"),
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
// Gamelift Aliases can be imported using the ID, e.g.,
//
// ```sh
//  $ pulumi import aws:gamelift/alias:Alias example <alias-id>
// ```
type Alias struct {
	pulumi.CustomResourceState

	// Alias ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the alias.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyOutput `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoutingStrategy == nil {
		return nil, errors.New("invalid value for required argument 'RoutingStrategy'")
	}
	var resource Alias
	err := ctx.RegisterResource("aws:gamelift/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:gamelift/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// Alias ARN.
	Arn *string `pulumi:"arn"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy *AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AliasState struct {
	// Alias ARN.
	Arn pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((*Alias)(nil))
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

func (i *Alias) ToAliasPtrOutput() AliasPtrOutput {
	return i.ToAliasPtrOutputWithContext(context.Background())
}

func (i *Alias) ToAliasPtrOutputWithContext(ctx context.Context) AliasPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPtrOutput)
}

type AliasPtrInput interface {
	pulumi.Input

	ToAliasPtrOutput() AliasPtrOutput
	ToAliasPtrOutputWithContext(ctx context.Context) AliasPtrOutput
}

type aliasPtrType AliasArgs

func (*aliasPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil))
}

func (i *aliasPtrType) ToAliasPtrOutput() AliasPtrOutput {
	return i.ToAliasPtrOutputWithContext(context.Background())
}

func (i *aliasPtrType) ToAliasPtrOutputWithContext(ctx context.Context) AliasPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasPtrOutput)
}

// AliasArrayInput is an input type that accepts AliasArray and AliasArrayOutput values.
// You can construct a concrete instance of `AliasArrayInput` via:
//
//          AliasArray{ AliasArgs{...} }
type AliasArrayInput interface {
	pulumi.Input

	ToAliasArrayOutput() AliasArrayOutput
	ToAliasArrayOutputWithContext(context.Context) AliasArrayOutput
}

type AliasArray []AliasInput

func (AliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alias)(nil)).Elem()
}

func (i AliasArray) ToAliasArrayOutput() AliasArrayOutput {
	return i.ToAliasArrayOutputWithContext(context.Background())
}

func (i AliasArray) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasArrayOutput)
}

// AliasMapInput is an input type that accepts AliasMap and AliasMapOutput values.
// You can construct a concrete instance of `AliasMapInput` via:
//
//          AliasMap{ "key": AliasArgs{...} }
type AliasMapInput interface {
	pulumi.Input

	ToAliasMapOutput() AliasMapOutput
	ToAliasMapOutputWithContext(context.Context) AliasMapOutput
}

type AliasMap map[string]AliasInput

func (AliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alias)(nil)).Elem()
}

func (i AliasMap) ToAliasMapOutput() AliasMapOutput {
	return i.ToAliasMapOutputWithContext(context.Background())
}

func (i AliasMap) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasMapOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Alias)(nil))
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func (o AliasOutput) ToAliasPtrOutput() AliasPtrOutput {
	return o.ToAliasPtrOutputWithContext(context.Background())
}

func (o AliasOutput) ToAliasPtrOutputWithContext(ctx context.Context) AliasPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Alias) *Alias {
		return &v
	}).(AliasPtrOutput)
}

type AliasPtrOutput struct{ *pulumi.OutputState }

func (AliasPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil))
}

func (o AliasPtrOutput) ToAliasPtrOutput() AliasPtrOutput {
	return o
}

func (o AliasPtrOutput) ToAliasPtrOutputWithContext(ctx context.Context) AliasPtrOutput {
	return o
}

func (o AliasPtrOutput) Elem() AliasOutput {
	return o.ApplyT(func(v *Alias) Alias {
		if v != nil {
			return *v
		}
		var ret Alias
		return ret
	}).(AliasOutput)
}

type AliasArrayOutput struct{ *pulumi.OutputState }

func (AliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Alias)(nil))
}

func (o AliasArrayOutput) ToAliasArrayOutput() AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) ToAliasArrayOutputWithContext(ctx context.Context) AliasArrayOutput {
	return o
}

func (o AliasArrayOutput) Index(i pulumi.IntInput) AliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Alias {
		return vs[0].([]Alias)[vs[1].(int)]
	}).(AliasOutput)
}

type AliasMapOutput struct{ *pulumi.OutputState }

func (AliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Alias)(nil))
}

func (o AliasMapOutput) ToAliasMapOutput() AliasMapOutput {
	return o
}

func (o AliasMapOutput) ToAliasMapOutputWithContext(ctx context.Context) AliasMapOutput {
	return o
}

func (o AliasMapOutput) MapIndex(k pulumi.StringInput) AliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Alias {
		return vs[0].(map[string]Alias)[vs[1].(string)]
	}).(AliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasPtrInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasArrayInput)(nil)).Elem(), AliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AliasMapInput)(nil)).Elem(), AliasMap{})
	pulumi.RegisterOutputType(AliasOutput{})
	pulumi.RegisterOutputType(AliasPtrOutput{})
	pulumi.RegisterOutputType(AliasArrayOutput{})
	pulumi.RegisterOutputType(AliasMapOutput{})
}
