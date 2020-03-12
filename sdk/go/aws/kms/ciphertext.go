// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The KMS ciphertext resource allows you to encrypt plaintext into ciphertext
// by using an AWS KMS customer master key. The value returned by this resource
// is stable across every apply. For a changing ciphertext value each apply, see
// the [`kms.Ciphertext` data source](https://www.terraform.io/docs/providers/aws/d/kms_ciphertext.html).
//
// > **Note:** All arguments including the plaintext be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_ciphertext.html.markdown.
type Ciphertext struct {
	pulumi.CustomResourceState

	// Base64 encoded ciphertext
	CiphertextBlob pulumi.StringOutput `pulumi:"ciphertextBlob"`
	// An optional mapping that makes up the encryption context.
	Context pulumi.StringMapOutput `pulumi:"context"`
	// Globally unique key ID for the customer master key.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext pulumi.StringOutput `pulumi:"plaintext"`
}

// NewCiphertext registers a new resource with the given unique name, arguments, and options.
func NewCiphertext(ctx *pulumi.Context,
	name string, args *CiphertextArgs, opts ...pulumi.ResourceOption) (*Ciphertext, error) {
	if args == nil || args.KeyId == nil {
		return nil, errors.New("missing required argument 'KeyId'")
	}
	if args == nil || args.Plaintext == nil {
		return nil, errors.New("missing required argument 'Plaintext'")
	}
	if args == nil {
		args = &CiphertextArgs{}
	}
	var resource Ciphertext
	err := ctx.RegisterResource("aws:kms/ciphertext:Ciphertext", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCiphertext gets an existing Ciphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCiphertext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CiphertextState, opts ...pulumi.ResourceOption) (*Ciphertext, error) {
	var resource Ciphertext
	err := ctx.ReadResource("aws:kms/ciphertext:Ciphertext", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ciphertext resources.
type ciphertextState struct {
	// Base64 encoded ciphertext
	CiphertextBlob *string `pulumi:"ciphertextBlob"`
	// An optional mapping that makes up the encryption context.
	Context map[string]string `pulumi:"context"`
	// Globally unique key ID for the customer master key.
	KeyId *string `pulumi:"keyId"`
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext *string `pulumi:"plaintext"`
}

type CiphertextState struct {
	// Base64 encoded ciphertext
	CiphertextBlob pulumi.StringPtrInput
	// An optional mapping that makes up the encryption context.
	Context pulumi.StringMapInput
	// Globally unique key ID for the customer master key.
	KeyId pulumi.StringPtrInput
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext pulumi.StringPtrInput
}

func (CiphertextState) ElementType() reflect.Type {
	return reflect.TypeOf((*ciphertextState)(nil)).Elem()
}

type ciphertextArgs struct {
	// An optional mapping that makes up the encryption context.
	Context map[string]string `pulumi:"context"`
	// Globally unique key ID for the customer master key.
	KeyId string `pulumi:"keyId"`
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext string `pulumi:"plaintext"`
}

// The set of arguments for constructing a Ciphertext resource.
type CiphertextArgs struct {
	// An optional mapping that makes up the encryption context.
	Context pulumi.StringMapInput
	// Globally unique key ID for the customer master key.
	KeyId pulumi.StringInput
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext pulumi.StringInput
}

func (CiphertextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ciphertextArgs)(nil)).Elem()
}

