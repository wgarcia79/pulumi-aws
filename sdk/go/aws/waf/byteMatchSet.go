// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF Byte Match Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.NewByteMatchSet(ctx, "byteSet", &waf.ByteMatchSetArgs{
// 			ByteMatchTuples: waf.ByteMatchSetByteMatchTupleArray{
// 				&waf.ByteMatchSetByteMatchTupleArgs{
// 					FieldToMatch: &waf.ByteMatchSetByteMatchTupleFieldToMatchArgs{
// 						Data: pulumi.String("referer"),
// 						Type: pulumi.String("HEADER"),
// 					},
// 					PositionalConstraint: pulumi.String("CONTAINS"),
// 					TargetString:         pulumi.String("badrefer1"),
// 					TextTransformation:   pulumi.String("NONE"),
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
// WAF Byte Match Set can be imported using the id, e.g.,
//
// ```sh
//  $ pulumi import aws:waf/byteMatchSet:ByteMatchSet byte_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
// ```
type ByteMatchSet struct {
	pulumi.CustomResourceState

	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayOutput `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewByteMatchSet registers a new resource with the given unique name, arguments, and options.
func NewByteMatchSet(ctx *pulumi.Context,
	name string, args *ByteMatchSetArgs, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	if args == nil {
		args = &ByteMatchSetArgs{}
	}

	var resource ByteMatchSet
	err := ctx.RegisterResource("aws:waf/byteMatchSet:ByteMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetByteMatchSet gets an existing ByteMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetByteMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ByteMatchSetState, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	var resource ByteMatchSet
	err := ctx.ReadResource("aws:waf/byteMatchSet:ByteMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ByteMatchSet resources.
type byteMatchSetState struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name *string `pulumi:"name"`
}

type ByteMatchSetState struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the Byte Match Set.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetState)(nil)).Elem()
}

type byteMatchSetArgs struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples []ByteMatchSetByteMatchTuple `pulumi:"byteMatchTuples"`
	// The name or description of the Byte Match Set.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	// Specifies the bytes (typically a string that corresponds
	// with ASCII characters) that you want to search for in web requests,
	// the location in requests that you want to search, and other settings.
	ByteMatchTuples ByteMatchSetByteMatchTupleArrayInput
	// The name or description of the Byte Match Set.
	Name pulumi.StringPtrInput
}

func (ByteMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetArgs)(nil)).Elem()
}

type ByteMatchSetInput interface {
	pulumi.Input

	ToByteMatchSetOutput() ByteMatchSetOutput
	ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput
}

func (*ByteMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSet)(nil))
}

func (i *ByteMatchSet) ToByteMatchSetOutput() ByteMatchSetOutput {
	return i.ToByteMatchSetOutputWithContext(context.Background())
}

func (i *ByteMatchSet) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetOutput)
}

func (i *ByteMatchSet) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return i.ToByteMatchSetPtrOutputWithContext(context.Background())
}

func (i *ByteMatchSet) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetPtrOutput)
}

type ByteMatchSetPtrInput interface {
	pulumi.Input

	ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput
	ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput
}

type byteMatchSetPtrType ByteMatchSetArgs

func (*byteMatchSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ByteMatchSet)(nil))
}

func (i *byteMatchSetPtrType) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return i.ToByteMatchSetPtrOutputWithContext(context.Background())
}

func (i *byteMatchSetPtrType) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetPtrOutput)
}

// ByteMatchSetArrayInput is an input type that accepts ByteMatchSetArray and ByteMatchSetArrayOutput values.
// You can construct a concrete instance of `ByteMatchSetArrayInput` via:
//
//          ByteMatchSetArray{ ByteMatchSetArgs{...} }
type ByteMatchSetArrayInput interface {
	pulumi.Input

	ToByteMatchSetArrayOutput() ByteMatchSetArrayOutput
	ToByteMatchSetArrayOutputWithContext(context.Context) ByteMatchSetArrayOutput
}

type ByteMatchSetArray []ByteMatchSetInput

func (ByteMatchSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ByteMatchSet)(nil)).Elem()
}

func (i ByteMatchSetArray) ToByteMatchSetArrayOutput() ByteMatchSetArrayOutput {
	return i.ToByteMatchSetArrayOutputWithContext(context.Background())
}

func (i ByteMatchSetArray) ToByteMatchSetArrayOutputWithContext(ctx context.Context) ByteMatchSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetArrayOutput)
}

// ByteMatchSetMapInput is an input type that accepts ByteMatchSetMap and ByteMatchSetMapOutput values.
// You can construct a concrete instance of `ByteMatchSetMapInput` via:
//
//          ByteMatchSetMap{ "key": ByteMatchSetArgs{...} }
type ByteMatchSetMapInput interface {
	pulumi.Input

	ToByteMatchSetMapOutput() ByteMatchSetMapOutput
	ToByteMatchSetMapOutputWithContext(context.Context) ByteMatchSetMapOutput
}

type ByteMatchSetMap map[string]ByteMatchSetInput

func (ByteMatchSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ByteMatchSet)(nil)).Elem()
}

func (i ByteMatchSetMap) ToByteMatchSetMapOutput() ByteMatchSetMapOutput {
	return i.ToByteMatchSetMapOutputWithContext(context.Background())
}

func (i ByteMatchSetMap) ToByteMatchSetMapOutputWithContext(ctx context.Context) ByteMatchSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetMapOutput)
}

type ByteMatchSetOutput struct{ *pulumi.OutputState }

func (ByteMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSet)(nil))
}

func (o ByteMatchSetOutput) ToByteMatchSetOutput() ByteMatchSetOutput {
	return o
}

func (o ByteMatchSetOutput) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return o
}

func (o ByteMatchSetOutput) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return o.ToByteMatchSetPtrOutputWithContext(context.Background())
}

func (o ByteMatchSetOutput) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ByteMatchSet) *ByteMatchSet {
		return &v
	}).(ByteMatchSetPtrOutput)
}

type ByteMatchSetPtrOutput struct{ *pulumi.OutputState }

func (ByteMatchSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ByteMatchSet)(nil))
}

func (o ByteMatchSetPtrOutput) ToByteMatchSetPtrOutput() ByteMatchSetPtrOutput {
	return o
}

func (o ByteMatchSetPtrOutput) ToByteMatchSetPtrOutputWithContext(ctx context.Context) ByteMatchSetPtrOutput {
	return o
}

func (o ByteMatchSetPtrOutput) Elem() ByteMatchSetOutput {
	return o.ApplyT(func(v *ByteMatchSet) ByteMatchSet {
		if v != nil {
			return *v
		}
		var ret ByteMatchSet
		return ret
	}).(ByteMatchSetOutput)
}

type ByteMatchSetArrayOutput struct{ *pulumi.OutputState }

func (ByteMatchSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ByteMatchSet)(nil))
}

func (o ByteMatchSetArrayOutput) ToByteMatchSetArrayOutput() ByteMatchSetArrayOutput {
	return o
}

func (o ByteMatchSetArrayOutput) ToByteMatchSetArrayOutputWithContext(ctx context.Context) ByteMatchSetArrayOutput {
	return o
}

func (o ByteMatchSetArrayOutput) Index(i pulumi.IntInput) ByteMatchSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ByteMatchSet {
		return vs[0].([]ByteMatchSet)[vs[1].(int)]
	}).(ByteMatchSetOutput)
}

type ByteMatchSetMapOutput struct{ *pulumi.OutputState }

func (ByteMatchSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ByteMatchSet)(nil))
}

func (o ByteMatchSetMapOutput) ToByteMatchSetMapOutput() ByteMatchSetMapOutput {
	return o
}

func (o ByteMatchSetMapOutput) ToByteMatchSetMapOutputWithContext(ctx context.Context) ByteMatchSetMapOutput {
	return o
}

func (o ByteMatchSetMapOutput) MapIndex(k pulumi.StringInput) ByteMatchSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ByteMatchSet {
		return vs[0].(map[string]ByteMatchSet)[vs[1].(string)]
	}).(ByteMatchSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ByteMatchSetInput)(nil)).Elem(), &ByteMatchSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*ByteMatchSetPtrInput)(nil)).Elem(), &ByteMatchSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*ByteMatchSetArrayInput)(nil)).Elem(), ByteMatchSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ByteMatchSetMapInput)(nil)).Elem(), ByteMatchSetMap{})
	pulumi.RegisterOutputType(ByteMatchSetOutput{})
	pulumi.RegisterOutputType(ByteMatchSetPtrOutput{})
	pulumi.RegisterOutputType(ByteMatchSetArrayOutput{})
	pulumi.RegisterOutputType(ByteMatchSetMapOutput{})
}
