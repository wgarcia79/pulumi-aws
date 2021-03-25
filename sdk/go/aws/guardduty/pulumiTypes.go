// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilterFindingCriteria struct {
	Criterions []FilterFindingCriteriaCriterion `pulumi:"criterions"`
}

// FilterFindingCriteriaInput is an input type that accepts FilterFindingCriteriaArgs and FilterFindingCriteriaOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaInput` via:
//
//          FilterFindingCriteriaArgs{...}
type FilterFindingCriteriaInput interface {
	pulumi.Input

	ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput
	ToFilterFindingCriteriaOutputWithContext(context.Context) FilterFindingCriteriaOutput
}

type FilterFindingCriteriaArgs struct {
	Criterions FilterFindingCriteriaCriterionArrayInput `pulumi:"criterions"`
}

func (FilterFindingCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteria)(nil)).Elem()
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput {
	return i.ToFilterFindingCriteriaOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaOutputWithContext(ctx context.Context) FilterFindingCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaOutput)
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return i.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaArgs) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaOutput).ToFilterFindingCriteriaPtrOutputWithContext(ctx)
}

// FilterFindingCriteriaPtrInput is an input type that accepts FilterFindingCriteriaArgs, FilterFindingCriteriaPtr and FilterFindingCriteriaPtrOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaPtrInput` via:
//
//          FilterFindingCriteriaArgs{...}
//
//  or:
//
//          nil
type FilterFindingCriteriaPtrInput interface {
	pulumi.Input

	ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput
	ToFilterFindingCriteriaPtrOutputWithContext(context.Context) FilterFindingCriteriaPtrOutput
}

type filterFindingCriteriaPtrType FilterFindingCriteriaArgs

func FilterFindingCriteriaPtr(v *FilterFindingCriteriaArgs) FilterFindingCriteriaPtrInput {
	return (*filterFindingCriteriaPtrType)(v)
}

func (*filterFindingCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterFindingCriteria)(nil)).Elem()
}

func (i *filterFindingCriteriaPtrType) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return i.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i *filterFindingCriteriaPtrType) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaPtrOutput)
}

type FilterFindingCriteriaOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteria)(nil)).Elem()
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaOutput() FilterFindingCriteriaOutput {
	return o
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaOutputWithContext(ctx context.Context) FilterFindingCriteriaOutput {
	return o
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return o.ToFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (o FilterFindingCriteriaOutput) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteria) *FilterFindingCriteria {
		return &v
	}).(FilterFindingCriteriaPtrOutput)
}
func (o FilterFindingCriteriaOutput) Criterions() FilterFindingCriteriaCriterionArrayOutput {
	return o.ApplyT(func(v FilterFindingCriteria) []FilterFindingCriteriaCriterion { return v.Criterions }).(FilterFindingCriteriaCriterionArrayOutput)
}

type FilterFindingCriteriaPtrOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterFindingCriteria)(nil)).Elem()
}

func (o FilterFindingCriteriaPtrOutput) ToFilterFindingCriteriaPtrOutput() FilterFindingCriteriaPtrOutput {
	return o
}

func (o FilterFindingCriteriaPtrOutput) ToFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FilterFindingCriteriaPtrOutput {
	return o
}

func (o FilterFindingCriteriaPtrOutput) Elem() FilterFindingCriteriaOutput {
	return o.ApplyT(func(v *FilterFindingCriteria) FilterFindingCriteria { return *v }).(FilterFindingCriteriaOutput)
}

func (o FilterFindingCriteriaPtrOutput) Criterions() FilterFindingCriteriaCriterionArrayOutput {
	return o.ApplyT(func(v *FilterFindingCriteria) []FilterFindingCriteriaCriterion {
		if v == nil {
			return nil
		}
		return v.Criterions
	}).(FilterFindingCriteriaCriterionArrayOutput)
}

type FilterFindingCriteriaCriterion struct {
	// List of string values to be evaluated.
	Equals []string `pulumi:"equals"`
	// The name of the field to be evaluated. The full list of field names can be found in [AWS documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria).
	Field string `pulumi:"field"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThan *string `pulumi:"greaterThan"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThanOrEqual *string `pulumi:"greaterThanOrEqual"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThan *string `pulumi:"lessThan"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThanOrEqual *string `pulumi:"lessThanOrEqual"`
	// List of string values to be evaluated.
	NotEquals []string `pulumi:"notEquals"`
}

// FilterFindingCriteriaCriterionInput is an input type that accepts FilterFindingCriteriaCriterionArgs and FilterFindingCriteriaCriterionOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaCriterionInput` via:
//
//          FilterFindingCriteriaCriterionArgs{...}
type FilterFindingCriteriaCriterionInput interface {
	pulumi.Input

	ToFilterFindingCriteriaCriterionOutput() FilterFindingCriteriaCriterionOutput
	ToFilterFindingCriteriaCriterionOutputWithContext(context.Context) FilterFindingCriteriaCriterionOutput
}

type FilterFindingCriteriaCriterionArgs struct {
	// List of string values to be evaluated.
	Equals pulumi.StringArrayInput `pulumi:"equals"`
	// The name of the field to be evaluated. The full list of field names can be found in [AWS documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria).
	Field pulumi.StringInput `pulumi:"field"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThan pulumi.StringPtrInput `pulumi:"greaterThan"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	GreaterThanOrEqual pulumi.StringPtrInput `pulumi:"greaterThanOrEqual"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThan pulumi.StringPtrInput `pulumi:"lessThan"`
	// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	LessThanOrEqual pulumi.StringPtrInput `pulumi:"lessThanOrEqual"`
	// List of string values to be evaluated.
	NotEquals pulumi.StringArrayInput `pulumi:"notEquals"`
}

func (FilterFindingCriteriaCriterionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteriaCriterion)(nil)).Elem()
}

func (i FilterFindingCriteriaCriterionArgs) ToFilterFindingCriteriaCriterionOutput() FilterFindingCriteriaCriterionOutput {
	return i.ToFilterFindingCriteriaCriterionOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaCriterionArgs) ToFilterFindingCriteriaCriterionOutputWithContext(ctx context.Context) FilterFindingCriteriaCriterionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaCriterionOutput)
}

// FilterFindingCriteriaCriterionArrayInput is an input type that accepts FilterFindingCriteriaCriterionArray and FilterFindingCriteriaCriterionArrayOutput values.
// You can construct a concrete instance of `FilterFindingCriteriaCriterionArrayInput` via:
//
//          FilterFindingCriteriaCriterionArray{ FilterFindingCriteriaCriterionArgs{...} }
type FilterFindingCriteriaCriterionArrayInput interface {
	pulumi.Input

	ToFilterFindingCriteriaCriterionArrayOutput() FilterFindingCriteriaCriterionArrayOutput
	ToFilterFindingCriteriaCriterionArrayOutputWithContext(context.Context) FilterFindingCriteriaCriterionArrayOutput
}

type FilterFindingCriteriaCriterionArray []FilterFindingCriteriaCriterionInput

func (FilterFindingCriteriaCriterionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFindingCriteriaCriterion)(nil)).Elem()
}

func (i FilterFindingCriteriaCriterionArray) ToFilterFindingCriteriaCriterionArrayOutput() FilterFindingCriteriaCriterionArrayOutput {
	return i.ToFilterFindingCriteriaCriterionArrayOutputWithContext(context.Background())
}

func (i FilterFindingCriteriaCriterionArray) ToFilterFindingCriteriaCriterionArrayOutputWithContext(ctx context.Context) FilterFindingCriteriaCriterionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterFindingCriteriaCriterionArrayOutput)
}

type FilterFindingCriteriaCriterionOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaCriterionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterFindingCriteriaCriterion)(nil)).Elem()
}

func (o FilterFindingCriteriaCriterionOutput) ToFilterFindingCriteriaCriterionOutput() FilterFindingCriteriaCriterionOutput {
	return o
}

func (o FilterFindingCriteriaCriterionOutput) ToFilterFindingCriteriaCriterionOutputWithContext(ctx context.Context) FilterFindingCriteriaCriterionOutput {
	return o
}

// List of string values to be evaluated.
func (o FilterFindingCriteriaCriterionOutput) Equals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) []string { return v.Equals }).(pulumi.StringArrayOutput)
}

// The name of the field to be evaluated. The full list of field names can be found in [AWS documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_filter-findings.html#filter_criteria).
func (o FilterFindingCriteriaCriterionOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) string { return v.Field }).(pulumi.StringOutput)
}

// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
func (o FilterFindingCriteriaCriterionOutput) GreaterThan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) *string { return v.GreaterThan }).(pulumi.StringPtrOutput)
}

// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
func (o FilterFindingCriteriaCriterionOutput) GreaterThanOrEqual() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) *string { return v.GreaterThanOrEqual }).(pulumi.StringPtrOutput)
}

// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
func (o FilterFindingCriteriaCriterionOutput) LessThan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) *string { return v.LessThan }).(pulumi.StringPtrOutput)
}

// A value to be evaluated. Accepts either an integer or a date in [RFC 3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
func (o FilterFindingCriteriaCriterionOutput) LessThanOrEqual() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) *string { return v.LessThanOrEqual }).(pulumi.StringPtrOutput)
}

// List of string values to be evaluated.
func (o FilterFindingCriteriaCriterionOutput) NotEquals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FilterFindingCriteriaCriterion) []string { return v.NotEquals }).(pulumi.StringArrayOutput)
}

type FilterFindingCriteriaCriterionArrayOutput struct{ *pulumi.OutputState }

func (FilterFindingCriteriaCriterionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilterFindingCriteriaCriterion)(nil)).Elem()
}

func (o FilterFindingCriteriaCriterionArrayOutput) ToFilterFindingCriteriaCriterionArrayOutput() FilterFindingCriteriaCriterionArrayOutput {
	return o
}

func (o FilterFindingCriteriaCriterionArrayOutput) ToFilterFindingCriteriaCriterionArrayOutputWithContext(ctx context.Context) FilterFindingCriteriaCriterionArrayOutput {
	return o
}

func (o FilterFindingCriteriaCriterionArrayOutput) Index(i pulumi.IntInput) FilterFindingCriteriaCriterionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilterFindingCriteriaCriterion {
		return vs[0].([]FilterFindingCriteriaCriterion)[vs[1].(int)]
	}).(FilterFindingCriteriaCriterionOutput)
}

func init() {
	pulumi.RegisterOutputType(FilterFindingCriteriaOutput{})
	pulumi.RegisterOutputType(FilterFindingCriteriaPtrOutput{})
	pulumi.RegisterOutputType(FilterFindingCriteriaCriterionOutput{})
	pulumi.RegisterOutputType(FilterFindingCriteriaCriterionArrayOutput{})
}
