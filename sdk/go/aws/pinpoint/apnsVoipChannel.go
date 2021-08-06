// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Pinpoint APNs VoIP Channel resource.
//
// ## Import
//
// Pinpoint APNs VoIP Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/apnsVoipChannel:ApnsVoipChannel apns_voip application-id
// ```
type ApnsVoipChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrOutput `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrOutput `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewApnsVoipChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsVoipChannel(ctx *pulumi.Context,
	name string, args *ApnsVoipChannelArgs, opts ...pulumi.ResourceOption) (*ApnsVoipChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource ApnsVoipChannel
	err := ctx.RegisterResource("aws:pinpoint/apnsVoipChannel:ApnsVoipChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsVoipChannel gets an existing ApnsVoipChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsVoipChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsVoipChannelState, opts ...pulumi.ResourceOption) (*ApnsVoipChannel, error) {
	var resource ApnsVoipChannel
	err := ctx.ReadResource("aws:pinpoint/apnsVoipChannel:ApnsVoipChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsVoipChannel resources.
type apnsVoipChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

type ApnsVoipChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsVoipChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsVoipChannelState)(nil)).Elem()
}

type apnsVoipChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a ApnsVoipChannel resource.
type ApnsVoipChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsVoipChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsVoipChannelArgs)(nil)).Elem()
}

type ApnsVoipChannelInput interface {
	pulumi.Input

	ToApnsVoipChannelOutput() ApnsVoipChannelOutput
	ToApnsVoipChannelOutputWithContext(ctx context.Context) ApnsVoipChannelOutput
}

func (*ApnsVoipChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsVoipChannel)(nil))
}

func (i *ApnsVoipChannel) ToApnsVoipChannelOutput() ApnsVoipChannelOutput {
	return i.ToApnsVoipChannelOutputWithContext(context.Background())
}

func (i *ApnsVoipChannel) ToApnsVoipChannelOutputWithContext(ctx context.Context) ApnsVoipChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipChannelOutput)
}

func (i *ApnsVoipChannel) ToApnsVoipChannelPtrOutput() ApnsVoipChannelPtrOutput {
	return i.ToApnsVoipChannelPtrOutputWithContext(context.Background())
}

func (i *ApnsVoipChannel) ToApnsVoipChannelPtrOutputWithContext(ctx context.Context) ApnsVoipChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipChannelPtrOutput)
}

type ApnsVoipChannelPtrInput interface {
	pulumi.Input

	ToApnsVoipChannelPtrOutput() ApnsVoipChannelPtrOutput
	ToApnsVoipChannelPtrOutputWithContext(ctx context.Context) ApnsVoipChannelPtrOutput
}

type apnsVoipChannelPtrType ApnsVoipChannelArgs

func (*apnsVoipChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsVoipChannel)(nil))
}

func (i *apnsVoipChannelPtrType) ToApnsVoipChannelPtrOutput() ApnsVoipChannelPtrOutput {
	return i.ToApnsVoipChannelPtrOutputWithContext(context.Background())
}

func (i *apnsVoipChannelPtrType) ToApnsVoipChannelPtrOutputWithContext(ctx context.Context) ApnsVoipChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipChannelPtrOutput)
}

// ApnsVoipChannelArrayInput is an input type that accepts ApnsVoipChannelArray and ApnsVoipChannelArrayOutput values.
// You can construct a concrete instance of `ApnsVoipChannelArrayInput` via:
//
//          ApnsVoipChannelArray{ ApnsVoipChannelArgs{...} }
type ApnsVoipChannelArrayInput interface {
	pulumi.Input

	ToApnsVoipChannelArrayOutput() ApnsVoipChannelArrayOutput
	ToApnsVoipChannelArrayOutputWithContext(context.Context) ApnsVoipChannelArrayOutput
}

type ApnsVoipChannelArray []ApnsVoipChannelInput

func (ApnsVoipChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApnsVoipChannel)(nil)).Elem()
}

func (i ApnsVoipChannelArray) ToApnsVoipChannelArrayOutput() ApnsVoipChannelArrayOutput {
	return i.ToApnsVoipChannelArrayOutputWithContext(context.Background())
}

func (i ApnsVoipChannelArray) ToApnsVoipChannelArrayOutputWithContext(ctx context.Context) ApnsVoipChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipChannelArrayOutput)
}

// ApnsVoipChannelMapInput is an input type that accepts ApnsVoipChannelMap and ApnsVoipChannelMapOutput values.
// You can construct a concrete instance of `ApnsVoipChannelMapInput` via:
//
//          ApnsVoipChannelMap{ "key": ApnsVoipChannelArgs{...} }
type ApnsVoipChannelMapInput interface {
	pulumi.Input

	ToApnsVoipChannelMapOutput() ApnsVoipChannelMapOutput
	ToApnsVoipChannelMapOutputWithContext(context.Context) ApnsVoipChannelMapOutput
}

type ApnsVoipChannelMap map[string]ApnsVoipChannelInput

func (ApnsVoipChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApnsVoipChannel)(nil)).Elem()
}

func (i ApnsVoipChannelMap) ToApnsVoipChannelMapOutput() ApnsVoipChannelMapOutput {
	return i.ToApnsVoipChannelMapOutputWithContext(context.Background())
}

func (i ApnsVoipChannelMap) ToApnsVoipChannelMapOutputWithContext(ctx context.Context) ApnsVoipChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsVoipChannelMapOutput)
}

type ApnsVoipChannelOutput struct {
	*pulumi.OutputState
}

func (ApnsVoipChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsVoipChannel)(nil))
}

func (o ApnsVoipChannelOutput) ToApnsVoipChannelOutput() ApnsVoipChannelOutput {
	return o
}

func (o ApnsVoipChannelOutput) ToApnsVoipChannelOutputWithContext(ctx context.Context) ApnsVoipChannelOutput {
	return o
}

func (o ApnsVoipChannelOutput) ToApnsVoipChannelPtrOutput() ApnsVoipChannelPtrOutput {
	return o.ToApnsVoipChannelPtrOutputWithContext(context.Background())
}

func (o ApnsVoipChannelOutput) ToApnsVoipChannelPtrOutputWithContext(ctx context.Context) ApnsVoipChannelPtrOutput {
	return o.ApplyT(func(v ApnsVoipChannel) *ApnsVoipChannel {
		return &v
	}).(ApnsVoipChannelPtrOutput)
}

type ApnsVoipChannelPtrOutput struct {
	*pulumi.OutputState
}

func (ApnsVoipChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsVoipChannel)(nil))
}

func (o ApnsVoipChannelPtrOutput) ToApnsVoipChannelPtrOutput() ApnsVoipChannelPtrOutput {
	return o
}

func (o ApnsVoipChannelPtrOutput) ToApnsVoipChannelPtrOutputWithContext(ctx context.Context) ApnsVoipChannelPtrOutput {
	return o
}

type ApnsVoipChannelArrayOutput struct{ *pulumi.OutputState }

func (ApnsVoipChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApnsVoipChannel)(nil))
}

func (o ApnsVoipChannelArrayOutput) ToApnsVoipChannelArrayOutput() ApnsVoipChannelArrayOutput {
	return o
}

func (o ApnsVoipChannelArrayOutput) ToApnsVoipChannelArrayOutputWithContext(ctx context.Context) ApnsVoipChannelArrayOutput {
	return o
}

func (o ApnsVoipChannelArrayOutput) Index(i pulumi.IntInput) ApnsVoipChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApnsVoipChannel {
		return vs[0].([]ApnsVoipChannel)[vs[1].(int)]
	}).(ApnsVoipChannelOutput)
}

type ApnsVoipChannelMapOutput struct{ *pulumi.OutputState }

func (ApnsVoipChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApnsVoipChannel)(nil))
}

func (o ApnsVoipChannelMapOutput) ToApnsVoipChannelMapOutput() ApnsVoipChannelMapOutput {
	return o
}

func (o ApnsVoipChannelMapOutput) ToApnsVoipChannelMapOutputWithContext(ctx context.Context) ApnsVoipChannelMapOutput {
	return o
}

func (o ApnsVoipChannelMapOutput) MapIndex(k pulumi.StringInput) ApnsVoipChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApnsVoipChannel {
		return vs[0].(map[string]ApnsVoipChannel)[vs[1].(string)]
	}).(ApnsVoipChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(ApnsVoipChannelOutput{})
	pulumi.RegisterOutputType(ApnsVoipChannelPtrOutput{})
	pulumi.RegisterOutputType(ApnsVoipChannelArrayOutput{})
	pulumi.RegisterOutputType(ApnsVoipChannelMapOutput{})
}
