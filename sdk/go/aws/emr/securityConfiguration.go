// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage AWS EMR Security Configurations
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/emr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := emr.NewSecurityConfiguration(ctx, "foo", &emr.SecurityConfigurationArgs{
// 			Configuration: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"EncryptionConfiguration\": {\n", "    \"AtRestEncryptionConfiguration\": {\n", "      \"S3EncryptionConfiguration\": {\n", "        \"EncryptionMode\": \"SSE-S3\"\n", "      },\n", "      \"LocalDiskEncryptionConfiguration\": {\n", "        \"EncryptionKeyProviderType\": \"AwsKms\",\n", "        \"AwsKmsKey\": \"arn:aws:kms:us-west-2:187416307283:alias/tf_emr_test_key\"\n", "      }\n", "    },\n", "    \"EnableInTransitEncryption\": false,\n", "    \"EnableAtRestEncryption\": true\n", "  }\n", "}\n", "\n")),
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
// EMR Security Configurations can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:emr/securityConfiguration:SecurityConfiguration sc example-sc-name
// ```
type SecurityConfiguration struct {
	pulumi.CustomResourceState

	// A JSON formatted Security Configuration
	Configuration pulumi.StringOutput `pulumi:"configuration"`
	// Date the Security Configuration was created
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The name of the EMR Security Configuration. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
}

// NewSecurityConfiguration registers a new resource with the given unique name, arguments, and options.
func NewSecurityConfiguration(ctx *pulumi.Context,
	name string, args *SecurityConfigurationArgs, opts ...pulumi.ResourceOption) (*SecurityConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	var resource SecurityConfiguration
	err := ctx.RegisterResource("aws:emr/securityConfiguration:SecurityConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityConfiguration gets an existing SecurityConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConfigurationState, opts ...pulumi.ResourceOption) (*SecurityConfiguration, error) {
	var resource SecurityConfiguration
	err := ctx.ReadResource("aws:emr/securityConfiguration:SecurityConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityConfiguration resources.
type securityConfigurationState struct {
	// A JSON formatted Security Configuration
	Configuration *string `pulumi:"configuration"`
	// Date the Security Configuration was created
	CreationDate *string `pulumi:"creationDate"`
	// The name of the EMR Security Configuration. By default generated by this provider.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
}

type SecurityConfigurationState struct {
	// A JSON formatted Security Configuration
	Configuration pulumi.StringPtrInput
	// Date the Security Configuration was created
	CreationDate pulumi.StringPtrInput
	// The name of the EMR Security Configuration. By default generated by this provider.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
}

func (SecurityConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigurationState)(nil)).Elem()
}

type securityConfigurationArgs struct {
	// A JSON formatted Security Configuration
	Configuration string `pulumi:"configuration"`
	// The name of the EMR Security Configuration. By default generated by this provider.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
}

// The set of arguments for constructing a SecurityConfiguration resource.
type SecurityConfigurationArgs struct {
	// A JSON formatted Security Configuration
	Configuration pulumi.StringInput
	// The name of the EMR Security Configuration. By default generated by this provider.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
}

func (SecurityConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigurationArgs)(nil)).Elem()
}

type SecurityConfigurationInput interface {
	pulumi.Input

	ToSecurityConfigurationOutput() SecurityConfigurationOutput
	ToSecurityConfigurationOutputWithContext(ctx context.Context) SecurityConfigurationOutput
}

func (*SecurityConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfiguration)(nil))
}

func (i *SecurityConfiguration) ToSecurityConfigurationOutput() SecurityConfigurationOutput {
	return i.ToSecurityConfigurationOutputWithContext(context.Background())
}

func (i *SecurityConfiguration) ToSecurityConfigurationOutputWithContext(ctx context.Context) SecurityConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigurationOutput)
}

func (i *SecurityConfiguration) ToSecurityConfigurationPtrOutput() SecurityConfigurationPtrOutput {
	return i.ToSecurityConfigurationPtrOutputWithContext(context.Background())
}

func (i *SecurityConfiguration) ToSecurityConfigurationPtrOutputWithContext(ctx context.Context) SecurityConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigurationPtrOutput)
}

type SecurityConfigurationPtrInput interface {
	pulumi.Input

	ToSecurityConfigurationPtrOutput() SecurityConfigurationPtrOutput
	ToSecurityConfigurationPtrOutputWithContext(ctx context.Context) SecurityConfigurationPtrOutput
}

type securityConfigurationPtrType SecurityConfigurationArgs

func (*securityConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfiguration)(nil))
}

func (i *securityConfigurationPtrType) ToSecurityConfigurationPtrOutput() SecurityConfigurationPtrOutput {
	return i.ToSecurityConfigurationPtrOutputWithContext(context.Background())
}

func (i *securityConfigurationPtrType) ToSecurityConfigurationPtrOutputWithContext(ctx context.Context) SecurityConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigurationPtrOutput)
}

// SecurityConfigurationArrayInput is an input type that accepts SecurityConfigurationArray and SecurityConfigurationArrayOutput values.
// You can construct a concrete instance of `SecurityConfigurationArrayInput` via:
//
//          SecurityConfigurationArray{ SecurityConfigurationArgs{...} }
type SecurityConfigurationArrayInput interface {
	pulumi.Input

	ToSecurityConfigurationArrayOutput() SecurityConfigurationArrayOutput
	ToSecurityConfigurationArrayOutputWithContext(context.Context) SecurityConfigurationArrayOutput
}

type SecurityConfigurationArray []SecurityConfigurationInput

func (SecurityConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityConfiguration)(nil)).Elem()
}

func (i SecurityConfigurationArray) ToSecurityConfigurationArrayOutput() SecurityConfigurationArrayOutput {
	return i.ToSecurityConfigurationArrayOutputWithContext(context.Background())
}

func (i SecurityConfigurationArray) ToSecurityConfigurationArrayOutputWithContext(ctx context.Context) SecurityConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigurationArrayOutput)
}

// SecurityConfigurationMapInput is an input type that accepts SecurityConfigurationMap and SecurityConfigurationMapOutput values.
// You can construct a concrete instance of `SecurityConfigurationMapInput` via:
//
//          SecurityConfigurationMap{ "key": SecurityConfigurationArgs{...} }
type SecurityConfigurationMapInput interface {
	pulumi.Input

	ToSecurityConfigurationMapOutput() SecurityConfigurationMapOutput
	ToSecurityConfigurationMapOutputWithContext(context.Context) SecurityConfigurationMapOutput
}

type SecurityConfigurationMap map[string]SecurityConfigurationInput

func (SecurityConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityConfiguration)(nil)).Elem()
}

func (i SecurityConfigurationMap) ToSecurityConfigurationMapOutput() SecurityConfigurationMapOutput {
	return i.ToSecurityConfigurationMapOutputWithContext(context.Background())
}

func (i SecurityConfigurationMap) ToSecurityConfigurationMapOutputWithContext(ctx context.Context) SecurityConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConfigurationMapOutput)
}

type SecurityConfigurationOutput struct {
	*pulumi.OutputState
}

func (SecurityConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConfiguration)(nil))
}

func (o SecurityConfigurationOutput) ToSecurityConfigurationOutput() SecurityConfigurationOutput {
	return o
}

func (o SecurityConfigurationOutput) ToSecurityConfigurationOutputWithContext(ctx context.Context) SecurityConfigurationOutput {
	return o
}

func (o SecurityConfigurationOutput) ToSecurityConfigurationPtrOutput() SecurityConfigurationPtrOutput {
	return o.ToSecurityConfigurationPtrOutputWithContext(context.Background())
}

func (o SecurityConfigurationOutput) ToSecurityConfigurationPtrOutputWithContext(ctx context.Context) SecurityConfigurationPtrOutput {
	return o.ApplyT(func(v SecurityConfiguration) *SecurityConfiguration {
		return &v
	}).(SecurityConfigurationPtrOutput)
}

type SecurityConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (SecurityConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConfiguration)(nil))
}

func (o SecurityConfigurationPtrOutput) ToSecurityConfigurationPtrOutput() SecurityConfigurationPtrOutput {
	return o
}

func (o SecurityConfigurationPtrOutput) ToSecurityConfigurationPtrOutputWithContext(ctx context.Context) SecurityConfigurationPtrOutput {
	return o
}

type SecurityConfigurationArrayOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityConfiguration)(nil))
}

func (o SecurityConfigurationArrayOutput) ToSecurityConfigurationArrayOutput() SecurityConfigurationArrayOutput {
	return o
}

func (o SecurityConfigurationArrayOutput) ToSecurityConfigurationArrayOutputWithContext(ctx context.Context) SecurityConfigurationArrayOutput {
	return o
}

func (o SecurityConfigurationArrayOutput) Index(i pulumi.IntInput) SecurityConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityConfiguration {
		return vs[0].([]SecurityConfiguration)[vs[1].(int)]
	}).(SecurityConfigurationOutput)
}

type SecurityConfigurationMapOutput struct{ *pulumi.OutputState }

func (SecurityConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecurityConfiguration)(nil))
}

func (o SecurityConfigurationMapOutput) ToSecurityConfigurationMapOutput() SecurityConfigurationMapOutput {
	return o
}

func (o SecurityConfigurationMapOutput) ToSecurityConfigurationMapOutputWithContext(ctx context.Context) SecurityConfigurationMapOutput {
	return o
}

func (o SecurityConfigurationMapOutput) MapIndex(k pulumi.StringInput) SecurityConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecurityConfiguration {
		return vs[0].(map[string]SecurityConfiguration)[vs[1].(string)]
	}).(SecurityConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityConfigurationOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationArrayOutput{})
	pulumi.RegisterOutputType(SecurityConfigurationMapOutput{})
}
