// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EKS add-on.
//
// > **Note:** Amazon EKS add-on can only be used with Amazon EKS Clusters
// running version 1.18 with platform version eks.3 or later
// because add-ons rely on the Server-side Apply Kubernetes feature,
// which is only available in Kubernetes 1.18 and later.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/eks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eks.NewAddon(ctx, "example", &eks.AddonArgs{
// 			ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
// 			AddonName:   pulumi.String("vpc-cni"),
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
// EKS add-on can be imported using the `cluster_name` and `addon_name` separated by a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:eks/addon:Addon my_eks_addon my_cluster_name:my_addon_name
// ```
type Addon struct {
	pulumi.CustomResourceState

	// Name of the EKS add-on. The name must match one of
	// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
	AddonName pulumi.StringOutput `pulumi:"addonName"`
	// The version of the EKS add-on. The version must
	// match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	AddonVersion pulumi.StringOutput `pulumi:"addonVersion"`
	// Amazon Resource Name (ARN) of the EKS add-on.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
	ModifiedAt pulumi.StringOutput `pulumi:"modifiedAt"`
	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are `NONE` and `OVERWRITE`.
	ResolveConflicts pulumi.StringPtrOutput `pulumi:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn pulumi.StringPtrOutput `pulumi:"serviceAccountRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// (Optional) Key-value map of resource tags, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewAddon registers a new resource with the given unique name, arguments, and options.
func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddonName == nil {
		return nil, errors.New("invalid value for required argument 'AddonName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	var resource Addon
	err := ctx.RegisterResource("aws:eks/addon:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddon gets an existing Addon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("aws:eks/addon:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Addon resources.
type addonState struct {
	// Name of the EKS add-on. The name must match one of
	// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
	AddonName *string `pulumi:"addonName"`
	// The version of the EKS add-on. The version must
	// match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	AddonVersion *string `pulumi:"addonVersion"`
	// Amazon Resource Name (ARN) of the EKS add-on.
	Arn *string `pulumi:"arn"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName *string `pulumi:"clusterName"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
	ModifiedAt *string `pulumi:"modifiedAt"`
	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are `NONE` and `OVERWRITE`.
	ResolveConflicts *string `pulumi:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn *string `pulumi:"serviceAccountRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// (Optional) Key-value map of resource tags, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type AddonState struct {
	// Name of the EKS add-on. The name must match one of
	// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
	AddonName pulumi.StringPtrInput
	// The version of the EKS add-on. The version must
	// match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	AddonVersion pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the EKS add-on.
	Arn pulumi.StringPtrInput
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringPtrInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
	CreatedAt pulumi.StringPtrInput
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
	ModifiedAt pulumi.StringPtrInput
	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are `NONE` and `OVERWRITE`.
	ResolveConflicts pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// (Optional) Key-value map of resource tags, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	// Name of the EKS add-on. The name must match one of
	// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
	AddonName string `pulumi:"addonName"`
	// The version of the EKS add-on. The version must
	// match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	AddonVersion *string `pulumi:"addonVersion"`
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName string `pulumi:"clusterName"`
	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are `NONE` and `OVERWRITE`.
	ResolveConflicts *string `pulumi:"resolveConflicts"`
	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn *string `pulumi:"serviceAccountRoleArn"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Addon resource.
type AddonArgs struct {
	// Name of the EKS add-on. The name must match one of
	// the names returned by [list-addon](https://docs.aws.amazon.com/cli/latest/reference/eks/list-addons.html).
	AddonName pulumi.StringInput
	// The version of the EKS add-on. The version must
	// match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
	AddonVersion pulumi.StringPtrInput
	// Name of the EKS Cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
	ClusterName pulumi.StringInput
	// Define how to resolve parameter value conflicts
	// when migrating an existing add-on to an Amazon EKS add-on or when applying
	// version updates to the add-on. Valid values are `NONE` and `OVERWRITE`.
	ResolveConflicts pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an
	// existing IAM role to bind to the add-on's service account. The role must be
	// assigned the IAM permissions required by the add-on. If you don't specify
	// an existing IAM role, then the add-on uses the permissions assigned to the node
	// IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// in the Amazon EKS User Guide.
	ServiceAccountRoleArn pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((*Addon)(nil))
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

func (i *Addon) ToAddonPtrOutput() AddonPtrOutput {
	return i.ToAddonPtrOutputWithContext(context.Background())
}

func (i *Addon) ToAddonPtrOutputWithContext(ctx context.Context) AddonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonPtrOutput)
}

type AddonPtrInput interface {
	pulumi.Input

	ToAddonPtrOutput() AddonPtrOutput
	ToAddonPtrOutputWithContext(ctx context.Context) AddonPtrOutput
}

type addonPtrType AddonArgs

func (*addonPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil))
}

func (i *addonPtrType) ToAddonPtrOutput() AddonPtrOutput {
	return i.ToAddonPtrOutputWithContext(context.Background())
}

func (i *addonPtrType) ToAddonPtrOutputWithContext(ctx context.Context) AddonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonPtrOutput)
}

// AddonArrayInput is an input type that accepts AddonArray and AddonArrayOutput values.
// You can construct a concrete instance of `AddonArrayInput` via:
//
//          AddonArray{ AddonArgs{...} }
type AddonArrayInput interface {
	pulumi.Input

	ToAddonArrayOutput() AddonArrayOutput
	ToAddonArrayOutputWithContext(context.Context) AddonArrayOutput
}

type AddonArray []AddonInput

func (AddonArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Addon)(nil)).Elem()
}

func (i AddonArray) ToAddonArrayOutput() AddonArrayOutput {
	return i.ToAddonArrayOutputWithContext(context.Background())
}

func (i AddonArray) ToAddonArrayOutputWithContext(ctx context.Context) AddonArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonArrayOutput)
}

// AddonMapInput is an input type that accepts AddonMap and AddonMapOutput values.
// You can construct a concrete instance of `AddonMapInput` via:
//
//          AddonMap{ "key": AddonArgs{...} }
type AddonMapInput interface {
	pulumi.Input

	ToAddonMapOutput() AddonMapOutput
	ToAddonMapOutputWithContext(context.Context) AddonMapOutput
}

type AddonMap map[string]AddonInput

func (AddonMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Addon)(nil)).Elem()
}

func (i AddonMap) ToAddonMapOutput() AddonMapOutput {
	return i.ToAddonMapOutputWithContext(context.Background())
}

func (i AddonMap) ToAddonMapOutputWithContext(ctx context.Context) AddonMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonMapOutput)
}

type AddonOutput struct{ *pulumi.OutputState }

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Addon)(nil))
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

func (o AddonOutput) ToAddonPtrOutput() AddonPtrOutput {
	return o.ToAddonPtrOutputWithContext(context.Background())
}

func (o AddonOutput) ToAddonPtrOutputWithContext(ctx context.Context) AddonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Addon) *Addon {
		return &v
	}).(AddonPtrOutput)
}

type AddonPtrOutput struct{ *pulumi.OutputState }

func (AddonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil))
}

func (o AddonPtrOutput) ToAddonPtrOutput() AddonPtrOutput {
	return o
}

func (o AddonPtrOutput) ToAddonPtrOutputWithContext(ctx context.Context) AddonPtrOutput {
	return o
}

func (o AddonPtrOutput) Elem() AddonOutput {
	return o.ApplyT(func(v *Addon) Addon {
		if v != nil {
			return *v
		}
		var ret Addon
		return ret
	}).(AddonOutput)
}

type AddonArrayOutput struct{ *pulumi.OutputState }

func (AddonArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Addon)(nil))
}

func (o AddonArrayOutput) ToAddonArrayOutput() AddonArrayOutput {
	return o
}

func (o AddonArrayOutput) ToAddonArrayOutputWithContext(ctx context.Context) AddonArrayOutput {
	return o
}

func (o AddonArrayOutput) Index(i pulumi.IntInput) AddonOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Addon {
		return vs[0].([]Addon)[vs[1].(int)]
	}).(AddonOutput)
}

type AddonMapOutput struct{ *pulumi.OutputState }

func (AddonMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Addon)(nil))
}

func (o AddonMapOutput) ToAddonMapOutput() AddonMapOutput {
	return o
}

func (o AddonMapOutput) ToAddonMapOutputWithContext(ctx context.Context) AddonMapOutput {
	return o
}

func (o AddonMapOutput) MapIndex(k pulumi.StringInput) AddonOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Addon {
		return vs[0].(map[string]Addon)[vs[1].(string)]
	}).(AddonOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddonInput)(nil)).Elem(), &Addon{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddonPtrInput)(nil)).Elem(), &Addon{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddonArrayInput)(nil)).Elem(), AddonArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddonMapInput)(nil)).Elem(), AddonMap{})
	pulumi.RegisterOutputType(AddonOutput{})
	pulumi.RegisterOutputType(AddonPtrOutput{})
	pulumi.RegisterOutputType(AddonArrayOutput{})
	pulumi.RegisterOutputType(AddonMapOutput{})
}
