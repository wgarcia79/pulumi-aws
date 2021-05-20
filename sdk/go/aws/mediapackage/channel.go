// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Elemental MediaPackage Channel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/mediapackage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mediapackage.NewChannel(ctx, "kittens", &mediapackage.ChannelArgs{
// 			ChannelId:   pulumi.String("kitten-channel"),
// 			Description: pulumi.String("A channel dedicated to amusing videos of kittens."),
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
// Media Package Channels can be imported via the channel ID, e.g.
//
// ```sh
//  $ pulumi import aws:mediapackage/channel:Channel kittens kittens-channel
// ```
type Channel struct {
	pulumi.CustomResourceState

	// The ARN of the channel
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A unique identifier describing the channel
	ChannelId pulumi.StringOutput `pulumi:"channelId"`
	// A description of the channel
	Description pulumi.StringOutput `pulumi:"description"`
	// A single item list of HLS ingest information
	HlsIngests ChannelHlsIngestArrayOutput `pulumi:"hlsIngests"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource Channel
	err := ctx.RegisterResource("aws:mediapackage/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("aws:mediapackage/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// The ARN of the channel
	Arn *string `pulumi:"arn"`
	// A unique identifier describing the channel
	ChannelId *string `pulumi:"channelId"`
	// A description of the channel
	Description *string `pulumi:"description"`
	// A single item list of HLS ingest information
	HlsIngests []ChannelHlsIngest `pulumi:"hlsIngests"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ChannelState struct {
	// The ARN of the channel
	Arn pulumi.StringPtrInput
	// A unique identifier describing the channel
	ChannelId pulumi.StringPtrInput
	// A description of the channel
	Description pulumi.StringPtrInput
	// A single item list of HLS ingest information
	HlsIngests ChannelHlsIngestArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// A unique identifier describing the channel
	ChannelId string `pulumi:"channelId"`
	// A description of the channel
	Description *string `pulumi:"description"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// A unique identifier describing the channel
	ChannelId pulumi.StringInput
	// A description of the channel
	Description pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((*Channel)(nil))
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

func (i *Channel) ToChannelPtrOutput() ChannelPtrOutput {
	return i.ToChannelPtrOutputWithContext(context.Background())
}

func (i *Channel) ToChannelPtrOutputWithContext(ctx context.Context) ChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelPtrOutput)
}

type ChannelPtrInput interface {
	pulumi.Input

	ToChannelPtrOutput() ChannelPtrOutput
	ToChannelPtrOutputWithContext(ctx context.Context) ChannelPtrOutput
}

type channelPtrType ChannelArgs

func (*channelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil))
}

func (i *channelPtrType) ToChannelPtrOutput() ChannelPtrOutput {
	return i.ToChannelPtrOutputWithContext(context.Background())
}

func (i *channelPtrType) ToChannelPtrOutputWithContext(ctx context.Context) ChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelPtrOutput)
}

// ChannelArrayInput is an input type that accepts ChannelArray and ChannelArrayOutput values.
// You can construct a concrete instance of `ChannelArrayInput` via:
//
//          ChannelArray{ ChannelArgs{...} }
type ChannelArrayInput interface {
	pulumi.Input

	ToChannelArrayOutput() ChannelArrayOutput
	ToChannelArrayOutputWithContext(context.Context) ChannelArrayOutput
}

type ChannelArray []ChannelInput

func (ChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Channel)(nil))
}

func (i ChannelArray) ToChannelArrayOutput() ChannelArrayOutput {
	return i.ToChannelArrayOutputWithContext(context.Background())
}

func (i ChannelArray) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelArrayOutput)
}

// ChannelMapInput is an input type that accepts ChannelMap and ChannelMapOutput values.
// You can construct a concrete instance of `ChannelMapInput` via:
//
//          ChannelMap{ "key": ChannelArgs{...} }
type ChannelMapInput interface {
	pulumi.Input

	ToChannelMapOutput() ChannelMapOutput
	ToChannelMapOutputWithContext(context.Context) ChannelMapOutput
}

type ChannelMap map[string]ChannelInput

func (ChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Channel)(nil))
}

func (i ChannelMap) ToChannelMapOutput() ChannelMapOutput {
	return i.ToChannelMapOutputWithContext(context.Background())
}

func (i ChannelMap) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelMapOutput)
}

type ChannelOutput struct {
	*pulumi.OutputState
}

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Channel)(nil))
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelPtrOutput() ChannelPtrOutput {
	return o.ToChannelPtrOutputWithContext(context.Background())
}

func (o ChannelOutput) ToChannelPtrOutputWithContext(ctx context.Context) ChannelPtrOutput {
	return o.ApplyT(func(v Channel) *Channel {
		return &v
	}).(ChannelPtrOutput)
}

type ChannelPtrOutput struct {
	*pulumi.OutputState
}

func (ChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil))
}

func (o ChannelPtrOutput) ToChannelPtrOutput() ChannelPtrOutput {
	return o
}

func (o ChannelPtrOutput) ToChannelPtrOutputWithContext(ctx context.Context) ChannelPtrOutput {
	return o
}

type ChannelArrayOutput struct{ *pulumi.OutputState }

func (ChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Channel)(nil))
}

func (o ChannelArrayOutput) ToChannelArrayOutput() ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) Index(i pulumi.IntInput) ChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Channel {
		return vs[0].([]Channel)[vs[1].(int)]
	}).(ChannelOutput)
}

type ChannelMapOutput struct{ *pulumi.OutputState }

func (ChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Channel)(nil))
}

func (o ChannelMapOutput) ToChannelMapOutput() ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) MapIndex(k pulumi.StringInput) ChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Channel {
		return vs[0].(map[string]Channel)[vs[1].(string)]
	}).(ChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
	pulumi.RegisterOutputType(ChannelPtrOutput{})
	pulumi.RegisterOutputType(ChannelArrayOutput{})
	pulumi.RegisterOutputType(ChannelMapOutput{})
}
