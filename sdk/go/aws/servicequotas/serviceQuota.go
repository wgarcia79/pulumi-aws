// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an individual Service Quota.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/servicequotas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicequotas.NewServiceQuota(ctx, "example", &servicequotas.ServiceQuotaArgs{
// 			QuotaCode:   pulumi.String("L-F678F1CE"),
// 			ServiceCode: pulumi.String("vpc"),
// 			Value:       pulumi.Float64(75),
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
// ~> *NOTE* This resource does not require explicit import and will assume management of an existing service quota on resource creation. `aws_servicequotas_service_quota` can be imported by using the service code and quota code, separated by a front slash (`/`), e.g.
//
// ```sh
//  $ pulumi import aws:servicequotas/serviceQuota:ServiceQuota example vpc/L-F678F1CE
// ```
type ServiceQuota struct {
	pulumi.CustomResourceState

	// Whether the service quota can be increased.
	Adjustable pulumi.BoolOutput `pulumi:"adjustable"`
	// Amazon Resource Name (ARN) of the service quota.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Default value of the service quota.
	DefaultValue pulumi.Float64Output `pulumi:"defaultValue"`
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode pulumi.StringOutput `pulumi:"quotaCode"`
	// Name of the quota.
	QuotaName     pulumi.StringOutput `pulumi:"quotaName"`
	RequestId     pulumi.StringOutput `pulumi:"requestId"`
	RequestStatus pulumi.StringOutput `pulumi:"requestStatus"`
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode pulumi.StringOutput `pulumi:"serviceCode"`
	// Name of the service.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value pulumi.Float64Output `pulumi:"value"`
}

// NewServiceQuota registers a new resource with the given unique name, arguments, and options.
func NewServiceQuota(ctx *pulumi.Context,
	name string, args *ServiceQuotaArgs, opts ...pulumi.ResourceOption) (*ServiceQuota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QuotaCode == nil {
		return nil, errors.New("invalid value for required argument 'QuotaCode'")
	}
	if args.ServiceCode == nil {
		return nil, errors.New("invalid value for required argument 'ServiceCode'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource ServiceQuota
	err := ctx.RegisterResource("aws:servicequotas/serviceQuota:ServiceQuota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceQuota gets an existing ServiceQuota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceQuotaState, opts ...pulumi.ResourceOption) (*ServiceQuota, error) {
	var resource ServiceQuota
	err := ctx.ReadResource("aws:servicequotas/serviceQuota:ServiceQuota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceQuota resources.
type serviceQuotaState struct {
	// Whether the service quota can be increased.
	Adjustable *bool `pulumi:"adjustable"`
	// Amazon Resource Name (ARN) of the service quota.
	Arn *string `pulumi:"arn"`
	// Default value of the service quota.
	DefaultValue *float64 `pulumi:"defaultValue"`
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode *string `pulumi:"quotaCode"`
	// Name of the quota.
	QuotaName     *string `pulumi:"quotaName"`
	RequestId     *string `pulumi:"requestId"`
	RequestStatus *string `pulumi:"requestStatus"`
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode *string `pulumi:"serviceCode"`
	// Name of the service.
	ServiceName *string `pulumi:"serviceName"`
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value *float64 `pulumi:"value"`
}

type ServiceQuotaState struct {
	// Whether the service quota can be increased.
	Adjustable pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the service quota.
	Arn pulumi.StringPtrInput
	// Default value of the service quota.
	DefaultValue pulumi.Float64PtrInput
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode pulumi.StringPtrInput
	// Name of the quota.
	QuotaName     pulumi.StringPtrInput
	RequestId     pulumi.StringPtrInput
	RequestStatus pulumi.StringPtrInput
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode pulumi.StringPtrInput
	// Name of the service.
	ServiceName pulumi.StringPtrInput
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value pulumi.Float64PtrInput
}

func (ServiceQuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceQuotaState)(nil)).Elem()
}

type serviceQuotaArgs struct {
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode string `pulumi:"quotaCode"`
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode string `pulumi:"serviceCode"`
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value float64 `pulumi:"value"`
}

// The set of arguments for constructing a ServiceQuota resource.
type ServiceQuotaArgs struct {
	// Code of the service quota to track. For example: `L-F678F1CE`. Available values can be found with the [AWS CLI service-quotas list-service-quotas command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-service-quotas.html).
	QuotaCode pulumi.StringInput
	// Code of the service to track. For example: `vpc`. Available values can be found with the [AWS CLI service-quotas list-services command](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/list-services.html).
	ServiceCode pulumi.StringInput
	// Float specifying the desired value for the service quota. If the desired value is higher than the current value, a quota increase request is submitted. When a known request is submitted and pending, the value reflects the desired value of the pending request.
	Value pulumi.Float64Input
}

func (ServiceQuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceQuotaArgs)(nil)).Elem()
}

type ServiceQuotaInput interface {
	pulumi.Input

	ToServiceQuotaOutput() ServiceQuotaOutput
	ToServiceQuotaOutputWithContext(ctx context.Context) ServiceQuotaOutput
}

func (*ServiceQuota) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuota)(nil))
}

func (i *ServiceQuota) ToServiceQuotaOutput() ServiceQuotaOutput {
	return i.ToServiceQuotaOutputWithContext(context.Background())
}

func (i *ServiceQuota) ToServiceQuotaOutputWithContext(ctx context.Context) ServiceQuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaOutput)
}

func (i *ServiceQuota) ToServiceQuotaPtrOutput() ServiceQuotaPtrOutput {
	return i.ToServiceQuotaPtrOutputWithContext(context.Background())
}

func (i *ServiceQuota) ToServiceQuotaPtrOutputWithContext(ctx context.Context) ServiceQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaPtrOutput)
}

type ServiceQuotaPtrInput interface {
	pulumi.Input

	ToServiceQuotaPtrOutput() ServiceQuotaPtrOutput
	ToServiceQuotaPtrOutputWithContext(ctx context.Context) ServiceQuotaPtrOutput
}

type serviceQuotaPtrType ServiceQuotaArgs

func (*serviceQuotaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceQuota)(nil))
}

func (i *serviceQuotaPtrType) ToServiceQuotaPtrOutput() ServiceQuotaPtrOutput {
	return i.ToServiceQuotaPtrOutputWithContext(context.Background())
}

func (i *serviceQuotaPtrType) ToServiceQuotaPtrOutputWithContext(ctx context.Context) ServiceQuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaPtrOutput)
}

// ServiceQuotaArrayInput is an input type that accepts ServiceQuotaArray and ServiceQuotaArrayOutput values.
// You can construct a concrete instance of `ServiceQuotaArrayInput` via:
//
//          ServiceQuotaArray{ ServiceQuotaArgs{...} }
type ServiceQuotaArrayInput interface {
	pulumi.Input

	ToServiceQuotaArrayOutput() ServiceQuotaArrayOutput
	ToServiceQuotaArrayOutputWithContext(context.Context) ServiceQuotaArrayOutput
}

type ServiceQuotaArray []ServiceQuotaInput

func (ServiceQuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ServiceQuota)(nil))
}

func (i ServiceQuotaArray) ToServiceQuotaArrayOutput() ServiceQuotaArrayOutput {
	return i.ToServiceQuotaArrayOutputWithContext(context.Background())
}

func (i ServiceQuotaArray) ToServiceQuotaArrayOutputWithContext(ctx context.Context) ServiceQuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaArrayOutput)
}

// ServiceQuotaMapInput is an input type that accepts ServiceQuotaMap and ServiceQuotaMapOutput values.
// You can construct a concrete instance of `ServiceQuotaMapInput` via:
//
//          ServiceQuotaMap{ "key": ServiceQuotaArgs{...} }
type ServiceQuotaMapInput interface {
	pulumi.Input

	ToServiceQuotaMapOutput() ServiceQuotaMapOutput
	ToServiceQuotaMapOutputWithContext(context.Context) ServiceQuotaMapOutput
}

type ServiceQuotaMap map[string]ServiceQuotaInput

func (ServiceQuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ServiceQuota)(nil))
}

func (i ServiceQuotaMap) ToServiceQuotaMapOutput() ServiceQuotaMapOutput {
	return i.ToServiceQuotaMapOutputWithContext(context.Background())
}

func (i ServiceQuotaMap) ToServiceQuotaMapOutputWithContext(ctx context.Context) ServiceQuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQuotaMapOutput)
}

type ServiceQuotaOutput struct {
	*pulumi.OutputState
}

func (ServiceQuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQuota)(nil))
}

func (o ServiceQuotaOutput) ToServiceQuotaOutput() ServiceQuotaOutput {
	return o
}

func (o ServiceQuotaOutput) ToServiceQuotaOutputWithContext(ctx context.Context) ServiceQuotaOutput {
	return o
}

func (o ServiceQuotaOutput) ToServiceQuotaPtrOutput() ServiceQuotaPtrOutput {
	return o.ToServiceQuotaPtrOutputWithContext(context.Background())
}

func (o ServiceQuotaOutput) ToServiceQuotaPtrOutputWithContext(ctx context.Context) ServiceQuotaPtrOutput {
	return o.ApplyT(func(v ServiceQuota) *ServiceQuota {
		return &v
	}).(ServiceQuotaPtrOutput)
}

type ServiceQuotaPtrOutput struct {
	*pulumi.OutputState
}

func (ServiceQuotaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceQuota)(nil))
}

func (o ServiceQuotaPtrOutput) ToServiceQuotaPtrOutput() ServiceQuotaPtrOutput {
	return o
}

func (o ServiceQuotaPtrOutput) ToServiceQuotaPtrOutputWithContext(ctx context.Context) ServiceQuotaPtrOutput {
	return o
}

type ServiceQuotaArrayOutput struct{ *pulumi.OutputState }

func (ServiceQuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQuota)(nil))
}

func (o ServiceQuotaArrayOutput) ToServiceQuotaArrayOutput() ServiceQuotaArrayOutput {
	return o
}

func (o ServiceQuotaArrayOutput) ToServiceQuotaArrayOutputWithContext(ctx context.Context) ServiceQuotaArrayOutput {
	return o
}

func (o ServiceQuotaArrayOutput) Index(i pulumi.IntInput) ServiceQuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceQuota {
		return vs[0].([]ServiceQuota)[vs[1].(int)]
	}).(ServiceQuotaOutput)
}

type ServiceQuotaMapOutput struct{ *pulumi.OutputState }

func (ServiceQuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceQuota)(nil))
}

func (o ServiceQuotaMapOutput) ToServiceQuotaMapOutput() ServiceQuotaMapOutput {
	return o
}

func (o ServiceQuotaMapOutput) ToServiceQuotaMapOutputWithContext(ctx context.Context) ServiceQuotaMapOutput {
	return o
}

func (o ServiceQuotaMapOutput) MapIndex(k pulumi.StringInput) ServiceQuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceQuota {
		return vs[0].(map[string]ServiceQuota)[vs[1].(string)]
	}).(ServiceQuotaOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceQuotaOutput{})
	pulumi.RegisterOutputType(ServiceQuotaPtrOutput{})
	pulumi.RegisterOutputType(ServiceQuotaArrayOutput{})
	pulumi.RegisterOutputType(ServiceQuotaMapOutput{})
}
