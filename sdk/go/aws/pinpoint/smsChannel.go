// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint SMS Channel resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewSmsChannel(ctx, "sms", &pinpoint.SmsChannelArgs{
// 			ApplicationId: app.ApplicationId,
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
// Pinpoint SMS Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/smsChannel:SmsChannel sms application-id
// ```
type SmsChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond pulumi.IntOutput `pulumi:"promotionalMessagesPerSecond"`
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrOutput `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrOutput `pulumi:"shortCode"`
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond pulumi.IntOutput `pulumi:"transactionalMessagesPerSecond"`
}

// NewSmsChannel registers a new resource with the given unique name, arguments, and options.
func NewSmsChannel(ctx *pulumi.Context,
	name string, args *SmsChannelArgs, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource SmsChannel
	err := ctx.RegisterResource("aws:pinpoint/smsChannel:SmsChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmsChannel gets an existing SmsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmsChannelState, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	var resource SmsChannel
	err := ctx.ReadResource("aws:pinpoint/smsChannel:SmsChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmsChannel resources.
type smsChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond *int `pulumi:"promotionalMessagesPerSecond"`
	// Sender identifier of your messages.
	SenderId *string `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode *string `pulumi:"shortCode"`
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond *int `pulumi:"transactionalMessagesPerSecond"`
}

type SmsChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Promotional messages per second that can be sent.
	PromotionalMessagesPerSecond pulumi.IntPtrInput
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrInput
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrInput
	// Transactional messages per second that can be sent.
	TransactionalMessagesPerSecond pulumi.IntPtrInput
}

func (SmsChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelState)(nil)).Elem()
}

type smsChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Sender identifier of your messages.
	SenderId *string `pulumi:"senderId"`
	// The Short Code registered with the phone provider.
	ShortCode *string `pulumi:"shortCode"`
}

// The set of arguments for constructing a SmsChannel resource.
type SmsChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Sender identifier of your messages.
	SenderId pulumi.StringPtrInput
	// The Short Code registered with the phone provider.
	ShortCode pulumi.StringPtrInput
}

func (SmsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelArgs)(nil)).Elem()
}

type SmsChannelInput interface {
	pulumi.Input

	ToSmsChannelOutput() SmsChannelOutput
	ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput
}

func (*SmsChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannel)(nil))
}

func (i *SmsChannel) ToSmsChannelOutput() SmsChannelOutput {
	return i.ToSmsChannelOutputWithContext(context.Background())
}

func (i *SmsChannel) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelOutput)
}

func (i *SmsChannel) ToSmsChannelPtrOutput() SmsChannelPtrOutput {
	return i.ToSmsChannelPtrOutputWithContext(context.Background())
}

func (i *SmsChannel) ToSmsChannelPtrOutputWithContext(ctx context.Context) SmsChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPtrOutput)
}

type SmsChannelPtrInput interface {
	pulumi.Input

	ToSmsChannelPtrOutput() SmsChannelPtrOutput
	ToSmsChannelPtrOutputWithContext(ctx context.Context) SmsChannelPtrOutput
}

type smsChannelPtrType SmsChannelArgs

func (*smsChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannel)(nil))
}

func (i *smsChannelPtrType) ToSmsChannelPtrOutput() SmsChannelPtrOutput {
	return i.ToSmsChannelPtrOutputWithContext(context.Background())
}

func (i *smsChannelPtrType) ToSmsChannelPtrOutputWithContext(ctx context.Context) SmsChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPtrOutput)
}

// SmsChannelArrayInput is an input type that accepts SmsChannelArray and SmsChannelArrayOutput values.
// You can construct a concrete instance of `SmsChannelArrayInput` via:
//
//          SmsChannelArray{ SmsChannelArgs{...} }
type SmsChannelArrayInput interface {
	pulumi.Input

	ToSmsChannelArrayOutput() SmsChannelArrayOutput
	ToSmsChannelArrayOutputWithContext(context.Context) SmsChannelArrayOutput
}

type SmsChannelArray []SmsChannelInput

func (SmsChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SmsChannel)(nil))
}

func (i SmsChannelArray) ToSmsChannelArrayOutput() SmsChannelArrayOutput {
	return i.ToSmsChannelArrayOutputWithContext(context.Background())
}

func (i SmsChannelArray) ToSmsChannelArrayOutputWithContext(ctx context.Context) SmsChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelArrayOutput)
}

// SmsChannelMapInput is an input type that accepts SmsChannelMap and SmsChannelMapOutput values.
// You can construct a concrete instance of `SmsChannelMapInput` via:
//
//          SmsChannelMap{ "key": SmsChannelArgs{...} }
type SmsChannelMapInput interface {
	pulumi.Input

	ToSmsChannelMapOutput() SmsChannelMapOutput
	ToSmsChannelMapOutputWithContext(context.Context) SmsChannelMapOutput
}

type SmsChannelMap map[string]SmsChannelInput

func (SmsChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SmsChannel)(nil))
}

func (i SmsChannelMap) ToSmsChannelMapOutput() SmsChannelMapOutput {
	return i.ToSmsChannelMapOutputWithContext(context.Background())
}

func (i SmsChannelMap) ToSmsChannelMapOutputWithContext(ctx context.Context) SmsChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelMapOutput)
}

type SmsChannelOutput struct {
	*pulumi.OutputState
}

func (SmsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannel)(nil))
}

func (o SmsChannelOutput) ToSmsChannelOutput() SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ToSmsChannelPtrOutput() SmsChannelPtrOutput {
	return o.ToSmsChannelPtrOutputWithContext(context.Background())
}

func (o SmsChannelOutput) ToSmsChannelPtrOutputWithContext(ctx context.Context) SmsChannelPtrOutput {
	return o.ApplyT(func(v SmsChannel) *SmsChannel {
		return &v
	}).(SmsChannelPtrOutput)
}

type SmsChannelPtrOutput struct {
	*pulumi.OutputState
}

func (SmsChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannel)(nil))
}

func (o SmsChannelPtrOutput) ToSmsChannelPtrOutput() SmsChannelPtrOutput {
	return o
}

func (o SmsChannelPtrOutput) ToSmsChannelPtrOutputWithContext(ctx context.Context) SmsChannelPtrOutput {
	return o
}

type SmsChannelArrayOutput struct{ *pulumi.OutputState }

func (SmsChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmsChannel)(nil))
}

func (o SmsChannelArrayOutput) ToSmsChannelArrayOutput() SmsChannelArrayOutput {
	return o
}

func (o SmsChannelArrayOutput) ToSmsChannelArrayOutputWithContext(ctx context.Context) SmsChannelArrayOutput {
	return o
}

func (o SmsChannelArrayOutput) Index(i pulumi.IntInput) SmsChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SmsChannel {
		return vs[0].([]SmsChannel)[vs[1].(int)]
	}).(SmsChannelOutput)
}

type SmsChannelMapOutput struct{ *pulumi.OutputState }

func (SmsChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SmsChannel)(nil))
}

func (o SmsChannelMapOutput) ToSmsChannelMapOutput() SmsChannelMapOutput {
	return o
}

func (o SmsChannelMapOutput) ToSmsChannelMapOutputWithContext(ctx context.Context) SmsChannelMapOutput {
	return o
}

func (o SmsChannelMapOutput) MapIndex(k pulumi.StringInput) SmsChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SmsChannel {
		return vs[0].(map[string]SmsChannel)[vs[1].(string)]
	}).(SmsChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(SmsChannelOutput{})
	pulumi.RegisterOutputType(SmsChannelPtrOutput{})
	pulumi.RegisterOutputType(SmsChannelArrayOutput{})
	pulumi.RegisterOutputType(SmsChannelMapOutput{})
}
