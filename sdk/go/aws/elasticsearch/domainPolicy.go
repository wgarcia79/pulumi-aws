// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows setting policy to an Elasticsearch domain while referencing domain attributes (e.g., ARN)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticsearch"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := elasticsearch.NewDomain(ctx, "example", &elasticsearch.DomainArgs{
// 			ElasticsearchVersion: pulumi.String("2.3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elasticsearch.NewDomainPolicy(ctx, "main", &elasticsearch.DomainPolicyArgs{
// 			DomainName: example.DomainName,
// 			AccessPolicies: example.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Action\": \"es:*\",\n", "            \"Principal\": \"*\",\n", "            \"Effect\": \"Allow\",\n", "            \"Condition\": {\n", "                \"IpAddress\": {\"aws:SourceIp\": \"127.0.0.1/32\"}\n", "            },\n", "            \"Resource\": \"", arn, "/*\"\n", "        }\n", "    ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DomainPolicy struct {
	pulumi.CustomResourceState

	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.StringOutput `pulumi:"accessPolicies"`
	// Name of the domain.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewDomainPolicy registers a new resource with the given unique name, arguments, and options.
func NewDomainPolicy(ctx *pulumi.Context,
	name string, args *DomainPolicyArgs, opts ...pulumi.ResourceOption) (*DomainPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicies == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicies'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource DomainPolicy
	err := ctx.RegisterResource("aws:elasticsearch/domainPolicy:DomainPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainPolicy gets an existing DomainPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainPolicyState, opts ...pulumi.ResourceOption) (*DomainPolicy, error) {
	var resource DomainPolicy
	err := ctx.ReadResource("aws:elasticsearch/domainPolicy:DomainPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainPolicy resources.
type domainPolicyState struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies interface{} `pulumi:"accessPolicies"`
	// Name of the domain.
	DomainName *string `pulumi:"domainName"`
}

type DomainPolicyState struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.Input
	// Name of the domain.
	DomainName pulumi.StringPtrInput
}

func (DomainPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainPolicyState)(nil)).Elem()
}

type domainPolicyArgs struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies interface{} `pulumi:"accessPolicies"`
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
}

// The set of arguments for constructing a DomainPolicy resource.
type DomainPolicyArgs struct {
	// IAM policy document specifying the access policies for the domain
	AccessPolicies pulumi.Input
	// Name of the domain.
	DomainName pulumi.StringInput
}

func (DomainPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainPolicyArgs)(nil)).Elem()
}

type DomainPolicyInput interface {
	pulumi.Input

	ToDomainPolicyOutput() DomainPolicyOutput
	ToDomainPolicyOutputWithContext(ctx context.Context) DomainPolicyOutput
}

func (*DomainPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPolicy)(nil))
}

func (i *DomainPolicy) ToDomainPolicyOutput() DomainPolicyOutput {
	return i.ToDomainPolicyOutputWithContext(context.Background())
}

func (i *DomainPolicy) ToDomainPolicyOutputWithContext(ctx context.Context) DomainPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPolicyOutput)
}

func (i *DomainPolicy) ToDomainPolicyPtrOutput() DomainPolicyPtrOutput {
	return i.ToDomainPolicyPtrOutputWithContext(context.Background())
}

func (i *DomainPolicy) ToDomainPolicyPtrOutputWithContext(ctx context.Context) DomainPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPolicyPtrOutput)
}

type DomainPolicyPtrInput interface {
	pulumi.Input

	ToDomainPolicyPtrOutput() DomainPolicyPtrOutput
	ToDomainPolicyPtrOutputWithContext(ctx context.Context) DomainPolicyPtrOutput
}

type domainPolicyPtrType DomainPolicyArgs

func (*domainPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPolicy)(nil))
}

func (i *domainPolicyPtrType) ToDomainPolicyPtrOutput() DomainPolicyPtrOutput {
	return i.ToDomainPolicyPtrOutputWithContext(context.Background())
}

func (i *domainPolicyPtrType) ToDomainPolicyPtrOutputWithContext(ctx context.Context) DomainPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPolicyPtrOutput)
}

// DomainPolicyArrayInput is an input type that accepts DomainPolicyArray and DomainPolicyArrayOutput values.
// You can construct a concrete instance of `DomainPolicyArrayInput` via:
//
//          DomainPolicyArray{ DomainPolicyArgs{...} }
type DomainPolicyArrayInput interface {
	pulumi.Input

	ToDomainPolicyArrayOutput() DomainPolicyArrayOutput
	ToDomainPolicyArrayOutputWithContext(context.Context) DomainPolicyArrayOutput
}

type DomainPolicyArray []DomainPolicyInput

func (DomainPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainPolicy)(nil)).Elem()
}

func (i DomainPolicyArray) ToDomainPolicyArrayOutput() DomainPolicyArrayOutput {
	return i.ToDomainPolicyArrayOutputWithContext(context.Background())
}

func (i DomainPolicyArray) ToDomainPolicyArrayOutputWithContext(ctx context.Context) DomainPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPolicyArrayOutput)
}

// DomainPolicyMapInput is an input type that accepts DomainPolicyMap and DomainPolicyMapOutput values.
// You can construct a concrete instance of `DomainPolicyMapInput` via:
//
//          DomainPolicyMap{ "key": DomainPolicyArgs{...} }
type DomainPolicyMapInput interface {
	pulumi.Input

	ToDomainPolicyMapOutput() DomainPolicyMapOutput
	ToDomainPolicyMapOutputWithContext(context.Context) DomainPolicyMapOutput
}

type DomainPolicyMap map[string]DomainPolicyInput

func (DomainPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainPolicy)(nil)).Elem()
}

func (i DomainPolicyMap) ToDomainPolicyMapOutput() DomainPolicyMapOutput {
	return i.ToDomainPolicyMapOutputWithContext(context.Background())
}

func (i DomainPolicyMap) ToDomainPolicyMapOutputWithContext(ctx context.Context) DomainPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPolicyMapOutput)
}

type DomainPolicyOutput struct{ *pulumi.OutputState }

func (DomainPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPolicy)(nil))
}

func (o DomainPolicyOutput) ToDomainPolicyOutput() DomainPolicyOutput {
	return o
}

func (o DomainPolicyOutput) ToDomainPolicyOutputWithContext(ctx context.Context) DomainPolicyOutput {
	return o
}

func (o DomainPolicyOutput) ToDomainPolicyPtrOutput() DomainPolicyPtrOutput {
	return o.ToDomainPolicyPtrOutputWithContext(context.Background())
}

func (o DomainPolicyOutput) ToDomainPolicyPtrOutputWithContext(ctx context.Context) DomainPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainPolicy) *DomainPolicy {
		return &v
	}).(DomainPolicyPtrOutput)
}

type DomainPolicyPtrOutput struct{ *pulumi.OutputState }

func (DomainPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPolicy)(nil))
}

func (o DomainPolicyPtrOutput) ToDomainPolicyPtrOutput() DomainPolicyPtrOutput {
	return o
}

func (o DomainPolicyPtrOutput) ToDomainPolicyPtrOutputWithContext(ctx context.Context) DomainPolicyPtrOutput {
	return o
}

func (o DomainPolicyPtrOutput) Elem() DomainPolicyOutput {
	return o.ApplyT(func(v *DomainPolicy) DomainPolicy {
		if v != nil {
			return *v
		}
		var ret DomainPolicy
		return ret
	}).(DomainPolicyOutput)
}

type DomainPolicyArrayOutput struct{ *pulumi.OutputState }

func (DomainPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainPolicy)(nil))
}

func (o DomainPolicyArrayOutput) ToDomainPolicyArrayOutput() DomainPolicyArrayOutput {
	return o
}

func (o DomainPolicyArrayOutput) ToDomainPolicyArrayOutputWithContext(ctx context.Context) DomainPolicyArrayOutput {
	return o
}

func (o DomainPolicyArrayOutput) Index(i pulumi.IntInput) DomainPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainPolicy {
		return vs[0].([]DomainPolicy)[vs[1].(int)]
	}).(DomainPolicyOutput)
}

type DomainPolicyMapOutput struct{ *pulumi.OutputState }

func (DomainPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainPolicy)(nil))
}

func (o DomainPolicyMapOutput) ToDomainPolicyMapOutput() DomainPolicyMapOutput {
	return o
}

func (o DomainPolicyMapOutput) ToDomainPolicyMapOutputWithContext(ctx context.Context) DomainPolicyMapOutput {
	return o
}

func (o DomainPolicyMapOutput) MapIndex(k pulumi.StringInput) DomainPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainPolicy {
		return vs[0].(map[string]DomainPolicy)[vs[1].(string)]
	}).(DomainPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainPolicyInput)(nil)).Elem(), &DomainPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainPolicyPtrInput)(nil)).Elem(), &DomainPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainPolicyArrayInput)(nil)).Elem(), DomainPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainPolicyMapInput)(nil)).Elem(), DomainPolicyMap{})
	pulumi.RegisterOutputType(DomainPolicyOutput{})
	pulumi.RegisterOutputType(DomainPolicyPtrOutput{})
	pulumi.RegisterOutputType(DomainPolicyArrayOutput{})
	pulumi.RegisterOutputType(DomainPolicyMapOutput{})
}
