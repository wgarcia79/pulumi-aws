// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 domain name.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
//
// > **Note:** This resource establishes ownership of and the TLS settings for
// a particular domain name. An API stage can be associated with the domain name using the `apigatewayv2.ApiMapping` resource.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewDomainName(ctx, "example", &apigatewayv2.DomainNameArgs{
// 			DomainName: pulumi.String("ws-api.example.com"),
// 			DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
// 				CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
// 				EndpointType:   pulumi.String("REGIONAL"),
// 				SecurityPolicy: pulumi.String("TLS_1_2"),
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
// `aws_apigatewayv2_domain_name` can be imported by using the domain name, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/domainName:DomainName example ws-api.example.com
// ```
type DomainName struct {
	pulumi.CustomResourceState

	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression pulumi.StringOutput `pulumi:"apiMappingSelectionExpression"`
	// The ARN of the domain name.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The domain name. Must be between 1 and 512 characters in length.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationOutput `pulumi:"domainNameConfiguration"`
	// The mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrOutput `pulumi:"mutualTlsAuthentication"`
	// A map of tags to assign to the domain name. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDomainName registers a new resource with the given unique name, arguments, and options.
func NewDomainName(ctx *pulumi.Context,
	name string, args *DomainNameArgs, opts ...pulumi.ResourceOption) (*DomainName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.DomainNameConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DomainNameConfiguration'")
	}
	var resource DomainName
	err := ctx.RegisterResource("aws:apigatewayv2/domainName:DomainName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainName gets an existing DomainName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNameState, opts ...pulumi.ResourceOption) (*DomainName, error) {
	var resource DomainName
	err := ctx.ReadResource("aws:apigatewayv2/domainName:DomainName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainName resources.
type domainNameState struct {
	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression *string `pulumi:"apiMappingSelectionExpression"`
	// The ARN of the domain name.
	Arn *string `pulumi:"arn"`
	// The domain name. Must be between 1 and 512 characters in length.
	DomainName *string `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration *DomainNameDomainNameConfiguration `pulumi:"domainNameConfiguration"`
	// The mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication *DomainNameMutualTlsAuthentication `pulumi:"mutualTlsAuthentication"`
	// A map of tags to assign to the domain name. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DomainNameState struct {
	// The [API mapping selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-mapping-selection-expressions) for the domain name.
	ApiMappingSelectionExpression pulumi.StringPtrInput
	// The ARN of the domain name.
	Arn pulumi.StringPtrInput
	// The domain name. Must be between 1 and 512 characters in length.
	DomainName pulumi.StringPtrInput
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationPtrInput
	// The mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrInput
	// A map of tags to assign to the domain name. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (DomainNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameState)(nil)).Elem()
}

type domainNameArgs struct {
	// The domain name. Must be between 1 and 512 characters in length.
	DomainName string `pulumi:"domainName"`
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfiguration `pulumi:"domainNameConfiguration"`
	// The mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication *DomainNameMutualTlsAuthentication `pulumi:"mutualTlsAuthentication"`
	// A map of tags to assign to the domain name. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a DomainName resource.
type DomainNameArgs struct {
	// The domain name. Must be between 1 and 512 characters in length.
	DomainName pulumi.StringInput
	// The domain name configuration.
	DomainNameConfiguration DomainNameDomainNameConfigurationInput
	// The mutual TLS authentication configuration for the domain name.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrInput
	// A map of tags to assign to the domain name. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (DomainNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameArgs)(nil)).Elem()
}

type DomainNameInput interface {
	pulumi.Input

	ToDomainNameOutput() DomainNameOutput
	ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput
}

func (*DomainName) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainName)(nil))
}

func (i *DomainName) ToDomainNameOutput() DomainNameOutput {
	return i.ToDomainNameOutputWithContext(context.Background())
}

func (i *DomainName) ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameOutput)
}

func (i *DomainName) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return i.ToDomainNamePtrOutputWithContext(context.Background())
}

func (i *DomainName) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNamePtrOutput)
}

type DomainNamePtrInput interface {
	pulumi.Input

	ToDomainNamePtrOutput() DomainNamePtrOutput
	ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput
}

type domainNamePtrType DomainNameArgs

func (*domainNamePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainName)(nil))
}

func (i *domainNamePtrType) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return i.ToDomainNamePtrOutputWithContext(context.Background())
}

func (i *domainNamePtrType) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNamePtrOutput)
}

// DomainNameArrayInput is an input type that accepts DomainNameArray and DomainNameArrayOutput values.
// You can construct a concrete instance of `DomainNameArrayInput` via:
//
//          DomainNameArray{ DomainNameArgs{...} }
type DomainNameArrayInput interface {
	pulumi.Input

	ToDomainNameArrayOutput() DomainNameArrayOutput
	ToDomainNameArrayOutputWithContext(context.Context) DomainNameArrayOutput
}

type DomainNameArray []DomainNameInput

func (DomainNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainName)(nil)).Elem()
}

func (i DomainNameArray) ToDomainNameArrayOutput() DomainNameArrayOutput {
	return i.ToDomainNameArrayOutputWithContext(context.Background())
}

func (i DomainNameArray) ToDomainNameArrayOutputWithContext(ctx context.Context) DomainNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameArrayOutput)
}

// DomainNameMapInput is an input type that accepts DomainNameMap and DomainNameMapOutput values.
// You can construct a concrete instance of `DomainNameMapInput` via:
//
//          DomainNameMap{ "key": DomainNameArgs{...} }
type DomainNameMapInput interface {
	pulumi.Input

	ToDomainNameMapOutput() DomainNameMapOutput
	ToDomainNameMapOutputWithContext(context.Context) DomainNameMapOutput
}

type DomainNameMap map[string]DomainNameInput

func (DomainNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainName)(nil)).Elem()
}

func (i DomainNameMap) ToDomainNameMapOutput() DomainNameMapOutput {
	return i.ToDomainNameMapOutputWithContext(context.Background())
}

func (i DomainNameMap) ToDomainNameMapOutputWithContext(ctx context.Context) DomainNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameMapOutput)
}

type DomainNameOutput struct {
	*pulumi.OutputState
}

func (DomainNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainName)(nil))
}

func (o DomainNameOutput) ToDomainNameOutput() DomainNameOutput {
	return o
}

func (o DomainNameOutput) ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput {
	return o
}

func (o DomainNameOutput) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return o.ToDomainNamePtrOutputWithContext(context.Background())
}

func (o DomainNameOutput) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return o.ApplyT(func(v DomainName) *DomainName {
		return &v
	}).(DomainNamePtrOutput)
}

type DomainNamePtrOutput struct {
	*pulumi.OutputState
}

func (DomainNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainName)(nil))
}

func (o DomainNamePtrOutput) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return o
}

func (o DomainNamePtrOutput) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return o
}

type DomainNameArrayOutput struct{ *pulumi.OutputState }

func (DomainNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainName)(nil))
}

func (o DomainNameArrayOutput) ToDomainNameArrayOutput() DomainNameArrayOutput {
	return o
}

func (o DomainNameArrayOutput) ToDomainNameArrayOutputWithContext(ctx context.Context) DomainNameArrayOutput {
	return o
}

func (o DomainNameArrayOutput) Index(i pulumi.IntInput) DomainNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainName {
		return vs[0].([]DomainName)[vs[1].(int)]
	}).(DomainNameOutput)
}

type DomainNameMapOutput struct{ *pulumi.OutputState }

func (DomainNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainName)(nil))
}

func (o DomainNameMapOutput) ToDomainNameMapOutput() DomainNameMapOutput {
	return o
}

func (o DomainNameMapOutput) ToDomainNameMapOutputWithContext(ctx context.Context) DomainNameMapOutput {
	return o
}

func (o DomainNameMapOutput) MapIndex(k pulumi.StringInput) DomainNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainName {
		return vs[0].(map[string]DomainName)[vs[1].(string)]
	}).(DomainNameOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainNameOutput{})
	pulumi.RegisterOutputType(DomainNamePtrOutput{})
	pulumi.RegisterOutputType(DomainNameArrayOutput{})
	pulumi.RegisterOutputType(DomainNameMapOutput{})
}
