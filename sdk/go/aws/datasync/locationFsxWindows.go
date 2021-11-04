// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS DataSync FSx Windows Location.
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
// 		_, err := datasync.NewLocationFsxWindows(ctx, "example", &datasync.LocationFsxWindowsArgs{
// 			FsxFilesystemArn: pulumi.Any(aws_fsx_windows_file_system.Example.Arn),
// 			User:             pulumi.String("SomeUser"),
// 			Password:         pulumi.String("SuperSecretPassw0rd"),
// 			SecurityGroupArns: pulumi.StringArray{
// 				pulumi.Any(aws_security_group.Example.Arn),
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
// `aws_datasync_location_fsx_windows_file_system` can be imported by using the `DataSync-ARN#FSx-Windows-ARN`, e.g.
//
// ```sh
//  $ pulumi import aws:datasync/locationFsxWindows:LocationFsxWindows example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a
// ```
type LocationFsxWindows struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time that the FSx for Windows location was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringOutput `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayOutput `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The URL of the FSx for Windows location that was described.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewLocationFsxWindows registers a new resource with the given unique name, arguments, and options.
func NewLocationFsxWindows(ctx *pulumi.Context,
	name string, args *LocationFsxWindowsArgs, opts ...pulumi.ResourceOption) (*LocationFsxWindows, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FsxFilesystemArn == nil {
		return nil, errors.New("invalid value for required argument 'FsxFilesystemArn'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.SecurityGroupArns == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupArns'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource LocationFsxWindows
	err := ctx.RegisterResource("aws:datasync/locationFsxWindows:LocationFsxWindows", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationFsxWindows gets an existing LocationFsxWindows resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationFsxWindows(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationFsxWindowsState, opts ...pulumi.ResourceOption) (*LocationFsxWindows, error) {
	var resource LocationFsxWindows
	err := ctx.ReadResource("aws:datasync/locationFsxWindows:LocationFsxWindows", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationFsxWindows resources.
type locationFsxWindowsState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// The time that the FSx for Windows location was created.
	CreationTime *string `pulumi:"creationTime"`
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn *string `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password *string `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The URL of the FSx for Windows location that was described.
	Uri *string `pulumi:"uri"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User *string `pulumi:"user"`
}

type LocationFsxWindowsState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// The time that the FSx for Windows location was created.
	CreationTime pulumi.StringPtrInput
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringPtrInput
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringPtrInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The URL of the FSx for Windows location that was described.
	Uri pulumi.StringPtrInput
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringPtrInput
}

func (LocationFsxWindowsState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxWindowsState)(nil)).Elem()
}

type locationFsxWindowsArgs struct {
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `pulumi:"domain"`
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn string `pulumi:"fsxFilesystemArn"`
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password string `pulumi:"password"`
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Subdirectory to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a LocationFsxWindows resource.
type LocationFsxWindowsArgs struct {
	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	FsxFilesystemArn pulumi.StringInput
	// The password of the user who has the permissions to access files and folders in the FSx for Windows file system.
	Password pulumi.StringInput
	// The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Windows file system.
	SecurityGroupArns pulumi.StringArrayInput
	// Subdirectory to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The user who has the permissions to access files and folders in the FSx for Windows file system.
	User pulumi.StringInput
}

func (LocationFsxWindowsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationFsxWindowsArgs)(nil)).Elem()
}

type LocationFsxWindowsInput interface {
	pulumi.Input

	ToLocationFsxWindowsOutput() LocationFsxWindowsOutput
	ToLocationFsxWindowsOutputWithContext(ctx context.Context) LocationFsxWindowsOutput
}

func (*LocationFsxWindows) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationFsxWindows)(nil))
}

func (i *LocationFsxWindows) ToLocationFsxWindowsOutput() LocationFsxWindowsOutput {
	return i.ToLocationFsxWindowsOutputWithContext(context.Background())
}

func (i *LocationFsxWindows) ToLocationFsxWindowsOutputWithContext(ctx context.Context) LocationFsxWindowsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxWindowsOutput)
}

func (i *LocationFsxWindows) ToLocationFsxWindowsPtrOutput() LocationFsxWindowsPtrOutput {
	return i.ToLocationFsxWindowsPtrOutputWithContext(context.Background())
}

func (i *LocationFsxWindows) ToLocationFsxWindowsPtrOutputWithContext(ctx context.Context) LocationFsxWindowsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxWindowsPtrOutput)
}

type LocationFsxWindowsPtrInput interface {
	pulumi.Input

	ToLocationFsxWindowsPtrOutput() LocationFsxWindowsPtrOutput
	ToLocationFsxWindowsPtrOutputWithContext(ctx context.Context) LocationFsxWindowsPtrOutput
}

type locationFsxWindowsPtrType LocationFsxWindowsArgs

func (*locationFsxWindowsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFsxWindows)(nil))
}

func (i *locationFsxWindowsPtrType) ToLocationFsxWindowsPtrOutput() LocationFsxWindowsPtrOutput {
	return i.ToLocationFsxWindowsPtrOutputWithContext(context.Background())
}

func (i *locationFsxWindowsPtrType) ToLocationFsxWindowsPtrOutputWithContext(ctx context.Context) LocationFsxWindowsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxWindowsPtrOutput)
}

// LocationFsxWindowsArrayInput is an input type that accepts LocationFsxWindowsArray and LocationFsxWindowsArrayOutput values.
// You can construct a concrete instance of `LocationFsxWindowsArrayInput` via:
//
//          LocationFsxWindowsArray{ LocationFsxWindowsArgs{...} }
type LocationFsxWindowsArrayInput interface {
	pulumi.Input

	ToLocationFsxWindowsArrayOutput() LocationFsxWindowsArrayOutput
	ToLocationFsxWindowsArrayOutputWithContext(context.Context) LocationFsxWindowsArrayOutput
}

type LocationFsxWindowsArray []LocationFsxWindowsInput

func (LocationFsxWindowsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocationFsxWindows)(nil)).Elem()
}

func (i LocationFsxWindowsArray) ToLocationFsxWindowsArrayOutput() LocationFsxWindowsArrayOutput {
	return i.ToLocationFsxWindowsArrayOutputWithContext(context.Background())
}

func (i LocationFsxWindowsArray) ToLocationFsxWindowsArrayOutputWithContext(ctx context.Context) LocationFsxWindowsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxWindowsArrayOutput)
}

// LocationFsxWindowsMapInput is an input type that accepts LocationFsxWindowsMap and LocationFsxWindowsMapOutput values.
// You can construct a concrete instance of `LocationFsxWindowsMapInput` via:
//
//          LocationFsxWindowsMap{ "key": LocationFsxWindowsArgs{...} }
type LocationFsxWindowsMapInput interface {
	pulumi.Input

	ToLocationFsxWindowsMapOutput() LocationFsxWindowsMapOutput
	ToLocationFsxWindowsMapOutputWithContext(context.Context) LocationFsxWindowsMapOutput
}

type LocationFsxWindowsMap map[string]LocationFsxWindowsInput

func (LocationFsxWindowsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocationFsxWindows)(nil)).Elem()
}

func (i LocationFsxWindowsMap) ToLocationFsxWindowsMapOutput() LocationFsxWindowsMapOutput {
	return i.ToLocationFsxWindowsMapOutputWithContext(context.Background())
}

func (i LocationFsxWindowsMap) ToLocationFsxWindowsMapOutputWithContext(ctx context.Context) LocationFsxWindowsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationFsxWindowsMapOutput)
}

type LocationFsxWindowsOutput struct{ *pulumi.OutputState }

func (LocationFsxWindowsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationFsxWindows)(nil))
}

func (o LocationFsxWindowsOutput) ToLocationFsxWindowsOutput() LocationFsxWindowsOutput {
	return o
}

func (o LocationFsxWindowsOutput) ToLocationFsxWindowsOutputWithContext(ctx context.Context) LocationFsxWindowsOutput {
	return o
}

func (o LocationFsxWindowsOutput) ToLocationFsxWindowsPtrOutput() LocationFsxWindowsPtrOutput {
	return o.ToLocationFsxWindowsPtrOutputWithContext(context.Background())
}

func (o LocationFsxWindowsOutput) ToLocationFsxWindowsPtrOutputWithContext(ctx context.Context) LocationFsxWindowsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationFsxWindows) *LocationFsxWindows {
		return &v
	}).(LocationFsxWindowsPtrOutput)
}

type LocationFsxWindowsPtrOutput struct{ *pulumi.OutputState }

func (LocationFsxWindowsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationFsxWindows)(nil))
}

func (o LocationFsxWindowsPtrOutput) ToLocationFsxWindowsPtrOutput() LocationFsxWindowsPtrOutput {
	return o
}

func (o LocationFsxWindowsPtrOutput) ToLocationFsxWindowsPtrOutputWithContext(ctx context.Context) LocationFsxWindowsPtrOutput {
	return o
}

func (o LocationFsxWindowsPtrOutput) Elem() LocationFsxWindowsOutput {
	return o.ApplyT(func(v *LocationFsxWindows) LocationFsxWindows {
		if v != nil {
			return *v
		}
		var ret LocationFsxWindows
		return ret
	}).(LocationFsxWindowsOutput)
}

type LocationFsxWindowsArrayOutput struct{ *pulumi.OutputState }

func (LocationFsxWindowsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationFsxWindows)(nil))
}

func (o LocationFsxWindowsArrayOutput) ToLocationFsxWindowsArrayOutput() LocationFsxWindowsArrayOutput {
	return o
}

func (o LocationFsxWindowsArrayOutput) ToLocationFsxWindowsArrayOutputWithContext(ctx context.Context) LocationFsxWindowsArrayOutput {
	return o
}

func (o LocationFsxWindowsArrayOutput) Index(i pulumi.IntInput) LocationFsxWindowsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocationFsxWindows {
		return vs[0].([]LocationFsxWindows)[vs[1].(int)]
	}).(LocationFsxWindowsOutput)
}

type LocationFsxWindowsMapOutput struct{ *pulumi.OutputState }

func (LocationFsxWindowsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocationFsxWindows)(nil))
}

func (o LocationFsxWindowsMapOutput) ToLocationFsxWindowsMapOutput() LocationFsxWindowsMapOutput {
	return o
}

func (o LocationFsxWindowsMapOutput) ToLocationFsxWindowsMapOutputWithContext(ctx context.Context) LocationFsxWindowsMapOutput {
	return o
}

func (o LocationFsxWindowsMapOutput) MapIndex(k pulumi.StringInput) LocationFsxWindowsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocationFsxWindows {
		return vs[0].(map[string]LocationFsxWindows)[vs[1].(string)]
	}).(LocationFsxWindowsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxWindowsInput)(nil)).Elem(), &LocationFsxWindows{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxWindowsPtrInput)(nil)).Elem(), &LocationFsxWindows{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxWindowsArrayInput)(nil)).Elem(), LocationFsxWindowsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocationFsxWindowsMapInput)(nil)).Elem(), LocationFsxWindowsMap{})
	pulumi.RegisterOutputType(LocationFsxWindowsOutput{})
	pulumi.RegisterOutputType(LocationFsxWindowsPtrOutput{})
	pulumi.RegisterOutputType(LocationFsxWindowsArrayOutput{})
	pulumi.RegisterOutputType(LocationFsxWindowsMapOutput{})
}
