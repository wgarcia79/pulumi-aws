// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
// e.g., to allow easier migration from another management solution
// or to make it easier for an operator to connect through bastion host(s).
//
// > **Note:** It's strongly discouraged to use this data source for querying ephemeral
// instances (e.g., managed via autoscaling group), as the output may change at any time
// and you'd need to re-run `apply` every time an instance comes up or dies.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("aws:ec2/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-instances in the AWS CLI reference][1].
	Filters []GetInstancesFilter `pulumi:"filters"`
	// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
	InstanceStateNames []string `pulumi:"instanceStateNames"`
	// A map of tags, each pair of which must
	// exactly match a pair on desired instances.
	InstanceTags map[string]string `pulumi:"instanceTags"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	Filters []GetInstancesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IDs of instances found through the filter
	Ids                []string          `pulumi:"ids"`
	InstanceStateNames []string          `pulumi:"instanceStateNames"`
	InstanceTags       map[string]string `pulumi:"instanceTags"`
	// Private IP addresses of instances found through the filter
	PrivateIps []string `pulumi:"privateIps"`
	// Public IP addresses of instances found through the filter
	PublicIps []string `pulumi:"publicIps"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			return *r, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-instances in the AWS CLI reference][1].
	Filters GetInstancesFilterArrayInput `pulumi:"filters"`
	// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
	InstanceStateNames pulumi.StringArrayInput `pulumi:"instanceStateNames"`
	// A map of tags, each pair of which must
	// exactly match a pair on desired instances.
	InstanceTags pulumi.StringMapInput `pulumi:"instanceTags"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) Filters() GetInstancesFilterArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesFilter { return v.Filters }).(GetInstancesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// IDs of instances found through the filter
func (o GetInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) InstanceStateNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.InstanceStateNames }).(pulumi.StringArrayOutput)
}

func (o GetInstancesResultOutput) InstanceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetInstancesResult) map[string]string { return v.InstanceTags }).(pulumi.StringMapOutput)
}

// Private IP addresses of instances found through the filter
func (o GetInstancesResultOutput) PrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.PrivateIps }).(pulumi.StringArrayOutput)
}

// Public IP addresses of instances found through the filter
func (o GetInstancesResultOutput) PublicIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []string { return v.PublicIps }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}
