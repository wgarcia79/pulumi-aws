// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AppSync GraphQL API.
//
// ## Example Usage
// ### API Key Authentication
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("API_KEY"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### AWS Cognito User Pool Authentication
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("AMAZON_COGNITO_USER_POOLS"),
// 			UserPoolConfig: &appsync.GraphQLApiUserPoolConfigArgs{
// 				AwsRegion:     pulumi.Any(data.Aws_region.Current.Name),
// 				DefaultAction: pulumi.String("DENY"),
// 				UserPoolId:    pulumi.Any(aws_cognito_user_pool.Example.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### AWS IAM Authentication
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("AWS_IAM"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Schema
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("AWS_IAM"),
// 			Schema: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "schema {\n", "	query: Query\n", "}\n", "type Query {\n", "  test: Int\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### OpenID Connect Authentication
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("OPENID_CONNECT"),
// 			OpenidConnectConfig: &appsync.GraphQLApiOpenidConnectConfigArgs{
// 				Issuer: pulumi.String("https://example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Multiple Authentication Providers
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appsync.NewGraphQLApi(ctx, "example", &appsync.GraphQLApiArgs{
// 			AdditionalAuthenticationProviders: appsync.GraphQLApiAdditionalAuthenticationProviderArray{
// 				&appsync.GraphQLApiAdditionalAuthenticationProviderArgs{
// 					AuthenticationType: pulumi.String("AWS_IAM"),
// 				},
// 			},
// 			AuthenticationType: pulumi.String("API_KEY"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Enabling Logging
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\n", "            \"Service\": \"appsync.amazonaws.com\"\n", "        },\n", "        \"Action\": \"sts:AssumeRole\"\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
// 			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs"),
// 			Role:      exampleRole.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
// 			LogConfig: &appsync.GraphQLApiLogConfigArgs{
// 				CloudwatchLogsRoleArn: exampleRole.Arn,
// 				FieldLogLevel:         pulumi.String("ERROR"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Associate Web ACL (v2)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/appsync"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
// 			AuthenticationType: pulumi.String("API_KEY"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleWebAcl, err := wafv2.NewWebAcl(ctx, "exampleWebAcl", &wafv2.WebAclArgs{
// 			Description: pulumi.String("Example of a managed rule."),
// 			Scope:       pulumi.String("REGIONAL"),
// 			DefaultAction: &wafv2.WebAclDefaultActionArgs{
// 				Allow: nil,
// 			},
// 			Rules: wafv2.WebAclRuleArray{
// 				&wafv2.WebAclRuleArgs{
// 					Name:     pulumi.String("rule-1"),
// 					Priority: pulumi.Int(1),
// 					OverrideAction: &wafv2.WebAclRuleOverrideActionArgs{
// 						Block: []map[string]interface{}{
// 							nil,
// 						},
// 					},
// 					Statement: &wafv2.WebAclRuleStatementArgs{
// 						ManagedRuleGroupStatement: &wafv2.WebAclRuleStatementManagedRuleGroupStatementArgs{
// 							Name:       pulumi.String("AWSManagedRulesCommonRuleSet"),
// 							VendorName: pulumi.String("AWS"),
// 						},
// 					},
// 					VisibilityConfig: &wafv2.WebAclRuleVisibilityConfigArgs{
// 						CloudwatchMetricsEnabled: pulumi.Bool(false),
// 						MetricName:               pulumi.String("friendly-rule-metric-name"),
// 						SampledRequestsEnabled:   pulumi.Bool(false),
// 					},
// 				},
// 			},
// 			VisibilityConfig: &wafv2.WebAclVisibilityConfigArgs{
// 				CloudwatchMetricsEnabled: pulumi.Bool(false),
// 				MetricName:               pulumi.String("friendly-metric-name"),
// 				SampledRequestsEnabled:   pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = wafv2.NewWebAclAssociation(ctx, "exampleWebAclAssociation", &wafv2.WebAclAssociationArgs{
// 			ResourceArn: exampleGraphQLApi.Arn,
// 			WebAclArn:   exampleWebAcl.Arn,
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
// AppSync GraphQL API can be imported using the GraphQL API ID, e.g.
//
// ```sh
//  $ pulumi import aws:appsync/graphQLApi:GraphQLApi example 0123456789
// ```
type GraphQLApi struct {
	pulumi.CustomResourceState

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayOutput `pulumi:"additionalAuthenticationProviders"`
	// The ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrOutput `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringOutput `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrOutput `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapOutput `pulumi:"uris"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrOutput `pulumi:"userPoolConfig"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrOutput `pulumi:"xrayEnabled"`
}

// NewGraphQLApi registers a new resource with the given unique name, arguments, and options.
func NewGraphQLApi(ctx *pulumi.Context,
	name string, args *GraphQLApiArgs, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthenticationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationType'")
	}
	var resource GraphQLApi
	err := ctx.RegisterResource("aws:appsync/graphQLApi:GraphQLApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLApi gets an existing GraphQLApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLApiState, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	var resource GraphQLApi
	err := ctx.ReadResource("aws:appsync/graphQLApi:GraphQLApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLApi resources.
type graphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// The ARN
	Arn *string `pulumi:"arn"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType *string `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *GraphQLApiLogConfig `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris map[string]string `pulumi:"uris"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `pulumi:"xrayEnabled"`
}

type GraphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayInput
	// The ARN
	Arn pulumi.StringPtrInput
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringPtrInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrInput
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrInput
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapInput
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrInput
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrInput
}

func (GraphQLApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiState)(nil)).Elem()
}

type graphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType string `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *GraphQLApiLogConfig `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `pulumi:"xrayEnabled"`
}

// The set of arguments for constructing a GraphQLApi resource.
type GraphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders GraphQLApiAdditionalAuthenticationProviderArrayInput
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig GraphQLApiLogConfigPtrInput
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig GraphQLApiOpenidConnectConfigPtrInput
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig GraphQLApiUserPoolConfigPtrInput
	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled pulumi.BoolPtrInput
}

func (GraphQLApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiArgs)(nil)).Elem()
}

type GraphQLApiInput interface {
	pulumi.Input

	ToGraphQLApiOutput() GraphQLApiOutput
	ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput
}

func (*GraphQLApi) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphQLApi)(nil))
}

func (i *GraphQLApi) ToGraphQLApiOutput() GraphQLApiOutput {
	return i.ToGraphQLApiOutputWithContext(context.Background())
}

func (i *GraphQLApi) ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiOutput)
}

func (i *GraphQLApi) ToGraphQLApiPtrOutput() GraphQLApiPtrOutput {
	return i.ToGraphQLApiPtrOutputWithContext(context.Background())
}

func (i *GraphQLApi) ToGraphQLApiPtrOutputWithContext(ctx context.Context) GraphQLApiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiPtrOutput)
}

type GraphQLApiPtrInput interface {
	pulumi.Input

	ToGraphQLApiPtrOutput() GraphQLApiPtrOutput
	ToGraphQLApiPtrOutputWithContext(ctx context.Context) GraphQLApiPtrOutput
}

type graphQLApiPtrType GraphQLApiArgs

func (*graphQLApiPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApi)(nil))
}

func (i *graphQLApiPtrType) ToGraphQLApiPtrOutput() GraphQLApiPtrOutput {
	return i.ToGraphQLApiPtrOutputWithContext(context.Background())
}

func (i *graphQLApiPtrType) ToGraphQLApiPtrOutputWithContext(ctx context.Context) GraphQLApiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiPtrOutput)
}

// GraphQLApiArrayInput is an input type that accepts GraphQLApiArray and GraphQLApiArrayOutput values.
// You can construct a concrete instance of `GraphQLApiArrayInput` via:
//
//          GraphQLApiArray{ GraphQLApiArgs{...} }
type GraphQLApiArrayInput interface {
	pulumi.Input

	ToGraphQLApiArrayOutput() GraphQLApiArrayOutput
	ToGraphQLApiArrayOutputWithContext(context.Context) GraphQLApiArrayOutput
}

type GraphQLApiArray []GraphQLApiInput

func (GraphQLApiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GraphQLApi)(nil)).Elem()
}

func (i GraphQLApiArray) ToGraphQLApiArrayOutput() GraphQLApiArrayOutput {
	return i.ToGraphQLApiArrayOutputWithContext(context.Background())
}

func (i GraphQLApiArray) ToGraphQLApiArrayOutputWithContext(ctx context.Context) GraphQLApiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiArrayOutput)
}

// GraphQLApiMapInput is an input type that accepts GraphQLApiMap and GraphQLApiMapOutput values.
// You can construct a concrete instance of `GraphQLApiMapInput` via:
//
//          GraphQLApiMap{ "key": GraphQLApiArgs{...} }
type GraphQLApiMapInput interface {
	pulumi.Input

	ToGraphQLApiMapOutput() GraphQLApiMapOutput
	ToGraphQLApiMapOutputWithContext(context.Context) GraphQLApiMapOutput
}

type GraphQLApiMap map[string]GraphQLApiInput

func (GraphQLApiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GraphQLApi)(nil)).Elem()
}

func (i GraphQLApiMap) ToGraphQLApiMapOutput() GraphQLApiMapOutput {
	return i.ToGraphQLApiMapOutputWithContext(context.Background())
}

func (i GraphQLApiMap) ToGraphQLApiMapOutputWithContext(ctx context.Context) GraphQLApiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiMapOutput)
}

type GraphQLApiOutput struct{ *pulumi.OutputState }

func (GraphQLApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphQLApi)(nil))
}

func (o GraphQLApiOutput) ToGraphQLApiOutput() GraphQLApiOutput {
	return o
}

func (o GraphQLApiOutput) ToGraphQLApiOutputWithContext(ctx context.Context) GraphQLApiOutput {
	return o
}

func (o GraphQLApiOutput) ToGraphQLApiPtrOutput() GraphQLApiPtrOutput {
	return o.ToGraphQLApiPtrOutputWithContext(context.Background())
}

func (o GraphQLApiOutput) ToGraphQLApiPtrOutputWithContext(ctx context.Context) GraphQLApiPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GraphQLApi) *GraphQLApi {
		return &v
	}).(GraphQLApiPtrOutput)
}

type GraphQLApiPtrOutput struct{ *pulumi.OutputState }

func (GraphQLApiPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApi)(nil))
}

func (o GraphQLApiPtrOutput) ToGraphQLApiPtrOutput() GraphQLApiPtrOutput {
	return o
}

func (o GraphQLApiPtrOutput) ToGraphQLApiPtrOutputWithContext(ctx context.Context) GraphQLApiPtrOutput {
	return o
}

func (o GraphQLApiPtrOutput) Elem() GraphQLApiOutput {
	return o.ApplyT(func(v *GraphQLApi) GraphQLApi {
		if v != nil {
			return *v
		}
		var ret GraphQLApi
		return ret
	}).(GraphQLApiOutput)
}

type GraphQLApiArrayOutput struct{ *pulumi.OutputState }

func (GraphQLApiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphQLApi)(nil))
}

func (o GraphQLApiArrayOutput) ToGraphQLApiArrayOutput() GraphQLApiArrayOutput {
	return o
}

func (o GraphQLApiArrayOutput) ToGraphQLApiArrayOutputWithContext(ctx context.Context) GraphQLApiArrayOutput {
	return o
}

func (o GraphQLApiArrayOutput) Index(i pulumi.IntInput) GraphQLApiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphQLApi {
		return vs[0].([]GraphQLApi)[vs[1].(int)]
	}).(GraphQLApiOutput)
}

type GraphQLApiMapOutput struct{ *pulumi.OutputState }

func (GraphQLApiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GraphQLApi)(nil))
}

func (o GraphQLApiMapOutput) ToGraphQLApiMapOutput() GraphQLApiMapOutput {
	return o
}

func (o GraphQLApiMapOutput) ToGraphQLApiMapOutputWithContext(ctx context.Context) GraphQLApiMapOutput {
	return o
}

func (o GraphQLApiMapOutput) MapIndex(k pulumi.StringInput) GraphQLApiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GraphQLApi {
		return vs[0].(map[string]GraphQLApi)[vs[1].(string)]
	}).(GraphQLApiOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQLApiOutput{})
	pulumi.RegisterOutputType(GraphQLApiPtrOutput{})
	pulumi.RegisterOutputType(GraphQLApiArrayOutput{})
	pulumi.RegisterOutputType(GraphQLApiMapOutput{})
}
