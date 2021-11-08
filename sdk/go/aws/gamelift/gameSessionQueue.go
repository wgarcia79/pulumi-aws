// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Gamelift Game Session Queue resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewGameSessionQueue(ctx, "test", &gamelift.GameSessionQueueArgs{
// 			Destinations: pulumi.StringArray{
// 				pulumi.Any(aws_gamelift_fleet.Us_west_2_fleet.Arn),
// 				pulumi.Any(aws_gamelift_fleet.Eu_central_1_fleet.Arn),
// 			},
// 			PlayerLatencyPolicies: gamelift.GameSessionQueuePlayerLatencyPolicyArray{
// 				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
// 					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(100),
// 					PolicyDurationSeconds:                      pulumi.Int(5),
// 				},
// 				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
// 					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(200),
// 				},
// 			},
// 			TimeoutInSeconds: pulumi.Int(60),
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
// Gamelift Game Session Queues can be imported by their `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:gamelift/gameSessionQueue:GameSessionQueue example example
// ```
type GameSessionQueue struct {
	pulumi.CustomResourceState

	// Game Session Queue ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayOutput `pulumi:"destinations"`
	// Name of the session queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayOutput `pulumi:"playerLatencyPolicies"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrOutput `pulumi:"timeoutInSeconds"`
}

// NewGameSessionQueue registers a new resource with the given unique name, arguments, and options.
func NewGameSessionQueue(ctx *pulumi.Context,
	name string, args *GameSessionQueueArgs, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	if args == nil {
		args = &GameSessionQueueArgs{}
	}

	var resource GameSessionQueue
	err := ctx.RegisterResource("aws:gamelift/gameSessionQueue:GameSessionQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameSessionQueue gets an existing GameSessionQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameSessionQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameSessionQueueState, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	var resource GameSessionQueue
	err := ctx.ReadResource("aws:gamelift/gameSessionQueue:GameSessionQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameSessionQueue resources.
type gameSessionQueueState struct {
	// Game Session Queue ARN.
	Arn *string `pulumi:"arn"`
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []string `pulumi:"destinations"`
	// Name of the session queue.
	Name *string `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
}

type GameSessionQueueState struct {
	// Game Session Queue ARN.
	Arn pulumi.StringPtrInput
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayInput
	// Name of the session queue.
	Name pulumi.StringPtrInput
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrInput
}

func (GameSessionQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueState)(nil)).Elem()
}

type gameSessionQueueArgs struct {
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []string `pulumi:"destinations"`
	// Name of the session queue.
	Name *string `pulumi:"name"`
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy `pulumi:"playerLatencyPolicies"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
}

// The set of arguments for constructing a GameSessionQueue resource.
type GameSessionQueueArgs struct {
	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations pulumi.StringArrayInput
	// Name of the session queue.
	Name pulumi.StringPtrInput
	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds pulumi.IntPtrInput
}

func (GameSessionQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueArgs)(nil)).Elem()
}

type GameSessionQueueInput interface {
	pulumi.Input

	ToGameSessionQueueOutput() GameSessionQueueOutput
	ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput
}

func (*GameSessionQueue) ElementType() reflect.Type {
	return reflect.TypeOf((*GameSessionQueue)(nil))
}

func (i *GameSessionQueue) ToGameSessionQueueOutput() GameSessionQueueOutput {
	return i.ToGameSessionQueueOutputWithContext(context.Background())
}

func (i *GameSessionQueue) ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueueOutput)
}

func (i *GameSessionQueue) ToGameSessionQueuePtrOutput() GameSessionQueuePtrOutput {
	return i.ToGameSessionQueuePtrOutputWithContext(context.Background())
}

func (i *GameSessionQueue) ToGameSessionQueuePtrOutputWithContext(ctx context.Context) GameSessionQueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueuePtrOutput)
}

type GameSessionQueuePtrInput interface {
	pulumi.Input

	ToGameSessionQueuePtrOutput() GameSessionQueuePtrOutput
	ToGameSessionQueuePtrOutputWithContext(ctx context.Context) GameSessionQueuePtrOutput
}

type gameSessionQueuePtrType GameSessionQueueArgs

func (*gameSessionQueuePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GameSessionQueue)(nil))
}

func (i *gameSessionQueuePtrType) ToGameSessionQueuePtrOutput() GameSessionQueuePtrOutput {
	return i.ToGameSessionQueuePtrOutputWithContext(context.Background())
}

func (i *gameSessionQueuePtrType) ToGameSessionQueuePtrOutputWithContext(ctx context.Context) GameSessionQueuePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueuePtrOutput)
}

// GameSessionQueueArrayInput is an input type that accepts GameSessionQueueArray and GameSessionQueueArrayOutput values.
// You can construct a concrete instance of `GameSessionQueueArrayInput` via:
//
//          GameSessionQueueArray{ GameSessionQueueArgs{...} }
type GameSessionQueueArrayInput interface {
	pulumi.Input

	ToGameSessionQueueArrayOutput() GameSessionQueueArrayOutput
	ToGameSessionQueueArrayOutputWithContext(context.Context) GameSessionQueueArrayOutput
}

type GameSessionQueueArray []GameSessionQueueInput

func (GameSessionQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GameSessionQueue)(nil)).Elem()
}

func (i GameSessionQueueArray) ToGameSessionQueueArrayOutput() GameSessionQueueArrayOutput {
	return i.ToGameSessionQueueArrayOutputWithContext(context.Background())
}

func (i GameSessionQueueArray) ToGameSessionQueueArrayOutputWithContext(ctx context.Context) GameSessionQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueueArrayOutput)
}

// GameSessionQueueMapInput is an input type that accepts GameSessionQueueMap and GameSessionQueueMapOutput values.
// You can construct a concrete instance of `GameSessionQueueMapInput` via:
//
//          GameSessionQueueMap{ "key": GameSessionQueueArgs{...} }
type GameSessionQueueMapInput interface {
	pulumi.Input

	ToGameSessionQueueMapOutput() GameSessionQueueMapOutput
	ToGameSessionQueueMapOutputWithContext(context.Context) GameSessionQueueMapOutput
}

type GameSessionQueueMap map[string]GameSessionQueueInput

func (GameSessionQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GameSessionQueue)(nil)).Elem()
}

func (i GameSessionQueueMap) ToGameSessionQueueMapOutput() GameSessionQueueMapOutput {
	return i.ToGameSessionQueueMapOutputWithContext(context.Background())
}

func (i GameSessionQueueMap) ToGameSessionQueueMapOutputWithContext(ctx context.Context) GameSessionQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueueMapOutput)
}

type GameSessionQueueOutput struct{ *pulumi.OutputState }

func (GameSessionQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GameSessionQueue)(nil))
}

func (o GameSessionQueueOutput) ToGameSessionQueueOutput() GameSessionQueueOutput {
	return o
}

func (o GameSessionQueueOutput) ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput {
	return o
}

func (o GameSessionQueueOutput) ToGameSessionQueuePtrOutput() GameSessionQueuePtrOutput {
	return o.ToGameSessionQueuePtrOutputWithContext(context.Background())
}

func (o GameSessionQueueOutput) ToGameSessionQueuePtrOutputWithContext(ctx context.Context) GameSessionQueuePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GameSessionQueue) *GameSessionQueue {
		return &v
	}).(GameSessionQueuePtrOutput)
}

type GameSessionQueuePtrOutput struct{ *pulumi.OutputState }

func (GameSessionQueuePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GameSessionQueue)(nil))
}

func (o GameSessionQueuePtrOutput) ToGameSessionQueuePtrOutput() GameSessionQueuePtrOutput {
	return o
}

func (o GameSessionQueuePtrOutput) ToGameSessionQueuePtrOutputWithContext(ctx context.Context) GameSessionQueuePtrOutput {
	return o
}

func (o GameSessionQueuePtrOutput) Elem() GameSessionQueueOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueue {
		if v != nil {
			return *v
		}
		var ret GameSessionQueue
		return ret
	}).(GameSessionQueueOutput)
}

type GameSessionQueueArrayOutput struct{ *pulumi.OutputState }

func (GameSessionQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GameSessionQueue)(nil))
}

func (o GameSessionQueueArrayOutput) ToGameSessionQueueArrayOutput() GameSessionQueueArrayOutput {
	return o
}

func (o GameSessionQueueArrayOutput) ToGameSessionQueueArrayOutputWithContext(ctx context.Context) GameSessionQueueArrayOutput {
	return o
}

func (o GameSessionQueueArrayOutput) Index(i pulumi.IntInput) GameSessionQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GameSessionQueue {
		return vs[0].([]GameSessionQueue)[vs[1].(int)]
	}).(GameSessionQueueOutput)
}

type GameSessionQueueMapOutput struct{ *pulumi.OutputState }

func (GameSessionQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GameSessionQueue)(nil))
}

func (o GameSessionQueueMapOutput) ToGameSessionQueueMapOutput() GameSessionQueueMapOutput {
	return o
}

func (o GameSessionQueueMapOutput) ToGameSessionQueueMapOutputWithContext(ctx context.Context) GameSessionQueueMapOutput {
	return o
}

func (o GameSessionQueueMapOutput) MapIndex(k pulumi.StringInput) GameSessionQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GameSessionQueue {
		return vs[0].(map[string]GameSessionQueue)[vs[1].(string)]
	}).(GameSessionQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GameSessionQueueInput)(nil)).Elem(), &GameSessionQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*GameSessionQueuePtrInput)(nil)).Elem(), &GameSessionQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*GameSessionQueueArrayInput)(nil)).Elem(), GameSessionQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GameSessionQueueMapInput)(nil)).Elem(), GameSessionQueueMap{})
	pulumi.RegisterOutputType(GameSessionQueueOutput{})
	pulumi.RegisterOutputType(GameSessionQueuePtrOutput{})
	pulumi.RegisterOutputType(GameSessionQueueArrayOutput{})
	pulumi.RegisterOutputType(GameSessionQueueMapOutput{})
}
