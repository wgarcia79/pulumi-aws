// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The KMS ciphertext resource allows you to encrypt plaintext into ciphertext
// by using an AWS KMS customer master key. The value returned by this resource
// is stable across every apply. For a changing ciphertext value each apply, see
// the `kms.Ciphertext` data source.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		oauthConfig, err := kms.NewKey(ctx, "oauthConfig", &kms.KeyArgs{
// 			Description: pulumi.String("oauth config"),
// 			IsEnabled:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCiphertext(ctx, "oauth", &kms.CiphertextArgs{
// 			KeyId:     oauthConfig.KeyId,
// 			Plaintext: pulumi.String(fmt.Sprintf("%v%v%v%v", "{\n", "  \"client_id\": \"e587dbae22222f55da22\",\n", "  \"client_secret\": \"8289575d00000ace55e1815ec13673955721b8a5\"\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyId == nil {
		return nil, errors.New("invalid value for required argument 'KeyId'")
	}
	if args.Plaintext == nil {
		return nil, errors.New("invalid value for required argument 'Plaintext'")
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

type CiphertextInput interface {
	pulumi.Input

	ToCiphertextOutput() CiphertextOutput
	ToCiphertextOutputWithContext(ctx context.Context) CiphertextOutput
}

func (*Ciphertext) ElementType() reflect.Type {
	return reflect.TypeOf((*Ciphertext)(nil))
}

func (i *Ciphertext) ToCiphertextOutput() CiphertextOutput {
	return i.ToCiphertextOutputWithContext(context.Background())
}

func (i *Ciphertext) ToCiphertextOutputWithContext(ctx context.Context) CiphertextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CiphertextOutput)
}

func (i *Ciphertext) ToCiphertextPtrOutput() CiphertextPtrOutput {
	return i.ToCiphertextPtrOutputWithContext(context.Background())
}

func (i *Ciphertext) ToCiphertextPtrOutputWithContext(ctx context.Context) CiphertextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CiphertextPtrOutput)
}

type CiphertextPtrInput interface {
	pulumi.Input

	ToCiphertextPtrOutput() CiphertextPtrOutput
	ToCiphertextPtrOutputWithContext(ctx context.Context) CiphertextPtrOutput
}

type ciphertextPtrType CiphertextArgs

func (*ciphertextPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ciphertext)(nil))
}

func (i *ciphertextPtrType) ToCiphertextPtrOutput() CiphertextPtrOutput {
	return i.ToCiphertextPtrOutputWithContext(context.Background())
}

func (i *ciphertextPtrType) ToCiphertextPtrOutputWithContext(ctx context.Context) CiphertextPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CiphertextPtrOutput)
}

// CiphertextArrayInput is an input type that accepts CiphertextArray and CiphertextArrayOutput values.
// You can construct a concrete instance of `CiphertextArrayInput` via:
//
//          CiphertextArray{ CiphertextArgs{...} }
type CiphertextArrayInput interface {
	pulumi.Input

	ToCiphertextArrayOutput() CiphertextArrayOutput
	ToCiphertextArrayOutputWithContext(context.Context) CiphertextArrayOutput
}

type CiphertextArray []CiphertextInput

func (CiphertextArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Ciphertext)(nil))
}

func (i CiphertextArray) ToCiphertextArrayOutput() CiphertextArrayOutput {
	return i.ToCiphertextArrayOutputWithContext(context.Background())
}

func (i CiphertextArray) ToCiphertextArrayOutputWithContext(ctx context.Context) CiphertextArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CiphertextArrayOutput)
}

// CiphertextMapInput is an input type that accepts CiphertextMap and CiphertextMapOutput values.
// You can construct a concrete instance of `CiphertextMapInput` via:
//
//          CiphertextMap{ "key": CiphertextArgs{...} }
type CiphertextMapInput interface {
	pulumi.Input

	ToCiphertextMapOutput() CiphertextMapOutput
	ToCiphertextMapOutputWithContext(context.Context) CiphertextMapOutput
}

type CiphertextMap map[string]CiphertextInput

func (CiphertextMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Ciphertext)(nil))
}

func (i CiphertextMap) ToCiphertextMapOutput() CiphertextMapOutput {
	return i.ToCiphertextMapOutputWithContext(context.Background())
}

func (i CiphertextMap) ToCiphertextMapOutputWithContext(ctx context.Context) CiphertextMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CiphertextMapOutput)
}

type CiphertextOutput struct {
	*pulumi.OutputState
}

func (CiphertextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ciphertext)(nil))
}

func (o CiphertextOutput) ToCiphertextOutput() CiphertextOutput {
	return o
}

func (o CiphertextOutput) ToCiphertextOutputWithContext(ctx context.Context) CiphertextOutput {
	return o
}

func (o CiphertextOutput) ToCiphertextPtrOutput() CiphertextPtrOutput {
	return o.ToCiphertextPtrOutputWithContext(context.Background())
}

func (o CiphertextOutput) ToCiphertextPtrOutputWithContext(ctx context.Context) CiphertextPtrOutput {
	return o.ApplyT(func(v Ciphertext) *Ciphertext {
		return &v
	}).(CiphertextPtrOutput)
}

type CiphertextPtrOutput struct {
	*pulumi.OutputState
}

func (CiphertextPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ciphertext)(nil))
}

func (o CiphertextPtrOutput) ToCiphertextPtrOutput() CiphertextPtrOutput {
	return o
}

func (o CiphertextPtrOutput) ToCiphertextPtrOutputWithContext(ctx context.Context) CiphertextPtrOutput {
	return o
}

type CiphertextArrayOutput struct{ *pulumi.OutputState }

func (CiphertextArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Ciphertext)(nil))
}

func (o CiphertextArrayOutput) ToCiphertextArrayOutput() CiphertextArrayOutput {
	return o
}

func (o CiphertextArrayOutput) ToCiphertextArrayOutputWithContext(ctx context.Context) CiphertextArrayOutput {
	return o
}

func (o CiphertextArrayOutput) Index(i pulumi.IntInput) CiphertextOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Ciphertext {
		return vs[0].([]Ciphertext)[vs[1].(int)]
	}).(CiphertextOutput)
}

type CiphertextMapOutput struct{ *pulumi.OutputState }

func (CiphertextMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Ciphertext)(nil))
}

func (o CiphertextMapOutput) ToCiphertextMapOutput() CiphertextMapOutput {
	return o
}

func (o CiphertextMapOutput) ToCiphertextMapOutputWithContext(ctx context.Context) CiphertextMapOutput {
	return o
}

func (o CiphertextMapOutput) MapIndex(k pulumi.StringInput) CiphertextOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Ciphertext {
		return vs[0].(map[string]Ciphertext)[vs[1].(string)]
	}).(CiphertextOutput)
}

func init() {
	pulumi.RegisterOutputType(CiphertextOutput{})
	pulumi.RegisterOutputType(CiphertextPtrOutput{})
	pulumi.RegisterOutputType(CiphertextArrayOutput{})
	pulumi.RegisterOutputType(CiphertextMapOutput{})
}
