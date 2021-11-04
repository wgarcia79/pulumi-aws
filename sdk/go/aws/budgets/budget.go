// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/budgets"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := budgets.NewBudget(ctx, "ec2", &budgets.BudgetArgs{
// 			BudgetType: pulumi.String("COST"),
// 			CostFilters: pulumi.StringMap{
// 				pulumi.String{
// 					Name: "Service",
// 					Values: []string{
// 						"Amazon Elastic Compute Cloud - Compute",
// 					},
// 				},
// 			},
// 			LimitAmount: pulumi.String("1200"),
// 			LimitUnit:   pulumi.String("USD"),
// 			Notifications: budgets.BudgetNotificationArray{
// 				&budgets.BudgetNotificationArgs{
// 					ComparisonOperator: pulumi.String("GREATER_THAN"),
// 					NotificationType:   pulumi.String("FORECASTED"),
// 					SubscriberEmailAddresses: pulumi.StringArray{
// 						pulumi.String("test@example.com"),
// 					},
// 					Threshold:     pulumi.Float64(100),
// 					ThresholdType: pulumi.String("PERCENTAGE"),
// 				},
// 			},
// 			TimePeriodEnd:   pulumi.String("2087-06-15_00:00"),
// 			TimePeriodStart: pulumi.String("2017-07-01_00:00"),
// 			TimeUnit:        pulumi.String("MONTHLY"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Create a budget for *$100*.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/budgets"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := budgets.NewBudget(ctx, "cost", &budgets.BudgetArgs{
// 			BudgetType:  pulumi.String("COST"),
// 			LimitAmount: pulumi.String("100"),
// 			LimitUnit:   pulumi.String("USD"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Create a budget for s3 with a limit of *3 GB* of storage.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/budgets"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := budgets.NewBudget(ctx, "s3", &budgets.BudgetArgs{
// 			BudgetType:  pulumi.String("USAGE"),
// 			LimitAmount: pulumi.String("3"),
// 			LimitUnit:   pulumi.String("GB"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Create a Savings Plan Utilization Budget
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/budgets"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := budgets.NewBudget(ctx, "savingsPlanUtilization", &budgets.BudgetArgs{
// 			BudgetType: pulumi.String("SAVINGS_PLANS_UTILIZATION"),
// 			CostTypes: &budgets.BudgetCostTypesArgs{
// 				IncludeCredit:            pulumi.Bool(false),
// 				IncludeDiscount:          pulumi.Bool(false),
// 				IncludeOtherSubscription: pulumi.Bool(false),
// 				IncludeRecurring:         pulumi.Bool(false),
// 				IncludeRefund:            pulumi.Bool(false),
// 				IncludeSubscription:      pulumi.Bool(true),
// 				IncludeSupport:           pulumi.Bool(false),
// 				IncludeTax:               pulumi.Bool(false),
// 				IncludeUpfront:           pulumi.Bool(false),
// 				UseBlended:               pulumi.Bool(false),
// 			},
// 			LimitAmount: pulumi.String("100.0"),
// 			LimitUnit:   pulumi.String("PERCENTAGE"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Create a RI Utilization Budget
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/budgets"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := budgets.NewBudget(ctx, "riUtilization", &budgets.BudgetArgs{
// 			BudgetType: pulumi.String("RI_UTILIZATION"),
// 			CostFilters: pulumi.StringMap{
// 				"Service": pulumi.String("Amazon Relational Database Service"),
// 			},
// 			CostTypes: &budgets.BudgetCostTypesArgs{
// 				IncludeCredit:            pulumi.Bool(false),
// 				IncludeDiscount:          pulumi.Bool(false),
// 				IncludeOtherSubscription: pulumi.Bool(false),
// 				IncludeRecurring:         pulumi.Bool(false),
// 				IncludeRefund:            pulumi.Bool(false),
// 				IncludeSubscription:      pulumi.Bool(true),
// 				IncludeSupport:           pulumi.Bool(false),
// 				IncludeTax:               pulumi.Bool(false),
// 				IncludeUpfront:           pulumi.Bool(false),
// 				UseBlended:               pulumi.Bool(false),
// 			},
// 			LimitAmount: pulumi.String("100.0"),
// 			LimitUnit:   pulumi.String("PERCENTAGE"),
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
// Budgets can be imported using `AccountID:BudgetName`, e.g.
//
// ```sh
//  $ pulumi import aws:budgets/budget:Budget myBudget 123456789012:myBudget`
// ```
type Budget struct {
	pulumi.CustomResourceState

	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The ARN of the budget.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringOutput `pulumi:"budgetType"`
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.StringMapOutput `pulumi:"costFilters"`
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes BudgetCostTypesOutput `pulumi:"costTypes"`
	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringOutput `pulumi:"limitAmount"`
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringOutput `pulumi:"limitUnit"`
	// The name of a budget. Unique within accounts.
	Name pulumi.StringOutput `pulumi:"name"`
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications BudgetNotificationArrayOutput `pulumi:"notifications"`
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringPtrOutput `pulumi:"timePeriodEnd"`
	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringOutput `pulumi:"timePeriodStart"`
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit pulumi.StringOutput `pulumi:"timeUnit"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOption) (*Budget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BudgetType == nil {
		return nil, errors.New("invalid value for required argument 'BudgetType'")
	}
	if args.LimitAmount == nil {
		return nil, errors.New("invalid value for required argument 'LimitAmount'")
	}
	if args.LimitUnit == nil {
		return nil, errors.New("invalid value for required argument 'LimitUnit'")
	}
	if args.TimeUnit == nil {
		return nil, errors.New("invalid value for required argument 'TimeUnit'")
	}
	var resource Budget
	err := ctx.RegisterResource("aws:budgets/budget:Budget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BudgetState, opts ...pulumi.ResourceOption) (*Budget, error) {
	var resource Budget
	err := ctx.ReadResource("aws:budgets/budget:Budget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Budget resources.
type budgetState struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId *string `pulumi:"accountId"`
	// The ARN of the budget.
	Arn *string `pulumi:"arn"`
	// Whether this budget tracks monetary cost or usage.
	BudgetType *string `pulumi:"budgetType"`
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters map[string]string `pulumi:"costFilters"`
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes *BudgetCostTypes `pulumi:"costTypes"`
	// The amount of cost or usage being measured for a budget.
	LimitAmount *string `pulumi:"limitAmount"`
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit *string `pulumi:"limitUnit"`
	// The name of a budget. Unique within accounts.
	Name *string `pulumi:"name"`
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix *string `pulumi:"namePrefix"`
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications []BudgetNotification `pulumi:"notifications"`
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd *string `pulumi:"timePeriodEnd"`
	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart *string `pulumi:"timePeriodStart"`
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit *string `pulumi:"timeUnit"`
}

type BudgetState struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringPtrInput
	// The ARN of the budget.
	Arn pulumi.StringPtrInput
	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringPtrInput
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.StringMapInput
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes BudgetCostTypesPtrInput
	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringPtrInput
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringPtrInput
	// The name of a budget. Unique within accounts.
	Name pulumi.StringPtrInput
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringPtrInput
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications BudgetNotificationArrayInput
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringPtrInput
	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringPtrInput
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit pulumi.StringPtrInput
}

func (BudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetState)(nil)).Elem()
}

type budgetArgs struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId *string `pulumi:"accountId"`
	// Whether this budget tracks monetary cost or usage.
	BudgetType string `pulumi:"budgetType"`
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters map[string]string `pulumi:"costFilters"`
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes *BudgetCostTypes `pulumi:"costTypes"`
	// The amount of cost or usage being measured for a budget.
	LimitAmount string `pulumi:"limitAmount"`
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit string `pulumi:"limitUnit"`
	// The name of a budget. Unique within accounts.
	Name *string `pulumi:"name"`
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix *string `pulumi:"namePrefix"`
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications []BudgetNotification `pulumi:"notifications"`
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd *string `pulumi:"timePeriodEnd"`
	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart *string `pulumi:"timePeriodStart"`
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit string `pulumi:"timeUnit"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringPtrInput
	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringInput
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.StringMapInput
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions.
	CostTypes BudgetCostTypesPtrInput
	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringInput
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringInput
	// The name of a budget. Unique within accounts.
	Name pulumi.StringPtrInput
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringPtrInput
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications BudgetNotificationArrayInput
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringPtrInput
	// The start of the time period covered by the budget. If you don't specify a start date, AWS defaults to the start of your chosen time period. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringPtrInput
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`, and `DAILY`.
	TimeUnit pulumi.StringInput
}

func (BudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*budgetArgs)(nil)).Elem()
}

type BudgetInput interface {
	pulumi.Input

	ToBudgetOutput() BudgetOutput
	ToBudgetOutputWithContext(ctx context.Context) BudgetOutput
}

func (*Budget) ElementType() reflect.Type {
	return reflect.TypeOf((*Budget)(nil))
}

func (i *Budget) ToBudgetOutput() BudgetOutput {
	return i.ToBudgetOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetOutput)
}

func (i *Budget) ToBudgetPtrOutput() BudgetPtrOutput {
	return i.ToBudgetPtrOutputWithContext(context.Background())
}

func (i *Budget) ToBudgetPtrOutputWithContext(ctx context.Context) BudgetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetPtrOutput)
}

type BudgetPtrInput interface {
	pulumi.Input

	ToBudgetPtrOutput() BudgetPtrOutput
	ToBudgetPtrOutputWithContext(ctx context.Context) BudgetPtrOutput
}

type budgetPtrType BudgetArgs

func (*budgetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil))
}

func (i *budgetPtrType) ToBudgetPtrOutput() BudgetPtrOutput {
	return i.ToBudgetPtrOutputWithContext(context.Background())
}

func (i *budgetPtrType) ToBudgetPtrOutputWithContext(ctx context.Context) BudgetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetPtrOutput)
}

// BudgetArrayInput is an input type that accepts BudgetArray and BudgetArrayOutput values.
// You can construct a concrete instance of `BudgetArrayInput` via:
//
//          BudgetArray{ BudgetArgs{...} }
type BudgetArrayInput interface {
	pulumi.Input

	ToBudgetArrayOutput() BudgetArrayOutput
	ToBudgetArrayOutputWithContext(context.Context) BudgetArrayOutput
}

type BudgetArray []BudgetInput

func (BudgetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Budget)(nil)).Elem()
}

func (i BudgetArray) ToBudgetArrayOutput() BudgetArrayOutput {
	return i.ToBudgetArrayOutputWithContext(context.Background())
}

func (i BudgetArray) ToBudgetArrayOutputWithContext(ctx context.Context) BudgetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetArrayOutput)
}

// BudgetMapInput is an input type that accepts BudgetMap and BudgetMapOutput values.
// You can construct a concrete instance of `BudgetMapInput` via:
//
//          BudgetMap{ "key": BudgetArgs{...} }
type BudgetMapInput interface {
	pulumi.Input

	ToBudgetMapOutput() BudgetMapOutput
	ToBudgetMapOutputWithContext(context.Context) BudgetMapOutput
}

type BudgetMap map[string]BudgetInput

func (BudgetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Budget)(nil)).Elem()
}

func (i BudgetMap) ToBudgetMapOutput() BudgetMapOutput {
	return i.ToBudgetMapOutputWithContext(context.Background())
}

func (i BudgetMap) ToBudgetMapOutputWithContext(ctx context.Context) BudgetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetMapOutput)
}

type BudgetOutput struct{ *pulumi.OutputState }

func (BudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Budget)(nil))
}

func (o BudgetOutput) ToBudgetOutput() BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetOutputWithContext(ctx context.Context) BudgetOutput {
	return o
}

func (o BudgetOutput) ToBudgetPtrOutput() BudgetPtrOutput {
	return o.ToBudgetPtrOutputWithContext(context.Background())
}

func (o BudgetOutput) ToBudgetPtrOutputWithContext(ctx context.Context) BudgetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Budget) *Budget {
		return &v
	}).(BudgetPtrOutput)
}

type BudgetPtrOutput struct{ *pulumi.OutputState }

func (BudgetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Budget)(nil))
}

func (o BudgetPtrOutput) ToBudgetPtrOutput() BudgetPtrOutput {
	return o
}

func (o BudgetPtrOutput) ToBudgetPtrOutputWithContext(ctx context.Context) BudgetPtrOutput {
	return o
}

func (o BudgetPtrOutput) Elem() BudgetOutput {
	return o.ApplyT(func(v *Budget) Budget {
		if v != nil {
			return *v
		}
		var ret Budget
		return ret
	}).(BudgetOutput)
}

type BudgetArrayOutput struct{ *pulumi.OutputState }

func (BudgetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Budget)(nil))
}

func (o BudgetArrayOutput) ToBudgetArrayOutput() BudgetArrayOutput {
	return o
}

func (o BudgetArrayOutput) ToBudgetArrayOutputWithContext(ctx context.Context) BudgetArrayOutput {
	return o
}

func (o BudgetArrayOutput) Index(i pulumi.IntInput) BudgetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Budget {
		return vs[0].([]Budget)[vs[1].(int)]
	}).(BudgetOutput)
}

type BudgetMapOutput struct{ *pulumi.OutputState }

func (BudgetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Budget)(nil))
}

func (o BudgetMapOutput) ToBudgetMapOutput() BudgetMapOutput {
	return o
}

func (o BudgetMapOutput) ToBudgetMapOutputWithContext(ctx context.Context) BudgetMapOutput {
	return o
}

func (o BudgetMapOutput) MapIndex(k pulumi.StringInput) BudgetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Budget {
		return vs[0].(map[string]Budget)[vs[1].(string)]
	}).(BudgetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetInput)(nil)).Elem(), &Budget{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetPtrInput)(nil)).Elem(), &Budget{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetArrayInput)(nil)).Elem(), BudgetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BudgetMapInput)(nil)).Elem(), BudgetMap{})
	pulumi.RegisterOutputType(BudgetOutput{})
	pulumi.RegisterOutputType(BudgetPtrOutput{})
	pulumi.RegisterOutputType(BudgetArrayOutput{})
	pulumi.RegisterOutputType(BudgetMapOutput{})
}
