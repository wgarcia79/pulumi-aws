// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elastictranscoder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic Transcoder preset resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elastictranscoder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elastictranscoder.NewPreset(ctx, "bar", &elastictranscoder.PresetArgs{
// 			Audio: &elastictranscoder.PresetAudioArgs{
// 				AudioPackingMode: pulumi.String("SingleTrack"),
// 				BitRate:          pulumi.String("96"),
// 				Channels:         pulumi.String("2"),
// 				Codec:            pulumi.String("AAC"),
// 				SampleRate:       pulumi.String("44100"),
// 			},
// 			AudioCodecOptions: &elastictranscoder.PresetAudioCodecOptionsArgs{
// 				Profile: pulumi.String("AAC-LC"),
// 			},
// 			Container:   pulumi.String("mp4"),
// 			Description: pulumi.String("Sample Preset"),
// 			Thumbnails: &elastictranscoder.PresetThumbnailsArgs{
// 				Format:        pulumi.String("png"),
// 				Interval:      pulumi.String("120"),
// 				MaxHeight:     pulumi.String("auto"),
// 				MaxWidth:      pulumi.String("auto"),
// 				PaddingPolicy: pulumi.String("Pad"),
// 				SizingPolicy:  pulumi.String("Fit"),
// 			},
// 			Video: &elastictranscoder.PresetVideoArgs{
// 				BitRate:            pulumi.String("1600"),
// 				Codec:              pulumi.String("H.264"),
// 				DisplayAspectRatio: pulumi.String("16:9"),
// 				FixedGop:           pulumi.String("false"),
// 				FrameRate:          pulumi.String("auto"),
// 				KeyframesMaxDist:   pulumi.String("240"),
// 				MaxFrameRate:       pulumi.String("60"),
// 				MaxHeight:          pulumi.String("auto"),
// 				MaxWidth:           pulumi.String("auto"),
// 				PaddingPolicy:      pulumi.String("Pad"),
// 				SizingPolicy:       pulumi.String("Fit"),
// 			},
// 			VideoCodecOptions: pulumi.StringMap{
// 				"ColorSpaceConversionMode": pulumi.String("None"),
// 				"InterlacedMode":           pulumi.String("Progressive"),
// 				"Level":                    pulumi.String("2.2"),
// 				"MaxReferenceFrames":       pulumi.String("3"),
// 				"Profile":                  pulumi.String("main"),
// 			},
// 			VideoWatermarks: elastictranscoder.PresetVideoWatermarkArray{
// 				&elastictranscoder.PresetVideoWatermarkArgs{
// 					HorizontalAlign:  pulumi.String("Right"),
// 					HorizontalOffset: pulumi.String("10px"),
// 					Id:               pulumi.String("Test"),
// 					MaxHeight:        pulumi.String(fmt.Sprintf("%v%v", "20", "%")),
// 					MaxWidth:         pulumi.String(fmt.Sprintf("%v%v", "20", "%")),
// 					Opacity:          pulumi.String("55.5"),
// 					SizingPolicy:     pulumi.String("ShrinkToFit"),
// 					Target:           pulumi.String("Content"),
// 					VerticalAlign:    pulumi.String("Bottom"),
// 					VerticalOffset:   pulumi.String("10px"),
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
// Elastic Transcoder presets can be imported using the `id`, e.g.,
//
// ```sh
//  $ pulumi import aws:elastictranscoder/preset:Preset basic_preset 1407981661351-cttk8b
// ```
type Preset struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Elastic Transcoder Preset.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Audio parameters object (documented below).
	Audio PresetAudioPtrOutput `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrOutput `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringOutput `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringOutput `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrOutput `pulumi:"thumbnails"`
	Type       pulumi.StringOutput       `pulumi:"type"`
	// Video parameters object (documented below)
	Video PresetVideoPtrOutput `pulumi:"video"`
	// Codec options for the video parameters
	VideoCodecOptions pulumi.StringMapOutput `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks PresetVideoWatermarkArrayOutput `pulumi:"videoWatermarks"`
}

// NewPreset registers a new resource with the given unique name, arguments, and options.
func NewPreset(ctx *pulumi.Context,
	name string, args *PresetArgs, opts ...pulumi.ResourceOption) (*Preset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Container == nil {
		return nil, errors.New("invalid value for required argument 'Container'")
	}
	var resource Preset
	err := ctx.RegisterResource("aws:elastictranscoder/preset:Preset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreset gets an existing Preset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PresetState, opts ...pulumi.ResourceOption) (*Preset, error) {
	var resource Preset
	err := ctx.ReadResource("aws:elastictranscoder/preset:Preset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Preset resources.
type presetState struct {
	// Amazon Resource Name (ARN) of the Elastic Transcoder Preset.
	Arn *string `pulumi:"arn"`
	// Audio parameters object (documented below).
	Audio *PresetAudio `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions *PresetAudioCodecOptions `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container *string `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description *string `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name *string `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails *PresetThumbnails `pulumi:"thumbnails"`
	Type       *string           `pulumi:"type"`
	// Video parameters object (documented below)
	Video *PresetVideo `pulumi:"video"`
	// Codec options for the video parameters
	VideoCodecOptions map[string]string `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks []PresetVideoWatermark `pulumi:"videoWatermarks"`
}

type PresetState struct {
	// Amazon Resource Name (ARN) of the Elastic Transcoder Preset.
	Arn pulumi.StringPtrInput
	// Audio parameters object (documented below).
	Audio PresetAudioPtrInput
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrInput
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringPtrInput
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrInput
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringPtrInput
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrInput
	Type       pulumi.StringPtrInput
	// Video parameters object (documented below)
	Video PresetVideoPtrInput
	// Codec options for the video parameters
	VideoCodecOptions pulumi.StringMapInput
	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks PresetVideoWatermarkArrayInput
}

func (PresetState) ElementType() reflect.Type {
	return reflect.TypeOf((*presetState)(nil)).Elem()
}

type presetArgs struct {
	// Audio parameters object (documented below).
	Audio *PresetAudio `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions *PresetAudioCodecOptions `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container string `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description *string `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name *string `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails *PresetThumbnails `pulumi:"thumbnails"`
	Type       *string           `pulumi:"type"`
	// Video parameters object (documented below)
	Video *PresetVideo `pulumi:"video"`
	// Codec options for the video parameters
	VideoCodecOptions map[string]string `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks []PresetVideoWatermark `pulumi:"videoWatermarks"`
}

// The set of arguments for constructing a Preset resource.
type PresetArgs struct {
	// Audio parameters object (documented below).
	Audio PresetAudioPtrInput
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrInput
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringInput
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrInput
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringPtrInput
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrInput
	Type       pulumi.StringPtrInput
	// Video parameters object (documented below)
	Video PresetVideoPtrInput
	// Codec options for the video parameters
	VideoCodecOptions pulumi.StringMapInput
	// Watermark parameters for the video parameters (documented below)
	VideoWatermarks PresetVideoWatermarkArrayInput
}

func (PresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*presetArgs)(nil)).Elem()
}

type PresetInput interface {
	pulumi.Input

	ToPresetOutput() PresetOutput
	ToPresetOutputWithContext(ctx context.Context) PresetOutput
}

func (*Preset) ElementType() reflect.Type {
	return reflect.TypeOf((*Preset)(nil))
}

func (i *Preset) ToPresetOutput() PresetOutput {
	return i.ToPresetOutputWithContext(context.Background())
}

func (i *Preset) ToPresetOutputWithContext(ctx context.Context) PresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetOutput)
}

func (i *Preset) ToPresetPtrOutput() PresetPtrOutput {
	return i.ToPresetPtrOutputWithContext(context.Background())
}

func (i *Preset) ToPresetPtrOutputWithContext(ctx context.Context) PresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetPtrOutput)
}

type PresetPtrInput interface {
	pulumi.Input

	ToPresetPtrOutput() PresetPtrOutput
	ToPresetPtrOutputWithContext(ctx context.Context) PresetPtrOutput
}

type presetPtrType PresetArgs

func (*presetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Preset)(nil))
}

func (i *presetPtrType) ToPresetPtrOutput() PresetPtrOutput {
	return i.ToPresetPtrOutputWithContext(context.Background())
}

func (i *presetPtrType) ToPresetPtrOutputWithContext(ctx context.Context) PresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetPtrOutput)
}

// PresetArrayInput is an input type that accepts PresetArray and PresetArrayOutput values.
// You can construct a concrete instance of `PresetArrayInput` via:
//
//          PresetArray{ PresetArgs{...} }
type PresetArrayInput interface {
	pulumi.Input

	ToPresetArrayOutput() PresetArrayOutput
	ToPresetArrayOutputWithContext(context.Context) PresetArrayOutput
}

type PresetArray []PresetInput

func (PresetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Preset)(nil)).Elem()
}

func (i PresetArray) ToPresetArrayOutput() PresetArrayOutput {
	return i.ToPresetArrayOutputWithContext(context.Background())
}

func (i PresetArray) ToPresetArrayOutputWithContext(ctx context.Context) PresetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetArrayOutput)
}

// PresetMapInput is an input type that accepts PresetMap and PresetMapOutput values.
// You can construct a concrete instance of `PresetMapInput` via:
//
//          PresetMap{ "key": PresetArgs{...} }
type PresetMapInput interface {
	pulumi.Input

	ToPresetMapOutput() PresetMapOutput
	ToPresetMapOutputWithContext(context.Context) PresetMapOutput
}

type PresetMap map[string]PresetInput

func (PresetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Preset)(nil)).Elem()
}

func (i PresetMap) ToPresetMapOutput() PresetMapOutput {
	return i.ToPresetMapOutputWithContext(context.Background())
}

func (i PresetMap) ToPresetMapOutputWithContext(ctx context.Context) PresetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PresetMapOutput)
}

type PresetOutput struct{ *pulumi.OutputState }

func (PresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Preset)(nil))
}

func (o PresetOutput) ToPresetOutput() PresetOutput {
	return o
}

func (o PresetOutput) ToPresetOutputWithContext(ctx context.Context) PresetOutput {
	return o
}

func (o PresetOutput) ToPresetPtrOutput() PresetPtrOutput {
	return o.ToPresetPtrOutputWithContext(context.Background())
}

func (o PresetOutput) ToPresetPtrOutputWithContext(ctx context.Context) PresetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Preset) *Preset {
		return &v
	}).(PresetPtrOutput)
}

type PresetPtrOutput struct{ *pulumi.OutputState }

func (PresetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Preset)(nil))
}

func (o PresetPtrOutput) ToPresetPtrOutput() PresetPtrOutput {
	return o
}

func (o PresetPtrOutput) ToPresetPtrOutputWithContext(ctx context.Context) PresetPtrOutput {
	return o
}

func (o PresetPtrOutput) Elem() PresetOutput {
	return o.ApplyT(func(v *Preset) Preset {
		if v != nil {
			return *v
		}
		var ret Preset
		return ret
	}).(PresetOutput)
}

type PresetArrayOutput struct{ *pulumi.OutputState }

func (PresetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Preset)(nil))
}

func (o PresetArrayOutput) ToPresetArrayOutput() PresetArrayOutput {
	return o
}

func (o PresetArrayOutput) ToPresetArrayOutputWithContext(ctx context.Context) PresetArrayOutput {
	return o
}

func (o PresetArrayOutput) Index(i pulumi.IntInput) PresetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Preset {
		return vs[0].([]Preset)[vs[1].(int)]
	}).(PresetOutput)
}

type PresetMapOutput struct{ *pulumi.OutputState }

func (PresetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Preset)(nil))
}

func (o PresetMapOutput) ToPresetMapOutput() PresetMapOutput {
	return o
}

func (o PresetMapOutput) ToPresetMapOutputWithContext(ctx context.Context) PresetMapOutput {
	return o
}

func (o PresetMapOutput) MapIndex(k pulumi.StringInput) PresetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Preset {
		return vs[0].(map[string]Preset)[vs[1].(string)]
	}).(PresetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PresetInput)(nil)).Elem(), &Preset{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresetPtrInput)(nil)).Elem(), &Preset{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresetArrayInput)(nil)).Elem(), PresetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PresetMapInput)(nil)).Elem(), PresetMap{})
	pulumi.RegisterOutputType(PresetOutput{})
	pulumi.RegisterOutputType(PresetPtrOutput{})
	pulumi.RegisterOutputType(PresetArrayOutput{})
	pulumi.RegisterOutputType(PresetMapOutput{})
}
