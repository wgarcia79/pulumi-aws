// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an RDS DB proxy resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html).
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
// 		_, err := rds.NewProxy(ctx, "example", &rds.ProxyArgs{
// 			DebugLogging:      pulumi.Bool(false),
// 			EngineFamily:      pulumi.String("MYSQL"),
// 			IdleClientTimeout: pulumi.Int(1800),
// 			RequireTls:        pulumi.Bool(true),
// 			RoleArn:           pulumi.Any(aws_iam_role.Example.Arn),
// 			VpcSecurityGroupIds: pulumi.StringArray{
// 				pulumi.Any(aws_security_group.Example.Id),
// 			},
// 			VpcSubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example.Id),
// 			},
// 			Auths: rds.ProxyAuthArray{
// 				&rds.ProxyAuthArgs{
// 					AuthScheme:  pulumi.String("SECRETS"),
// 					Description: pulumi.String("example"),
// 					IamAuth:     pulumi.String("DISABLED"),
// 					SecretArn:   pulumi.Any(aws_secretsmanager_secret.Example.Arn),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("example"),
// 				"Key":  pulumi.String("value"),
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
// DB proxies can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:rds/proxy:Proxy example example
// ```
type Proxy struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the proxy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayOutput `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrOutput `pulumi:"debugLogging"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringOutput `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntOutput `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringOutput `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrOutput `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayOutput `pulumi:"vpcSubnetIds"`
}

// NewProxy registers a new resource with the given unique name, arguments, and options.
func NewProxy(ctx *pulumi.Context,
	name string, args *ProxyArgs, opts ...pulumi.ResourceOption) (*Proxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Auths == nil {
		return nil, errors.New("invalid value for required argument 'Auths'")
	}
	if args.EngineFamily == nil {
		return nil, errors.New("invalid value for required argument 'EngineFamily'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.VpcSubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'VpcSubnetIds'")
	}
	var resource Proxy
	err := ctx.RegisterResource("aws:rds/proxy:Proxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxy gets an existing Proxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyState, opts ...pulumi.ResourceOption) (*Proxy, error) {
	var resource Proxy
	err := ctx.ReadResource("aws:rds/proxy:Proxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Proxy resources.
type proxyState struct {
	// The Amazon Resource Name (ARN) for the proxy.
	Arn *string `pulumi:"arn"`
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths []ProxyAuth `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `pulumi:"debugLogging"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `pulumi:"endpoint"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily *string `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *int `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name *string `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls *bool `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn *string `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

type ProxyState struct {
	// The Amazon Resource Name (ARN) for the proxy.
	Arn pulumi.StringPtrInput
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayInput
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrInput
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringPtrInput
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringPtrInput
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntPtrInput
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringPtrInput
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyState)(nil)).Elem()
}

type proxyArgs struct {
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths []ProxyAuth `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging *bool `pulumi:"debugLogging"`
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily string `pulumi:"engineFamily"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout *int `pulumi:"idleClientTimeout"`
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name *string `pulumi:"name"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls *bool `pulumi:"requireTls"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn string `pulumi:"roleArn"`
	// A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

// The set of arguments for constructing a Proxy resource.
type ProxyArgs struct {
	// Configuration block(s) with authorization mechanisms to connect to the associated instances or clusters. Described below.
	Auths ProxyAuthArrayInput
	// Whether the proxy includes detailed information about SQL statements in its logs. This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections. The debug information includes the text of SQL statements that you submit through the proxy. Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive information that appears in the logs.
	DebugLogging pulumi.BoolPtrInput
	// The kinds of databases that the proxy can connect to. This value determines which database network protocol the proxy recognizes when it interprets network traffic to and from the database. The engine family applies to MySQL and PostgreSQL for both RDS and Aurora. Valid values are `MYSQL` and `POSTGRESQL`.
	EngineFamily pulumi.StringInput
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this value higher or lower than the connection timeout limit for the associated database.
	IdleClientTimeout pulumi.IntPtrInput
	// The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	Name pulumi.StringPtrInput
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	RequireTls pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	RoleArn pulumi.StringInput
	// A mapping of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyArgs)(nil)).Elem()
}

type ProxyInput interface {
	pulumi.Input

	ToProxyOutput() ProxyOutput
	ToProxyOutputWithContext(ctx context.Context) ProxyOutput
}

func (*Proxy) ElementType() reflect.Type {
	return reflect.TypeOf((*Proxy)(nil))
}

func (i *Proxy) ToProxyOutput() ProxyOutput {
	return i.ToProxyOutputWithContext(context.Background())
}

func (i *Proxy) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyOutput)
}

func (i *Proxy) ToProxyPtrOutput() ProxyPtrOutput {
	return i.ToProxyPtrOutputWithContext(context.Background())
}

func (i *Proxy) ToProxyPtrOutputWithContext(ctx context.Context) ProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyPtrOutput)
}

type ProxyPtrInput interface {
	pulumi.Input

	ToProxyPtrOutput() ProxyPtrOutput
	ToProxyPtrOutputWithContext(ctx context.Context) ProxyPtrOutput
}

type proxyPtrType ProxyArgs

func (*proxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil))
}

func (i *proxyPtrType) ToProxyPtrOutput() ProxyPtrOutput {
	return i.ToProxyPtrOutputWithContext(context.Background())
}

func (i *proxyPtrType) ToProxyPtrOutputWithContext(ctx context.Context) ProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyPtrOutput)
}

// ProxyArrayInput is an input type that accepts ProxyArray and ProxyArrayOutput values.
// You can construct a concrete instance of `ProxyArrayInput` via:
//
//          ProxyArray{ ProxyArgs{...} }
type ProxyArrayInput interface {
	pulumi.Input

	ToProxyArrayOutput() ProxyArrayOutput
	ToProxyArrayOutputWithContext(context.Context) ProxyArrayOutput
}

type ProxyArray []ProxyInput

func (ProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Proxy)(nil)).Elem()
}

func (i ProxyArray) ToProxyArrayOutput() ProxyArrayOutput {
	return i.ToProxyArrayOutputWithContext(context.Background())
}

func (i ProxyArray) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyArrayOutput)
}

// ProxyMapInput is an input type that accepts ProxyMap and ProxyMapOutput values.
// You can construct a concrete instance of `ProxyMapInput` via:
//
//          ProxyMap{ "key": ProxyArgs{...} }
type ProxyMapInput interface {
	pulumi.Input

	ToProxyMapOutput() ProxyMapOutput
	ToProxyMapOutputWithContext(context.Context) ProxyMapOutput
}

type ProxyMap map[string]ProxyInput

func (ProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Proxy)(nil)).Elem()
}

func (i ProxyMap) ToProxyMapOutput() ProxyMapOutput {
	return i.ToProxyMapOutputWithContext(context.Background())
}

func (i ProxyMap) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyMapOutput)
}

type ProxyOutput struct{ *pulumi.OutputState }

func (ProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Proxy)(nil))
}

func (o ProxyOutput) ToProxyOutput() ProxyOutput {
	return o
}

func (o ProxyOutput) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return o
}

func (o ProxyOutput) ToProxyPtrOutput() ProxyPtrOutput {
	return o.ToProxyPtrOutputWithContext(context.Background())
}

func (o ProxyOutput) ToProxyPtrOutputWithContext(ctx context.Context) ProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Proxy) *Proxy {
		return &v
	}).(ProxyPtrOutput)
}

type ProxyPtrOutput struct{ *pulumi.OutputState }

func (ProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil))
}

func (o ProxyPtrOutput) ToProxyPtrOutput() ProxyPtrOutput {
	return o
}

func (o ProxyPtrOutput) ToProxyPtrOutputWithContext(ctx context.Context) ProxyPtrOutput {
	return o
}

func (o ProxyPtrOutput) Elem() ProxyOutput {
	return o.ApplyT(func(v *Proxy) Proxy {
		if v != nil {
			return *v
		}
		var ret Proxy
		return ret
	}).(ProxyOutput)
}

type ProxyArrayOutput struct{ *pulumi.OutputState }

func (ProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Proxy)(nil))
}

func (o ProxyArrayOutput) ToProxyArrayOutput() ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) Index(i pulumi.IntInput) ProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Proxy {
		return vs[0].([]Proxy)[vs[1].(int)]
	}).(ProxyOutput)
}

type ProxyMapOutput struct{ *pulumi.OutputState }

func (ProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Proxy)(nil))
}

func (o ProxyMapOutput) ToProxyMapOutput() ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) MapIndex(k pulumi.StringInput) ProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Proxy {
		return vs[0].(map[string]Proxy)[vs[1].(string)]
	}).(ProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyInput)(nil)).Elem(), &Proxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyPtrInput)(nil)).Elem(), &Proxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyArrayInput)(nil)).Elem(), ProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyMapInput)(nil)).Elem(), ProxyMap{})
	pulumi.RegisterOutputType(ProxyOutput{})
	pulumi.RegisterOutputType(ProxyPtrOutput{})
	pulumi.RegisterOutputType(ProxyArrayOutput{})
	pulumi.RegisterOutputType(ProxyMapOutput{})
}
