// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Backup Global Settings resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.NewGlobalSettings(ctx, "test", &backup.GlobalSettingsArgs{
// 			GlobalSettings: pulumi.StringMap{
// 				"isCrossAccountBackupEnabled": pulumi.String("true"),
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
// Backup Global Settings can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:backup/globalSettings:GlobalSettings example 123456789012
// ```
type GlobalSettings struct {
	pulumi.CustomResourceState

	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings pulumi.StringMapOutput `pulumi:"globalSettings"`
}

// NewGlobalSettings registers a new resource with the given unique name, arguments, and options.
func NewGlobalSettings(ctx *pulumi.Context,
	name string, args *GlobalSettingsArgs, opts ...pulumi.ResourceOption) (*GlobalSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalSettings == nil {
		return nil, errors.New("invalid value for required argument 'GlobalSettings'")
	}
	var resource GlobalSettings
	err := ctx.RegisterResource("aws:backup/globalSettings:GlobalSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalSettings gets an existing GlobalSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalSettingsState, opts ...pulumi.ResourceOption) (*GlobalSettings, error) {
	var resource GlobalSettings
	err := ctx.ReadResource("aws:backup/globalSettings:GlobalSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalSettings resources.
type globalSettingsState struct {
	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings map[string]string `pulumi:"globalSettings"`
}

type GlobalSettingsState struct {
	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings pulumi.StringMapInput
}

func (GlobalSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalSettingsState)(nil)).Elem()
}

type globalSettingsArgs struct {
	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings map[string]string `pulumi:"globalSettings"`
}

// The set of arguments for constructing a GlobalSettings resource.
type GlobalSettingsArgs struct {
	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings pulumi.StringMapInput
}

func (GlobalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalSettingsArgs)(nil)).Elem()
}

type GlobalSettingsInput interface {
	pulumi.Input

	ToGlobalSettingsOutput() GlobalSettingsOutput
	ToGlobalSettingsOutputWithContext(ctx context.Context) GlobalSettingsOutput
}

func (*GlobalSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalSettings)(nil))
}

func (i *GlobalSettings) ToGlobalSettingsOutput() GlobalSettingsOutput {
	return i.ToGlobalSettingsOutputWithContext(context.Background())
}

func (i *GlobalSettings) ToGlobalSettingsOutputWithContext(ctx context.Context) GlobalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSettingsOutput)
}

func (i *GlobalSettings) ToGlobalSettingsPtrOutput() GlobalSettingsPtrOutput {
	return i.ToGlobalSettingsPtrOutputWithContext(context.Background())
}

func (i *GlobalSettings) ToGlobalSettingsPtrOutputWithContext(ctx context.Context) GlobalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSettingsPtrOutput)
}

type GlobalSettingsPtrInput interface {
	pulumi.Input

	ToGlobalSettingsPtrOutput() GlobalSettingsPtrOutput
	ToGlobalSettingsPtrOutputWithContext(ctx context.Context) GlobalSettingsPtrOutput
}

type globalSettingsPtrType GlobalSettingsArgs

func (*globalSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalSettings)(nil))
}

func (i *globalSettingsPtrType) ToGlobalSettingsPtrOutput() GlobalSettingsPtrOutput {
	return i.ToGlobalSettingsPtrOutputWithContext(context.Background())
}

func (i *globalSettingsPtrType) ToGlobalSettingsPtrOutputWithContext(ctx context.Context) GlobalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSettingsPtrOutput)
}

// GlobalSettingsArrayInput is an input type that accepts GlobalSettingsArray and GlobalSettingsArrayOutput values.
// You can construct a concrete instance of `GlobalSettingsArrayInput` via:
//
//          GlobalSettingsArray{ GlobalSettingsArgs{...} }
type GlobalSettingsArrayInput interface {
	pulumi.Input

	ToGlobalSettingsArrayOutput() GlobalSettingsArrayOutput
	ToGlobalSettingsArrayOutputWithContext(context.Context) GlobalSettingsArrayOutput
}

type GlobalSettingsArray []GlobalSettingsInput

func (GlobalSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalSettings)(nil)).Elem()
}

func (i GlobalSettingsArray) ToGlobalSettingsArrayOutput() GlobalSettingsArrayOutput {
	return i.ToGlobalSettingsArrayOutputWithContext(context.Background())
}

func (i GlobalSettingsArray) ToGlobalSettingsArrayOutputWithContext(ctx context.Context) GlobalSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSettingsArrayOutput)
}

// GlobalSettingsMapInput is an input type that accepts GlobalSettingsMap and GlobalSettingsMapOutput values.
// You can construct a concrete instance of `GlobalSettingsMapInput` via:
//
//          GlobalSettingsMap{ "key": GlobalSettingsArgs{...} }
type GlobalSettingsMapInput interface {
	pulumi.Input

	ToGlobalSettingsMapOutput() GlobalSettingsMapOutput
	ToGlobalSettingsMapOutputWithContext(context.Context) GlobalSettingsMapOutput
}

type GlobalSettingsMap map[string]GlobalSettingsInput

func (GlobalSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalSettings)(nil)).Elem()
}

func (i GlobalSettingsMap) ToGlobalSettingsMapOutput() GlobalSettingsMapOutput {
	return i.ToGlobalSettingsMapOutputWithContext(context.Background())
}

func (i GlobalSettingsMap) ToGlobalSettingsMapOutputWithContext(ctx context.Context) GlobalSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalSettingsMapOutput)
}

type GlobalSettingsOutput struct{ *pulumi.OutputState }

func (GlobalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalSettings)(nil))
}

func (o GlobalSettingsOutput) ToGlobalSettingsOutput() GlobalSettingsOutput {
	return o
}

func (o GlobalSettingsOutput) ToGlobalSettingsOutputWithContext(ctx context.Context) GlobalSettingsOutput {
	return o
}

func (o GlobalSettingsOutput) ToGlobalSettingsPtrOutput() GlobalSettingsPtrOutput {
	return o.ToGlobalSettingsPtrOutputWithContext(context.Background())
}

func (o GlobalSettingsOutput) ToGlobalSettingsPtrOutputWithContext(ctx context.Context) GlobalSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalSettings) *GlobalSettings {
		return &v
	}).(GlobalSettingsPtrOutput)
}

type GlobalSettingsPtrOutput struct{ *pulumi.OutputState }

func (GlobalSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalSettings)(nil))
}

func (o GlobalSettingsPtrOutput) ToGlobalSettingsPtrOutput() GlobalSettingsPtrOutput {
	return o
}

func (o GlobalSettingsPtrOutput) ToGlobalSettingsPtrOutputWithContext(ctx context.Context) GlobalSettingsPtrOutput {
	return o
}

func (o GlobalSettingsPtrOutput) Elem() GlobalSettingsOutput {
	return o.ApplyT(func(v *GlobalSettings) GlobalSettings {
		if v != nil {
			return *v
		}
		var ret GlobalSettings
		return ret
	}).(GlobalSettingsOutput)
}

type GlobalSettingsArrayOutput struct{ *pulumi.OutputState }

func (GlobalSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalSettings)(nil))
}

func (o GlobalSettingsArrayOutput) ToGlobalSettingsArrayOutput() GlobalSettingsArrayOutput {
	return o
}

func (o GlobalSettingsArrayOutput) ToGlobalSettingsArrayOutputWithContext(ctx context.Context) GlobalSettingsArrayOutput {
	return o
}

func (o GlobalSettingsArrayOutput) Index(i pulumi.IntInput) GlobalSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalSettings {
		return vs[0].([]GlobalSettings)[vs[1].(int)]
	}).(GlobalSettingsOutput)
}

type GlobalSettingsMapOutput struct{ *pulumi.OutputState }

func (GlobalSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalSettings)(nil))
}

func (o GlobalSettingsMapOutput) ToGlobalSettingsMapOutput() GlobalSettingsMapOutput {
	return o
}

func (o GlobalSettingsMapOutput) ToGlobalSettingsMapOutputWithContext(ctx context.Context) GlobalSettingsMapOutput {
	return o
}

func (o GlobalSettingsMapOutput) MapIndex(k pulumi.StringInput) GlobalSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalSettings {
		return vs[0].(map[string]GlobalSettings)[vs[1].(string)]
	}).(GlobalSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalSettingsOutput{})
	pulumi.RegisterOutputType(GlobalSettingsPtrOutput{})
	pulumi.RegisterOutputType(GlobalSettingsArrayOutput{})
	pulumi.RegisterOutputType(GlobalSettingsMapOutput{})
}
