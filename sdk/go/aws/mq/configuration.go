// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mq

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an MQ Configuration Resource.
//
// For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/mq"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mq.NewConfiguration(ctx, "example", &mq.ConfigurationArgs{
// 			Data:          pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n", "<broker xmlns=\"http://activemq.apache.org/schema/core\">\n", "  <plugins>\n", "    <forcePersistencyModeBrokerPlugin persistenceFlag=\"true\"/>\n", "    <statisticsBrokerPlugin/>\n", "    <timeStampingBrokerPlugin ttlCeiling=\"86400000\" zeroExpirationOverride=\"86400000\"/>\n", "  </plugins>\n", "</broker>\n", "\n")),
// 			Description:   pulumi.String("Example Configuration"),
// 			EngineType:    pulumi.String("ActiveMQ"),
// 			EngineVersion: pulumi.String("5.15.0"),
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
// MQ Configurations can be imported using the configuration ID, e.g.
//
// ```sh
//  $ pulumi import aws:mq/configuration:Configuration example c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// ARN of the configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
	AuthenticationStrategy pulumi.StringOutput `pulumi:"authenticationStrategy"`
	// Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data pulumi.StringOutput `pulumi:"data"`
	// Description of the configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType pulumi.StringOutput `pulumi:"engineType"`
	// Version of the broker engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Latest revision of the configuration.
	LatestRevision pulumi.IntOutput `pulumi:"latestRevision"`
	// Name of the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.EngineType == nil {
		return nil, errors.New("invalid value for required argument 'EngineType'")
	}
	if args.EngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'EngineVersion'")
	}
	var resource Configuration
	err := ctx.RegisterResource("aws:mq/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("aws:mq/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// ARN of the configuration.
	Arn *string `pulumi:"arn"`
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
	AuthenticationStrategy *string `pulumi:"authenticationStrategy"`
	// Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data *string `pulumi:"data"`
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType *string `pulumi:"engineType"`
	// Version of the broker engine.
	EngineVersion *string `pulumi:"engineVersion"`
	// Latest revision of the configuration.
	LatestRevision *int `pulumi:"latestRevision"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ConfigurationState struct {
	// ARN of the configuration.
	Arn pulumi.StringPtrInput
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
	AuthenticationStrategy pulumi.StringPtrInput
	// Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data pulumi.StringPtrInput
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType pulumi.StringPtrInput
	// Version of the broker engine.
	EngineVersion pulumi.StringPtrInput
	// Latest revision of the configuration.
	LatestRevision pulumi.IntPtrInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
	AuthenticationStrategy *string `pulumi:"authenticationStrategy"`
	// Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data string `pulumi:"data"`
	// Description of the configuration.
	Description *string `pulumi:"description"`
	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType string `pulumi:"engineType"`
	// Version of the broker engine.
	EngineVersion string `pulumi:"engineVersion"`
	// Name of the configuration.
	Name *string `pulumi:"name"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Authentication strategy associated with the configuration. Valid values are `simple` and `ldap`. `ldap` is not supported for `engineType` `RabbitMQ`.
	AuthenticationStrategy pulumi.StringPtrInput
	// Broker configuration in XML format. See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html) for supported parameters and format of the XML.
	Data pulumi.StringInput
	// Description of the configuration.
	Description pulumi.StringPtrInput
	// Type of broker engine. Valid values are `ActiveMQ` and `RabbitMQ`.
	EngineType pulumi.StringInput
	// Version of the broker engine.
	EngineVersion pulumi.StringInput
	// Name of the configuration.
	Name pulumi.StringPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil))
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

func (i *Configuration) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPtrOutput)
}

type ConfigurationPtrInput interface {
	pulumi.Input

	ToConfigurationPtrOutput() ConfigurationPtrOutput
	ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput
}

type configurationPtrType ConfigurationArgs

func (*configurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil))
}

func (i *configurationPtrType) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *configurationPtrType) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPtrOutput)
}

// ConfigurationArrayInput is an input type that accepts ConfigurationArray and ConfigurationArrayOutput values.
// You can construct a concrete instance of `ConfigurationArrayInput` via:
//
//          ConfigurationArray{ ConfigurationArgs{...} }
type ConfigurationArrayInput interface {
	pulumi.Input

	ToConfigurationArrayOutput() ConfigurationArrayOutput
	ToConfigurationArrayOutputWithContext(context.Context) ConfigurationArrayOutput
}

type ConfigurationArray []ConfigurationInput

func (ConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Configuration)(nil)).Elem()
}

func (i ConfigurationArray) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return i.ToConfigurationArrayOutputWithContext(context.Background())
}

func (i ConfigurationArray) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationArrayOutput)
}

// ConfigurationMapInput is an input type that accepts ConfigurationMap and ConfigurationMapOutput values.
// You can construct a concrete instance of `ConfigurationMapInput` via:
//
//          ConfigurationMap{ "key": ConfigurationArgs{...} }
type ConfigurationMapInput interface {
	pulumi.Input

	ToConfigurationMapOutput() ConfigurationMapOutput
	ToConfigurationMapOutputWithContext(context.Context) ConfigurationMapOutput
}

type ConfigurationMap map[string]ConfigurationInput

func (ConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Configuration)(nil)).Elem()
}

func (i ConfigurationMap) ToConfigurationMapOutput() ConfigurationMapOutput {
	return i.ToConfigurationMapOutputWithContext(context.Background())
}

func (i ConfigurationMap) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationMapOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil))
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o.ToConfigurationPtrOutputWithContext(context.Background())
}

func (o ConfigurationOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Configuration) *Configuration {
		return &v
	}).(ConfigurationPtrOutput)
}

type ConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil))
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) Elem() ConfigurationOutput {
	return o.ApplyT(func(v *Configuration) Configuration {
		if v != nil {
			return *v
		}
		var ret Configuration
		return ret
	}).(ConfigurationOutput)
}

type ConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Configuration)(nil))
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) Index(i pulumi.IntInput) ConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Configuration {
		return vs[0].([]Configuration)[vs[1].(int)]
	}).(ConfigurationOutput)
}

type ConfigurationMapOutput struct{ *pulumi.OutputState }

func (ConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Configuration)(nil))
}

func (o ConfigurationMapOutput) ToConfigurationMapOutput() ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) MapIndex(k pulumi.StringInput) ConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Configuration {
		return vs[0].(map[string]Configuration)[vs[1].(string)]
	}).(ConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationMapOutput{})
}
