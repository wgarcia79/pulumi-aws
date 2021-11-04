// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks User Profile resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewUserProfile(ctx, "myProfile", &opsworks.UserProfileArgs{
// 			UserArn:     pulumi.Any(aws_iam_user.User.Arn),
// 			SshUsername: pulumi.String("my_user"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type UserProfile struct {
	pulumi.CustomResourceState

	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrOutput `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey pulumi.StringPtrOutput `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringOutput `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn pulumi.StringOutput `pulumi:"userArn"`
}

// NewUserProfile registers a new resource with the given unique name, arguments, and options.
func NewUserProfile(ctx *pulumi.Context,
	name string, args *UserProfileArgs, opts ...pulumi.ResourceOption) (*UserProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SshUsername == nil {
		return nil, errors.New("invalid value for required argument 'SshUsername'")
	}
	if args.UserArn == nil {
		return nil, errors.New("invalid value for required argument 'UserArn'")
	}
	var resource UserProfile
	err := ctx.RegisterResource("aws:opsworks/userProfile:UserProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserProfile gets an existing UserProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserProfileState, opts ...pulumi.ResourceOption) (*UserProfile, error) {
	var resource UserProfile
	err := ctx.ReadResource("aws:opsworks/userProfile:UserProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserProfile resources.
type userProfileState struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername *string `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn *string `pulumi:"userArn"`
}

type UserProfileState struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrInput
	// The users public key
	SshPublicKey pulumi.StringPtrInput
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringPtrInput
	// The user's IAM ARN
	UserArn pulumi.StringPtrInput
}

func (UserProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*userProfileState)(nil)).Elem()
}

type userProfileArgs struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername string `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn string `pulumi:"userArn"`
}

// The set of arguments for constructing a UserProfile resource.
type UserProfileArgs struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrInput
	// The users public key
	SshPublicKey pulumi.StringPtrInput
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringInput
	// The user's IAM ARN
	UserArn pulumi.StringInput
}

func (UserProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userProfileArgs)(nil)).Elem()
}

type UserProfileInput interface {
	pulumi.Input

	ToUserProfileOutput() UserProfileOutput
	ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput
}

func (*UserProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfile)(nil))
}

func (i *UserProfile) ToUserProfileOutput() UserProfileOutput {
	return i.ToUserProfileOutputWithContext(context.Background())
}

func (i *UserProfile) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfileOutput)
}

func (i *UserProfile) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return i.ToUserProfilePtrOutputWithContext(context.Background())
}

func (i *UserProfile) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfilePtrOutput)
}

type UserProfilePtrInput interface {
	pulumi.Input

	ToUserProfilePtrOutput() UserProfilePtrOutput
	ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput
}

type userProfilePtrType UserProfileArgs

func (*userProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProfile)(nil))
}

func (i *userProfilePtrType) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return i.ToUserProfilePtrOutputWithContext(context.Background())
}

func (i *userProfilePtrType) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfilePtrOutput)
}

// UserProfileArrayInput is an input type that accepts UserProfileArray and UserProfileArrayOutput values.
// You can construct a concrete instance of `UserProfileArrayInput` via:
//
//          UserProfileArray{ UserProfileArgs{...} }
type UserProfileArrayInput interface {
	pulumi.Input

	ToUserProfileArrayOutput() UserProfileArrayOutput
	ToUserProfileArrayOutputWithContext(context.Context) UserProfileArrayOutput
}

type UserProfileArray []UserProfileInput

func (UserProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserProfile)(nil)).Elem()
}

func (i UserProfileArray) ToUserProfileArrayOutput() UserProfileArrayOutput {
	return i.ToUserProfileArrayOutputWithContext(context.Background())
}

func (i UserProfileArray) ToUserProfileArrayOutputWithContext(ctx context.Context) UserProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfileArrayOutput)
}

// UserProfileMapInput is an input type that accepts UserProfileMap and UserProfileMapOutput values.
// You can construct a concrete instance of `UserProfileMapInput` via:
//
//          UserProfileMap{ "key": UserProfileArgs{...} }
type UserProfileMapInput interface {
	pulumi.Input

	ToUserProfileMapOutput() UserProfileMapOutput
	ToUserProfileMapOutputWithContext(context.Context) UserProfileMapOutput
}

type UserProfileMap map[string]UserProfileInput

func (UserProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserProfile)(nil)).Elem()
}

func (i UserProfileMap) ToUserProfileMapOutput() UserProfileMapOutput {
	return i.ToUserProfileMapOutputWithContext(context.Background())
}

func (i UserProfileMap) ToUserProfileMapOutputWithContext(ctx context.Context) UserProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfileMapOutput)
}

type UserProfileOutput struct{ *pulumi.OutputState }

func (UserProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfile)(nil))
}

func (o UserProfileOutput) ToUserProfileOutput() UserProfileOutput {
	return o
}

func (o UserProfileOutput) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return o
}

func (o UserProfileOutput) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return o.ToUserProfilePtrOutputWithContext(context.Background())
}

func (o UserProfileOutput) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserProfile) *UserProfile {
		return &v
	}).(UserProfilePtrOutput)
}

type UserProfilePtrOutput struct{ *pulumi.OutputState }

func (UserProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProfile)(nil))
}

func (o UserProfilePtrOutput) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return o
}

func (o UserProfilePtrOutput) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return o
}

func (o UserProfilePtrOutput) Elem() UserProfileOutput {
	return o.ApplyT(func(v *UserProfile) UserProfile {
		if v != nil {
			return *v
		}
		var ret UserProfile
		return ret
	}).(UserProfileOutput)
}

type UserProfileArrayOutput struct{ *pulumi.OutputState }

func (UserProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserProfile)(nil))
}

func (o UserProfileArrayOutput) ToUserProfileArrayOutput() UserProfileArrayOutput {
	return o
}

func (o UserProfileArrayOutput) ToUserProfileArrayOutputWithContext(ctx context.Context) UserProfileArrayOutput {
	return o
}

func (o UserProfileArrayOutput) Index(i pulumi.IntInput) UserProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserProfile {
		return vs[0].([]UserProfile)[vs[1].(int)]
	}).(UserProfileOutput)
}

type UserProfileMapOutput struct{ *pulumi.OutputState }

func (UserProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserProfile)(nil))
}

func (o UserProfileMapOutput) ToUserProfileMapOutput() UserProfileMapOutput {
	return o
}

func (o UserProfileMapOutput) ToUserProfileMapOutputWithContext(ctx context.Context) UserProfileMapOutput {
	return o
}

func (o UserProfileMapOutput) MapIndex(k pulumi.StringInput) UserProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserProfile {
		return vs[0].(map[string]UserProfile)[vs[1].(string)]
	}).(UserProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserProfileInput)(nil)).Elem(), &UserProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserProfilePtrInput)(nil)).Elem(), &UserProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserProfileArrayInput)(nil)).Elem(), UserProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserProfileMapInput)(nil)).Elem(), UserProfileMap{})
	pulumi.RegisterOutputType(UserProfileOutput{})
	pulumi.RegisterOutputType(UserProfilePtrOutput{})
	pulumi.RegisterOutputType(UserProfileArrayOutput{})
	pulumi.RegisterOutputType(UserProfileMapOutput{})
}
