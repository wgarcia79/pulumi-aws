// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Log subscription for AWS Directory Service that pushes logs to cloudwatch.
//
// ## Import
//
// Directory Service Log Subscriptions can be imported using the directory id, e.g.
//
// ```sh
//  $ pulumi import aws:directoryservice/logService:LogService msad d-1234567890
// ```
type LogService struct {
	pulumi.CustomResourceState

	// The id of directory.
	DirectoryId pulumi.StringOutput `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
}

// NewLogService registers a new resource with the given unique name, arguments, and options.
func NewLogService(ctx *pulumi.Context,
	name string, args *LogServiceArgs, opts ...pulumi.ResourceOption) (*LogService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectoryId == nil {
		return nil, errors.New("invalid value for required argument 'DirectoryId'")
	}
	if args.LogGroupName == nil {
		return nil, errors.New("invalid value for required argument 'LogGroupName'")
	}
	var resource LogService
	err := ctx.RegisterResource("aws:directoryservice/logService:LogService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogService gets an existing LogService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogServiceState, opts ...pulumi.ResourceOption) (*LogService, error) {
	var resource LogService
	err := ctx.ReadResource("aws:directoryservice/logService:LogService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogService resources.
type logServiceState struct {
	// The id of directory.
	DirectoryId *string `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName *string `pulumi:"logGroupName"`
}

type LogServiceState struct {
	// The id of directory.
	DirectoryId pulumi.StringPtrInput
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringPtrInput
}

func (LogServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*logServiceState)(nil)).Elem()
}

type logServiceArgs struct {
	// The id of directory.
	DirectoryId string `pulumi:"directoryId"`
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName string `pulumi:"logGroupName"`
}

// The set of arguments for constructing a LogService resource.
type LogServiceArgs struct {
	// The id of directory.
	DirectoryId pulumi.StringInput
	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName pulumi.StringInput
}

func (LogServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logServiceArgs)(nil)).Elem()
}

type LogServiceInput interface {
	pulumi.Input

	ToLogServiceOutput() LogServiceOutput
	ToLogServiceOutputWithContext(ctx context.Context) LogServiceOutput
}

func (*LogService) ElementType() reflect.Type {
	return reflect.TypeOf((*LogService)(nil))
}

func (i *LogService) ToLogServiceOutput() LogServiceOutput {
	return i.ToLogServiceOutputWithContext(context.Background())
}

func (i *LogService) ToLogServiceOutputWithContext(ctx context.Context) LogServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogServiceOutput)
}

func (i *LogService) ToLogServicePtrOutput() LogServicePtrOutput {
	return i.ToLogServicePtrOutputWithContext(context.Background())
}

func (i *LogService) ToLogServicePtrOutputWithContext(ctx context.Context) LogServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogServicePtrOutput)
}

type LogServicePtrInput interface {
	pulumi.Input

	ToLogServicePtrOutput() LogServicePtrOutput
	ToLogServicePtrOutputWithContext(ctx context.Context) LogServicePtrOutput
}

type logServicePtrType LogServiceArgs

func (*logServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogService)(nil))
}

func (i *logServicePtrType) ToLogServicePtrOutput() LogServicePtrOutput {
	return i.ToLogServicePtrOutputWithContext(context.Background())
}

func (i *logServicePtrType) ToLogServicePtrOutputWithContext(ctx context.Context) LogServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogServicePtrOutput)
}

// LogServiceArrayInput is an input type that accepts LogServiceArray and LogServiceArrayOutput values.
// You can construct a concrete instance of `LogServiceArrayInput` via:
//
//          LogServiceArray{ LogServiceArgs{...} }
type LogServiceArrayInput interface {
	pulumi.Input

	ToLogServiceArrayOutput() LogServiceArrayOutput
	ToLogServiceArrayOutputWithContext(context.Context) LogServiceArrayOutput
}

type LogServiceArray []LogServiceInput

func (LogServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LogService)(nil))
}

func (i LogServiceArray) ToLogServiceArrayOutput() LogServiceArrayOutput {
	return i.ToLogServiceArrayOutputWithContext(context.Background())
}

func (i LogServiceArray) ToLogServiceArrayOutputWithContext(ctx context.Context) LogServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogServiceArrayOutput)
}

// LogServiceMapInput is an input type that accepts LogServiceMap and LogServiceMapOutput values.
// You can construct a concrete instance of `LogServiceMapInput` via:
//
//          LogServiceMap{ "key": LogServiceArgs{...} }
type LogServiceMapInput interface {
	pulumi.Input

	ToLogServiceMapOutput() LogServiceMapOutput
	ToLogServiceMapOutputWithContext(context.Context) LogServiceMapOutput
}

type LogServiceMap map[string]LogServiceInput

func (LogServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LogService)(nil))
}

func (i LogServiceMap) ToLogServiceMapOutput() LogServiceMapOutput {
	return i.ToLogServiceMapOutputWithContext(context.Background())
}

func (i LogServiceMap) ToLogServiceMapOutputWithContext(ctx context.Context) LogServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogServiceMapOutput)
}

type LogServiceOutput struct {
	*pulumi.OutputState
}

func (LogServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogService)(nil))
}

func (o LogServiceOutput) ToLogServiceOutput() LogServiceOutput {
	return o
}

func (o LogServiceOutput) ToLogServiceOutputWithContext(ctx context.Context) LogServiceOutput {
	return o
}

func (o LogServiceOutput) ToLogServicePtrOutput() LogServicePtrOutput {
	return o.ToLogServicePtrOutputWithContext(context.Background())
}

func (o LogServiceOutput) ToLogServicePtrOutputWithContext(ctx context.Context) LogServicePtrOutput {
	return o.ApplyT(func(v LogService) *LogService {
		return &v
	}).(LogServicePtrOutput)
}

type LogServicePtrOutput struct {
	*pulumi.OutputState
}

func (LogServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogService)(nil))
}

func (o LogServicePtrOutput) ToLogServicePtrOutput() LogServicePtrOutput {
	return o
}

func (o LogServicePtrOutput) ToLogServicePtrOutputWithContext(ctx context.Context) LogServicePtrOutput {
	return o
}

type LogServiceArrayOutput struct{ *pulumi.OutputState }

func (LogServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogService)(nil))
}

func (o LogServiceArrayOutput) ToLogServiceArrayOutput() LogServiceArrayOutput {
	return o
}

func (o LogServiceArrayOutput) ToLogServiceArrayOutputWithContext(ctx context.Context) LogServiceArrayOutput {
	return o
}

func (o LogServiceArrayOutput) Index(i pulumi.IntInput) LogServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogService {
		return vs[0].([]LogService)[vs[1].(int)]
	}).(LogServiceOutput)
}

type LogServiceMapOutput struct{ *pulumi.OutputState }

func (LogServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogService)(nil))
}

func (o LogServiceMapOutput) ToLogServiceMapOutput() LogServiceMapOutput {
	return o
}

func (o LogServiceMapOutput) ToLogServiceMapOutputWithContext(ctx context.Context) LogServiceMapOutput {
	return o
}

func (o LogServiceMapOutput) MapIndex(k pulumi.StringInput) LogServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogService {
		return vs[0].(map[string]LogService)[vs[1].(string)]
	}).(LogServiceOutput)
}

func init() {
	pulumi.RegisterOutputType(LogServiceOutput{})
	pulumi.RegisterOutputType(LogServicePtrOutput{})
	pulumi.RegisterOutputType(LogServiceArrayOutput{})
	pulumi.RegisterOutputType(LogServiceMapOutput{})
}
