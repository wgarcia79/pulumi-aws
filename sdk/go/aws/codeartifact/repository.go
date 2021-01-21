// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeArtifact Repository Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codeartifact"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
// 			Description: pulumi.String("domain key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDomain, err := codeartifact.NewDomain(ctx, "exampleDomain", &codeartifact.DomainArgs{
// 			Domain:        pulumi.String("example"),
// 			EncryptionKey: exampleKey.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("example"),
// 			Domain:     exampleDomain.Domain,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Upstream Repository
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codeartifact"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		upstream, err := codeartifact.NewRepository(ctx, "upstream", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("upstream"),
// 			Domain:     pulumi.Any(aws_codeartifact_domain.Test.Domain),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("example"),
// 			Domain:     pulumi.Any(aws_codeartifact_domain.Example.Domain),
// 			Upstreams: codeartifact.RepositoryUpstreamArray{
// 				&codeartifact.RepositoryUpstreamArgs{
// 					RepositoryName: upstream.Repository,
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
// ### With External Connection
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codeartifact"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codeartifact.NewRepository(ctx, "upstream", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("upstream"),
// 			Domain:     pulumi.Any(aws_codeartifact_domain.Test.Domain),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("example"),
// 			Domain:     pulumi.Any(aws_codeartifact_domain.Example.Domain),
// 			ExternalConnections: &codeartifact.RepositoryExternalConnectionsArgs{
// 				ExternalConnectionName: pulumi.String("public:npmjs"),
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
// CodeArtifact Repository can be imported using the CodeArtifact Repository ARN, e.g.
//
// ```sh
//  $ pulumi import aws:codeartifact/repository:Repository example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
// ```
type Repository struct {
	pulumi.CustomResourceState

	// The account number of the AWS account that manages the repository.
	AdministratorAccount pulumi.StringOutput `pulumi:"administratorAccount"`
	// The ARN of the repository.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain that contains the created repository.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringOutput `pulumi:"domainOwner"`
	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections RepositoryExternalConnectionsPtrOutput `pulumi:"externalConnections"`
	// The name of the repository to create.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams RepositoryUpstreamArrayOutput `pulumi:"upstreams"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource Repository
	err := ctx.RegisterResource("aws:codeartifact/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("aws:codeartifact/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// The account number of the AWS account that manages the repository.
	AdministratorAccount *string `pulumi:"administratorAccount"`
	// The ARN of the repository.
	Arn *string `pulumi:"arn"`
	// The description of the repository.
	Description *string `pulumi:"description"`
	// The domain that contains the created repository.
	Domain *string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections *RepositoryExternalConnections `pulumi:"externalConnections"`
	// The name of the repository to create.
	Repository *string `pulumi:"repository"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams []RepositoryUpstream `pulumi:"upstreams"`
}

type RepositoryState struct {
	// The account number of the AWS account that manages the repository.
	AdministratorAccount pulumi.StringPtrInput
	// The ARN of the repository.
	Arn pulumi.StringPtrInput
	// The description of the repository.
	Description pulumi.StringPtrInput
	// The domain that contains the created repository.
	Domain pulumi.StringPtrInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections RepositoryExternalConnectionsPtrInput
	// The name of the repository to create.
	Repository pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams RepositoryUpstreamArrayInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The description of the repository.
	Description *string `pulumi:"description"`
	// The domain that contains the created repository.
	Domain string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections *RepositoryExternalConnections `pulumi:"externalConnections"`
	// The name of the repository to create.
	Repository string `pulumi:"repository"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams []RepositoryUpstream `pulumi:"upstreams"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The description of the repository.
	Description pulumi.StringPtrInput
	// The domain that contains the created repository.
	Domain pulumi.StringInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections RepositoryExternalConnectionsPtrInput
	// The name of the repository to create.
	Repository pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams RepositoryUpstreamArrayInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

func (i *Repository) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPtrOutput)
}

type RepositoryPtrInput interface {
	pulumi.Input

	ToRepositoryPtrOutput() RepositoryPtrOutput
	ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput
}

type repositoryPtrType RepositoryArgs

func (*repositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil))
}

func (i *repositoryPtrType) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i *repositoryPtrType) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPtrOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//          RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Repository)(nil))
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//          RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Repository)(nil))
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct {
	*pulumi.OutputState
}

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o.ToRepositoryPtrOutputWithContext(context.Background())
}

func (o RepositoryOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o.ApplyT(func(v Repository) *Repository {
		return &v
	}).(RepositoryPtrOutput)
}

type RepositoryPtrOutput struct {
	*pulumi.OutputState
}

func (RepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil))
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Repository)(nil))
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Repository {
		return vs[0].([]Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Repository)(nil))
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Repository {
		return vs[0].(map[string]Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryPtrOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
