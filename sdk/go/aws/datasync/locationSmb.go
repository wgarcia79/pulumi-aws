// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a SMB Location within AWS DataSync.
//
// > **NOTE:** The DataSync Agents must be available before creating this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/datasync"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datasync.NewLocationSmb(ctx, "example", &datasync.LocationSmbArgs{
// 			ServerHostname: pulumi.String("smb.example.com"),
// 			Subdirectory:   pulumi.String("/exported/path"),
// 			User:           pulumi.String("Guest"),
// 			Password:       pulumi.String("ANotGreatPassword"),
// 			AgentArns: pulumi.StringArray{
// 				pulumi.Any(aws_datasync_agent.Example.Arn),
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
// `aws_datasync_location_smb` can be imported by using the Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:datasync/locationSmb:LocationSmb example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
// ```
type LocationSmb struct {
	pulumi.CustomResourceState

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayOutput `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrOutput `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringOutput `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringOutput `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	Uri     pulumi.StringOutput    `pulumi:"uri"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewLocationSmb registers a new resource with the given unique name, arguments, and options.
func NewLocationSmb(ctx *pulumi.Context,
	name string, args *LocationSmbArgs, opts ...pulumi.ResourceOption) (*LocationSmb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentArns == nil {
		return nil, errors.New("invalid value for required argument 'AgentArns'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.ServerHostname == nil {
		return nil, errors.New("invalid value for required argument 'ServerHostname'")
	}
	if args.Subdirectory == nil {
		return nil, errors.New("invalid value for required argument 'Subdirectory'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource LocationSmb
	err := ctx.RegisterResource("aws:datasync/locationSmb:LocationSmb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationSmb gets an existing LocationSmb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationSmb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationSmbState, opts ...pulumi.ResourceOption) (*LocationSmb, error) {
	var resource LocationSmb
	err := ctx.ReadResource("aws:datasync/locationSmb:LocationSmb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationSmb resources.
type locationSmbState struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The name of the Windows domain the SMB server belongs to.
	Domain *string `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions *LocationSmbMountOptions `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password *string `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname *string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	Uri     *string           `pulumi:"uri"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User *string `pulumi:"user"`
}

type LocationSmbState struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringPtrInput
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrInput
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringPtrInput
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringPtrInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	Uri     pulumi.StringPtrInput
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringPtrInput
}

func (LocationSmbState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationSmbState)(nil)).Elem()
}

type locationSmbArgs struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `pulumi:"agentArns"`
	// The name of the Windows domain the SMB server belongs to.
	Domain *string `pulumi:"domain"`
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions *LocationSmbMountOptions `pulumi:"mountOptions"`
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password string `pulumi:"password"`
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname string `pulumi:"serverHostname"`
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a LocationSmb resource.
type LocationSmbArgs struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns pulumi.StringArrayInput
	// The name of the Windows domain the SMB server belongs to.
	Domain pulumi.StringPtrInput
	// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
	MountOptions LocationSmbMountOptionsPtrInput
	// The password of the user who can mount the share and has file permissions in the SMB.
	Password pulumi.StringInput
	// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
	ServerHostname pulumi.StringInput
	// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
	Subdirectory pulumi.StringInput
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The user who can mount the share and has file and folder permissions in the SMB share.
	User pulumi.StringInput
}

func (LocationSmbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationSmbArgs)(nil)).Elem()
}

type LocationSmbInput interface {
	pulumi.Input

	ToLocationSmbOutput() LocationSmbOutput
	ToLocationSmbOutputWithContext(ctx context.Context) LocationSmbOutput
}

func (*LocationSmb) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationSmb)(nil))
}

func (i *LocationSmb) ToLocationSmbOutput() LocationSmbOutput {
	return i.ToLocationSmbOutputWithContext(context.Background())
}

func (i *LocationSmb) ToLocationSmbOutputWithContext(ctx context.Context) LocationSmbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbOutput)
}

func (i *LocationSmb) ToLocationSmbPtrOutput() LocationSmbPtrOutput {
	return i.ToLocationSmbPtrOutputWithContext(context.Background())
}

func (i *LocationSmb) ToLocationSmbPtrOutputWithContext(ctx context.Context) LocationSmbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbPtrOutput)
}

type LocationSmbPtrInput interface {
	pulumi.Input

	ToLocationSmbPtrOutput() LocationSmbPtrOutput
	ToLocationSmbPtrOutputWithContext(ctx context.Context) LocationSmbPtrOutput
}

type locationSmbPtrType LocationSmbArgs

func (*locationSmbPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationSmb)(nil))
}

func (i *locationSmbPtrType) ToLocationSmbPtrOutput() LocationSmbPtrOutput {
	return i.ToLocationSmbPtrOutputWithContext(context.Background())
}

func (i *locationSmbPtrType) ToLocationSmbPtrOutputWithContext(ctx context.Context) LocationSmbPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbPtrOutput)
}

// LocationSmbArrayInput is an input type that accepts LocationSmbArray and LocationSmbArrayOutput values.
// You can construct a concrete instance of `LocationSmbArrayInput` via:
//
//          LocationSmbArray{ LocationSmbArgs{...} }
type LocationSmbArrayInput interface {
	pulumi.Input

	ToLocationSmbArrayOutput() LocationSmbArrayOutput
	ToLocationSmbArrayOutputWithContext(context.Context) LocationSmbArrayOutput
}

type LocationSmbArray []LocationSmbInput

func (LocationSmbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationSmb)(nil)).Elem()
}

func (i LocationSmbArray) ToLocationSmbArrayOutput() LocationSmbArrayOutput {
	return i.ToLocationSmbArrayOutputWithContext(context.Background())
}

func (i LocationSmbArray) ToLocationSmbArrayOutputWithContext(ctx context.Context) LocationSmbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbArrayOutput)
}

// LocationSmbMapInput is an input type that accepts LocationSmbMap and LocationSmbMapOutput values.
// You can construct a concrete instance of `LocationSmbMapInput` via:
//
//          LocationSmbMap{ "key": LocationSmbArgs{...} }
type LocationSmbMapInput interface {
	pulumi.Input

	ToLocationSmbMapOutput() LocationSmbMapOutput
	ToLocationSmbMapOutputWithContext(context.Context) LocationSmbMapOutput
}

type LocationSmbMap map[string]LocationSmbInput

func (LocationSmbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationSmb)(nil)).Elem()
}

func (i LocationSmbMap) ToLocationSmbMapOutput() LocationSmbMapOutput {
	return i.ToLocationSmbMapOutputWithContext(context.Background())
}

func (i LocationSmbMap) ToLocationSmbMapOutputWithContext(ctx context.Context) LocationSmbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbMapOutput)
}

type LocationSmbOutput struct{ *pulumi.OutputState }

func (LocationSmbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationSmb)(nil))
}

func (o LocationSmbOutput) ToLocationSmbOutput() LocationSmbOutput {
	return o
}

func (o LocationSmbOutput) ToLocationSmbOutputWithContext(ctx context.Context) LocationSmbOutput {
	return o
}

func (o LocationSmbOutput) ToLocationSmbPtrOutput() LocationSmbPtrOutput {
	return o.ToLocationSmbPtrOutputWithContext(context.Background())
}

func (o LocationSmbOutput) ToLocationSmbPtrOutputWithContext(ctx context.Context) LocationSmbPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationSmb) *LocationSmb {
		return &v
	}).(LocationSmbPtrOutput)
}

type LocationSmbPtrOutput struct{ *pulumi.OutputState }

func (LocationSmbPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationSmb)(nil))
}

func (o LocationSmbPtrOutput) ToLocationSmbPtrOutput() LocationSmbPtrOutput {
	return o
}

func (o LocationSmbPtrOutput) ToLocationSmbPtrOutputWithContext(ctx context.Context) LocationSmbPtrOutput {
	return o
}

func (o LocationSmbPtrOutput) Elem() LocationSmbOutput {
	return o.ApplyT(func(v *LocationSmb) LocationSmb {
		if v != nil {
			return *v
		}
		var ret LocationSmb
		return ret
	}).(LocationSmbOutput)
}

type LocationSmbArrayOutput struct{ *pulumi.OutputState }

func (LocationSmbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationSmb)(nil))
}

func (o LocationSmbArrayOutput) ToLocationSmbArrayOutput() LocationSmbArrayOutput {
	return o
}

func (o LocationSmbArrayOutput) ToLocationSmbArrayOutputWithContext(ctx context.Context) LocationSmbArrayOutput {
	return o
}

func (o LocationSmbArrayOutput) Index(i pulumi.IntInput) LocationSmbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocationSmb {
		return vs[0].([]LocationSmb)[vs[1].(int)]
	}).(LocationSmbOutput)
}

type LocationSmbMapOutput struct{ *pulumi.OutputState }

func (LocationSmbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocationSmb)(nil))
}

func (o LocationSmbMapOutput) ToLocationSmbMapOutput() LocationSmbMapOutput {
	return o
}

func (o LocationSmbMapOutput) ToLocationSmbMapOutputWithContext(ctx context.Context) LocationSmbMapOutput {
	return o
}

func (o LocationSmbMapOutput) MapIndex(k pulumi.StringInput) LocationSmbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocationSmb {
		return vs[0].(map[string]LocationSmb)[vs[1].(string)]
	}).(LocationSmbOutput)
}

func init() {
	pulumi.RegisterOutputType(LocationSmbOutput{})
	pulumi.RegisterOutputType(LocationSmbPtrOutput{})
	pulumi.RegisterOutputType(LocationSmbArrayOutput{})
	pulumi.RegisterOutputType(LocationSmbMapOutput{})
}
