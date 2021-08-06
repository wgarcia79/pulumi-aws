// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Batch Job Queue resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/batch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewJobQueue(ctx, "testQueue", &batch.JobQueueArgs{
// 			State:    pulumi.String("ENABLED"),
// 			Priority: pulumi.Int(1),
// 			ComputeEnvironments: pulumi.StringArray{
// 				pulumi.Any(aws_batch_compute_environment.Test_environment_1.Arn),
// 				pulumi.Any(aws_batch_compute_environment.Test_environment_2.Arn),
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
// Batch Job Queue can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:batch/jobQueue:JobQueue test_queue arn:aws:batch:us-east-1:123456789012:job-queue/sample
// ```
type JobQueue struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the job queue.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayOutput `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewJobQueue registers a new resource with the given unique name, arguments, and options.
func NewJobQueue(ctx *pulumi.Context,
	name string, args *JobQueueArgs, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeEnvironments == nil {
		return nil, errors.New("invalid value for required argument 'ComputeEnvironments'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	var resource JobQueue
	err := ctx.RegisterResource("aws:batch/jobQueue:JobQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobQueue gets an existing JobQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobQueueState, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	var resource JobQueue
	err := ctx.ReadResource("aws:batch/jobQueue:JobQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobQueue resources.
type jobQueueState struct {
	// The Amazon Resource Name of the job queue.
	Arn *string `pulumi:"arn"`
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments []string `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name *string `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority *int `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State *string `pulumi:"state"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type JobQueueState struct {
	// The Amazon Resource Name of the job queue.
	Arn pulumi.StringPtrInput
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayInput
	// Specifies the name of the job queue.
	Name pulumi.StringPtrInput
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntPtrInput
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (JobQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueState)(nil)).Elem()
}

type jobQueueArgs struct {
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments []string `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name *string `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority int `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State string `pulumi:"state"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a JobQueue resource.
type JobQueueArgs struct {
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayInput
	// Specifies the name of the job queue.
	Name pulumi.StringPtrInput
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntInput
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (JobQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueArgs)(nil)).Elem()
}

type JobQueueInput interface {
	pulumi.Input

	ToJobQueueOutput() JobQueueOutput
	ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput
}

func (*JobQueue) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueue)(nil))
}

func (i *JobQueue) ToJobQueueOutput() JobQueueOutput {
	return i.ToJobQueueOutputWithContext(context.Background())
}

func (i *JobQueue) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueueOutput)
}

func (i *JobQueue) ToJobQueuePtrOutput() JobQueuePtrOutput {
	return i.ToJobQueuePtrOutputWithContext(context.Background())
}

func (i *JobQueue) ToJobQueuePtrOutputWithContext(ctx context.Context) JobQueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueuePtrOutput)
}

type JobQueuePtrInput interface {
	pulumi.Input

	ToJobQueuePtrOutput() JobQueuePtrOutput
	ToJobQueuePtrOutputWithContext(ctx context.Context) JobQueuePtrOutput
}

type jobQueuePtrType JobQueueArgs

func (*jobQueuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobQueue)(nil))
}

func (i *jobQueuePtrType) ToJobQueuePtrOutput() JobQueuePtrOutput {
	return i.ToJobQueuePtrOutputWithContext(context.Background())
}

func (i *jobQueuePtrType) ToJobQueuePtrOutputWithContext(ctx context.Context) JobQueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueuePtrOutput)
}

// JobQueueArrayInput is an input type that accepts JobQueueArray and JobQueueArrayOutput values.
// You can construct a concrete instance of `JobQueueArrayInput` via:
//
//          JobQueueArray{ JobQueueArgs{...} }
type JobQueueArrayInput interface {
	pulumi.Input

	ToJobQueueArrayOutput() JobQueueArrayOutput
	ToJobQueueArrayOutputWithContext(context.Context) JobQueueArrayOutput
}

type JobQueueArray []JobQueueInput

func (JobQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*JobQueue)(nil)).Elem()
}

func (i JobQueueArray) ToJobQueueArrayOutput() JobQueueArrayOutput {
	return i.ToJobQueueArrayOutputWithContext(context.Background())
}

func (i JobQueueArray) ToJobQueueArrayOutputWithContext(ctx context.Context) JobQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueueArrayOutput)
}

// JobQueueMapInput is an input type that accepts JobQueueMap and JobQueueMapOutput values.
// You can construct a concrete instance of `JobQueueMapInput` via:
//
//          JobQueueMap{ "key": JobQueueArgs{...} }
type JobQueueMapInput interface {
	pulumi.Input

	ToJobQueueMapOutput() JobQueueMapOutput
	ToJobQueueMapOutputWithContext(context.Context) JobQueueMapOutput
}

type JobQueueMap map[string]JobQueueInput

func (JobQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*JobQueue)(nil)).Elem()
}

func (i JobQueueMap) ToJobQueueMapOutput() JobQueueMapOutput {
	return i.ToJobQueueMapOutputWithContext(context.Background())
}

func (i JobQueueMap) ToJobQueueMapOutputWithContext(ctx context.Context) JobQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueueMapOutput)
}

type JobQueueOutput struct {
	*pulumi.OutputState
}

func (JobQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueue)(nil))
}

func (o JobQueueOutput) ToJobQueueOutput() JobQueueOutput {
	return o
}

func (o JobQueueOutput) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return o
}

func (o JobQueueOutput) ToJobQueuePtrOutput() JobQueuePtrOutput {
	return o.ToJobQueuePtrOutputWithContext(context.Background())
}

func (o JobQueueOutput) ToJobQueuePtrOutputWithContext(ctx context.Context) JobQueuePtrOutput {
	return o.ApplyT(func(v JobQueue) *JobQueue {
		return &v
	}).(JobQueuePtrOutput)
}

type JobQueuePtrOutput struct {
	*pulumi.OutputState
}

func (JobQueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobQueue)(nil))
}

func (o JobQueuePtrOutput) ToJobQueuePtrOutput() JobQueuePtrOutput {
	return o
}

func (o JobQueuePtrOutput) ToJobQueuePtrOutputWithContext(ctx context.Context) JobQueuePtrOutput {
	return o
}

type JobQueueArrayOutput struct{ *pulumi.OutputState }

func (JobQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobQueue)(nil))
}

func (o JobQueueArrayOutput) ToJobQueueArrayOutput() JobQueueArrayOutput {
	return o
}

func (o JobQueueArrayOutput) ToJobQueueArrayOutputWithContext(ctx context.Context) JobQueueArrayOutput {
	return o
}

func (o JobQueueArrayOutput) Index(i pulumi.IntInput) JobQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobQueue {
		return vs[0].([]JobQueue)[vs[1].(int)]
	}).(JobQueueOutput)
}

type JobQueueMapOutput struct{ *pulumi.OutputState }

func (JobQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JobQueue)(nil))
}

func (o JobQueueMapOutput) ToJobQueueMapOutput() JobQueueMapOutput {
	return o
}

func (o JobQueueMapOutput) ToJobQueueMapOutputWithContext(ctx context.Context) JobQueueMapOutput {
	return o
}

func (o JobQueueMapOutput) MapIndex(k pulumi.StringInput) JobQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JobQueue {
		return vs[0].(map[string]JobQueue)[vs[1].(string)]
	}).(JobQueueOutput)
}

func init() {
	pulumi.RegisterOutputType(JobQueueOutput{})
	pulumi.RegisterOutputType(JobQueuePtrOutput{})
	pulumi.RegisterOutputType(JobQueueArrayOutput{})
	pulumi.RegisterOutputType(JobQueueMapOutput{})
}
