// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFormation Stack resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudformation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudformation.NewStack(ctx, "network", &cloudformation.StackArgs{
// 			Parameters: pulumi.StringMap{
// 				"VPCCidr": pulumi.String("10.0.0.0/16"),
// 			},
// 			TemplateBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Parameters\" : {\n", "    \"VPCCidr\" : {\n", "      \"Type\" : \"String\",\n", "      \"Default\" : \"10.0.0.0/16\",\n", "      \"Description\" : \"Enter the CIDR block for the VPC. Default is 10.0.0.0/16.\"\n", "    }\n", "  },\n", "  \"Resources\" : {\n", "    \"myVpc\": {\n", "      \"Type\" : \"AWS::EC2::VPC\",\n", "      \"Properties\" : {\n", "        \"CidrBlock\" : { \"Ref\" : \"VPCCidr\" },\n", "        \"Tags\" : [\n", "          {\"Key\": \"Name\", \"Value\": \"Primary_CF_VPC\"}\n", "        ]\n", "      }\n", "    }\n", "  }\n", "}\n", "\n")),
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
// Cloudformation Stacks can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:cloudformation/stack:Stack stack networking-stack
// ```
type Stack struct {
	pulumi.CustomResourceState

	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrOutput `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrOutput `pulumi:"iamRoleArn"`
	// Stack name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayOutput `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrOutput `pulumi:"onFailure"`
	// A map of outputs from the stack.
	Outputs pulumi.StringMapOutput `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringOutput `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrOutput `pulumi:"policyUrl"`
	// Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringOutput `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrOutput `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrOutput `pulumi:"timeoutInMinutes"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	if args == nil {
		args = &StackArgs{}
	}

	var resource Stack
	err := ctx.RegisterResource("aws:cloudformation/stack:Stack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	var resource Stack
	err := ctx.ReadResource("aws:cloudformation/stack:Stack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type stackState struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Stack name.
	Name *string `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure *string `pulumi:"onFailure"`
	// A map of outputs from the stack.
	Outputs map[string]string `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody *string `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl *string `pulumi:"policyUrl"`
	// Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl *string `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
}

type StackState struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrInput
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrInput
	// Stack name.
	Name pulumi.StringPtrInput
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrInput
	// A map of outputs from the stack.
	Outputs pulumi.StringMapInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringPtrInput
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrInput
	// Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringPtrInput
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrInput
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrInput
}

func (StackState) ElementType() reflect.Type {
	return reflect.TypeOf((*stackState)(nil)).Elem()
}

type stackArgs struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback *bool `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Stack name.
	Name *string `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []string `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure *string `pulumi:"onFailure"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody *string `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl *string `pulumi:"policyUrl"`
	// Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody *string `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl *string `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes *int `pulumi:"timeoutInMinutes"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolPtrInput
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringPtrInput
	// Stack name.
	Name pulumi.StringPtrInput
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringPtrInput
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.StringMapInput
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringPtrInput
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringPtrInput
	// Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringPtrInput
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringPtrInput
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntPtrInput
}

func (StackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stackArgs)(nil)).Elem()
}

type StackInput interface {
	pulumi.Input

	ToStackOutput() StackOutput
	ToStackOutputWithContext(ctx context.Context) StackOutput
}

func (*Stack) ElementType() reflect.Type {
	return reflect.TypeOf((*Stack)(nil))
}

func (i *Stack) ToStackOutput() StackOutput {
	return i.ToStackOutputWithContext(context.Background())
}

func (i *Stack) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackOutput)
}

func (i *Stack) ToStackPtrOutput() StackPtrOutput {
	return i.ToStackPtrOutputWithContext(context.Background())
}

func (i *Stack) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPtrOutput)
}

type StackPtrInput interface {
	pulumi.Input

	ToStackPtrOutput() StackPtrOutput
	ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput
}

type stackPtrType StackArgs

func (*stackPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil))
}

func (i *stackPtrType) ToStackPtrOutput() StackPtrOutput {
	return i.ToStackPtrOutputWithContext(context.Background())
}

func (i *stackPtrType) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackPtrOutput)
}

// StackArrayInput is an input type that accepts StackArray and StackArrayOutput values.
// You can construct a concrete instance of `StackArrayInput` via:
//
//          StackArray{ StackArgs{...} }
type StackArrayInput interface {
	pulumi.Input

	ToStackArrayOutput() StackArrayOutput
	ToStackArrayOutputWithContext(context.Context) StackArrayOutput
}

type StackArray []StackInput

func (StackArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stack)(nil)).Elem()
}

func (i StackArray) ToStackArrayOutput() StackArrayOutput {
	return i.ToStackArrayOutputWithContext(context.Background())
}

func (i StackArray) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackArrayOutput)
}

// StackMapInput is an input type that accepts StackMap and StackMapOutput values.
// You can construct a concrete instance of `StackMapInput` via:
//
//          StackMap{ "key": StackArgs{...} }
type StackMapInput interface {
	pulumi.Input

	ToStackMapOutput() StackMapOutput
	ToStackMapOutputWithContext(context.Context) StackMapOutput
}

type StackMap map[string]StackInput

func (StackMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stack)(nil)).Elem()
}

func (i StackMap) ToStackMapOutput() StackMapOutput {
	return i.ToStackMapOutputWithContext(context.Background())
}

func (i StackMap) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackMapOutput)
}

type StackOutput struct{ *pulumi.OutputState }

func (StackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Stack)(nil))
}

func (o StackOutput) ToStackOutput() StackOutput {
	return o
}

func (o StackOutput) ToStackOutputWithContext(ctx context.Context) StackOutput {
	return o
}

func (o StackOutput) ToStackPtrOutput() StackPtrOutput {
	return o.ToStackPtrOutputWithContext(context.Background())
}

func (o StackOutput) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Stack) *Stack {
		return &v
	}).(StackPtrOutput)
}

type StackPtrOutput struct{ *pulumi.OutputState }

func (StackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stack)(nil))
}

func (o StackPtrOutput) ToStackPtrOutput() StackPtrOutput {
	return o
}

func (o StackPtrOutput) ToStackPtrOutputWithContext(ctx context.Context) StackPtrOutput {
	return o
}

func (o StackPtrOutput) Elem() StackOutput {
	return o.ApplyT(func(v *Stack) Stack {
		if v != nil {
			return *v
		}
		var ret Stack
		return ret
	}).(StackOutput)
}

type StackArrayOutput struct{ *pulumi.OutputState }

func (StackArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Stack)(nil))
}

func (o StackArrayOutput) ToStackArrayOutput() StackArrayOutput {
	return o
}

func (o StackArrayOutput) ToStackArrayOutputWithContext(ctx context.Context) StackArrayOutput {
	return o
}

func (o StackArrayOutput) Index(i pulumi.IntInput) StackOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Stack {
		return vs[0].([]Stack)[vs[1].(int)]
	}).(StackOutput)
}

type StackMapOutput struct{ *pulumi.OutputState }

func (StackMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Stack)(nil))
}

func (o StackMapOutput) ToStackMapOutput() StackMapOutput {
	return o
}

func (o StackMapOutput) ToStackMapOutputWithContext(ctx context.Context) StackMapOutput {
	return o
}

func (o StackMapOutput) MapIndex(k pulumi.StringInput) StackOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Stack {
		return vs[0].(map[string]Stack)[vs[1].(string)]
	}).(StackOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackInput)(nil)).Elem(), &Stack{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackPtrInput)(nil)).Elem(), &Stack{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackArrayInput)(nil)).Elem(), StackArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackMapInput)(nil)).Elem(), StackMap{})
	pulumi.RegisterOutputType(StackOutput{})
	pulumi.RegisterOutputType(StackPtrOutput{})
	pulumi.RegisterOutputType(StackArrayOutput{})
	pulumi.RegisterOutputType(StackMapOutput{})
}
