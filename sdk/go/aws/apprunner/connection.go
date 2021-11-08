// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Runner Connection.
//
// > **NOTE:** After creation, you must complete the authentication handshake using the App Runner console.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apprunner"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apprunner.NewConnection(ctx, "example", &apprunner.ConnectionArgs{
// 			ConnectionName: pulumi.String("example"),
// 			ProviderType:   pulumi.String("GITHUB"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("example-apprunner-connection"),
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
// App Runner Connections can be imported by using the `connection_name`, e.g.,
//
// ```sh
//  $ pulumi import aws:apprunner/connection:Connection example example
// ```
type Connection struct {
	pulumi.CustomResourceState

	// ARN of the connection.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the connection.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// The source repository provider. Valid values: `GITHUB`.
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
	// The current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an `apprunner.Service` resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.ProviderType == nil {
		return nil, errors.New("invalid value for required argument 'ProviderType'")
	}
	var resource Connection
	err := ctx.RegisterResource("aws:apprunner/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("aws:apprunner/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// ARN of the connection.
	Arn *string `pulumi:"arn"`
	// Name of the connection.
	ConnectionName *string `pulumi:"connectionName"`
	// The source repository provider. Valid values: `GITHUB`.
	ProviderType *string `pulumi:"providerType"`
	// The current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an `apprunner.Service` resource.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ConnectionState struct {
	// ARN of the connection.
	Arn pulumi.StringPtrInput
	// Name of the connection.
	ConnectionName pulumi.StringPtrInput
	// The source repository provider. Valid values: `GITHUB`.
	ProviderType pulumi.StringPtrInput
	// The current state of the App Runner connection. When the state is `AVAILABLE`, you can use the connection to create an `apprunner.Service` resource.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Name of the connection.
	ConnectionName string `pulumi:"connectionName"`
	// The source repository provider. Valid values: `GITHUB`.
	ProviderType string `pulumi:"providerType"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Name of the connection.
	ConnectionName pulumi.StringInput
	// The source repository provider. Valid values: `GITHUB`.
	ProviderType pulumi.StringInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

type ConnectionPtrInput interface {
	pulumi.Input

	ToConnectionPtrOutput() ConnectionPtrOutput
	ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput
}

type connectionPtrType ConnectionArgs

func (*connectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (i *connectionPtrType) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *connectionPtrType) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//          ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//          ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o.ToConnectionPtrOutputWithContext(context.Background())
}

func (o ConnectionOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Connection) *Connection {
		return &v
	}).(ConnectionPtrOutput)
}

type ConnectionPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (o ConnectionPtrOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) Elem() ConnectionOutput {
	return o.ApplyT(func(v *Connection) Connection {
		if v != nil {
			return *v
		}
		var ret Connection
		return ret
	}).(ConnectionOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Connection)(nil))
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Connection {
		return vs[0].([]Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Connection)(nil))
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Connection {
		return vs[0].(map[string]Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionPtrInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionPtrOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
