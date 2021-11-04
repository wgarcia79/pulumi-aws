// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package worklink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// WorkLink can be imported using the ARN, e.g.
//
// ```sh
//  $ pulumi import aws:worklink/fleet:Fleet test arn:aws:worklink::123456789012:fleet/example
// ```
type Fleet struct {
	pulumi.CustomResourceState

	// The ARN of the created WorkLink Fleet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn pulumi.StringPtrOutput `pulumi:"auditStreamArn"`
	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode pulumi.StringOutput `pulumi:"companyCode"`
	// The time that the fleet was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringPtrOutput `pulumi:"deviceCaCertificate"`
	// The name of the fleet.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider FleetIdentityProviderPtrOutput `pulumi:"identityProvider"`
	// The time that the fleet was last updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// A region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network FleetNetworkPtrOutput `pulumi:"network"`
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolPtrOutput `pulumi:"optimizeForEndUserLocation"`
}

// NewFleet registers a new resource with the given unique name, arguments, and options.
func NewFleet(ctx *pulumi.Context,
	name string, args *FleetArgs, opts ...pulumi.ResourceOption) (*Fleet, error) {
	if args == nil {
		args = &FleetArgs{}
	}

	var resource Fleet
	err := ctx.RegisterResource("aws:worklink/fleet:Fleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleet gets an existing Fleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetState, opts ...pulumi.ResourceOption) (*Fleet, error) {
	var resource Fleet
	err := ctx.ReadResource("aws:worklink/fleet:Fleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fleet resources.
type fleetState struct {
	// The ARN of the created WorkLink Fleet.
	Arn *string `pulumi:"arn"`
	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn *string `pulumi:"auditStreamArn"`
	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode *string `pulumi:"companyCode"`
	// The time that the fleet was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate *string `pulumi:"deviceCaCertificate"`
	// The name of the fleet.
	DisplayName *string `pulumi:"displayName"`
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider *FleetIdentityProvider `pulumi:"identityProvider"`
	// The time that the fleet was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network *FleetNetwork `pulumi:"network"`
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation *bool `pulumi:"optimizeForEndUserLocation"`
}

type FleetState struct {
	// The ARN of the created WorkLink Fleet.
	Arn pulumi.StringPtrInput
	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn pulumi.StringPtrInput
	// The identifier used by users to sign in to the Amazon WorkLink app.
	CompanyCode pulumi.StringPtrInput
	// The time that the fleet was created.
	CreatedTime pulumi.StringPtrInput
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringPtrInput
	// The name of the fleet.
	DisplayName pulumi.StringPtrInput
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider FleetIdentityProviderPtrInput
	// The time that the fleet was last updated.
	LastUpdatedTime pulumi.StringPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network FleetNetworkPtrInput
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolPtrInput
}

func (FleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetState)(nil)).Elem()
}

type fleetArgs struct {
	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn *string `pulumi:"auditStreamArn"`
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate *string `pulumi:"deviceCaCertificate"`
	// The name of the fleet.
	DisplayName *string `pulumi:"displayName"`
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider *FleetIdentityProvider `pulumi:"identityProvider"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network *FleetNetwork `pulumi:"network"`
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation *bool `pulumi:"optimizeForEndUserLocation"`
}

// The set of arguments for constructing a Fleet resource.
type FleetArgs struct {
	// The ARN of the Amazon Kinesis data stream that receives the audit events. Kinesis data stream name must begin with `"AmazonWorkLink-"`.
	AuditStreamArn pulumi.StringPtrInput
	// The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
	DeviceCaCertificate pulumi.StringPtrInput
	// The name of the fleet.
	DisplayName pulumi.StringPtrInput
	// Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
	IdentityProvider FleetIdentityProviderPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// Provide this to allow manage the company network configuration for the fleet. Fields documented below.
	Network FleetNetworkPtrInput
	// The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
	OptimizeForEndUserLocation pulumi.BoolPtrInput
}

func (FleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetArgs)(nil)).Elem()
}

type FleetInput interface {
	pulumi.Input

	ToFleetOutput() FleetOutput
	ToFleetOutputWithContext(ctx context.Context) FleetOutput
}

func (*Fleet) ElementType() reflect.Type {
	return reflect.TypeOf((*Fleet)(nil))
}

func (i *Fleet) ToFleetOutput() FleetOutput {
	return i.ToFleetOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetOutput)
}

func (i *Fleet) ToFleetPtrOutput() FleetPtrOutput {
	return i.ToFleetPtrOutputWithContext(context.Background())
}

func (i *Fleet) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetPtrOutput)
}

type FleetPtrInput interface {
	pulumi.Input

	ToFleetPtrOutput() FleetPtrOutput
	ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput
}

type fleetPtrType FleetArgs

func (*fleetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil))
}

func (i *fleetPtrType) ToFleetPtrOutput() FleetPtrOutput {
	return i.ToFleetPtrOutputWithContext(context.Background())
}

func (i *fleetPtrType) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetPtrOutput)
}

// FleetArrayInput is an input type that accepts FleetArray and FleetArrayOutput values.
// You can construct a concrete instance of `FleetArrayInput` via:
//
//          FleetArray{ FleetArgs{...} }
type FleetArrayInput interface {
	pulumi.Input

	ToFleetArrayOutput() FleetArrayOutput
	ToFleetArrayOutputWithContext(context.Context) FleetArrayOutput
}

type FleetArray []FleetInput

func (FleetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fleet)(nil)).Elem()
}

func (i FleetArray) ToFleetArrayOutput() FleetArrayOutput {
	return i.ToFleetArrayOutputWithContext(context.Background())
}

func (i FleetArray) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetArrayOutput)
}

// FleetMapInput is an input type that accepts FleetMap and FleetMapOutput values.
// You can construct a concrete instance of `FleetMapInput` via:
//
//          FleetMap{ "key": FleetArgs{...} }
type FleetMapInput interface {
	pulumi.Input

	ToFleetMapOutput() FleetMapOutput
	ToFleetMapOutputWithContext(context.Context) FleetMapOutput
}

type FleetMap map[string]FleetInput

func (FleetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fleet)(nil)).Elem()
}

func (i FleetMap) ToFleetMapOutput() FleetMapOutput {
	return i.ToFleetMapOutputWithContext(context.Background())
}

func (i FleetMap) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetMapOutput)
}

type FleetOutput struct{ *pulumi.OutputState }

func (FleetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Fleet)(nil))
}

func (o FleetOutput) ToFleetOutput() FleetOutput {
	return o
}

func (o FleetOutput) ToFleetOutputWithContext(ctx context.Context) FleetOutput {
	return o
}

func (o FleetOutput) ToFleetPtrOutput() FleetPtrOutput {
	return o.ToFleetPtrOutputWithContext(context.Background())
}

func (o FleetOutput) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Fleet) *Fleet {
		return &v
	}).(FleetPtrOutput)
}

type FleetPtrOutput struct{ *pulumi.OutputState }

func (FleetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fleet)(nil))
}

func (o FleetPtrOutput) ToFleetPtrOutput() FleetPtrOutput {
	return o
}

func (o FleetPtrOutput) ToFleetPtrOutputWithContext(ctx context.Context) FleetPtrOutput {
	return o
}

func (o FleetPtrOutput) Elem() FleetOutput {
	return o.ApplyT(func(v *Fleet) Fleet {
		if v != nil {
			return *v
		}
		var ret Fleet
		return ret
	}).(FleetOutput)
}

type FleetArrayOutput struct{ *pulumi.OutputState }

func (FleetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Fleet)(nil))
}

func (o FleetArrayOutput) ToFleetArrayOutput() FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) ToFleetArrayOutputWithContext(ctx context.Context) FleetArrayOutput {
	return o
}

func (o FleetArrayOutput) Index(i pulumi.IntInput) FleetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Fleet {
		return vs[0].([]Fleet)[vs[1].(int)]
	}).(FleetOutput)
}

type FleetMapOutput struct{ *pulumi.OutputState }

func (FleetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Fleet)(nil))
}

func (o FleetMapOutput) ToFleetMapOutput() FleetMapOutput {
	return o
}

func (o FleetMapOutput) ToFleetMapOutputWithContext(ctx context.Context) FleetMapOutput {
	return o
}

func (o FleetMapOutput) MapIndex(k pulumi.StringInput) FleetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Fleet {
		return vs[0].(map[string]Fleet)[vs[1].(string)]
	}).(FleetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FleetInput)(nil)).Elem(), &Fleet{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetPtrInput)(nil)).Elem(), &Fleet{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetArrayInput)(nil)).Elem(), FleetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetMapInput)(nil)).Elem(), FleetMap{})
	pulumi.RegisterOutputType(FleetOutput{})
	pulumi.RegisterOutputType(FleetPtrOutput{})
	pulumi.RegisterOutputType(FleetArrayOutput{})
	pulumi.RegisterOutputType(FleetMapOutput{})
}
