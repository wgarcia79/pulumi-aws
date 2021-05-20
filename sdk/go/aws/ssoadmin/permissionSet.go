// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssoadmin

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Single Sign-On (SSO) Permission Set resource
//
// > **NOTE:** Updating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.
//
// ## Import
//
// SSO Permission Sets can be imported using the `arn` and `instance_arn` separated by a comma (`,`) e.g.
//
// ```sh
//  $ pulumi import aws:ssoadmin/permissionSet:PermissionSet example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
// ```
type PermissionSet struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Permission Set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date the Permission Set was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the Permission Set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The name of the Permission Set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState pulumi.StringPtrOutput `pulumi:"relayState"`
	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration pulumi.StringPtrOutput `pulumi:"sessionDuration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPermissionSet registers a new resource with the given unique name, arguments, and options.
func NewPermissionSet(ctx *pulumi.Context,
	name string, args *PermissionSetArgs, opts ...pulumi.ResourceOption) (*PermissionSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	var resource PermissionSet
	err := ctx.RegisterResource("aws:ssoadmin/permissionSet:PermissionSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionSet gets an existing PermissionSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionSetState, opts ...pulumi.ResourceOption) (*PermissionSet, error) {
	var resource PermissionSet
	err := ctx.ReadResource("aws:ssoadmin/permissionSet:PermissionSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionSet resources.
type permissionSetState struct {
	// The Amazon Resource Name (ARN) of the Permission Set.
	Arn *string `pulumi:"arn"`
	// The date the Permission Set was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the Permission Set.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn *string `pulumi:"instanceArn"`
	// The name of the Permission Set.
	Name *string `pulumi:"name"`
	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState *string `pulumi:"relayState"`
	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration *string `pulumi:"sessionDuration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PermissionSetState struct {
	// The Amazon Resource Name (ARN) of the Permission Set.
	Arn pulumi.StringPtrInput
	// The date the Permission Set was created in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreatedDate pulumi.StringPtrInput
	// The description of the Permission Set.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringPtrInput
	// The name of the Permission Set.
	Name pulumi.StringPtrInput
	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState pulumi.StringPtrInput
	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PermissionSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetState)(nil)).Elem()
}

type permissionSetArgs struct {
	// The description of the Permission Set.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn string `pulumi:"instanceArn"`
	// The name of the Permission Set.
	Name *string `pulumi:"name"`
	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState *string `pulumi:"relayState"`
	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration *string `pulumi:"sessionDuration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a PermissionSet resource.
type PermissionSetArgs struct {
	// The description of the Permission Set.
	Description pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
	InstanceArn pulumi.StringInput
	// The name of the Permission Set.
	Name pulumi.StringPtrInput
	// The relay state URL used to redirect users within the application during the federation authentication process.
	RelayState pulumi.StringPtrInput
	// The length of time that the application user sessions are valid in the ISO-8601 standard. Default: `PT1H`.
	SessionDuration pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (PermissionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionSetArgs)(nil)).Elem()
}

type PermissionSetInput interface {
	pulumi.Input

	ToPermissionSetOutput() PermissionSetOutput
	ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput
}

func (*PermissionSet) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSet)(nil))
}

func (i *PermissionSet) ToPermissionSetOutput() PermissionSetOutput {
	return i.ToPermissionSetOutputWithContext(context.Background())
}

func (i *PermissionSet) ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetOutput)
}

func (i *PermissionSet) ToPermissionSetPtrOutput() PermissionSetPtrOutput {
	return i.ToPermissionSetPtrOutputWithContext(context.Background())
}

func (i *PermissionSet) ToPermissionSetPtrOutputWithContext(ctx context.Context) PermissionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetPtrOutput)
}

type PermissionSetPtrInput interface {
	pulumi.Input

	ToPermissionSetPtrOutput() PermissionSetPtrOutput
	ToPermissionSetPtrOutputWithContext(ctx context.Context) PermissionSetPtrOutput
}

type permissionSetPtrType PermissionSetArgs

func (*permissionSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSet)(nil))
}

func (i *permissionSetPtrType) ToPermissionSetPtrOutput() PermissionSetPtrOutput {
	return i.ToPermissionSetPtrOutputWithContext(context.Background())
}

func (i *permissionSetPtrType) ToPermissionSetPtrOutputWithContext(ctx context.Context) PermissionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetPtrOutput)
}

// PermissionSetArrayInput is an input type that accepts PermissionSetArray and PermissionSetArrayOutput values.
// You can construct a concrete instance of `PermissionSetArrayInput` via:
//
//          PermissionSetArray{ PermissionSetArgs{...} }
type PermissionSetArrayInput interface {
	pulumi.Input

	ToPermissionSetArrayOutput() PermissionSetArrayOutput
	ToPermissionSetArrayOutputWithContext(context.Context) PermissionSetArrayOutput
}

type PermissionSetArray []PermissionSetInput

func (PermissionSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PermissionSet)(nil))
}

func (i PermissionSetArray) ToPermissionSetArrayOutput() PermissionSetArrayOutput {
	return i.ToPermissionSetArrayOutputWithContext(context.Background())
}

func (i PermissionSetArray) ToPermissionSetArrayOutputWithContext(ctx context.Context) PermissionSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetArrayOutput)
}

// PermissionSetMapInput is an input type that accepts PermissionSetMap and PermissionSetMapOutput values.
// You can construct a concrete instance of `PermissionSetMapInput` via:
//
//          PermissionSetMap{ "key": PermissionSetArgs{...} }
type PermissionSetMapInput interface {
	pulumi.Input

	ToPermissionSetMapOutput() PermissionSetMapOutput
	ToPermissionSetMapOutputWithContext(context.Context) PermissionSetMapOutput
}

type PermissionSetMap map[string]PermissionSetInput

func (PermissionSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PermissionSet)(nil))
}

func (i PermissionSetMap) ToPermissionSetMapOutput() PermissionSetMapOutput {
	return i.ToPermissionSetMapOutputWithContext(context.Background())
}

func (i PermissionSetMap) ToPermissionSetMapOutputWithContext(ctx context.Context) PermissionSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionSetMapOutput)
}

type PermissionSetOutput struct {
	*pulumi.OutputState
}

func (PermissionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSet)(nil))
}

func (o PermissionSetOutput) ToPermissionSetOutput() PermissionSetOutput {
	return o
}

func (o PermissionSetOutput) ToPermissionSetOutputWithContext(ctx context.Context) PermissionSetOutput {
	return o
}

func (o PermissionSetOutput) ToPermissionSetPtrOutput() PermissionSetPtrOutput {
	return o.ToPermissionSetPtrOutputWithContext(context.Background())
}

func (o PermissionSetOutput) ToPermissionSetPtrOutputWithContext(ctx context.Context) PermissionSetPtrOutput {
	return o.ApplyT(func(v PermissionSet) *PermissionSet {
		return &v
	}).(PermissionSetPtrOutput)
}

type PermissionSetPtrOutput struct {
	*pulumi.OutputState
}

func (PermissionSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionSet)(nil))
}

func (o PermissionSetPtrOutput) ToPermissionSetPtrOutput() PermissionSetPtrOutput {
	return o
}

func (o PermissionSetPtrOutput) ToPermissionSetPtrOutputWithContext(ctx context.Context) PermissionSetPtrOutput {
	return o
}

type PermissionSetArrayOutput struct{ *pulumi.OutputState }

func (PermissionSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionSet)(nil))
}

func (o PermissionSetArrayOutput) ToPermissionSetArrayOutput() PermissionSetArrayOutput {
	return o
}

func (o PermissionSetArrayOutput) ToPermissionSetArrayOutputWithContext(ctx context.Context) PermissionSetArrayOutput {
	return o
}

func (o PermissionSetArrayOutput) Index(i pulumi.IntInput) PermissionSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionSet {
		return vs[0].([]PermissionSet)[vs[1].(int)]
	}).(PermissionSetOutput)
}

type PermissionSetMapOutput struct{ *pulumi.OutputState }

func (PermissionSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PermissionSet)(nil))
}

func (o PermissionSetMapOutput) ToPermissionSetMapOutput() PermissionSetMapOutput {
	return o
}

func (o PermissionSetMapOutput) ToPermissionSetMapOutputWithContext(ctx context.Context) PermissionSetMapOutput {
	return o
}

func (o PermissionSetMapOutput) MapIndex(k pulumi.StringInput) PermissionSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PermissionSet {
		return vs[0].(map[string]PermissionSet)[vs[1].(string)]
	}).(PermissionSetOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionSetOutput{})
	pulumi.RegisterOutputType(PermissionSetPtrOutput{})
	pulumi.RegisterOutputType(PermissionSetArrayOutput{})
	pulumi.RegisterOutputType(PermissionSetMapOutput{})
}
