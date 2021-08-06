// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EKS Fargate Profile.
//
// ## Example Usage
// ### Example IAM Role for EKS Fargate Profile
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": "sts:AssumeRole",
// 					"Effect": "Allow",
// 					"Principal": map[string]interface{}{
// 						"Service": "eks-fargate-pods.amazonaws.com",
// 					},
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "example_AmazonEKSFargatePodExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy"),
// 			Role:      example.Name,
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
// EKS Fargate Profiles can be imported using the `cluster_name` and `fargate_profile_name` separated by a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:eks/fargateProfile:FargateProfile my_fargate_profile my_cluster:my_fargate_profile
// ```
type FargateProfile struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringOutput `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringOutput `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayOutput `pulumi:"selectors"`
	// Status of the EKS Fargate Profile.
	Status pulumi.StringOutput `pulumi:"status"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	Tags      pulumi.StringMapOutput   `pulumi:"tags"`
	TagsAll   pulumi.StringMapOutput   `pulumi:"tagsAll"`
}

// NewFargateProfile registers a new resource with the given unique name, arguments, and options.
func NewFargateProfile(ctx *pulumi.Context,
	name string, args *FargateProfileArgs, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PodExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'PodExecutionRoleArn'")
	}
	if args.Selectors == nil {
		return nil, errors.New("invalid value for required argument 'Selectors'")
	}
	var resource FargateProfile
	err := ctx.RegisterResource("aws:eks/fargateProfile:FargateProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFargateProfile gets an existing FargateProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFargateProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FargateProfileState, opts ...pulumi.ResourceOption) (*FargateProfile, error) {
	var resource FargateProfile
	err := ctx.ReadResource("aws:eks/fargateProfile:FargateProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FargateProfile resources.
type fargateProfileState struct {
	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn *string `pulumi:"arn"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName *string `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName *string `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn *string `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors []FargateProfileSelector `pulumi:"selectors"`
	// Status of the EKS Fargate Profile.
	Status *string `pulumi:"status"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	TagsAll   map[string]string `pulumi:"tagsAll"`
}

type FargateProfileState struct {
	// Amazon Resource Name (ARN) of the EKS Fargate Profile.
	Arn pulumi.StringPtrInput
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringPtrInput
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringPtrInput
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayInput
	// Status of the EKS Fargate Profile.
	Status pulumi.StringPtrInput
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayInput
	Tags      pulumi.StringMapInput
	TagsAll   pulumi.StringMapInput
}

func (FargateProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileState)(nil)).Elem()
}

type fargateProfileArgs struct {
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName string `pulumi:"clusterName"`
	// Name of the EKS Fargate Profile.
	FargateProfileName *string `pulumi:"fargateProfileName"`
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn string `pulumi:"podExecutionRoleArn"`
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors []FargateProfileSelector `pulumi:"selectors"`
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	TagsAll   map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a FargateProfile resource.
type FargateProfileArgs struct {
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringInput
	// Name of the EKS Fargate Profile.
	FargateProfileName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Fargate Profile.
	PodExecutionRoleArn pulumi.StringInput
	// Configuration block(s) for selecting Kubernetes Pods to execute with this EKS Fargate Profile. Detailed below.
	Selectors FargateProfileSelectorArrayInput
	// Identifiers of private EC2 Subnets to associate with the EKS Fargate Profile. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
	SubnetIds pulumi.StringArrayInput
	Tags      pulumi.StringMapInput
	TagsAll   pulumi.StringMapInput
}

func (FargateProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fargateProfileArgs)(nil)).Elem()
}

type FargateProfileInput interface {
	pulumi.Input

	ToFargateProfileOutput() FargateProfileOutput
	ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput
}

func (*FargateProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateProfile)(nil))
}

func (i *FargateProfile) ToFargateProfileOutput() FargateProfileOutput {
	return i.ToFargateProfileOutputWithContext(context.Background())
}

func (i *FargateProfile) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfileOutput)
}

func (i *FargateProfile) ToFargateProfilePtrOutput() FargateProfilePtrOutput {
	return i.ToFargateProfilePtrOutputWithContext(context.Background())
}

func (i *FargateProfile) ToFargateProfilePtrOutputWithContext(ctx context.Context) FargateProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfilePtrOutput)
}

type FargateProfilePtrInput interface {
	pulumi.Input

	ToFargateProfilePtrOutput() FargateProfilePtrOutput
	ToFargateProfilePtrOutputWithContext(ctx context.Context) FargateProfilePtrOutput
}

type fargateProfilePtrType FargateProfileArgs

func (*fargateProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateProfile)(nil))
}

func (i *fargateProfilePtrType) ToFargateProfilePtrOutput() FargateProfilePtrOutput {
	return i.ToFargateProfilePtrOutputWithContext(context.Background())
}

func (i *fargateProfilePtrType) ToFargateProfilePtrOutputWithContext(ctx context.Context) FargateProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfilePtrOutput)
}

// FargateProfileArrayInput is an input type that accepts FargateProfileArray and FargateProfileArrayOutput values.
// You can construct a concrete instance of `FargateProfileArrayInput` via:
//
//          FargateProfileArray{ FargateProfileArgs{...} }
type FargateProfileArrayInput interface {
	pulumi.Input

	ToFargateProfileArrayOutput() FargateProfileArrayOutput
	ToFargateProfileArrayOutputWithContext(context.Context) FargateProfileArrayOutput
}

type FargateProfileArray []FargateProfileInput

func (FargateProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FargateProfile)(nil)).Elem()
}

func (i FargateProfileArray) ToFargateProfileArrayOutput() FargateProfileArrayOutput {
	return i.ToFargateProfileArrayOutputWithContext(context.Background())
}

func (i FargateProfileArray) ToFargateProfileArrayOutputWithContext(ctx context.Context) FargateProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfileArrayOutput)
}

// FargateProfileMapInput is an input type that accepts FargateProfileMap and FargateProfileMapOutput values.
// You can construct a concrete instance of `FargateProfileMapInput` via:
//
//          FargateProfileMap{ "key": FargateProfileArgs{...} }
type FargateProfileMapInput interface {
	pulumi.Input

	ToFargateProfileMapOutput() FargateProfileMapOutput
	ToFargateProfileMapOutputWithContext(context.Context) FargateProfileMapOutput
}

type FargateProfileMap map[string]FargateProfileInput

func (FargateProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FargateProfile)(nil)).Elem()
}

func (i FargateProfileMap) ToFargateProfileMapOutput() FargateProfileMapOutput {
	return i.ToFargateProfileMapOutputWithContext(context.Background())
}

func (i FargateProfileMap) ToFargateProfileMapOutputWithContext(ctx context.Context) FargateProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FargateProfileMapOutput)
}

type FargateProfileOutput struct {
	*pulumi.OutputState
}

func (FargateProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FargateProfile)(nil))
}

func (o FargateProfileOutput) ToFargateProfileOutput() FargateProfileOutput {
	return o
}

func (o FargateProfileOutput) ToFargateProfileOutputWithContext(ctx context.Context) FargateProfileOutput {
	return o
}

func (o FargateProfileOutput) ToFargateProfilePtrOutput() FargateProfilePtrOutput {
	return o.ToFargateProfilePtrOutputWithContext(context.Background())
}

func (o FargateProfileOutput) ToFargateProfilePtrOutputWithContext(ctx context.Context) FargateProfilePtrOutput {
	return o.ApplyT(func(v FargateProfile) *FargateProfile {
		return &v
	}).(FargateProfilePtrOutput)
}

type FargateProfilePtrOutput struct {
	*pulumi.OutputState
}

func (FargateProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FargateProfile)(nil))
}

func (o FargateProfilePtrOutput) ToFargateProfilePtrOutput() FargateProfilePtrOutput {
	return o
}

func (o FargateProfilePtrOutput) ToFargateProfilePtrOutputWithContext(ctx context.Context) FargateProfilePtrOutput {
	return o
}

type FargateProfileArrayOutput struct{ *pulumi.OutputState }

func (FargateProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FargateProfile)(nil))
}

func (o FargateProfileArrayOutput) ToFargateProfileArrayOutput() FargateProfileArrayOutput {
	return o
}

func (o FargateProfileArrayOutput) ToFargateProfileArrayOutputWithContext(ctx context.Context) FargateProfileArrayOutput {
	return o
}

func (o FargateProfileArrayOutput) Index(i pulumi.IntInput) FargateProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FargateProfile {
		return vs[0].([]FargateProfile)[vs[1].(int)]
	}).(FargateProfileOutput)
}

type FargateProfileMapOutput struct{ *pulumi.OutputState }

func (FargateProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FargateProfile)(nil))
}

func (o FargateProfileMapOutput) ToFargateProfileMapOutput() FargateProfileMapOutput {
	return o
}

func (o FargateProfileMapOutput) ToFargateProfileMapOutputWithContext(ctx context.Context) FargateProfileMapOutput {
	return o
}

func (o FargateProfileMapOutput) MapIndex(k pulumi.StringInput) FargateProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FargateProfile {
		return vs[0].(map[string]FargateProfile)[vs[1].(string)]
	}).(FargateProfileOutput)
}

func init() {
	pulumi.RegisterOutputType(FargateProfileOutput{})
	pulumi.RegisterOutputType(FargateProfilePtrOutput{})
	pulumi.RegisterOutputType(FargateProfileArrayOutput{})
	pulumi.RegisterOutputType(FargateProfileMapOutput{})
}
