// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS App Mesh virtual service resource.
//
// ## Example Usage
// ### Virtual Node Provider
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualServiceSpecArgs{
// 				Provider: &appmesh.VirtualServiceSpecProviderArgs{
// 					VirtualNode: &appmesh.VirtualServiceSpecProviderVirtualNodeArgs{
// 						VirtualNodeName: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
// 					},
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
// ### Virtual Router Provider
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualServiceSpecArgs{
// 				Provider: &appmesh.VirtualServiceSpecProviderArgs{
// 					VirtualRouter: &appmesh.VirtualServiceSpecProviderVirtualRouterArgs{
// 						VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
// 					},
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
// App Mesh virtual services can be imported using `mesh_name` together with the virtual service's `name`, e.g.
//
// ```sh
//  $ pulumi import aws:appmesh/virtualService:VirtualService servicea simpleapp/servicea.simpleapp.local
// ```
//
//  [1]/docs/providers/aws/index.html
type VirtualService struct {
	pulumi.CustomResourceState

	// The ARN of the virtual service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource VirtualService
	err := ctx.RegisterResource("aws:appmesh/virtualService:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("aws:appmesh/virtualService:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
	// The ARN of the virtual service.
	Arn *string `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The virtual service specification to apply.
	Spec *VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VirtualServiceState struct {
	// The ARN of the virtual service.
	Arn pulumi.StringPtrInput
	// The creation date of the virtual service.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}

type VirtualServiceInput interface {
	pulumi.Input

	ToVirtualServiceOutput() VirtualServiceOutput
	ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput
}

func (*VirtualService) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualService)(nil))
}

func (i *VirtualService) ToVirtualServiceOutput() VirtualServiceOutput {
	return i.ToVirtualServiceOutputWithContext(context.Background())
}

func (i *VirtualService) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceOutput)
}

func (i *VirtualService) ToVirtualServicePtrOutput() VirtualServicePtrOutput {
	return i.ToVirtualServicePtrOutputWithContext(context.Background())
}

func (i *VirtualService) ToVirtualServicePtrOutputWithContext(ctx context.Context) VirtualServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServicePtrOutput)
}

type VirtualServicePtrInput interface {
	pulumi.Input

	ToVirtualServicePtrOutput() VirtualServicePtrOutput
	ToVirtualServicePtrOutputWithContext(ctx context.Context) VirtualServicePtrOutput
}

type virtualServicePtrType VirtualServiceArgs

func (*virtualServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil))
}

func (i *virtualServicePtrType) ToVirtualServicePtrOutput() VirtualServicePtrOutput {
	return i.ToVirtualServicePtrOutputWithContext(context.Background())
}

func (i *virtualServicePtrType) ToVirtualServicePtrOutputWithContext(ctx context.Context) VirtualServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServicePtrOutput)
}

// VirtualServiceArrayInput is an input type that accepts VirtualServiceArray and VirtualServiceArrayOutput values.
// You can construct a concrete instance of `VirtualServiceArrayInput` via:
//
//          VirtualServiceArray{ VirtualServiceArgs{...} }
type VirtualServiceArrayInput interface {
	pulumi.Input

	ToVirtualServiceArrayOutput() VirtualServiceArrayOutput
	ToVirtualServiceArrayOutputWithContext(context.Context) VirtualServiceArrayOutput
}

type VirtualServiceArray []VirtualServiceInput

func (VirtualServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualService)(nil)).Elem()
}

func (i VirtualServiceArray) ToVirtualServiceArrayOutput() VirtualServiceArrayOutput {
	return i.ToVirtualServiceArrayOutputWithContext(context.Background())
}

func (i VirtualServiceArray) ToVirtualServiceArrayOutputWithContext(ctx context.Context) VirtualServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceArrayOutput)
}

// VirtualServiceMapInput is an input type that accepts VirtualServiceMap and VirtualServiceMapOutput values.
// You can construct a concrete instance of `VirtualServiceMapInput` via:
//
//          VirtualServiceMap{ "key": VirtualServiceArgs{...} }
type VirtualServiceMapInput interface {
	pulumi.Input

	ToVirtualServiceMapOutput() VirtualServiceMapOutput
	ToVirtualServiceMapOutputWithContext(context.Context) VirtualServiceMapOutput
}

type VirtualServiceMap map[string]VirtualServiceInput

func (VirtualServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualService)(nil)).Elem()
}

func (i VirtualServiceMap) ToVirtualServiceMapOutput() VirtualServiceMapOutput {
	return i.ToVirtualServiceMapOutputWithContext(context.Background())
}

func (i VirtualServiceMap) ToVirtualServiceMapOutputWithContext(ctx context.Context) VirtualServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceMapOutput)
}

type VirtualServiceOutput struct{ *pulumi.OutputState }

func (VirtualServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualService)(nil))
}

func (o VirtualServiceOutput) ToVirtualServiceOutput() VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServicePtrOutput() VirtualServicePtrOutput {
	return o.ToVirtualServicePtrOutputWithContext(context.Background())
}

func (o VirtualServiceOutput) ToVirtualServicePtrOutputWithContext(ctx context.Context) VirtualServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualService) *VirtualService {
		return &v
	}).(VirtualServicePtrOutput)
}

type VirtualServicePtrOutput struct{ *pulumi.OutputState }

func (VirtualServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil))
}

func (o VirtualServicePtrOutput) ToVirtualServicePtrOutput() VirtualServicePtrOutput {
	return o
}

func (o VirtualServicePtrOutput) ToVirtualServicePtrOutputWithContext(ctx context.Context) VirtualServicePtrOutput {
	return o
}

func (o VirtualServicePtrOutput) Elem() VirtualServiceOutput {
	return o.ApplyT(func(v *VirtualService) VirtualService {
		if v != nil {
			return *v
		}
		var ret VirtualService
		return ret
	}).(VirtualServiceOutput)
}

type VirtualServiceArrayOutput struct{ *pulumi.OutputState }

func (VirtualServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualService)(nil))
}

func (o VirtualServiceArrayOutput) ToVirtualServiceArrayOutput() VirtualServiceArrayOutput {
	return o
}

func (o VirtualServiceArrayOutput) ToVirtualServiceArrayOutputWithContext(ctx context.Context) VirtualServiceArrayOutput {
	return o
}

func (o VirtualServiceArrayOutput) Index(i pulumi.IntInput) VirtualServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualService {
		return vs[0].([]VirtualService)[vs[1].(int)]
	}).(VirtualServiceOutput)
}

type VirtualServiceMapOutput struct{ *pulumi.OutputState }

func (VirtualServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualService)(nil))
}

func (o VirtualServiceMapOutput) ToVirtualServiceMapOutput() VirtualServiceMapOutput {
	return o
}

func (o VirtualServiceMapOutput) ToVirtualServiceMapOutputWithContext(ctx context.Context) VirtualServiceMapOutput {
	return o
}

func (o VirtualServiceMapOutput) MapIndex(k pulumi.StringInput) VirtualServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualService {
		return vs[0].(map[string]VirtualService)[vs[1].(string)]
	}).(VirtualServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualServiceOutput{})
	pulumi.RegisterOutputType(VirtualServicePtrOutput{})
	pulumi.RegisterOutputType(VirtualServiceArrayOutput{})
	pulumi.RegisterOutputType(VirtualServiceMapOutput{})
}
