// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Pause Cluster Action
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": [\n", "          \"scheduler.redshift.amazonaws.com\"\n", "        ]\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePolicy, err := iam.NewPolicy(ctx, "examplePolicy", &iam.PolicyArgs{
// 			Policy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "      {\n", "          \"Sid\": \"VisualEditor0\",\n", "          \"Effect\": \"Allow\",\n", "          \"Action\": [\n", "              \"redshift:PauseCluster\",\n", "              \"redshift:ResumeCluster\",\n", "              \"redshift:ResizeCluster\"\n", "          ],\n", "          \"Resource\": \"*\"\n", "      }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: examplePolicy.Arn,
// 			Role:      exampleRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redshift.NewScheduledAction(ctx, "exampleScheduledAction", &redshift.ScheduledActionArgs{
// 			Schedule: pulumi.String("cron(00 23 * * ? *)"),
// 			IamRole:  exampleRole.Arn,
// 			TargetAction: &redshift.ScheduledActionTargetActionArgs{
// 				PauseCluster: &redshift.ScheduledActionTargetActionPauseClusterArgs{
// 					ClusterIdentifier: pulumi.String("tf-redshift001"),
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
// ### Resize Cluster Action
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.NewScheduledAction(ctx, "example", &redshift.ScheduledActionArgs{
// 			Schedule: pulumi.String("cron(00 23 * * ? *)"),
// 			IamRole:  pulumi.Any(aws_iam_role.Example.Arn),
// 			TargetAction: &redshift.ScheduledActionTargetActionArgs{
// 				ResizeCluster: &redshift.ScheduledActionTargetActionResizeClusterArgs{
// 					ClusterIdentifier: pulumi.String("tf-redshift001"),
// 					ClusterType:       pulumi.String("multi-node"),
// 					NodeType:          pulumi.String("dc1.large"),
// 					NumberOfNodes:     pulumi.Int(2),
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
// Redshift Scheduled Action can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:redshift/scheduledAction:ScheduledAction example tf-redshift-scheduled-action
// ```
type ScheduledAction struct {
	pulumi.CustomResourceState

	// The description of the scheduled action.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to enable the scheduled action. Default is `true` .
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The IAM role to assume to run the scheduled action.
	IamRole pulumi.StringOutput `pulumi:"iamRole"`
	// The scheduled action name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// Target action. Documented below.
	TargetAction ScheduledActionTargetActionOutput `pulumi:"targetAction"`
}

// NewScheduledAction registers a new resource with the given unique name, arguments, and options.
func NewScheduledAction(ctx *pulumi.Context,
	name string, args *ScheduledActionArgs, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IamRole == nil {
		return nil, errors.New("invalid value for required argument 'IamRole'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.TargetAction == nil {
		return nil, errors.New("invalid value for required argument 'TargetAction'")
	}
	var resource ScheduledAction
	err := ctx.RegisterResource("aws:redshift/scheduledAction:ScheduledAction", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:redshift/scheduledAction:ScheduledAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledAction resources.
type scheduledActionState struct {
	// The description of the scheduled action.
	Description *string `pulumi:"description"`
	// Whether to enable the scheduled action. Default is `true` .
	Enable *bool `pulumi:"enable"`
	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *string `pulumi:"endTime"`
	// The IAM role to assume to run the scheduled action.
	IamRole *string `pulumi:"iamRole"`
	// The scheduled action name.
	Name *string `pulumi:"name"`
	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule *string `pulumi:"schedule"`
	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *string `pulumi:"startTime"`
	// Target action. Documented below.
	TargetAction *ScheduledActionTargetAction `pulumi:"targetAction"`
}

type ScheduledActionState struct {
	// The description of the scheduled action.
	Description pulumi.StringPtrInput
	// Whether to enable the scheduled action. Default is `true` .
	Enable pulumi.BoolPtrInput
	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime pulumi.StringPtrInput
	// The IAM role to assume to run the scheduled action.
	IamRole pulumi.StringPtrInput
	// The scheduled action name.
	Name pulumi.StringPtrInput
	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule pulumi.StringPtrInput
	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime pulumi.StringPtrInput
	// Target action. Documented below.
	TargetAction ScheduledActionTargetActionPtrInput
}

func (ScheduledActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionState)(nil)).Elem()
}

type scheduledActionArgs struct {
	// The description of the scheduled action.
	Description *string `pulumi:"description"`
	// Whether to enable the scheduled action. Default is `true` .
	Enable *bool `pulumi:"enable"`
	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *string `pulumi:"endTime"`
	// The IAM role to assume to run the scheduled action.
	IamRole string `pulumi:"iamRole"`
	// The scheduled action name.
	Name *string `pulumi:"name"`
	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule string `pulumi:"schedule"`
	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *string `pulumi:"startTime"`
	// Target action. Documented below.
	TargetAction ScheduledActionTargetAction `pulumi:"targetAction"`
}

// The set of arguments for constructing a ScheduledAction resource.
type ScheduledActionArgs struct {
	// The description of the scheduled action.
	Description pulumi.StringPtrInput
	// Whether to enable the scheduled action. Default is `true` .
	Enable pulumi.BoolPtrInput
	// The end time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime pulumi.StringPtrInput
	// The IAM role to assume to run the scheduled action.
	IamRole pulumi.StringInput
	// The scheduled action name.
	Name pulumi.StringPtrInput
	// The schedule of action. The schedule is defined format of "at expression" or "cron expression", for example `at(2016-03-04T17:27:00)` or `cron(0 10 ? * MON *)`. See [Scheduled Action](https://docs.aws.amazon.com/redshift/latest/APIReference/API_ScheduledAction.html) for more information.
	Schedule pulumi.StringInput
	// The start time in UTC when the schedule is active, in UTC RFC3339 format(for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime pulumi.StringPtrInput
	// Target action. Documented below.
	TargetAction ScheduledActionTargetActionInput
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

type ScheduledActionOutput struct{ *pulumi.OutputState }

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
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledAction) *ScheduledAction {
		return &v
	}).(ScheduledActionPtrOutput)
}

type ScheduledActionPtrOutput struct{ *pulumi.OutputState }

func (ScheduledActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil))
}

func (o ScheduledActionPtrOutput) ToScheduledActionPtrOutput() ScheduledActionPtrOutput {
	return o
}

func (o ScheduledActionPtrOutput) ToScheduledActionPtrOutputWithContext(ctx context.Context) ScheduledActionPtrOutput {
	return o
}

func (o ScheduledActionPtrOutput) Elem() ScheduledActionOutput {
	return o.ApplyT(func(v *ScheduledAction) ScheduledAction {
		if v != nil {
			return *v
		}
		var ret ScheduledAction
		return ret
	}).(ScheduledActionOutput)
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
