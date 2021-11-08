// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Service Catalog Budget Resource Association.
//
// > **Tip:** A "resource" is either a Service Catalog portfolio or product.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicecatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicecatalog.NewBudgetResourceAssociation(ctx, "example", &servicecatalog.BudgetResourceAssociationArgs{
// 			BudgetName: pulumi.String("budget-pjtvyakdlyo3m"),
// 			ResourceId: pulumi.String("prod-dnigbtea24ste"),
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
// `aws_servicecatalog_budget_resource_association` can be imported using the budget name and resource ID, e.g.,
//
// ```sh
//  $ pulumi import aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation example budget-pjtvyakdlyo3m:prod-dnigbtea24ste
// ```
type BudgetResourceAssociation struct {
	pulumi.CustomResourceState

	// Budget name.
	BudgetName pulumi.StringOutput `pulumi:"budgetName"`
	// Resource identifier.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewBudgetResourceAssociation registers a new resource with the given unique name, arguments, and options.
func NewBudgetResourceAssociation(ctx *pulumi.Context,
	name string, args *BudgetResourceAssociationArgs, opts ...pulumi.ResourceOption) (*BudgetResourceAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BudgetName == nil {
		return nil, errors.New("invalid value for required argument 'BudgetName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource BudgetResourceAssociation
	err := ctx.RegisterResource("aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudgetResourceAssociation gets an existing BudgetResourceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudgetResourceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetResourceAssociationState, opts ...pulumi.ResourceOption) (*BudgetResourceAssociation, error) {
	var resource BudgetResourceAssociation
	err := ctx.ReadResource("aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BudgetResourceAssociation resources.
type budgetResourceAssociationState struct {
	// Budget name.
	BudgetName *string `pulumi:"budgetName"`
	// Resource identifier.
	ResourceId *string `pulumi:"resourceId"`
}

type BudgetResourceAssociationState struct {
	// Budget name.
	BudgetName pulumi.StringPtrInput
	// Resource identifier.
	ResourceId pulumi.StringPtrInput
}

func (BudgetResourceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetResourceAssociationState)(nil)).Elem()
}

type budgetResourceAssociationArgs struct {
	// Budget name.
	BudgetName string `pulumi:"budgetName"`
	// Resource identifier.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a BudgetResourceAssociation resource.
type BudgetResourceAssociationArgs struct {
	// Budget name.
	BudgetName pulumi.StringInput
	// Resource identifier.
	ResourceId pulumi.StringInput
}

func (BudgetResourceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetResourceAssociationArgs)(nil)).Elem()
}

type BudgetResourceAssociationInput interface {
	pulumi.Input

	ToBudgetResourceAssociationOutput() BudgetResourceAssociationOutput
	ToBudgetResourceAssociationOutputWithContext(ctx context.Context) BudgetResourceAssociationOutput
}

func (*BudgetResourceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetResourceAssociation)(nil))
}

func (i *BudgetResourceAssociation) ToBudgetResourceAssociationOutput() BudgetResourceAssociationOutput {
	return i.ToBudgetResourceAssociationOutputWithContext(context.Background())
}

func (i *BudgetResourceAssociation) ToBudgetResourceAssociationOutputWithContext(ctx context.Context) BudgetResourceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceAssociationOutput)
}

func (i *BudgetResourceAssociation) ToBudgetResourceAssociationPtrOutput() BudgetResourceAssociationPtrOutput {
	return i.ToBudgetResourceAssociationPtrOutputWithContext(context.Background())
}

func (i *BudgetResourceAssociation) ToBudgetResourceAssociationPtrOutputWithContext(ctx context.Context) BudgetResourceAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceAssociationPtrOutput)
}

type BudgetResourceAssociationPtrInput interface {
	pulumi.Input

	ToBudgetResourceAssociationPtrOutput() BudgetResourceAssociationPtrOutput
	ToBudgetResourceAssociationPtrOutputWithContext(ctx context.Context) BudgetResourceAssociationPtrOutput
}

type budgetResourceAssociationPtrType BudgetResourceAssociationArgs

func (*budgetResourceAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetResourceAssociation)(nil))
}

func (i *budgetResourceAssociationPtrType) ToBudgetResourceAssociationPtrOutput() BudgetResourceAssociationPtrOutput {
	return i.ToBudgetResourceAssociationPtrOutputWithContext(context.Background())
}

func (i *budgetResourceAssociationPtrType) ToBudgetResourceAssociationPtrOutputWithContext(ctx context.Context) BudgetResourceAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceAssociationPtrOutput)
}

// BudgetResourceAssociationArrayInput is an input type that accepts BudgetResourceAssociationArray and BudgetResourceAssociationArrayOutput values.
// You can construct a concrete instance of `BudgetResourceAssociationArrayInput` via:
//
//          BudgetResourceAssociationArray{ BudgetResourceAssociationArgs{...} }
type BudgetResourceAssociationArrayInput interface {
	pulumi.Input

	ToBudgetResourceAssociationArrayOutput() BudgetResourceAssociationArrayOutput
	ToBudgetResourceAssociationArrayOutputWithContext(context.Context) BudgetResourceAssociationArrayOutput
}

type BudgetResourceAssociationArray []BudgetResourceAssociationInput

func (BudgetResourceAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BudgetResourceAssociation)(nil)).Elem()
}

func (i BudgetResourceAssociationArray) ToBudgetResourceAssociationArrayOutput() BudgetResourceAssociationArrayOutput {
	return i.ToBudgetResourceAssociationArrayOutputWithContext(context.Background())
}

func (i BudgetResourceAssociationArray) ToBudgetResourceAssociationArrayOutputWithContext(ctx context.Context) BudgetResourceAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceAssociationArrayOutput)
}

// BudgetResourceAssociationMapInput is an input type that accepts BudgetResourceAssociationMap and BudgetResourceAssociationMapOutput values.
// You can construct a concrete instance of `BudgetResourceAssociationMapInput` via:
//
//          BudgetResourceAssociationMap{ "key": BudgetResourceAssociationArgs{...} }
type BudgetResourceAssociationMapInput interface {
	pulumi.Input

	ToBudgetResourceAssociationMapOutput() BudgetResourceAssociationMapOutput
	ToBudgetResourceAssociationMapOutputWithContext(context.Context) BudgetResourceAssociationMapOutput
}

type BudgetResourceAssociationMap map[string]BudgetResourceAssociationInput

func (BudgetResourceAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BudgetResourceAssociation)(nil)).Elem()
}

func (i BudgetResourceAssociationMap) ToBudgetResourceAssociationMapOutput() BudgetResourceAssociationMapOutput {
	return i.ToBudgetResourceAssociationMapOutputWithContext(context.Background())
}

func (i BudgetResourceAssociationMap) ToBudgetResourceAssociationMapOutputWithContext(ctx context.Context) BudgetResourceAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetResourceAssociationMapOutput)
}

type BudgetResourceAssociationOutput struct{ *pulumi.OutputState }

func (BudgetResourceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetResourceAssociation)(nil))
}

func (o BudgetResourceAssociationOutput) ToBudgetResourceAssociationOutput() BudgetResourceAssociationOutput {
	return o
}

func (o BudgetResourceAssociationOutput) ToBudgetResourceAssociationOutputWithContext(ctx context.Context) BudgetResourceAssociationOutput {
	return o
}

func (o BudgetResourceAssociationOutput) ToBudgetResourceAssociationPtrOutput() BudgetResourceAssociationPtrOutput {
	return o.ToBudgetResourceAssociationPtrOutputWithContext(context.Background())
}

func (o BudgetResourceAssociationOutput) ToBudgetResourceAssociationPtrOutputWithContext(ctx context.Context) BudgetResourceAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetResourceAssociation) *BudgetResourceAssociation {
		return &v
	}).(BudgetResourceAssociationPtrOutput)
}

type BudgetResourceAssociationPtrOutput struct{ *pulumi.OutputState }

func (BudgetResourceAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetResourceAssociation)(nil))
}

func (o BudgetResourceAssociationPtrOutput) ToBudgetResourceAssociationPtrOutput() BudgetResourceAssociationPtrOutput {
	return o
}

func (o BudgetResourceAssociationPtrOutput) ToBudgetResourceAssociationPtrOutputWithContext(ctx context.Context) BudgetResourceAssociationPtrOutput {
	return o
}

func (o BudgetResourceAssociationPtrOutput) Elem() BudgetResourceAssociationOutput {
	return o.ApplyT(func(v *BudgetResourceAssociation) BudgetResourceAssociation {
		if v != nil {
			return *v
		}
		var ret BudgetResourceAssociation
		return ret
	}).(BudgetResourceAssociationOutput)
}

type BudgetResourceAssociationArrayOutput struct{ *pulumi.OutputState }

func (BudgetResourceAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BudgetResourceAssociation)(nil))
}

func (o BudgetResourceAssociationArrayOutput) ToBudgetResourceAssociationArrayOutput() BudgetResourceAssociationArrayOutput {
	return o
}

func (o BudgetResourceAssociationArrayOutput) ToBudgetResourceAssociationArrayOutputWithContext(ctx context.Context) BudgetResourceAssociationArrayOutput {
	return o
}

func (o BudgetResourceAssociationArrayOutput) Index(i pulumi.IntInput) BudgetResourceAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BudgetResourceAssociation {
		return vs[0].([]BudgetResourceAssociation)[vs[1].(int)]
	}).(BudgetResourceAssociationOutput)
}

type BudgetResourceAssociationMapOutput struct{ *pulumi.OutputState }

func (BudgetResourceAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BudgetResourceAssociation)(nil))
}

func (o BudgetResourceAssociationMapOutput) ToBudgetResourceAssociationMapOutput() BudgetResourceAssociationMapOutput {
	return o
}

func (o BudgetResourceAssociationMapOutput) ToBudgetResourceAssociationMapOutputWithContext(ctx context.Context) BudgetResourceAssociationMapOutput {
	return o
}

func (o BudgetResourceAssociationMapOutput) MapIndex(k pulumi.StringInput) BudgetResourceAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BudgetResourceAssociation {
		return vs[0].(map[string]BudgetResourceAssociation)[vs[1].(string)]
	}).(BudgetResourceAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceAssociationInput)(nil)).Elem(), &BudgetResourceAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceAssociationPtrInput)(nil)).Elem(), &BudgetResourceAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceAssociationArrayInput)(nil)).Elem(), BudgetResourceAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetResourceAssociationMapInput)(nil)).Elem(), BudgetResourceAssociationMap{})
	pulumi.RegisterOutputType(BudgetResourceAssociationOutput{})
	pulumi.RegisterOutputType(BudgetResourceAssociationPtrOutput{})
	pulumi.RegisterOutputType(BudgetResourceAssociationArrayOutput{})
	pulumi.RegisterOutputType(BudgetResourceAssociationMapOutput{})
}
