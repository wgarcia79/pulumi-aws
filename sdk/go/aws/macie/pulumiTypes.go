// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FindingsFilterFindingCriteria struct {
	// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
	Criterions []FindingsFilterFindingCriteriaCriterion `pulumi:"criterions"`
}

// FindingsFilterFindingCriteriaInput is an input type that accepts FindingsFilterFindingCriteriaArgs and FindingsFilterFindingCriteriaOutput values.
// You can construct a concrete instance of `FindingsFilterFindingCriteriaInput` via:
//
//          FindingsFilterFindingCriteriaArgs{...}
type FindingsFilterFindingCriteriaInput interface {
	pulumi.Input

	ToFindingsFilterFindingCriteriaOutput() FindingsFilterFindingCriteriaOutput
	ToFindingsFilterFindingCriteriaOutputWithContext(context.Context) FindingsFilterFindingCriteriaOutput
}

type FindingsFilterFindingCriteriaArgs struct {
	// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
	Criterions FindingsFilterFindingCriteriaCriterionArrayInput `pulumi:"criterions"`
}

func (FindingsFilterFindingCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilterFindingCriteria)(nil)).Elem()
}

func (i FindingsFilterFindingCriteriaArgs) ToFindingsFilterFindingCriteriaOutput() FindingsFilterFindingCriteriaOutput {
	return i.ToFindingsFilterFindingCriteriaOutputWithContext(context.Background())
}

func (i FindingsFilterFindingCriteriaArgs) ToFindingsFilterFindingCriteriaOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterFindingCriteriaOutput)
}

func (i FindingsFilterFindingCriteriaArgs) ToFindingsFilterFindingCriteriaPtrOutput() FindingsFilterFindingCriteriaPtrOutput {
	return i.ToFindingsFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i FindingsFilterFindingCriteriaArgs) ToFindingsFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterFindingCriteriaOutput).ToFindingsFilterFindingCriteriaPtrOutputWithContext(ctx)
}

// FindingsFilterFindingCriteriaPtrInput is an input type that accepts FindingsFilterFindingCriteriaArgs, FindingsFilterFindingCriteriaPtr and FindingsFilterFindingCriteriaPtrOutput values.
// You can construct a concrete instance of `FindingsFilterFindingCriteriaPtrInput` via:
//
//          FindingsFilterFindingCriteriaArgs{...}
//
//  or:
//
//          nil
type FindingsFilterFindingCriteriaPtrInput interface {
	pulumi.Input

	ToFindingsFilterFindingCriteriaPtrOutput() FindingsFilterFindingCriteriaPtrOutput
	ToFindingsFilterFindingCriteriaPtrOutputWithContext(context.Context) FindingsFilterFindingCriteriaPtrOutput
}

type findingsFilterFindingCriteriaPtrType FindingsFilterFindingCriteriaArgs

func FindingsFilterFindingCriteriaPtr(v *FindingsFilterFindingCriteriaArgs) FindingsFilterFindingCriteriaPtrInput {
	return (*findingsFilterFindingCriteriaPtrType)(v)
}

func (*findingsFilterFindingCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FindingsFilterFindingCriteria)(nil)).Elem()
}

func (i *findingsFilterFindingCriteriaPtrType) ToFindingsFilterFindingCriteriaPtrOutput() FindingsFilterFindingCriteriaPtrOutput {
	return i.ToFindingsFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (i *findingsFilterFindingCriteriaPtrType) ToFindingsFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterFindingCriteriaPtrOutput)
}

type FindingsFilterFindingCriteriaOutput struct{ *pulumi.OutputState }

func (FindingsFilterFindingCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilterFindingCriteria)(nil)).Elem()
}

func (o FindingsFilterFindingCriteriaOutput) ToFindingsFilterFindingCriteriaOutput() FindingsFilterFindingCriteriaOutput {
	return o
}

func (o FindingsFilterFindingCriteriaOutput) ToFindingsFilterFindingCriteriaOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaOutput {
	return o
}

func (o FindingsFilterFindingCriteriaOutput) ToFindingsFilterFindingCriteriaPtrOutput() FindingsFilterFindingCriteriaPtrOutput {
	return o.ToFindingsFilterFindingCriteriaPtrOutputWithContext(context.Background())
}

func (o FindingsFilterFindingCriteriaOutput) ToFindingsFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FindingsFilterFindingCriteria) *FindingsFilterFindingCriteria {
		return &v
	}).(FindingsFilterFindingCriteriaPtrOutput)
}

// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
func (o FindingsFilterFindingCriteriaOutput) Criterions() FindingsFilterFindingCriteriaCriterionArrayOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteria) []FindingsFilterFindingCriteriaCriterion { return v.Criterions }).(FindingsFilterFindingCriteriaCriterionArrayOutput)
}

type FindingsFilterFindingCriteriaPtrOutput struct{ *pulumi.OutputState }

func (FindingsFilterFindingCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FindingsFilterFindingCriteria)(nil)).Elem()
}

func (o FindingsFilterFindingCriteriaPtrOutput) ToFindingsFilterFindingCriteriaPtrOutput() FindingsFilterFindingCriteriaPtrOutput {
	return o
}

func (o FindingsFilterFindingCriteriaPtrOutput) ToFindingsFilterFindingCriteriaPtrOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaPtrOutput {
	return o
}

func (o FindingsFilterFindingCriteriaPtrOutput) Elem() FindingsFilterFindingCriteriaOutput {
	return o.ApplyT(func(v *FindingsFilterFindingCriteria) FindingsFilterFindingCriteria {
		if v != nil {
			return *v
		}
		var ret FindingsFilterFindingCriteria
		return ret
	}).(FindingsFilterFindingCriteriaOutput)
}

// A condition that specifies the property, operator, and one or more values to use to filter the results.  (documented below)
func (o FindingsFilterFindingCriteriaPtrOutput) Criterions() FindingsFilterFindingCriteriaCriterionArrayOutput {
	return o.ApplyT(func(v *FindingsFilterFindingCriteria) []FindingsFilterFindingCriteriaCriterion {
		if v == nil {
			return nil
		}
		return v.Criterions
	}).(FindingsFilterFindingCriteriaCriterionArrayOutput)
}

type FindingsFilterFindingCriteriaCriterion struct {
	// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
	EqExactMatches []string `pulumi:"eqExactMatches"`
	// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Eqs []string `pulumi:"eqs"`
	// The name of the field to be evaluated.
	Field string `pulumi:"field"`
	// The value for the property is greater than the specified value.
	Gt *string `pulumi:"gt"`
	// The value for the property is greater than or equal to the specified value.
	Gte *string `pulumi:"gte"`
	// The value for the property is less than the specified value.
	Lt *string `pulumi:"lt"`
	// The value for the property is less than or equal to the specified value.
	Lte *string `pulumi:"lte"`
	// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Neqs []string `pulumi:"neqs"`
}

// FindingsFilterFindingCriteriaCriterionInput is an input type that accepts FindingsFilterFindingCriteriaCriterionArgs and FindingsFilterFindingCriteriaCriterionOutput values.
// You can construct a concrete instance of `FindingsFilterFindingCriteriaCriterionInput` via:
//
//          FindingsFilterFindingCriteriaCriterionArgs{...}
type FindingsFilterFindingCriteriaCriterionInput interface {
	pulumi.Input

	ToFindingsFilterFindingCriteriaCriterionOutput() FindingsFilterFindingCriteriaCriterionOutput
	ToFindingsFilterFindingCriteriaCriterionOutputWithContext(context.Context) FindingsFilterFindingCriteriaCriterionOutput
}

type FindingsFilterFindingCriteriaCriterionArgs struct {
	// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
	EqExactMatches pulumi.StringArrayInput `pulumi:"eqExactMatches"`
	// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Eqs pulumi.StringArrayInput `pulumi:"eqs"`
	// The name of the field to be evaluated.
	Field pulumi.StringInput `pulumi:"field"`
	// The value for the property is greater than the specified value.
	Gt pulumi.StringPtrInput `pulumi:"gt"`
	// The value for the property is greater than or equal to the specified value.
	Gte pulumi.StringPtrInput `pulumi:"gte"`
	// The value for the property is less than the specified value.
	Lt pulumi.StringPtrInput `pulumi:"lt"`
	// The value for the property is less than or equal to the specified value.
	Lte pulumi.StringPtrInput `pulumi:"lte"`
	// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
	Neqs pulumi.StringArrayInput `pulumi:"neqs"`
}

func (FindingsFilterFindingCriteriaCriterionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilterFindingCriteriaCriterion)(nil)).Elem()
}

func (i FindingsFilterFindingCriteriaCriterionArgs) ToFindingsFilterFindingCriteriaCriterionOutput() FindingsFilterFindingCriteriaCriterionOutput {
	return i.ToFindingsFilterFindingCriteriaCriterionOutputWithContext(context.Background())
}

func (i FindingsFilterFindingCriteriaCriterionArgs) ToFindingsFilterFindingCriteriaCriterionOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaCriterionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterFindingCriteriaCriterionOutput)
}

// FindingsFilterFindingCriteriaCriterionArrayInput is an input type that accepts FindingsFilterFindingCriteriaCriterionArray and FindingsFilterFindingCriteriaCriterionArrayOutput values.
// You can construct a concrete instance of `FindingsFilterFindingCriteriaCriterionArrayInput` via:
//
//          FindingsFilterFindingCriteriaCriterionArray{ FindingsFilterFindingCriteriaCriterionArgs{...} }
type FindingsFilterFindingCriteriaCriterionArrayInput interface {
	pulumi.Input

	ToFindingsFilterFindingCriteriaCriterionArrayOutput() FindingsFilterFindingCriteriaCriterionArrayOutput
	ToFindingsFilterFindingCriteriaCriterionArrayOutputWithContext(context.Context) FindingsFilterFindingCriteriaCriterionArrayOutput
}

type FindingsFilterFindingCriteriaCriterionArray []FindingsFilterFindingCriteriaCriterionInput

func (FindingsFilterFindingCriteriaCriterionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FindingsFilterFindingCriteriaCriterion)(nil)).Elem()
}

func (i FindingsFilterFindingCriteriaCriterionArray) ToFindingsFilterFindingCriteriaCriterionArrayOutput() FindingsFilterFindingCriteriaCriterionArrayOutput {
	return i.ToFindingsFilterFindingCriteriaCriterionArrayOutputWithContext(context.Background())
}

func (i FindingsFilterFindingCriteriaCriterionArray) ToFindingsFilterFindingCriteriaCriterionArrayOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaCriterionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterFindingCriteriaCriterionArrayOutput)
}

type FindingsFilterFindingCriteriaCriterionOutput struct{ *pulumi.OutputState }

func (FindingsFilterFindingCriteriaCriterionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilterFindingCriteriaCriterion)(nil)).Elem()
}

func (o FindingsFilterFindingCriteriaCriterionOutput) ToFindingsFilterFindingCriteriaCriterionOutput() FindingsFilterFindingCriteriaCriterionOutput {
	return o
}

func (o FindingsFilterFindingCriteriaCriterionOutput) ToFindingsFilterFindingCriteriaCriterionOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaCriterionOutput {
	return o
}

// The value for the property exclusively matches (equals an exact match for) all the specified values. If you specify multiple values, Amazon Macie uses AND logic to join the values.
func (o FindingsFilterFindingCriteriaCriterionOutput) EqExactMatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) []string { return v.EqExactMatches }).(pulumi.StringArrayOutput)
}

// The value for the property matches (equals) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
func (o FindingsFilterFindingCriteriaCriterionOutput) Eqs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) []string { return v.Eqs }).(pulumi.StringArrayOutput)
}

// The name of the field to be evaluated.
func (o FindingsFilterFindingCriteriaCriterionOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) string { return v.Field }).(pulumi.StringOutput)
}

// The value for the property is greater than the specified value.
func (o FindingsFilterFindingCriteriaCriterionOutput) Gt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) *string { return v.Gt }).(pulumi.StringPtrOutput)
}

// The value for the property is greater than or equal to the specified value.
func (o FindingsFilterFindingCriteriaCriterionOutput) Gte() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) *string { return v.Gte }).(pulumi.StringPtrOutput)
}

// The value for the property is less than the specified value.
func (o FindingsFilterFindingCriteriaCriterionOutput) Lt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) *string { return v.Lt }).(pulumi.StringPtrOutput)
}

// The value for the property is less than or equal to the specified value.
func (o FindingsFilterFindingCriteriaCriterionOutput) Lte() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) *string { return v.Lte }).(pulumi.StringPtrOutput)
}

// The value for the property doesn't match (doesn't equal) the specified value. If you specify multiple values, Amazon Macie uses OR logic to join the values.
func (o FindingsFilterFindingCriteriaCriterionOutput) Neqs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FindingsFilterFindingCriteriaCriterion) []string { return v.Neqs }).(pulumi.StringArrayOutput)
}

type FindingsFilterFindingCriteriaCriterionArrayOutput struct{ *pulumi.OutputState }

func (FindingsFilterFindingCriteriaCriterionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FindingsFilterFindingCriteriaCriterion)(nil)).Elem()
}

func (o FindingsFilterFindingCriteriaCriterionArrayOutput) ToFindingsFilterFindingCriteriaCriterionArrayOutput() FindingsFilterFindingCriteriaCriterionArrayOutput {
	return o
}

func (o FindingsFilterFindingCriteriaCriterionArrayOutput) ToFindingsFilterFindingCriteriaCriterionArrayOutputWithContext(ctx context.Context) FindingsFilterFindingCriteriaCriterionArrayOutput {
	return o
}

func (o FindingsFilterFindingCriteriaCriterionArrayOutput) Index(i pulumi.IntInput) FindingsFilterFindingCriteriaCriterionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FindingsFilterFindingCriteriaCriterion {
		return vs[0].([]FindingsFilterFindingCriteriaCriterion)[vs[1].(int)]
	}).(FindingsFilterFindingCriteriaCriterionOutput)
}

type S3BucketAssociationClassificationType struct {
	// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
	// The only valid value is the default value, `FULL`.
	Continuous *string `pulumi:"continuous"`
	// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
	// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
	OneTime *string `pulumi:"oneTime"`
}

// S3BucketAssociationClassificationTypeInput is an input type that accepts S3BucketAssociationClassificationTypeArgs and S3BucketAssociationClassificationTypeOutput values.
// You can construct a concrete instance of `S3BucketAssociationClassificationTypeInput` via:
//
//          S3BucketAssociationClassificationTypeArgs{...}
type S3BucketAssociationClassificationTypeInput interface {
	pulumi.Input

	ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput
	ToS3BucketAssociationClassificationTypeOutputWithContext(context.Context) S3BucketAssociationClassificationTypeOutput
}

type S3BucketAssociationClassificationTypeArgs struct {
	// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
	// The only valid value is the default value, `FULL`.
	Continuous pulumi.StringPtrInput `pulumi:"continuous"`
	// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
	// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
	OneTime pulumi.StringPtrInput `pulumi:"oneTime"`
}

func (S3BucketAssociationClassificationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*S3BucketAssociationClassificationType)(nil)).Elem()
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput {
	return i.ToS3BucketAssociationClassificationTypeOutputWithContext(context.Background())
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypeOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypeOutput)
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return i.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypeOutput).ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx)
}

// S3BucketAssociationClassificationTypePtrInput is an input type that accepts S3BucketAssociationClassificationTypeArgs, S3BucketAssociationClassificationTypePtr and S3BucketAssociationClassificationTypePtrOutput values.
// You can construct a concrete instance of `S3BucketAssociationClassificationTypePtrInput` via:
//
//          S3BucketAssociationClassificationTypeArgs{...}
//
//  or:
//
//          nil
type S3BucketAssociationClassificationTypePtrInput interface {
	pulumi.Input

	ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput
	ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Context) S3BucketAssociationClassificationTypePtrOutput
}

type s3bucketAssociationClassificationTypePtrType S3BucketAssociationClassificationTypeArgs

func S3BucketAssociationClassificationTypePtr(v *S3BucketAssociationClassificationTypeArgs) S3BucketAssociationClassificationTypePtrInput {
	return (*s3bucketAssociationClassificationTypePtrType)(v)
}

func (*s3bucketAssociationClassificationTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketAssociationClassificationType)(nil)).Elem()
}

func (i *s3bucketAssociationClassificationTypePtrType) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return i.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (i *s3bucketAssociationClassificationTypePtrType) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypePtrOutput)
}

type S3BucketAssociationClassificationTypeOutput struct{ *pulumi.OutputState }

func (S3BucketAssociationClassificationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*S3BucketAssociationClassificationType)(nil)).Elem()
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput {
	return o
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypeOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypeOutput {
	return o
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return o.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v S3BucketAssociationClassificationType) *S3BucketAssociationClassificationType {
		return &v
	}).(S3BucketAssociationClassificationTypePtrOutput)
}

// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
// The only valid value is the default value, `FULL`.
func (o S3BucketAssociationClassificationTypeOutput) Continuous() pulumi.StringPtrOutput {
	return o.ApplyT(func(v S3BucketAssociationClassificationType) *string { return v.Continuous }).(pulumi.StringPtrOutput)
}

// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
func (o S3BucketAssociationClassificationTypeOutput) OneTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v S3BucketAssociationClassificationType) *string { return v.OneTime }).(pulumi.StringPtrOutput)
}

type S3BucketAssociationClassificationTypePtrOutput struct{ *pulumi.OutputState }

func (S3BucketAssociationClassificationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketAssociationClassificationType)(nil)).Elem()
}

func (o S3BucketAssociationClassificationTypePtrOutput) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return o
}

func (o S3BucketAssociationClassificationTypePtrOutput) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return o
}

func (o S3BucketAssociationClassificationTypePtrOutput) Elem() S3BucketAssociationClassificationTypeOutput {
	return o.ApplyT(func(v *S3BucketAssociationClassificationType) S3BucketAssociationClassificationType {
		if v != nil {
			return *v
		}
		var ret S3BucketAssociationClassificationType
		return ret
	}).(S3BucketAssociationClassificationTypeOutput)
}

// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
// The only valid value is the default value, `FULL`.
func (o S3BucketAssociationClassificationTypePtrOutput) Continuous() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3BucketAssociationClassificationType) *string {
		if v == nil {
			return nil
		}
		return v.Continuous
	}).(pulumi.StringPtrOutput)
}

// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
func (o S3BucketAssociationClassificationTypePtrOutput) OneTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3BucketAssociationClassificationType) *string {
		if v == nil {
			return nil
		}
		return v.OneTime
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FindingsFilterFindingCriteriaOutput{})
	pulumi.RegisterOutputType(FindingsFilterFindingCriteriaPtrOutput{})
	pulumi.RegisterOutputType(FindingsFilterFindingCriteriaCriterionOutput{})
	pulumi.RegisterOutputType(FindingsFilterFindingCriteriaCriterionArrayOutput{})
	pulumi.RegisterOutputType(S3BucketAssociationClassificationTypeOutput{})
	pulumi.RegisterOutputType(S3BucketAssociationClassificationTypePtrOutput{})
}
