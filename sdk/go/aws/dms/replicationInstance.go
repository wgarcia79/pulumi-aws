// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.
//
// ## Example Usage
//
// Create required roles and then create a DMS instance, setting the dependsOn to the required role policy attachments.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dms"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dmsAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"sts:AssumeRole",
// 					},
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Identifiers: []string{
// 								"dms.amazonaws.com",
// 							},
// 							Type: "Service",
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "dms_access_for_endpoint", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "dms_access_for_endpoint_AmazonDMSRedshiftS3Role", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role"),
// 			Role:      dms_access_for_endpoint.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "dms_cloudwatch_logs_role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole"),
// 			Role:      dms_cloudwatch_logs_role.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "dms_vpc_role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(dmsAssumeRole.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "dms_vpc_role_AmazonDMSVPCManagementRole", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"),
// 			Role:      dms_vpc_role.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dms.NewReplicationInstance(ctx, "test", &dms.ReplicationInstanceArgs{
// 			AllocatedStorage:           pulumi.Int(20),
// 			ApplyImmediately:           pulumi.Bool(true),
// 			AutoMinorVersionUpgrade:    pulumi.Bool(true),
// 			AvailabilityZone:           pulumi.String("us-west-2c"),
// 			EngineVersion:              pulumi.String("3.1.4"),
// 			KmsKeyArn:                  pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
// 			MultiAz:                    pulumi.Bool(false),
// 			PreferredMaintenanceWindow: pulumi.String("sun:10:30-sun:14:30"),
// 			PubliclyAccessible:         pulumi.Bool(true),
// 			ReplicationInstanceClass:   pulumi.String("dms.t2.micro"),
// 			ReplicationInstanceId:      pulumi.String("test-dms-replication-instance-tf"),
// 			ReplicationSubnetGroupId:   pulumi.Any(aws_dms_replication_subnet_group.Test - dms - replication - subnet - group - tf.Id),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("test"),
// 			},
// 			VpcSecurityGroupIds: pulumi.StringArray{
// 				pulumi.String("sg-12345678"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
// 			dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
// 			dms_vpc_role_AmazonDMSVPCManagementRole,
// 		}))
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
// Replication instances can be imported using the `replication_instance_id`, e.g.
//
// ```sh
//  $ pulumi import aws:dms/replicationInstance:ReplicationInstance test test-dms-replication-instance-tf
// ```
type ReplicationInstance struct {
	pulumi.CustomResourceState

	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"allowMajorVersionUpgrade"`
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolPtrOutput `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolOutput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The engine version number of the replication instance.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolOutput `pulumi:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolOutput `pulumi:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringOutput `pulumi:"replicationInstanceClass"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringOutput `pulumi:"replicationInstanceId"`
	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps pulumi.StringArrayOutput `pulumi:"replicationInstancePrivateIps"`
	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps pulumi.StringArrayOutput `pulumi:"replicationInstancePublicIps"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringOutput `pulumi:"replicationSubnetGroupId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
}

// NewReplicationInstance registers a new resource with the given unique name, arguments, and options.
func NewReplicationInstance(ctx *pulumi.Context,
	name string, args *ReplicationInstanceArgs, opts ...pulumi.ResourceOption) (*ReplicationInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicationInstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationInstanceClass'")
	}
	if args.ReplicationInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationInstanceId'")
	}
	var resource ReplicationInstance
	err := ctx.RegisterResource("aws:dms/replicationInstance:ReplicationInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationInstance gets an existing ReplicationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationInstanceState, opts ...pulumi.ResourceOption) (*ReplicationInstance, error) {
	var resource ReplicationInstance
	err := ctx.ReadResource("aws:dms/replicationInstance:ReplicationInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationInstance resources.
type replicationInstanceState struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade *bool `pulumi:"allowMajorVersionUpgrade"`
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The engine version number of the replication instance.
	EngineVersion *string `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz *bool `pulumi:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string `pulumi:"replicationInstanceArn"`
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass *string `pulumi:"replicationInstanceClass"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId *string `pulumi:"replicationInstanceId"`
	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps []string `pulumi:"replicationInstancePrivateIps"`
	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps []string `pulumi:"replicationInstancePublicIps"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId *string `pulumi:"replicationSubnetGroupId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

type ReplicationInstanceState struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntPtrInput
	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade pulumi.BoolPtrInput
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringPtrInput
	// The engine version number of the replication instance.
	EngineVersion pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringPtrInput
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolPtrInput
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringPtrInput
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringPtrInput
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringPtrInput
	// A list of the private IP addresses of the replication instance.
	ReplicationInstancePrivateIps pulumi.StringArrayInput
	// A list of the public IP addresses of the replication instance.
	ReplicationInstancePublicIps pulumi.StringArrayInput
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (ReplicationInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationInstanceState)(nil)).Elem()
}

type replicationInstanceArgs struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade *bool `pulumi:"allowMajorVersionUpgrade"`
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The engine version number of the replication instance.
	EngineVersion *string `pulumi:"engineVersion"`
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz *bool `pulumi:"multiAz"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass string `pulumi:"replicationInstanceClass"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId string `pulumi:"replicationInstanceId"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId *string `pulumi:"replicationSubnetGroupId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a ReplicationInstance resource.
type ReplicationInstanceArgs struct {
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage pulumi.IntPtrInput
	// Indicates that major version upgrades are allowed.
	AllowMajorVersionUpgrade pulumi.BoolPtrInput
	// Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
	ApplyImmediately pulumi.BoolPtrInput
	// Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the replication instance will be created in.
	AvailabilityZone pulumi.StringPtrInput
	// The engine version number of the replication instance.
	EngineVersion pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kmsKeyArn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn pulumi.StringPtrInput
	// Specifies if the replication instance is a multi-az deployment. You cannot set the `availabilityZone` parameter if the `multiAz` parameter is set to `true`.
	MultiAz pulumi.BoolPtrInput
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
	PubliclyAccessible pulumi.BoolPtrInput
	// The compute and memory capacity of the replication instance as specified by the replication instance class. Can be one of `dms.t2.micro | dms.t2.small | dms.t2.medium | dms.t2.large | dms.c4.large | dms.c4.xlarge | dms.c4.2xlarge | dms.c4.4xlarge`
	ReplicationInstanceClass pulumi.StringInput
	// The replication instance identifier. This parameter is stored as a lowercase string.
	ReplicationInstanceId pulumi.StringInput
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupId pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.
	VpcSecurityGroupIds pulumi.StringArrayInput
}

func (ReplicationInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationInstanceArgs)(nil)).Elem()
}

type ReplicationInstanceInput interface {
	pulumi.Input

	ToReplicationInstanceOutput() ReplicationInstanceOutput
	ToReplicationInstanceOutputWithContext(ctx context.Context) ReplicationInstanceOutput
}

func (*ReplicationInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationInstance)(nil))
}

func (i *ReplicationInstance) ToReplicationInstanceOutput() ReplicationInstanceOutput {
	return i.ToReplicationInstanceOutputWithContext(context.Background())
}

func (i *ReplicationInstance) ToReplicationInstanceOutputWithContext(ctx context.Context) ReplicationInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationInstanceOutput)
}

func (i *ReplicationInstance) ToReplicationInstancePtrOutput() ReplicationInstancePtrOutput {
	return i.ToReplicationInstancePtrOutputWithContext(context.Background())
}

func (i *ReplicationInstance) ToReplicationInstancePtrOutputWithContext(ctx context.Context) ReplicationInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationInstancePtrOutput)
}

type ReplicationInstancePtrInput interface {
	pulumi.Input

	ToReplicationInstancePtrOutput() ReplicationInstancePtrOutput
	ToReplicationInstancePtrOutputWithContext(ctx context.Context) ReplicationInstancePtrOutput
}

type replicationInstancePtrType ReplicationInstanceArgs

func (*replicationInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationInstance)(nil))
}

func (i *replicationInstancePtrType) ToReplicationInstancePtrOutput() ReplicationInstancePtrOutput {
	return i.ToReplicationInstancePtrOutputWithContext(context.Background())
}

func (i *replicationInstancePtrType) ToReplicationInstancePtrOutputWithContext(ctx context.Context) ReplicationInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationInstancePtrOutput)
}

// ReplicationInstanceArrayInput is an input type that accepts ReplicationInstanceArray and ReplicationInstanceArrayOutput values.
// You can construct a concrete instance of `ReplicationInstanceArrayInput` via:
//
//          ReplicationInstanceArray{ ReplicationInstanceArgs{...} }
type ReplicationInstanceArrayInput interface {
	pulumi.Input

	ToReplicationInstanceArrayOutput() ReplicationInstanceArrayOutput
	ToReplicationInstanceArrayOutputWithContext(context.Context) ReplicationInstanceArrayOutput
}

type ReplicationInstanceArray []ReplicationInstanceInput

func (ReplicationInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReplicationInstance)(nil)).Elem()
}

func (i ReplicationInstanceArray) ToReplicationInstanceArrayOutput() ReplicationInstanceArrayOutput {
	return i.ToReplicationInstanceArrayOutputWithContext(context.Background())
}

func (i ReplicationInstanceArray) ToReplicationInstanceArrayOutputWithContext(ctx context.Context) ReplicationInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationInstanceArrayOutput)
}

// ReplicationInstanceMapInput is an input type that accepts ReplicationInstanceMap and ReplicationInstanceMapOutput values.
// You can construct a concrete instance of `ReplicationInstanceMapInput` via:
//
//          ReplicationInstanceMap{ "key": ReplicationInstanceArgs{...} }
type ReplicationInstanceMapInput interface {
	pulumi.Input

	ToReplicationInstanceMapOutput() ReplicationInstanceMapOutput
	ToReplicationInstanceMapOutputWithContext(context.Context) ReplicationInstanceMapOutput
}

type ReplicationInstanceMap map[string]ReplicationInstanceInput

func (ReplicationInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReplicationInstance)(nil)).Elem()
}

func (i ReplicationInstanceMap) ToReplicationInstanceMapOutput() ReplicationInstanceMapOutput {
	return i.ToReplicationInstanceMapOutputWithContext(context.Background())
}

func (i ReplicationInstanceMap) ToReplicationInstanceMapOutputWithContext(ctx context.Context) ReplicationInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationInstanceMapOutput)
}

type ReplicationInstanceOutput struct{ *pulumi.OutputState }

func (ReplicationInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationInstance)(nil))
}

func (o ReplicationInstanceOutput) ToReplicationInstanceOutput() ReplicationInstanceOutput {
	return o
}

func (o ReplicationInstanceOutput) ToReplicationInstanceOutputWithContext(ctx context.Context) ReplicationInstanceOutput {
	return o
}

func (o ReplicationInstanceOutput) ToReplicationInstancePtrOutput() ReplicationInstancePtrOutput {
	return o.ToReplicationInstancePtrOutputWithContext(context.Background())
}

func (o ReplicationInstanceOutput) ToReplicationInstancePtrOutputWithContext(ctx context.Context) ReplicationInstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationInstance) *ReplicationInstance {
		return &v
	}).(ReplicationInstancePtrOutput)
}

type ReplicationInstancePtrOutput struct{ *pulumi.OutputState }

func (ReplicationInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationInstance)(nil))
}

func (o ReplicationInstancePtrOutput) ToReplicationInstancePtrOutput() ReplicationInstancePtrOutput {
	return o
}

func (o ReplicationInstancePtrOutput) ToReplicationInstancePtrOutputWithContext(ctx context.Context) ReplicationInstancePtrOutput {
	return o
}

func (o ReplicationInstancePtrOutput) Elem() ReplicationInstanceOutput {
	return o.ApplyT(func(v *ReplicationInstance) ReplicationInstance {
		if v != nil {
			return *v
		}
		var ret ReplicationInstance
		return ret
	}).(ReplicationInstanceOutput)
}

type ReplicationInstanceArrayOutput struct{ *pulumi.OutputState }

func (ReplicationInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicationInstance)(nil))
}

func (o ReplicationInstanceArrayOutput) ToReplicationInstanceArrayOutput() ReplicationInstanceArrayOutput {
	return o
}

func (o ReplicationInstanceArrayOutput) ToReplicationInstanceArrayOutputWithContext(ctx context.Context) ReplicationInstanceArrayOutput {
	return o
}

func (o ReplicationInstanceArrayOutput) Index(i pulumi.IntInput) ReplicationInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicationInstance {
		return vs[0].([]ReplicationInstance)[vs[1].(int)]
	}).(ReplicationInstanceOutput)
}

type ReplicationInstanceMapOutput struct{ *pulumi.OutputState }

func (ReplicationInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReplicationInstance)(nil))
}

func (o ReplicationInstanceMapOutput) ToReplicationInstanceMapOutput() ReplicationInstanceMapOutput {
	return o
}

func (o ReplicationInstanceMapOutput) ToReplicationInstanceMapOutputWithContext(ctx context.Context) ReplicationInstanceMapOutput {
	return o
}

func (o ReplicationInstanceMapOutput) MapIndex(k pulumi.StringInput) ReplicationInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReplicationInstance {
		return vs[0].(map[string]ReplicationInstance)[vs[1].(string)]
	}).(ReplicationInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationInstanceInput)(nil)).Elem(), &ReplicationInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationInstancePtrInput)(nil)).Elem(), &ReplicationInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationInstanceArrayInput)(nil)).Elem(), ReplicationInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationInstanceMapInput)(nil)).Elem(), ReplicationInstanceMap{})
	pulumi.RegisterOutputType(ReplicationInstanceOutput{})
	pulumi.RegisterOutputType(ReplicationInstancePtrOutput{})
	pulumi.RegisterOutputType(ReplicationInstanceArrayOutput{})
	pulumi.RegisterOutputType(ReplicationInstanceMapOutput{})
}
