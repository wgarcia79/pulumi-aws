// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Application AutoScaling ScheduledAction resource.
//
// ## Example Usage
// ### DynamoDB Table Autoscaling
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appautoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dynamodbTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTarget", &appautoscaling.TargetArgs{
// 			MaxCapacity:       pulumi.Int(100),
// 			MinCapacity:       pulumi.Int(5),
// 			ResourceId:        pulumi.String("table/tableName"),
// 			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
// 			ServiceNamespace:  pulumi.String("dynamodb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appautoscaling.NewScheduledAction(ctx, "dynamodbScheduledAction", &appautoscaling.ScheduledActionArgs{
// 			ServiceNamespace:  dynamodbTarget.ServiceNamespace,
// 			ResourceId:        dynamodbTarget.ResourceId,
// 			ScalableDimension: dynamodbTarget.ScalableDimension,
// 			Schedule:          pulumi.String("at(2006-01-02T15:04:05)"),
// 			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
// 				MinCapacity: pulumi.Int(1),
// 				MaxCapacity: pulumi.Int(200),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### ECS Service Autoscaling
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appautoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ecsTarget, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
// 			MaxCapacity:       pulumi.Int(4),
// 			MinCapacity:       pulumi.Int(1),
// 			ResourceId:        pulumi.String("service/clusterName/serviceName"),
// 			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
// 			ServiceNamespace:  pulumi.String("ecs"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appautoscaling.NewScheduledAction(ctx, "ecsScheduledAction", &appautoscaling.ScheduledActionArgs{
// 			ServiceNamespace:  ecsTarget.ServiceNamespace,
// 			ResourceId:        ecsTarget.ResourceId,
// 			ScalableDimension: ecsTarget.ScalableDimension,
// 			Schedule:          pulumi.String("at(2006-01-02T15:04:05)"),
// 			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
// 				MinCapacity: pulumi.Int(1),
// 				MaxCapacity: pulumi.Int(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ScheduledAction struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The name of the scheduled action.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction ScheduledActionScalableTargetActionOutput `pulumi:"scalableTargetAction"`
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`
	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `startTime` and `endTime`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewScheduledAction registers a new resource with the given unique name, arguments, and options.
func NewScheduledAction(ctx *pulumi.Context,
	name string, args *ScheduledActionArgs, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ScalableDimension == nil {
		return nil, errors.New("invalid value for required argument 'ScalableDimension'")
	}
	if args.ScalableTargetAction == nil {
		return nil, errors.New("invalid value for required argument 'ScalableTargetAction'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.ServiceNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServiceNamespace'")
	}
	var resource ScheduledAction
	err := ctx.RegisterResource("aws:appautoscaling/scheduledAction:ScheduledAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledAction gets an existing ScheduledAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledActionState, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	var resource ScheduledAction
	err := ctx.ReadResource("aws:appautoscaling/scheduledAction:ScheduledAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledAction resources.
type scheduledActionState struct {
	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn *string `pulumi:"arn"`
	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime *string `pulumi:"endTime"`
	// The name of the scheduled action.
	Name *string `pulumi:"name"`
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId *string `pulumi:"resourceId"`
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension *string `pulumi:"scalableDimension"`
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction *ScheduledActionScalableTargetAction `pulumi:"scalableTargetAction"`
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule *string `pulumi:"schedule"`
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace *string `pulumi:"serviceNamespace"`
	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime *string `pulumi:"startTime"`
	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `startTime` and `endTime`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone *string `pulumi:"timezone"`
}

type ScheduledActionState struct {
	// The Amazon Resource Name (ARN) of the scheduled action.
	Arn pulumi.StringPtrInput
	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime pulumi.StringPtrInput
	// The name of the scheduled action.
	Name pulumi.StringPtrInput
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringPtrInput
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringPtrInput
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction ScheduledActionScalableTargetActionPtrInput
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringPtrInput
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringPtrInput
	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime pulumi.StringPtrInput
	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `startTime` and `endTime`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone pulumi.StringPtrInput
}

func (ScheduledActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionState)(nil)).Elem()
}

type scheduledActionArgs struct {
	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime *string `pulumi:"endTime"`
	// The name of the scheduled action.
	Name *string `pulumi:"name"`
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId string `pulumi:"resourceId"`
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension string `pulumi:"scalableDimension"`
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction ScheduledActionScalableTargetAction `pulumi:"scalableTargetAction"`
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule string `pulumi:"schedule"`
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace string `pulumi:"serviceNamespace"`
	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime *string `pulumi:"startTime"`
	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `startTime` and `endTime`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a ScheduledAction resource.
type ScheduledActionArgs struct {
	// The date and time for the scheduled action to end in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	EndTime pulumi.StringPtrInput
	// The name of the scheduled action.
	Name pulumi.StringPtrInput
	// The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
	ResourceId pulumi.StringInput
	// The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
	ScalableDimension pulumi.StringInput
	// The new minimum and maximum capacity. You can set both values or just one. See below
	ScalableTargetAction ScheduledActionScalableTargetActionInput
	// The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). Times for at expressions and cron expressions are evaluated using the time zone configured in `timezone`. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
	Schedule pulumi.StringInput
	// The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
	ServiceNamespace pulumi.StringInput
	// The date and time for the scheduled action to start in RFC 3339 format. The timezone is not affected by the setting of `timezone`.
	StartTime pulumi.StringPtrInput
	// The time zone used when setting a scheduled action by using an at or cron expression. Does not affect timezone for `startTime` and `endTime`. Valid values are the [canonical names of the IANA time zones supported by Joda-Time](https://www.joda.org/joda-time/timezones.html), such as `Etc/GMT+9` or `Pacific/Tahiti`. Default is `UTC`.
	Timezone pulumi.StringPtrInput
}

func (ScheduledActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionArgs)(nil)).Elem()
}

type ScheduledActionInput interface {
	pulumi.Input

	ToScheduledActionOutput() ScheduledActionOutput
	ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput
}

func (*ScheduledAction) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledAction)(nil))
}

func (i *ScheduledAction) ToScheduledActionOutput() ScheduledActionOutput {
	return i.ToScheduledActionOutputWithContext(context.Background())
}

func (i *ScheduledAction) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionOutput)
}

func (i *ScheduledAction) ToScheduledActionPtrOutput() ScheduledActionPtrOutput {
	return i.ToScheduledActionPtrOutputWithContext(context.Background())
}

func (i *ScheduledAction) ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionPtrOutput)
}

type ScheduledActionPtrInput interface {
	pulumi.Input

	ToScheduledActionPtrOutput() ScheduledActionPtrOutput
	ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput
}

type scheduledActionPtrType ScheduledActionArgs

func (*scheduledActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil))
}

func (i *scheduledActionPtrType) ToScheduledActionPtrOutput() ScheduledActionPtrOutput {
	return i.ToScheduledActionPtrOutputWithContext(context.Background())
}

func (i *scheduledActionPtrType) ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionPtrOutput)
}

// ScheduledActionArrayInput is an input type that accepts ScheduledActionArray and ScheduledActionArrayOutput values.
// You can construct a concrete instance of `ScheduledActionArrayInput` via:
//
//          ScheduledActionArray{ ScheduledActionArgs{...} }
type ScheduledActionArrayInput interface {
	pulumi.Input

	ToScheduledActionArrayOutput() ScheduledActionArrayOutput
	ToScheduledActionArrayOutputWithContext(context.Context) ScheduledActionArrayOutput
}

type ScheduledActionArray []ScheduledActionInput

func (ScheduledActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScheduledAction)(nil)).Elem()
}

func (i ScheduledActionArray) ToScheduledActionArrayOutput() ScheduledActionArrayOutput {
	return i.ToScheduledActionArrayOutputWithContext(context.Background())
}

func (i ScheduledActionArray) ToScheduledActionArrayOutputWithContext(ctx context.Context) ScheduledActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionArrayOutput)
}

// ScheduledActionMapInput is an input type that accepts ScheduledActionMap and ScheduledActionMapOutput values.
// You can construct a concrete instance of `ScheduledActionMapInput` via:
//
//          ScheduledActionMap{ "key": ScheduledActionArgs{...} }
type ScheduledActionMapInput interface {
	pulumi.Input

	ToScheduledActionMapOutput() ScheduledActionMapOutput
	ToScheduledActionMapOutputWithContext(context.Context) ScheduledActionMapOutput
}

type ScheduledActionMap map[string]ScheduledActionInput

func (ScheduledActionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScheduledAction)(nil)).Elem()
}

func (i ScheduledActionMap) ToScheduledActionMapOutput() ScheduledActionMapOutput {
	return i.ToScheduledActionMapOutputWithContext(context.Background())
}

func (i ScheduledActionMap) ToScheduledActionMapOutputWithContext(ctx context.Context) ScheduledActionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionMapOutput)
}

type ScheduledActionOutput struct {
	*pulumi.OutputState
}

func (ScheduledActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledAction)(nil))
}

func (o ScheduledActionOutput) ToScheduledActionOutput() ScheduledActionOutput {
	return o
}

func (o ScheduledActionOutput) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return o
}

func (o ScheduledActionOutput) ToScheduledActionPtrOutput() ScheduledActionPtrOutput {
	return o.ToScheduledActionPtrOutputWithContext(context.Background())
}

func (o ScheduledActionOutput) ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput {
	return o.ApplyT(func(v ScheduledAction) *ScheduledAction {
		return &v
	}).(ScheduledActionPtrOutput)
}

type ScheduledActionPtrOutput struct {
	*pulumi.OutputState
}

func (ScheduledActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil))
}

func (o ScheduledActionPtrOutput) ToScheduledActionPtrOutput() ScheduledActionPtrOutput {
	return o
}

func (o ScheduledActionPtrOutput) ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput {
	return o
}

type ScheduledActionArrayOutput struct{ *pulumi.OutputState }

func (ScheduledActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduledAction)(nil))
}

func (o ScheduledActionArrayOutput) ToScheduledActionArrayOutput() ScheduledActionArrayOutput {
	return o
}

func (o ScheduledActionArrayOutput) ToScheduledActionArrayOutputWithContext(ctx context.Context) ScheduledActionArrayOutput {
	return o
}

func (o ScheduledActionArrayOutput) Index(i pulumi.IntInput) ScheduledActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduledAction {
		return vs[0].([]ScheduledAction)[vs[1].(int)]
	}).(ScheduledActionOutput)
}

type ScheduledActionMapOutput struct{ *pulumi.OutputState }

func (ScheduledActionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ScheduledAction)(nil))
}

func (o ScheduledActionMapOutput) ToScheduledActionMapOutput() ScheduledActionMapOutput {
	return o
}

func (o ScheduledActionMapOutput) ToScheduledActionMapOutputWithContext(ctx context.Context) ScheduledActionMapOutput {
	return o
}

func (o ScheduledActionMapOutput) MapIndex(k pulumi.StringInput) ScheduledActionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ScheduledAction {
		return vs[0].(map[string]ScheduledAction)[vs[1].(string)]
	}).(ScheduledActionOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledActionOutput{})
	pulumi.RegisterOutputType(ScheduledActionPtrOutput{})
	pulumi.RegisterOutputType(ScheduledActionArrayOutput{})
	pulumi.RegisterOutputType(ScheduledActionMapOutput{})
}
