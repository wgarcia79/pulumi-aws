// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about a Lambda Function.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		functionName := cfg.Require("functionName")
// 		_, err := lambda.LookupFunction(ctx, &lambda.LookupFunctionArgs{
// 			FunctionName: functionName,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("aws:lambda/getFunction:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunction.
type LookupFunctionArgs struct {
	// Name of the lambda function.
	FunctionName string `pulumi:"functionName"`
	// Alias name or version number of the lambda function. e.g. `$LATEST`, `my-alias`, or `1`
	Qualifier *string           `pulumi:"qualifier"`
	Tags      map[string]string `pulumi:"tags"`
}

// A collection of values returned by getFunction.
type LookupFunctionResult struct {
	// The instruction set architecture for the Lambda function.
	Architectures []string `pulumi:"architectures"`
	// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `qualifiedArn`.
	Arn string `pulumi:"arn"`
	// Amazon Resource Name (ARN) for a Code Signing Configuration.
	CodeSigningConfigArn string `pulumi:"codeSigningConfigArn"`
	// Configure the function's *dead letter queue*.
	DeadLetterConfig GetFunctionDeadLetterConfig `pulumi:"deadLetterConfig"`
	// Description of what your Lambda Function does.
	Description string `pulumi:"description"`
	// The Lambda environment's configuration settings.
	Environment GetFunctionEnvironment `pulumi:"environment"`
	// The connection settings for an Amazon EFS file system.
	FileSystemConfigs []GetFunctionFileSystemConfig `pulumi:"fileSystemConfigs"`
	FunctionName      string                        `pulumi:"functionName"`
	// The function entrypoint in your code.
	Handler string `pulumi:"handler"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ARN to be used for invoking Lambda Function from API Gateway.
	InvokeArn string `pulumi:"invokeArn"`
	// The ARN for the KMS encryption key.
	KmsKeyArn string `pulumi:"kmsKeyArn"`
	// The date this resource was last modified.
	LastModified string `pulumi:"lastModified"`
	// A list of Lambda Layer ARNs attached to your Lambda Function.
	Layers []string `pulumi:"layers"`
	// Amount of memory in MB your Lambda Function can use at runtime.
	MemorySize int `pulumi:"memorySize"`
	// Qualified (`:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `arn`.
	QualifiedArn string  `pulumi:"qualifiedArn"`
	Qualifier    *string `pulumi:"qualifier"`
	// The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
	ReservedConcurrentExecutions int `pulumi:"reservedConcurrentExecutions"`
	// IAM role attached to the Lambda Function.
	Role string `pulumi:"role"`
	// The runtime environment for the Lambda function.
	Runtime string `pulumi:"runtime"`
	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn string `pulumi:"signingJobArn"`
	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn string `pulumi:"signingProfileVersionArn"`
	// Base64-encoded representation of raw SHA-256 sum of the zip file.
	SourceCodeHash string `pulumi:"sourceCodeHash"`
	// The size in bytes of the function .zip file.
	SourceCodeSize int               `pulumi:"sourceCodeSize"`
	Tags           map[string]string `pulumi:"tags"`
	// The function execution time at which Lambda should terminate the function.
	Timeout int `pulumi:"timeout"`
	// Tracing settings of the function.
	TracingConfig GetFunctionTracingConfig `pulumi:"tracingConfig"`
	// The version of the Lambda function.
	Version string `pulumi:"version"`
	// VPC configuration associated with your Lambda function.
	VpcConfig GetFunctionVpcConfig `pulumi:"vpcConfig"`
}

func LookupFunctionOutput(ctx *pulumi.Context, args LookupFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionResult, error) {
			args := v.(LookupFunctionArgs)
			r, err := LookupFunction(ctx, &args, opts...)
			return *r, err
		}).(LookupFunctionResultOutput)
}

// A collection of arguments for invoking getFunction.
type LookupFunctionOutputArgs struct {
	// Name of the lambda function.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Alias name or version number of the lambda function. e.g. `$LATEST`, `my-alias`, or `1`
	Qualifier pulumi.StringPtrInput `pulumi:"qualifier"`
	Tags      pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionArgs)(nil)).Elem()
}

// A collection of values returned by getFunction.
type LookupFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionResult)(nil)).Elem()
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutput() LookupFunctionResultOutput {
	return o
}

func (o LookupFunctionResultOutput) ToLookupFunctionResultOutputWithContext(ctx context.Context) LookupFunctionResultOutput {
	return o
}

// The instruction set architecture for the Lambda function.
func (o LookupFunctionResultOutput) Architectures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []string { return v.Architectures }).(pulumi.StringArrayOutput)
}

// Unqualified (no `:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `qualifiedArn`.
func (o LookupFunctionResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) for a Code Signing Configuration.
func (o LookupFunctionResultOutput) CodeSigningConfigArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.CodeSigningConfigArn }).(pulumi.StringOutput)
}

// Configure the function's *dead letter queue*.
func (o LookupFunctionResultOutput) DeadLetterConfig() GetFunctionDeadLetterConfigOutput {
	return o.ApplyT(func(v LookupFunctionResult) GetFunctionDeadLetterConfig { return v.DeadLetterConfig }).(GetFunctionDeadLetterConfigOutput)
}

// Description of what your Lambda Function does.
func (o LookupFunctionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Lambda environment's configuration settings.
func (o LookupFunctionResultOutput) Environment() GetFunctionEnvironmentOutput {
	return o.ApplyT(func(v LookupFunctionResult) GetFunctionEnvironment { return v.Environment }).(GetFunctionEnvironmentOutput)
}

// The connection settings for an Amazon EFS file system.
func (o LookupFunctionResultOutput) FileSystemConfigs() GetFunctionFileSystemConfigArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []GetFunctionFileSystemConfig { return v.FileSystemConfigs }).(GetFunctionFileSystemConfigArrayOutput)
}

func (o LookupFunctionResultOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.FunctionName }).(pulumi.StringOutput)
}

// The function entrypoint in your code.
func (o LookupFunctionResultOutput) Handler() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Handler }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ARN to be used for invoking Lambda Function from API Gateway.
func (o LookupFunctionResultOutput) InvokeArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.InvokeArn }).(pulumi.StringOutput)
}

// The ARN for the KMS encryption key.
func (o LookupFunctionResultOutput) KmsKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.KmsKeyArn }).(pulumi.StringOutput)
}

// The date this resource was last modified.
func (o LookupFunctionResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// A list of Lambda Layer ARNs attached to your Lambda Function.
func (o LookupFunctionResultOutput) Layers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFunctionResult) []string { return v.Layers }).(pulumi.StringArrayOutput)
}

// Amount of memory in MB your Lambda Function can use at runtime.
func (o LookupFunctionResultOutput) MemorySize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.MemorySize }).(pulumi.IntOutput)
}

// Qualified (`:QUALIFIER` or `:VERSION` suffix) Amazon Resource Name (ARN) identifying your Lambda Function. See also `arn`.
func (o LookupFunctionResultOutput) QualifiedArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.QualifiedArn }).(pulumi.StringOutput)
}

func (o LookupFunctionResultOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionResult) *string { return v.Qualifier }).(pulumi.StringPtrOutput)
}

// The amount of reserved concurrent executions for this lambda function or `-1` if unreserved.
func (o LookupFunctionResultOutput) ReservedConcurrentExecutions() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.ReservedConcurrentExecutions }).(pulumi.IntOutput)
}

// IAM role attached to the Lambda Function.
func (o LookupFunctionResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Role }).(pulumi.StringOutput)
}

// The runtime environment for the Lambda function.
func (o LookupFunctionResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Runtime }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of a signing job.
func (o LookupFunctionResultOutput) SigningJobArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.SigningJobArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for a signing profile version.
func (o LookupFunctionResultOutput) SigningProfileVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.SigningProfileVersionArn }).(pulumi.StringOutput)
}

// Base64-encoded representation of raw SHA-256 sum of the zip file.
func (o LookupFunctionResultOutput) SourceCodeHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.SourceCodeHash }).(pulumi.StringOutput)
}

// The size in bytes of the function .zip file.
func (o LookupFunctionResultOutput) SourceCodeSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.SourceCodeSize }).(pulumi.IntOutput)
}

func (o LookupFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The function execution time at which Lambda should terminate the function.
func (o LookupFunctionResultOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFunctionResult) int { return v.Timeout }).(pulumi.IntOutput)
}

// Tracing settings of the function.
func (o LookupFunctionResultOutput) TracingConfig() GetFunctionTracingConfigOutput {
	return o.ApplyT(func(v LookupFunctionResult) GetFunctionTracingConfig { return v.TracingConfig }).(GetFunctionTracingConfigOutput)
}

// The version of the Lambda function.
func (o LookupFunctionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFunctionResult) string { return v.Version }).(pulumi.StringOutput)
}

// VPC configuration associated with your Lambda function.
func (o LookupFunctionResultOutput) VpcConfig() GetFunctionVpcConfigOutput {
	return o.ApplyT(func(v LookupFunctionResult) GetFunctionVpcConfig { return v.VpcConfig }).(GetFunctionVpcConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionResultOutput{})
}
