// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package glue

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Glue Security Configuration.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_security_configuration.html.markdown.
type SecurityConfiguration struct {
	pulumi.CustomResourceState

	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration SecurityConfigurationEncryptionConfigurationOutput `pulumi:"encryptionConfiguration"`
	// Name of the security configuration.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSecurityConfiguration registers a new resource with the given unique name, arguments, and options.
func NewSecurityConfiguration(ctx *pulumi.Context,
	name string, args *SecurityConfigurationArgs, opts ...pulumi.ResourceOption) (*SecurityConfiguration, error) {
	if args == nil || args.EncryptionConfiguration == nil {
		return nil, errors.New("missing required argument 'EncryptionConfiguration'")
	}
	if args == nil {
		args = &SecurityConfigurationArgs{}
	}
	var resource SecurityConfiguration
	err := ctx.RegisterResource("aws:glue/securityConfiguration:SecurityConfiguration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:glue/securityConfiguration:SecurityConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityConfiguration resources.
type securityConfigurationState struct {
	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration *SecurityConfigurationEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// Name of the security configuration.
	Name *string `pulumi:"name"`
}

type SecurityConfigurationState struct {
	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration SecurityConfigurationEncryptionConfigurationPtrInput
	// Name of the security configuration.
	Name pulumi.StringPtrInput
}

func (SecurityConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigurationState)(nil)).Elem()
}

type securityConfigurationArgs struct {
	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration SecurityConfigurationEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// Name of the security configuration.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SecurityConfiguration resource.
type SecurityConfigurationArgs struct {
	// Configuration block containing encryption configuration. Detailed below.
	EncryptionConfiguration SecurityConfigurationEncryptionConfigurationInput
	// Name of the security configuration.
	Name pulumi.StringPtrInput
}

func (SecurityConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConfigurationArgs)(nil)).Elem()
}

