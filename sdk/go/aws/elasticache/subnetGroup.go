// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ElastiCache Subnet Group resource.
//
// > **NOTE:** ElastiCache Subnet Groups are only for use when working with an
// ElastiCache cluster **inside** of a VPC. If you are on EC2 Classic, see the
// ElastiCache Security Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("tf-test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
// 			VpcId:            fooVpc.ID(),
// 			CidrBlock:        pulumi.String("10.0.0.0/24"),
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("tf-test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticache.NewSubnetGroup(ctx, "bar", &elasticache.SubnetGroupArgs{
// 			SubnetIds: pulumi.StringArray{
// 				fooSubnet.ID(),
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
// ElastiCache Subnet Groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:elasticache/subnetGroup:SubnetGroup bar tf-test-cache-subnet
// ```
type SubnetGroup struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// Name for the cache subnet group. Elasticache converts this name to lowercase.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:elasticache/subnetGroup:SubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetGroupState, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	var resource SubnetGroup
	err := ctx.ReadResource("aws:elasticache/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	Arn *string `pulumi:"arn"`
	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Name for the cache subnet group. Elasticache converts this name to lowercase.
	Name *string `pulumi:"name"`
	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type SubnetGroupState struct {
	Arn pulumi.StringPtrInput
	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Name for the cache subnet group. Elasticache converts this name to lowercase.
	Name pulumi.StringPtrInput
	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds pulumi.StringArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Name for the cache subnet group. Elasticache converts this name to lowercase.
	Name *string `pulumi:"name"`
	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// Description for the cache subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Name for the cache subnet group. Elasticache converts this name to lowercase.
	Name pulumi.StringPtrInput
	// List of VPC Subnet IDs for the cache subnet group
	SubnetIds pulumi.StringArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}

type SubnetGroupInput interface {
	pulumi.Input

	ToSubnetGroupOutput() SubnetGroupOutput
	ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput
}

func (*SubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroup)(nil))
}

func (i *SubnetGroup) ToSubnetGroupOutput() SubnetGroupOutput {
	return i.ToSubnetGroupOutputWithContext(context.Background())
}

func (i *SubnetGroup) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupOutput)
}

func (i *SubnetGroup) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return i.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *SubnetGroup) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupPtrOutput)
}

type SubnetGroupPtrInput interface {
	pulumi.Input

	ToSubnetGroupPtrOutput() SubnetGroupPtrOutput
	ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput
}

type subnetGroupPtrType SubnetGroupArgs

func (*subnetGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetGroup)(nil))
}

func (i *subnetGroupPtrType) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return i.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (i *subnetGroupPtrType) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupPtrOutput)
}

// SubnetGroupArrayInput is an input type that accepts SubnetGroupArray and SubnetGroupArrayOutput values.
// You can construct a concrete instance of `SubnetGroupArrayInput` via:
//
//          SubnetGroupArray{ SubnetGroupArgs{...} }
type SubnetGroupArrayInput interface {
	pulumi.Input

	ToSubnetGroupArrayOutput() SubnetGroupArrayOutput
	ToSubnetGroupArrayOutputWithContext(context.Context) SubnetGroupArrayOutput
}

type SubnetGroupArray []SubnetGroupInput

func (SubnetGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetGroup)(nil)).Elem()
}

func (i SubnetGroupArray) ToSubnetGroupArrayOutput() SubnetGroupArrayOutput {
	return i.ToSubnetGroupArrayOutputWithContext(context.Background())
}

func (i SubnetGroupArray) ToSubnetGroupArrayOutputWithContext(ctx context.Context) SubnetGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupArrayOutput)
}

// SubnetGroupMapInput is an input type that accepts SubnetGroupMap and SubnetGroupMapOutput values.
// You can construct a concrete instance of `SubnetGroupMapInput` via:
//
//          SubnetGroupMap{ "key": SubnetGroupArgs{...} }
type SubnetGroupMapInput interface {
	pulumi.Input

	ToSubnetGroupMapOutput() SubnetGroupMapOutput
	ToSubnetGroupMapOutputWithContext(context.Context) SubnetGroupMapOutput
}

type SubnetGroupMap map[string]SubnetGroupInput

func (SubnetGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetGroup)(nil)).Elem()
}

func (i SubnetGroupMap) ToSubnetGroupMapOutput() SubnetGroupMapOutput {
	return i.ToSubnetGroupMapOutputWithContext(context.Background())
}

func (i SubnetGroupMap) ToSubnetGroupMapOutputWithContext(ctx context.Context) SubnetGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupMapOutput)
}

type SubnetGroupOutput struct {
	*pulumi.OutputState
}

func (SubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroup)(nil))
}

func (o SubnetGroupOutput) ToSubnetGroupOutput() SubnetGroupOutput {
	return o
}

func (o SubnetGroupOutput) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return o
}

func (o SubnetGroupOutput) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return o.ToSubnetGroupPtrOutputWithContext(context.Background())
}

func (o SubnetGroupOutput) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return o.ApplyT(func(v SubnetGroup) *SubnetGroup {
		return &v
	}).(SubnetGroupPtrOutput)
}

type SubnetGroupPtrOutput struct {
	*pulumi.OutputState
}

func (SubnetGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetGroup)(nil))
}

func (o SubnetGroupPtrOutput) ToSubnetGroupPtrOutput() SubnetGroupPtrOutput {
	return o
}

func (o SubnetGroupPtrOutput) ToSubnetGroupPtrOutputWithContext(ctx context.Context) SubnetGroupPtrOutput {
	return o
}

type SubnetGroupArrayOutput struct{ *pulumi.OutputState }

func (SubnetGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetGroup)(nil))
}

func (o SubnetGroupArrayOutput) ToSubnetGroupArrayOutput() SubnetGroupArrayOutput {
	return o
}

func (o SubnetGroupArrayOutput) ToSubnetGroupArrayOutputWithContext(ctx context.Context) SubnetGroupArrayOutput {
	return o
}

func (o SubnetGroupArrayOutput) Index(i pulumi.IntInput) SubnetGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetGroup {
		return vs[0].([]SubnetGroup)[vs[1].(int)]
	}).(SubnetGroupOutput)
}

type SubnetGroupMapOutput struct{ *pulumi.OutputState }

func (SubnetGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SubnetGroup)(nil))
}

func (o SubnetGroupMapOutput) ToSubnetGroupMapOutput() SubnetGroupMapOutput {
	return o
}

func (o SubnetGroupMapOutput) ToSubnetGroupMapOutputWithContext(ctx context.Context) SubnetGroupMapOutput {
	return o
}

func (o SubnetGroupMapOutput) MapIndex(k pulumi.StringInput) SubnetGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SubnetGroup {
		return vs[0].(map[string]SubnetGroup)[vs[1].(string)]
	}).(SubnetGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetGroupOutput{})
	pulumi.RegisterOutputType(SubnetGroupPtrOutput{})
	pulumi.RegisterOutputType(SubnetGroupArrayOutput{})
	pulumi.RegisterOutputType(SubnetGroupMapOutput{})
}
