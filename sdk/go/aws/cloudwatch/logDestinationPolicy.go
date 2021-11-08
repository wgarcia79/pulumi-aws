// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Logs destination policy resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testDestination, err := cloudwatch.NewLogDestination(ctx, "testDestination", &cloudwatch.LogDestinationArgs{
// 			RoleArn:   pulumi.Any(aws_iam_role.Iam_for_cloudwatch.Arn),
// 			TargetArn: pulumi.Any(aws_kinesis_stream.Kinesis_for_cloudwatch.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogDestinationPolicy(ctx, "testDestinationPolicyLogDestinationPolicy", &cloudwatch.LogDestinationPolicyArgs{
// 			DestinationName: testDestination.Name,
// 			AccessPolicy: testDestinationPolicyPolicyDocument.ApplyT(func(testDestinationPolicyPolicyDocument iam.GetPolicyDocumentResult) (string, error) {
// 				return testDestinationPolicyPolicyDocument.Json, nil
// 			}).(pulumi.StringOutput),
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
// CloudWatch Logs destination policies can be imported using the `destination_name`, e.g.,
//
// ```sh
//  $ pulumi import aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy test_destination_policy test_destination
// ```
type LogDestinationPolicy struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringOutput `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName pulumi.StringOutput `pulumi:"destinationName"`
}

// NewLogDestinationPolicy registers a new resource with the given unique name, arguments, and options.
func NewLogDestinationPolicy(ctx *pulumi.Context,
	name string, args *LogDestinationPolicyArgs, opts ...pulumi.ResourceOption) (*LogDestinationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicy == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicy'")
	}
	if args.DestinationName == nil {
		return nil, errors.New("invalid value for required argument 'DestinationName'")
	}
	var resource LogDestinationPolicy
	err := ctx.RegisterResource("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDestinationPolicy gets an existing LogDestinationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDestinationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDestinationPolicyState, opts ...pulumi.ResourceOption) (*LogDestinationPolicy, error) {
	var resource LogDestinationPolicy
	err := ctx.ReadResource("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDestinationPolicy resources.
type logDestinationPolicyState struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy *string `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName *string `pulumi:"destinationName"`
}

type LogDestinationPolicyState struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringPtrInput
	// A name for the subscription filter
	DestinationName pulumi.StringPtrInput
}

func (LogDestinationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationPolicyState)(nil)).Elem()
}

type logDestinationPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy string `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName string `pulumi:"destinationName"`
}

// The set of arguments for constructing a LogDestinationPolicy resource.
type LogDestinationPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringInput
	// A name for the subscription filter
	DestinationName pulumi.StringInput
}

func (LogDestinationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationPolicyArgs)(nil)).Elem()
}

type LogDestinationPolicyInput interface {
	pulumi.Input

	ToLogDestinationPolicyOutput() LogDestinationPolicyOutput
	ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput
}

func (*LogDestinationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDestinationPolicy)(nil))
}

func (i *LogDestinationPolicy) ToLogDestinationPolicyOutput() LogDestinationPolicyOutput {
	return i.ToLogDestinationPolicyOutputWithContext(context.Background())
}

func (i *LogDestinationPolicy) ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyOutput)
}

func (i *LogDestinationPolicy) ToLogDestinationPolicyPtrOutput() LogDestinationPolicyPtrOutput {
	return i.ToLogDestinationPolicyPtrOutputWithContext(context.Background())
}

func (i *LogDestinationPolicy) ToLogDestinationPolicyPtrOutputWithContext(ctx context.Context) LogDestinationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyPtrOutput)
}

type LogDestinationPolicyPtrInput interface {
	pulumi.Input

	ToLogDestinationPolicyPtrOutput() LogDestinationPolicyPtrOutput
	ToLogDestinationPolicyPtrOutputWithContext(ctx context.Context) LogDestinationPolicyPtrOutput
}

type logDestinationPolicyPtrType LogDestinationPolicyArgs

func (*logDestinationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDestinationPolicy)(nil))
}

func (i *logDestinationPolicyPtrType) ToLogDestinationPolicyPtrOutput() LogDestinationPolicyPtrOutput {
	return i.ToLogDestinationPolicyPtrOutputWithContext(context.Background())
}

func (i *logDestinationPolicyPtrType) ToLogDestinationPolicyPtrOutputWithContext(ctx context.Context) LogDestinationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyPtrOutput)
}

// LogDestinationPolicyArrayInput is an input type that accepts LogDestinationPolicyArray and LogDestinationPolicyArrayOutput values.
// You can construct a concrete instance of `LogDestinationPolicyArrayInput` via:
//
//          LogDestinationPolicyArray{ LogDestinationPolicyArgs{...} }
type LogDestinationPolicyArrayInput interface {
	pulumi.Input

	ToLogDestinationPolicyArrayOutput() LogDestinationPolicyArrayOutput
	ToLogDestinationPolicyArrayOutputWithContext(context.Context) LogDestinationPolicyArrayOutput
}

type LogDestinationPolicyArray []LogDestinationPolicyInput

func (LogDestinationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogDestinationPolicy)(nil)).Elem()
}

func (i LogDestinationPolicyArray) ToLogDestinationPolicyArrayOutput() LogDestinationPolicyArrayOutput {
	return i.ToLogDestinationPolicyArrayOutputWithContext(context.Background())
}

func (i LogDestinationPolicyArray) ToLogDestinationPolicyArrayOutputWithContext(ctx context.Context) LogDestinationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyArrayOutput)
}

// LogDestinationPolicyMapInput is an input type that accepts LogDestinationPolicyMap and LogDestinationPolicyMapOutput values.
// You can construct a concrete instance of `LogDestinationPolicyMapInput` via:
//
//          LogDestinationPolicyMap{ "key": LogDestinationPolicyArgs{...} }
type LogDestinationPolicyMapInput interface {
	pulumi.Input

	ToLogDestinationPolicyMapOutput() LogDestinationPolicyMapOutput
	ToLogDestinationPolicyMapOutputWithContext(context.Context) LogDestinationPolicyMapOutput
}

type LogDestinationPolicyMap map[string]LogDestinationPolicyInput

func (LogDestinationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogDestinationPolicy)(nil)).Elem()
}

func (i LogDestinationPolicyMap) ToLogDestinationPolicyMapOutput() LogDestinationPolicyMapOutput {
	return i.ToLogDestinationPolicyMapOutputWithContext(context.Background())
}

func (i LogDestinationPolicyMap) ToLogDestinationPolicyMapOutputWithContext(ctx context.Context) LogDestinationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyMapOutput)
}

type LogDestinationPolicyOutput struct{ *pulumi.OutputState }

func (LogDestinationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDestinationPolicy)(nil))
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyOutput() LogDestinationPolicyOutput {
	return o
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput {
	return o
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyPtrOutput() LogDestinationPolicyPtrOutput {
	return o.ToLogDestinationPolicyPtrOutputWithContext(context.Background())
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyPtrOutputWithContext(ctx context.Context) LogDestinationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogDestinationPolicy) *LogDestinationPolicy {
		return &v
	}).(LogDestinationPolicyPtrOutput)
}

type LogDestinationPolicyPtrOutput struct{ *pulumi.OutputState }

func (LogDestinationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogDestinationPolicy)(nil))
}

func (o LogDestinationPolicyPtrOutput) ToLogDestinationPolicyPtrOutput() LogDestinationPolicyPtrOutput {
	return o
}

func (o LogDestinationPolicyPtrOutput) ToLogDestinationPolicyPtrOutputWithContext(ctx context.Context) LogDestinationPolicyPtrOutput {
	return o
}

func (o LogDestinationPolicyPtrOutput) Elem() LogDestinationPolicyOutput {
	return o.ApplyT(func(v *LogDestinationPolicy) LogDestinationPolicy {
		if v != nil {
			return *v
		}
		var ret LogDestinationPolicy
		return ret
	}).(LogDestinationPolicyOutput)
}

type LogDestinationPolicyArrayOutput struct{ *pulumi.OutputState }

func (LogDestinationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogDestinationPolicy)(nil))
}

func (o LogDestinationPolicyArrayOutput) ToLogDestinationPolicyArrayOutput() LogDestinationPolicyArrayOutput {
	return o
}

func (o LogDestinationPolicyArrayOutput) ToLogDestinationPolicyArrayOutputWithContext(ctx context.Context) LogDestinationPolicyArrayOutput {
	return o
}

func (o LogDestinationPolicyArrayOutput) Index(i pulumi.IntInput) LogDestinationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogDestinationPolicy {
		return vs[0].([]LogDestinationPolicy)[vs[1].(int)]
	}).(LogDestinationPolicyOutput)
}

type LogDestinationPolicyMapOutput struct{ *pulumi.OutputState }

func (LogDestinationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogDestinationPolicy)(nil))
}

func (o LogDestinationPolicyMapOutput) ToLogDestinationPolicyMapOutput() LogDestinationPolicyMapOutput {
	return o
}

func (o LogDestinationPolicyMapOutput) ToLogDestinationPolicyMapOutputWithContext(ctx context.Context) LogDestinationPolicyMapOutput {
	return o
}

func (o LogDestinationPolicyMapOutput) MapIndex(k pulumi.StringInput) LogDestinationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogDestinationPolicy {
		return vs[0].(map[string]LogDestinationPolicy)[vs[1].(string)]
	}).(LogDestinationPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogDestinationPolicyInput)(nil)).Elem(), &LogDestinationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDestinationPolicyPtrInput)(nil)).Elem(), &LogDestinationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDestinationPolicyArrayInput)(nil)).Elem(), LogDestinationPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogDestinationPolicyMapInput)(nil)).Elem(), LogDestinationPolicyMap{})
	pulumi.RegisterOutputType(LogDestinationPolicyOutput{})
	pulumi.RegisterOutputType(LogDestinationPolicyPtrOutput{})
	pulumi.RegisterOutputType(LogDestinationPolicyArrayOutput{})
	pulumi.RegisterOutputType(LogDestinationPolicyMapOutput{})
}
