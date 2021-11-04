// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Application AutoScaling Policy resource.
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
// 		dynamodbTableReadTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTableReadTarget", &appautoscaling.TargetArgs{
// 			MaxCapacity:       pulumi.Int(100),
// 			MinCapacity:       pulumi.Int(5),
// 			ResourceId:        pulumi.String("table/tableName"),
// 			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
// 			ServiceNamespace:  pulumi.String("dynamodb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appautoscaling.NewPolicy(ctx, "dynamodbTableReadPolicy", &appautoscaling.PolicyArgs{
// 			PolicyType:        pulumi.String("TargetTrackingScaling"),
// 			ResourceId:        dynamodbTableReadTarget.ResourceId,
// 			ScalableDimension: dynamodbTableReadTarget.ScalableDimension,
// 			ServiceNamespace:  dynamodbTableReadTarget.ServiceNamespace,
// 			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
// 				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
// 					PredefinedMetricType: pulumi.String("DynamoDBReadCapacityUtilization"),
// 				},
// 				TargetValue: pulumi.Float64(70),
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
// 		_, err = appautoscaling.NewPolicy(ctx, "ecsPolicy", &appautoscaling.PolicyArgs{
// 			PolicyType:        pulumi.String("StepScaling"),
// 			ResourceId:        ecsTarget.ResourceId,
// 			ScalableDimension: ecsTarget.ScalableDimension,
// 			ServiceNamespace:  ecsTarget.ServiceNamespace,
// 			StepScalingPolicyConfiguration: &appautoscaling.PolicyStepScalingPolicyConfigurationArgs{
// 				AdjustmentType:        pulumi.String("ChangeInCapacity"),
// 				Cooldown:              pulumi.Int(60),
// 				MetricAggregationType: pulumi.String("Maximum"),
// 				StepAdjustments: appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArray{
// 					&appautoscaling.PolicyStepScalingPolicyConfigurationStepAdjustmentArgs{
// 						MetricIntervalUpperBound: pulumi.String("0"),
// 						ScalingAdjustment:        -1,
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
// ### Preserve desired count when updating an autoscaled ECS Service
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecs.NewService(ctx, "ecsService", &ecs.ServiceArgs{
// 			Cluster:        pulumi.String("clusterName"),
// 			TaskDefinition: pulumi.String("taskDefinitionFamily:1"),
// 			DesiredCount:   pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Aurora Read Replica Autoscaling
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appautoscaling"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		replicasTarget, err := appautoscaling.NewTarget(ctx, "replicasTarget", &appautoscaling.TargetArgs{
// 			ServiceNamespace:  pulumi.String("rds"),
// 			ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
// 			ResourceId:        pulumi.String(fmt.Sprintf("%v%v", "cluster:", aws_rds_cluster.Example.Id)),
// 			MinCapacity:       pulumi.Int(1),
// 			MaxCapacity:       pulumi.Int(15),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appautoscaling.NewPolicy(ctx, "replicasPolicy", &appautoscaling.PolicyArgs{
// 			ServiceNamespace:  replicasTarget.ServiceNamespace,
// 			ScalableDimension: replicasTarget.ScalableDimension,
// 			ResourceId:        replicasTarget.ResourceId,
// 			PolicyType:        pulumi.String("TargetTrackingScaling"),
// 			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
// 				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
// 					PredefinedMetricType: pulumi.String("RDSReaderAverageCPUUtilization"),
// 				},
// 				TargetValue:      pulumi.Float64(75),
// 				ScaleInCooldown:  pulumi.Int(300),
// 				ScaleOutCooldown: pulumi.Int(300),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### MSK / Kafka Autoscaling
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
// 		mskTarget, err := appautoscaling.NewTarget(ctx, "mskTarget", &appautoscaling.TargetArgs{
// 			ServiceNamespace:  pulumi.String("kafka"),
// 			ScalableDimension: pulumi.String("kafka:broker-storage:VolumeSize"),
// 			ResourceId:        pulumi.Any(aws_msk_cluster.Example.Arn),
// 			MinCapacity:       pulumi.Int(1),
// 			MaxCapacity:       pulumi.Int(8),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appautoscaling.NewPolicy(ctx, "targets", &appautoscaling.PolicyArgs{
// 			ServiceNamespace:  mskTarget.ServiceNamespace,
// 			ScalableDimension: mskTarget.ScalableDimension,
// 			ResourceId:        mskTarget.ResourceId,
// 			PolicyType:        pulumi.String("TargetTrackingScaling"),
// 			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
// 				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
// 					PredefinedMetricType: pulumi.String("KafkaBrokerStorageUtilization"),
// 				},
// 				TargetValue: pulumi.Float64(55),
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
// Application AutoScaling Policy can be imported using the `service-namespace` , `resource-id`, `scalable-dimension` and `policy-name` separated by `/`.
//
// ```sh
//  $ pulumi import aws:appautoscaling/policy:Policy test-policy service-namespace/resource-id/scalable-dimension/policy-name
// ```
type Policy struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the policy. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration PolicyStepScalingPolicyConfigurationPtrOutput `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration PolicyTargetTrackingScalingPolicyConfigurationPtrOutput `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ScalableDimension == nil {
		return nil, errors.New("invalid value for required argument 'ScalableDimension'")
	}
	if args.ServiceNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServiceNamespace'")
	}
	var resource Policy
	err := ctx.RegisterResource("aws:appautoscaling/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws:appautoscaling/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The ARN assigned by AWS to the scaling policy.
	Arn *string `pulumi:"arn"`
	// The name of the policy. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType *string `pulumi:"policyType"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId *string `pulumi:"resourceId"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension *string `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace *string `pulumi:"serviceNamespace"`
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration *PolicyStepScalingPolicyConfiguration `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration *PolicyTargetTrackingScalingPolicyConfiguration `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

type PolicyState struct {
	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringPtrInput
	// The name of the policy. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType pulumi.StringPtrInput
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringPtrInput
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringPtrInput
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringPtrInput
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration PolicyStepScalingPolicyConfigurationPtrInput
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration PolicyTargetTrackingScalingPolicyConfigurationPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// The name of the policy. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType *string `pulumi:"policyType"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId string `pulumi:"resourceId"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension string `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace string `pulumi:"serviceNamespace"`
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration *PolicyStepScalingPolicyConfiguration `pulumi:"stepScalingPolicyConfiguration"`
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration *PolicyTargetTrackingScalingPolicyConfiguration `pulumi:"targetTrackingScalingPolicyConfiguration"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The name of the policy. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
	PolicyType pulumi.StringPtrInput
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringInput
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringInput
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringInput
	// Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
	StepScalingPolicyConfiguration PolicyStepScalingPolicyConfigurationPtrInput
	// A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
	TargetTrackingScalingPolicyConfiguration PolicyTargetTrackingScalingPolicyConfigurationPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i *Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//          PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//          PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}

type PolicyPtrOutput struct{ *pulumi.OutputState }

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) Elem() PolicyOutput {
	return o.ApplyT(func(v *Policy) Policy {
		if v != nil {
			return *v
		}
		var ret Policy
		return ret
	}).(PolicyOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Policy {
		return vs[0].([]Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Policy {
		return vs[0].(map[string]Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPtrInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
