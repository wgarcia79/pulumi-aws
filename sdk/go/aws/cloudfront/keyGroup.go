// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// The following example below creates a CloudFront key group.
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		examplePublicKey, err := cloudfront.NewPublicKey(ctx, "examplePublicKey", &cloudfront.PublicKeyArgs{
// 			Comment:    pulumi.String("example public key"),
// 			EncodedKey: readFileOrPanic("public_key.pem"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudfront.NewKeyGroup(ctx, "exampleKeyGroup", &cloudfront.KeyGroupArgs{
// 			Comment: pulumi.String("example key group"),
// 			Items: pulumi.StringArray{
// 				examplePublicKey.ID(),
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
// CloudFront Key Group can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudfront/keyGroup:KeyGroup example 4b4f2r1c-315d-5c2e-f093-216t50jed10f
// ```
type KeyGroup struct {
	pulumi.CustomResourceState

	// A comment to describe the key group..
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The identifier for this version of the key group.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A list of the identifiers of the public keys in the key group.
	Items pulumi.StringArrayOutput `pulumi:"items"`
	// A name to identify the key group.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewKeyGroup registers a new resource with the given unique name, arguments, and options.
func NewKeyGroup(ctx *pulumi.Context,
	name string, args *KeyGroupArgs, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	var resource KeyGroup
	err := ctx.RegisterResource("aws:cloudfront/keyGroup:KeyGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyGroup gets an existing KeyGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyGroupState, opts ...pulumi.ResourceOption) (*KeyGroup, error) {
	var resource KeyGroup
	err := ctx.ReadResource("aws:cloudfront/keyGroup:KeyGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyGroup resources.
type keyGroupState struct {
	// A comment to describe the key group..
	Comment *string `pulumi:"comment"`
	// The identifier for this version of the key group.
	Etag *string `pulumi:"etag"`
	// A list of the identifiers of the public keys in the key group.
	Items []string `pulumi:"items"`
	// A name to identify the key group.
	Name *string `pulumi:"name"`
}

type KeyGroupState struct {
	// A comment to describe the key group..
	Comment pulumi.StringPtrInput
	// The identifier for this version of the key group.
	Etag pulumi.StringPtrInput
	// A list of the identifiers of the public keys in the key group.
	Items pulumi.StringArrayInput
	// A name to identify the key group.
	Name pulumi.StringPtrInput
}

func (KeyGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupState)(nil)).Elem()
}

type keyGroupArgs struct {
	// A comment to describe the key group..
	Comment *string `pulumi:"comment"`
	// A list of the identifiers of the public keys in the key group.
	Items []string `pulumi:"items"`
	// A name to identify the key group.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a KeyGroup resource.
type KeyGroupArgs struct {
	// A comment to describe the key group..
	Comment pulumi.StringPtrInput
	// A list of the identifiers of the public keys in the key group.
	Items pulumi.StringArrayInput
	// A name to identify the key group.
	Name pulumi.StringPtrInput
}

func (KeyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyGroupArgs)(nil)).Elem()
}

type KeyGroupInput interface {
	pulumi.Input

	ToKeyGroupOutput() KeyGroupOutput
	ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput
}

func (*KeyGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyGroup)(nil))
}

func (i *KeyGroup) ToKeyGroupOutput() KeyGroupOutput {
	return i.ToKeyGroupOutputWithContext(context.Background())
}

func (i *KeyGroup) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupOutput)
}

func (i *KeyGroup) ToKeyGroupPtrOutput() KeyGroupPtrOutput {
	return i.ToKeyGroupPtrOutputWithContext(context.Background())
}

func (i *KeyGroup) ToKeyGroupPtrOutputWithContext(ctx context.Context) KeyGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupPtrOutput)
}

type KeyGroupPtrInput interface {
	pulumi.Input

	ToKeyGroupPtrOutput() KeyGroupPtrOutput
	ToKeyGroupPtrOutputWithContext(ctx context.Context) KeyGroupPtrOutput
}

type keyGroupPtrType KeyGroupArgs

func (*keyGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyGroup)(nil))
}

func (i *keyGroupPtrType) ToKeyGroupPtrOutput() KeyGroupPtrOutput {
	return i.ToKeyGroupPtrOutputWithContext(context.Background())
}

func (i *keyGroupPtrType) ToKeyGroupPtrOutputWithContext(ctx context.Context) KeyGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupPtrOutput)
}

// KeyGroupArrayInput is an input type that accepts KeyGroupArray and KeyGroupArrayOutput values.
// You can construct a concrete instance of `KeyGroupArrayInput` via:
//
//          KeyGroupArray{ KeyGroupArgs{...} }
type KeyGroupArrayInput interface {
	pulumi.Input

	ToKeyGroupArrayOutput() KeyGroupArrayOutput
	ToKeyGroupArrayOutputWithContext(context.Context) KeyGroupArrayOutput
}

type KeyGroupArray []KeyGroupInput

func (KeyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyGroup)(nil)).Elem()
}

func (i KeyGroupArray) ToKeyGroupArrayOutput() KeyGroupArrayOutput {
	return i.ToKeyGroupArrayOutputWithContext(context.Background())
}

func (i KeyGroupArray) ToKeyGroupArrayOutputWithContext(ctx context.Context) KeyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupArrayOutput)
}

// KeyGroupMapInput is an input type that accepts KeyGroupMap and KeyGroupMapOutput values.
// You can construct a concrete instance of `KeyGroupMapInput` via:
//
//          KeyGroupMap{ "key": KeyGroupArgs{...} }
type KeyGroupMapInput interface {
	pulumi.Input

	ToKeyGroupMapOutput() KeyGroupMapOutput
	ToKeyGroupMapOutputWithContext(context.Context) KeyGroupMapOutput
}

type KeyGroupMap map[string]KeyGroupInput

func (KeyGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyGroup)(nil)).Elem()
}

func (i KeyGroupMap) ToKeyGroupMapOutput() KeyGroupMapOutput {
	return i.ToKeyGroupMapOutputWithContext(context.Background())
}

func (i KeyGroupMap) ToKeyGroupMapOutputWithContext(ctx context.Context) KeyGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyGroupMapOutput)
}

type KeyGroupOutput struct{ *pulumi.OutputState }

func (KeyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyGroup)(nil))
}

func (o KeyGroupOutput) ToKeyGroupOutput() KeyGroupOutput {
	return o
}

func (o KeyGroupOutput) ToKeyGroupOutputWithContext(ctx context.Context) KeyGroupOutput {
	return o
}

func (o KeyGroupOutput) ToKeyGroupPtrOutput() KeyGroupPtrOutput {
	return o.ToKeyGroupPtrOutputWithContext(context.Background())
}

func (o KeyGroupOutput) ToKeyGroupPtrOutputWithContext(ctx context.Context) KeyGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyGroup) *KeyGroup {
		return &v
	}).(KeyGroupPtrOutput)
}

type KeyGroupPtrOutput struct{ *pulumi.OutputState }

func (KeyGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyGroup)(nil))
}

func (o KeyGroupPtrOutput) ToKeyGroupPtrOutput() KeyGroupPtrOutput {
	return o
}

func (o KeyGroupPtrOutput) ToKeyGroupPtrOutputWithContext(ctx context.Context) KeyGroupPtrOutput {
	return o
}

func (o KeyGroupPtrOutput) Elem() KeyGroupOutput {
	return o.ApplyT(func(v *KeyGroup) KeyGroup {
		if v != nil {
			return *v
		}
		var ret KeyGroup
		return ret
	}).(KeyGroupOutput)
}

type KeyGroupArrayOutput struct{ *pulumi.OutputState }

func (KeyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyGroup)(nil))
}

func (o KeyGroupArrayOutput) ToKeyGroupArrayOutput() KeyGroupArrayOutput {
	return o
}

func (o KeyGroupArrayOutput) ToKeyGroupArrayOutputWithContext(ctx context.Context) KeyGroupArrayOutput {
	return o
}

func (o KeyGroupArrayOutput) Index(i pulumi.IntInput) KeyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyGroup {
		return vs[0].([]KeyGroup)[vs[1].(int)]
	}).(KeyGroupOutput)
}

type KeyGroupMapOutput struct{ *pulumi.OutputState }

func (KeyGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyGroup)(nil))
}

func (o KeyGroupMapOutput) ToKeyGroupMapOutput() KeyGroupMapOutput {
	return o
}

func (o KeyGroupMapOutput) ToKeyGroupMapOutputWithContext(ctx context.Context) KeyGroupMapOutput {
	return o
}

func (o KeyGroupMapOutput) MapIndex(k pulumi.StringInput) KeyGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyGroup {
		return vs[0].(map[string]KeyGroup)[vs[1].(string)]
	}).(KeyGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyGroupInput)(nil)).Elem(), &KeyGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyGroupPtrInput)(nil)).Elem(), &KeyGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyGroupArrayInput)(nil)).Elem(), KeyGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyGroupMapInput)(nil)).Elem(), KeyGroupMap{})
	pulumi.RegisterOutputType(KeyGroupOutput{})
	pulumi.RegisterOutputType(KeyGroupPtrOutput{})
	pulumi.RegisterOutputType(KeyGroupArrayOutput{})
	pulumi.RegisterOutputType(KeyGroupMapOutput{})
}
