// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecrpublic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Public Elastic Container Registry Repository.
//
// > **NOTE:** This resource can only be used with `us-east-1` region.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/base64"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecrpublic"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func filebase64OrPanic(path string) pulumi.StringPtrInput {
// 	if fileData, err := ioutil.ReadFile(path); err == nil {
// 		return pulumi.String(base64.StdEncoding.EncodeToString(fileData[:]))
// 	} else {
// 		panic(err.Error())
// 	}
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "usEast1", &providers.awsArgs{
// 			Region: "us-east-1",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecrpublic.NewRepository(ctx, "foo", &ecrpublic.RepositoryArgs{
// 			RepositoryName: pulumi.String("bar"),
// 			CatalogData: &ecrpublic.RepositoryCatalogDataArgs{
// 				AboutText: pulumi.String("About Text"),
// 				Architectures: pulumi.StringArray{
// 					pulumi.String("ARM"),
// 				},
// 				Description:   pulumi.String("Description"),
// 				LogoImageBlob: filebase64OrPanic(image.Png),
// 				OperatingSystems: pulumi.StringArray{
// 					pulumi.String("Linux"),
// 				},
// 				UsageText: pulumi.String("Usage Text"),
// 			},
// 		}, pulumi.Provider(aws.Us_east_1))
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
// ECR Public Repositories can be imported using the `repository_name`, e.g.,
//
// ```sh
//  $ pulumi import aws:ecrpublic/repository:Repository example example
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Full ARN of the repository.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Catalog data configuration for the repository. See below for schema.
	CatalogData  RepositoryCatalogDataPtrOutput `pulumi:"catalogData"`
	ForceDestroy pulumi.BoolPtrOutput           `pulumi:"forceDestroy"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Name of the repository.
	RepositoryName pulumi.StringOutput `pulumi:"repositoryName"`
	// The URI of the repository.
	RepositoryUri pulumi.StringOutput `pulumi:"repositoryUri"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryName == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryName'")
	}
	var resource Repository
	err := ctx.RegisterResource("aws:ecrpublic/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws:ecrpublic/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Full ARN of the repository.
	Arn *string `pulumi:"arn"`
	// Catalog data configuration for the repository. See below for schema.
	CatalogData  *RepositoryCatalogData `pulumi:"catalogData"`
	ForceDestroy *bool                  `pulumi:"forceDestroy"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// Name of the repository.
	RepositoryName *string `pulumi:"repositoryName"`
	// The URI of the repository.
	RepositoryUri *string `pulumi:"repositoryUri"`
}

type RepositoryState struct {
	// Full ARN of the repository.
	Arn pulumi.StringPtrInput
	// Catalog data configuration for the repository. See below for schema.
	CatalogData  RepositoryCatalogDataPtrInput
	ForceDestroy pulumi.BoolPtrInput
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// Name of the repository.
	RepositoryName pulumi.StringPtrInput
	// The URI of the repository.
	RepositoryUri pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Catalog data configuration for the repository. See below for schema.
	CatalogData  *RepositoryCatalogData `pulumi:"catalogData"`
	ForceDestroy *bool                  `pulumi:"forceDestroy"`
	// Name of the repository.
	RepositoryName string `pulumi:"repositoryName"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Catalog data configuration for the repository. See below for schema.
	CatalogData  RepositoryCatalogDataPtrInput
	ForceDestroy pulumi.BoolPtrInput
	// Name of the repository.
	RepositoryName pulumi.StringInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

func (i *Repository) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPtrOutput)
}

type RepositoryPtrInput interface {
	pulumi.Input

	ToRepositoryPtrOutput() RepositoryPtrOutput
	ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput
}

type repositoryPtrType RepositoryArgs

func (*repositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil))
}

func (i *repositoryPtrType) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i *repositoryPtrType) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPtrOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//          RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//          RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o.ToRepositoryPtrOutputWithContext(context.Background())
}

func (o RepositoryOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Repository) *Repository {
		return &v
	}).(RepositoryPtrOutput)
}

type RepositoryPtrOutput struct{ *pulumi.OutputState }

func (RepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil))
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o
}

func (o RepositoryPtrOutput) Elem() RepositoryOutput {
	return o.ApplyT(func(v *Repository) Repository {
		if v != nil {
			return *v
		}
		var ret Repository
		return ret
	}).(RepositoryOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Repository)(nil))
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Repository {
		return vs[0].([]Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Repository)(nil))
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Repository {
		return vs[0].(map[string]Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryPtrInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryPtrOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
