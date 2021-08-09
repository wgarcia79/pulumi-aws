// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amplify Branch resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/amplify"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := amplify.NewApp(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = amplify.NewBranch(ctx, "master", &amplify.BranchArgs{
// 			AppId:      example.ID(),
// 			BranchName: pulumi.String("master"),
// 			Framework:  pulumi.String("React"),
// 			Stage:      pulumi.String("PRODUCTION"),
// 			EnvironmentVariables: pulumi.StringMap{
// 				"REACT_APP_API_SERVER": pulumi.String("https://api.example.com"),
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
// Amplify branch can be imported using `app_id` and `branch_name`, e.g.
//
// ```sh
//  $ pulumi import aws:amplify/branch:Branch master d2ypk4k47z8u6/master
// ```
type Branch struct {
	pulumi.CustomResourceState

	// The unique ID for an Amplify app.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for the branch.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of custom resources that are linked to this branch.
	AssociatedResources pulumi.StringArrayOutput `pulumi:"associatedResources"`
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn pulumi.StringPtrOutput `pulumi:"backendEnvironmentArn"`
	// The basic authorization credentials for the branch.
	BasicAuthCredentials pulumi.StringPtrOutput `pulumi:"basicAuthCredentials"`
	// The name for the branch.
	BranchName pulumi.StringOutput `pulumi:"branchName"`
	// The custom domains for the branch.
	CustomDomains pulumi.StringArrayOutput `pulumi:"customDomains"`
	// The description for the branch.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination branch if the branch is a pull request branch.
	DestinationBranch pulumi.StringOutput `pulumi:"destinationBranch"`
	// The display name for a branch. This is used as the default domain prefix.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Enables auto building for the branch.
	EnableAutoBuild pulumi.BoolPtrOutput `pulumi:"enableAutoBuild"`
	// Enables basic authorization for the branch.
	EnableBasicAuth pulumi.BoolPtrOutput `pulumi:"enableBasicAuth"`
	// Enables notifications for the branch.
	EnableNotification pulumi.BoolPtrOutput `pulumi:"enableNotification"`
	// Enables performance mode for the branch.
	EnablePerformanceMode pulumi.BoolPtrOutput `pulumi:"enablePerformanceMode"`
	// Enables pull request previews for this branch.
	EnablePullRequestPreview pulumi.BoolPtrOutput `pulumi:"enablePullRequestPreview"`
	// The environment variables for the branch.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// The framework for the branch.
	Framework pulumi.StringPtrOutput `pulumi:"framework"`
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName pulumi.StringPtrOutput `pulumi:"pullRequestEnvironmentName"`
	// The source branch if the branch is a pull request branch.
	SourceBranch pulumi.StringOutput `pulumi:"sourceBranch"`
	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage pulumi.StringPtrOutput `pulumi:"stage"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The content Time To Live (TTL) for the website in seconds.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
}

// NewBranch registers a new resource with the given unique name, arguments, and options.
func NewBranch(ctx *pulumi.Context,
	name string, args *BranchArgs, opts ...pulumi.ResourceOption) (*Branch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.BranchName == nil {
		return nil, errors.New("invalid value for required argument 'BranchName'")
	}
	var resource Branch
	err := ctx.RegisterResource("aws:amplify/branch:Branch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranch gets an existing Branch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchState, opts ...pulumi.ResourceOption) (*Branch, error) {
	var resource Branch
	err := ctx.ReadResource("aws:amplify/branch:Branch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Branch resources.
type branchState struct {
	// The unique ID for an Amplify app.
	AppId *string `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for the branch.
	Arn *string `pulumi:"arn"`
	// A list of custom resources that are linked to this branch.
	AssociatedResources []string `pulumi:"associatedResources"`
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn *string `pulumi:"backendEnvironmentArn"`
	// The basic authorization credentials for the branch.
	BasicAuthCredentials *string `pulumi:"basicAuthCredentials"`
	// The name for the branch.
	BranchName *string `pulumi:"branchName"`
	// The custom domains for the branch.
	CustomDomains []string `pulumi:"customDomains"`
	// The description for the branch.
	Description *string `pulumi:"description"`
	// The destination branch if the branch is a pull request branch.
	DestinationBranch *string `pulumi:"destinationBranch"`
	// The display name for a branch. This is used as the default domain prefix.
	DisplayName *string `pulumi:"displayName"`
	// Enables auto building for the branch.
	EnableAutoBuild *bool `pulumi:"enableAutoBuild"`
	// Enables basic authorization for the branch.
	EnableBasicAuth *bool `pulumi:"enableBasicAuth"`
	// Enables notifications for the branch.
	EnableNotification *bool `pulumi:"enableNotification"`
	// Enables performance mode for the branch.
	EnablePerformanceMode *bool `pulumi:"enablePerformanceMode"`
	// Enables pull request previews for this branch.
	EnablePullRequestPreview *bool `pulumi:"enablePullRequestPreview"`
	// The environment variables for the branch.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The framework for the branch.
	Framework *string `pulumi:"framework"`
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string `pulumi:"pullRequestEnvironmentName"`
	// The source branch if the branch is a pull request branch.
	SourceBranch *string `pulumi:"sourceBranch"`
	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage *string `pulumi:"stage"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The content Time To Live (TTL) for the website in seconds.
	Ttl *string `pulumi:"ttl"`
}

type BranchState struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the branch.
	Arn pulumi.StringPtrInput
	// A list of custom resources that are linked to this branch.
	AssociatedResources pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn pulumi.StringPtrInput
	// The basic authorization credentials for the branch.
	BasicAuthCredentials pulumi.StringPtrInput
	// The name for the branch.
	BranchName pulumi.StringPtrInput
	// The custom domains for the branch.
	CustomDomains pulumi.StringArrayInput
	// The description for the branch.
	Description pulumi.StringPtrInput
	// The destination branch if the branch is a pull request branch.
	DestinationBranch pulumi.StringPtrInput
	// The display name for a branch. This is used as the default domain prefix.
	DisplayName pulumi.StringPtrInput
	// Enables auto building for the branch.
	EnableAutoBuild pulumi.BoolPtrInput
	// Enables basic authorization for the branch.
	EnableBasicAuth pulumi.BoolPtrInput
	// Enables notifications for the branch.
	EnableNotification pulumi.BoolPtrInput
	// Enables performance mode for the branch.
	EnablePerformanceMode pulumi.BoolPtrInput
	// Enables pull request previews for this branch.
	EnablePullRequestPreview pulumi.BoolPtrInput
	// The environment variables for the branch.
	EnvironmentVariables pulumi.StringMapInput
	// The framework for the branch.
	Framework pulumi.StringPtrInput
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName pulumi.StringPtrInput
	// The source branch if the branch is a pull request branch.
	SourceBranch pulumi.StringPtrInput
	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The content Time To Live (TTL) for the website in seconds.
	Ttl pulumi.StringPtrInput
}

func (BranchState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchState)(nil)).Elem()
}

type branchArgs struct {
	// The unique ID for an Amplify app.
	AppId string `pulumi:"appId"`
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn *string `pulumi:"backendEnvironmentArn"`
	// The basic authorization credentials for the branch.
	BasicAuthCredentials *string `pulumi:"basicAuthCredentials"`
	// The name for the branch.
	BranchName string `pulumi:"branchName"`
	// The description for the branch.
	Description *string `pulumi:"description"`
	// The display name for a branch. This is used as the default domain prefix.
	DisplayName *string `pulumi:"displayName"`
	// Enables auto building for the branch.
	EnableAutoBuild *bool `pulumi:"enableAutoBuild"`
	// Enables basic authorization for the branch.
	EnableBasicAuth *bool `pulumi:"enableBasicAuth"`
	// Enables notifications for the branch.
	EnableNotification *bool `pulumi:"enableNotification"`
	// Enables performance mode for the branch.
	EnablePerformanceMode *bool `pulumi:"enablePerformanceMode"`
	// Enables pull request previews for this branch.
	EnablePullRequestPreview *bool `pulumi:"enablePullRequestPreview"`
	// The environment variables for the branch.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The framework for the branch.
	Framework *string `pulumi:"framework"`
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string `pulumi:"pullRequestEnvironmentName"`
	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage *string `pulumi:"stage"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The content Time To Live (TTL) for the website in seconds.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Branch resource.
type BranchArgs struct {
	// The unique ID for an Amplify app.
	AppId pulumi.StringInput
	// The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	BackendEnvironmentArn pulumi.StringPtrInput
	// The basic authorization credentials for the branch.
	BasicAuthCredentials pulumi.StringPtrInput
	// The name for the branch.
	BranchName pulumi.StringInput
	// The description for the branch.
	Description pulumi.StringPtrInput
	// The display name for a branch. This is used as the default domain prefix.
	DisplayName pulumi.StringPtrInput
	// Enables auto building for the branch.
	EnableAutoBuild pulumi.BoolPtrInput
	// Enables basic authorization for the branch.
	EnableBasicAuth pulumi.BoolPtrInput
	// Enables notifications for the branch.
	EnableNotification pulumi.BoolPtrInput
	// Enables performance mode for the branch.
	EnablePerformanceMode pulumi.BoolPtrInput
	// Enables pull request previews for this branch.
	EnablePullRequestPreview pulumi.BoolPtrInput
	// The environment variables for the branch.
	EnvironmentVariables pulumi.StringMapInput
	// The framework for the branch.
	Framework pulumi.StringPtrInput
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName pulumi.StringPtrInput
	// Describes the current stage for the branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The content Time To Live (TTL) for the website in seconds.
	Ttl pulumi.StringPtrInput
}

func (BranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchArgs)(nil)).Elem()
}

type BranchInput interface {
	pulumi.Input

	ToBranchOutput() BranchOutput
	ToBranchOutputWithContext(ctx context.Context) BranchOutput
}

func (*Branch) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil))
}

func (i *Branch) ToBranchOutput() BranchOutput {
	return i.ToBranchOutputWithContext(context.Background())
}

func (i *Branch) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchOutput)
}

func (i *Branch) ToBranchPtrOutput() BranchPtrOutput {
	return i.ToBranchPtrOutputWithContext(context.Background())
}

func (i *Branch) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPtrOutput)
}

type BranchPtrInput interface {
	pulumi.Input

	ToBranchPtrOutput() BranchPtrOutput
	ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput
}

type branchPtrType BranchArgs

func (*branchPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil))
}

func (i *branchPtrType) ToBranchPtrOutput() BranchPtrOutput {
	return i.ToBranchPtrOutputWithContext(context.Background())
}

func (i *branchPtrType) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchPtrOutput)
}

// BranchArrayInput is an input type that accepts BranchArray and BranchArrayOutput values.
// You can construct a concrete instance of `BranchArrayInput` via:
//
//          BranchArray{ BranchArgs{...} }
type BranchArrayInput interface {
	pulumi.Input

	ToBranchArrayOutput() BranchArrayOutput
	ToBranchArrayOutputWithContext(context.Context) BranchArrayOutput
}

type BranchArray []BranchInput

func (BranchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Branch)(nil)).Elem()
}

func (i BranchArray) ToBranchArrayOutput() BranchArrayOutput {
	return i.ToBranchArrayOutputWithContext(context.Background())
}

func (i BranchArray) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchArrayOutput)
}

// BranchMapInput is an input type that accepts BranchMap and BranchMapOutput values.
// You can construct a concrete instance of `BranchMapInput` via:
//
//          BranchMap{ "key": BranchArgs{...} }
type BranchMapInput interface {
	pulumi.Input

	ToBranchMapOutput() BranchMapOutput
	ToBranchMapOutputWithContext(context.Context) BranchMapOutput
}

type BranchMap map[string]BranchInput

func (BranchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Branch)(nil)).Elem()
}

func (i BranchMap) ToBranchMapOutput() BranchMapOutput {
	return i.ToBranchMapOutputWithContext(context.Background())
}

func (i BranchMap) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchMapOutput)
}

type BranchOutput struct{ *pulumi.OutputState }

func (BranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil))
}

func (o BranchOutput) ToBranchOutput() BranchOutput {
	return o
}

func (o BranchOutput) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return o
}

func (o BranchOutput) ToBranchPtrOutput() BranchPtrOutput {
	return o.ToBranchPtrOutputWithContext(context.Background())
}

func (o BranchOutput) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Branch) *Branch {
		return &v
	}).(BranchPtrOutput)
}

type BranchPtrOutput struct{ *pulumi.OutputState }

func (BranchPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil))
}

func (o BranchPtrOutput) ToBranchPtrOutput() BranchPtrOutput {
	return o
}

func (o BranchPtrOutput) ToBranchPtrOutputWithContext(ctx context.Context) BranchPtrOutput {
	return o
}

func (o BranchPtrOutput) Elem() BranchOutput {
	return o.ApplyT(func(v *Branch) Branch {
		if v != nil {
			return *v
		}
		var ret Branch
		return ret
	}).(BranchOutput)
}

type BranchArrayOutput struct{ *pulumi.OutputState }

func (BranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Branch)(nil))
}

func (o BranchArrayOutput) ToBranchArrayOutput() BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) Index(i pulumi.IntInput) BranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Branch {
		return vs[0].([]Branch)[vs[1].(int)]
	}).(BranchOutput)
}

type BranchMapOutput struct{ *pulumi.OutputState }

func (BranchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Branch)(nil))
}

func (o BranchMapOutput) ToBranchMapOutput() BranchMapOutput {
	return o
}

func (o BranchMapOutput) ToBranchMapOutputWithContext(ctx context.Context) BranchMapOutput {
	return o
}

func (o BranchMapOutput) MapIndex(k pulumi.StringInput) BranchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Branch {
		return vs[0].(map[string]Branch)[vs[1].(string)]
	}).(BranchOutput)
}

func init() {
	pulumi.RegisterOutputType(BranchOutput{})
	pulumi.RegisterOutputType(BranchPtrOutput{})
	pulumi.RegisterOutputType(BranchArrayOutput{})
	pulumi.RegisterOutputType(BranchMapOutput{})
}
