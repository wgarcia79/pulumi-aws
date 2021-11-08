// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Pinpoint Event Stream resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kinesis"
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
// 		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			ShardCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"pinpoint.us-east-1.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewEventStream(ctx, "stream", &pinpoint.EventStreamArgs{
// 			ApplicationId:        app.ApplicationId,
// 			DestinationStreamArn: testStream.Arn,
// 			RoleArn:              testRole.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
// 			Role:   testRole.ID(),
// 			Policy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Action\": [\n", "      \"kinesis:PutRecords\",\n", "      \"kinesis:DescribeStream\"\n", "    ],\n", "    \"Effect\": \"Allow\",\n", "    \"Resource\": [\n", "      \"arn:aws:kinesis:us-east-1:*:*/*\"\n", "    ]\n", "  }\n", "}\n")),
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
// Pinpoint Event Stream can be imported using the `application-id`, e.g.,
//
// ```sh
//  $ pulumi import aws:pinpoint/eventStream:EventStream stream application-id
// ```
type EventStream struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringOutput `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewEventStream registers a new resource with the given unique name, arguments, and options.
func NewEventStream(ctx *pulumi.Context,
	name string, args *EventStreamArgs, opts ...pulumi.ResourceOption) (*EventStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.DestinationStreamArn == nil {
		return nil, errors.New("invalid value for required argument 'DestinationStreamArn'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource EventStream
	err := ctx.RegisterResource("aws:pinpoint/eventStream:EventStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventStream gets an existing EventStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventStreamState, opts ...pulumi.ResourceOption) (*EventStream, error) {
	var resource EventStream
	err := ctx.ReadResource("aws:pinpoint/eventStream:EventStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventStream resources.
type eventStreamState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn *string `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn *string `pulumi:"roleArn"`
}

type EventStreamState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringPtrInput
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringPtrInput
}

func (EventStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventStreamState)(nil)).Elem()
}

type eventStreamArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn string `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a EventStream resource.
type EventStreamArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringInput
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringInput
}

func (EventStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventStreamArgs)(nil)).Elem()
}

type EventStreamInput interface {
	pulumi.Input

	ToEventStreamOutput() EventStreamOutput
	ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput
}

func (*EventStream) ElementType() reflect.Type {
	return reflect.TypeOf((*EventStream)(nil))
}

func (i *EventStream) ToEventStreamOutput() EventStreamOutput {
	return i.ToEventStreamOutputWithContext(context.Background())
}

func (i *EventStream) ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamOutput)
}

func (i *EventStream) ToEventStreamPtrOutput() EventStreamPtrOutput {
	return i.ToEventStreamPtrOutputWithContext(context.Background())
}

func (i *EventStream) ToEventStreamPtrOutputWithContext(ctx context.Context) EventStreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamPtrOutput)
}

type EventStreamPtrInput interface {
	pulumi.Input

	ToEventStreamPtrOutput() EventStreamPtrOutput
	ToEventStreamPtrOutputWithContext(ctx context.Context) EventStreamPtrOutput
}

type eventStreamPtrType EventStreamArgs

func (*eventStreamPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventStream)(nil))
}

func (i *eventStreamPtrType) ToEventStreamPtrOutput() EventStreamPtrOutput {
	return i.ToEventStreamPtrOutputWithContext(context.Background())
}

func (i *eventStreamPtrType) ToEventStreamPtrOutputWithContext(ctx context.Context) EventStreamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamPtrOutput)
}

// EventStreamArrayInput is an input type that accepts EventStreamArray and EventStreamArrayOutput values.
// You can construct a concrete instance of `EventStreamArrayInput` via:
//
//          EventStreamArray{ EventStreamArgs{...} }
type EventStreamArrayInput interface {
	pulumi.Input

	ToEventStreamArrayOutput() EventStreamArrayOutput
	ToEventStreamArrayOutputWithContext(context.Context) EventStreamArrayOutput
}

type EventStreamArray []EventStreamInput

func (EventStreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventStream)(nil)).Elem()
}

func (i EventStreamArray) ToEventStreamArrayOutput() EventStreamArrayOutput {
	return i.ToEventStreamArrayOutputWithContext(context.Background())
}

func (i EventStreamArray) ToEventStreamArrayOutputWithContext(ctx context.Context) EventStreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamArrayOutput)
}

// EventStreamMapInput is an input type that accepts EventStreamMap and EventStreamMapOutput values.
// You can construct a concrete instance of `EventStreamMapInput` via:
//
//          EventStreamMap{ "key": EventStreamArgs{...} }
type EventStreamMapInput interface {
	pulumi.Input

	ToEventStreamMapOutput() EventStreamMapOutput
	ToEventStreamMapOutputWithContext(context.Context) EventStreamMapOutput
}

type EventStreamMap map[string]EventStreamInput

func (EventStreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventStream)(nil)).Elem()
}

func (i EventStreamMap) ToEventStreamMapOutput() EventStreamMapOutput {
	return i.ToEventStreamMapOutputWithContext(context.Background())
}

func (i EventStreamMap) ToEventStreamMapOutputWithContext(ctx context.Context) EventStreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamMapOutput)
}

type EventStreamOutput struct{ *pulumi.OutputState }

func (EventStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventStream)(nil))
}

func (o EventStreamOutput) ToEventStreamOutput() EventStreamOutput {
	return o
}

func (o EventStreamOutput) ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput {
	return o
}

func (o EventStreamOutput) ToEventStreamPtrOutput() EventStreamPtrOutput {
	return o.ToEventStreamPtrOutputWithContext(context.Background())
}

func (o EventStreamOutput) ToEventStreamPtrOutputWithContext(ctx context.Context) EventStreamPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventStream) *EventStream {
		return &v
	}).(EventStreamPtrOutput)
}

type EventStreamPtrOutput struct{ *pulumi.OutputState }

func (EventStreamPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventStream)(nil))
}

func (o EventStreamPtrOutput) ToEventStreamPtrOutput() EventStreamPtrOutput {
	return o
}

func (o EventStreamPtrOutput) ToEventStreamPtrOutputWithContext(ctx context.Context) EventStreamPtrOutput {
	return o
}

func (o EventStreamPtrOutput) Elem() EventStreamOutput {
	return o.ApplyT(func(v *EventStream) EventStream {
		if v != nil {
			return *v
		}
		var ret EventStream
		return ret
	}).(EventStreamOutput)
}

type EventStreamArrayOutput struct{ *pulumi.OutputState }

func (EventStreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventStream)(nil))
}

func (o EventStreamArrayOutput) ToEventStreamArrayOutput() EventStreamArrayOutput {
	return o
}

func (o EventStreamArrayOutput) ToEventStreamArrayOutputWithContext(ctx context.Context) EventStreamArrayOutput {
	return o
}

func (o EventStreamArrayOutput) Index(i pulumi.IntInput) EventStreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventStream {
		return vs[0].([]EventStream)[vs[1].(int)]
	}).(EventStreamOutput)
}

type EventStreamMapOutput struct{ *pulumi.OutputState }

func (EventStreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventStream)(nil))
}

func (o EventStreamMapOutput) ToEventStreamMapOutput() EventStreamMapOutput {
	return o
}

func (o EventStreamMapOutput) ToEventStreamMapOutputWithContext(ctx context.Context) EventStreamMapOutput {
	return o
}

func (o EventStreamMapOutput) MapIndex(k pulumi.StringInput) EventStreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventStream {
		return vs[0].(map[string]EventStream)[vs[1].(string)]
	}).(EventStreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventStreamInput)(nil)).Elem(), &EventStream{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventStreamPtrInput)(nil)).Elem(), &EventStream{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventStreamArrayInput)(nil)).Elem(), EventStreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventStreamMapInput)(nil)).Elem(), EventStreamMap{})
	pulumi.RegisterOutputType(EventStreamOutput{})
	pulumi.RegisterOutputType(EventStreamPtrOutput{})
	pulumi.RegisterOutputType(EventStreamArrayOutput{})
	pulumi.RegisterOutputType(EventStreamMapOutput{})
}
