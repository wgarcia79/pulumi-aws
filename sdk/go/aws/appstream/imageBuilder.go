// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AppStream image builder.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appstream"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appstream.NewImageBuilder(ctx, "testFleet", &appstream.ImageBuilderArgs{
// 			Description:                 pulumi.String("Description of a ImageBuilder"),
// 			DisplayName:                 pulumi.String("Display name of a ImageBuilder"),
// 			EnableDefaultInternetAccess: pulumi.Bool(false),
// 			ImageName:                   pulumi.String("AppStream-WinServer2012R2-07-19-2021"),
// 			InstanceType:                pulumi.String("stream.standard.large"),
// 			VpcConfig: &appstream.ImageBuilderVpcConfigArgs{
// 				SubnetIds: pulumi.StringArray{
// 					pulumi.Any(aws_subnet.Example.Id),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("Example Image Builder"),
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
// `aws_appstream_image_builder` can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:appstream/imageBuilder:ImageBuilder example imageBuilderExample
// ```
type ImageBuilder struct {
	pulumi.CustomResourceState

	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints ImageBuilderAccessEndpointArrayOutput `pulumi:"accessEndpoints"`
	// The version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion pulumi.StringOutput `pulumi:"appstreamAgentVersion"`
	// ARN of the appstream image builder.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Date and time, in UTC and extended RFC 3339 format, when the image builder was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Description to display.
	Description pulumi.StringOutput `pulumi:"description"`
	// Human-readable friendly name for the AppStream image builder.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo ImageBuilderDomainJoinInfoOutput `pulumi:"domainJoinInfo"`
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess pulumi.BoolOutput `pulumi:"enableDefaultInternetAccess"`
	// ARN of the IAM role to apply to the image builder.
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`
	// ARN of the public, private, or shared image to use.
	ImageArn pulumi.StringOutput `pulumi:"imageArn"`
	// Name of the image used to create the image builder.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The instance type to use when launching the image builder.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Unique name for the image builder.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the image builder. Can be: `PENDING`, `UPDATING_AGENT`, `RUNNING`, `STOPPING`, `STOPPED`, `REBOOTING`, `SNAPSHOTTING`, `DELETING`, `FAILED`, `UPDATING`, `PENDING_QUALIFICATION`
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags to assign to the instance. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig ImageBuilderVpcConfigOutput `pulumi:"vpcConfig"`
}

// NewImageBuilder registers a new resource with the given unique name, arguments, and options.
func NewImageBuilder(ctx *pulumi.Context,
	name string, args *ImageBuilderArgs, opts ...pulumi.ResourceOption) (*ImageBuilder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource ImageBuilder
	err := ctx.RegisterResource("aws:appstream/imageBuilder:ImageBuilder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageBuilder gets an existing ImageBuilder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageBuilder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageBuilderState, opts ...pulumi.ResourceOption) (*ImageBuilder, error) {
	var resource ImageBuilder
	err := ctx.ReadResource("aws:appstream/imageBuilder:ImageBuilder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageBuilder resources.
type imageBuilderState struct {
	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints []ImageBuilderAccessEndpoint `pulumi:"accessEndpoints"`
	// The version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion *string `pulumi:"appstreamAgentVersion"`
	// ARN of the appstream image builder.
	Arn *string `pulumi:"arn"`
	// Date and time, in UTC and extended RFC 3339 format, when the image builder was created.
	CreatedTime *string `pulumi:"createdTime"`
	// Description to display.
	Description *string `pulumi:"description"`
	// Human-readable friendly name for the AppStream image builder.
	DisplayName *string `pulumi:"displayName"`
	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo *ImageBuilderDomainJoinInfo `pulumi:"domainJoinInfo"`
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool `pulumi:"enableDefaultInternetAccess"`
	// ARN of the IAM role to apply to the image builder.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// ARN of the public, private, or shared image to use.
	ImageArn *string `pulumi:"imageArn"`
	// Name of the image used to create the image builder.
	ImageName *string `pulumi:"imageName"`
	// The instance type to use when launching the image builder.
	InstanceType *string `pulumi:"instanceType"`
	// Unique name for the image builder.
	Name *string `pulumi:"name"`
	// State of the image builder. Can be: `PENDING`, `UPDATING_AGENT`, `RUNNING`, `STOPPING`, `STOPPED`, `REBOOTING`, `SNAPSHOTTING`, `DELETING`, `FAILED`, `UPDATING`, `PENDING_QUALIFICATION`
	State *string `pulumi:"state"`
	// A map of tags to assign to the instance. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig *ImageBuilderVpcConfig `pulumi:"vpcConfig"`
}

type ImageBuilderState struct {
	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints ImageBuilderAccessEndpointArrayInput
	// The version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion pulumi.StringPtrInput
	// ARN of the appstream image builder.
	Arn pulumi.StringPtrInput
	// Date and time, in UTC and extended RFC 3339 format, when the image builder was created.
	CreatedTime pulumi.StringPtrInput
	// Description to display.
	Description pulumi.StringPtrInput
	// Human-readable friendly name for the AppStream image builder.
	DisplayName pulumi.StringPtrInput
	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo ImageBuilderDomainJoinInfoPtrInput
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess pulumi.BoolPtrInput
	// ARN of the IAM role to apply to the image builder.
	IamRoleArn pulumi.StringPtrInput
	// ARN of the public, private, or shared image to use.
	ImageArn pulumi.StringPtrInput
	// Name of the image used to create the image builder.
	ImageName pulumi.StringPtrInput
	// The instance type to use when launching the image builder.
	InstanceType pulumi.StringPtrInput
	// Unique name for the image builder.
	Name pulumi.StringPtrInput
	// State of the image builder. Can be: `PENDING`, `UPDATING_AGENT`, `RUNNING`, `STOPPING`, `STOPPED`, `REBOOTING`, `SNAPSHOTTING`, `DELETING`, `FAILED`, `UPDATING`, `PENDING_QUALIFICATION`
	State pulumi.StringPtrInput
	// A map of tags to assign to the instance. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig ImageBuilderVpcConfigPtrInput
}

func (ImageBuilderState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageBuilderState)(nil)).Elem()
}

type imageBuilderArgs struct {
	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints []ImageBuilderAccessEndpoint `pulumi:"accessEndpoints"`
	// The version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion *string `pulumi:"appstreamAgentVersion"`
	// Description to display.
	Description *string `pulumi:"description"`
	// Human-readable friendly name for the AppStream image builder.
	DisplayName *string `pulumi:"displayName"`
	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo *ImageBuilderDomainJoinInfo `pulumi:"domainJoinInfo"`
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool `pulumi:"enableDefaultInternetAccess"`
	// ARN of the IAM role to apply to the image builder.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// ARN of the public, private, or shared image to use.
	ImageArn *string `pulumi:"imageArn"`
	// Name of the image used to create the image builder.
	ImageName *string `pulumi:"imageName"`
	// The instance type to use when launching the image builder.
	InstanceType string `pulumi:"instanceType"`
	// Unique name for the image builder.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the instance. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig *ImageBuilderVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a ImageBuilder resource.
type ImageBuilderArgs struct {
	// Set of interface VPC endpoint (interface endpoint) objects. Maximum of 4. See below.
	AccessEndpoints ImageBuilderAccessEndpointArrayInput
	// The version of the AppStream 2.0 agent to use for this image builder.
	AppstreamAgentVersion pulumi.StringPtrInput
	// Description to display.
	Description pulumi.StringPtrInput
	// Human-readable friendly name for the AppStream image builder.
	DisplayName pulumi.StringPtrInput
	// Configuration block for the name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain. See below.
	DomainJoinInfo ImageBuilderDomainJoinInfoPtrInput
	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess pulumi.BoolPtrInput
	// ARN of the IAM role to apply to the image builder.
	IamRoleArn pulumi.StringPtrInput
	// ARN of the public, private, or shared image to use.
	ImageArn pulumi.StringPtrInput
	// Name of the image used to create the image builder.
	ImageName pulumi.StringPtrInput
	// The instance type to use when launching the image builder.
	InstanceType pulumi.StringInput
	// Unique name for the image builder.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the instance. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
	// Configuration block for the VPC configuration for the image builder. See below.
	VpcConfig ImageBuilderVpcConfigPtrInput
}

func (ImageBuilderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageBuilderArgs)(nil)).Elem()
}

type ImageBuilderInput interface {
	pulumi.Input

	ToImageBuilderOutput() ImageBuilderOutput
	ToImageBuilderOutputWithContext(ctx context.Context) ImageBuilderOutput
}

func (*ImageBuilder) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageBuilder)(nil))
}

func (i *ImageBuilder) ToImageBuilderOutput() ImageBuilderOutput {
	return i.ToImageBuilderOutputWithContext(context.Background())
}

func (i *ImageBuilder) ToImageBuilderOutputWithContext(ctx context.Context) ImageBuilderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageBuilderOutput)
}

func (i *ImageBuilder) ToImageBuilderPtrOutput() ImageBuilderPtrOutput {
	return i.ToImageBuilderPtrOutputWithContext(context.Background())
}

func (i *ImageBuilder) ToImageBuilderPtrOutputWithContext(ctx context.Context) ImageBuilderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageBuilderPtrOutput)
}

type ImageBuilderPtrInput interface {
	pulumi.Input

	ToImageBuilderPtrOutput() ImageBuilderPtrOutput
	ToImageBuilderPtrOutputWithContext(ctx context.Context) ImageBuilderPtrOutput
}

type imageBuilderPtrType ImageBuilderArgs

func (*imageBuilderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageBuilder)(nil))
}

func (i *imageBuilderPtrType) ToImageBuilderPtrOutput() ImageBuilderPtrOutput {
	return i.ToImageBuilderPtrOutputWithContext(context.Background())
}

func (i *imageBuilderPtrType) ToImageBuilderPtrOutputWithContext(ctx context.Context) ImageBuilderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageBuilderPtrOutput)
}

// ImageBuilderArrayInput is an input type that accepts ImageBuilderArray and ImageBuilderArrayOutput values.
// You can construct a concrete instance of `ImageBuilderArrayInput` via:
//
//          ImageBuilderArray{ ImageBuilderArgs{...} }
type ImageBuilderArrayInput interface {
	pulumi.Input

	ToImageBuilderArrayOutput() ImageBuilderArrayOutput
	ToImageBuilderArrayOutputWithContext(context.Context) ImageBuilderArrayOutput
}

type ImageBuilderArray []ImageBuilderInput

func (ImageBuilderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageBuilder)(nil)).Elem()
}

func (i ImageBuilderArray) ToImageBuilderArrayOutput() ImageBuilderArrayOutput {
	return i.ToImageBuilderArrayOutputWithContext(context.Background())
}

func (i ImageBuilderArray) ToImageBuilderArrayOutputWithContext(ctx context.Context) ImageBuilderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageBuilderArrayOutput)
}

// ImageBuilderMapInput is an input type that accepts ImageBuilderMap and ImageBuilderMapOutput values.
// You can construct a concrete instance of `ImageBuilderMapInput` via:
//
//          ImageBuilderMap{ "key": ImageBuilderArgs{...} }
type ImageBuilderMapInput interface {
	pulumi.Input

	ToImageBuilderMapOutput() ImageBuilderMapOutput
	ToImageBuilderMapOutputWithContext(context.Context) ImageBuilderMapOutput
}

type ImageBuilderMap map[string]ImageBuilderInput

func (ImageBuilderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageBuilder)(nil)).Elem()
}

func (i ImageBuilderMap) ToImageBuilderMapOutput() ImageBuilderMapOutput {
	return i.ToImageBuilderMapOutputWithContext(context.Background())
}

func (i ImageBuilderMap) ToImageBuilderMapOutputWithContext(ctx context.Context) ImageBuilderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageBuilderMapOutput)
}

type ImageBuilderOutput struct{ *pulumi.OutputState }

func (ImageBuilderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageBuilder)(nil))
}

func (o ImageBuilderOutput) ToImageBuilderOutput() ImageBuilderOutput {
	return o
}

func (o ImageBuilderOutput) ToImageBuilderOutputWithContext(ctx context.Context) ImageBuilderOutput {
	return o
}

func (o ImageBuilderOutput) ToImageBuilderPtrOutput() ImageBuilderPtrOutput {
	return o.ToImageBuilderPtrOutputWithContext(context.Background())
}

func (o ImageBuilderOutput) ToImageBuilderPtrOutputWithContext(ctx context.Context) ImageBuilderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageBuilder) *ImageBuilder {
		return &v
	}).(ImageBuilderPtrOutput)
}

type ImageBuilderPtrOutput struct{ *pulumi.OutputState }

func (ImageBuilderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageBuilder)(nil))
}

func (o ImageBuilderPtrOutput) ToImageBuilderPtrOutput() ImageBuilderPtrOutput {
	return o
}

func (o ImageBuilderPtrOutput) ToImageBuilderPtrOutputWithContext(ctx context.Context) ImageBuilderPtrOutput {
	return o
}

func (o ImageBuilderPtrOutput) Elem() ImageBuilderOutput {
	return o.ApplyT(func(v *ImageBuilder) ImageBuilder {
		if v != nil {
			return *v
		}
		var ret ImageBuilder
		return ret
	}).(ImageBuilderOutput)
}

type ImageBuilderArrayOutput struct{ *pulumi.OutputState }

func (ImageBuilderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageBuilder)(nil))
}

func (o ImageBuilderArrayOutput) ToImageBuilderArrayOutput() ImageBuilderArrayOutput {
	return o
}

func (o ImageBuilderArrayOutput) ToImageBuilderArrayOutputWithContext(ctx context.Context) ImageBuilderArrayOutput {
	return o
}

func (o ImageBuilderArrayOutput) Index(i pulumi.IntInput) ImageBuilderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageBuilder {
		return vs[0].([]ImageBuilder)[vs[1].(int)]
	}).(ImageBuilderOutput)
}

type ImageBuilderMapOutput struct{ *pulumi.OutputState }

func (ImageBuilderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageBuilder)(nil))
}

func (o ImageBuilderMapOutput) ToImageBuilderMapOutput() ImageBuilderMapOutput {
	return o
}

func (o ImageBuilderMapOutput) ToImageBuilderMapOutputWithContext(ctx context.Context) ImageBuilderMapOutput {
	return o
}

func (o ImageBuilderMapOutput) MapIndex(k pulumi.StringInput) ImageBuilderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageBuilder {
		return vs[0].(map[string]ImageBuilder)[vs[1].(string)]
	}).(ImageBuilderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageBuilderInput)(nil)).Elem(), &ImageBuilder{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageBuilderPtrInput)(nil)).Elem(), &ImageBuilder{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageBuilderArrayInput)(nil)).Elem(), ImageBuilderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageBuilderMapInput)(nil)).Elem(), ImageBuilderMap{})
	pulumi.RegisterOutputType(ImageBuilderOutput{})
	pulumi.RegisterOutputType(ImageBuilderPtrOutput{})
	pulumi.RegisterOutputType(ImageBuilderArrayOutput{})
	pulumi.RegisterOutputType(ImageBuilderMapOutput{})
}
