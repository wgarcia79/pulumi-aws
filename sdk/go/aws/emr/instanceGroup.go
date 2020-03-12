// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package emr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic MapReduce Cluster Instance Group configuration.
// See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.
//
// > **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
// web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
// this provider will resize any Instance Group to zero when destroying the resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/emr_instance_group.html.markdown.
type InstanceGroup struct {
	pulumi.CustomResourceState

	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy pulumi.StringPtrOutput `pulumi:"autoscalingPolicy"`
	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice pulumi.StringPtrOutput `pulumi:"bidPrice"`
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
	ConfigurationsJson pulumi.StringPtrOutput `pulumi:"configurationsJson"`
	// One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs InstanceGroupEbsConfigArrayOutput `pulumi:"ebsConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized pulumi.BoolPtrOutput `pulumi:"ebsOptimized"`
	// target number of instances for the instance group. defaults to 0.
	InstanceCount pulumi.IntPtrOutput `pulumi:"instanceCount"`
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	RunningInstanceCount pulumi.IntOutput `pulumi:"runningInstanceCount"`
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	if args == nil {
		args = &InstanceGroupArgs{}
	}
	var resource InstanceGroup
	err := ctx.RegisterResource("aws:emr/instanceGroup:InstanceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupState, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	var resource InstanceGroup
	err := ctx.ReadResource("aws:emr/instanceGroup:InstanceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroup resources.
type instanceGroupState struct {
	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy *string `pulumi:"autoscalingPolicy"`
	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice *string `pulumi:"bidPrice"`
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId *string `pulumi:"clusterId"`
	// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
	ConfigurationsJson *string `pulumi:"configurationsJson"`
	// One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs []InstanceGroupEbsConfig `pulumi:"ebsConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// target number of instances for the instance group. defaults to 0.
	InstanceCount *int `pulumi:"instanceCount"`
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType *string `pulumi:"instanceType"`
	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	RunningInstanceCount *int `pulumi:"runningInstanceCount"`
	Status *string `pulumi:"status"`
}

type InstanceGroupState struct {
	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy pulumi.StringPtrInput
	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice pulumi.StringPtrInput
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringPtrInput
	// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
	ConfigurationsJson pulumi.StringPtrInput
	// One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs InstanceGroupEbsConfigArrayInput
	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized pulumi.BoolPtrInput
	// target number of instances for the instance group. defaults to 0.
	InstanceCount pulumi.IntPtrInput
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType pulumi.StringPtrInput
	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	RunningInstanceCount pulumi.IntPtrInput
	Status pulumi.StringPtrInput
}

func (InstanceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupState)(nil)).Elem()
}

type instanceGroupArgs struct {
	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy *string `pulumi:"autoscalingPolicy"`
	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice *string `pulumi:"bidPrice"`
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId string `pulumi:"clusterId"`
	// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
	ConfigurationsJson *string `pulumi:"configurationsJson"`
	// One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs []InstanceGroupEbsConfig `pulumi:"ebsConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized *bool `pulumi:"ebsOptimized"`
	// target number of instances for the instance group. defaults to 0.
	InstanceCount *int `pulumi:"instanceCount"`
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType string `pulumi:"instanceType"`
	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
	AutoscalingPolicy pulumi.StringPtrInput
	// If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice pulumi.StringPtrInput
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringInput
	// A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
	ConfigurationsJson pulumi.StringPtrInput
	// One or more `ebsConfig` blocks as defined below. Changing this forces a new resource to be created.
	EbsConfigs InstanceGroupEbsConfigArrayInput
	// Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
	EbsOptimized pulumi.BoolPtrInput
	// target number of instances for the instance group. defaults to 0.
	InstanceCount pulumi.IntPtrInput
	// The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
	InstanceType pulumi.StringInput
	// Human friendly name given to the instance group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
}

func (InstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupArgs)(nil)).Elem()
}

