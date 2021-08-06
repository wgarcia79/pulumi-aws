// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a RDS DB Cluster association with an IAM Role. Example use cases:
//
// * [Creating an IAM Role to Allow Amazon Aurora to Access AWS Services](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.Authorizing.IAM.CreateRole.html)
// * [Importing Amazon S3 Data into an RDS PostgreSQL DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewClusterRoleAssociation(ctx, "example", &rds.ClusterRoleAssociationArgs{
// 			DbClusterIdentifier: pulumi.Any(aws_rds_cluster.Example.Id),
// 			FeatureName:         pulumi.String("S3_INTEGRATION"),
// 			RoleArn:             pulumi.Any(aws_iam_role.Example.Id),
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
// `aws_rds_cluster_role_association` can be imported using the DB Cluster Identifier and IAM Role ARN separated by a comma (`,`), e.g.
//
// ```sh
//  $ pulumi import aws:rds/clusterRoleAssociation:ClusterRoleAssociation example my-db-cluster,arn:aws:iam::123456789012:role/my-role
// ```
type ClusterRoleAssociation struct {
	pulumi.CustomResourceState

	// DB Cluster Identifier to associate with the IAM Role.
	DbClusterIdentifier pulumi.StringOutput `pulumi:"dbClusterIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringOutput `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewClusterRoleAssociation registers a new resource with the given unique name, arguments, and options.
func NewClusterRoleAssociation(ctx *pulumi.Context,
	name string, args *ClusterRoleAssociationArgs, opts ...pulumi.ResourceOption) (*ClusterRoleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DbClusterIdentifier'")
	}
	if args.FeatureName == nil {
		return nil, errors.New("invalid value for required argument 'FeatureName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource ClusterRoleAssociation
	err := ctx.RegisterResource("aws:rds/clusterRoleAssociation:ClusterRoleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRoleAssociation gets an existing ClusterRoleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRoleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleAssociationState, opts ...pulumi.ResourceOption) (*ClusterRoleAssociation, error) {
	var resource ClusterRoleAssociation
	err := ctx.ReadResource("aws:rds/clusterRoleAssociation:ClusterRoleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRoleAssociation resources.
type clusterRoleAssociationState struct {
	// DB Cluster Identifier to associate with the IAM Role.
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName *string `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn *string `pulumi:"roleArn"`
}

type ClusterRoleAssociationState struct {
	// DB Cluster Identifier to associate with the IAM Role.
	DbClusterIdentifier pulumi.StringPtrInput
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn pulumi.StringPtrInput
}

func (ClusterRoleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleAssociationState)(nil)).Elem()
}

type clusterRoleAssociationArgs struct {
	// DB Cluster Identifier to associate with the IAM Role.
	DbClusterIdentifier string `pulumi:"dbClusterIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName string `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a ClusterRoleAssociation resource.
type ClusterRoleAssociationArgs struct {
	// DB Cluster Identifier to associate with the IAM Role.
	DbClusterIdentifier pulumi.StringInput
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringInput
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Cluster.
	RoleArn pulumi.StringInput
}

func (ClusterRoleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleAssociationArgs)(nil)).Elem()
}

type ClusterRoleAssociationInput interface {
	pulumi.Input

	ToClusterRoleAssociationOutput() ClusterRoleAssociationOutput
	ToClusterRoleAssociationOutputWithContext(ctx context.Context) ClusterRoleAssociationOutput
}

func (*ClusterRoleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleAssociation)(nil))
}

func (i *ClusterRoleAssociation) ToClusterRoleAssociationOutput() ClusterRoleAssociationOutput {
	return i.ToClusterRoleAssociationOutputWithContext(context.Background())
}

func (i *ClusterRoleAssociation) ToClusterRoleAssociationOutputWithContext(ctx context.Context) ClusterRoleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleAssociationOutput)
}

func (i *ClusterRoleAssociation) ToClusterRoleAssociationPtrOutput() ClusterRoleAssociationPtrOutput {
	return i.ToClusterRoleAssociationPtrOutputWithContext(context.Background())
}

func (i *ClusterRoleAssociation) ToClusterRoleAssociationPtrOutputWithContext(ctx context.Context) ClusterRoleAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleAssociationPtrOutput)
}

type ClusterRoleAssociationPtrInput interface {
	pulumi.Input

	ToClusterRoleAssociationPtrOutput() ClusterRoleAssociationPtrOutput
	ToClusterRoleAssociationPtrOutputWithContext(ctx context.Context) ClusterRoleAssociationPtrOutput
}

type clusterRoleAssociationPtrType ClusterRoleAssociationArgs

func (*clusterRoleAssociationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleAssociation)(nil))
}

func (i *clusterRoleAssociationPtrType) ToClusterRoleAssociationPtrOutput() ClusterRoleAssociationPtrOutput {
	return i.ToClusterRoleAssociationPtrOutputWithContext(context.Background())
}

func (i *clusterRoleAssociationPtrType) ToClusterRoleAssociationPtrOutputWithContext(ctx context.Context) ClusterRoleAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleAssociationPtrOutput)
}

// ClusterRoleAssociationArrayInput is an input type that accepts ClusterRoleAssociationArray and ClusterRoleAssociationArrayOutput values.
// You can construct a concrete instance of `ClusterRoleAssociationArrayInput` via:
//
//          ClusterRoleAssociationArray{ ClusterRoleAssociationArgs{...} }
type ClusterRoleAssociationArrayInput interface {
	pulumi.Input

	ToClusterRoleAssociationArrayOutput() ClusterRoleAssociationArrayOutput
	ToClusterRoleAssociationArrayOutputWithContext(context.Context) ClusterRoleAssociationArrayOutput
}

type ClusterRoleAssociationArray []ClusterRoleAssociationInput

func (ClusterRoleAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRoleAssociation)(nil)).Elem()
}

func (i ClusterRoleAssociationArray) ToClusterRoleAssociationArrayOutput() ClusterRoleAssociationArrayOutput {
	return i.ToClusterRoleAssociationArrayOutputWithContext(context.Background())
}

func (i ClusterRoleAssociationArray) ToClusterRoleAssociationArrayOutputWithContext(ctx context.Context) ClusterRoleAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleAssociationArrayOutput)
}

// ClusterRoleAssociationMapInput is an input type that accepts ClusterRoleAssociationMap and ClusterRoleAssociationMapOutput values.
// You can construct a concrete instance of `ClusterRoleAssociationMapInput` via:
//
//          ClusterRoleAssociationMap{ "key": ClusterRoleAssociationArgs{...} }
type ClusterRoleAssociationMapInput interface {
	pulumi.Input

	ToClusterRoleAssociationMapOutput() ClusterRoleAssociationMapOutput
	ToClusterRoleAssociationMapOutputWithContext(context.Context) ClusterRoleAssociationMapOutput
}

type ClusterRoleAssociationMap map[string]ClusterRoleAssociationInput

func (ClusterRoleAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRoleAssociation)(nil)).Elem()
}

func (i ClusterRoleAssociationMap) ToClusterRoleAssociationMapOutput() ClusterRoleAssociationMapOutput {
	return i.ToClusterRoleAssociationMapOutputWithContext(context.Background())
}

func (i ClusterRoleAssociationMap) ToClusterRoleAssociationMapOutputWithContext(ctx context.Context) ClusterRoleAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleAssociationMapOutput)
}

type ClusterRoleAssociationOutput struct {
	*pulumi.OutputState
}

func (ClusterRoleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoleAssociation)(nil))
}

func (o ClusterRoleAssociationOutput) ToClusterRoleAssociationOutput() ClusterRoleAssociationOutput {
	return o
}

func (o ClusterRoleAssociationOutput) ToClusterRoleAssociationOutputWithContext(ctx context.Context) ClusterRoleAssociationOutput {
	return o
}

func (o ClusterRoleAssociationOutput) ToClusterRoleAssociationPtrOutput() ClusterRoleAssociationPtrOutput {
	return o.ToClusterRoleAssociationPtrOutputWithContext(context.Background())
}

func (o ClusterRoleAssociationOutput) ToClusterRoleAssociationPtrOutputWithContext(ctx context.Context) ClusterRoleAssociationPtrOutput {
	return o.ApplyT(func(v ClusterRoleAssociation) *ClusterRoleAssociation {
		return &v
	}).(ClusterRoleAssociationPtrOutput)
}

type ClusterRoleAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (ClusterRoleAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRoleAssociation)(nil))
}

func (o ClusterRoleAssociationPtrOutput) ToClusterRoleAssociationPtrOutput() ClusterRoleAssociationPtrOutput {
	return o
}

func (o ClusterRoleAssociationPtrOutput) ToClusterRoleAssociationPtrOutputWithContext(ctx context.Context) ClusterRoleAssociationPtrOutput {
	return o
}

type ClusterRoleAssociationArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoleAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoleAssociation)(nil))
}

func (o ClusterRoleAssociationArrayOutput) ToClusterRoleAssociationArrayOutput() ClusterRoleAssociationArrayOutput {
	return o
}

func (o ClusterRoleAssociationArrayOutput) ToClusterRoleAssociationArrayOutputWithContext(ctx context.Context) ClusterRoleAssociationArrayOutput {
	return o
}

func (o ClusterRoleAssociationArrayOutput) Index(i pulumi.IntInput) ClusterRoleAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRoleAssociation {
		return vs[0].([]ClusterRoleAssociation)[vs[1].(int)]
	}).(ClusterRoleAssociationOutput)
}

type ClusterRoleAssociationMapOutput struct{ *pulumi.OutputState }

func (ClusterRoleAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterRoleAssociation)(nil))
}

func (o ClusterRoleAssociationMapOutput) ToClusterRoleAssociationMapOutput() ClusterRoleAssociationMapOutput {
	return o
}

func (o ClusterRoleAssociationMapOutput) ToClusterRoleAssociationMapOutputWithContext(ctx context.Context) ClusterRoleAssociationMapOutput {
	return o
}

func (o ClusterRoleAssociationMapOutput) MapIndex(k pulumi.StringInput) ClusterRoleAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterRoleAssociation {
		return vs[0].(map[string]ClusterRoleAssociation)[vs[1].(string)]
	}).(ClusterRoleAssociationOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterRoleAssociationOutput{})
	pulumi.RegisterOutputType(ClusterRoleAssociationPtrOutput{})
	pulumi.RegisterOutputType(ClusterRoleAssociationArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoleAssociationMapOutput{})
}
