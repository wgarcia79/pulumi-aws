// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway Method Settings, e.g. logging or monitoring.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method_settings.html.markdown.
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
	if args == nil || args.MethodPath == nil {
		return nil, errors.New("missing required argument 'MethodPath'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	if args == nil || args.StageName == nil {
		return nil, errors.New("missing required argument 'StageName'")
	}
	if args == nil {
		args = &MethodSettingsArgs{}
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

