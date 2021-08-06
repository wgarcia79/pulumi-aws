// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewCapacityReservation(ctx, "_default", &ec2.CapacityReservationArgs{
// 			AvailabilityZone: pulumi.String("eu-west-1a"),
// 			InstanceCount:    pulumi.Int(1),
// 			InstancePlatform: pulumi.String("Linux/UNIX"),
// 			InstanceType:     pulumi.String("t2.micro"),
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
// Capacity Reservations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/capacityReservation:CapacityReservation web cr-0123456789abcdef0
// ```
type CapacityReservation struct {
	pulumi.CustomResourceState

	// The ARN of the Capacity Reservation.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized pulumi.BoolPtrOutput `pulumi:"ebsOptimized"`
	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType pulumi.StringPtrOutput `pulumi:"endDateType"`
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage pulumi.BoolPtrOutput `pulumi:"ephemeralStorage"`
	// The number of instances for which to reserve capacity.
	InstanceCount pulumi.IntOutput `pulumi:"instanceCount"`
	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria pulumi.StringPtrOutput `pulumi:"instanceMatchCriteria"`
	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform pulumi.StringOutput `pulumi:"instancePlatform"`
	// The instance type for which to reserve capacity.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn pulumi.StringPtrOutput `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the Capacity Reservation.
	OwnerId pulumi.StringOutput    `pulumi:"ownerId"`
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy pulumi.StringPtrOutput `pulumi:"tenancy"`
}

// NewCapacityReservation registers a new resource with the given unique name, arguments, and options.
func NewCapacityReservation(ctx *pulumi.Context,
	name string, args *CapacityReservationArgs, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.InstanceCount == nil {
		return nil, errors.New("invalid value for required argument 'InstanceCount'")
	}
	if args.InstancePlatform == nil {
		return nil, errors.New("invalid value for required argument 'InstancePlatform'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource CapacityReservation
	err := ctx.RegisterResource("aws:ec2/capacityReservation:CapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityReservation gets an existing CapacityReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationState, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	var resource CapacityReservation
	err := ctx.ReadResource("aws:ec2/capacityReservation:CapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityReservation resources.
type capacityReservationState struct {
	// The ARN of the Capacity Reservation.
	Arn *string `pulumi:"arn"`
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate *string `pulumi:"endDate"`
	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType *string `pulumi:"endDateType"`
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage *bool `pulumi:"ephemeralStorage"`
	// The number of instances for which to reserve capacity.
	InstanceCount *int `pulumi:"instanceCount"`
	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria *string `pulumi:"instanceMatchCriteria"`
	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform *string `pulumi:"instancePlatform"`
	// The instance type for which to reserve capacity.
	InstanceType *string `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn *string `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the Capacity Reservation.
	OwnerId *string           `pulumi:"ownerId"`
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy *string `pulumi:"tenancy"`
}

type CapacityReservationState struct {
	// The ARN of the Capacity Reservation.
	Arn pulumi.StringPtrInput
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone pulumi.StringPtrInput
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized pulumi.BoolPtrInput
	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate pulumi.StringPtrInput
	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType pulumi.StringPtrInput
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage pulumi.BoolPtrInput
	// The number of instances for which to reserve capacity.
	InstanceCount pulumi.IntPtrInput
	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria pulumi.StringPtrInput
	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform pulumi.StringPtrInput
	// The instance type for which to reserve capacity.
	InstanceType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn pulumi.StringPtrInput
	// The ID of the AWS account that owns the Capacity Reservation.
	OwnerId pulumi.StringPtrInput
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy pulumi.StringPtrInput
}

func (CapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationState)(nil)).Elem()
}

type capacityReservationArgs struct {
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate *string `pulumi:"endDate"`
	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType *string `pulumi:"endDateType"`
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage *bool `pulumi:"ephemeralStorage"`
	// The number of instances for which to reserve capacity.
	InstanceCount int `pulumi:"instanceCount"`
	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria *string `pulumi:"instanceMatchCriteria"`
	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform string `pulumi:"instancePlatform"`
	// The instance type for which to reserve capacity.
	InstanceType string `pulumi:"instanceType"`
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn *string           `pulumi:"outpostArn"`
	Tags       map[string]string `pulumi:"tags"`
	TagsAll    map[string]string `pulumi:"tagsAll"`
	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy *string `pulumi:"tenancy"`
}

// The set of arguments for constructing a CapacityReservation resource.
type CapacityReservationArgs struct {
	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone pulumi.StringInput
	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized pulumi.BoolPtrInput
	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate pulumi.StringPtrInput
	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType pulumi.StringPtrInput
	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage pulumi.BoolPtrInput
	// The number of instances for which to reserve capacity.
	InstanceCount pulumi.IntInput
	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria pulumi.StringPtrInput
	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform pulumi.StringInput
	// The instance type for which to reserve capacity.
	InstanceType pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn pulumi.StringPtrInput
	Tags       pulumi.StringMapInput
	TagsAll    pulumi.StringMapInput
	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy pulumi.StringPtrInput
}

func (CapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationArgs)(nil)).Elem()
}

type CapacityReservationInput interface {
	pulumi.Input

	ToCapacityReservationOutput() CapacityReservationOutput
	ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput
}

func (*CapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservation)(nil))
}

func (i *CapacityReservation) ToCapacityReservationOutput() CapacityReservationOutput {
	return i.ToCapacityReservationOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationOutput)
}

func (i *CapacityReservation) ToCapacityReservationPtrOutput() CapacityReservationPtrOutput {
	return i.ToCapacityReservationPtrOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationPtrOutputWithContext(ctx context.Context) CapacityReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationPtrOutput)
}

type CapacityReservationPtrInput interface {
	pulumi.Input

	ToCapacityReservationPtrOutput() CapacityReservationPtrOutput
	ToCapacityReservationPtrOutputWithContext(ctx context.Context) CapacityReservationPtrOutput
}

type capacityReservationPtrType CapacityReservationArgs

func (*capacityReservationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil))
}

func (i *capacityReservationPtrType) ToCapacityReservationPtrOutput() CapacityReservationPtrOutput {
	return i.ToCapacityReservationPtrOutputWithContext(context.Background())
}

func (i *capacityReservationPtrType) ToCapacityReservationPtrOutputWithContext(ctx context.Context) CapacityReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationPtrOutput)
}

// CapacityReservationArrayInput is an input type that accepts CapacityReservationArray and CapacityReservationArrayOutput values.
// You can construct a concrete instance of `CapacityReservationArrayInput` via:
//
//          CapacityReservationArray{ CapacityReservationArgs{...} }
type CapacityReservationArrayInput interface {
	pulumi.Input

	ToCapacityReservationArrayOutput() CapacityReservationArrayOutput
	ToCapacityReservationArrayOutputWithContext(context.Context) CapacityReservationArrayOutput
}

type CapacityReservationArray []CapacityReservationInput

func (CapacityReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CapacityReservation)(nil)).Elem()
}

func (i CapacityReservationArray) ToCapacityReservationArrayOutput() CapacityReservationArrayOutput {
	return i.ToCapacityReservationArrayOutputWithContext(context.Background())
}

func (i CapacityReservationArray) ToCapacityReservationArrayOutputWithContext(ctx context.Context) CapacityReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationArrayOutput)
}

// CapacityReservationMapInput is an input type that accepts CapacityReservationMap and CapacityReservationMapOutput values.
// You can construct a concrete instance of `CapacityReservationMapInput` via:
//
//          CapacityReservationMap{ "key": CapacityReservationArgs{...} }
type CapacityReservationMapInput interface {
	pulumi.Input

	ToCapacityReservationMapOutput() CapacityReservationMapOutput
	ToCapacityReservationMapOutputWithContext(context.Context) CapacityReservationMapOutput
}

type CapacityReservationMap map[string]CapacityReservationInput

func (CapacityReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CapacityReservation)(nil)).Elem()
}

func (i CapacityReservationMap) ToCapacityReservationMapOutput() CapacityReservationMapOutput {
	return i.ToCapacityReservationMapOutputWithContext(context.Background())
}

func (i CapacityReservationMap) ToCapacityReservationMapOutputWithContext(ctx context.Context) CapacityReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationMapOutput)
}

type CapacityReservationOutput struct {
	*pulumi.OutputState
}

func (CapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservation)(nil))
}

func (o CapacityReservationOutput) ToCapacityReservationOutput() CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationPtrOutput() CapacityReservationPtrOutput {
	return o.ToCapacityReservationPtrOutputWithContext(context.Background())
}

func (o CapacityReservationOutput) ToCapacityReservationPtrOutputWithContext(ctx context.Context) CapacityReservationPtrOutput {
	return o.ApplyT(func(v CapacityReservation) *CapacityReservation {
		return &v
	}).(CapacityReservationPtrOutput)
}

type CapacityReservationPtrOutput struct {
	*pulumi.OutputState
}

func (CapacityReservationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil))
}

func (o CapacityReservationPtrOutput) ToCapacityReservationPtrOutput() CapacityReservationPtrOutput {
	return o
}

func (o CapacityReservationPtrOutput) ToCapacityReservationPtrOutputWithContext(ctx context.Context) CapacityReservationPtrOutput {
	return o
}

type CapacityReservationArrayOutput struct{ *pulumi.OutputState }

func (CapacityReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapacityReservation)(nil))
}

func (o CapacityReservationArrayOutput) ToCapacityReservationArrayOutput() CapacityReservationArrayOutput {
	return o
}

func (o CapacityReservationArrayOutput) ToCapacityReservationArrayOutputWithContext(ctx context.Context) CapacityReservationArrayOutput {
	return o
}

func (o CapacityReservationArrayOutput) Index(i pulumi.IntInput) CapacityReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapacityReservation {
		return vs[0].([]CapacityReservation)[vs[1].(int)]
	}).(CapacityReservationOutput)
}

type CapacityReservationMapOutput struct{ *pulumi.OutputState }

func (CapacityReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CapacityReservation)(nil))
}

func (o CapacityReservationMapOutput) ToCapacityReservationMapOutput() CapacityReservationMapOutput {
	return o
}

func (o CapacityReservationMapOutput) ToCapacityReservationMapOutputWithContext(ctx context.Context) CapacityReservationMapOutput {
	return o
}

func (o CapacityReservationMapOutput) MapIndex(k pulumi.StringInput) CapacityReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CapacityReservation {
		return vs[0].(map[string]CapacityReservation)[vs[1].(string)]
	}).(CapacityReservationOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationOutput{})
	pulumi.RegisterOutputType(CapacityReservationPtrOutput{})
	pulumi.RegisterOutputType(CapacityReservationArrayOutput{})
	pulumi.RegisterOutputType(CapacityReservationMapOutput{})
}
