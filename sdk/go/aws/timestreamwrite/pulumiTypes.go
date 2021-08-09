// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package timestreamwrite

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TableRetentionProperties struct {
	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	MagneticStoreRetentionPeriodInDays int `pulumi:"magneticStoreRetentionPeriodInDays"`
	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	MemoryStoreRetentionPeriodInHours int `pulumi:"memoryStoreRetentionPeriodInHours"`
}

// TableRetentionPropertiesInput is an input type that accepts TableRetentionPropertiesArgs and TableRetentionPropertiesOutput values.
// You can construct a concrete instance of `TableRetentionPropertiesInput` via:
//
//          TableRetentionPropertiesArgs{...}
type TableRetentionPropertiesInput interface {
	pulumi.Input

	ToTableRetentionPropertiesOutput() TableRetentionPropertiesOutput
	ToTableRetentionPropertiesOutputWithContext(context.Context) TableRetentionPropertiesOutput
}

type TableRetentionPropertiesArgs struct {
	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	MagneticStoreRetentionPeriodInDays pulumi.IntInput `pulumi:"magneticStoreRetentionPeriodInDays"`
	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	MemoryStoreRetentionPeriodInHours pulumi.IntInput `pulumi:"memoryStoreRetentionPeriodInHours"`
}

func (TableRetentionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableRetentionProperties)(nil)).Elem()
}

func (i TableRetentionPropertiesArgs) ToTableRetentionPropertiesOutput() TableRetentionPropertiesOutput {
	return i.ToTableRetentionPropertiesOutputWithContext(context.Background())
}

func (i TableRetentionPropertiesArgs) ToTableRetentionPropertiesOutputWithContext(ctx context.Context) TableRetentionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableRetentionPropertiesOutput)
}

func (i TableRetentionPropertiesArgs) ToTableRetentionPropertiesPtrOutput() TableRetentionPropertiesPtrOutput {
	return i.ToTableRetentionPropertiesPtrOutputWithContext(context.Background())
}

func (i TableRetentionPropertiesArgs) ToTableRetentionPropertiesPtrOutputWithContext(ctx context.Context) TableRetentionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableRetentionPropertiesOutput).ToTableRetentionPropertiesPtrOutputWithContext(ctx)
}

// TableRetentionPropertiesPtrInput is an input type that accepts TableRetentionPropertiesArgs, TableRetentionPropertiesPtr and TableRetentionPropertiesPtrOutput values.
// You can construct a concrete instance of `TableRetentionPropertiesPtrInput` via:
//
//          TableRetentionPropertiesArgs{...}
//
//  or:
//
//          nil
type TableRetentionPropertiesPtrInput interface {
	pulumi.Input

	ToTableRetentionPropertiesPtrOutput() TableRetentionPropertiesPtrOutput
	ToTableRetentionPropertiesPtrOutputWithContext(context.Context) TableRetentionPropertiesPtrOutput
}

type tableRetentionPropertiesPtrType TableRetentionPropertiesArgs

func TableRetentionPropertiesPtr(v *TableRetentionPropertiesArgs) TableRetentionPropertiesPtrInput {
	return (*tableRetentionPropertiesPtrType)(v)
}

func (*tableRetentionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableRetentionProperties)(nil)).Elem()
}

func (i *tableRetentionPropertiesPtrType) ToTableRetentionPropertiesPtrOutput() TableRetentionPropertiesPtrOutput {
	return i.ToTableRetentionPropertiesPtrOutputWithContext(context.Background())
}

func (i *tableRetentionPropertiesPtrType) ToTableRetentionPropertiesPtrOutputWithContext(ctx context.Context) TableRetentionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableRetentionPropertiesPtrOutput)
}

type TableRetentionPropertiesOutput struct{ *pulumi.OutputState }

func (TableRetentionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableRetentionProperties)(nil)).Elem()
}

func (o TableRetentionPropertiesOutput) ToTableRetentionPropertiesOutput() TableRetentionPropertiesOutput {
	return o
}

func (o TableRetentionPropertiesOutput) ToTableRetentionPropertiesOutputWithContext(ctx context.Context) TableRetentionPropertiesOutput {
	return o
}

func (o TableRetentionPropertiesOutput) ToTableRetentionPropertiesPtrOutput() TableRetentionPropertiesPtrOutput {
	return o.ToTableRetentionPropertiesPtrOutputWithContext(context.Background())
}

func (o TableRetentionPropertiesOutput) ToTableRetentionPropertiesPtrOutputWithContext(ctx context.Context) TableRetentionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableRetentionProperties) *TableRetentionProperties {
		return &v
	}).(TableRetentionPropertiesPtrOutput)
}

// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
func (o TableRetentionPropertiesOutput) MagneticStoreRetentionPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v TableRetentionProperties) int { return v.MagneticStoreRetentionPeriodInDays }).(pulumi.IntOutput)
}

// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
func (o TableRetentionPropertiesOutput) MemoryStoreRetentionPeriodInHours() pulumi.IntOutput {
	return o.ApplyT(func(v TableRetentionProperties) int { return v.MemoryStoreRetentionPeriodInHours }).(pulumi.IntOutput)
}

type TableRetentionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TableRetentionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableRetentionProperties)(nil)).Elem()
}

func (o TableRetentionPropertiesPtrOutput) ToTableRetentionPropertiesPtrOutput() TableRetentionPropertiesPtrOutput {
	return o
}

func (o TableRetentionPropertiesPtrOutput) ToTableRetentionPropertiesPtrOutputWithContext(ctx context.Context) TableRetentionPropertiesPtrOutput {
	return o
}

func (o TableRetentionPropertiesPtrOutput) Elem() TableRetentionPropertiesOutput {
	return o.ApplyT(func(v *TableRetentionProperties) TableRetentionProperties {
		if v != nil {
			return *v
		}
		var ret TableRetentionProperties
		return ret
	}).(TableRetentionPropertiesOutput)
}

// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
func (o TableRetentionPropertiesPtrOutput) MagneticStoreRetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableRetentionProperties) *int {
		if v == nil {
			return nil
		}
		return &v.MagneticStoreRetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
func (o TableRetentionPropertiesPtrOutput) MemoryStoreRetentionPeriodInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableRetentionProperties) *int {
		if v == nil {
			return nil
		}
		return &v.MemoryStoreRetentionPeriodInHours
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TableRetentionPropertiesOutput{})
	pulumi.RegisterOutputType(TableRetentionPropertiesPtrOutput{})
}
