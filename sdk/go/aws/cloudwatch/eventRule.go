// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an EventBridge Rule resource.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
// 			Description:  pulumi.String("Capture each AWS Console Sign In"),
// 			EventPattern: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "{\n", "  \"detail-type\": [\n", "    \"AWS Console Sign In via CloudTrail\"\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		awsLogins, err := sns.NewTopic(ctx, "awsLogins", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewEventTarget(ctx, "sns", &cloudwatch.EventTargetArgs{
// 			Rule: console.Name,
// 			Arn:  awsLogins.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sns.NewTopicPolicy(ctx, "_default", &sns.TopicPolicyArgs{
// 			Arn: awsLogins.Arn,
// 			Policy: snsTopicPolicy.ApplyT(func(snsTopicPolicy iam.GetPolicyDocumentResult) (string, error) {
// 				return snsTopicPolicy.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EventRule struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The event bus to associate with this rule. If you omit this, the `default` event bus is used.
	EventBusName pulumi.StringPtrOutput `pulumi:"eventBusName"`
	// The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required.
	// See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
	EventPattern pulumi.StringPtrOutput `pulumi:"eventPattern"`
	// Whether the rule should be enabled (defaults to `true`).
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The scheduling expression.
	// For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus.
	ScheduleExpression pulumi.StringPtrOutput `pulumi:"scheduleExpression"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEventRule registers a new resource with the given unique name, arguments, and options.
func NewEventRule(ctx *pulumi.Context,
	name string, args *EventRuleArgs, opts ...pulumi.ResourceOption) (*EventRule, error) {
	if args == nil {
		args = &EventRuleArgs{}
	}
	var resource EventRule
	err := ctx.RegisterResource("aws:cloudwatch/eventRule:EventRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventRule gets an existing EventRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventRuleState, opts ...pulumi.ResourceOption) (*EventRule, error) {
	var resource EventRule
	err := ctx.ReadResource("aws:cloudwatch/eventRule:EventRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventRule resources.
type eventRuleState struct {
	// The Amazon Resource Name (ARN) of the rule.
	Arn *string `pulumi:"arn"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The event bus to associate with this rule. If you omit this, the `default` event bus is used.
	EventBusName *string `pulumi:"eventBusName"`
	// The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required.
	// See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
	EventPattern *string `pulumi:"eventPattern"`
	// Whether the rule should be enabled (defaults to `true`).
	IsEnabled *bool `pulumi:"isEnabled"`
	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn *string `pulumi:"roleArn"`
	// The scheduling expression.
	// For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus.
	ScheduleExpression *string `pulumi:"scheduleExpression"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EventRuleState struct {
	// The Amazon Resource Name (ARN) of the rule.
	Arn pulumi.StringPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The event bus to associate with this rule. If you omit this, the `default` event bus is used.
	EventBusName pulumi.StringPtrInput
	// The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required.
	// See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
	EventPattern pulumi.StringPtrInput
	// Whether the rule should be enabled (defaults to `true`).
	IsEnabled pulumi.BoolPtrInput
	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn pulumi.StringPtrInput
	// The scheduling expression.
	// For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus.
	ScheduleExpression pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventRuleState)(nil)).Elem()
}

type eventRuleArgs struct {
	// The description of the rule.
	Description *string `pulumi:"description"`
	// The event bus to associate with this rule. If you omit this, the `default` event bus is used.
	EventBusName *string `pulumi:"eventBusName"`
	// The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required.
	// See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
	EventPattern *string `pulumi:"eventPattern"`
	// Whether the rule should be enabled (defaults to `true`).
	IsEnabled *bool `pulumi:"isEnabled"`
	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn *string `pulumi:"roleArn"`
	// The scheduling expression.
	// For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus.
	ScheduleExpression *string `pulumi:"scheduleExpression"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventRule resource.
type EventRuleArgs struct {
	// The description of the rule.
	Description pulumi.StringPtrInput
	// The event bus to associate with this rule. If you omit this, the `default` event bus is used.
	EventBusName pulumi.StringPtrInput
	// The event pattern described a JSON object. At least one of `scheduleExpression` or `eventPattern` is required.
	// See full documentation of [Events and Event Patterns in EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) for details.
	EventPattern pulumi.StringPtrInput
	// Whether the rule should be enabled (defaults to `true`).
	IsEnabled pulumi.BoolPtrInput
	// The name of the rule. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
	RoleArn pulumi.StringPtrInput
	// The scheduling expression.
	// For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`. At least one of `scheduleExpression` or `eventPattern` is required. Can only be used on the default event bus.
	ScheduleExpression pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventRuleArgs)(nil)).Elem()
}
