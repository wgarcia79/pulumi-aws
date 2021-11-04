// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package chime

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Amazon Chime Voice Connector group under the administrator's AWS account. You can associate Amazon Chime Voice Connectors with the Amazon Chime Voice Connector group by including VoiceConnectorItems in the request.
//
// You can include Amazon Chime Voice Connectors from different AWS Regions in your group. This creates a fault tolerant mechanism for fallback in case of availability events.
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
// 		vc1, err := chime.NewVoiceConnector(ctx, "vc1", &chime.VoiceConnectorArgs{
// 			RequireEncryption: pulumi.Bool(true),
// 			AwsRegion:         pulumi.String("us-east-1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vc2, err := chime.NewVoiceConnector(ctx, "vc2", &chime.VoiceConnectorArgs{
// 			RequireEncryption: pulumi.Bool(true),
// 			AwsRegion:         pulumi.String("us-west-2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = chime.NewVoiceConnectorGroup(ctx, "group", &chime.VoiceConnectorGroupArgs{
// 			Connectors: chime.VoiceConnectorGroupConnectorArray{
// 				&chime.VoiceConnectorGroupConnectorArgs{
// 					VoiceConnectorId: vc1.ID(),
// 					Priority:         pulumi.Int(1),
// 				},
// 				&chime.VoiceConnectorGroupConnectorArgs{
// 					VoiceConnectorId: vc2.ID(),
// 					Priority:         pulumi.Int(3),
// 				},
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
// Configuration Recorder can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:chime/voiceConnectorGroup:VoiceConnectorGroup default example
// ```
type VoiceConnectorGroup struct {
	pulumi.CustomResourceState

	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors VoiceConnectorGroupConnectorArrayOutput `pulumi:"connectors"`
	// The name of the Amazon Chime Voice Connector group.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewVoiceConnectorGroup registers a new resource with the given unique name, arguments, and options.
func NewVoiceConnectorGroup(ctx *pulumi.Context,
	name string, args *VoiceConnectorGroupArgs, opts ...pulumi.ResourceOption) (*VoiceConnectorGroup, error) {
	if args == nil {
		args = &VoiceConnectorGroupArgs{}
	}

	var resource VoiceConnectorGroup
	err := ctx.RegisterResource("aws:chime/voiceConnectorGroup:VoiceConnectorGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVoiceConnectorGroup gets an existing VoiceConnectorGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVoiceConnectorGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VoiceConnectorGroupState, opts ...pulumi.ResourceOption) (*VoiceConnectorGroup, error) {
	var resource VoiceConnectorGroup
	err := ctx.ReadResource("aws:chime/voiceConnectorGroup:VoiceConnectorGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VoiceConnectorGroup resources.
type voiceConnectorGroupState struct {
	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors []VoiceConnectorGroupConnector `pulumi:"connectors"`
	// The name of the Amazon Chime Voice Connector group.
	Name *string `pulumi:"name"`
}

type VoiceConnectorGroupState struct {
	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors VoiceConnectorGroupConnectorArrayInput
	// The name of the Amazon Chime Voice Connector group.
	Name pulumi.StringPtrInput
}

func (VoiceConnectorGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorGroupState)(nil)).Elem()
}

type voiceConnectorGroupArgs struct {
	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors []VoiceConnectorGroupConnector `pulumi:"connectors"`
	// The name of the Amazon Chime Voice Connector group.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a VoiceConnectorGroup resource.
type VoiceConnectorGroupArgs struct {
	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors VoiceConnectorGroupConnectorArrayInput
	// The name of the Amazon Chime Voice Connector group.
	Name pulumi.StringPtrInput
}

func (VoiceConnectorGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*voiceConnectorGroupArgs)(nil)).Elem()
}

type VoiceConnectorGroupInput interface {
	pulumi.Input

	ToVoiceConnectorGroupOutput() VoiceConnectorGroupOutput
	ToVoiceConnectorGroupOutputWithContext(ctx context.Context) VoiceConnectorGroupOutput
}

func (*VoiceConnectorGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceConnectorGroup)(nil))
}

func (i *VoiceConnectorGroup) ToVoiceConnectorGroupOutput() VoiceConnectorGroupOutput {
	return i.ToVoiceConnectorGroupOutputWithContext(context.Background())
}

func (i *VoiceConnectorGroup) ToVoiceConnectorGroupOutputWithContext(ctx context.Context) VoiceConnectorGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorGroupOutput)
}

func (i *VoiceConnectorGroup) ToVoiceConnectorGroupPtrOutput() VoiceConnectorGroupPtrOutput {
	return i.ToVoiceConnectorGroupPtrOutputWithContext(context.Background())
}

func (i *VoiceConnectorGroup) ToVoiceConnectorGroupPtrOutputWithContext(ctx context.Context) VoiceConnectorGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorGroupPtrOutput)
}

type VoiceConnectorGroupPtrInput interface {
	pulumi.Input

	ToVoiceConnectorGroupPtrOutput() VoiceConnectorGroupPtrOutput
	ToVoiceConnectorGroupPtrOutputWithContext(ctx context.Context) VoiceConnectorGroupPtrOutput
}

type voiceConnectorGroupPtrType VoiceConnectorGroupArgs

func (*voiceConnectorGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorGroup)(nil))
}

func (i *voiceConnectorGroupPtrType) ToVoiceConnectorGroupPtrOutput() VoiceConnectorGroupPtrOutput {
	return i.ToVoiceConnectorGroupPtrOutputWithContext(context.Background())
}

func (i *voiceConnectorGroupPtrType) ToVoiceConnectorGroupPtrOutputWithContext(ctx context.Context) VoiceConnectorGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorGroupPtrOutput)
}

// VoiceConnectorGroupArrayInput is an input type that accepts VoiceConnectorGroupArray and VoiceConnectorGroupArrayOutput values.
// You can construct a concrete instance of `VoiceConnectorGroupArrayInput` via:
//
//          VoiceConnectorGroupArray{ VoiceConnectorGroupArgs{...} }
type VoiceConnectorGroupArrayInput interface {
	pulumi.Input

	ToVoiceConnectorGroupArrayOutput() VoiceConnectorGroupArrayOutput
	ToVoiceConnectorGroupArrayOutputWithContext(context.Context) VoiceConnectorGroupArrayOutput
}

type VoiceConnectorGroupArray []VoiceConnectorGroupInput

func (VoiceConnectorGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VoiceConnectorGroup)(nil)).Elem()
}

func (i VoiceConnectorGroupArray) ToVoiceConnectorGroupArrayOutput() VoiceConnectorGroupArrayOutput {
	return i.ToVoiceConnectorGroupArrayOutputWithContext(context.Background())
}

func (i VoiceConnectorGroupArray) ToVoiceConnectorGroupArrayOutputWithContext(ctx context.Context) VoiceConnectorGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorGroupArrayOutput)
}

// VoiceConnectorGroupMapInput is an input type that accepts VoiceConnectorGroupMap and VoiceConnectorGroupMapOutput values.
// You can construct a concrete instance of `VoiceConnectorGroupMapInput` via:
//
//          VoiceConnectorGroupMap{ "key": VoiceConnectorGroupArgs{...} }
type VoiceConnectorGroupMapInput interface {
	pulumi.Input

	ToVoiceConnectorGroupMapOutput() VoiceConnectorGroupMapOutput
	ToVoiceConnectorGroupMapOutputWithContext(context.Context) VoiceConnectorGroupMapOutput
}

type VoiceConnectorGroupMap map[string]VoiceConnectorGroupInput

func (VoiceConnectorGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VoiceConnectorGroup)(nil)).Elem()
}

func (i VoiceConnectorGroupMap) ToVoiceConnectorGroupMapOutput() VoiceConnectorGroupMapOutput {
	return i.ToVoiceConnectorGroupMapOutputWithContext(context.Background())
}

func (i VoiceConnectorGroupMap) ToVoiceConnectorGroupMapOutputWithContext(ctx context.Context) VoiceConnectorGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceConnectorGroupMapOutput)
}

type VoiceConnectorGroupOutput struct{ *pulumi.OutputState }

func (VoiceConnectorGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceConnectorGroup)(nil))
}

func (o VoiceConnectorGroupOutput) ToVoiceConnectorGroupOutput() VoiceConnectorGroupOutput {
	return o
}

func (o VoiceConnectorGroupOutput) ToVoiceConnectorGroupOutputWithContext(ctx context.Context) VoiceConnectorGroupOutput {
	return o
}

func (o VoiceConnectorGroupOutput) ToVoiceConnectorGroupPtrOutput() VoiceConnectorGroupPtrOutput {
	return o.ToVoiceConnectorGroupPtrOutputWithContext(context.Background())
}

func (o VoiceConnectorGroupOutput) ToVoiceConnectorGroupPtrOutputWithContext(ctx context.Context) VoiceConnectorGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VoiceConnectorGroup) *VoiceConnectorGroup {
		return &v
	}).(VoiceConnectorGroupPtrOutput)
}

type VoiceConnectorGroupPtrOutput struct{ *pulumi.OutputState }

func (VoiceConnectorGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VoiceConnectorGroup)(nil))
}

func (o VoiceConnectorGroupPtrOutput) ToVoiceConnectorGroupPtrOutput() VoiceConnectorGroupPtrOutput {
	return o
}

func (o VoiceConnectorGroupPtrOutput) ToVoiceConnectorGroupPtrOutputWithContext(ctx context.Context) VoiceConnectorGroupPtrOutput {
	return o
}

func (o VoiceConnectorGroupPtrOutput) Elem() VoiceConnectorGroupOutput {
	return o.ApplyT(func(v *VoiceConnectorGroup) VoiceConnectorGroup {
		if v != nil {
			return *v
		}
		var ret VoiceConnectorGroup
		return ret
	}).(VoiceConnectorGroupOutput)
}

type VoiceConnectorGroupArrayOutput struct{ *pulumi.OutputState }

func (VoiceConnectorGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceConnectorGroup)(nil))
}

func (o VoiceConnectorGroupArrayOutput) ToVoiceConnectorGroupArrayOutput() VoiceConnectorGroupArrayOutput {
	return o
}

func (o VoiceConnectorGroupArrayOutput) ToVoiceConnectorGroupArrayOutputWithContext(ctx context.Context) VoiceConnectorGroupArrayOutput {
	return o
}

func (o VoiceConnectorGroupArrayOutput) Index(i pulumi.IntInput) VoiceConnectorGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VoiceConnectorGroup {
		return vs[0].([]VoiceConnectorGroup)[vs[1].(int)]
	}).(VoiceConnectorGroupOutput)
}

type VoiceConnectorGroupMapOutput struct{ *pulumi.OutputState }

func (VoiceConnectorGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VoiceConnectorGroup)(nil))
}

func (o VoiceConnectorGroupMapOutput) ToVoiceConnectorGroupMapOutput() VoiceConnectorGroupMapOutput {
	return o
}

func (o VoiceConnectorGroupMapOutput) ToVoiceConnectorGroupMapOutputWithContext(ctx context.Context) VoiceConnectorGroupMapOutput {
	return o
}

func (o VoiceConnectorGroupMapOutput) MapIndex(k pulumi.StringInput) VoiceConnectorGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VoiceConnectorGroup {
		return vs[0].(map[string]VoiceConnectorGroup)[vs[1].(string)]
	}).(VoiceConnectorGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorGroupInput)(nil)).Elem(), &VoiceConnectorGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorGroupPtrInput)(nil)).Elem(), &VoiceConnectorGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorGroupArrayInput)(nil)).Elem(), VoiceConnectorGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceConnectorGroupMapInput)(nil)).Elem(), VoiceConnectorGroupMap{})
	pulumi.RegisterOutputType(VoiceConnectorGroupOutput{})
	pulumi.RegisterOutputType(VoiceConnectorGroupPtrOutput{})
	pulumi.RegisterOutputType(VoiceConnectorGroupArrayOutput{})
	pulumi.RegisterOutputType(VoiceConnectorGroupMapOutput{})
}
