// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Sagemaker Image resource.
//
// ## Example Usage
// ### Basic usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sagemaker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sagemaker.NewImage(ctx, "example", &sagemaker.ImageArgs{
// 			ImageName: pulumi.String("example"),
// 			RoleArn:   pulumi.Any(aws_iam_role.Test.Arn),
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
// Sagemaker Code Images can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:sagemaker/image:Image test_image my-code-repo
// ```
type Image struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Image
	err := ctx.RegisterResource("aws:sagemaker/image:Image", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:sagemaker/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn *string `pulumi:"arn"`
	// The description of the image.
	Description *string `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName *string `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName *string `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ImageState struct {
	// The Amazon Resource Name (ARN) assigned by AWS to this Image.
	Arn pulumi.StringPtrInput
	// The description of the image.
	Description pulumi.StringPtrInput
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrInput
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The description of the image.
	Description *string `pulumi:"description"`
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName *string `pulumi:"displayName"`
	// The name of the image. Must be unique to your account.
	ImageName string `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The description of the image.
	Description pulumi.StringPtrInput
	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName pulumi.StringPtrInput
	// The name of the image. Must be unique to your account.
	ImageName pulumi.StringInput
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
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

type ImageOutput struct{ *pulumi.OutputState }

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
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Image) *Image {
		return &v
	}).(ImagePtrOutput)
}

type ImagePtrOutput struct{ *pulumi.OutputState }

func (ImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil))
}

func (o ImagePtrOutput) ToImagePtrOutput() ImagePtrOutput {
	return o
}

func (o ImagePtrOutput) ToImagePtrOutputWithContext(ctx context.Context) ImagePtrOutput {
	return o
}

func (o ImagePtrOutput) Elem() ImageOutput {
	return o.ApplyT(func(v *Image) Image {
		if v != nil {
			return *v
		}
		var ret Image
		return ret
	}).(ImageOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImagePtrInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImagePtrOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
