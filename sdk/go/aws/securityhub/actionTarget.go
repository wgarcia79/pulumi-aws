// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates Security Hub custom action.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/securityhub"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = securityhub.NewActionTarget(ctx, "exampleActionTarget", &securityhub.ActionTargetArgs{
// 			Identifier:  pulumi.String("SendToChat"),
// 			Description: pulumi.String("This is custom action sends selected findings to chat"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleAccount,
// 		}))
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
// Security Hub custom action can be imported using the action target ARN e.g.
//
// ```sh
//  $ pulumi import aws:securityhub/actionTarget:ActionTarget example arn:aws:securityhub:eu-west-1:312940875350:action/custom/a
// ```
type ActionTarget struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the custom action target.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ID for the custom action target.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// The description for the custom action target.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewActionTarget registers a new resource with the given unique name, arguments, and options.
func NewActionTarget(ctx *pulumi.Context,
	name string, args *ActionTargetArgs, opts ...pulumi.ResourceOption) (*ActionTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	var resource ActionTarget
	err := ctx.RegisterResource("aws:securityhub/actionTarget:ActionTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActionTarget gets an existing ActionTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActionTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActionTargetState, opts ...pulumi.ResourceOption) (*ActionTarget, error) {
	var resource ActionTarget
	err := ctx.ReadResource("aws:securityhub/actionTarget:ActionTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActionTarget resources.
type actionTargetState struct {
	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn *string `pulumi:"arn"`
	// The name of the custom action target.
	Description *string `pulumi:"description"`
	// The ID for the custom action target.
	Identifier *string `pulumi:"identifier"`
	// The description for the custom action target.
	Name *string `pulumi:"name"`
}

type ActionTargetState struct {
	// Amazon Resource Name (ARN) of the Security Hub custom action target.
	Arn pulumi.StringPtrInput
	// The name of the custom action target.
	Description pulumi.StringPtrInput
	// The ID for the custom action target.
	Identifier pulumi.StringPtrInput
	// The description for the custom action target.
	Name pulumi.StringPtrInput
}

func (ActionTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*actionTargetState)(nil)).Elem()
}

type actionTargetArgs struct {
	// The name of the custom action target.
	Description string `pulumi:"description"`
	// The ID for the custom action target.
	Identifier string `pulumi:"identifier"`
	// The description for the custom action target.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ActionTarget resource.
type ActionTargetArgs struct {
	// The name of the custom action target.
	Description pulumi.StringInput
	// The ID for the custom action target.
	Identifier pulumi.StringInput
	// The description for the custom action target.
	Name pulumi.StringPtrInput
}

func (ActionTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*actionTargetArgs)(nil)).Elem()
}

type ActionTargetInput interface {
	pulumi.Input

	ToActionTargetOutput() ActionTargetOutput
	ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput
}

func (*ActionTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionTarget)(nil))
}

func (i *ActionTarget) ToActionTargetOutput() ActionTargetOutput {
	return i.ToActionTargetOutputWithContext(context.Background())
}

func (i *ActionTarget) ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetOutput)
}

func (i *ActionTarget) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return i.ToActionTargetPtrOutputWithContext(context.Background())
}

func (i *ActionTarget) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetPtrOutput)
}

type ActionTargetPtrInput interface {
	pulumi.Input

	ToActionTargetPtrOutput() ActionTargetPtrOutput
	ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput
}

type actionTargetPtrType ActionTargetArgs

func (*actionTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionTarget)(nil))
}

func (i *actionTargetPtrType) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return i.ToActionTargetPtrOutputWithContext(context.Background())
}

func (i *actionTargetPtrType) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetPtrOutput)
}

// ActionTargetArrayInput is an input type that accepts ActionTargetArray and ActionTargetArrayOutput values.
// You can construct a concrete instance of `ActionTargetArrayInput` via:
//
//          ActionTargetArray{ ActionTargetArgs{...} }
type ActionTargetArrayInput interface {
	pulumi.Input

	ToActionTargetArrayOutput() ActionTargetArrayOutput
	ToActionTargetArrayOutputWithContext(context.Context) ActionTargetArrayOutput
}

type ActionTargetArray []ActionTargetInput

func (ActionTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ActionTarget)(nil)).Elem()
}

func (i ActionTargetArray) ToActionTargetArrayOutput() ActionTargetArrayOutput {
	return i.ToActionTargetArrayOutputWithContext(context.Background())
}

func (i ActionTargetArray) ToActionTargetArrayOutputWithContext(ctx context.Context) ActionTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetArrayOutput)
}

// ActionTargetMapInput is an input type that accepts ActionTargetMap and ActionTargetMapOutput values.
// You can construct a concrete instance of `ActionTargetMapInput` via:
//
//          ActionTargetMap{ "key": ActionTargetArgs{...} }
type ActionTargetMapInput interface {
	pulumi.Input

	ToActionTargetMapOutput() ActionTargetMapOutput
	ToActionTargetMapOutputWithContext(context.Context) ActionTargetMapOutput
}

type ActionTargetMap map[string]ActionTargetInput

func (ActionTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ActionTarget)(nil)).Elem()
}

func (i ActionTargetMap) ToActionTargetMapOutput() ActionTargetMapOutput {
	return i.ToActionTargetMapOutputWithContext(context.Background())
}

func (i ActionTargetMap) ToActionTargetMapOutputWithContext(ctx context.Context) ActionTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionTargetMapOutput)
}

type ActionTargetOutput struct{ *pulumi.OutputState }

func (ActionTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionTarget)(nil))
}

func (o ActionTargetOutput) ToActionTargetOutput() ActionTargetOutput {
	return o
}

func (o ActionTargetOutput) ToActionTargetOutputWithContext(ctx context.Context) ActionTargetOutput {
	return o
}

func (o ActionTargetOutput) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return o.ToActionTargetPtrOutputWithContext(context.Background())
}

func (o ActionTargetOutput) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionTarget) *ActionTarget {
		return &v
	}).(ActionTargetPtrOutput)
}

type ActionTargetPtrOutput struct{ *pulumi.OutputState }

func (ActionTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionTarget)(nil))
}

func (o ActionTargetPtrOutput) ToActionTargetPtrOutput() ActionTargetPtrOutput {
	return o
}

func (o ActionTargetPtrOutput) ToActionTargetPtrOutputWithContext(ctx context.Context) ActionTargetPtrOutput {
	return o
}

func (o ActionTargetPtrOutput) Elem() ActionTargetOutput {
	return o.ApplyT(func(v *ActionTarget) ActionTarget {
		if v != nil {
			return *v
		}
		var ret ActionTarget
		return ret
	}).(ActionTargetOutput)
}

type ActionTargetArrayOutput struct{ *pulumi.OutputState }

func (ActionTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionTarget)(nil))
}

func (o ActionTargetArrayOutput) ToActionTargetArrayOutput() ActionTargetArrayOutput {
	return o
}

func (o ActionTargetArrayOutput) ToActionTargetArrayOutputWithContext(ctx context.Context) ActionTargetArrayOutput {
	return o
}

func (o ActionTargetArrayOutput) Index(i pulumi.IntInput) ActionTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionTarget {
		return vs[0].([]ActionTarget)[vs[1].(int)]
	}).(ActionTargetOutput)
}

type ActionTargetMapOutput struct{ *pulumi.OutputState }

func (ActionTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ActionTarget)(nil))
}

func (o ActionTargetMapOutput) ToActionTargetMapOutput() ActionTargetMapOutput {
	return o
}

func (o ActionTargetMapOutput) ToActionTargetMapOutputWithContext(ctx context.Context) ActionTargetMapOutput {
	return o
}

func (o ActionTargetMapOutput) MapIndex(k pulumi.StringInput) ActionTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ActionTarget {
		return vs[0].(map[string]ActionTarget)[vs[1].(string)]
	}).(ActionTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionTargetInput)(nil)).Elem(), &ActionTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionTargetPtrInput)(nil)).Elem(), &ActionTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionTargetArrayInput)(nil)).Elem(), ActionTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionTargetMapInput)(nil)).Elem(), ActionTargetMap{})
	pulumi.RegisterOutputType(ActionTargetOutput{})
	pulumi.RegisterOutputType(ActionTargetPtrOutput{})
	pulumi.RegisterOutputType(ActionTargetArrayOutput{})
	pulumi.RegisterOutputType(ActionTargetMapOutput{})
}
