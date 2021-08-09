// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic Container Registry Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentRegion, err := aws.GetRegion(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		currentPartition, err := aws.GetPartition(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Sid":    "testpolicy",
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"AWS": fmt.Sprintf("%v%v%v%v%v", "arn:", currentPartition.Partition, ":iam::", currentCallerIdentity.AccountId, ":root"),
// 					},
// 					"Action": []string{
// 						"ecr:ReplicateImage",
// 					},
// 					"Resource": []string{
// 						fmt.Sprintf("%v%v%v%v%v%v%v", "arn:", currentPartition.Partition, ":ecr:", currentRegion.Name, ":", currentCallerIdentity.AccountId, ":repository/*"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err = ecr.NewRegistryPolicy(ctx, "example", &ecr.RegistryPolicyArgs{
// 			Policy: pulumi.String(json0),
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
// ECR Registry Policy can be imported using the registry id, e.g.
//
// ```sh
//  $ pulumi import aws:ecr/registryPolicy:RegistryPolicy example 123456789012
// ```
type RegistryPolicy struct {
	pulumi.CustomResourceState

	Policy pulumi.StringOutput `pulumi:"policy"`
	// The registry ID where the registry was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
}

// NewRegistryPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegistryPolicy(ctx *pulumi.Context,
	name string, args *RegistryPolicyArgs, opts ...pulumi.ResourceOption) (*RegistryPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource RegistryPolicy
	err := ctx.RegisterResource("aws:ecr/registryPolicy:RegistryPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryPolicy gets an existing RegistryPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryPolicyState, opts ...pulumi.ResourceOption) (*RegistryPolicy, error) {
	var resource RegistryPolicy
	err := ctx.ReadResource("aws:ecr/registryPolicy:RegistryPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryPolicy resources.
type registryPolicyState struct {
	Policy *string `pulumi:"policy"`
	// The registry ID where the registry was created.
	RegistryId *string `pulumi:"registryId"`
}

type RegistryPolicyState struct {
	Policy pulumi.StringPtrInput
	// The registry ID where the registry was created.
	RegistryId pulumi.StringPtrInput
}

func (RegistryPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryPolicyState)(nil)).Elem()
}

type registryPolicyArgs struct {
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a RegistryPolicy resource.
type RegistryPolicyArgs struct {
	Policy pulumi.StringInput
}

func (RegistryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryPolicyArgs)(nil)).Elem()
}

type RegistryPolicyInput interface {
	pulumi.Input

	ToRegistryPolicyOutput() RegistryPolicyOutput
	ToRegistryPolicyOutputWithContext(ctx context.Context) RegistryPolicyOutput
}

func (*RegistryPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPolicy)(nil))
}

func (i *RegistryPolicy) ToRegistryPolicyOutput() RegistryPolicyOutput {
	return i.ToRegistryPolicyOutputWithContext(context.Background())
}

func (i *RegistryPolicy) ToRegistryPolicyOutputWithContext(ctx context.Context) RegistryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPolicyOutput)
}

func (i *RegistryPolicy) ToRegistryPolicyPtrOutput() RegistryPolicyPtrOutput {
	return i.ToRegistryPolicyPtrOutputWithContext(context.Background())
}

func (i *RegistryPolicy) ToRegistryPolicyPtrOutputWithContext(ctx context.Context) RegistryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPolicyPtrOutput)
}

type RegistryPolicyPtrInput interface {
	pulumi.Input

	ToRegistryPolicyPtrOutput() RegistryPolicyPtrOutput
	ToRegistryPolicyPtrOutputWithContext(ctx context.Context) RegistryPolicyPtrOutput
}

type registryPolicyPtrType RegistryPolicyArgs

func (*registryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryPolicy)(nil))
}

func (i *registryPolicyPtrType) ToRegistryPolicyPtrOutput() RegistryPolicyPtrOutput {
	return i.ToRegistryPolicyPtrOutputWithContext(context.Background())
}

func (i *registryPolicyPtrType) ToRegistryPolicyPtrOutputWithContext(ctx context.Context) RegistryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPolicyPtrOutput)
}

// RegistryPolicyArrayInput is an input type that accepts RegistryPolicyArray and RegistryPolicyArrayOutput values.
// You can construct a concrete instance of `RegistryPolicyArrayInput` via:
//
//          RegistryPolicyArray{ RegistryPolicyArgs{...} }
type RegistryPolicyArrayInput interface {
	pulumi.Input

	ToRegistryPolicyArrayOutput() RegistryPolicyArrayOutput
	ToRegistryPolicyArrayOutputWithContext(context.Context) RegistryPolicyArrayOutput
}

type RegistryPolicyArray []RegistryPolicyInput

func (RegistryPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryPolicy)(nil)).Elem()
}

func (i RegistryPolicyArray) ToRegistryPolicyArrayOutput() RegistryPolicyArrayOutput {
	return i.ToRegistryPolicyArrayOutputWithContext(context.Background())
}

func (i RegistryPolicyArray) ToRegistryPolicyArrayOutputWithContext(ctx context.Context) RegistryPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPolicyArrayOutput)
}

// RegistryPolicyMapInput is an input type that accepts RegistryPolicyMap and RegistryPolicyMapOutput values.
// You can construct a concrete instance of `RegistryPolicyMapInput` via:
//
//          RegistryPolicyMap{ "key": RegistryPolicyArgs{...} }
type RegistryPolicyMapInput interface {
	pulumi.Input

	ToRegistryPolicyMapOutput() RegistryPolicyMapOutput
	ToRegistryPolicyMapOutputWithContext(context.Context) RegistryPolicyMapOutput
}

type RegistryPolicyMap map[string]RegistryPolicyInput

func (RegistryPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryPolicy)(nil)).Elem()
}

func (i RegistryPolicyMap) ToRegistryPolicyMapOutput() RegistryPolicyMapOutput {
	return i.ToRegistryPolicyMapOutputWithContext(context.Background())
}

func (i RegistryPolicyMap) ToRegistryPolicyMapOutputWithContext(ctx context.Context) RegistryPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPolicyMapOutput)
}

type RegistryPolicyOutput struct{ *pulumi.OutputState }

func (RegistryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPolicy)(nil))
}

func (o RegistryPolicyOutput) ToRegistryPolicyOutput() RegistryPolicyOutput {
	return o
}

func (o RegistryPolicyOutput) ToRegistryPolicyOutputWithContext(ctx context.Context) RegistryPolicyOutput {
	return o
}

func (o RegistryPolicyOutput) ToRegistryPolicyPtrOutput() RegistryPolicyPtrOutput {
	return o.ToRegistryPolicyPtrOutputWithContext(context.Background())
}

func (o RegistryPolicyOutput) ToRegistryPolicyPtrOutputWithContext(ctx context.Context) RegistryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistryPolicy) *RegistryPolicy {
		return &v
	}).(RegistryPolicyPtrOutput)
}

type RegistryPolicyPtrOutput struct{ *pulumi.OutputState }

func (RegistryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryPolicy)(nil))
}

func (o RegistryPolicyPtrOutput) ToRegistryPolicyPtrOutput() RegistryPolicyPtrOutput {
	return o
}

func (o RegistryPolicyPtrOutput) ToRegistryPolicyPtrOutputWithContext(ctx context.Context) RegistryPolicyPtrOutput {
	return o
}

func (o RegistryPolicyPtrOutput) Elem() RegistryPolicyOutput {
	return o.ApplyT(func(v *RegistryPolicy) RegistryPolicy {
		if v != nil {
			return *v
		}
		var ret RegistryPolicy
		return ret
	}).(RegistryPolicyOutput)
}

type RegistryPolicyArrayOutput struct{ *pulumi.OutputState }

func (RegistryPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPolicy)(nil))
}

func (o RegistryPolicyArrayOutput) ToRegistryPolicyArrayOutput() RegistryPolicyArrayOutput {
	return o
}

func (o RegistryPolicyArrayOutput) ToRegistryPolicyArrayOutputWithContext(ctx context.Context) RegistryPolicyArrayOutput {
	return o
}

func (o RegistryPolicyArrayOutput) Index(i pulumi.IntInput) RegistryPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryPolicy {
		return vs[0].([]RegistryPolicy)[vs[1].(int)]
	}).(RegistryPolicyOutput)
}

type RegistryPolicyMapOutput struct{ *pulumi.OutputState }

func (RegistryPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RegistryPolicy)(nil))
}

func (o RegistryPolicyMapOutput) ToRegistryPolicyMapOutput() RegistryPolicyMapOutput {
	return o
}

func (o RegistryPolicyMapOutput) ToRegistryPolicyMapOutputWithContext(ctx context.Context) RegistryPolicyMapOutput {
	return o
}

func (o RegistryPolicyMapOutput) MapIndex(k pulumi.StringInput) RegistryPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RegistryPolicy {
		return vs[0].(map[string]RegistryPolicy)[vs[1].(string)]
	}).(RegistryPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryPolicyOutput{})
	pulumi.RegisterOutputType(RegistryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RegistryPolicyArrayOutput{})
	pulumi.RegisterOutputType(RegistryPolicyMapOutput{})
}
