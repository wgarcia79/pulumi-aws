// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synthetics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Synthetics Canary resource.
//
// > **NOTE:** When you create a canary, AWS creates supporting implicit resources. See the Amazon CloudWatch Synthetics documentation on [DeleteCanary](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DeleteCanary.html) for a full list. Neither AWS nor this provider deletes these implicit resources automatically when the canary is deleted. Before deleting a canary, ensure you have all the information about the canary that you need to delete the implicit resources using the AWS Console, or AWS CLI.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/synthetics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := synthetics.NewCanary(ctx, "some", &synthetics.CanaryArgs{
// 			ArtifactS3Location: pulumi.String("s3://some-bucket/"),
// 			ExecutionRoleArn:   pulumi.String("some-role"),
// 			Handler:            pulumi.String("exports.handler"),
// 			RuntimeVersion:     pulumi.String("syn-1.0"),
// 			Schedule: &synthetics.CanaryScheduleArgs{
// 				Expression: pulumi.String("rate(0 minute)"),
// 			},
// 			ZipFile: pulumi.String("test-fixtures/lambdatest.zip"),
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
// Synthetics Canaries can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:synthetics/canary:Canary some some-canary
// ```
type Canary struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Canary.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location pulumi.StringOutput `pulumi:"artifactS3Location"`
	// ARN of the Lambda function that is used as your canary's engine.
	EngineArn pulumi.StringOutput `pulumi:"engineArn"`
	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod pulumi.IntPtrOutput `pulumi:"failureRetentionPeriod"`
	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler pulumi.StringOutput `pulumi:"handler"`
	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration block for individual canary runs. Detailed below.
	RunConfig CanaryRunConfigOutput `pulumi:"runConfig"`
	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion pulumi.StringOutput `pulumi:"runtimeVersion"`
	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zipFile`.**
	S3Bucket pulumi.StringPtrOutput `pulumi:"s3Bucket"`
	// S3 key of your script. **Conflicts with `zipFile`.**
	S3Key pulumi.StringPtrOutput `pulumi:"s3Key"`
	// S3 version ID of your script. **Conflicts with `zipFile`.**
	S3Version pulumi.StringPtrOutput `pulumi:"s3Version"`
	// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
	Schedule CanaryScheduleOutput `pulumi:"schedule"`
	// ARN of the Lambda layer where Synthetics stores the canary script code.
	SourceLocationArn pulumi.StringOutput `pulumi:"sourceLocationArn"`
	// Whether to run or stop the canary.
	StartCanary pulumi.BoolPtrOutput `pulumi:"startCanary"`
	// Canary status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod pulumi.IntPtrOutput `pulumi:"successRetentionPeriod"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Structure that contains information about when the canary was created, modified, and most recently run. see Timeline.
	Timelines CanaryTimelineArrayOutput `pulumi:"timelines"`
	// Configuration block. Detailed below.
	VpcConfig CanaryVpcConfigPtrOutput `pulumi:"vpcConfig"`
	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3Bucket`, `s3Key`, and `s3Version`.**
	ZipFile pulumi.StringPtrOutput `pulumi:"zipFile"`
}

// NewCanary registers a new resource with the given unique name, arguments, and options.
func NewCanary(ctx *pulumi.Context,
	name string, args *CanaryArgs, opts ...pulumi.ResourceOption) (*Canary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArtifactS3Location == nil {
		return nil, errors.New("invalid value for required argument 'ArtifactS3Location'")
	}
	if args.ExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionRoleArn'")
	}
	if args.Handler == nil {
		return nil, errors.New("invalid value for required argument 'Handler'")
	}
	if args.RuntimeVersion == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeVersion'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	var resource Canary
	err := ctx.RegisterResource("aws:synthetics/canary:Canary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCanary gets an existing Canary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCanary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CanaryState, opts ...pulumi.ResourceOption) (*Canary, error) {
	var resource Canary
	err := ctx.ReadResource("aws:synthetics/canary:Canary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Canary resources.
type canaryState struct {
	// Amazon Resource Name (ARN) of the Canary.
	Arn *string `pulumi:"arn"`
	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location *string `pulumi:"artifactS3Location"`
	// ARN of the Lambda function that is used as your canary's engine.
	EngineArn *string `pulumi:"engineArn"`
	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod *int `pulumi:"failureRetentionPeriod"`
	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler *string `pulumi:"handler"`
	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name *string `pulumi:"name"`
	// Configuration block for individual canary runs. Detailed below.
	RunConfig *CanaryRunConfig `pulumi:"runConfig"`
	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zipFile`.**
	S3Bucket *string `pulumi:"s3Bucket"`
	// S3 key of your script. **Conflicts with `zipFile`.**
	S3Key *string `pulumi:"s3Key"`
	// S3 version ID of your script. **Conflicts with `zipFile`.**
	S3Version *string `pulumi:"s3Version"`
	// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
	Schedule *CanarySchedule `pulumi:"schedule"`
	// ARN of the Lambda layer where Synthetics stores the canary script code.
	SourceLocationArn *string `pulumi:"sourceLocationArn"`
	// Whether to run or stop the canary.
	StartCanary *bool `pulumi:"startCanary"`
	// Canary status.
	Status *string `pulumi:"status"`
	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod *int `pulumi:"successRetentionPeriod"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Structure that contains information about when the canary was created, modified, and most recently run. see Timeline.
	Timelines []CanaryTimeline `pulumi:"timelines"`
	// Configuration block. Detailed below.
	VpcConfig *CanaryVpcConfig `pulumi:"vpcConfig"`
	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3Bucket`, `s3Key`, and `s3Version`.**
	ZipFile *string `pulumi:"zipFile"`
}

type CanaryState struct {
	// Amazon Resource Name (ARN) of the Canary.
	Arn pulumi.StringPtrInput
	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location pulumi.StringPtrInput
	// ARN of the Lambda function that is used as your canary's engine.
	EngineArn pulumi.StringPtrInput
	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn pulumi.StringPtrInput
	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod pulumi.IntPtrInput
	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler pulumi.StringPtrInput
	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name pulumi.StringPtrInput
	// Configuration block for individual canary runs. Detailed below.
	RunConfig CanaryRunConfigPtrInput
	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion pulumi.StringPtrInput
	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zipFile`.**
	S3Bucket pulumi.StringPtrInput
	// S3 key of your script. **Conflicts with `zipFile`.**
	S3Key pulumi.StringPtrInput
	// S3 version ID of your script. **Conflicts with `zipFile`.**
	S3Version pulumi.StringPtrInput
	// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
	Schedule CanarySchedulePtrInput
	// ARN of the Lambda layer where Synthetics stores the canary script code.
	SourceLocationArn pulumi.StringPtrInput
	// Whether to run or stop the canary.
	StartCanary pulumi.BoolPtrInput
	// Canary status.
	Status pulumi.StringPtrInput
	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod pulumi.IntPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Structure that contains information about when the canary was created, modified, and most recently run. see Timeline.
	Timelines CanaryTimelineArrayInput
	// Configuration block. Detailed below.
	VpcConfig CanaryVpcConfigPtrInput
	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3Bucket`, `s3Key`, and `s3Version`.**
	ZipFile pulumi.StringPtrInput
}

func (CanaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*canaryState)(nil)).Elem()
}

type canaryArgs struct {
	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location string `pulumi:"artifactS3Location"`
	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod *int `pulumi:"failureRetentionPeriod"`
	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler string `pulumi:"handler"`
	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name *string `pulumi:"name"`
	// Configuration block for individual canary runs. Detailed below.
	RunConfig *CanaryRunConfig `pulumi:"runConfig"`
	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion string `pulumi:"runtimeVersion"`
	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zipFile`.**
	S3Bucket *string `pulumi:"s3Bucket"`
	// S3 key of your script. **Conflicts with `zipFile`.**
	S3Key *string `pulumi:"s3Key"`
	// S3 version ID of your script. **Conflicts with `zipFile`.**
	S3Version *string `pulumi:"s3Version"`
	// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
	Schedule CanarySchedule `pulumi:"schedule"`
	// Whether to run or stop the canary.
	StartCanary *bool `pulumi:"startCanary"`
	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod *int `pulumi:"successRetentionPeriod"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Configuration block. Detailed below.
	VpcConfig *CanaryVpcConfig `pulumi:"vpcConfig"`
	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3Bucket`, `s3Key`, and `s3Version`.**
	ZipFile *string `pulumi:"zipFile"`
}

// The set of arguments for constructing a Canary resource.
type CanaryArgs struct {
	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location pulumi.StringInput
	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn pulumi.StringInput
	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod pulumi.IntPtrInput
	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler pulumi.StringInput
	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name pulumi.StringPtrInput
	// Configuration block for individual canary runs. Detailed below.
	RunConfig CanaryRunConfigPtrInput
	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion pulumi.StringInput
	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. Specify the full bucket name including s3:// as the start of the bucket name. **Conflicts with `zipFile`.**
	S3Bucket pulumi.StringPtrInput
	// S3 key of your script. **Conflicts with `zipFile`.**
	S3Key pulumi.StringPtrInput
	// S3 version ID of your script. **Conflicts with `zipFile`.**
	S3Version pulumi.StringPtrInput
	// Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.
	Schedule CanaryScheduleInput
	// Whether to run or stop the canary.
	StartCanary pulumi.BoolPtrInput
	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod pulumi.IntPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Configuration block. Detailed below.
	VpcConfig CanaryVpcConfigPtrInput
	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 5 MB. **Conflicts with `s3Bucket`, `s3Key`, and `s3Version`.**
	ZipFile pulumi.StringPtrInput
}

func (CanaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*canaryArgs)(nil)).Elem()
}

type CanaryInput interface {
	pulumi.Input

	ToCanaryOutput() CanaryOutput
	ToCanaryOutputWithContext(ctx context.Context) CanaryOutput
}

func (*Canary) ElementType() reflect.Type {
	return reflect.TypeOf((*Canary)(nil))
}

func (i *Canary) ToCanaryOutput() CanaryOutput {
	return i.ToCanaryOutputWithContext(context.Background())
}

func (i *Canary) ToCanaryOutputWithContext(ctx context.Context) CanaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CanaryOutput)
}

func (i *Canary) ToCanaryPtrOutput() CanaryPtrOutput {
	return i.ToCanaryPtrOutputWithContext(context.Background())
}

func (i *Canary) ToCanaryPtrOutputWithContext(ctx context.Context) CanaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CanaryPtrOutput)
}

type CanaryPtrInput interface {
	pulumi.Input

	ToCanaryPtrOutput() CanaryPtrOutput
	ToCanaryPtrOutputWithContext(ctx context.Context) CanaryPtrOutput
}

type canaryPtrType CanaryArgs

func (*canaryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Canary)(nil))
}

func (i *canaryPtrType) ToCanaryPtrOutput() CanaryPtrOutput {
	return i.ToCanaryPtrOutputWithContext(context.Background())
}

func (i *canaryPtrType) ToCanaryPtrOutputWithContext(ctx context.Context) CanaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CanaryPtrOutput)
}

// CanaryArrayInput is an input type that accepts CanaryArray and CanaryArrayOutput values.
// You can construct a concrete instance of `CanaryArrayInput` via:
//
//          CanaryArray{ CanaryArgs{...} }
type CanaryArrayInput interface {
	pulumi.Input

	ToCanaryArrayOutput() CanaryArrayOutput
	ToCanaryArrayOutputWithContext(context.Context) CanaryArrayOutput
}

type CanaryArray []CanaryInput

func (CanaryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Canary)(nil)).Elem()
}

func (i CanaryArray) ToCanaryArrayOutput() CanaryArrayOutput {
	return i.ToCanaryArrayOutputWithContext(context.Background())
}

func (i CanaryArray) ToCanaryArrayOutputWithContext(ctx context.Context) CanaryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CanaryArrayOutput)
}

// CanaryMapInput is an input type that accepts CanaryMap and CanaryMapOutput values.
// You can construct a concrete instance of `CanaryMapInput` via:
//
//          CanaryMap{ "key": CanaryArgs{...} }
type CanaryMapInput interface {
	pulumi.Input

	ToCanaryMapOutput() CanaryMapOutput
	ToCanaryMapOutputWithContext(context.Context) CanaryMapOutput
}

type CanaryMap map[string]CanaryInput

func (CanaryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Canary)(nil)).Elem()
}

func (i CanaryMap) ToCanaryMapOutput() CanaryMapOutput {
	return i.ToCanaryMapOutputWithContext(context.Background())
}

func (i CanaryMap) ToCanaryMapOutputWithContext(ctx context.Context) CanaryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CanaryMapOutput)
}

type CanaryOutput struct{ *pulumi.OutputState }

func (CanaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Canary)(nil))
}

func (o CanaryOutput) ToCanaryOutput() CanaryOutput {
	return o
}

func (o CanaryOutput) ToCanaryOutputWithContext(ctx context.Context) CanaryOutput {
	return o
}

func (o CanaryOutput) ToCanaryPtrOutput() CanaryPtrOutput {
	return o.ToCanaryPtrOutputWithContext(context.Background())
}

func (o CanaryOutput) ToCanaryPtrOutputWithContext(ctx context.Context) CanaryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Canary) *Canary {
		return &v
	}).(CanaryPtrOutput)
}

type CanaryPtrOutput struct{ *pulumi.OutputState }

func (CanaryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Canary)(nil))
}

func (o CanaryPtrOutput) ToCanaryPtrOutput() CanaryPtrOutput {
	return o
}

func (o CanaryPtrOutput) ToCanaryPtrOutputWithContext(ctx context.Context) CanaryPtrOutput {
	return o
}

func (o CanaryPtrOutput) Elem() CanaryOutput {
	return o.ApplyT(func(v *Canary) Canary {
		if v != nil {
			return *v
		}
		var ret Canary
		return ret
	}).(CanaryOutput)
}

type CanaryArrayOutput struct{ *pulumi.OutputState }

func (CanaryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Canary)(nil))
}

func (o CanaryArrayOutput) ToCanaryArrayOutput() CanaryArrayOutput {
	return o
}

func (o CanaryArrayOutput) ToCanaryArrayOutputWithContext(ctx context.Context) CanaryArrayOutput {
	return o
}

func (o CanaryArrayOutput) Index(i pulumi.IntInput) CanaryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Canary {
		return vs[0].([]Canary)[vs[1].(int)]
	}).(CanaryOutput)
}

type CanaryMapOutput struct{ *pulumi.OutputState }

func (CanaryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Canary)(nil))
}

func (o CanaryMapOutput) ToCanaryMapOutput() CanaryMapOutput {
	return o
}

func (o CanaryMapOutput) ToCanaryMapOutputWithContext(ctx context.Context) CanaryMapOutput {
	return o
}

func (o CanaryMapOutput) MapIndex(k pulumi.StringInput) CanaryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Canary {
		return vs[0].(map[string]Canary)[vs[1].(string)]
	}).(CanaryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CanaryInput)(nil)).Elem(), &Canary{})
	pulumi.RegisterInputType(reflect.TypeOf((*CanaryPtrInput)(nil)).Elem(), &Canary{})
	pulumi.RegisterInputType(reflect.TypeOf((*CanaryArrayInput)(nil)).Elem(), CanaryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CanaryMapInput)(nil)).Elem(), CanaryMap{})
	pulumi.RegisterOutputType(CanaryOutput{})
	pulumi.RegisterOutputType(CanaryPtrOutput{})
	pulumi.RegisterOutputType(CanaryArrayOutput{})
	pulumi.RegisterOutputType(CanaryMapOutput{})
}
