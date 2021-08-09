// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Pinpoint GCM Channel resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewGcmChannel(ctx, "gcm", &pinpoint.GcmChannelArgs{
// 			ApplicationId: app.ApplicationId,
// 			ApiKey:        pulumi.String("api_key"),
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
// Pinpoint GCM Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/gcmChannel:GcmChannel gcm application-id
// ```
type GcmChannel struct {
	pulumi.CustomResourceState

	// Platform credential API key from Google.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewGcmChannel registers a new resource with the given unique name, arguments, and options.
func NewGcmChannel(ctx *pulumi.Context,
	name string, args *GcmChannelArgs, opts ...pulumi.ResourceOption) (*GcmChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource GcmChannel
	err := ctx.RegisterResource("aws:pinpoint/gcmChannel:GcmChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcmChannel gets an existing GcmChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcmChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcmChannelState, opts ...pulumi.ResourceOption) (*GcmChannel, error) {
	var resource GcmChannel
	err := ctx.ReadResource("aws:pinpoint/gcmChannel:GcmChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GcmChannel resources.
type gcmChannelState struct {
	// Platform credential API key from Google.
	ApiKey *string `pulumi:"apiKey"`
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

type GcmChannelState struct {
	// Platform credential API key from Google.
	ApiKey pulumi.StringPtrInput
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (GcmChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcmChannelState)(nil)).Elem()
}

type gcmChannelArgs struct {
	// Platform credential API key from Google.
	ApiKey string `pulumi:"apiKey"`
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

// The set of arguments for constructing a GcmChannel resource.
type GcmChannelArgs struct {
	// Platform credential API key from Google.
	ApiKey pulumi.StringInput
	// The application ID.
	ApplicationId pulumi.StringInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (GcmChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcmChannelArgs)(nil)).Elem()
}

type GcmChannelInput interface {
	pulumi.Input

	ToGcmChannelOutput() GcmChannelOutput
	ToGcmChannelOutputWithContext(ctx context.Context) GcmChannelOutput
}

func (*GcmChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmChannel)(nil))
}

func (i *GcmChannel) ToGcmChannelOutput() GcmChannelOutput {
	return i.ToGcmChannelOutputWithContext(context.Background())
}

func (i *GcmChannel) ToGcmChannelOutputWithContext(ctx context.Context) GcmChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmChannelOutput)
}

func (i *GcmChannel) ToGcmChannelPtrOutput() GcmChannelPtrOutput {
	return i.ToGcmChannelPtrOutputWithContext(context.Background())
}

func (i *GcmChannel) ToGcmChannelPtrOutputWithContext(ctx context.Context) GcmChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmChannelPtrOutput)
}

type GcmChannelPtrInput interface {
	pulumi.Input

	ToGcmChannelPtrOutput() GcmChannelPtrOutput
	ToGcmChannelPtrOutputWithContext(ctx context.Context) GcmChannelPtrOutput
}

type gcmChannelPtrType GcmChannelArgs

func (*gcmChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmChannel)(nil))
}

func (i *gcmChannelPtrType) ToGcmChannelPtrOutput() GcmChannelPtrOutput {
	return i.ToGcmChannelPtrOutputWithContext(context.Background())
}

func (i *gcmChannelPtrType) ToGcmChannelPtrOutputWithContext(ctx context.Context) GcmChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmChannelPtrOutput)
}

// GcmChannelArrayInput is an input type that accepts GcmChannelArray and GcmChannelArrayOutput values.
// You can construct a concrete instance of `GcmChannelArrayInput` via:
//
//          GcmChannelArray{ GcmChannelArgs{...} }
type GcmChannelArrayInput interface {
	pulumi.Input

	ToGcmChannelArrayOutput() GcmChannelArrayOutput
	ToGcmChannelArrayOutputWithContext(context.Context) GcmChannelArrayOutput
}

type GcmChannelArray []GcmChannelInput

func (GcmChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GcmChannel)(nil)).Elem()
}

func (i GcmChannelArray) ToGcmChannelArrayOutput() GcmChannelArrayOutput {
	return i.ToGcmChannelArrayOutputWithContext(context.Background())
}

func (i GcmChannelArray) ToGcmChannelArrayOutputWithContext(ctx context.Context) GcmChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmChannelArrayOutput)
}

// GcmChannelMapInput is an input type that accepts GcmChannelMap and GcmChannelMapOutput values.
// You can construct a concrete instance of `GcmChannelMapInput` via:
//
//          GcmChannelMap{ "key": GcmChannelArgs{...} }
type GcmChannelMapInput interface {
	pulumi.Input

	ToGcmChannelMapOutput() GcmChannelMapOutput
	ToGcmChannelMapOutputWithContext(context.Context) GcmChannelMapOutput
}

type GcmChannelMap map[string]GcmChannelInput

func (GcmChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GcmChannel)(nil)).Elem()
}

func (i GcmChannelMap) ToGcmChannelMapOutput() GcmChannelMapOutput {
	return i.ToGcmChannelMapOutputWithContext(context.Background())
}

func (i GcmChannelMap) ToGcmChannelMapOutputWithContext(ctx context.Context) GcmChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcmChannelMapOutput)
}

type GcmChannelOutput struct{ *pulumi.OutputState }

func (GcmChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcmChannel)(nil))
}

func (o GcmChannelOutput) ToGcmChannelOutput() GcmChannelOutput {
	return o
}

func (o GcmChannelOutput) ToGcmChannelOutputWithContext(ctx context.Context) GcmChannelOutput {
	return o
}

func (o GcmChannelOutput) ToGcmChannelPtrOutput() GcmChannelPtrOutput {
	return o.ToGcmChannelPtrOutputWithContext(context.Background())
}

func (o GcmChannelOutput) ToGcmChannelPtrOutputWithContext(ctx context.Context) GcmChannelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GcmChannel) *GcmChannel {
		return &v
	}).(GcmChannelPtrOutput)
}

type GcmChannelPtrOutput struct{ *pulumi.OutputState }

func (GcmChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcmChannel)(nil))
}

func (o GcmChannelPtrOutput) ToGcmChannelPtrOutput() GcmChannelPtrOutput {
	return o
}

func (o GcmChannelPtrOutput) ToGcmChannelPtrOutputWithContext(ctx context.Context) GcmChannelPtrOutput {
	return o
}

func (o GcmChannelPtrOutput) Elem() GcmChannelOutput {
	return o.ApplyT(func(v *GcmChannel) GcmChannel {
		if v != nil {
			return *v
		}
		var ret GcmChannel
		return ret
	}).(GcmChannelOutput)
}

type GcmChannelArrayOutput struct{ *pulumi.OutputState }

func (GcmChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GcmChannel)(nil))
}

func (o GcmChannelArrayOutput) ToGcmChannelArrayOutput() GcmChannelArrayOutput {
	return o
}

func (o GcmChannelArrayOutput) ToGcmChannelArrayOutputWithContext(ctx context.Context) GcmChannelArrayOutput {
	return o
}

func (o GcmChannelArrayOutput) Index(i pulumi.IntInput) GcmChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GcmChannel {
		return vs[0].([]GcmChannel)[vs[1].(int)]
	}).(GcmChannelOutput)
}

type GcmChannelMapOutput struct{ *pulumi.OutputState }

func (GcmChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GcmChannel)(nil))
}

func (o GcmChannelMapOutput) ToGcmChannelMapOutput() GcmChannelMapOutput {
	return o
}

func (o GcmChannelMapOutput) ToGcmChannelMapOutputWithContext(ctx context.Context) GcmChannelMapOutput {
	return o
}

func (o GcmChannelMapOutput) MapIndex(k pulumi.StringInput) GcmChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GcmChannel {
		return vs[0].(map[string]GcmChannel)[vs[1].(string)]
	}).(GcmChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(GcmChannelOutput{})
	pulumi.RegisterOutputType(GcmChannelPtrOutput{})
	pulumi.RegisterOutputType(GcmChannelArrayOutput{})
	pulumi.RegisterOutputType(GcmChannelMapOutput{})
}
