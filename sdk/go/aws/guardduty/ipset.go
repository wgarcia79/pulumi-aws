// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a GuardDuty IPSet.
//
// > **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/guardduty"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		primary, err := guardduty.NewDetector(ctx, "primary", &guardduty.DetectorArgs{
// 			Enable: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myIPSet, err := s3.NewBucketObject(ctx, "myIPSet", &s3.BucketObjectArgs{
// 			Acl:     pulumi.String("public-read"),
// 			Content: pulumi.String("10.0.0.0/8\n"),
// 			Bucket:  bucket.ID(),
// 			Key:     pulumi.String("MyIPSet"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = guardduty.NewIPSet(ctx, "example", &guardduty.IPSetArgs{
// 			Activate:   pulumi.Bool(true),
// 			DetectorId: primary.ID(),
// 			Format:     pulumi.String("TXT"),
// 			Location: pulumi.All(myIPSet.Bucket, myIPSet.Key).ApplyT(func(_args []interface{}) (string, error) {
// 				bucket := _args[0].(string)
// 				key := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v", "https://s3.amazonaws.com/", bucket, "/", key), nil
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
// GuardDuty IPSet can be imported using the the primary GuardDuty detector ID and IPSet ID, e.g.
//
// ```sh
//  $ pulumi import aws:guardduty/iPSet:IPSet MyIPSet 00b00fd5aecc0ab60a708659477e9617:123456789012
// ```
type IPSet struct {
	pulumi.CustomResourceState

	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolOutput `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty IPSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringOutput `pulumi:"detectorId"`
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringOutput `pulumi:"format"`
	// The URI of the file that contains the IPSet.
	Location pulumi.StringOutput `pulumi:"location"`
	// The friendly name to identify the IPSet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewIPSet registers a new resource with the given unique name, arguments, and options.
func NewIPSet(ctx *pulumi.Context,
	name string, args *IPSetArgs, opts ...pulumi.ResourceOption) (*IPSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activate == nil {
		return nil, errors.New("invalid value for required argument 'Activate'")
	}
	if args.DetectorId == nil {
		return nil, errors.New("invalid value for required argument 'DetectorId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource IPSet
	err := ctx.RegisterResource("aws:guardduty/iPSet:IPSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPSet gets an existing IPSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPSetState, opts ...pulumi.ResourceOption) (*IPSet, error) {
	var resource IPSet
	err := ctx.ReadResource("aws:guardduty/iPSet:IPSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPSet resources.
type ipsetState struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate *bool `pulumi:"activate"`
	// Amazon Resource Name (ARN) of the GuardDuty IPSet.
	Arn *string `pulumi:"arn"`
	// The detector ID of the GuardDuty.
	DetectorId *string `pulumi:"detectorId"`
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format *string `pulumi:"format"`
	// The URI of the file that contains the IPSet.
	Location *string `pulumi:"location"`
	// The friendly name to identify the IPSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type IPSetState struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of the GuardDuty IPSet.
	Arn pulumi.StringPtrInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringPtrInput
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringPtrInput
	// The URI of the file that contains the IPSet.
	Location pulumi.StringPtrInput
	// The friendly name to identify the IPSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (IPSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetState)(nil)).Elem()
}

type ipsetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate bool `pulumi:"activate"`
	// The detector ID of the GuardDuty.
	DetectorId string `pulumi:"detectorId"`
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format string `pulumi:"format"`
	// The URI of the file that contains the IPSet.
	Location string `pulumi:"location"`
	// The friendly name to identify the IPSet.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a IPSet resource.
type IPSetArgs struct {
	// Specifies whether GuardDuty is to start using the uploaded IPSet.
	Activate pulumi.BoolInput
	// The detector ID of the GuardDuty.
	DetectorId pulumi.StringInput
	// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
	Format pulumi.StringInput
	// The URI of the file that contains the IPSet.
	Location pulumi.StringInput
	// The friendly name to identify the IPSet.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (IPSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsetArgs)(nil)).Elem()
}

type IPSetInput interface {
	pulumi.Input

	ToIPSetOutput() IPSetOutput
	ToIPSetOutputWithContext(ctx context.Context) IPSetOutput
}

func (*IPSet) ElementType() reflect.Type {
	return reflect.TypeOf((*IPSet)(nil))
}

func (i *IPSet) ToIPSetOutput() IPSetOutput {
	return i.ToIPSetOutputWithContext(context.Background())
}

func (i *IPSet) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetOutput)
}

func (i *IPSet) ToIPSetPtrOutput() IPSetPtrOutput {
	return i.ToIPSetPtrOutputWithContext(context.Background())
}

func (i *IPSet) ToIPSetPtrOutputWithContext(ctx context.Context) IPSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetPtrOutput)
}

type IPSetPtrInput interface {
	pulumi.Input

	ToIPSetPtrOutput() IPSetPtrOutput
	ToIPSetPtrOutputWithContext(ctx context.Context) IPSetPtrOutput
}

type ipsetPtrType IPSetArgs

func (*ipsetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSet)(nil))
}

func (i *ipsetPtrType) ToIPSetPtrOutput() IPSetPtrOutput {
	return i.ToIPSetPtrOutputWithContext(context.Background())
}

func (i *ipsetPtrType) ToIPSetPtrOutputWithContext(ctx context.Context) IPSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetPtrOutput)
}

// IPSetArrayInput is an input type that accepts IPSetArray and IPSetArrayOutput values.
// You can construct a concrete instance of `IPSetArrayInput` via:
//
//          IPSetArray{ IPSetArgs{...} }
type IPSetArrayInput interface {
	pulumi.Input

	ToIPSetArrayOutput() IPSetArrayOutput
	ToIPSetArrayOutputWithContext(context.Context) IPSetArrayOutput
}

type IPSetArray []IPSetInput

func (IPSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IPSet)(nil))
}

func (i IPSetArray) ToIPSetArrayOutput() IPSetArrayOutput {
	return i.ToIPSetArrayOutputWithContext(context.Background())
}

func (i IPSetArray) ToIPSetArrayOutputWithContext(ctx context.Context) IPSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetArrayOutput)
}

// IPSetMapInput is an input type that accepts IPSetMap and IPSetMapOutput values.
// You can construct a concrete instance of `IPSetMapInput` via:
//
//          IPSetMap{ "key": IPSetArgs{...} }
type IPSetMapInput interface {
	pulumi.Input

	ToIPSetMapOutput() IPSetMapOutput
	ToIPSetMapOutputWithContext(context.Context) IPSetMapOutput
}

type IPSetMap map[string]IPSetInput

func (IPSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IPSet)(nil))
}

func (i IPSetMap) ToIPSetMapOutput() IPSetMapOutput {
	return i.ToIPSetMapOutputWithContext(context.Background())
}

func (i IPSetMap) ToIPSetMapOutputWithContext(ctx context.Context) IPSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPSetMapOutput)
}

type IPSetOutput struct {
	*pulumi.OutputState
}

func (IPSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPSet)(nil))
}

func (o IPSetOutput) ToIPSetOutput() IPSetOutput {
	return o
}

func (o IPSetOutput) ToIPSetOutputWithContext(ctx context.Context) IPSetOutput {
	return o
}

func (o IPSetOutput) ToIPSetPtrOutput() IPSetPtrOutput {
	return o.ToIPSetPtrOutputWithContext(context.Background())
}

func (o IPSetOutput) ToIPSetPtrOutputWithContext(ctx context.Context) IPSetPtrOutput {
	return o.ApplyT(func(v IPSet) *IPSet {
		return &v
	}).(IPSetPtrOutput)
}

type IPSetPtrOutput struct {
	*pulumi.OutputState
}

func (IPSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPSet)(nil))
}

func (o IPSetPtrOutput) ToIPSetPtrOutput() IPSetPtrOutput {
	return o
}

func (o IPSetPtrOutput) ToIPSetPtrOutputWithContext(ctx context.Context) IPSetPtrOutput {
	return o
}

type IPSetArrayOutput struct{ *pulumi.OutputState }

func (IPSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPSet)(nil))
}

func (o IPSetArrayOutput) ToIPSetArrayOutput() IPSetArrayOutput {
	return o
}

func (o IPSetArrayOutput) ToIPSetArrayOutputWithContext(ctx context.Context) IPSetArrayOutput {
	return o
}

func (o IPSetArrayOutput) Index(i pulumi.IntInput) IPSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPSet {
		return vs[0].([]IPSet)[vs[1].(int)]
	}).(IPSetOutput)
}

type IPSetMapOutput struct{ *pulumi.OutputState }

func (IPSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IPSet)(nil))
}

func (o IPSetMapOutput) ToIPSetMapOutput() IPSetMapOutput {
	return o
}

func (o IPSetMapOutput) ToIPSetMapOutputWithContext(ctx context.Context) IPSetMapOutput {
	return o
}

func (o IPSetMapOutput) MapIndex(k pulumi.StringInput) IPSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IPSet {
		return vs[0].(map[string]IPSet)[vs[1].(string)]
	}).(IPSetOutput)
}

func init() {
	pulumi.RegisterOutputType(IPSetOutput{})
	pulumi.RegisterOutputType(IPSetPtrOutput{})
	pulumi.RegisterOutputType(IPSetArrayOutput{})
	pulumi.RegisterOutputType(IPSetMapOutput{})
}
