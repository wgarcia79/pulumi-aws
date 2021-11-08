// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Step Function Activity resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.NewActivity(ctx, "sfnActivity", nil)
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
// Activities can be imported using the `arn`, e.g.,
//
// ```sh
//  $ pulumi import aws:sfn/activity:Activity foo arn:aws:states:eu-west-1:123456789098:activity:bar
// ```
type Activity struct {
	pulumi.CustomResourceState

	// The date the activity was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The name of the activity to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewActivity registers a new resource with the given unique name, arguments, and options.
func NewActivity(ctx *pulumi.Context,
	name string, args *ActivityArgs, opts ...pulumi.ResourceOption) (*Activity, error) {
	if args == nil {
		args = &ActivityArgs{}
	}

	var resource Activity
	err := ctx.RegisterResource("aws:sfn/activity:Activity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivity gets an existing Activity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityState, opts ...pulumi.ResourceOption) (*Activity, error) {
	var resource Activity
	err := ctx.ReadResource("aws:sfn/activity:Activity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Activity resources.
type activityState struct {
	// The date the activity was created.
	CreationDate *string `pulumi:"creationDate"`
	// The name of the activity to create.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ActivityState struct {
	// The date the activity was created.
	CreationDate pulumi.StringPtrInput
	// The name of the activity to create.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ActivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityState)(nil)).Elem()
}

type activityArgs struct {
	// The name of the activity to create.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Activity resource.
type ActivityArgs struct {
	// The name of the activity to create.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ActivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityArgs)(nil)).Elem()
}

type ActivityInput interface {
	pulumi.Input

	ToActivityOutput() ActivityOutput
	ToActivityOutputWithContext(ctx context.Context) ActivityOutput
}

func (*Activity) ElementType() reflect.Type {
	return reflect.TypeOf((*Activity)(nil))
}

func (i *Activity) ToActivityOutput() ActivityOutput {
	return i.ToActivityOutputWithContext(context.Background())
}

func (i *Activity) ToActivityOutputWithContext(ctx context.Context) ActivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityOutput)
}

func (i *Activity) ToActivityPtrOutput() ActivityPtrOutput {
	return i.ToActivityPtrOutputWithContext(context.Background())
}

func (i *Activity) ToActivityPtrOutputWithContext(ctx context.Context) ActivityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityPtrOutput)
}

type ActivityPtrInput interface {
	pulumi.Input

	ToActivityPtrOutput() ActivityPtrOutput
	ToActivityPtrOutputWithContext(ctx context.Context) ActivityPtrOutput
}

type activityPtrType ActivityArgs

func (*activityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Activity)(nil))
}

func (i *activityPtrType) ToActivityPtrOutput() ActivityPtrOutput {
	return i.ToActivityPtrOutputWithContext(context.Background())
}

func (i *activityPtrType) ToActivityPtrOutputWithContext(ctx context.Context) ActivityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityPtrOutput)
}

// ActivityArrayInput is an input type that accepts ActivityArray and ActivityArrayOutput values.
// You can construct a concrete instance of `ActivityArrayInput` via:
//
//          ActivityArray{ ActivityArgs{...} }
type ActivityArrayInput interface {
	pulumi.Input

	ToActivityArrayOutput() ActivityArrayOutput
	ToActivityArrayOutputWithContext(context.Context) ActivityArrayOutput
}

type ActivityArray []ActivityInput

func (ActivityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Activity)(nil)).Elem()
}

func (i ActivityArray) ToActivityArrayOutput() ActivityArrayOutput {
	return i.ToActivityArrayOutputWithContext(context.Background())
}

func (i ActivityArray) ToActivityArrayOutputWithContext(ctx context.Context) ActivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityArrayOutput)
}

// ActivityMapInput is an input type that accepts ActivityMap and ActivityMapOutput values.
// You can construct a concrete instance of `ActivityMapInput` via:
//
//          ActivityMap{ "key": ActivityArgs{...} }
type ActivityMapInput interface {
	pulumi.Input

	ToActivityMapOutput() ActivityMapOutput
	ToActivityMapOutputWithContext(context.Context) ActivityMapOutput
}

type ActivityMap map[string]ActivityInput

func (ActivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Activity)(nil)).Elem()
}

func (i ActivityMap) ToActivityMapOutput() ActivityMapOutput {
	return i.ToActivityMapOutputWithContext(context.Background())
}

func (i ActivityMap) ToActivityMapOutputWithContext(ctx context.Context) ActivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityMapOutput)
}

type ActivityOutput struct{ *pulumi.OutputState }

func (ActivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Activity)(nil))
}

func (o ActivityOutput) ToActivityOutput() ActivityOutput {
	return o
}

func (o ActivityOutput) ToActivityOutputWithContext(ctx context.Context) ActivityOutput {
	return o
}

func (o ActivityOutput) ToActivityPtrOutput() ActivityPtrOutput {
	return o.ToActivityPtrOutputWithContext(context.Background())
}

func (o ActivityOutput) ToActivityPtrOutputWithContext(ctx context.Context) ActivityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Activity) *Activity {
		return &v
	}).(ActivityPtrOutput)
}

type ActivityPtrOutput struct{ *pulumi.OutputState }

func (ActivityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Activity)(nil))
}

func (o ActivityPtrOutput) ToActivityPtrOutput() ActivityPtrOutput {
	return o
}

func (o ActivityPtrOutput) ToActivityPtrOutputWithContext(ctx context.Context) ActivityPtrOutput {
	return o
}

func (o ActivityPtrOutput) Elem() ActivityOutput {
	return o.ApplyT(func(v *Activity) Activity {
		if v != nil {
			return *v
		}
		var ret Activity
		return ret
	}).(ActivityOutput)
}

type ActivityArrayOutput struct{ *pulumi.OutputState }

func (ActivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Activity)(nil))
}

func (o ActivityArrayOutput) ToActivityArrayOutput() ActivityArrayOutput {
	return o
}

func (o ActivityArrayOutput) ToActivityArrayOutputWithContext(ctx context.Context) ActivityArrayOutput {
	return o
}

func (o ActivityArrayOutput) Index(i pulumi.IntInput) ActivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Activity {
		return vs[0].([]Activity)[vs[1].(int)]
	}).(ActivityOutput)
}

type ActivityMapOutput struct{ *pulumi.OutputState }

func (ActivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Activity)(nil))
}

func (o ActivityMapOutput) ToActivityMapOutput() ActivityMapOutput {
	return o
}

func (o ActivityMapOutput) ToActivityMapOutputWithContext(ctx context.Context) ActivityMapOutput {
	return o
}

func (o ActivityMapOutput) MapIndex(k pulumi.StringInput) ActivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Activity {
		return vs[0].(map[string]Activity)[vs[1].(string)]
	}).(ActivityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityInput)(nil)).Elem(), &Activity{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityPtrInput)(nil)).Elem(), &Activity{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityArrayInput)(nil)).Elem(), ActivityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityMapInput)(nil)).Elem(), ActivityMap{})
	pulumi.RegisterOutputType(ActivityOutput{})
	pulumi.RegisterOutputType(ActivityPtrOutput{})
	pulumi.RegisterOutputType(ActivityArrayOutput{})
	pulumi.RegisterOutputType(ActivityMapOutput{})
}
