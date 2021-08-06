// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC DHCP Options resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewVpcDhcpOptions(ctx, "dnsResolver", &ec2.VpcDhcpOptionsArgs{
// 			DomainNameServers: pulumi.StringArray{
// 				pulumi.String("8.8.8.8"),
// 				pulumi.String("8.8.4.4"),
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
// Full usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewVpcDhcpOptions(ctx, "foo", &ec2.VpcDhcpOptionsArgs{
// 			DomainName: pulumi.String("service.consul"),
// 			DomainNameServers: pulumi.StringArray{
// 				pulumi.String("127.0.0.1"),
// 				pulumi.String("10.0.0.2"),
// 			},
// 			NetbiosNameServers: pulumi.StringArray{
// 				pulumi.String("127.0.0.1"),
// 			},
// 			NetbiosNodeType: pulumi.String("2"),
// 			NtpServers: pulumi.StringArray{
// 				pulumi.String("127.0.0.1"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("foo-name"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Remarks
//
// * Notice that all arguments are optional but you have to specify at least one argument.
// * `domainNameServers`, `netbiosNameServers`, `ntpServers` are limited by AWS to maximum four servers only.
// * To actually use the DHCP Options Set you need to associate it to a VPC using `ec2.VpcDhcpOptionsAssociation`.
// * If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
// * In most cases unless you're configuring your own DNS you'll want to set `domainNameServers` to `AmazonProvidedDNS`.
//
// ## Import
//
// VPC DHCP Options can be imported using the `dhcp options id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/vpcDhcpOptions:VpcDhcpOptions my_options dopt-d9070ebb
// ```
type VpcDhcpOptions struct {
	pulumi.CustomResourceState

	// The ARN of the DHCP Options Set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayOutput `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayOutput `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrOutput `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayOutput `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcDhcpOptions registers a new resource with the given unique name, arguments, and options.
func NewVpcDhcpOptions(ctx *pulumi.Context,
	name string, args *VpcDhcpOptionsArgs, opts ...pulumi.ResourceOption) (*VpcDhcpOptions, error) {
	if args == nil {
		args = &VpcDhcpOptionsArgs{}
	}

	var resource VpcDhcpOptions
	err := ctx.RegisterResource("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcDhcpOptions gets an existing VpcDhcpOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcDhcpOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcDhcpOptionsState, opts ...pulumi.ResourceOption) (*VpcDhcpOptions, error) {
	var resource VpcDhcpOptions
	err := ctx.ReadResource("aws:ec2/vpcDhcpOptions:VpcDhcpOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcDhcpOptions resources.
type vpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn *string `pulumi:"arn"`
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName *string `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers []string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers []string `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn pulumi.StringPtrInput
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrInput
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayInput
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VpcDhcpOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsState)(nil)).Elem()
}

type vpcDhcpOptionsArgs struct {
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName *string `pulumi:"domainName"`
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers []string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers []string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	// List of NTP servers to configure.
	NtpServers []string `pulumi:"ntpServers"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a VpcDhcpOptions resource.
type VpcDhcpOptionsArgs struct {
	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
	DomainName pulumi.StringPtrInput
	// List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
	DomainNameServers pulumi.StringArrayInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringArrayInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	// List of NTP servers to configure.
	NtpServers pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (VpcDhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsArgs)(nil)).Elem()
}

type VpcDhcpOptionsInput interface {
	pulumi.Input

	ToVpcDhcpOptionsOutput() VpcDhcpOptionsOutput
	ToVpcDhcpOptionsOutputWithContext(ctx context.Context) VpcDhcpOptionsOutput
}

func (*VpcDhcpOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcDhcpOptions)(nil))
}

func (i *VpcDhcpOptions) ToVpcDhcpOptionsOutput() VpcDhcpOptionsOutput {
	return i.ToVpcDhcpOptionsOutputWithContext(context.Background())
}

func (i *VpcDhcpOptions) ToVpcDhcpOptionsOutputWithContext(ctx context.Context) VpcDhcpOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcDhcpOptionsOutput)
}

func (i *VpcDhcpOptions) ToVpcDhcpOptionsPtrOutput() VpcDhcpOptionsPtrOutput {
	return i.ToVpcDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i *VpcDhcpOptions) ToVpcDhcpOptionsPtrOutputWithContext(ctx context.Context) VpcDhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcDhcpOptionsPtrOutput)
}

type VpcDhcpOptionsPtrInput interface {
	pulumi.Input

	ToVpcDhcpOptionsPtrOutput() VpcDhcpOptionsPtrOutput
	ToVpcDhcpOptionsPtrOutputWithContext(ctx context.Context) VpcDhcpOptionsPtrOutput
}

type vpcDhcpOptionsPtrType VpcDhcpOptionsArgs

func (*vpcDhcpOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcDhcpOptions)(nil))
}

func (i *vpcDhcpOptionsPtrType) ToVpcDhcpOptionsPtrOutput() VpcDhcpOptionsPtrOutput {
	return i.ToVpcDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i *vpcDhcpOptionsPtrType) ToVpcDhcpOptionsPtrOutputWithContext(ctx context.Context) VpcDhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcDhcpOptionsPtrOutput)
}

// VpcDhcpOptionsArrayInput is an input type that accepts VpcDhcpOptionsArray and VpcDhcpOptionsArrayOutput values.
// You can construct a concrete instance of `VpcDhcpOptionsArrayInput` via:
//
//          VpcDhcpOptionsArray{ VpcDhcpOptionsArgs{...} }
type VpcDhcpOptionsArrayInput interface {
	pulumi.Input

	ToVpcDhcpOptionsArrayOutput() VpcDhcpOptionsArrayOutput
	ToVpcDhcpOptionsArrayOutputWithContext(context.Context) VpcDhcpOptionsArrayOutput
}

type VpcDhcpOptionsArray []VpcDhcpOptionsInput

func (VpcDhcpOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcDhcpOptions)(nil)).Elem()
}

func (i VpcDhcpOptionsArray) ToVpcDhcpOptionsArrayOutput() VpcDhcpOptionsArrayOutput {
	return i.ToVpcDhcpOptionsArrayOutputWithContext(context.Background())
}

func (i VpcDhcpOptionsArray) ToVpcDhcpOptionsArrayOutputWithContext(ctx context.Context) VpcDhcpOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcDhcpOptionsArrayOutput)
}

// VpcDhcpOptionsMapInput is an input type that accepts VpcDhcpOptionsMap and VpcDhcpOptionsMapOutput values.
// You can construct a concrete instance of `VpcDhcpOptionsMapInput` via:
//
//          VpcDhcpOptionsMap{ "key": VpcDhcpOptionsArgs{...} }
type VpcDhcpOptionsMapInput interface {
	pulumi.Input

	ToVpcDhcpOptionsMapOutput() VpcDhcpOptionsMapOutput
	ToVpcDhcpOptionsMapOutputWithContext(context.Context) VpcDhcpOptionsMapOutput
}

type VpcDhcpOptionsMap map[string]VpcDhcpOptionsInput

func (VpcDhcpOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcDhcpOptions)(nil)).Elem()
}

func (i VpcDhcpOptionsMap) ToVpcDhcpOptionsMapOutput() VpcDhcpOptionsMapOutput {
	return i.ToVpcDhcpOptionsMapOutputWithContext(context.Background())
}

func (i VpcDhcpOptionsMap) ToVpcDhcpOptionsMapOutputWithContext(ctx context.Context) VpcDhcpOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcDhcpOptionsMapOutput)
}

type VpcDhcpOptionsOutput struct {
	*pulumi.OutputState
}

func (VpcDhcpOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcDhcpOptions)(nil))
}

func (o VpcDhcpOptionsOutput) ToVpcDhcpOptionsOutput() VpcDhcpOptionsOutput {
	return o
}

func (o VpcDhcpOptionsOutput) ToVpcDhcpOptionsOutputWithContext(ctx context.Context) VpcDhcpOptionsOutput {
	return o
}

func (o VpcDhcpOptionsOutput) ToVpcDhcpOptionsPtrOutput() VpcDhcpOptionsPtrOutput {
	return o.ToVpcDhcpOptionsPtrOutputWithContext(context.Background())
}

func (o VpcDhcpOptionsOutput) ToVpcDhcpOptionsPtrOutputWithContext(ctx context.Context) VpcDhcpOptionsPtrOutput {
	return o.ApplyT(func(v VpcDhcpOptions) *VpcDhcpOptions {
		return &v
	}).(VpcDhcpOptionsPtrOutput)
}

type VpcDhcpOptionsPtrOutput struct {
	*pulumi.OutputState
}

func (VpcDhcpOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcDhcpOptions)(nil))
}

func (o VpcDhcpOptionsPtrOutput) ToVpcDhcpOptionsPtrOutput() VpcDhcpOptionsPtrOutput {
	return o
}

func (o VpcDhcpOptionsPtrOutput) ToVpcDhcpOptionsPtrOutputWithContext(ctx context.Context) VpcDhcpOptionsPtrOutput {
	return o
}

type VpcDhcpOptionsArrayOutput struct{ *pulumi.OutputState }

func (VpcDhcpOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcDhcpOptions)(nil))
}

func (o VpcDhcpOptionsArrayOutput) ToVpcDhcpOptionsArrayOutput() VpcDhcpOptionsArrayOutput {
	return o
}

func (o VpcDhcpOptionsArrayOutput) ToVpcDhcpOptionsArrayOutputWithContext(ctx context.Context) VpcDhcpOptionsArrayOutput {
	return o
}

func (o VpcDhcpOptionsArrayOutput) Index(i pulumi.IntInput) VpcDhcpOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcDhcpOptions {
		return vs[0].([]VpcDhcpOptions)[vs[1].(int)]
	}).(VpcDhcpOptionsOutput)
}

type VpcDhcpOptionsMapOutput struct{ *pulumi.OutputState }

func (VpcDhcpOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcDhcpOptions)(nil))
}

func (o VpcDhcpOptionsMapOutput) ToVpcDhcpOptionsMapOutput() VpcDhcpOptionsMapOutput {
	return o
}

func (o VpcDhcpOptionsMapOutput) ToVpcDhcpOptionsMapOutputWithContext(ctx context.Context) VpcDhcpOptionsMapOutput {
	return o
}

func (o VpcDhcpOptionsMapOutput) MapIndex(k pulumi.StringInput) VpcDhcpOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcDhcpOptions {
		return vs[0].(map[string]VpcDhcpOptions)[vs[1].(string)]
	}).(VpcDhcpOptionsOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcDhcpOptionsOutput{})
	pulumi.RegisterOutputType(VpcDhcpOptionsPtrOutput{})
	pulumi.RegisterOutputType(VpcDhcpOptionsArrayOutput{})
	pulumi.RegisterOutputType(VpcDhcpOptionsMapOutput{})
}
