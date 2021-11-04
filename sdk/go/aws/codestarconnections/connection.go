// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeStar Connection.
//
// > **NOTE:** The `codestarconnections.Connection` resource is created in the state `PENDING`. Authentication with the connection provider must be completed in the AWS Console.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codepipeline"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codestarconnections"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleConnection, err := codestarconnections.NewConnection(ctx, "exampleConnection", &codestarconnections.ConnectionArgs{
// 			ProviderType: pulumi.String("Bitbucket"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codepipeline.NewPipeline(ctx, "examplePipeline", &codepipeline.PipelineArgs{
// 			RoleArn:       pulumi.Any(aws_iam_role.Codepipeline_role.Arn),
// 			ArtifactStore: nil,
// 			Stages: codepipeline.PipelineStageArray{
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Source"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Source"),
// 							Category: pulumi.String("Source"),
// 							Owner:    pulumi.String("AWS"),
// 							Provider: pulumi.String("CodeStarSourceConnection"),
// 							Version:  pulumi.String("1"),
// 							OutputArtifacts: pulumi.StringArray{
// 								pulumi.String("source_output"),
// 							},
// 							Configuration: pulumi.StringMap{
// 								"ConnectionArn":    exampleConnection.Arn,
// 								"FullRepositoryId": pulumi.String("my-organization/test"),
// 								"BranchName":       pulumi.String("main"),
// 							},
// 						},
// 					},
// 				},
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Build"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						nil,
// 					},
// 				},
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Deploy"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						nil,
// 					},
// 				},
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
// CodeStar connections can be imported using the ARN, e.g.
//
// ```sh
//  $ pulumi import aws:codestarconnections/connection:Connection test-connection arn:aws:codestar-connections:us-west-1:0123456789:connection/79d4d357-a2ee-41e4-b350-2fe39ae59448
// ```
type Connection struct {
	pulumi.CustomResourceState

	// The codestar connection ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `providerType`
	HostArn pulumi.StringPtrOutput `pulumi:"hostArn"`
	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `providerType` will create a new resource. Conflicts with `hostArn`
	ProviderType pulumi.StringOutput `pulumi:"providerType"`
	// Map of key-value resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		args = &ConnectionArgs{}
	}

	var resource Connection
	err := ctx.RegisterResource("aws:codestarconnections/connection:Connection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:codestarconnections/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// The codestar connection ARN.
	Arn *string `pulumi:"arn"`
	// The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `providerType`
	HostArn *string `pulumi:"hostArn"`
	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name *string `pulumi:"name"`
	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `providerType` will create a new resource. Conflicts with `hostArn`
	ProviderType *string `pulumi:"providerType"`
	// Map of key-value resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ConnectionState struct {
	// The codestar connection ARN.
	Arn pulumi.StringPtrInput
	// The codestar connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
	ConnectionStatus pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `providerType`
	HostArn pulumi.StringPtrInput
	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name pulumi.StringPtrInput
	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `providerType` will create a new resource. Conflicts with `hostArn`
	ProviderType pulumi.StringPtrInput
	// Map of key-value resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `providerType`
	HostArn *string `pulumi:"hostArn"`
	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name *string `pulumi:"name"`
	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `providerType` will create a new resource. Conflicts with `hostArn`
	ProviderType *string `pulumi:"providerType"`
	// Map of key-value resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// The Amazon Resource Name (ARN) of the host associated with the connection. Conflicts with `providerType`
	HostArn pulumi.StringPtrInput
	// The name of the connection to be created. The name must be unique in the calling AWS account. Changing `name` will create a new resource.
	Name pulumi.StringPtrInput
	// The name of the external provider where your third-party code repository is configured. Valid values are `Bitbucket`, `GitHub` or `GitHubEnterpriseServer`. Changing `providerType` will create a new resource. Conflicts with `hostArn`
	ProviderType pulumi.StringPtrInput
	// Map of key-value resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
