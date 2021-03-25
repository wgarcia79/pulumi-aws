// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EventBridge event archive resource.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		orderEventBus, err := cloudwatch.NewEventBus(ctx, "orderEventBus", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventArchive(ctx, "orderEventArchive", &cloudwatch.EventArchiveArgs{
// 			EventSourceArn: orderEventBus.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Example all optional arguments
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		orderEventBus, err := cloudwatch.NewEventBus(ctx, "orderEventBus", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventArchive(ctx, "orderEventArchive", &cloudwatch.EventArchiveArgs{
// 			Description:    pulumi.String("Archived events from order service"),
// 			EventSourceArn: orderEventBus.Arn,
// 			RetentionDays:  pulumi.Int(7),
// 			EventPattern:   pulumi.String(fmt.Sprintf("%v%v%v", "{\n", "  \"source\": [\"company.team.order\"]\n", "}\n")),
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
// Event Archive can be imported using their name, for example bash
//
// ```sh
//  $ pulumi import aws:cloudwatch/eventArchive:EventArchive imported_event_archive order-archive
// ```
type EventArchive struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the event archive.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the new event archive.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrOutput `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringOutput `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewEventArchive registers a new resource with the given unique name, arguments, and options.
func NewEventArchive(ctx *pulumi.Context,
	name string, args *EventArchiveArgs, opts ...pulumi.ResourceOption) (*EventArchive, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventSourceArn == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceArn'")
	}
	var resource EventArchive
	err := ctx.RegisterResource("aws:cloudwatch/eventArchive:EventArchive", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventArchive gets an existing EventArchive resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventArchive(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventArchiveState, opts ...pulumi.ResourceOption) (*EventArchive, error) {
	var resource EventArchive
	err := ctx.ReadResource("aws:cloudwatch/eventArchive:EventArchive", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventArchive resources.
type eventArchiveState struct {
	// The Amazon Resource Name (ARN) of the event archive.
	Arn *string `pulumi:"arn"`
	// The description of the new event archive.
	Description *string `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern *string `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn *string `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name *string `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *int `pulumi:"retentionDays"`
}

type EventArchiveState struct {
	// The Amazon Resource Name (ARN) of the event archive.
	Arn pulumi.StringPtrInput
	// The description of the new event archive.
	Description pulumi.StringPtrInput
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrInput
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringPtrInput
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringPtrInput
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrInput
}

func (EventArchiveState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArchiveState)(nil)).Elem()
}

type eventArchiveArgs struct {
	// The description of the new event archive.
	Description *string `pulumi:"description"`
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern *string `pulumi:"eventPattern"`
	// Event bus source ARN from where these events should be archived.
	EventSourceArn string `pulumi:"eventSourceArn"`
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name *string `pulumi:"name"`
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a EventArchive resource.
type EventArchiveArgs struct {
	// The description of the new event archive.
	Description pulumi.StringPtrInput
	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `eventSourceArn`.
	EventPattern pulumi.StringPtrInput
	// Event bus source ARN from where these events should be archived.
	EventSourceArn pulumi.StringInput
	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name pulumi.StringPtrInput
	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays pulumi.IntPtrInput
}

func (EventArchiveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArchiveArgs)(nil)).Elem()
}

type EventArchiveInput interface {
	pulumi.Input

	ToEventArchiveOutput() EventArchiveOutput
	ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput
}

func (*EventArchive) ElementType() reflect.Type {
	return reflect.TypeOf((*EventArchive)(nil))
}

func (i *EventArchive) ToEventArchiveOutput() EventArchiveOutput {
	return i.ToEventArchiveOutputWithContext(context.Background())
}

func (i *EventArchive) ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveOutput)
}

func (i *EventArchive) ToEventArchivePtrOutput() EventArchivePtrOutput {
	return i.ToEventArchivePtrOutputWithContext(context.Background())
}

func (i *EventArchive) ToEventArchivePtrOutputWithContext(ctx context.Context) EventArchivePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchivePtrOutput)
}

type EventArchivePtrInput interface {
	pulumi.Input

	ToEventArchivePtrOutput() EventArchivePtrOutput
	ToEventArchivePtrOutputWithContext(ctx context.Context) EventArchivePtrOutput
}

type eventArchivePtrType EventArchiveArgs

func (*eventArchivePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventArchive)(nil))
}

func (i *eventArchivePtrType) ToEventArchivePtrOutput() EventArchivePtrOutput {
	return i.ToEventArchivePtrOutputWithContext(context.Background())
}

func (i *eventArchivePtrType) ToEventArchivePtrOutputWithContext(ctx context.Context) EventArchivePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchivePtrOutput)
}

// EventArchiveArrayInput is an input type that accepts EventArchiveArray and EventArchiveArrayOutput values.
// You can construct a concrete instance of `EventArchiveArrayInput` via:
//
//          EventArchiveArray{ EventArchiveArgs{...} }
type EventArchiveArrayInput interface {
	pulumi.Input

	ToEventArchiveArrayOutput() EventArchiveArrayOutput
	ToEventArchiveArrayOutputWithContext(context.Context) EventArchiveArrayOutput
}

type EventArchiveArray []EventArchiveInput

func (EventArchiveArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EventArchive)(nil))
}

func (i EventArchiveArray) ToEventArchiveArrayOutput() EventArchiveArrayOutput {
	return i.ToEventArchiveArrayOutputWithContext(context.Background())
}

func (i EventArchiveArray) ToEventArchiveArrayOutputWithContext(ctx context.Context) EventArchiveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveArrayOutput)
}

// EventArchiveMapInput is an input type that accepts EventArchiveMap and EventArchiveMapOutput values.
// You can construct a concrete instance of `EventArchiveMapInput` via:
//
//          EventArchiveMap{ "key": EventArchiveArgs{...} }
type EventArchiveMapInput interface {
	pulumi.Input

	ToEventArchiveMapOutput() EventArchiveMapOutput
	ToEventArchiveMapOutputWithContext(context.Context) EventArchiveMapOutput
}

type EventArchiveMap map[string]EventArchiveInput

func (EventArchiveMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EventArchive)(nil))
}

func (i EventArchiveMap) ToEventArchiveMapOutput() EventArchiveMapOutput {
	return i.ToEventArchiveMapOutputWithContext(context.Background())
}

func (i EventArchiveMap) ToEventArchiveMapOutputWithContext(ctx context.Context) EventArchiveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArchiveMapOutput)
}

type EventArchiveOutput struct {
	*pulumi.OutputState
}

func (EventArchiveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventArchive)(nil))
}

func (o EventArchiveOutput) ToEventArchiveOutput() EventArchiveOutput {
	return o
}

func (o EventArchiveOutput) ToEventArchiveOutputWithContext(ctx context.Context) EventArchiveOutput {
	return o
}

func (o EventArchiveOutput) ToEventArchivePtrOutput() EventArchivePtrOutput {
	return o.ToEventArchivePtrOutputWithContext(context.Background())
}

func (o EventArchiveOutput) ToEventArchivePtrOutputWithContext(ctx context.Context) EventArchivePtrOutput {
	return o.ApplyT(func(v EventArchive) *EventArchive {
		return &v
	}).(EventArchivePtrOutput)
}

type EventArchivePtrOutput struct {
	*pulumi.OutputState
}

func (EventArchivePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventArchive)(nil))
}

func (o EventArchivePtrOutput) ToEventArchivePtrOutput() EventArchivePtrOutput {
	return o
}

func (o EventArchivePtrOutput) ToEventArchivePtrOutputWithContext(ctx context.Context) EventArchivePtrOutput {
	return o
}

type EventArchiveArrayOutput struct{ *pulumi.OutputState }

func (EventArchiveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventArchive)(nil))
}

func (o EventArchiveArrayOutput) ToEventArchiveArrayOutput() EventArchiveArrayOutput {
	return o
}

func (o EventArchiveArrayOutput) ToEventArchiveArrayOutputWithContext(ctx context.Context) EventArchiveArrayOutput {
	return o
}

func (o EventArchiveArrayOutput) Index(i pulumi.IntInput) EventArchiveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventArchive {
		return vs[0].([]EventArchive)[vs[1].(int)]
	}).(EventArchiveOutput)
}

type EventArchiveMapOutput struct{ *pulumi.OutputState }

func (EventArchiveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventArchive)(nil))
}

func (o EventArchiveMapOutput) ToEventArchiveMapOutput() EventArchiveMapOutput {
	return o
}

func (o EventArchiveMapOutput) ToEventArchiveMapOutputWithContext(ctx context.Context) EventArchiveMapOutput {
	return o
}

func (o EventArchiveMapOutput) MapIndex(k pulumi.StringInput) EventArchiveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventArchive {
		return vs[0].(map[string]EventArchive)[vs[1].(string)]
	}).(EventArchiveOutput)
}

func init() {
	pulumi.RegisterOutputType(EventArchiveOutput{})
	pulumi.RegisterOutputType(EventArchivePtrOutput{})
	pulumi.RegisterOutputType(EventArchiveArrayOutput{})
	pulumi.RegisterOutputType(EventArchiveMapOutput{})
}
