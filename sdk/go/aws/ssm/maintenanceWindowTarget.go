// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Maintenance Window Target resource
//
// ## Example Usage
// ### Instance Target
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		window, err := ssm.NewMaintenanceWindow(ctx, "window", &ssm.MaintenanceWindowArgs{
// 			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
// 			Duration: pulumi.Int(3),
// 			Cutoff:   pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewMaintenanceWindowTarget(ctx, "target1", &ssm.MaintenanceWindowTargetArgs{
// 			WindowId:     window.ID(),
// 			Description:  pulumi.String("This is a maintenance window target"),
// 			ResourceType: pulumi.String("INSTANCE"),
// 			Targets: ssm.MaintenanceWindowTargetTargetArray{
// 				&ssm.MaintenanceWindowTargetTargetArgs{
// 					Key: pulumi.String("tag:Name"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("acceptance_test"),
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
// ### Resource Group Target
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		window, err := ssm.NewMaintenanceWindow(ctx, "window", &ssm.MaintenanceWindowArgs{
// 			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
// 			Duration: pulumi.Int(3),
// 			Cutoff:   pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewMaintenanceWindowTarget(ctx, "target1", &ssm.MaintenanceWindowTargetArgs{
// 			WindowId:     window.ID(),
// 			Description:  pulumi.String("This is a maintenance window target"),
// 			ResourceType: pulumi.String("RESOURCE_GROUP"),
// 			Targets: ssm.MaintenanceWindowTargetTargetArray{
// 				&ssm.MaintenanceWindowTargetTargetArgs{
// 					Key: pulumi.String("resource-groups:ResourceTypeFilters"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("AWS::EC2::Instance"),
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
//
// ## Import
//
// SSM Maintenance Window targets can be imported using `WINDOW_ID/WINDOW_TARGET_ID`, e.g.
//
// ```sh
//  $ pulumi import aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
// ```
type MaintenanceWindowTarget struct {
	pulumi.CustomResourceState

	// The description of the maintenance window target.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the maintenance window target.
	Name pulumi.StringOutput `pulumi:"name"`
	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation pulumi.StringPtrOutput `pulumi:"ownerInformation"`
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets MaintenanceWindowTargetTargetArrayOutput `pulumi:"targets"`
	// The Id of the maintenance window to register the target with.
	WindowId pulumi.StringOutput `pulumi:"windowId"`
}

// NewMaintenanceWindowTarget registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindowTarget(ctx *pulumi.Context,
	name string, args *MaintenanceWindowTargetArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindowTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	if args.WindowId == nil {
		return nil, errors.New("invalid value for required argument 'WindowId'")
	}
	var resource MaintenanceWindowTarget
	err := ctx.RegisterResource("aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindowTarget gets an existing MaintenanceWindowTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindowTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowTargetState, opts ...pulumi.ResourceOption) (*MaintenanceWindowTarget, error) {
	var resource MaintenanceWindowTarget
	err := ctx.ReadResource("aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindowTarget resources.
type maintenanceWindowTargetState struct {
	// The description of the maintenance window target.
	Description *string `pulumi:"description"`
	// The name of the maintenance window target.
	Name *string `pulumi:"name"`
	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation *string `pulumi:"ownerInformation"`
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType *string `pulumi:"resourceType"`
	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets []MaintenanceWindowTargetTarget `pulumi:"targets"`
	// The Id of the maintenance window to register the target with.
	WindowId *string `pulumi:"windowId"`
}

type MaintenanceWindowTargetState struct {
	// The description of the maintenance window target.
	Description pulumi.StringPtrInput
	// The name of the maintenance window target.
	Name pulumi.StringPtrInput
	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation pulumi.StringPtrInput
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType pulumi.StringPtrInput
	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets MaintenanceWindowTargetTargetArrayInput
	// The Id of the maintenance window to register the target with.
	WindowId pulumi.StringPtrInput
}

func (MaintenanceWindowTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTargetState)(nil)).Elem()
}

type maintenanceWindowTargetArgs struct {
	// The description of the maintenance window target.
	Description *string `pulumi:"description"`
	// The name of the maintenance window target.
	Name *string `pulumi:"name"`
	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation *string `pulumi:"ownerInformation"`
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType string `pulumi:"resourceType"`
	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets []MaintenanceWindowTargetTarget `pulumi:"targets"`
	// The Id of the maintenance window to register the target with.
	WindowId string `pulumi:"windowId"`
}

// The set of arguments for constructing a MaintenanceWindowTarget resource.
type MaintenanceWindowTargetArgs struct {
	// The description of the maintenance window target.
	Description pulumi.StringPtrInput
	// The name of the maintenance window target.
	Name pulumi.StringPtrInput
	// User-provided value that will be included in any CloudWatch events raised while running tasks for these targets in this Maintenance Window.
	OwnerInformation pulumi.StringPtrInput
	// The type of target being registered with the Maintenance Window. Possible values are `INSTANCE` and `RESOURCE_GROUP`.
	ResourceType pulumi.StringInput
	// The targets to register with the maintenance window. In other words, the instances to run commands on when the maintenance window runs. You can specify targets using instance IDs, resource group names, or tags that have been applied to instances. For more information about these examples formats see
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/mw-cli-tutorial-targets-examples.html)
	Targets MaintenanceWindowTargetTargetArrayInput
	// The Id of the maintenance window to register the target with.
	WindowId pulumi.StringInput
}

func (MaintenanceWindowTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTargetArgs)(nil)).Elem()
}

type MaintenanceWindowTargetInput interface {
	pulumi.Input

	ToMaintenanceWindowTargetOutput() MaintenanceWindowTargetOutput
	ToMaintenanceWindowTargetOutputWithContext(ctx context.Context) MaintenanceWindowTargetOutput
}

func (*MaintenanceWindowTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTarget)(nil))
}

func (i *MaintenanceWindowTarget) ToMaintenanceWindowTargetOutput() MaintenanceWindowTargetOutput {
	return i.ToMaintenanceWindowTargetOutputWithContext(context.Background())
}

func (i *MaintenanceWindowTarget) ToMaintenanceWindowTargetOutputWithContext(ctx context.Context) MaintenanceWindowTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTargetOutput)
}

func (i *MaintenanceWindowTarget) ToMaintenanceWindowTargetPtrOutput() MaintenanceWindowTargetPtrOutput {
	return i.ToMaintenanceWindowTargetPtrOutputWithContext(context.Background())
}

func (i *MaintenanceWindowTarget) ToMaintenanceWindowTargetPtrOutputWithContext(ctx context.Context) MaintenanceWindowTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTargetPtrOutput)
}

type MaintenanceWindowTargetPtrInput interface {
	pulumi.Input

	ToMaintenanceWindowTargetPtrOutput() MaintenanceWindowTargetPtrOutput
	ToMaintenanceWindowTargetPtrOutputWithContext(ctx context.Context) MaintenanceWindowTargetPtrOutput
}

type maintenanceWindowTargetPtrType MaintenanceWindowTargetArgs

func (*maintenanceWindowTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowTarget)(nil))
}

func (i *maintenanceWindowTargetPtrType) ToMaintenanceWindowTargetPtrOutput() MaintenanceWindowTargetPtrOutput {
	return i.ToMaintenanceWindowTargetPtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowTargetPtrType) ToMaintenanceWindowTargetPtrOutputWithContext(ctx context.Context) MaintenanceWindowTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTargetPtrOutput)
}

// MaintenanceWindowTargetArrayInput is an input type that accepts MaintenanceWindowTargetArray and MaintenanceWindowTargetArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowTargetArrayInput` via:
//
//          MaintenanceWindowTargetArray{ MaintenanceWindowTargetArgs{...} }
type MaintenanceWindowTargetArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowTargetArrayOutput() MaintenanceWindowTargetArrayOutput
	ToMaintenanceWindowTargetArrayOutputWithContext(context.Context) MaintenanceWindowTargetArrayOutput
}

type MaintenanceWindowTargetArray []MaintenanceWindowTargetInput

func (MaintenanceWindowTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindowTarget)(nil)).Elem()
}

func (i MaintenanceWindowTargetArray) ToMaintenanceWindowTargetArrayOutput() MaintenanceWindowTargetArrayOutput {
	return i.ToMaintenanceWindowTargetArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowTargetArray) ToMaintenanceWindowTargetArrayOutputWithContext(ctx context.Context) MaintenanceWindowTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTargetArrayOutput)
}

// MaintenanceWindowTargetMapInput is an input type that accepts MaintenanceWindowTargetMap and MaintenanceWindowTargetMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowTargetMapInput` via:
//
//          MaintenanceWindowTargetMap{ "key": MaintenanceWindowTargetArgs{...} }
type MaintenanceWindowTargetMapInput interface {
	pulumi.Input

	ToMaintenanceWindowTargetMapOutput() MaintenanceWindowTargetMapOutput
	ToMaintenanceWindowTargetMapOutputWithContext(context.Context) MaintenanceWindowTargetMapOutput
}

type MaintenanceWindowTargetMap map[string]MaintenanceWindowTargetInput

func (MaintenanceWindowTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindowTarget)(nil)).Elem()
}

func (i MaintenanceWindowTargetMap) ToMaintenanceWindowTargetMapOutput() MaintenanceWindowTargetMapOutput {
	return i.ToMaintenanceWindowTargetMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowTargetMap) ToMaintenanceWindowTargetMapOutputWithContext(ctx context.Context) MaintenanceWindowTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTargetMapOutput)
}

type MaintenanceWindowTargetOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTarget)(nil))
}

func (o MaintenanceWindowTargetOutput) ToMaintenanceWindowTargetOutput() MaintenanceWindowTargetOutput {
	return o
}

func (o MaintenanceWindowTargetOutput) ToMaintenanceWindowTargetOutputWithContext(ctx context.Context) MaintenanceWindowTargetOutput {
	return o
}

func (o MaintenanceWindowTargetOutput) ToMaintenanceWindowTargetPtrOutput() MaintenanceWindowTargetPtrOutput {
	return o.ToMaintenanceWindowTargetPtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowTargetOutput) ToMaintenanceWindowTargetPtrOutputWithContext(ctx context.Context) MaintenanceWindowTargetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindowTarget) *MaintenanceWindowTarget {
		return &v
	}).(MaintenanceWindowTargetPtrOutput)
}

type MaintenanceWindowTargetPtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowTarget)(nil))
}

func (o MaintenanceWindowTargetPtrOutput) ToMaintenanceWindowTargetPtrOutput() MaintenanceWindowTargetPtrOutput {
	return o
}

func (o MaintenanceWindowTargetPtrOutput) ToMaintenanceWindowTargetPtrOutputWithContext(ctx context.Context) MaintenanceWindowTargetPtrOutput {
	return o
}

func (o MaintenanceWindowTargetPtrOutput) Elem() MaintenanceWindowTargetOutput {
	return o.ApplyT(func(v *MaintenanceWindowTarget) MaintenanceWindowTarget {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindowTarget
		return ret
	}).(MaintenanceWindowTargetOutput)
}

type MaintenanceWindowTargetArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceWindowTarget)(nil))
}

func (o MaintenanceWindowTargetArrayOutput) ToMaintenanceWindowTargetArrayOutput() MaintenanceWindowTargetArrayOutput {
	return o
}

func (o MaintenanceWindowTargetArrayOutput) ToMaintenanceWindowTargetArrayOutputWithContext(ctx context.Context) MaintenanceWindowTargetArrayOutput {
	return o
}

func (o MaintenanceWindowTargetArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MaintenanceWindowTarget {
		return vs[0].([]MaintenanceWindowTarget)[vs[1].(int)]
	}).(MaintenanceWindowTargetOutput)
}

type MaintenanceWindowTargetMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MaintenanceWindowTarget)(nil))
}

func (o MaintenanceWindowTargetMapOutput) ToMaintenanceWindowTargetMapOutput() MaintenanceWindowTargetMapOutput {
	return o
}

func (o MaintenanceWindowTargetMapOutput) ToMaintenanceWindowTargetMapOutputWithContext(ctx context.Context) MaintenanceWindowTargetMapOutput {
	return o
}

func (o MaintenanceWindowTargetMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MaintenanceWindowTarget {
		return vs[0].(map[string]MaintenanceWindowTarget)[vs[1].(string)]
	}).(MaintenanceWindowTargetOutput)
}

func init() {
	pulumi.RegisterOutputType(MaintenanceWindowTargetOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTargetPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTargetArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTargetMapOutput{})
}
