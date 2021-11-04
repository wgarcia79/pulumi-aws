// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudWatch Logs query definition resource.
//
// ## Example Usage
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
// 		_, err := cloudwatch.NewQueryDefinition(ctx, "example", &cloudwatch.QueryDefinitionArgs{
// 			LogGroupNames: pulumi.StringArray{
// 				pulumi.String("/aws/logGroup1"),
// 				pulumi.String("/aws/logGroup2"),
// 			},
// 			QueryString: pulumi.String(fmt.Sprintf("%v%v%v%v", "fields @timestamp, @message\n", "| sort @timestamp desc\n", "| limit 25\n", "\n")),
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
// CloudWatch query definitions can be imported using the query definition ARN. The ARN can be found on the "Edit Query" page for the query in the AWS Console.
//
// ```sh
//  $ pulumi import aws:cloudwatch/queryDefinition:QueryDefinition example arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd
// ```
type QueryDefinition struct {
	pulumi.CustomResourceState

	// Specific log groups to use with the query.
	LogGroupNames pulumi.StringArrayOutput `pulumi:"logGroupNames"`
	// The name of the query.
	Name pulumi.StringOutput `pulumi:"name"`
	// The query definition ID.
	QueryDefinitionId pulumi.StringOutput `pulumi:"queryDefinitionId"`
	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString pulumi.StringOutput `pulumi:"queryString"`
}

// NewQueryDefinition registers a new resource with the given unique name, arguments, and options.
func NewQueryDefinition(ctx *pulumi.Context,
	name string, args *QueryDefinitionArgs, opts ...pulumi.ResourceOption) (*QueryDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QueryString == nil {
		return nil, errors.New("invalid value for required argument 'QueryString'")
	}
	var resource QueryDefinition
	err := ctx.RegisterResource("aws:cloudwatch/queryDefinition:QueryDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueryDefinition gets an existing QueryDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryDefinitionState, opts ...pulumi.ResourceOption) (*QueryDefinition, error) {
	var resource QueryDefinition
	err := ctx.ReadResource("aws:cloudwatch/queryDefinition:QueryDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueryDefinition resources.
type queryDefinitionState struct {
	// Specific log groups to use with the query.
	LogGroupNames []string `pulumi:"logGroupNames"`
	// The name of the query.
	Name *string `pulumi:"name"`
	// The query definition ID.
	QueryDefinitionId *string `pulumi:"queryDefinitionId"`
	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString *string `pulumi:"queryString"`
}

type QueryDefinitionState struct {
	// Specific log groups to use with the query.
	LogGroupNames pulumi.StringArrayInput
	// The name of the query.
	Name pulumi.StringPtrInput
	// The query definition ID.
	QueryDefinitionId pulumi.StringPtrInput
	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString pulumi.StringPtrInput
}

func (QueryDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryDefinitionState)(nil)).Elem()
}

type queryDefinitionArgs struct {
	// Specific log groups to use with the query.
	LogGroupNames []string `pulumi:"logGroupNames"`
	// The name of the query.
	Name *string `pulumi:"name"`
	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString string `pulumi:"queryString"`
}

// The set of arguments for constructing a QueryDefinition resource.
type QueryDefinitionArgs struct {
	// Specific log groups to use with the query.
	LogGroupNames pulumi.StringArrayInput
	// The name of the query.
	Name pulumi.StringPtrInput
	// The query to save. You can read more about CloudWatch Logs Query Syntax in the [documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	QueryString pulumi.StringInput
}

func (QueryDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryDefinitionArgs)(nil)).Elem()
}

type QueryDefinitionInput interface {
	pulumi.Input

	ToQueryDefinitionOutput() QueryDefinitionOutput
	ToQueryDefinitionOutputWithContext(ctx context.Context) QueryDefinitionOutput
}

func (*QueryDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinition)(nil))
}

func (i *QueryDefinition) ToQueryDefinitionOutput() QueryDefinitionOutput {
	return i.ToQueryDefinitionOutputWithContext(context.Background())
}

func (i *QueryDefinition) ToQueryDefinitionOutputWithContext(ctx context.Context) QueryDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionOutput)
}

func (i *QueryDefinition) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return i.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (i *QueryDefinition) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionPtrOutput)
}

type QueryDefinitionPtrInput interface {
	pulumi.Input

	ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput
	ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput
}

type queryDefinitionPtrType QueryDefinitionArgs

func (*queryDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinition)(nil))
}

func (i *queryDefinitionPtrType) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return i.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (i *queryDefinitionPtrType) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionPtrOutput)
}

// QueryDefinitionArrayInput is an input type that accepts QueryDefinitionArray and QueryDefinitionArrayOutput values.
// You can construct a concrete instance of `QueryDefinitionArrayInput` via:
//
//          QueryDefinitionArray{ QueryDefinitionArgs{...} }
type QueryDefinitionArrayInput interface {
	pulumi.Input

	ToQueryDefinitionArrayOutput() QueryDefinitionArrayOutput
	ToQueryDefinitionArrayOutputWithContext(context.Context) QueryDefinitionArrayOutput
}

type QueryDefinitionArray []QueryDefinitionInput

func (QueryDefinitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueryDefinition)(nil)).Elem()
}

func (i QueryDefinitionArray) ToQueryDefinitionArrayOutput() QueryDefinitionArrayOutput {
	return i.ToQueryDefinitionArrayOutputWithContext(context.Background())
}

func (i QueryDefinitionArray) ToQueryDefinitionArrayOutputWithContext(ctx context.Context) QueryDefinitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionArrayOutput)
}

// QueryDefinitionMapInput is an input type that accepts QueryDefinitionMap and QueryDefinitionMapOutput values.
// You can construct a concrete instance of `QueryDefinitionMapInput` via:
//
//          QueryDefinitionMap{ "key": QueryDefinitionArgs{...} }
type QueryDefinitionMapInput interface {
	pulumi.Input

	ToQueryDefinitionMapOutput() QueryDefinitionMapOutput
	ToQueryDefinitionMapOutputWithContext(context.Context) QueryDefinitionMapOutput
}

type QueryDefinitionMap map[string]QueryDefinitionInput

func (QueryDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueryDefinition)(nil)).Elem()
}

func (i QueryDefinitionMap) ToQueryDefinitionMapOutput() QueryDefinitionMapOutput {
	return i.ToQueryDefinitionMapOutputWithContext(context.Background())
}

func (i QueryDefinitionMap) ToQueryDefinitionMapOutputWithContext(ctx context.Context) QueryDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDefinitionMapOutput)
}

type QueryDefinitionOutput struct{ *pulumi.OutputState }

func (QueryDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDefinition)(nil))
}

func (o QueryDefinitionOutput) ToQueryDefinitionOutput() QueryDefinitionOutput {
	return o
}

func (o QueryDefinitionOutput) ToQueryDefinitionOutputWithContext(ctx context.Context) QueryDefinitionOutput {
	return o
}

func (o QueryDefinitionOutput) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return o.ToQueryDefinitionPtrOutputWithContext(context.Background())
}

func (o QueryDefinitionOutput) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDefinition) *QueryDefinition {
		return &v
	}).(QueryDefinitionPtrOutput)
}

type QueryDefinitionPtrOutput struct{ *pulumi.OutputState }

func (QueryDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDefinition)(nil))
}

func (o QueryDefinitionPtrOutput) ToQueryDefinitionPtrOutput() QueryDefinitionPtrOutput {
	return o
}

func (o QueryDefinitionPtrOutput) ToQueryDefinitionPtrOutputWithContext(ctx context.Context) QueryDefinitionPtrOutput {
	return o
}

func (o QueryDefinitionPtrOutput) Elem() QueryDefinitionOutput {
	return o.ApplyT(func(v *QueryDefinition) QueryDefinition {
		if v != nil {
			return *v
		}
		var ret QueryDefinition
		return ret
	}).(QueryDefinitionOutput)
}

type QueryDefinitionArrayOutput struct{ *pulumi.OutputState }

func (QueryDefinitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryDefinition)(nil))
}

func (o QueryDefinitionArrayOutput) ToQueryDefinitionArrayOutput() QueryDefinitionArrayOutput {
	return o
}

func (o QueryDefinitionArrayOutput) ToQueryDefinitionArrayOutputWithContext(ctx context.Context) QueryDefinitionArrayOutput {
	return o
}

func (o QueryDefinitionArrayOutput) Index(i pulumi.IntInput) QueryDefinitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryDefinition {
		return vs[0].([]QueryDefinition)[vs[1].(int)]
	}).(QueryDefinitionOutput)
}

type QueryDefinitionMapOutput struct{ *pulumi.OutputState }

func (QueryDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryDefinition)(nil))
}

func (o QueryDefinitionMapOutput) ToQueryDefinitionMapOutput() QueryDefinitionMapOutput {
	return o
}

func (o QueryDefinitionMapOutput) ToQueryDefinitionMapOutputWithContext(ctx context.Context) QueryDefinitionMapOutput {
	return o
}

func (o QueryDefinitionMapOutput) MapIndex(k pulumi.StringInput) QueryDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueryDefinition {
		return vs[0].(map[string]QueryDefinition)[vs[1].(string)]
	}).(QueryDefinitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueryDefinitionInput)(nil)).Elem(), &QueryDefinition{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueryDefinitionPtrInput)(nil)).Elem(), &QueryDefinition{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueryDefinitionArrayInput)(nil)).Elem(), QueryDefinitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueryDefinitionMapInput)(nil)).Elem(), QueryDefinitionMap{})
	pulumi.RegisterOutputType(QueryDefinitionOutput{})
	pulumi.RegisterOutputType(QueryDefinitionPtrOutput{})
	pulumi.RegisterOutputType(QueryDefinitionArrayOutput{})
	pulumi.RegisterOutputType(QueryDefinitionMapOutput{})
}
