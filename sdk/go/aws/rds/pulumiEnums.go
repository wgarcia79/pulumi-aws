// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EngineMode pulumi.String

const (
	EngineModeProvisioned   = EngineMode("provisioned")
	EngineModeServerless    = EngineMode("serverless")
	EngineModeParallelQuery = EngineMode("parallelquery")
	EngineModeGlobal        = EngineMode("global")
)

func (EngineMode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EngineMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EngineMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EngineType pulumi.String

const (
	EngineTypeAurora           = EngineType("aurora")
	EngineTypeAuroraMysql      = EngineType("aurora-mysql")
	EngineTypeAuroraPostgresql = EngineType("aurora-postgresql")
)

func (EngineType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EngineType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EngineType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceType pulumi.String

const (
	InstanceType_T3_Micro     = InstanceType("db.t3.micro")
	InstanceType_T3_Small     = InstanceType("db.t3.small")
	InstanceType_T3_Medium    = InstanceType("db.t3.medium")
	InstanceType_T3_Large     = InstanceType("db.t3.large")
	InstanceType_T3_XLarge    = InstanceType("db.t3.xlarge")
	InstanceType_T3_2XLarge   = InstanceType("db.t3.2xlarge")
	InstanceType_T2_Micro     = InstanceType("db.t2.micro")
	InstanceType_T2_Small     = InstanceType("db.t2.small")
	InstanceType_T2_Medium    = InstanceType("db.t2.medium")
	InstanceType_T2_Large     = InstanceType("db.t2.large")
	InstanceType_T2_XLarge    = InstanceType("db.t2.xlarge")
	InstanceType_T2_2XLarge   = InstanceType("db.t2.2xlarge")
	InstanceType_M1_Small     = InstanceType("db.m1.small")
	InstanceType_M1_Medium    = InstanceType("db.m1.medium")
	InstanceType_M1_Large     = InstanceType("db.m1.large")
	InstanceType_M1_XLarge    = InstanceType("db.m1.xlarge")
	InstanceType_M2_XLarge    = InstanceType("db.m2.xlarge")
	InstanceType_M2_2XLarge   = InstanceType("db.m2.2xlarge")
	InstanceType_M2_4XLarge   = InstanceType("db.m2.4xlarge")
	InstanceType_M3_Medium    = InstanceType("db.m3.medium")
	InstanceType_M3_Large     = InstanceType("db.m3.large")
	InstanceType_M3_XLarge    = InstanceType("db.m3.xlarge")
	InstanceType_M3_2XLarge   = InstanceType("db.m3.2xlarge")
	InstanceType_M4_Large     = InstanceType("db.m4.large")
	InstanceType_M4_XLarge    = InstanceType("db.m4.xlarge")
	InstanceType_M4_2XLarge   = InstanceType("db.m4.2xlarge")
	InstanceType_M4_4XLarge   = InstanceType("db.m4.4xlarge")
	InstanceType_M4_10XLarge  = InstanceType("db.m4.10xlarge")
	InstanceType_M4_16XLarge  = InstanceType("db.m4.10xlarge")
	InstanceType_M5_Large     = InstanceType("db.m5.large")
	InstanceType_M5_XLarge    = InstanceType("db.m5.xlarge")
	InstanceType_M5_2XLarge   = InstanceType("db.m5.2xlarge")
	InstanceType_M5_4XLarge   = InstanceType("db.m5.4xlarge")
	InstanceType_M5_12XLarge  = InstanceType("db.m5.12xlarge")
	InstanceType_M5_24XLarge  = InstanceType("db.m5.24xlarge")
	InstanceType_R3_Large     = InstanceType("db.r3.large")
	InstanceType_R3_XLarge    = InstanceType("db.r3.xlarge")
	InstanceType_R3_2XLarge   = InstanceType("db.r3.2xlarge")
	InstanceType_R3_4XLarge   = InstanceType("db.r3.4xlarge")
	InstanceType_R3_8XLarge   = InstanceType("db.r3.8xlarge")
	InstanceType_R4_Large     = InstanceType("db.r4.large")
	InstanceType_R4_XLarge    = InstanceType("db.r4.xlarge")
	InstanceType_R4_2XLarge   = InstanceType("db.r4.2xlarge")
	InstanceType_R4_4XLarge   = InstanceType("db.r4.4xlarge")
	InstanceType_R4_8XLarge   = InstanceType("db.r4.8xlarge")
	InstanceType_R4_16XLarge  = InstanceType("db.r4.16xlarge")
	InstanceType_R5_Large     = InstanceType("db.r5.large")
	InstanceType_R5_XLarge    = InstanceType("db.r5.xlarge")
	InstanceType_R5_2XLarge   = InstanceType("db.r5.2xlarge")
	InstanceType_R5_4XLarge   = InstanceType("db.r5.4xlarge")
	InstanceType_R5_12XLarge  = InstanceType("db.r5.12xlarge")
	InstanceType_R5_24XLarge  = InstanceType("db.r5.24xlarge")
	InstanceType_X1_16XLarge  = InstanceType("db.x1.16xlarge")
	InstanceType_X1_32XLarge  = InstanceType("db.x1.32xlarge")
	InstanceType_X1E_XLarge   = InstanceType("db.x1e.xlarge")
	InstanceType_X1E_2XLarge  = InstanceType("db.x1e.2xlarge")
	InstanceType_X1E_4XLarge  = InstanceType("db.x1e.4xlarge")
	InstanceType_X1E_8XLarge  = InstanceType("db.x1e.8xlarge")
	InstanceType_X1E_32XLarge = InstanceType("db.x1e.32xlarge")
)

func (InstanceType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e InstanceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageType pulumi.String

const (
	StorageTypeStandard = StorageType("standard")
	StorageTypeGP2      = StorageType("gp2")
	StorageTypeIO1      = StorageType("io1")
)

func (StorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e StorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
