// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ElastiCache user resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticache"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elasticache.NewUser(ctx, "test", &elasticache.UserArgs{
// 			AccessString: pulumi.String("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember"),
// 			Engine:       pulumi.String("REDIS"),
// 			Passwords: pulumi.StringArray{
// 				pulumi.String("password123456789"),
// 			},
// 			UserId:   pulumi.String("testUserId"),
// 			UserName: pulumi.String("testUserName"),
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
// ElastiCache users can be imported using the `user_id`, e.g.
//
// ```sh
//  $ pulumi import aws:elasticache/user:User my_user userId1
// ```
type User struct {
	pulumi.CustomResourceState

	// Access permissions string used for this user.
	AccessString pulumi.StringOutput `pulumi:"accessString"`
	// The ARN of the created ElastiCache User.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The current supported value is `REDIS`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Indicates a password is not required for this user.
	NoPasswordRequired pulumi.BoolPtrOutput `pulumi:"noPasswordRequired"`
	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayOutput `pulumi:"passwords"`
	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the user.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The username of the user.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessString == nil {
		return nil, errors.New("invalid value for required argument 'AccessString'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource User
	err := ctx.RegisterResource("aws:elasticache/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:elasticache/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Access permissions string used for this user.
	AccessString *string `pulumi:"accessString"`
	// The ARN of the created ElastiCache User.
	Arn *string `pulumi:"arn"`
	// The current supported value is `REDIS`.
	Engine *string `pulumi:"engine"`
	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `pulumi:"noPasswordRequired"`
	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords []string `pulumi:"passwords"`
	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the user.
	UserId *string `pulumi:"userId"`
	// The username of the user.
	UserName *string `pulumi:"userName"`
}

type UserState struct {
	// Access permissions string used for this user.
	AccessString pulumi.StringPtrInput
	// The ARN of the created ElastiCache User.
	Arn pulumi.StringPtrInput
	// The current supported value is `REDIS`.
	Engine pulumi.StringPtrInput
	// Indicates a password is not required for this user.
	NoPasswordRequired pulumi.BoolPtrInput
	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayInput
	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The ID of the user.
	UserId pulumi.StringPtrInput
	// The username of the user.
	UserName pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Access permissions string used for this user.
	AccessString string `pulumi:"accessString"`
	// The ARN of the created ElastiCache User.
	Arn *string `pulumi:"arn"`
	// The current supported value is `REDIS`.
	Engine string `pulumi:"engine"`
	// Indicates a password is not required for this user.
	NoPasswordRequired *bool `pulumi:"noPasswordRequired"`
	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords []string `pulumi:"passwords"`
	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the user.
	UserId string `pulumi:"userId"`
	// The username of the user.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Access permissions string used for this user.
	AccessString pulumi.StringInput
	// The ARN of the created ElastiCache User.
	Arn pulumi.StringPtrInput
	// The current supported value is `REDIS`.
	Engine pulumi.StringInput
	// Indicates a password is not required for this user.
	NoPasswordRequired pulumi.BoolPtrInput
	// Passwords used for this user. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayInput
	// A list of tags to be added to this resource. A tag is a key-value pair.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
	// The ID of the user.
	UserId pulumi.StringInput
	// The username of the user.
	UserName pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

func (i *User) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *User) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

type UserPtrInput interface {
	pulumi.Input

	ToUserPtrOutput() UserPtrOutput
	ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput
}

type userPtrType UserArgs

func (*userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (i *userPtrType) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *userPtrType) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct {
	*pulumi.OutputState
}

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) ToUserPtrOutput() UserPtrOutput {
	return o.ToUserPtrOutputWithContext(context.Background())
}

func (o UserOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o.ApplyT(func(v User) *User {
		return &v
	}).(UserPtrOutput)
}

type UserPtrOutput struct {
	*pulumi.OutputState
}

func (UserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (o UserPtrOutput) ToUserPtrOutput() UserPtrOutput {
	return o
}

func (o UserPtrOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]User)(nil))
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) User {
		return vs[0].([]User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]User)(nil))
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) User {
		return vs[0].(map[string]User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserPtrOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
