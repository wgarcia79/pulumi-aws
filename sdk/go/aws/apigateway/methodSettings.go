// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Method Settings, e.g. logging or monitoring.
//
// ## Import
//
// `aws_api_gateway_method_settings` can be imported using `REST-API-ID/STAGE-NAME/METHOD-PATH`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/methodSettings:MethodSettings example 12345abcde/example/test/GET
// ```
type MethodSettings struct {
	pulumi.CustomResourceState

	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringOutput `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The settings block, see below.
	Settings MethodSettingsSettingsOutput `pulumi:"settings"`
	// The name of the stage
	StageName pulumi.StringOutput `pulumi:"stageName"`
}

// NewMethodSettings registers a new resource with the given unique name, arguments, and options.
func NewMethodSettings(ctx *pulumi.Context,
	name string, args *MethodSettingsArgs, opts ...pulumi.ResourceOption) (*MethodSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MethodPath == nil {
		return nil, errors.New("invalid value for required argument 'MethodPath'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	if args.StageName == nil {
		return nil, errors.New("invalid value for required argument 'StageName'")
	}
	var resource MethodSettings
	err := ctx.RegisterResource("aws:apigateway/methodSettings:MethodSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethodSettings gets an existing MethodSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodSettingsState, opts ...pulumi.ResourceOption) (*MethodSettings, error) {
	var resource MethodSettings
	err := ctx.ReadResource("aws:apigateway/methodSettings:MethodSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MethodSettings resources.
type methodSettingsState struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath *string `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi *string `pulumi:"restApi"`
	// The settings block, see below.
	Settings *MethodSettingsSettings `pulumi:"settings"`
	// The name of the stage
	StageName *string `pulumi:"stageName"`
}

type MethodSettingsState struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringPtrInput
	// The ID of the REST API
	RestApi pulumi.StringPtrInput
	// The settings block, see below.
	Settings MethodSettingsSettingsPtrInput
	// The name of the stage
	StageName pulumi.StringPtrInput
}

func (MethodSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*methodSettingsState)(nil)).Elem()
}

type methodSettingsArgs struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath string `pulumi:"methodPath"`
	// The ID of the REST API
	RestApi interface{} `pulumi:"restApi"`
	// The settings block, see below.
	Settings MethodSettingsSettings `pulumi:"settings"`
	// The name of the stage
	StageName string `pulumi:"stageName"`
}

// The set of arguments for constructing a MethodSettings resource.
type MethodSettingsArgs struct {
	// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
	MethodPath pulumi.StringInput
	// The ID of the REST API
	RestApi pulumi.Input
	// The settings block, see below.
	Settings MethodSettingsSettingsInput
	// The name of the stage
	StageName pulumi.StringInput
}

func (MethodSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*methodSettingsArgs)(nil)).Elem()
}

type MethodSettingsInput interface {
	pulumi.Input

	ToMethodSettingsOutput() MethodSettingsOutput
	ToMethodSettingsOutputWithContext(ctx context.Context) MethodSettingsOutput
}

func (*MethodSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*MethodSettings)(nil))
}

func (i *MethodSettings) ToMethodSettingsOutput() MethodSettingsOutput {
	return i.ToMethodSettingsOutputWithContext(context.Background())
}

func (i *MethodSettings) ToMethodSettingsOutputWithContext(ctx context.Context) MethodSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodSettingsOutput)
}

func (i *MethodSettings) ToMethodSettingsPtrOutput() MethodSettingsPtrOutput {
	return i.ToMethodSettingsPtrOutputWithContext(context.Background())
}

func (i *MethodSettings) ToMethodSettingsPtrOutputWithContext(ctx context.Context) MethodSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodSettingsPtrOutput)
}

type MethodSettingsPtrInput interface {
	pulumi.Input

	ToMethodSettingsPtrOutput() MethodSettingsPtrOutput
	ToMethodSettingsPtrOutputWithContext(ctx context.Context) MethodSettingsPtrOutput
}

type methodSettingsPtrType MethodSettingsArgs

func (*methodSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MethodSettings)(nil))
}

func (i *methodSettingsPtrType) ToMethodSettingsPtrOutput() MethodSettingsPtrOutput {
	return i.ToMethodSettingsPtrOutputWithContext(context.Background())
}

func (i *methodSettingsPtrType) ToMethodSettingsPtrOutputWithContext(ctx context.Context) MethodSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodSettingsPtrOutput)
}

// MethodSettingsArrayInput is an input type that accepts MethodSettingsArray and MethodSettingsArrayOutput values.
// You can construct a concrete instance of `MethodSettingsArrayInput` via:
//
//          MethodSettingsArray{ MethodSettingsArgs{...} }
type MethodSettingsArrayInput interface {
	pulumi.Input

	ToMethodSettingsArrayOutput() MethodSettingsArrayOutput
	ToMethodSettingsArrayOutputWithContext(context.Context) MethodSettingsArrayOutput
}

type MethodSettingsArray []MethodSettingsInput

func (MethodSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*MethodSettings)(nil))
}

func (i MethodSettingsArray) ToMethodSettingsArrayOutput() MethodSettingsArrayOutput {
	return i.ToMethodSettingsArrayOutputWithContext(context.Background())
}

func (i MethodSettingsArray) ToMethodSettingsArrayOutputWithContext(ctx context.Context) MethodSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodSettingsArrayOutput)
}

// MethodSettingsMapInput is an input type that accepts MethodSettingsMap and MethodSettingsMapOutput values.
// You can construct a concrete instance of `MethodSettingsMapInput` via:
//
//          MethodSettingsMap{ "key": MethodSettingsArgs{...} }
type MethodSettingsMapInput interface {
	pulumi.Input

	ToMethodSettingsMapOutput() MethodSettingsMapOutput
	ToMethodSettingsMapOutputWithContext(context.Context) MethodSettingsMapOutput
}

type MethodSettingsMap map[string]MethodSettingsInput

func (MethodSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*MethodSettings)(nil))
}

func (i MethodSettingsMap) ToMethodSettingsMapOutput() MethodSettingsMapOutput {
	return i.ToMethodSettingsMapOutputWithContext(context.Background())
}

func (i MethodSettingsMap) ToMethodSettingsMapOutputWithContext(ctx context.Context) MethodSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MethodSettingsMapOutput)
}

type MethodSettingsOutput struct {
	*pulumi.OutputState
}

func (MethodSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MethodSettings)(nil))
}

func (o MethodSettingsOutput) ToMethodSettingsOutput() MethodSettingsOutput {
	return o
}

func (o MethodSettingsOutput) ToMethodSettingsOutputWithContext(ctx context.Context) MethodSettingsOutput {
	return o
}

func (o MethodSettingsOutput) ToMethodSettingsPtrOutput() MethodSettingsPtrOutput {
	return o.ToMethodSettingsPtrOutputWithContext(context.Background())
}

func (o MethodSettingsOutput) ToMethodSettingsPtrOutputWithContext(ctx context.Context) MethodSettingsPtrOutput {
	return o.ApplyT(func(v MethodSettings) *MethodSettings {
		return &v
	}).(MethodSettingsPtrOutput)
}

type MethodSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (MethodSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MethodSettings)(nil))
}

func (o MethodSettingsPtrOutput) ToMethodSettingsPtrOutput() MethodSettingsPtrOutput {
	return o
}

func (o MethodSettingsPtrOutput) ToMethodSettingsPtrOutputWithContext(ctx context.Context) MethodSettingsPtrOutput {
	return o
}

type MethodSettingsArrayOutput struct{ *pulumi.OutputState }

func (MethodSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MethodSettings)(nil))
}

func (o MethodSettingsArrayOutput) ToMethodSettingsArrayOutput() MethodSettingsArrayOutput {
	return o
}

func (o MethodSettingsArrayOutput) ToMethodSettingsArrayOutputWithContext(ctx context.Context) MethodSettingsArrayOutput {
	return o
}

func (o MethodSettingsArrayOutput) Index(i pulumi.IntInput) MethodSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MethodSettings {
		return vs[0].([]MethodSettings)[vs[1].(int)]
	}).(MethodSettingsOutput)
}

type MethodSettingsMapOutput struct{ *pulumi.OutputState }

func (MethodSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MethodSettings)(nil))
}

func (o MethodSettingsMapOutput) ToMethodSettingsMapOutput() MethodSettingsMapOutput {
	return o
}

func (o MethodSettingsMapOutput) ToMethodSettingsMapOutputWithContext(ctx context.Context) MethodSettingsMapOutput {
	return o
}

func (o MethodSettingsMapOutput) MapIndex(k pulumi.StringInput) MethodSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MethodSettings {
		return vs[0].(map[string]MethodSettings)[vs[1].(string)]
	}).(MethodSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(MethodSettingsOutput{})
	pulumi.RegisterOutputType(MethodSettingsPtrOutput{})
	pulumi.RegisterOutputType(MethodSettingsArrayOutput{})
	pulumi.RegisterOutputType(MethodSettingsMapOutput{})
}
