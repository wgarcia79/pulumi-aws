// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
// scales elastically for real-time processing of streaming big data.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kinesis"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			RetentionPeriod: pulumi.Int(48),
// 			ShardCount:      pulumi.Int(1),
// 			ShardLevelMetrics: pulumi.StringArray{
// 				pulumi.String("IncomingBytes"),
// 				pulumi.String("OutgoingBytes"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("test"),
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
// Kinesis Streams can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:kinesis/stream:Stream test_stream kinesis-test
// ```
//
//  [1]https://aws.amazon.com/documentation/kinesis/ [2]https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html [3]https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html
type Stream struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrOutput `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringOutput `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntOutput `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayOutput `pulumi:"shardLevelMetrics"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	TagsAll           pulumi.StringMapOutput   `pulumi:"tagsAll"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ShardCount == nil {
		return nil, errors.New("invalid value for required argument 'ShardCount'")
	}
	var resource Stream
	err := ctx.RegisterResource("aws:kinesis/stream:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("aws:kinesis/stream:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType *string `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion *bool `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount *int `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics []string          `pulumi:"shardLevelMetrics"`
	Tags              map[string]string `pulumi:"tags"`
	TagsAll           map[string]string `pulumi:"tagsAll"`
}

type StreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrInput
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrInput
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntPtrInput
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayInput
	Tags              pulumi.StringMapInput
	TagsAll           pulumi.StringMapInput
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType *string `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion *bool `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount int `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics []string          `pulumi:"shardLevelMetrics"`
	Tags              map[string]string `pulumi:"tags"`
	TagsAll           map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrInput
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrInput
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntInput
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayInput
	Tags              pulumi.StringMapInput
	TagsAll           pulumi.StringMapInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((*Stream)(nil))
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

func (i *Stream) ToStreamPtrOutput() StreamPtrOutput {
	return i.ToStreamPtrOutputWithContext(context.Background())
}

func (i *Stream) ToStreamPtrOutputWithContext(ctx context.Context) StreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamPtrOutput)
}

type StreamPtrInput interface {
	pulumi.Input

	ToStreamPtrOutput() StreamPtrOutput
	ToStreamPtrOutputWithContext(ctx context.Context) StreamPtrOutput
}

type streamPtrType StreamArgs

func (*streamPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil))
}

func (i *streamPtrType) ToStreamPtrOutput() StreamPtrOutput {
	return i.ToStreamPtrOutputWithContext(context.Background())
}

func (i *streamPtrType) ToStreamPtrOutputWithContext(ctx context.Context) StreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamPtrOutput)
}

// StreamArrayInput is an input type that accepts StreamArray and StreamArrayOutput values.
// You can construct a concrete instance of `StreamArrayInput` via:
//
//          StreamArray{ StreamArgs{...} }
type StreamArrayInput interface {
	pulumi.Input

	ToStreamArrayOutput() StreamArrayOutput
	ToStreamArrayOutputWithContext(context.Context) StreamArrayOutput
}

type StreamArray []StreamInput

func (StreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stream)(nil)).Elem()
}

func (i StreamArray) ToStreamArrayOutput() StreamArrayOutput {
	return i.ToStreamArrayOutputWithContext(context.Background())
}

func (i StreamArray) ToStreamArrayOutputWithContext(ctx context.Context) StreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamArrayOutput)
}

// StreamMapInput is an input type that accepts StreamMap and StreamMapOutput values.
// You can construct a concrete instance of `StreamMapInput` via:
//
//          StreamMap{ "key": StreamArgs{...} }
type StreamMapInput interface {
	pulumi.Input

	ToStreamMapOutput() StreamMapOutput
	ToStreamMapOutputWithContext(context.Context) StreamMapOutput
}

type StreamMap map[string]StreamInput

func (StreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stream)(nil)).Elem()
}

func (i StreamMap) ToStreamMapOutput() StreamMapOutput {
	return i.ToStreamMapOutputWithContext(context.Background())
}

func (i StreamMap) ToStreamMapOutputWithContext(ctx context.Context) StreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamMapOutput)
}

type StreamOutput struct {
	*pulumi.OutputState
}

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Stream)(nil))
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

func (o StreamOutput) ToStreamPtrOutput() StreamPtrOutput {
	return o.ToStreamPtrOutputWithContext(context.Background())
}

func (o StreamOutput) ToStreamPtrOutputWithContext(ctx context.Context) StreamPtrOutput {
	return o.ApplyT(func(v Stream) *Stream {
		return &v
	}).(StreamPtrOutput)
}

type StreamPtrOutput struct {
	*pulumi.OutputState
}

func (StreamPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil))
}

func (o StreamPtrOutput) ToStreamPtrOutput() StreamPtrOutput {
	return o
}

func (o StreamPtrOutput) ToStreamPtrOutputWithContext(ctx context.Context) StreamPtrOutput {
	return o
}

type StreamArrayOutput struct{ *pulumi.OutputState }

func (StreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Stream)(nil))
}

func (o StreamArrayOutput) ToStreamArrayOutput() StreamArrayOutput {
	return o
}

func (o StreamArrayOutput) ToStreamArrayOutputWithContext(ctx context.Context) StreamArrayOutput {
	return o
}

func (o StreamArrayOutput) Index(i pulumi.IntInput) StreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Stream {
		return vs[0].([]Stream)[vs[1].(int)]
	}).(StreamOutput)
}

type StreamMapOutput struct{ *pulumi.OutputState }

func (StreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Stream)(nil))
}

func (o StreamMapOutput) ToStreamMapOutput() StreamMapOutput {
	return o
}

func (o StreamMapOutput) ToStreamMapOutputWithContext(ctx context.Context) StreamMapOutput {
	return o
}

func (o StreamMapOutput) MapIndex(k pulumi.StringInput) StreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Stream {
		return vs[0].(map[string]Stream)[vs[1].(string)]
	}).(StreamOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamOutput{})
	pulumi.RegisterOutputType(StreamPtrOutput{})
	pulumi.RegisterOutputType(StreamArrayOutput{})
	pulumi.RegisterOutputType(StreamMapOutput{})
}
