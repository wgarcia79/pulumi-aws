// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeArtifact Domains Permissions Policy Resource.
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
// 		_, err = codeartifact.NewDomainPermissions(ctx, "test", &codeartifact.DomainPermissionsArgs{
// 			Domain: exampleDomain.Domain,
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
// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN, e.g.
//
// ```sh
//  $ pulumi import aws:codeartifact/domainPermissions:DomainPermissions example arn:aws:codeartifact:us-west-2:012345678912:domain/tf-acc-test-1928056699409417367
// ```
type DomainPermissions struct {
	pulumi.CustomResourceState

	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringOutput `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringOutput `pulumi:"policyRevision"`
	// The ARN of the resource associated with the resource policy.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewDomainPermissions registers a new resource with the given unique name, arguments, and options.
func NewDomainPermissions(ctx *pulumi.Context,
	name string, args *DomainPermissionsArgs, opts ...pulumi.ResourceOption) (*DomainPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	var resource DomainPermissions
	err := ctx.RegisterResource("aws:codeartifact/domainPermissions:DomainPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainPermissions gets an existing DomainPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainPermissionsState, opts ...pulumi.ResourceOption) (*DomainPermissions, error) {
	var resource DomainPermissions
	err := ctx.ReadResource("aws:codeartifact/domainPermissions:DomainPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainPermissions resources.
type domainPermissionsState struct {
	// The name of the domain on which to set the resource policy.
	Domain *string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `pulumi:"policyRevision"`
	// The ARN of the resource associated with the resource policy.
	ResourceArn *string `pulumi:"resourceArn"`
}

type DomainPermissionsState struct {
	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringPtrInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringPtrInput
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringPtrInput
	// The ARN of the resource associated with the resource policy.
	ResourceArn pulumi.StringPtrInput
}

func (DomainPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainPermissionsState)(nil)).Elem()
}

type domainPermissionsArgs struct {
	// The name of the domain on which to set the resource policy.
	Domain string `pulumi:"domain"`
	// The account number of the AWS account that owns the domain.
	DomainOwner *string `pulumi:"domainOwner"`
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument string `pulumi:"policyDocument"`
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision *string `pulumi:"policyRevision"`
}

// The set of arguments for constructing a DomainPermissions resource.
type DomainPermissionsArgs struct {
	// The name of the domain on which to set the resource policy.
	Domain pulumi.StringInput
	// The account number of the AWS account that owns the domain.
	DomainOwner pulumi.StringPtrInput
	// A JSON policy string to be set as the access control resource policy on the provided domain.
	PolicyDocument pulumi.StringInput
	// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
	PolicyRevision pulumi.StringPtrInput
}

func (DomainPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainPermissionsArgs)(nil)).Elem()
}

type DomainPermissionsInput interface {
	pulumi.Input

	ToDomainPermissionsOutput() DomainPermissionsOutput
	ToDomainPermissionsOutputWithContext(ctx context.Context) DomainPermissionsOutput
}

func (*DomainPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPermissions)(nil))
}

func (i *DomainPermissions) ToDomainPermissionsOutput() DomainPermissionsOutput {
	return i.ToDomainPermissionsOutputWithContext(context.Background())
}

func (i *DomainPermissions) ToDomainPermissionsOutputWithContext(ctx context.Context) DomainPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPermissionsOutput)
}

func (i *DomainPermissions) ToDomainPermissionsPtrOutput() DomainPermissionsPtrOutput {
	return i.ToDomainPermissionsPtrOutputWithContext(context.Background())
}

func (i *DomainPermissions) ToDomainPermissionsPtrOutputWithContext(ctx context.Context) DomainPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPermissionsPtrOutput)
}

type DomainPermissionsPtrInput interface {
	pulumi.Input

	ToDomainPermissionsPtrOutput() DomainPermissionsPtrOutput
	ToDomainPermissionsPtrOutputWithContext(ctx context.Context) DomainPermissionsPtrOutput
}

type domainPermissionsPtrType DomainPermissionsArgs

func (*domainPermissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPermissions)(nil))
}

func (i *domainPermissionsPtrType) ToDomainPermissionsPtrOutput() DomainPermissionsPtrOutput {
	return i.ToDomainPermissionsPtrOutputWithContext(context.Background())
}

func (i *domainPermissionsPtrType) ToDomainPermissionsPtrOutputWithContext(ctx context.Context) DomainPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPermissionsPtrOutput)
}

// DomainPermissionsArrayInput is an input type that accepts DomainPermissionsArray and DomainPermissionsArrayOutput values.
// You can construct a concrete instance of `DomainPermissionsArrayInput` via:
//
//          DomainPermissionsArray{ DomainPermissionsArgs{...} }
type DomainPermissionsArrayInput interface {
	pulumi.Input

	ToDomainPermissionsArrayOutput() DomainPermissionsArrayOutput
	ToDomainPermissionsArrayOutputWithContext(context.Context) DomainPermissionsArrayOutput
}

type DomainPermissionsArray []DomainPermissionsInput

func (DomainPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainPermissions)(nil)).Elem()
}

func (i DomainPermissionsArray) ToDomainPermissionsArrayOutput() DomainPermissionsArrayOutput {
	return i.ToDomainPermissionsArrayOutputWithContext(context.Background())
}

func (i DomainPermissionsArray) ToDomainPermissionsArrayOutputWithContext(ctx context.Context) DomainPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPermissionsArrayOutput)
}

// DomainPermissionsMapInput is an input type that accepts DomainPermissionsMap and DomainPermissionsMapOutput values.
// You can construct a concrete instance of `DomainPermissionsMapInput` via:
//
//          DomainPermissionsMap{ "key": DomainPermissionsArgs{...} }
type DomainPermissionsMapInput interface {
	pulumi.Input

	ToDomainPermissionsMapOutput() DomainPermissionsMapOutput
	ToDomainPermissionsMapOutputWithContext(context.Context) DomainPermissionsMapOutput
}

type DomainPermissionsMap map[string]DomainPermissionsInput

func (DomainPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainPermissions)(nil)).Elem()
}

func (i DomainPermissionsMap) ToDomainPermissionsMapOutput() DomainPermissionsMapOutput {
	return i.ToDomainPermissionsMapOutputWithContext(context.Background())
}

func (i DomainPermissionsMap) ToDomainPermissionsMapOutputWithContext(ctx context.Context) DomainPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPermissionsMapOutput)
}

type DomainPermissionsOutput struct{ *pulumi.OutputState }

func (DomainPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPermissions)(nil))
}

func (o DomainPermissionsOutput) ToDomainPermissionsOutput() DomainPermissionsOutput {
	return o
}

func (o DomainPermissionsOutput) ToDomainPermissionsOutputWithContext(ctx context.Context) DomainPermissionsOutput {
	return o
}

func (o DomainPermissionsOutput) ToDomainPermissionsPtrOutput() DomainPermissionsPtrOutput {
	return o.ToDomainPermissionsPtrOutputWithContext(context.Background())
}

func (o DomainPermissionsOutput) ToDomainPermissionsPtrOutputWithContext(ctx context.Context) DomainPermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainPermissions) *DomainPermissions {
		return &v
	}).(DomainPermissionsPtrOutput)
}

type DomainPermissionsPtrOutput struct{ *pulumi.OutputState }

func (DomainPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPermissions)(nil))
}

func (o DomainPermissionsPtrOutput) ToDomainPermissionsPtrOutput() DomainPermissionsPtrOutput {
	return o
}

func (o DomainPermissionsPtrOutput) ToDomainPermissionsPtrOutputWithContext(ctx context.Context) DomainPermissionsPtrOutput {
	return o
}

func (o DomainPermissionsPtrOutput) Elem() DomainPermissionsOutput {
	return o.ApplyT(func(v *DomainPermissions) DomainPermissions {
		if v != nil {
			return *v
		}
		var ret DomainPermissions
		return ret
	}).(DomainPermissionsOutput)
}

type DomainPermissionsArrayOutput struct{ *pulumi.OutputState }

func (DomainPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainPermissions)(nil))
}

func (o DomainPermissionsArrayOutput) ToDomainPermissionsArrayOutput() DomainPermissionsArrayOutput {
	return o
}

func (o DomainPermissionsArrayOutput) ToDomainPermissionsArrayOutputWithContext(ctx context.Context) DomainPermissionsArrayOutput {
	return o
}

func (o DomainPermissionsArrayOutput) Index(i pulumi.IntInput) DomainPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainPermissions {
		return vs[0].([]DomainPermissions)[vs[1].(int)]
	}).(DomainPermissionsOutput)
}

type DomainPermissionsMapOutput struct{ *pulumi.OutputState }

func (DomainPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainPermissions)(nil))
}

func (o DomainPermissionsMapOutput) ToDomainPermissionsMapOutput() DomainPermissionsMapOutput {
	return o
}

func (o DomainPermissionsMapOutput) ToDomainPermissionsMapOutputWithContext(ctx context.Context) DomainPermissionsMapOutput {
	return o
}

func (o DomainPermissionsMapOutput) MapIndex(k pulumi.StringInput) DomainPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainPermissions {
		return vs[0].(map[string]DomainPermissions)[vs[1].(string)]
	}).(DomainPermissionsOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainPermissionsOutput{})
	pulumi.RegisterOutputType(DomainPermissionsPtrOutput{})
	pulumi.RegisterOutputType(DomainPermissionsArrayOutput{})
	pulumi.RegisterOutputType(DomainPermissionsMapOutput{})
}
