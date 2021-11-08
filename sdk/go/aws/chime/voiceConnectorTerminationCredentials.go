// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package chime

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
//
// > **Note:** Voice Connector Termination Credentials requires a [Voice Connector Termination](https://www.terraform.io/docs/providers/aws/r/chime_voice_connector_termination.html) to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/chime"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultVoiceConnector, err := chime.NewVoiceConnector(ctx, "defaultVoiceConnector", &chime.VoiceConnectorArgs{
// 			RequireEncryption: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultVoiceConnectorTermination, err := chime.NewVoiceConnectorTermination(ctx, "defaultVoiceConnectorTermination", &chime.VoiceConnectorTerminationArgs{
// 			Disabled: pulumi.Bool(true),
// 			CpsLimit: pulumi.Int(1),
// 			CidrAllowLists: pulumi.StringArray{
// 				pulumi.String("50.35.78.96/31"),
// 			},
// 			CallingRegions: pulumi.StringArray{
// 				pulumi.String("US"),
// 				pulumi.String("CA"),
// 			},
// 			VoiceConnectorId: defaultVoiceConnector.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = chime.NewVoiceConnectorTerminationCredentials(ctx, "defaultVoiceConnectorTerminationCredentials", &chime.VoiceConnectorTerminationCredentialsArgs{
// 			VoiceConnectorId: defaultVoiceConnector.ID(),
// 			Credentials: chime.VoiceConnectorTerminationCredentialsCredentialArray{
// 				&chime.VoiceConnectorTerminationCredentialsCredentialArgs{
// 					Username: pulumi.String("test"),
// 					Password: pulumi.String("test!"),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			defaultVoiceConnectorTermination,
// 		}))
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
// Chime Voice Connector Termination Credentials can be imported using the `voice_connector_id`, e.g.,
//
// ```sh
//  $ pulumi import aws:chime/voiceConnectorTerminationCredentials:VoiceConnectorTerminationCredentials default abcdef1ghij2klmno3pqr4
// ```
type VoiceConnectorTerminationCredentials struct {
	pulumi.CustomResourceState

	// List of termination SIP credentials.
	Credentials VoiceConnectorTerminationCredentialsCredentialArrayOutput `pulumi:"credentials"`
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringOutput `pulumi:"voiceConnectorId"`
}

// NewVoiceConnectorTerminationCredentials registers a new resource with the given unique name, arguments, and options.
func NewVoiceConnectorTerminationCredentials(ctx *pulumi.Context,
	name string, args *VoiceConnectorTerminationCredentialsArgs, opts ...pulumi.ResourceOption) (*VoiceConnectorTerminationCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Credentials == nil {
		return nil, errors.New("invalid value for required argument 'Credentials'")
	}
	if args.VoiceConnectorId == nil {
		return nil, errors.New("invalid value for required argument 'VoiceConnectorId'")
	}
	var resource VoiceConnectorTerminationCredentials
	err := ctx.RegisterResource("aws:chime/voiceConnectorTerminationCredentials:VoiceConnectorTerminationCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVoiceConnectorTerminationCredentials gets an existing VoiceConnectorTerminationCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVoiceConnectorTerminationCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VoiceConnectorTerminationCredentialsState, opts ...pulumi.ResourceOption) (*VoiceConnectorTerminationCredentials, error) {
	var resource VoiceConnectorTerminationCredentials
	err := ctx.ReadResource("aws:chime/voiceConnectorTerminationCredentials:VoiceConnectorTerminationCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VoiceConnectorTerminationCredentials resources.
type voiceConnectorTerminationCredentialsState struct {
	// List of termination SIP credentials.
	Credentials []VoiceConnectorTerminationCredentialsCredential `pulumi:"credentials"`
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId *string `pulumi:"voiceConnectorId"`
}

type VoiceConnectorTerminationCredentialsState struct {
	// List of termination SIP credentials.
	Credentials VoiceConnectorTerminationCredentialsCredentialArrayInput
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringPtrInput
}

func (VoiceConnectorTerminationCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorTerminationCredentialsState)(nil)).Elem()
}

type voiceConnectorTerminationCredentialsArgs struct {
	// List of termination SIP credentials.
	Credentials []VoiceConnectorTerminationCredentialsCredential `pulumi:"credentials"`
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId string `pulumi:"voiceConnectorId"`
}

// The set of arguments for constructing a VoiceConnectorTerminationCredentials resource.
type VoiceConnectorTerminationCredentialsArgs struct {
	// List of termination SIP credentials.
	Credentials VoiceConnectorTerminationCredentialsCredentialArrayInput
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId pulumi.StringInput
}

func (VoiceConnectorTerminationCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorTerminationCredentialsArgs)(nil)).Elem()
}

type VoiceConnectorTerminationCredentialsInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationCredentialsOutput() VoiceConnectorTerminationCredentialsOutput
	ToVoiceConnectorTerminationCredentialsOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsOutput
}

func (*VoiceConnectorTerminationCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceConnectorTerminationCredentials)(nil))
}

func (i *VoiceConnectorTerminationCredentials) ToVoiceConnectorTerminationCredentialsOutput() VoiceConnectorTerminationCredentialsOutput {
	return i.ToVoiceConnectorTerminationCredentialsOutputWithContext(context.Background())
}

func (i *VoiceConnectorTerminationCredentials) ToVoiceConnectorTerminationCredentialsOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationCredentialsOutput)
}

func (i *VoiceConnectorTerminationCredentials) ToVoiceConnectorTerminationCredentialsPtrOutput() VoiceConnectorTerminationCredentialsPtrOutput {
	return i.ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(context.Background())
}

func (i *VoiceConnectorTerminationCredentials) ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationCredentialsPtrOutput)
}

type VoiceConnectorTerminationCredentialsPtrInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationCredentialsPtrOutput() VoiceConnectorTerminationCredentialsPtrOutput
	ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsPtrOutput
}

type voiceConnectorTerminationCredentialsPtrType VoiceConnectorTerminationCredentialsArgs

func (*voiceConnectorTerminationCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorTerminationCredentials)(nil))
}

func (i *voiceConnectorTerminationCredentialsPtrType) ToVoiceConnectorTerminationCredentialsPtrOutput() VoiceConnectorTerminationCredentialsPtrOutput {
	return i.ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(context.Background())
}

func (i *voiceConnectorTerminationCredentialsPtrType) ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationCredentialsPtrOutput)
}

// VoiceConnectorTerminationCredentialsArrayInput is an input type that accepts VoiceConnectorTerminationCredentialsArray and VoiceConnectorTerminationCredentialsArrayOutput values.
// You can construct a concrete instance of `VoiceConnectorTerminationCredentialsArrayInput` via:
//
//          VoiceConnectorTerminationCredentialsArray{ VoiceConnectorTerminationCredentialsArgs{...} }
type VoiceConnectorTerminationCredentialsArrayInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationCredentialsArrayOutput() VoiceConnectorTerminationCredentialsArrayOutput
	ToVoiceConnectorTerminationCredentialsArrayOutputWithContext(context.Context) VoiceConnectorTerminationCredentialsArrayOutput
}

type VoiceConnectorTerminationCredentialsArray []VoiceConnectorTerminationCredentialsInput

func (VoiceConnectorTerminationCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoiceConnectorTerminationCredentials)(nil)).Elem()
}

func (i VoiceConnectorTerminationCredentialsArray) ToVoiceConnectorTerminationCredentialsArrayOutput() VoiceConnectorTerminationCredentialsArrayOutput {
	return i.ToVoiceConnectorTerminationCredentialsArrayOutputWithContext(context.Background())
}

func (i VoiceConnectorTerminationCredentialsArray) ToVoiceConnectorTerminationCredentialsArrayOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationCredentialsArrayOutput)
}

// VoiceConnectorTerminationCredentialsMapInput is an input type that accepts VoiceConnectorTerminationCredentialsMap and VoiceConnectorTerminationCredentialsMapOutput values.
// You can construct a concrete instance of `VoiceConnectorTerminationCredentialsMapInput` via:
//
//          VoiceConnectorTerminationCredentialsMap{ "key": VoiceConnectorTerminationCredentialsArgs{...} }
type VoiceConnectorTerminationCredentialsMapInput interface {
	pulumi.Input

	ToVoiceConnectorTerminationCredentialsMapOutput() VoiceConnectorTerminationCredentialsMapOutput
	ToVoiceConnectorTerminationCredentialsMapOutputWithContext(context.Context) VoiceConnectorTerminationCredentialsMapOutput
}

type VoiceConnectorTerminationCredentialsMap map[string]VoiceConnectorTerminationCredentialsInput

func (VoiceConnectorTerminationCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoiceConnectorTerminationCredentials)(nil)).Elem()
}

func (i VoiceConnectorTerminationCredentialsMap) ToVoiceConnectorTerminationCredentialsMapOutput() VoiceConnectorTerminationCredentialsMapOutput {
	return i.ToVoiceConnectorTerminationCredentialsMapOutputWithContext(context.Background())
}

func (i VoiceConnectorTerminationCredentialsMap) ToVoiceConnectorTerminationCredentialsMapOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorTerminationCredentialsMapOutput)
}

type VoiceConnectorTerminationCredentialsOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceConnectorTerminationCredentials)(nil))
}

func (o VoiceConnectorTerminationCredentialsOutput) ToVoiceConnectorTerminationCredentialsOutput() VoiceConnectorTerminationCredentialsOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsOutput) ToVoiceConnectorTerminationCredentialsOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsOutput) ToVoiceConnectorTerminationCredentialsPtrOutput() VoiceConnectorTerminationCredentialsPtrOutput {
	return o.ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(context.Background())
}

func (o VoiceConnectorTerminationCredentialsOutput) ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VoiceConnectorTerminationCredentials) *VoiceConnectorTerminationCredentials {
		return &v
	}).(VoiceConnectorTerminationCredentialsPtrOutput)
}

type VoiceConnectorTerminationCredentialsPtrOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorTerminationCredentials)(nil))
}

func (o VoiceConnectorTerminationCredentialsPtrOutput) ToVoiceConnectorTerminationCredentialsPtrOutput() VoiceConnectorTerminationCredentialsPtrOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsPtrOutput) ToVoiceConnectorTerminationCredentialsPtrOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsPtrOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsPtrOutput) Elem() VoiceConnectorTerminationCredentialsOutput {
	return o.ApplyT(func(v *VoiceConnectorTerminationCredentials) VoiceConnectorTerminationCredentials {
		if v != nil {
			return *v
		}
		var ret VoiceConnectorTerminationCredentials
		return ret
	}).(VoiceConnectorTerminationCredentialsOutput)
}

type VoiceConnectorTerminationCredentialsArrayOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceConnectorTerminationCredentials)(nil))
}

func (o VoiceConnectorTerminationCredentialsArrayOutput) ToVoiceConnectorTerminationCredentialsArrayOutput() VoiceConnectorTerminationCredentialsArrayOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsArrayOutput) ToVoiceConnectorTerminationCredentialsArrayOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsArrayOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsArrayOutput) Index(i pulumi.IntInput) VoiceConnectorTerminationCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VoiceConnectorTerminationCredentials {
		return vs[0].([]VoiceConnectorTerminationCredentials)[vs[1].(int)]
	}).(VoiceConnectorTerminationCredentialsOutput)
}

type VoiceConnectorTerminationCredentialsMapOutput struct{ *pulumi.OutputState }

func (VoiceConnectorTerminationCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VoiceConnectorTerminationCredentials)(nil))
}

func (o VoiceConnectorTerminationCredentialsMapOutput) ToVoiceConnectorTerminationCredentialsMapOutput() VoiceConnectorTerminationCredentialsMapOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsMapOutput) ToVoiceConnectorTerminationCredentialsMapOutputWithContext(ctx context.Context) VoiceConnectorTerminationCredentialsMapOutput {
	return o
}

func (o VoiceConnectorTerminationCredentialsMapOutput) MapIndex(k pulumi.StringInput) VoiceConnectorTerminationCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VoiceConnectorTerminationCredentials {
		return vs[0].(map[string]VoiceConnectorTerminationCredentials)[vs[1].(string)]
	}).(VoiceConnectorTerminationCredentialsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationCredentialsInput)(nil)).Elem(), &VoiceConnectorTerminationCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationCredentialsPtrInput)(nil)).Elem(), &VoiceConnectorTerminationCredentials{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationCredentialsArrayInput)(nil)).Elem(), VoiceConnectorTerminationCredentialsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorTerminationCredentialsMapInput)(nil)).Elem(), VoiceConnectorTerminationCredentialsMap{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationCredentialsOutput{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationCredentialsPtrOutput{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationCredentialsArrayOutput{})
	pulumi.RegisterOutputType(VoiceConnectorTerminationCredentialsMapOutput{})
}
