// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecrpublic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RepositoryCatalogData struct {
	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	AboutText *string `pulumi:"aboutText"`
	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
	Architectures []string `pulumi:"architectures"`
	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	Description *string `pulumi:"description"`
	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	LogoImageBlob *string `pulumi:"logoImageBlob"`
	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. `ARM`, `ARM 64`, `x86`, `x86-64`
	OperatingSystems []string `pulumi:"operatingSystems"`
	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	UsageText *string `pulumi:"usageText"`
}

// RepositoryCatalogDataInput is an input type that accepts RepositoryCatalogDataArgs and RepositoryCatalogDataOutput values.
// You can construct a concrete instance of `RepositoryCatalogDataInput` via:
//
//          RepositoryCatalogDataArgs{...}
type RepositoryCatalogDataInput interface {
	pulumi.Input

	ToRepositoryCatalogDataOutput() RepositoryCatalogDataOutput
	ToRepositoryCatalogDataOutputWithContext(context.Context) RepositoryCatalogDataOutput
}

type RepositoryCatalogDataArgs struct {
	// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
	AboutText pulumi.StringPtrInput `pulumi:"aboutText"`
	// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
	Architectures pulumi.StringArrayInput `pulumi:"architectures"`
	// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
	LogoImageBlob pulumi.StringPtrInput `pulumi:"logoImageBlob"`
	// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. `ARM`, `ARM 64`, `x86`, `x86-64`
	OperatingSystems pulumi.StringArrayInput `pulumi:"operatingSystems"`
	// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
	UsageText pulumi.StringPtrInput `pulumi:"usageText"`
}

func (RepositoryCatalogDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryCatalogData)(nil)).Elem()
}

func (i RepositoryCatalogDataArgs) ToRepositoryCatalogDataOutput() RepositoryCatalogDataOutput {
	return i.ToRepositoryCatalogDataOutputWithContext(context.Background())
}

func (i RepositoryCatalogDataArgs) ToRepositoryCatalogDataOutputWithContext(ctx context.Context) RepositoryCatalogDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCatalogDataOutput)
}

func (i RepositoryCatalogDataArgs) ToRepositoryCatalogDataPtrOutput() RepositoryCatalogDataPtrOutput {
	return i.ToRepositoryCatalogDataPtrOutputWithContext(context.Background())
}

func (i RepositoryCatalogDataArgs) ToRepositoryCatalogDataPtrOutputWithContext(ctx context.Context) RepositoryCatalogDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCatalogDataOutput).ToRepositoryCatalogDataPtrOutputWithContext(ctx)
}

// RepositoryCatalogDataPtrInput is an input type that accepts RepositoryCatalogDataArgs, RepositoryCatalogDataPtr and RepositoryCatalogDataPtrOutput values.
// You can construct a concrete instance of `RepositoryCatalogDataPtrInput` via:
//
//          RepositoryCatalogDataArgs{...}
//
//  or:
//
//          nil
type RepositoryCatalogDataPtrInput interface {
	pulumi.Input

	ToRepositoryCatalogDataPtrOutput() RepositoryCatalogDataPtrOutput
	ToRepositoryCatalogDataPtrOutputWithContext(context.Context) RepositoryCatalogDataPtrOutput
}

type repositoryCatalogDataPtrType RepositoryCatalogDataArgs

func RepositoryCatalogDataPtr(v *RepositoryCatalogDataArgs) RepositoryCatalogDataPtrInput {
	return (*repositoryCatalogDataPtrType)(v)
}

func (*repositoryCatalogDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCatalogData)(nil)).Elem()
}

func (i *repositoryCatalogDataPtrType) ToRepositoryCatalogDataPtrOutput() RepositoryCatalogDataPtrOutput {
	return i.ToRepositoryCatalogDataPtrOutputWithContext(context.Background())
}

func (i *repositoryCatalogDataPtrType) ToRepositoryCatalogDataPtrOutputWithContext(ctx context.Context) RepositoryCatalogDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryCatalogDataPtrOutput)
}

type RepositoryCatalogDataOutput struct{ *pulumi.OutputState }

func (RepositoryCatalogDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryCatalogData)(nil)).Elem()
}

func (o RepositoryCatalogDataOutput) ToRepositoryCatalogDataOutput() RepositoryCatalogDataOutput {
	return o
}

func (o RepositoryCatalogDataOutput) ToRepositoryCatalogDataOutputWithContext(ctx context.Context) RepositoryCatalogDataOutput {
	return o
}

func (o RepositoryCatalogDataOutput) ToRepositoryCatalogDataPtrOutput() RepositoryCatalogDataPtrOutput {
	return o.ToRepositoryCatalogDataPtrOutputWithContext(context.Background())
}

func (o RepositoryCatalogDataOutput) ToRepositoryCatalogDataPtrOutputWithContext(ctx context.Context) RepositoryCatalogDataPtrOutput {
	return o.ApplyT(func(v RepositoryCatalogData) *RepositoryCatalogData {
		return &v
	}).(RepositoryCatalogDataPtrOutput)
}

// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
func (o RepositoryCatalogDataOutput) AboutText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryCatalogData) *string { return v.AboutText }).(pulumi.StringPtrOutput)
}

// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
func (o RepositoryCatalogDataOutput) Architectures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepositoryCatalogData) []string { return v.Architectures }).(pulumi.StringArrayOutput)
}

// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
func (o RepositoryCatalogDataOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryCatalogData) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
func (o RepositoryCatalogDataOutput) LogoImageBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryCatalogData) *string { return v.LogoImageBlob }).(pulumi.StringPtrOutput)
}

// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. `ARM`, `ARM 64`, `x86`, `x86-64`
func (o RepositoryCatalogDataOutput) OperatingSystems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepositoryCatalogData) []string { return v.OperatingSystems }).(pulumi.StringArrayOutput)
}

// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
func (o RepositoryCatalogDataOutput) UsageText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryCatalogData) *string { return v.UsageText }).(pulumi.StringPtrOutput)
}

type RepositoryCatalogDataPtrOutput struct{ *pulumi.OutputState }

func (RepositoryCatalogDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryCatalogData)(nil)).Elem()
}

func (o RepositoryCatalogDataPtrOutput) ToRepositoryCatalogDataPtrOutput() RepositoryCatalogDataPtrOutput {
	return o
}

func (o RepositoryCatalogDataPtrOutput) ToRepositoryCatalogDataPtrOutputWithContext(ctx context.Context) RepositoryCatalogDataPtrOutput {
	return o
}

func (o RepositoryCatalogDataPtrOutput) Elem() RepositoryCatalogDataOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) RepositoryCatalogData { return *v }).(RepositoryCatalogDataOutput)
}

// A detailed description of the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The text must be in markdown format.
func (o RepositoryCatalogDataPtrOutput) AboutText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) *string {
		if v == nil {
			return nil
		}
		return v.AboutText
	}).(pulumi.StringPtrOutput)
}

// The system architecture that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported architectures will appear as badges on the repository and are used as search filters: `Linux`, `Windows`
func (o RepositoryCatalogDataPtrOutput) Architectures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) []string {
		if v == nil {
			return nil
		}
		return v.Architectures
	}).(pulumi.StringArrayOutput)
}

// A short description of the contents of the repository. This text appears in both the image details and also when searching for repositories on the Amazon ECR Public Gallery.
func (o RepositoryCatalogDataPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The base64-encoded repository logo payload. (Only visible for verified accounts) Note that drift detection is disabled for this attribute.
func (o RepositoryCatalogDataPtrOutput) LogoImageBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) *string {
		if v == nil {
			return nil
		}
		return v.LogoImageBlob
	}).(pulumi.StringPtrOutput)
}

// The operating systems that the images in the repository are compatible with. On the Amazon ECR Public Gallery, the following supported operating systems will appear as badges on the repository and are used as search filters. `ARM`, `ARM 64`, `x86`, `x86-64`
func (o RepositoryCatalogDataPtrOutput) OperatingSystems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) []string {
		if v == nil {
			return nil
		}
		return v.OperatingSystems
	}).(pulumi.StringArrayOutput)
}

// Detailed information on how to use the contents of the repository. It is publicly visible in the Amazon ECR Public Gallery. The usage text provides context, support information, and additional usage details for users of the repository. The text must be in markdown format.
func (o RepositoryCatalogDataPtrOutput) UsageText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryCatalogData) *string {
		if v == nil {
			return nil
		}
		return v.UsageText
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryCatalogDataOutput{})
	pulumi.RegisterOutputType(RepositoryCatalogDataPtrOutput{})
}
