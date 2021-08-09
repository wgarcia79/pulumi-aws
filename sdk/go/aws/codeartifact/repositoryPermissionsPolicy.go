// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeArtifact Repostory Permissions Policy Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codeartifact"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 			Domain:        pulumi.String("example.com"),
// 			EncryptionKey: exampleKey.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleRepository, err := codeartifact.NewRepository(ctx, "exampleRepository", &codeartifact.RepositoryArgs{
// 			Repository: pulumi.String("example"),
// 			Domain:     exampleDomain.Domain,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codeartifact.NewRepositoryPermissionsPolicy(ctx, "exampleRepositoryPermissionsPolicy", &codeartifact.RepositoryPermissionsPolicyArgs{
// 			Repository: exampleRepository.Repository,
// 			Domain:     exampleDomain.Domain,
// 			PolicyDocument: exampleDomain.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Action\": \"codeartifact:CreateRepository\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": \"*\",\n", "            \"Resource\": \"", arn, "\"\n", "        }\n", "    ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
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
// CodeArtifact Repository Permissions Policies can be imported using the CodeArtifact Repository ARN, e.g.
//
// ```sh
//  $ pulumi import aws:codeartifact/repositoryPermissionsPolicy:RepositoryPermissionsPolicy example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
// ```
type RepositoryPermissionsPolicy struct {
	pulumi.CustomResourceState

	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringOutput `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringOutput `pulumi:"policyRevision"`
	// The name of the repository to set the resource policy on.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The ARN of the resource associated with the resource policy.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewRepositoryPermissionsPolicy registers a new resource with the given unique name, arguments, and options.
func NewRepositoryPermissionsPolicy(ctx *pulumi.Context,
	name string, args *RepositoryPermissionsPolicyArgs, opts ...pulumi.ResourceOption) (*RepositoryPermissionsPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource RepositoryPermissionsPolicy
	err := ctx.RegisterResource("aws:codeartifact/repositoryPermissionsPolicy:RepositoryPermissionsPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryPermissionsPolicy gets an existing RepositoryPermissionsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryPermissionsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryPermissionsPolicyState, opts ...pulumi.ResourceOption) (*RepositoryPermissionsPolicy, error) {
	var resource RepositoryPermissionsPolicy
	err := ctx.ReadResource("aws:codeartifact/repositoryPermissionsPolicy:RepositoryPermissionsPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryPermissionsPolicy resources.
type repositoryPermissionsPolicyState struct {
	// The name of the domain on which to set the resource policy.
	Domain *string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `pulumi:"policyRevision"`
	// The name of the repository to set the resource policy on.
	Repository *string `pulumi:"repository"`
	// The ARN of the resource associated with the resource policy.
	ResourceArn *string `pulumi:"resourceArn"`
}

type RepositoryPermissionsPolicyState struct {
	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringPtrInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringPtrInput
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringPtrInput
	// The name of the repository to set the resource policy on.
	Repository pulumi.StringPtrInput
	// The ARN of the resource associated with the resource policy.
	ResourceArn pulumi.StringPtrInput
}

func (RepositoryPermissionsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPermissionsPolicyState)(nil)).Elem()
}

type repositoryPermissionsPolicyArgs struct {
	// The name of the domain on which to set the resource policy.
	Domain string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument string `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `pulumi:"policyRevision"`
	// The name of the repository to set the resource policy on.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a RepositoryPermissionsPolicy resource.
type RepositoryPermissionsPolicyArgs struct {
	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringInput
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringPtrInput
	// The name of the repository to set the resource policy on.
	Repository pulumi.StringInput
}

func (RepositoryPermissionsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryPermissionsPolicyArgs)(nil)).Elem()
}

type RepositoryPermissionsPolicyInput interface {
	pulumi.Input

	ToRepositoryPermissionsPolicyOutput() RepositoryPermissionsPolicyOutput
	ToRepositoryPermissionsPolicyOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyOutput
}

func (*RepositoryPermissionsPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPermissionsPolicy)(nil))
}

func (i *RepositoryPermissionsPolicy) ToRepositoryPermissionsPolicyOutput() RepositoryPermissionsPolicyOutput {
	return i.ToRepositoryPermissionsPolicyOutputWithContext(context.Background())
}

func (i *RepositoryPermissionsPolicy) ToRepositoryPermissionsPolicyOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPermissionsPolicyOutput)
}

func (i *RepositoryPermissionsPolicy) ToRepositoryPermissionsPolicyPtrOutput() RepositoryPermissionsPolicyPtrOutput {
	return i.ToRepositoryPermissionsPolicyPtrOutputWithContext(context.Background())
}

func (i *RepositoryPermissionsPolicy) ToRepositoryPermissionsPolicyPtrOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPermissionsPolicyPtrOutput)
}

type RepositoryPermissionsPolicyPtrInput interface {
	pulumi.Input

	ToRepositoryPermissionsPolicyPtrOutput() RepositoryPermissionsPolicyPtrOutput
	ToRepositoryPermissionsPolicyPtrOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyPtrOutput
}

type repositoryPermissionsPolicyPtrType RepositoryPermissionsPolicyArgs

func (*repositoryPermissionsPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPermissionsPolicy)(nil))
}

func (i *repositoryPermissionsPolicyPtrType) ToRepositoryPermissionsPolicyPtrOutput() RepositoryPermissionsPolicyPtrOutput {
	return i.ToRepositoryPermissionsPolicyPtrOutputWithContext(context.Background())
}

func (i *repositoryPermissionsPolicyPtrType) ToRepositoryPermissionsPolicyPtrOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPermissionsPolicyPtrOutput)
}

// RepositoryPermissionsPolicyArrayInput is an input type that accepts RepositoryPermissionsPolicyArray and RepositoryPermissionsPolicyArrayOutput values.
// You can construct a concrete instance of `RepositoryPermissionsPolicyArrayInput` via:
//
//          RepositoryPermissionsPolicyArray{ RepositoryPermissionsPolicyArgs{...} }
type RepositoryPermissionsPolicyArrayInput interface {
	pulumi.Input

	ToRepositoryPermissionsPolicyArrayOutput() RepositoryPermissionsPolicyArrayOutput
	ToRepositoryPermissionsPolicyArrayOutputWithContext(context.Context) RepositoryPermissionsPolicyArrayOutput
}

type RepositoryPermissionsPolicyArray []RepositoryPermissionsPolicyInput

func (RepositoryPermissionsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryPermissionsPolicy)(nil)).Elem()
}

func (i RepositoryPermissionsPolicyArray) ToRepositoryPermissionsPolicyArrayOutput() RepositoryPermissionsPolicyArrayOutput {
	return i.ToRepositoryPermissionsPolicyArrayOutputWithContext(context.Background())
}

func (i RepositoryPermissionsPolicyArray) ToRepositoryPermissionsPolicyArrayOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPermissionsPolicyArrayOutput)
}

// RepositoryPermissionsPolicyMapInput is an input type that accepts RepositoryPermissionsPolicyMap and RepositoryPermissionsPolicyMapOutput values.
// You can construct a concrete instance of `RepositoryPermissionsPolicyMapInput` via:
//
//          RepositoryPermissionsPolicyMap{ "key": RepositoryPermissionsPolicyArgs{...} }
type RepositoryPermissionsPolicyMapInput interface {
	pulumi.Input

	ToRepositoryPermissionsPolicyMapOutput() RepositoryPermissionsPolicyMapOutput
	ToRepositoryPermissionsPolicyMapOutputWithContext(context.Context) RepositoryPermissionsPolicyMapOutput
}

type RepositoryPermissionsPolicyMap map[string]RepositoryPermissionsPolicyInput

func (RepositoryPermissionsPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryPermissionsPolicy)(nil)).Elem()
}

func (i RepositoryPermissionsPolicyMap) ToRepositoryPermissionsPolicyMapOutput() RepositoryPermissionsPolicyMapOutput {
	return i.ToRepositoryPermissionsPolicyMapOutputWithContext(context.Background())
}

func (i RepositoryPermissionsPolicyMap) ToRepositoryPermissionsPolicyMapOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPermissionsPolicyMapOutput)
}

type RepositoryPermissionsPolicyOutput struct{ *pulumi.OutputState }

func (RepositoryPermissionsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryPermissionsPolicy)(nil))
}

func (o RepositoryPermissionsPolicyOutput) ToRepositoryPermissionsPolicyOutput() RepositoryPermissionsPolicyOutput {
	return o
}

func (o RepositoryPermissionsPolicyOutput) ToRepositoryPermissionsPolicyOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyOutput {
	return o
}

func (o RepositoryPermissionsPolicyOutput) ToRepositoryPermissionsPolicyPtrOutput() RepositoryPermissionsPolicyPtrOutput {
	return o.ToRepositoryPermissionsPolicyPtrOutputWithContext(context.Background())
}

func (o RepositoryPermissionsPolicyOutput) ToRepositoryPermissionsPolicyPtrOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryPermissionsPolicy) *RepositoryPermissionsPolicy {
		return &v
	}).(RepositoryPermissionsPolicyPtrOutput)
}

type RepositoryPermissionsPolicyPtrOutput struct{ *pulumi.OutputState }

func (RepositoryPermissionsPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryPermissionsPolicy)(nil))
}

func (o RepositoryPermissionsPolicyPtrOutput) ToRepositoryPermissionsPolicyPtrOutput() RepositoryPermissionsPolicyPtrOutput {
	return o
}

func (o RepositoryPermissionsPolicyPtrOutput) ToRepositoryPermissionsPolicyPtrOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyPtrOutput {
	return o
}

func (o RepositoryPermissionsPolicyPtrOutput) Elem() RepositoryPermissionsPolicyOutput {
	return o.ApplyT(func(v *RepositoryPermissionsPolicy) RepositoryPermissionsPolicy {
		if v != nil {
			return *v
		}
		var ret RepositoryPermissionsPolicy
		return ret
	}).(RepositoryPermissionsPolicyOutput)
}

type RepositoryPermissionsPolicyArrayOutput struct{ *pulumi.OutputState }

func (RepositoryPermissionsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepositoryPermissionsPolicy)(nil))
}

func (o RepositoryPermissionsPolicyArrayOutput) ToRepositoryPermissionsPolicyArrayOutput() RepositoryPermissionsPolicyArrayOutput {
	return o
}

func (o RepositoryPermissionsPolicyArrayOutput) ToRepositoryPermissionsPolicyArrayOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyArrayOutput {
	return o
}

func (o RepositoryPermissionsPolicyArrayOutput) Index(i pulumi.IntInput) RepositoryPermissionsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepositoryPermissionsPolicy {
		return vs[0].([]RepositoryPermissionsPolicy)[vs[1].(int)]
	}).(RepositoryPermissionsPolicyOutput)
}

type RepositoryPermissionsPolicyMapOutput struct{ *pulumi.OutputState }

func (RepositoryPermissionsPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RepositoryPermissionsPolicy)(nil))
}

func (o RepositoryPermissionsPolicyMapOutput) ToRepositoryPermissionsPolicyMapOutput() RepositoryPermissionsPolicyMapOutput {
	return o
}

func (o RepositoryPermissionsPolicyMapOutput) ToRepositoryPermissionsPolicyMapOutputWithContext(ctx context.Context) RepositoryPermissionsPolicyMapOutput {
	return o
}

func (o RepositoryPermissionsPolicyMapOutput) MapIndex(k pulumi.StringInput) RepositoryPermissionsPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RepositoryPermissionsPolicy {
		return vs[0].(map[string]RepositoryPermissionsPolicy)[vs[1].(string)]
	}).(RepositoryPermissionsPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoryPermissionsPolicyOutput{})
	pulumi.RegisterOutputType(RepositoryPermissionsPolicyPtrOutput{})
	pulumi.RegisterOutputType(RepositoryPermissionsPolicyArrayOutput{})
	pulumi.RegisterOutputType(RepositoryPermissionsPolicyMapOutput{})
}
