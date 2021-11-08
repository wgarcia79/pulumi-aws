// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gives an external source (like a CloudWatch Event Rule, SNS, or S3) permission to access the Lambda function.
//
// ## Example Usage
// ### Specify Lambda permissions for API Gateway REST API
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lambda.NewPermission(ctx, "lambdaPermission", &lambda.PermissionArgs{
// 			Action:    pulumi.String("lambda:InvokeFunction"),
// 			Function:  pulumi.Any("MyDemoFunction"),
// 			Principal: pulumi.String("apigateway.amazonaws.com"),
// 			SourceArn: myDemoAPI.ExecutionArn.ApplyT(func(executionArn string) (string, error) {
// 				return fmt.Sprintf("%v%v", executionArn, "/*/*/*"), nil
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
// Lambda permission statements can be imported using function_name/statement_id, with an optional qualifier, e.g.,
//
// ```sh
//  $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function/AllowExecutionFromCloudWatch
// ```
//
// ```sh
//  $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function:qualifier_name/AllowExecutionFromCloudWatch
// ```
type Permission struct {
	pulumi.CustomResourceState

	// The AWS Lambda action you want to allow in this statement. (e.g., `lambda:InvokeFunction`)
	Action pulumi.StringOutput `pulumi:"action"`
	// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
	EventSourceToken pulumi.StringPtrOutput `pulumi:"eventSourceToken"`
	// Name of the Lambda function whose resource policy you are updating
	Function pulumi.StringOutput `pulumi:"function"`
	// The principal who is getting this permissionE.g., `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARNE.g., `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
	Qualifier pulumi.StringPtrOutput `pulumi:"qualifier"`
	// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
	SourceAccount pulumi.StringPtrOutput `pulumi:"sourceAccount"`
	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
	// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
	SourceArn pulumi.StringPtrOutput `pulumi:"sourceArn"`
	// A unique statement identifier. By default generated by this provider.
	StatementId pulumi.StringOutput `pulumi:"statementId"`
	// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrOutput `pulumi:"statementIdPrefix"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Function == nil {
		return nil, errors.New("invalid value for required argument 'Function'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	var resource Permission
	err := ctx.RegisterResource("aws:lambda/permission:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws:lambda/permission:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
	// The AWS Lambda action you want to allow in this statement. (e.g., `lambda:InvokeFunction`)
	Action *string `pulumi:"action"`
	// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
	EventSourceToken *string `pulumi:"eventSourceToken"`
	// Name of the Lambda function whose resource policy you are updating
	Function interface{} `pulumi:"function"`
	// The principal who is getting this permissionE.g., `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
	Principal *string `pulumi:"principal"`
	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARNE.g., `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
	Qualifier *string `pulumi:"qualifier"`
	// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
	SourceAccount *string `pulumi:"sourceAccount"`
	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
	// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
	SourceArn *string `pulumi:"sourceArn"`
	// A unique statement identifier. By default generated by this provider.
	StatementId *string `pulumi:"statementId"`
	// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix *string `pulumi:"statementIdPrefix"`
}

type PermissionState struct {
	// The AWS Lambda action you want to allow in this statement. (e.g., `lambda:InvokeFunction`)
	Action pulumi.StringPtrInput
	// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
	EventSourceToken pulumi.StringPtrInput
	// Name of the Lambda function whose resource policy you are updating
	Function pulumi.Input
	// The principal who is getting this permissionE.g., `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
	Principal pulumi.StringPtrInput
	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARNE.g., `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
	Qualifier pulumi.StringPtrInput
	// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
	SourceAccount pulumi.StringPtrInput
	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
	// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
	SourceArn pulumi.StringPtrInput
	// A unique statement identifier. By default generated by this provider.
	StatementId pulumi.StringPtrInput
	// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrInput
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	// The AWS Lambda action you want to allow in this statement. (e.g., `lambda:InvokeFunction`)
	Action string `pulumi:"action"`
	// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
	EventSourceToken *string `pulumi:"eventSourceToken"`
	// Name of the Lambda function whose resource policy you are updating
	Function interface{} `pulumi:"function"`
	// The principal who is getting this permissionE.g., `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
	Principal string `pulumi:"principal"`
	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARNE.g., `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
	Qualifier *string `pulumi:"qualifier"`
	// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
	SourceAccount *string `pulumi:"sourceAccount"`
	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
	// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
	SourceArn *string `pulumi:"sourceArn"`
	// A unique statement identifier. By default generated by this provider.
	StatementId *string `pulumi:"statementId"`
	// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix *string `pulumi:"statementIdPrefix"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	// The AWS Lambda action you want to allow in this statement. (e.g., `lambda:InvokeFunction`)
	Action pulumi.StringInput
	// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
	EventSourceToken pulumi.StringPtrInput
	// Name of the Lambda function whose resource policy you are updating
	Function pulumi.Input
	// The principal who is getting this permissionE.g., `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
	Principal pulumi.StringInput
	// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARNE.g., `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
	Qualifier pulumi.StringPtrInput
	// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
	SourceAccount pulumi.StringPtrInput
	// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
	// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
	// For S3, this should be the ARN of the S3 Bucket.
	// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
	// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
	SourceArn pulumi.StringPtrInput
	// A unique statement identifier. By default generated by this provider.
	StatementId pulumi.StringPtrInput
	// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statementId`.
	StatementIdPrefix pulumi.StringPtrInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}

type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(ctx context.Context) PermissionOutput
}

func (*Permission) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil))
}

func (i *Permission) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}

func (i *Permission) ToPermissionPtrOutput() PermissionPtrOutput {
	return i.ToPermissionPtrOutputWithContext(context.Background())
}

func (i *Permission) ToPermissionPtrOutputWithContext(ctx context.Context) PermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionPtrOutput)
}

type PermissionPtrInput interface {
	pulumi.Input

	ToPermissionPtrOutput() PermissionPtrOutput
	ToPermissionPtrOutputWithContext(ctx context.Context) PermissionPtrOutput
}

type permissionPtrType PermissionArgs

func (*permissionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil))
}

func (i *permissionPtrType) ToPermissionPtrOutput() PermissionPtrOutput {
	return i.ToPermissionPtrOutputWithContext(context.Background())
}

func (i *permissionPtrType) ToPermissionPtrOutputWithContext(ctx context.Context) PermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionPtrOutput)
}

// PermissionArrayInput is an input type that accepts PermissionArray and PermissionArrayOutput values.
// You can construct a concrete instance of `PermissionArrayInput` via:
//
//          PermissionArray{ PermissionArgs{...} }
type PermissionArrayInput interface {
	pulumi.Input

	ToPermissionArrayOutput() PermissionArrayOutput
	ToPermissionArrayOutputWithContext(context.Context) PermissionArrayOutput
}

type PermissionArray []PermissionInput

func (PermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Permission)(nil)).Elem()
}

func (i PermissionArray) ToPermissionArrayOutput() PermissionArrayOutput {
	return i.ToPermissionArrayOutputWithContext(context.Background())
}

func (i PermissionArray) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionArrayOutput)
}

// PermissionMapInput is an input type that accepts PermissionMap and PermissionMapOutput values.
// You can construct a concrete instance of `PermissionMapInput` via:
//
//          PermissionMap{ "key": PermissionArgs{...} }
type PermissionMapInput interface {
	pulumi.Input

	ToPermissionMapOutput() PermissionMapOutput
	ToPermissionMapOutputWithContext(context.Context) PermissionMapOutput
}

type PermissionMap map[string]PermissionInput

func (PermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Permission)(nil)).Elem()
}

func (i PermissionMap) ToPermissionMapOutput() PermissionMapOutput {
	return i.ToPermissionMapOutputWithContext(context.Background())
}

func (i PermissionMap) ToPermissionMapOutputWithContext(ctx context.Context) PermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionMapOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil))
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionPtrOutput() PermissionPtrOutput {
	return o.ToPermissionPtrOutputWithContext(context.Background())
}

func (o PermissionOutput) ToPermissionPtrOutputWithContext(ctx context.Context) PermissionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Permission) *Permission {
		return &v
	}).(PermissionPtrOutput)
}

type PermissionPtrOutput struct{ *pulumi.OutputState }

func (PermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permission)(nil))
}

func (o PermissionPtrOutput) ToPermissionPtrOutput() PermissionPtrOutput {
	return o
}

func (o PermissionPtrOutput) ToPermissionPtrOutputWithContext(ctx context.Context) PermissionPtrOutput {
	return o
}

func (o PermissionPtrOutput) Elem() PermissionOutput {
	return o.ApplyT(func(v *Permission) Permission {
		if v != nil {
			return *v
		}
		var ret Permission
		return ret
	}).(PermissionOutput)
}

type PermissionArrayOutput struct{ *pulumi.OutputState }

func (PermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil))
}

func (o PermissionArrayOutput) ToPermissionArrayOutput() PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) Index(i pulumi.IntInput) PermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Permission {
		return vs[0].([]Permission)[vs[1].(int)]
	}).(PermissionOutput)
}

type PermissionMapOutput struct{ *pulumi.OutputState }

func (PermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Permission)(nil))
}

func (o PermissionMapOutput) ToPermissionMapOutput() PermissionMapOutput {
	return o
}

func (o PermissionMapOutput) ToPermissionMapOutputWithContext(ctx context.Context) PermissionMapOutput {
	return o
}

func (o PermissionMapOutput) MapIndex(k pulumi.StringInput) PermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Permission {
		return vs[0].(map[string]Permission)[vs[1].(string)]
	}).(PermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionPtrInput)(nil)).Elem(), &Permission{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionArrayInput)(nil)).Elem(), PermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionMapInput)(nil)).Elem(), PermissionMap{})
	pulumi.RegisterOutputType(PermissionOutput{})
	pulumi.RegisterOutputType(PermissionPtrOutput{})
	pulumi.RegisterOutputType(PermissionArrayOutput{})
	pulumi.RegisterOutputType(PermissionMapOutput{})
}
