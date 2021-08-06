// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image Builder Image.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := imagebuilder.NewImage(ctx, "example", &imagebuilder.ImageArgs{
// 			DistributionConfigurationArn:   pulumi.Any(aws_imagebuilder_distribution_configuration.Example.Arn),
// 			ImageRecipeArn:                 pulumi.Any(aws_imagebuilder_image_recipe.Example.Arn),
// 			InfrastructureConfigurationArn: pulumi.Any(aws_imagebuilder_infrastructure_configuration.Example.Arn),
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
// `aws_imagebuilder_image` resources can be imported using the Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:imagebuilder/image:Image example arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
// ```
type Image struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the image.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Date the image was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrOutput `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrOutput `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn pulumi.StringOutput `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImageImageTestsConfigurationOutput `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringOutput `pulumi:"infrastructureConfigurationArn"`
	// Name of the AMI.
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating System version of the image.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// List of objects with resources created by the image.
	OutputResources ImageOutputResourceArrayOutput `pulumi:"outputResources"`
	// Platform of the image.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Key-value map of resource tags for the Image Builder Image. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Version of the image.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageRecipeArn == nil {
		return nil, errors.New("invalid value for required argument 'ImageRecipeArn'")
	}
	if args.InfrastructureConfigurationArn == nil {
		return nil, errors.New("invalid value for required argument 'InfrastructureConfigurationArn'")
	}
	var resource Image
	err := ctx.RegisterResource("aws:imagebuilder/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("aws:imagebuilder/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// Amazon Resource Name (ARN) of the image.
	Arn *string `pulumi:"arn"`
	// Date the image was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn *string `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration *ImageImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn *string `pulumi:"infrastructureConfigurationArn"`
	// Name of the AMI.
	Name *string `pulumi:"name"`
	// Operating System version of the image.
	OsVersion *string `pulumi:"osVersion"`
	// List of objects with resources created by the image.
	OutputResources []ImageOutputResource `pulumi:"outputResources"`
	// Platform of the image.
	Platform *string `pulumi:"platform"`
	// Key-value map of resource tags for the Image Builder Image. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Version of the image.
	Version *string `pulumi:"version"`
}

type ImageState struct {
	// Amazon Resource Name (ARN) of the image.
	Arn pulumi.StringPtrInput
	// Date the image was created.
	DateCreated pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrInput
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn pulumi.StringPtrInput
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImageImageTestsConfigurationPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringPtrInput
	// Name of the AMI.
	Name pulumi.StringPtrInput
	// Operating System version of the image.
	OsVersion pulumi.StringPtrInput
	// List of objects with resources created by the image.
	OutputResources ImageOutputResourceArrayInput
	// Platform of the image.
	Platform pulumi.StringPtrInput
	// Key-value map of resource tags for the Image Builder Image. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Version of the image.
	Version pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn *string `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled *bool `pulumi:"enhancedImageMetadataEnabled"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn string `pulumi:"imageRecipeArn"`
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration *ImageImageTestsConfiguration `pulumi:"imageTestsConfiguration"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn string `pulumi:"infrastructureConfigurationArn"`
	// Key-value map of resource tags for the Image Builder Image. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn pulumi.StringPtrInput
	// Whether additional information about the image being created is collected. Defaults to `true`.
	EnhancedImageMetadataEnabled pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn pulumi.StringInput
	// Configuration block with image tests configuration. Detailed below.
	ImageTestsConfiguration ImageImageTestsConfigurationPtrInput
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn pulumi.StringInput
	// Key-value map of resource tags for the Image Builder Image. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

func (i *Image) ToImagePtrOutput() ImagePtrOutput {
	return i.ToImagePtrOutputWithContext(context.Background())
}

func (i *Image) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePtrOutput)
}

type ImagePtrInput interface {
	pulumi.Input

	ToImagePtrOutput() ImagePtrOutput
	ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput
}

type imagePtrType ImageArgs

func (*imagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil))
}

func (i *imagePtrType) ToImagePtrOutput() ImagePtrOutput {
	return i.ToImagePtrOutputWithContext(context.Background())
}

func (i *imagePtrType) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePtrOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//          ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//          ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct {
	*pulumi.OutputState
}

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) ToImagePtrOutput() ImagePtrOutput {
	return o.ToImagePtrOutputWithContext(context.Background())
}

func (o ImageOutput) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return o.ApplyT(func(v Image) *Image {
		return &v
	}).(ImagePtrOutput)
}

type ImagePtrOutput struct {
	*pulumi.OutputState
}

func (ImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil))
}

func (o ImagePtrOutput) ToImagePtrOutput() ImagePtrOutput {
	return o
}

func (o ImagePtrOutput) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return o
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Image)(nil))
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Image {
		return vs[0].([]Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Image)(nil))
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Image {
		return vs[0].(map[string]Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImagePtrOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
