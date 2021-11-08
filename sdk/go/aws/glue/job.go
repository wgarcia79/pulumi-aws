// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue Job resource.
//
// > Glue functionality, such as monitoring and logging of jobs, is typically managed with the `defaultArguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.
//
// ## Example Usage
// ### Python Job
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewJob(ctx, "example", &glue.JobArgs{
// 			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			Command: &glue.JobCommandArgs{
// 				ScriptLocation: pulumi.String(fmt.Sprintf("%v%v%v", "s3://", aws_s3_bucket.Example.Bucket, "/example.py")),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Scala Job
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := glue.NewJob(ctx, "example", &glue.JobArgs{
// 			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
// 			Command: &glue.JobCommandArgs{
// 				ScriptLocation: pulumi.String(fmt.Sprintf("%v%v%v", "s3://", aws_s3_bucket.Example.Bucket, "/example.scala")),
// 			},
// 			DefaultArguments: pulumi.StringMap{
// 				"--job-language": pulumi.String("scala"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enabling CloudWatch Logs and Metrics
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/glue"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
// 			RetentionInDays: pulumi.Int(14),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = glue.NewJob(ctx, "exampleJob", &glue.JobArgs{
// 			DefaultArguments: pulumi.StringMap{
// 				"--continuous-log-logGroup":          exampleLogGroup.Name,
// 				"--enable-continuous-cloudwatch-log": pulumi.String("true"),
// 				"--enable-continuous-log-filter":     pulumi.String("true"),
// 				"--enable-metrics":                   pulumi.String(""),
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
// Glue Jobs can be imported using `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:glue/job:Job MyJob MyJob
// ```
type Job struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of Glue Job
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The command of the job. Defined below.
	Command JobCommandOutput `pulumi:"command"`
	// The list of connections used for this job.
	Connections pulumi.StringArrayOutput `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.StringMapOutput `pulumi:"defaultArguments"`
	// Description of the job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty JobExecutionPropertyOutput `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringOutput `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
	MaxCapacity pulumi.Float64Output `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrOutput `pulumi:"maxRetries"`
	// The name you assign to this job. It must be unique in your account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments pulumi.StringMapOutput `pulumi:"nonOverridableArguments"`
	// Notification property of the job. Defined below.
	NotificationProperty JobNotificationPropertyOutput `pulumi:"notificationProperty"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrOutput `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrOutput `pulumi:"securityConfiguration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Command == nil {
		return nil, errors.New("invalid value for required argument 'Command'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource Job
	err := ctx.RegisterResource("aws:glue/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("aws:glue/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Amazon Resource Name (ARN) of Glue Job
	Arn *string `pulumi:"arn"`
	// The command of the job. Defined below.
	Command *JobCommand `pulumi:"command"`
	// The list of connections used for this job.
	Connections []string `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments map[string]string `pulumi:"defaultArguments"`
	// Description of the job.
	Description *string `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty *JobExecutionProperty `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name you assign to this job. It must be unique in your account.
	Name *string `pulumi:"name"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments map[string]string `pulumi:"nonOverridableArguments"`
	// Notification property of the job. Defined below.
	NotificationProperty *JobNotificationProperty `pulumi:"notificationProperty"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
}

type JobState struct {
	// Amazon Resource Name (ARN) of Glue Job
	Arn pulumi.StringPtrInput
	// The command of the job. Defined below.
	Command JobCommandPtrInput
	// The list of connections used for this job.
	Connections pulumi.StringArrayInput
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.StringMapInput
	// Description of the job.
	Description pulumi.StringPtrInput
	// Execution property of the job. Defined below.
	ExecutionProperty JobExecutionPropertyPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name you assign to this job. It must be unique in your account.
	Name pulumi.StringPtrInput
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments pulumi.StringMapInput
	// Notification property of the job. Defined below.
	NotificationProperty JobNotificationPropertyPtrInput
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrInput
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringPtrInput
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// The command of the job. Defined below.
	Command JobCommand `pulumi:"command"`
	// The list of connections used for this job.
	Connections []string `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments map[string]string `pulumi:"defaultArguments"`
	// Description of the job.
	Description *string `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty *JobExecutionProperty `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name you assign to this job. It must be unique in your account.
	Name *string `pulumi:"name"`
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments map[string]string `pulumi:"nonOverridableArguments"`
	// Notification property of the job. Defined below.
	NotificationProperty *JobNotificationProperty `pulumi:"notificationProperty"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn string `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// The command of the job. Defined below.
	Command JobCommandInput
	// The list of connections used for this job.
	Connections pulumi.StringArrayInput
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.StringMapInput
	// Description of the job.
	Description pulumi.StringPtrInput
	// Execution property of the job. Defined below.
	ExecutionProperty JobExecutionPropertyPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name you assign to this job. It must be unique in your account.
	Name pulumi.StringPtrInput
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments pulumi.StringMapInput
	// Notification property of the job. Defined below.
	NotificationProperty JobNotificationPropertyPtrInput
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrInput
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringInput
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

func (i *Job) ToJobPtrOutput() JobPtrOutput {
	return i.ToJobPtrOutputWithContext(context.Background())
}

func (i *Job) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPtrOutput)
}

type JobPtrInput interface {
	pulumi.Input

	ToJobPtrOutput() JobPtrOutput
	ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput
}

type jobPtrType JobArgs

func (*jobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil))
}

func (i *jobPtrType) ToJobPtrOutput() JobPtrOutput {
	return i.ToJobPtrOutputWithContext(context.Background())
}

func (i *jobPtrType) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPtrOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//          JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//          JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) ToJobPtrOutput() JobPtrOutput {
	return o.ToJobPtrOutputWithContext(context.Background())
}

func (o JobOutput) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Job) *Job {
		return &v
	}).(JobPtrOutput)
}

type JobPtrOutput struct{ *pulumi.OutputState }

func (JobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil))
}

func (o JobPtrOutput) ToJobPtrOutput() JobPtrOutput {
	return o
}

func (o JobPtrOutput) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return o
}

func (o JobPtrOutput) Elem() JobOutput {
	return o.ApplyT(func(v *Job) Job {
		if v != nil {
			return *v
		}
		var ret Job
		return ret
	}).(JobOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Job)(nil))
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Job {
		return vs[0].([]Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Job)(nil))
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Job {
		return vs[0].(map[string]Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobPtrInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobPtrOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
