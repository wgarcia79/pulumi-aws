// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables AWS Shield Advanced for a specific AWS resource.
// The resource can be an Amazon CloudFront distribution, Elastic Load Balancing load balancer, AWS Global Accelerator accelerator, Elastic IP Address, or an Amazon Route 53 hosted zone.
//
// ## Example Usage
// ### Create protection
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/shield"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aws.GetAvailabilityZones(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentRegion, err := aws.GetRegion(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleEip, err := ec2.NewEip(ctx, "exampleEip", &ec2.EipArgs{
// 			Vpc: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = shield.NewProtection(ctx, "exampleProtection", &shield.ProtectionArgs{
// 			ResourceArn: exampleEip.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:ec2:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":eip-allocation/", id), nil
// 			}).(pulumi.StringOutput),
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("Dev"),
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
// Shield protection resources can be imported by specifying their ID e.g.
//
// ```sh
//  $ pulumi import aws:shield/protection:Protection example ff9592dc-22f3-4e88-afa1-7b29fde9669a
// ```
type Protection struct {
	pulumi.CustomResourceState

	// The ARN of the Protection.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly name for the Protection you are creating.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewProtection registers a new resource with the given unique name, arguments, and options.
func NewProtection(ctx *pulumi.Context,
	name string, args *ProtectionArgs, opts ...pulumi.ResourceOption) (*Protection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	var resource Protection
	err := ctx.RegisterResource("aws:shield/protection:Protection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtection gets an existing Protection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionState, opts ...pulumi.ResourceOption) (*Protection, error) {
	var resource Protection
	err := ctx.ReadResource("aws:shield/protection:Protection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Protection resources.
type protectionState struct {
	// The ARN of the Protection.
	Arn *string `pulumi:"arn"`
	// A friendly name for the Protection you are creating.
	Name *string `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn *string `pulumi:"resourceArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ProtectionState struct {
	// The ARN of the Protection.
	Arn pulumi.StringPtrInput
	// A friendly name for the Protection you are creating.
	Name pulumi.StringPtrInput
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ProtectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionState)(nil)).Elem()
}

type protectionArgs struct {
	// A friendly name for the Protection you are creating.
	Name *string `pulumi:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn string `pulumi:"resourceArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Protection resource.
type ProtectionArgs struct {
	// A friendly name for the Protection you are creating.
	Name pulumi.StringPtrInput
	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ProtectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionArgs)(nil)).Elem()
}

type ProtectionInput interface {
	pulumi.Input

	ToProtectionOutput() ProtectionOutput
	ToProtectionOutputWithContext(ctx context.Context) ProtectionOutput
}

func (*Protection) ElementType() reflect.Type {
	return reflect.TypeOf((*Protection)(nil))
}

func (i *Protection) ToProtectionOutput() ProtectionOutput {
	return i.ToProtectionOutputWithContext(context.Background())
}

func (i *Protection) ToProtectionOutputWithContext(ctx context.Context) ProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionOutput)
}

func (i *Protection) ToProtectionPtrOutput() ProtectionPtrOutput {
	return i.ToProtectionPtrOutputWithContext(context.Background())
}

func (i *Protection) ToProtectionPtrOutputWithContext(ctx context.Context) ProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionPtrOutput)
}

type ProtectionPtrInput interface {
	pulumi.Input

	ToProtectionPtrOutput() ProtectionPtrOutput
	ToProtectionPtrOutputWithContext(ctx context.Context) ProtectionPtrOutput
}

type protectionPtrType ProtectionArgs

func (*protectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Protection)(nil))
}

func (i *protectionPtrType) ToProtectionPtrOutput() ProtectionPtrOutput {
	return i.ToProtectionPtrOutputWithContext(context.Background())
}

func (i *protectionPtrType) ToProtectionPtrOutputWithContext(ctx context.Context) ProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionPtrOutput)
}

// ProtectionArrayInput is an input type that accepts ProtectionArray and ProtectionArrayOutput values.
// You can construct a concrete instance of `ProtectionArrayInput` via:
//
//          ProtectionArray{ ProtectionArgs{...} }
type ProtectionArrayInput interface {
	pulumi.Input

	ToProtectionArrayOutput() ProtectionArrayOutput
	ToProtectionArrayOutputWithContext(context.Context) ProtectionArrayOutput
}

type ProtectionArray []ProtectionInput

func (ProtectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Protection)(nil)).Elem()
}

func (i ProtectionArray) ToProtectionArrayOutput() ProtectionArrayOutput {
	return i.ToProtectionArrayOutputWithContext(context.Background())
}

func (i ProtectionArray) ToProtectionArrayOutputWithContext(ctx context.Context) ProtectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionArrayOutput)
}

// ProtectionMapInput is an input type that accepts ProtectionMap and ProtectionMapOutput values.
// You can construct a concrete instance of `ProtectionMapInput` via:
//
//          ProtectionMap{ "key": ProtectionArgs{...} }
type ProtectionMapInput interface {
	pulumi.Input

	ToProtectionMapOutput() ProtectionMapOutput
	ToProtectionMapOutputWithContext(context.Context) ProtectionMapOutput
}

type ProtectionMap map[string]ProtectionInput

func (ProtectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Protection)(nil)).Elem()
}

func (i ProtectionMap) ToProtectionMapOutput() ProtectionMapOutput {
	return i.ToProtectionMapOutputWithContext(context.Background())
}

func (i ProtectionMap) ToProtectionMapOutputWithContext(ctx context.Context) ProtectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionMapOutput)
}

type ProtectionOutput struct{ *pulumi.OutputState }

func (ProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protection)(nil))
}

func (o ProtectionOutput) ToProtectionOutput() ProtectionOutput {
	return o
}

func (o ProtectionOutput) ToProtectionOutputWithContext(ctx context.Context) ProtectionOutput {
	return o
}

func (o ProtectionOutput) ToProtectionPtrOutput() ProtectionPtrOutput {
	return o.ToProtectionPtrOutputWithContext(context.Background())
}

func (o ProtectionOutput) ToProtectionPtrOutputWithContext(ctx context.Context) ProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protection) *Protection {
		return &v
	}).(ProtectionPtrOutput)
}

type ProtectionPtrOutput struct{ *pulumi.OutputState }

func (ProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protection)(nil))
}

func (o ProtectionPtrOutput) ToProtectionPtrOutput() ProtectionPtrOutput {
	return o
}

func (o ProtectionPtrOutput) ToProtectionPtrOutputWithContext(ctx context.Context) ProtectionPtrOutput {
	return o
}

func (o ProtectionPtrOutput) Elem() ProtectionOutput {
	return o.ApplyT(func(v *Protection) Protection {
		if v != nil {
			return *v
		}
		var ret Protection
		return ret
	}).(ProtectionOutput)
}

type ProtectionArrayOutput struct{ *pulumi.OutputState }

func (ProtectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Protection)(nil))
}

func (o ProtectionArrayOutput) ToProtectionArrayOutput() ProtectionArrayOutput {
	return o
}

func (o ProtectionArrayOutput) ToProtectionArrayOutputWithContext(ctx context.Context) ProtectionArrayOutput {
	return o
}

func (o ProtectionArrayOutput) Index(i pulumi.IntInput) ProtectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Protection {
		return vs[0].([]Protection)[vs[1].(int)]
	}).(ProtectionOutput)
}

type ProtectionMapOutput struct{ *pulumi.OutputState }

func (ProtectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Protection)(nil))
}

func (o ProtectionMapOutput) ToProtectionMapOutput() ProtectionMapOutput {
	return o
}

func (o ProtectionMapOutput) ToProtectionMapOutputWithContext(ctx context.Context) ProtectionMapOutput {
	return o
}

func (o ProtectionMapOutput) MapIndex(k pulumi.StringInput) ProtectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Protection {
		return vs[0].(map[string]Protection)[vs[1].(string)]
	}).(ProtectionOutput)
}

func init() {
	pulumi.RegisterOutputType(ProtectionOutput{})
	pulumi.RegisterOutputType(ProtectionPtrOutput{})
	pulumi.RegisterOutputType(ProtectionArrayOutput{})
	pulumi.RegisterOutputType(ProtectionMapOutput{})
}
