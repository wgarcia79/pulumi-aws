// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Athena Named Query resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/athena"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		hogeBucket, err := s3.NewBucket(ctx, "hogeBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testKey, err := kms.NewKey(ctx, "testKey", &kms.KeyArgs{
// 			DeletionWindowInDays: pulumi.Int(7),
// 			Description:          pulumi.String("Athena KMS Key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testWorkgroup, err := athena.NewWorkgroup(ctx, "testWorkgroup", &athena.WorkgroupArgs{
// 			Configuration: &athena.WorkgroupConfigurationArgs{
// 				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
// 					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
// 						EncryptionOption: pulumi.String("SSE_KMS"),
// 						KmsKeyArn:        testKey.Arn,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		hogeDatabase, err := athena.NewDatabase(ctx, "hogeDatabase", &athena.DatabaseArgs{
// 			Name:   pulumi.String("users"),
// 			Bucket: hogeBucket.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = athena.NewNamedQuery(ctx, "foo", &athena.NamedQueryArgs{
// 			Workgroup: testWorkgroup.ID(),
// 			Database:  hogeDatabase.Name,
// 			Query: hogeDatabase.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v%v", "SELECT * FROM ", name, " limit 10;"), nil
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
// Athena Named Query can be imported using the query ID, e.g.
//
// ```sh
//  $ pulumi import aws:athena/namedQuery:NamedQuery example 0123456789
// ```
type NamedQuery struct {
	pulumi.CustomResourceState

	// The database to which the query belongs.
	Database pulumi.StringOutput `pulumi:"database"`
	// A brief explanation of the query. Maximum length of 1024.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The plain language name for the query. Maximum length of 128.
	Name pulumi.StringOutput `pulumi:"name"`
	// The text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query pulumi.StringOutput `pulumi:"query"`
	// The workgroup to which the query belongs. Defaults to `primary`
	Workgroup pulumi.StringPtrOutput `pulumi:"workgroup"`
}

// NewNamedQuery registers a new resource with the given unique name, arguments, and options.
func NewNamedQuery(ctx *pulumi.Context,
	name string, args *NamedQueryArgs, opts ...pulumi.ResourceOption) (*NamedQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	var resource NamedQuery
	err := ctx.RegisterResource("aws:athena/namedQuery:NamedQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedQuery gets an existing NamedQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedQueryState, opts ...pulumi.ResourceOption) (*NamedQuery, error) {
	var resource NamedQuery
	err := ctx.ReadResource("aws:athena/namedQuery:NamedQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedQuery resources.
type namedQueryState struct {
	// The database to which the query belongs.
	Database *string `pulumi:"database"`
	// A brief explanation of the query. Maximum length of 1024.
	Description *string `pulumi:"description"`
	// The plain language name for the query. Maximum length of 128.
	Name *string `pulumi:"name"`
	// The text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query *string `pulumi:"query"`
	// The workgroup to which the query belongs. Defaults to `primary`
	Workgroup *string `pulumi:"workgroup"`
}

type NamedQueryState struct {
	// The database to which the query belongs.
	Database pulumi.StringPtrInput
	// A brief explanation of the query. Maximum length of 1024.
	Description pulumi.StringPtrInput
	// The plain language name for the query. Maximum length of 128.
	Name pulumi.StringPtrInput
	// The text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query pulumi.StringPtrInput
	// The workgroup to which the query belongs. Defaults to `primary`
	Workgroup pulumi.StringPtrInput
}

func (NamedQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedQueryState)(nil)).Elem()
}

type namedQueryArgs struct {
	// The database to which the query belongs.
	Database string `pulumi:"database"`
	// A brief explanation of the query. Maximum length of 1024.
	Description *string `pulumi:"description"`
	// The plain language name for the query. Maximum length of 128.
	Name *string `pulumi:"name"`
	// The text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query string `pulumi:"query"`
	// The workgroup to which the query belongs. Defaults to `primary`
	Workgroup *string `pulumi:"workgroup"`
}

// The set of arguments for constructing a NamedQuery resource.
type NamedQueryArgs struct {
	// The database to which the query belongs.
	Database pulumi.StringInput
	// A brief explanation of the query. Maximum length of 1024.
	Description pulumi.StringPtrInput
	// The plain language name for the query. Maximum length of 128.
	Name pulumi.StringPtrInput
	// The text of the query itself. In other words, all query statements. Maximum length of 262144.
	Query pulumi.StringInput
	// The workgroup to which the query belongs. Defaults to `primary`
	Workgroup pulumi.StringPtrInput
}

func (NamedQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedQueryArgs)(nil)).Elem()
}

type NamedQueryInput interface {
	pulumi.Input

	ToNamedQueryOutput() NamedQueryOutput
	ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput
}

func (*NamedQuery) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedQuery)(nil))
}

func (i *NamedQuery) ToNamedQueryOutput() NamedQueryOutput {
	return i.ToNamedQueryOutputWithContext(context.Background())
}

func (i *NamedQuery) ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryOutput)
}

func (i *NamedQuery) ToNamedQueryPtrOutput() NamedQueryPtrOutput {
	return i.ToNamedQueryPtrOutputWithContext(context.Background())
}

func (i *NamedQuery) ToNamedQueryPtrOutputWithContext(ctx context.Context) NamedQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryPtrOutput)
}

type NamedQueryPtrInput interface {
	pulumi.Input

	ToNamedQueryPtrOutput() NamedQueryPtrOutput
	ToNamedQueryPtrOutputWithContext(ctx context.Context) NamedQueryPtrOutput
}

type namedQueryPtrType NamedQueryArgs

func (*namedQueryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedQuery)(nil))
}

func (i *namedQueryPtrType) ToNamedQueryPtrOutput() NamedQueryPtrOutput {
	return i.ToNamedQueryPtrOutputWithContext(context.Background())
}

func (i *namedQueryPtrType) ToNamedQueryPtrOutputWithContext(ctx context.Context) NamedQueryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryPtrOutput)
}

// NamedQueryArrayInput is an input type that accepts NamedQueryArray and NamedQueryArrayOutput values.
// You can construct a concrete instance of `NamedQueryArrayInput` via:
//
//          NamedQueryArray{ NamedQueryArgs{...} }
type NamedQueryArrayInput interface {
	pulumi.Input

	ToNamedQueryArrayOutput() NamedQueryArrayOutput
	ToNamedQueryArrayOutputWithContext(context.Context) NamedQueryArrayOutput
}

type NamedQueryArray []NamedQueryInput

func (NamedQueryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamedQuery)(nil)).Elem()
}

func (i NamedQueryArray) ToNamedQueryArrayOutput() NamedQueryArrayOutput {
	return i.ToNamedQueryArrayOutputWithContext(context.Background())
}

func (i NamedQueryArray) ToNamedQueryArrayOutputWithContext(ctx context.Context) NamedQueryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryArrayOutput)
}

// NamedQueryMapInput is an input type that accepts NamedQueryMap and NamedQueryMapOutput values.
// You can construct a concrete instance of `NamedQueryMapInput` via:
//
//          NamedQueryMap{ "key": NamedQueryArgs{...} }
type NamedQueryMapInput interface {
	pulumi.Input

	ToNamedQueryMapOutput() NamedQueryMapOutput
	ToNamedQueryMapOutputWithContext(context.Context) NamedQueryMapOutput
}

type NamedQueryMap map[string]NamedQueryInput

func (NamedQueryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamedQuery)(nil)).Elem()
}

func (i NamedQueryMap) ToNamedQueryMapOutput() NamedQueryMapOutput {
	return i.ToNamedQueryMapOutputWithContext(context.Background())
}

func (i NamedQueryMap) ToNamedQueryMapOutputWithContext(ctx context.Context) NamedQueryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedQueryMapOutput)
}

type NamedQueryOutput struct{ *pulumi.OutputState }

func (NamedQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedQuery)(nil))
}

func (o NamedQueryOutput) ToNamedQueryOutput() NamedQueryOutput {
	return o
}

func (o NamedQueryOutput) ToNamedQueryOutputWithContext(ctx context.Context) NamedQueryOutput {
	return o
}

func (o NamedQueryOutput) ToNamedQueryPtrOutput() NamedQueryPtrOutput {
	return o.ToNamedQueryPtrOutputWithContext(context.Background())
}

func (o NamedQueryOutput) ToNamedQueryPtrOutputWithContext(ctx context.Context) NamedQueryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NamedQuery) *NamedQuery {
		return &v
	}).(NamedQueryPtrOutput)
}

type NamedQueryPtrOutput struct{ *pulumi.OutputState }

func (NamedQueryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedQuery)(nil))
}

func (o NamedQueryPtrOutput) ToNamedQueryPtrOutput() NamedQueryPtrOutput {
	return o
}

func (o NamedQueryPtrOutput) ToNamedQueryPtrOutputWithContext(ctx context.Context) NamedQueryPtrOutput {
	return o
}

func (o NamedQueryPtrOutput) Elem() NamedQueryOutput {
	return o.ApplyT(func(v *NamedQuery) NamedQuery {
		if v != nil {
			return *v
		}
		var ret NamedQuery
		return ret
	}).(NamedQueryOutput)
}

type NamedQueryArrayOutput struct{ *pulumi.OutputState }

func (NamedQueryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamedQuery)(nil))
}

func (o NamedQueryArrayOutput) ToNamedQueryArrayOutput() NamedQueryArrayOutput {
	return o
}

func (o NamedQueryArrayOutput) ToNamedQueryArrayOutputWithContext(ctx context.Context) NamedQueryArrayOutput {
	return o
}

func (o NamedQueryArrayOutput) Index(i pulumi.IntInput) NamedQueryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamedQuery {
		return vs[0].([]NamedQuery)[vs[1].(int)]
	}).(NamedQueryOutput)
}

type NamedQueryMapOutput struct{ *pulumi.OutputState }

func (NamedQueryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NamedQuery)(nil))
}

func (o NamedQueryMapOutput) ToNamedQueryMapOutput() NamedQueryMapOutput {
	return o
}

func (o NamedQueryMapOutput) ToNamedQueryMapOutputWithContext(ctx context.Context) NamedQueryMapOutput {
	return o
}

func (o NamedQueryMapOutput) MapIndex(k pulumi.StringInput) NamedQueryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NamedQuery {
		return vs[0].(map[string]NamedQuery)[vs[1].(string)]
	}).(NamedQueryOutput)
}

func init() {
	pulumi.RegisterOutputType(NamedQueryOutput{})
	pulumi.RegisterOutputType(NamedQueryPtrOutput{})
	pulumi.RegisterOutputType(NamedQueryArrayOutput{})
	pulumi.RegisterOutputType(NamedQueryMapOutput{})
}
