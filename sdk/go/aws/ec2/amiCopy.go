// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
// including cross-region copies.
//
// If the source AMI has associated EBS snapshots, those will also be duplicated
// along with the AMI.
//
// This is useful for taking a single AMI provisioned in one region and making
// it available in another for a multi-region deployment.
//
// Copying an AMI can take several minutes. The creation of this resource will
// block until the new AMI is available for use on new instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewAmiCopy(ctx, "example", &ec2.AmiCopyArgs{
// 			Description:     pulumi.String("A copy of ami-xxxxxxxx"),
// 			SourceAmiId:     pulumi.String("ami-xxxxxxxx"),
// 			SourceAmiRegion: pulumi.String("us-west-1"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AmiCopy struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// The ARN of the AMI.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn pulumi.StringPtrOutput `pulumi:"destinationOutpostArn"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiCopyEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolOutput `pulumi:"enaSupport"`
	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiCopyEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	Hypervisor            pulumi.StringOutput                    `pulumi:"hypervisor"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation   pulumi.StringOutput `pulumi:"imageLocation"`
	ImageOwnerAlias pulumi.StringOutput `pulumi:"imageOwnerAlias"`
	ImageType       pulumi.StringOutput `pulumi:"imageType"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringOutput `pulumi:"kernelId"`
	// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
	// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
	// if this parameter is not specified, the default CMK for EBS is used
	KmsKeyId           pulumi.StringOutput `pulumi:"kmsKeyId"`
	ManageEbsSnapshots pulumi.BoolOutput   `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name            pulumi.StringOutput `pulumi:"name"`
	OwnerId         pulumi.StringOutput `pulumi:"ownerId"`
	Platform        pulumi.StringOutput `pulumi:"platform"`
	PlatformDetails pulumi.StringOutput `pulumi:"platformDetails"`
	Public          pulumi.BoolOutput   `pulumi:"public"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringOutput `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringOutput `pulumi:"rootDeviceName"`
	RootSnapshotId pulumi.StringOutput `pulumi:"rootSnapshotId"`
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId pulumi.StringOutput `pulumi:"sourceAmiId"`
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion pulumi.StringOutput `pulumi:"sourceAmiRegion"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringOutput `pulumi:"sriovNetSupport"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll        pulumi.StringMapOutput `pulumi:"tagsAll"`
	UsageOperation pulumi.StringOutput    `pulumi:"usageOperation"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringOutput `pulumi:"virtualizationType"`
}

// NewAmiCopy registers a new resource with the given unique name, arguments, and options.
func NewAmiCopy(ctx *pulumi.Context,
	name string, args *AmiCopyArgs, opts ...pulumi.ResourceOption) (*AmiCopy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceAmiId == nil {
		return nil, errors.New("invalid value for required argument 'SourceAmiId'")
	}
	if args.SourceAmiRegion == nil {
		return nil, errors.New("invalid value for required argument 'SourceAmiRegion'")
	}
	var resource AmiCopy
	err := ctx.RegisterResource("aws:ec2/amiCopy:AmiCopy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmiCopy gets an existing AmiCopy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmiCopy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiCopyState, opts ...pulumi.ResourceOption) (*AmiCopy, error) {
	var resource AmiCopy
	err := ctx.ReadResource("aws:ec2/amiCopy:AmiCopy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AmiCopy resources.
type amiCopyState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// The ARN of the AMI.
	Arn *string `pulumi:"arn"`
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// The ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn *string `pulumi:"destinationOutpostArn"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiCopyEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
	Encrypted *bool `pulumi:"encrypted"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiCopyEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	Hypervisor            *string                       `pulumi:"hypervisor"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation   *string `pulumi:"imageLocation"`
	ImageOwnerAlias *string `pulumi:"imageOwnerAlias"`
	ImageType       *string `pulumi:"imageType"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId *string `pulumi:"kernelId"`
	// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
	// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
	// if this parameter is not specified, the default CMK for EBS is used
	KmsKeyId           *string `pulumi:"kmsKeyId"`
	ManageEbsSnapshots *bool   `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name            *string `pulumi:"name"`
	OwnerId         *string `pulumi:"ownerId"`
	Platform        *string `pulumi:"platform"`
	PlatformDetails *string `pulumi:"platformDetails"`
	Public          *bool   `pulumi:"public"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	RootSnapshotId *string `pulumi:"rootSnapshotId"`
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId *string `pulumi:"sourceAmiId"`
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion *string `pulumi:"sourceAmiRegion"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags           map[string]string `pulumi:"tags"`
	TagsAll        map[string]string `pulumi:"tagsAll"`
	UsageOperation *string           `pulumi:"usageOperation"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type AmiCopyState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// The ARN of the AMI.
	Arn pulumi.StringPtrInput
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// The ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiCopyEbsBlockDeviceArrayInput
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
	Encrypted pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiCopyEphemeralBlockDeviceArrayInput
	Hypervisor            pulumi.StringPtrInput
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation   pulumi.StringPtrInput
	ImageOwnerAlias pulumi.StringPtrInput
	ImageType       pulumi.StringPtrInput
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringPtrInput
	// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
	// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
	// if this parameter is not specified, the default CMK for EBS is used
	KmsKeyId           pulumi.StringPtrInput
	ManageEbsSnapshots pulumi.BoolPtrInput
	// A region-unique name for the AMI.
	Name            pulumi.StringPtrInput
	OwnerId         pulumi.StringPtrInput
	Platform        pulumi.StringPtrInput
	PlatformDetails pulumi.StringPtrInput
	Public          pulumi.BoolPtrInput
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	RootSnapshotId pulumi.StringPtrInput
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId pulumi.StringPtrInput
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags           pulumi.StringMapInput
	TagsAll        pulumi.StringMapInput
	UsageOperation pulumi.StringPtrInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiCopyState) ElementType() reflect.Type {
	return reflect.TypeOf((*amiCopyState)(nil)).Elem()
}

type amiCopyArgs struct {
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// The ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn *string `pulumi:"destinationOutpostArn"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiCopyEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
	Encrypted *bool `pulumi:"encrypted"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiCopyEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
	// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
	// if this parameter is not specified, the default CMK for EBS is used
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId string `pulumi:"sourceAmiId"`
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion string `pulumi:"sourceAmiRegion"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AmiCopy resource.
type AmiCopyArgs struct {
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// The ARN of the Outpost to which to copy the AMI.
	// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	DestinationOutpostArn pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiCopyEbsBlockDeviceArrayInput
	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshotId`.
	Encrypted pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiCopyEphemeralBlockDeviceArrayInput
	// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
	// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
	// if this parameter is not specified, the default CMK for EBS is used
	KmsKeyId pulumi.StringPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// The id of the AMI to copy. This id must be valid in the region
	// given by `sourceAmiRegion`.
	SourceAmiId pulumi.StringInput
	// The region from which the AMI will be copied. This may be the
	// same as the AWS provider region in order to create a copy within the same region.
	SourceAmiRegion pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AmiCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiCopyArgs)(nil)).Elem()
}

type AmiCopyInput interface {
	pulumi.Input

	ToAmiCopyOutput() AmiCopyOutput
	ToAmiCopyOutputWithContext(ctx context.Context) AmiCopyOutput
}

func (*AmiCopy) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiCopy)(nil))
}

func (i *AmiCopy) ToAmiCopyOutput() AmiCopyOutput {
	return i.ToAmiCopyOutputWithContext(context.Background())
}

func (i *AmiCopy) ToAmiCopyOutputWithContext(ctx context.Context) AmiCopyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiCopyOutput)
}

func (i *AmiCopy) ToAmiCopyPtrOutput() AmiCopyPtrOutput {
	return i.ToAmiCopyPtrOutputWithContext(context.Background())
}

func (i *AmiCopy) ToAmiCopyPtrOutputWithContext(ctx context.Context) AmiCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiCopyPtrOutput)
}

type AmiCopyPtrInput interface {
	pulumi.Input

	ToAmiCopyPtrOutput() AmiCopyPtrOutput
	ToAmiCopyPtrOutputWithContext(ctx context.Context) AmiCopyPtrOutput
}

type amiCopyPtrType AmiCopyArgs

func (*amiCopyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiCopy)(nil))
}

func (i *amiCopyPtrType) ToAmiCopyPtrOutput() AmiCopyPtrOutput {
	return i.ToAmiCopyPtrOutputWithContext(context.Background())
}

func (i *amiCopyPtrType) ToAmiCopyPtrOutputWithContext(ctx context.Context) AmiCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiCopyPtrOutput)
}

// AmiCopyArrayInput is an input type that accepts AmiCopyArray and AmiCopyArrayOutput values.
// You can construct a concrete instance of `AmiCopyArrayInput` via:
//
//          AmiCopyArray{ AmiCopyArgs{...} }
type AmiCopyArrayInput interface {
	pulumi.Input

	ToAmiCopyArrayOutput() AmiCopyArrayOutput
	ToAmiCopyArrayOutputWithContext(context.Context) AmiCopyArrayOutput
}

type AmiCopyArray []AmiCopyInput

func (AmiCopyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AmiCopy)(nil)).Elem()
}

func (i AmiCopyArray) ToAmiCopyArrayOutput() AmiCopyArrayOutput {
	return i.ToAmiCopyArrayOutputWithContext(context.Background())
}

func (i AmiCopyArray) ToAmiCopyArrayOutputWithContext(ctx context.Context) AmiCopyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiCopyArrayOutput)
}

// AmiCopyMapInput is an input type that accepts AmiCopyMap and AmiCopyMapOutput values.
// You can construct a concrete instance of `AmiCopyMapInput` via:
//
//          AmiCopyMap{ "key": AmiCopyArgs{...} }
type AmiCopyMapInput interface {
	pulumi.Input

	ToAmiCopyMapOutput() AmiCopyMapOutput
	ToAmiCopyMapOutputWithContext(context.Context) AmiCopyMapOutput
}

type AmiCopyMap map[string]AmiCopyInput

func (AmiCopyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AmiCopy)(nil)).Elem()
}

func (i AmiCopyMap) ToAmiCopyMapOutput() AmiCopyMapOutput {
	return i.ToAmiCopyMapOutputWithContext(context.Background())
}

func (i AmiCopyMap) ToAmiCopyMapOutputWithContext(ctx context.Context) AmiCopyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiCopyMapOutput)
}

type AmiCopyOutput struct{ *pulumi.OutputState }

func (AmiCopyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiCopy)(nil))
}

func (o AmiCopyOutput) ToAmiCopyOutput() AmiCopyOutput {
	return o
}

func (o AmiCopyOutput) ToAmiCopyOutputWithContext(ctx context.Context) AmiCopyOutput {
	return o
}

func (o AmiCopyOutput) ToAmiCopyPtrOutput() AmiCopyPtrOutput {
	return o.ToAmiCopyPtrOutputWithContext(context.Background())
}

func (o AmiCopyOutput) ToAmiCopyPtrOutputWithContext(ctx context.Context) AmiCopyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmiCopy) *AmiCopy {
		return &v
	}).(AmiCopyPtrOutput)
}

type AmiCopyPtrOutput struct{ *pulumi.OutputState }

func (AmiCopyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiCopy)(nil))
}

func (o AmiCopyPtrOutput) ToAmiCopyPtrOutput() AmiCopyPtrOutput {
	return o
}

func (o AmiCopyPtrOutput) ToAmiCopyPtrOutputWithContext(ctx context.Context) AmiCopyPtrOutput {
	return o
}

func (o AmiCopyPtrOutput) Elem() AmiCopyOutput {
	return o.ApplyT(func(v *AmiCopy) AmiCopy {
		if v != nil {
			return *v
		}
		var ret AmiCopy
		return ret
	}).(AmiCopyOutput)
}

type AmiCopyArrayOutput struct{ *pulumi.OutputState }

func (AmiCopyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmiCopy)(nil))
}

func (o AmiCopyArrayOutput) ToAmiCopyArrayOutput() AmiCopyArrayOutput {
	return o
}

func (o AmiCopyArrayOutput) ToAmiCopyArrayOutputWithContext(ctx context.Context) AmiCopyArrayOutput {
	return o
}

func (o AmiCopyArrayOutput) Index(i pulumi.IntInput) AmiCopyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AmiCopy {
		return vs[0].([]AmiCopy)[vs[1].(int)]
	}).(AmiCopyOutput)
}

type AmiCopyMapOutput struct{ *pulumi.OutputState }

func (AmiCopyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AmiCopy)(nil))
}

func (o AmiCopyMapOutput) ToAmiCopyMapOutput() AmiCopyMapOutput {
	return o
}

func (o AmiCopyMapOutput) ToAmiCopyMapOutputWithContext(ctx context.Context) AmiCopyMapOutput {
	return o
}

func (o AmiCopyMapOutput) MapIndex(k pulumi.StringInput) AmiCopyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AmiCopy {
		return vs[0].(map[string]AmiCopy)[vs[1].(string)]
	}).(AmiCopyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AmiCopyInput)(nil)).Elem(), &AmiCopy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmiCopyPtrInput)(nil)).Elem(), &AmiCopy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmiCopyArrayInput)(nil)).Elem(), AmiCopyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AmiCopyMapInput)(nil)).Elem(), AmiCopyMap{})
	pulumi.RegisterOutputType(AmiCopyOutput{})
	pulumi.RegisterOutputType(AmiCopyPtrOutput{})
	pulumi.RegisterOutputType(AmiCopyArrayOutput{})
	pulumi.RegisterOutputType(AmiCopyMapOutput{})
}
