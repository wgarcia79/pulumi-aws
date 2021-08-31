// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoveryreadiness

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Route 53 Recovery Readiness Readiness Check.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53recoveryreadiness"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53recoveryreadiness.NewReadinessCheck(ctx, "example", &route53recoveryreadiness.ReadinessCheckArgs{
// 			ReadinessCheckName: pulumi.Any(my_cw_alarm_check),
// 			ResourceSetName:    pulumi.Any(my_cw_alarm_set),
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
// Route53 Recovery Readiness readiness checks can be imported via the readiness check name, e.g.
//
// ```sh
//  $ pulumi import aws:route53recoveryreadiness/readinessCheck:ReadinessCheck my-cw-alarm-check
// ```
type ReadinessCheck struct {
	pulumi.CustomResourceState

	// ARN of the readiness_check
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Unique name describing the readiness check.
	ReadinessCheckName pulumi.StringOutput `pulumi:"readinessCheckName"`
	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName pulumi.StringOutput `pulumi:"resourceSetName"`
	// Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewReadinessCheck registers a new resource with the given unique name, arguments, and options.
func NewReadinessCheck(ctx *pulumi.Context,
	name string, args *ReadinessCheckArgs, opts ...pulumi.ResourceOption) (*ReadinessCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReadinessCheckName == nil {
		return nil, errors.New("invalid value for required argument 'ReadinessCheckName'")
	}
	if args.ResourceSetName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSetName'")
	}
	var resource ReadinessCheck
	err := ctx.RegisterResource("aws:route53recoveryreadiness/readinessCheck:ReadinessCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadinessCheck gets an existing ReadinessCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadinessCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadinessCheckState, opts ...pulumi.ResourceOption) (*ReadinessCheck, error) {
	var resource ReadinessCheck
	err := ctx.ReadResource("aws:route53recoveryreadiness/readinessCheck:ReadinessCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadinessCheck resources.
type readinessCheckState struct {
	// ARN of the readiness_check
	Arn *string `pulumi:"arn"`
	// Unique name describing the readiness check.
	ReadinessCheckName *string `pulumi:"readinessCheckName"`
	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName *string `pulumi:"resourceSetName"`
	// Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ReadinessCheckState struct {
	// ARN of the readiness_check
	Arn pulumi.StringPtrInput
	// Unique name describing the readiness check.
	ReadinessCheckName pulumi.StringPtrInput
	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (ReadinessCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*readinessCheckState)(nil)).Elem()
}

type readinessCheckArgs struct {
	// Unique name describing the readiness check.
	ReadinessCheckName string `pulumi:"readinessCheckName"`
	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName string `pulumi:"resourceSetName"`
	// Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a ReadinessCheck resource.
type ReadinessCheckArgs struct {
	// Unique name describing the readiness check.
	ReadinessCheckName pulumi.StringInput
	// Name describing the resource set that will be monitored for readiness.
	ResourceSetName pulumi.StringInput
	// Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
	TagsAll pulumi.StringMapInput
}

func (ReadinessCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readinessCheckArgs)(nil)).Elem()
}

type ReadinessCheckInput interface {
	pulumi.Input

	ToReadinessCheckOutput() ReadinessCheckOutput
	ToReadinessCheckOutputWithContext(ctx context.Context) ReadinessCheckOutput
}

func (*ReadinessCheck) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadinessCheck)(nil))
}

func (i *ReadinessCheck) ToReadinessCheckOutput() ReadinessCheckOutput {
	return i.ToReadinessCheckOutputWithContext(context.Background())
}

func (i *ReadinessCheck) ToReadinessCheckOutputWithContext(ctx context.Context) ReadinessCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadinessCheckOutput)
}

func (i *ReadinessCheck) ToReadinessCheckPtrOutput() ReadinessCheckPtrOutput {
	return i.ToReadinessCheckPtrOutputWithContext(context.Background())
}

func (i *ReadinessCheck) ToReadinessCheckPtrOutputWithContext(ctx context.Context) ReadinessCheckPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadinessCheckPtrOutput)
}

type ReadinessCheckPtrInput interface {
	pulumi.Input

	ToReadinessCheckPtrOutput() ReadinessCheckPtrOutput
	ToReadinessCheckPtrOutputWithContext(ctx context.Context) ReadinessCheckPtrOutput
}

type readinessCheckPtrType ReadinessCheckArgs

func (*readinessCheckPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadinessCheck)(nil))
}

func (i *readinessCheckPtrType) ToReadinessCheckPtrOutput() ReadinessCheckPtrOutput {
	return i.ToReadinessCheckPtrOutputWithContext(context.Background())
}

func (i *readinessCheckPtrType) ToReadinessCheckPtrOutputWithContext(ctx context.Context) ReadinessCheckPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadinessCheckPtrOutput)
}

// ReadinessCheckArrayInput is an input type that accepts ReadinessCheckArray and ReadinessCheckArrayOutput values.
// You can construct a concrete instance of `ReadinessCheckArrayInput` via:
//
//          ReadinessCheckArray{ ReadinessCheckArgs{...} }
type ReadinessCheckArrayInput interface {
	pulumi.Input

	ToReadinessCheckArrayOutput() ReadinessCheckArrayOutput
	ToReadinessCheckArrayOutputWithContext(context.Context) ReadinessCheckArrayOutput
}

type ReadinessCheckArray []ReadinessCheckInput

func (ReadinessCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadinessCheck)(nil)).Elem()
}

func (i ReadinessCheckArray) ToReadinessCheckArrayOutput() ReadinessCheckArrayOutput {
	return i.ToReadinessCheckArrayOutputWithContext(context.Background())
}

func (i ReadinessCheckArray) ToReadinessCheckArrayOutputWithContext(ctx context.Context) ReadinessCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadinessCheckArrayOutput)
}

// ReadinessCheckMapInput is an input type that accepts ReadinessCheckMap and ReadinessCheckMapOutput values.
// You can construct a concrete instance of `ReadinessCheckMapInput` via:
//
//          ReadinessCheckMap{ "key": ReadinessCheckArgs{...} }
type ReadinessCheckMapInput interface {
	pulumi.Input

	ToReadinessCheckMapOutput() ReadinessCheckMapOutput
	ToReadinessCheckMapOutputWithContext(context.Context) ReadinessCheckMapOutput
}

type ReadinessCheckMap map[string]ReadinessCheckInput

func (ReadinessCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadinessCheck)(nil)).Elem()
}

func (i ReadinessCheckMap) ToReadinessCheckMapOutput() ReadinessCheckMapOutput {
	return i.ToReadinessCheckMapOutputWithContext(context.Background())
}

func (i ReadinessCheckMap) ToReadinessCheckMapOutputWithContext(ctx context.Context) ReadinessCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadinessCheckMapOutput)
}

type ReadinessCheckOutput struct{ *pulumi.OutputState }

func (ReadinessCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadinessCheck)(nil))
}

func (o ReadinessCheckOutput) ToReadinessCheckOutput() ReadinessCheckOutput {
	return o
}

func (o ReadinessCheckOutput) ToReadinessCheckOutputWithContext(ctx context.Context) ReadinessCheckOutput {
	return o
}

func (o ReadinessCheckOutput) ToReadinessCheckPtrOutput() ReadinessCheckPtrOutput {
	return o.ToReadinessCheckPtrOutputWithContext(context.Background())
}

func (o ReadinessCheckOutput) ToReadinessCheckPtrOutputWithContext(ctx context.Context) ReadinessCheckPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReadinessCheck) *ReadinessCheck {
		return &v
	}).(ReadinessCheckPtrOutput)
}

type ReadinessCheckPtrOutput struct{ *pulumi.OutputState }

func (ReadinessCheckPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadinessCheck)(nil))
}

func (o ReadinessCheckPtrOutput) ToReadinessCheckPtrOutput() ReadinessCheckPtrOutput {
	return o
}

func (o ReadinessCheckPtrOutput) ToReadinessCheckPtrOutputWithContext(ctx context.Context) ReadinessCheckPtrOutput {
	return o
}

func (o ReadinessCheckPtrOutput) Elem() ReadinessCheckOutput {
	return o.ApplyT(func(v *ReadinessCheck) ReadinessCheck {
		if v != nil {
			return *v
		}
		var ret ReadinessCheck
		return ret
	}).(ReadinessCheckOutput)
}

type ReadinessCheckArrayOutput struct{ *pulumi.OutputState }

func (ReadinessCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReadinessCheck)(nil))
}

func (o ReadinessCheckArrayOutput) ToReadinessCheckArrayOutput() ReadinessCheckArrayOutput {
	return o
}

func (o ReadinessCheckArrayOutput) ToReadinessCheckArrayOutputWithContext(ctx context.Context) ReadinessCheckArrayOutput {
	return o
}

func (o ReadinessCheckArrayOutput) Index(i pulumi.IntInput) ReadinessCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReadinessCheck {
		return vs[0].([]ReadinessCheck)[vs[1].(int)]
	}).(ReadinessCheckOutput)
}

type ReadinessCheckMapOutput struct{ *pulumi.OutputState }

func (ReadinessCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReadinessCheck)(nil))
}

func (o ReadinessCheckMapOutput) ToReadinessCheckMapOutput() ReadinessCheckMapOutput {
	return o
}

func (o ReadinessCheckMapOutput) ToReadinessCheckMapOutputWithContext(ctx context.Context) ReadinessCheckMapOutput {
	return o
}

func (o ReadinessCheckMapOutput) MapIndex(k pulumi.StringInput) ReadinessCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReadinessCheck {
		return vs[0].(map[string]ReadinessCheck)[vs[1].(string)]
	}).(ReadinessCheckOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadinessCheckOutput{})
	pulumi.RegisterOutputType(ReadinessCheckPtrOutput{})
	pulumi.RegisterOutputType(ReadinessCheckArrayOutput{})
	pulumi.RegisterOutputType(ReadinessCheckMapOutput{})
}
