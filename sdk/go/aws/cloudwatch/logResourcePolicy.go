// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a CloudWatch log resource policy.
//
// ## Example Usage
// ### Elasticsearch Log Publishing
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
// 		elasticsearch_log_publishing_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"logs:CreateLogStream",
// 						"logs:PutLogEvents",
// 						"logs:PutLogEventsBatch",
// 					},
// 					Resources: []string{
// 						"arn:aws:logs:*",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Identifiers: []string{
// 								"es.amazonaws.com",
// 							},
// 							Type: "Service",
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogResourcePolicy(ctx, "elasticsearch_log_publishing_policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
// 			PolicyDocument: pulumi.String(elasticsearch_log_publishing_policyPolicyDocument.Json),
// 			PolicyName:     pulumi.String("elasticsearch-log-publishing-policy"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Route53 Query Logging
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
// 		route53_query_logging_policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"logs:CreateLogStream",
// 						"logs:PutLogEvents",
// 					},
// 					Resources: []string{
// 						"arn:aws:logs:*:*:log-group:/aws/route53/*",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Identifiers: []string{
// 								"route53.amazonaws.com",
// 							},
// 							Type: "Service",
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogResourcePolicy(ctx, "route53_query_logging_policyLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
// 			PolicyDocument: pulumi.String(route53_query_logging_policyPolicyDocument.Json),
// 			PolicyName:     pulumi.String("route53-query-logging-policy"),
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
// CloudWatch log resource policies can be imported using the policy name, e.g.,
//
// ```sh
//  $ pulumi import aws:cloudwatch/logResourcePolicy:LogResourcePolicy MyPolicy MyPolicy
// ```
type LogResourcePolicy struct {
	pulumi.CustomResourceState

	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// Name of the resource policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewLogResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewLogResourcePolicy(ctx *pulumi.Context,
	name string, args *LogResourcePolicyArgs, opts ...pulumi.ResourceOption) (*LogResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	var resource LogResourcePolicy
	err := ctx.RegisterResource("aws:cloudwatch/logResourcePolicy:LogResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogResourcePolicy gets an existing LogResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogResourcePolicyState, opts ...pulumi.ResourceOption) (*LogResourcePolicy, error) {
	var resource LogResourcePolicy
	err := ctx.ReadResource("aws:cloudwatch/logResourcePolicy:LogResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogResourcePolicy resources.
type logResourcePolicyState struct {
	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Name of the resource policy.
	PolicyName *string `pulumi:"policyName"`
}

type LogResourcePolicyState struct {
	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument pulumi.StringPtrInput
	// Name of the resource policy.
	PolicyName pulumi.StringPtrInput
}

func (LogResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*logResourcePolicyState)(nil)).Elem()
}

type logResourcePolicyArgs struct {
	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument string `pulumi:"policyDocument"`
	// Name of the resource policy.
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a LogResourcePolicy resource.
type LogResourcePolicyArgs struct {
	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument pulumi.StringInput
	// Name of the resource policy.
	PolicyName pulumi.StringInput
}

func (LogResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logResourcePolicyArgs)(nil)).Elem()
}

type LogResourcePolicyInput interface {
	pulumi.Input

	ToLogResourcePolicyOutput() LogResourcePolicyOutput
	ToLogResourcePolicyOutputWithContext(ctx context.Context) LogResourcePolicyOutput
}

func (*LogResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LogResourcePolicy)(nil))
}

func (i *LogResourcePolicy) ToLogResourcePolicyOutput() LogResourcePolicyOutput {
	return i.ToLogResourcePolicyOutputWithContext(context.Background())
}

func (i *LogResourcePolicy) ToLogResourcePolicyOutputWithContext(ctx context.Context) LogResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogResourcePolicyOutput)
}

func (i *LogResourcePolicy) ToLogResourcePolicyPtrOutput() LogResourcePolicyPtrOutput {
	return i.ToLogResourcePolicyPtrOutputWithContext(context.Background())
}

func (i *LogResourcePolicy) ToLogResourcePolicyPtrOutputWithContext(ctx context.Context) LogResourcePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogResourcePolicyPtrOutput)
}

type LogResourcePolicyPtrInput interface {
	pulumi.Input

	ToLogResourcePolicyPtrOutput() LogResourcePolicyPtrOutput
	ToLogResourcePolicyPtrOutputWithContext(ctx context.Context) LogResourcePolicyPtrOutput
}

type logResourcePolicyPtrType LogResourcePolicyArgs

func (*logResourcePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogResourcePolicy)(nil))
}

func (i *logResourcePolicyPtrType) ToLogResourcePolicyPtrOutput() LogResourcePolicyPtrOutput {
	return i.ToLogResourcePolicyPtrOutputWithContext(context.Background())
}

func (i *logResourcePolicyPtrType) ToLogResourcePolicyPtrOutputWithContext(ctx context.Context) LogResourcePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogResourcePolicyPtrOutput)
}

// LogResourcePolicyArrayInput is an input type that accepts LogResourcePolicyArray and LogResourcePolicyArrayOutput values.
// You can construct a concrete instance of `LogResourcePolicyArrayInput` via:
//
//          LogResourcePolicyArray{ LogResourcePolicyArgs{...} }
type LogResourcePolicyArrayInput interface {
	pulumi.Input

	ToLogResourcePolicyArrayOutput() LogResourcePolicyArrayOutput
	ToLogResourcePolicyArrayOutputWithContext(context.Context) LogResourcePolicyArrayOutput
}

type LogResourcePolicyArray []LogResourcePolicyInput

func (LogResourcePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogResourcePolicy)(nil)).Elem()
}

func (i LogResourcePolicyArray) ToLogResourcePolicyArrayOutput() LogResourcePolicyArrayOutput {
	return i.ToLogResourcePolicyArrayOutputWithContext(context.Background())
}

func (i LogResourcePolicyArray) ToLogResourcePolicyArrayOutputWithContext(ctx context.Context) LogResourcePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogResourcePolicyArrayOutput)
}

// LogResourcePolicyMapInput is an input type that accepts LogResourcePolicyMap and LogResourcePolicyMapOutput values.
// You can construct a concrete instance of `LogResourcePolicyMapInput` via:
//
//          LogResourcePolicyMap{ "key": LogResourcePolicyArgs{...} }
type LogResourcePolicyMapInput interface {
	pulumi.Input

	ToLogResourcePolicyMapOutput() LogResourcePolicyMapOutput
	ToLogResourcePolicyMapOutputWithContext(context.Context) LogResourcePolicyMapOutput
}

type LogResourcePolicyMap map[string]LogResourcePolicyInput

func (LogResourcePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogResourcePolicy)(nil)).Elem()
}

func (i LogResourcePolicyMap) ToLogResourcePolicyMapOutput() LogResourcePolicyMapOutput {
	return i.ToLogResourcePolicyMapOutputWithContext(context.Background())
}

func (i LogResourcePolicyMap) ToLogResourcePolicyMapOutputWithContext(ctx context.Context) LogResourcePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogResourcePolicyMapOutput)
}

type LogResourcePolicyOutput struct{ *pulumi.OutputState }

func (LogResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogResourcePolicy)(nil))
}

func (o LogResourcePolicyOutput) ToLogResourcePolicyOutput() LogResourcePolicyOutput {
	return o
}

func (o LogResourcePolicyOutput) ToLogResourcePolicyOutputWithContext(ctx context.Context) LogResourcePolicyOutput {
	return o
}

func (o LogResourcePolicyOutput) ToLogResourcePolicyPtrOutput() LogResourcePolicyPtrOutput {
	return o.ToLogResourcePolicyPtrOutputWithContext(context.Background())
}

func (o LogResourcePolicyOutput) ToLogResourcePolicyPtrOutputWithContext(ctx context.Context) LogResourcePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogResourcePolicy) *LogResourcePolicy {
		return &v
	}).(LogResourcePolicyPtrOutput)
}

type LogResourcePolicyPtrOutput struct{ *pulumi.OutputState }

func (LogResourcePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogResourcePolicy)(nil))
}

func (o LogResourcePolicyPtrOutput) ToLogResourcePolicyPtrOutput() LogResourcePolicyPtrOutput {
	return o
}

func (o LogResourcePolicyPtrOutput) ToLogResourcePolicyPtrOutputWithContext(ctx context.Context) LogResourcePolicyPtrOutput {
	return o
}

func (o LogResourcePolicyPtrOutput) Elem() LogResourcePolicyOutput {
	return o.ApplyT(func(v *LogResourcePolicy) LogResourcePolicy {
		if v != nil {
			return *v
		}
		var ret LogResourcePolicy
		return ret
	}).(LogResourcePolicyOutput)
}

type LogResourcePolicyArrayOutput struct{ *pulumi.OutputState }

func (LogResourcePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogResourcePolicy)(nil))
}

func (o LogResourcePolicyArrayOutput) ToLogResourcePolicyArrayOutput() LogResourcePolicyArrayOutput {
	return o
}

func (o LogResourcePolicyArrayOutput) ToLogResourcePolicyArrayOutputWithContext(ctx context.Context) LogResourcePolicyArrayOutput {
	return o
}

func (o LogResourcePolicyArrayOutput) Index(i pulumi.IntInput) LogResourcePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogResourcePolicy {
		return vs[0].([]LogResourcePolicy)[vs[1].(int)]
	}).(LogResourcePolicyOutput)
}

type LogResourcePolicyMapOutput struct{ *pulumi.OutputState }

func (LogResourcePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogResourcePolicy)(nil))
}

func (o LogResourcePolicyMapOutput) ToLogResourcePolicyMapOutput() LogResourcePolicyMapOutput {
	return o
}

func (o LogResourcePolicyMapOutput) ToLogResourcePolicyMapOutputWithContext(ctx context.Context) LogResourcePolicyMapOutput {
	return o
}

func (o LogResourcePolicyMapOutput) MapIndex(k pulumi.StringInput) LogResourcePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogResourcePolicy {
		return vs[0].(map[string]LogResourcePolicy)[vs[1].(string)]
	}).(LogResourcePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogResourcePolicyInput)(nil)).Elem(), &LogResourcePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogResourcePolicyPtrInput)(nil)).Elem(), &LogResourcePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogResourcePolicyArrayInput)(nil)).Elem(), LogResourcePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogResourcePolicyMapInput)(nil)).Elem(), LogResourcePolicyMap{})
	pulumi.RegisterOutputType(LogResourcePolicyOutput{})
	pulumi.RegisterOutputType(LogResourcePolicyPtrOutput{})
	pulumi.RegisterOutputType(LogResourcePolicyArrayOutput{})
	pulumi.RegisterOutputType(LogResourcePolicyMapOutput{})
}
