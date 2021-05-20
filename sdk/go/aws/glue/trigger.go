// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Glue Trigger resource.
//
// ## Example Usage
// ### Conditional Trigger
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
// 			Type: pulumi.String("CONDITIONAL"),
// 			Actions: glue.TriggerActionArray{
// 				&glue.TriggerActionArgs{
// 					JobName: pulumi.Any(aws_glue_job.Example1.Name),
// 				},
// 			},
// 			Predicate: &glue.TriggerPredicateArgs{
// 				Conditions: glue.TriggerPredicateConditionArray{
// 					&glue.TriggerPredicateConditionArgs{
// 						JobName: pulumi.Any(aws_glue_job.Example2.Name),
// 						State:   pulumi.String("SUCCEEDED"),
// 					},
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
// ### On-Demand Trigger
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
// 			Type: pulumi.String("ON_DEMAND"),
// 			Actions: glue.TriggerActionArray{
// 				&glue.TriggerActionArgs{
// 					JobName: pulumi.Any(aws_glue_job.Example.Name),
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
// ### Scheduled Trigger
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
// 			Schedule: pulumi.String("cron(15 12 * * ? *)"),
// 			Type:     pulumi.String("SCHEDULED"),
// 			Actions: glue.TriggerActionArray{
// 				&glue.TriggerActionArgs{
// 					JobName: pulumi.Any(aws_glue_job.Example.Name),
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
// ### Conditional Trigger with Crawler Action
//
// **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
// 			Type: pulumi.String("CONDITIONAL"),
// 			Actions: glue.TriggerActionArray{
// 				&glue.TriggerActionArgs{
// 					CrawlerName: pulumi.Any(aws_glue_crawler.Example1.Name),
// 				},
// 			},
// 			Predicate: &glue.TriggerPredicateArgs{
// 				Conditions: glue.TriggerPredicateConditionArray{
// 					&glue.TriggerPredicateConditionArgs{
// 						JobName: pulumi.Any(aws_glue_job.Example2.Name),
// 						State:   pulumi.String("SUCCEEDED"),
// 					},
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
// ### Conditional Trigger with Crawler Condition
//
// **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewTrigger(ctx, "example", &glue.TriggerArgs{
// 			Type: pulumi.String("CONDITIONAL"),
// 			Actions: glue.TriggerActionArray{
// 				&glue.TriggerActionArgs{
// 					JobName: pulumi.Any(aws_glue_job.Example1.Name),
// 				},
// 			},
// 			Predicate: &glue.TriggerPredicateArgs{
// 				Conditions: glue.TriggerPredicateConditionArray{
// 					&glue.TriggerPredicateConditionArgs{
// 						CrawlerName: pulumi.Any(aws_glue_crawler.Example2.Name),
// 						CrawlState:  pulumi.String("SUCCEEDED"),
// 					},
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
// Glue Triggers can be imported using `name`, e.g.
//
// ```sh
//  $ pulumi import aws:glue/trigger:Trigger MyTrigger MyTrigger
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayOutput `pulumi:"actions"`
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the new trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the trigger.
	Name pulumi.StringOutput `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrOutput `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringOutput `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrOutput `pulumi:"workflowName"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Trigger
	err := ctx.RegisterResource("aws:glue/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("aws:glue/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []TriggerAction `pulumi:"actions"`
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn *string `pulumi:"arn"`
	// A description of the new trigger.
	Description *string `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the trigger.
	Name *string `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate *TriggerPredicate `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule *string `pulumi:"schedule"`
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
	Type *string `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName *string `pulumi:"workflowName"`
}

type TriggerState struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayInput
	// Amazon Resource Name (ARN) of Glue Trigger
	Arn pulumi.StringPtrInput
	// A description of the new trigger.
	Description pulumi.StringPtrInput
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the trigger.
	Name pulumi.StringPtrInput
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrInput
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrInput
	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `jobName` must also be specified. Conflicts with `crawlerState`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringPtrInput
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions []TriggerAction `pulumi:"actions"`
	// A description of the new trigger.
	Description *string `pulumi:"description"`
	// Start the trigger. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the trigger.
	Name *string `pulumi:"name"`
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate *TriggerPredicate `pulumi:"predicate"`
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule *string `pulumi:"schedule"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
	Type string `pulumi:"type"`
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// List of actions initiated by this trigger when it fires. See Actions Below.
	Actions TriggerActionArrayInput
	// A description of the new trigger.
	Description pulumi.StringPtrInput
	// Start the trigger. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the trigger.
	Name pulumi.StringPtrInput
	// A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. See Predicate Below.
	Predicate TriggerPredicatePtrInput
	// A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
	Schedule pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
	Type pulumi.StringInput
	// A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
	WorkflowName pulumi.StringPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

func (i *Trigger) ToTriggerPtrOutput() TriggerPtrOutput {
	return i.ToTriggerPtrOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPtrOutput)
}

type TriggerPtrInput interface {
	pulumi.Input

	ToTriggerPtrOutput() TriggerPtrOutput
	ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput
}

type triggerPtrType TriggerArgs

func (*triggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil))
}

func (i *triggerPtrType) ToTriggerPtrOutput() TriggerPtrOutput {
	return i.ToTriggerPtrOutputWithContext(context.Background())
}

func (i *triggerPtrType) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPtrOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//          TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Trigger)(nil))
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//          TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Trigger)(nil))
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct {
	*pulumi.OutputState
}

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerPtrOutput() TriggerPtrOutput {
	return o.ToTriggerPtrOutputWithContext(context.Background())
}

func (o TriggerOutput) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return o.ApplyT(func(v Trigger) *Trigger {
		return &v
	}).(TriggerPtrOutput)
}

type TriggerPtrOutput struct {
	*pulumi.OutputState
}

func (TriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil))
}

func (o TriggerPtrOutput) ToTriggerPtrOutput() TriggerPtrOutput {
	return o
}

func (o TriggerPtrOutput) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return o
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Trigger)(nil))
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Trigger {
		return vs[0].([]Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Trigger)(nil))
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Trigger {
		return vs[0].(map[string]Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerPtrOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
