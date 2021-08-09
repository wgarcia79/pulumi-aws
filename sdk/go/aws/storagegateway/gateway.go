// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.
//
// > **NOTE:** The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).
//
// ## Example Usage
// ### Local Cache
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testVolumeAttachment, err := ec2.NewVolumeAttachment(ctx, "testVolumeAttachment", &ec2.VolumeAttachmentArgs{
// 			DeviceName: pulumi.String("/dev/xvdb"),
// 			VolumeId:   pulumi.Any(aws_ebs_volume.Test.Id),
// 			InstanceId: pulumi.Any(aws_instance.Test.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storagegateway.NewCache(ctx, "testCache", &storagegateway.CacheArgs{
// 			DiskId: testLocalDisk.ApplyT(func(testLocalDisk storagegateway.GetLocalDiskResult) (string, error) {
// 				return testLocalDisk.DiskId, nil
// 			}).(pulumi.StringOutput),
// 			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Test.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### FSx File Gateway
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("FILE_FSX_SMB"),
// 			SmbActiveDirectorySettings: &storagegateway.GatewaySmbActiveDirectorySettingsArgs{
// 				DomainName: pulumi.String("corp.example.com"),
// 				Password:   pulumi.String("avoid-plaintext-passwords"),
// 				Username:   pulumi.String("Admin"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### S3 File Gateway
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("FILE_S3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Tape Gateway
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress:  pulumi.String("1.2.3.4"),
// 			GatewayName:       pulumi.String("example"),
// 			GatewayTimezone:   pulumi.String("GMT"),
// 			GatewayType:       pulumi.String("VTL"),
// 			MediumChangerType: pulumi.String("AWS-Gateway-VTL"),
// 			TapeDriveType:     pulumi.String("IBM-ULT3580-TD5"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Volume Gateway (Cached)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("CACHED"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Volume Gateway (Stored)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
// 			GatewayIpAddress: pulumi.String("1.2.3.4"),
// 			GatewayName:      pulumi.String("example"),
// 			GatewayTimezone:  pulumi.String("GMT"),
// 			GatewayType:      pulumi.String("STORED"),
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
// `aws_storagegateway_gateway` can be imported by using the gateway Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:storagegateway/gateway:Gateway example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
// ```
//
//  Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the provider configuration or use `ignoreChanges` to hide the difference.
type Gateway struct {
	pulumi.CustomResourceState

	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringOutput `pulumi:"activationKey"`
	// Amazon Resource Name (ARN) of the gateway.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrOutput `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrOutput `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrOutput `pulumi:"cloudwatchLogGroupArn"`
	// The ID of the Amazon EC2 instance that was used to launch the gateway.
	Ec2InstanceId pulumi.StringOutput `pulumi:"ec2InstanceId"`
	// The type of endpoint for your gateway.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// Identifier of the gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringOutput `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName pulumi.StringOutput `pulumi:"gatewayName"`
	// An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
	GatewayNetworkInterfaces GatewayGatewayNetworkInterfaceArrayOutput `pulumi:"gatewayNetworkInterfaces"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringOutput `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrOutput `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrOutput `pulumi:"gatewayVpcEndpoint"`
	// The type of hypervisor environment used by the host.
	HostEnvironment pulumi.StringOutput `pulumi:"hostEnvironment"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType pulumi.StringPtrOutput `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrOutput `pulumi:"smbActiveDirectorySettings"`
	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility pulumi.BoolPtrOutput `pulumi:"smbFileShareVisibility"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrOutput `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringOutput `pulumi:"smbSecurityStrategy"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrOutput `pulumi:"tapeDriveType"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.GatewayTimezone == nil {
		return nil, errors.New("invalid value for required argument 'GatewayTimezone'")
	}
	var resource Gateway
	err := ctx.RegisterResource("aws:storagegateway/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("aws:storagegateway/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey *string `pulumi:"activationKey"`
	// Amazon Resource Name (ARN) of the gateway.
	Arn *string `pulumi:"arn"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec *int `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec *int `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// The ID of the Amazon EC2 instance that was used to launch the gateway.
	Ec2InstanceId *string `pulumi:"ec2InstanceId"`
	// The type of endpoint for your gateway.
	EndpointType *string `pulumi:"endpointType"`
	// Identifier of the gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName *string `pulumi:"gatewayName"`
	// An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
	GatewayNetworkInterfaces []GatewayGatewayNetworkInterface `pulumi:"gatewayNetworkInterfaces"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone *string `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType *string `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint *string `pulumi:"gatewayVpcEndpoint"`
	// The type of hypervisor environment used by the host.
	HostEnvironment *string `pulumi:"hostEnvironment"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType *string `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings *GatewaySmbActiveDirectorySettings `pulumi:"smbActiveDirectorySettings"`
	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility *bool `pulumi:"smbFileShareVisibility"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword *string `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy *string `pulumi:"smbSecurityStrategy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType *string `pulumi:"tapeDriveType"`
}

type GatewayState struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the gateway.
	Arn pulumi.StringPtrInput
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// The ID of the Amazon EC2 instance that was used to launch the gateway.
	Ec2InstanceId pulumi.StringPtrInput
	// The type of endpoint for your gateway.
	EndpointType pulumi.StringPtrInput
	// Identifier of the gateway.
	GatewayId pulumi.StringPtrInput
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringPtrInput
	// Name of the gateway.
	GatewayName pulumi.StringPtrInput
	// An array that contains descriptions of the gateway network interfaces. See Gateway Network Interface.
	GatewayNetworkInterfaces GatewayGatewayNetworkInterfaceArrayInput
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringPtrInput
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrInput
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrInput
	// The type of hypervisor environment used by the host.
	HostEnvironment pulumi.StringPtrInput
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType pulumi.StringPtrInput
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrInput
	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility pulumi.BoolPtrInput
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrInput
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey *string `pulumi:"activationKey"`
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec *int `pulumi:"averageDownloadRateLimitInBitsPerSec"`
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec *int `pulumi:"averageUploadRateLimitInBitsPerSec"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn *string `pulumi:"cloudwatchLogGroupArn"`
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Name of the gateway.
	GatewayName string `pulumi:"gatewayName"`
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone string `pulumi:"gatewayTimezone"`
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType *string `pulumi:"gatewayType"`
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint *string `pulumi:"gatewayVpcEndpoint"`
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType *string `pulumi:"mediumChangerType"`
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings *GatewaySmbActiveDirectorySettings `pulumi:"smbActiveDirectorySettings"`
	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility *bool `pulumi:"smbFileShareVisibility"`
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword *string `pulumi:"smbGuestPassword"`
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy *string `pulumi:"smbSecurityStrategy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType *string `pulumi:"tapeDriveType"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// Gateway activation key during resource creation. Conflicts with `gatewayIpAddress`. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	ActivationKey pulumi.StringPtrInput
	// The average download bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageDownloadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The average upload bandwidth rate limit in bits per second. This is supported for the `CACHED`, `STORED`, and `VTL` gateway types.
	AverageUploadRateLimitInBitsPerSec pulumi.IntPtrInput
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group to use to monitor and log events in the gateway.
	CloudwatchLogGroupArn pulumi.StringPtrInput
	// Gateway IP address to retrieve activation key during resource creation. Conflicts with `activationKey`. Gateway must be accessible on port 80 from where this provider is running. Additional information is available in the [Storage Gateway User Guide](https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html).
	GatewayIpAddress pulumi.StringPtrInput
	// Name of the gateway.
	GatewayName pulumi.StringInput
	// Time zone for the gateway. The time zone is of the format "GMT", "GMT-hr:mm", or "GMT+hr:mm". For example, `GMT-4:00` indicates the time is 4 hours behind GMT. The time zone is used, for example, for scheduling snapshots and your gateway's maintenance schedule.
	GatewayTimezone pulumi.StringInput
	// Type of the gateway. The default value is `STORED`. Valid values: `CACHED`, `FILE_FSX_SMB`, `FILE_S3`, `STORED`, `VTL`.
	GatewayType pulumi.StringPtrInput
	// VPC endpoint address to be used when activating your gateway. This should be used when your instance is in a private subnet. Requires HTTP access from client computer running this provider. More info on what ports are required by your VPC Endpoint Security group in [Activating a Gateway in a Virtual Private Cloud](https://docs.aws.amazon.com/storagegateway/latest/userguide/gateway-private-link.html).
	GatewayVpcEndpoint pulumi.StringPtrInput
	// Type of medium changer to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `STK-L700`, `AWS-Gateway-VTL`, `IBM-03584L32-0402`.
	MediumChangerType pulumi.StringPtrInput
	// Nested argument with Active Directory domain join information for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `ActiveDirectory` authentication SMB file shares. More details below.
	SmbActiveDirectorySettings GatewaySmbActiveDirectorySettingsPtrInput
	// Specifies whether the shares on this gateway appear when listing shares.
	SmbFileShareVisibility pulumi.BoolPtrInput
	// Guest password for Server Message Block (SMB) file shares. Only valid for `FILE_S3` and `FILE_FSX_SMB` gateway types. Must be set before creating `GuestAccess` authentication SMB file shares. This provider can only detect drift of the existence of a guest password, not its actual value from the gateway. This provider can however update the password with changing the argument.
	SmbGuestPassword pulumi.StringPtrInput
	// Specifies the type of security strategy. Valid values are: `ClientSpecified`, `MandatorySigning`, and `MandatoryEncryption`. See [Setting a Security Level for Your Gateway](https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-gateway-file.html#security-strategy) for more information.
	SmbSecurityStrategy pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Type of tape drive to use for tape gateway. This provider cannot detect drift of this argument. Valid values: `IBM-ULT3580-TD5`.
	TapeDriveType pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

func (i *Gateway) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

type GatewayPtrInput interface {
	pulumi.Input

	ToGatewayPtrOutput() GatewayPtrOutput
	ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput
}

type gatewayPtrType GatewayArgs

func (*gatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (i *gatewayPtrType) ToGatewayPtrOutput() GatewayPtrOutput {
	return i.ToGatewayPtrOutputWithContext(context.Background())
}

func (i *gatewayPtrType) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayPtrOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//          GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//          GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o.ToGatewayPtrOutputWithContext(context.Background())
}

func (o GatewayOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Gateway) *Gateway {
		return &v
	}).(GatewayPtrOutput)
}

type GatewayPtrOutput struct{ *pulumi.OutputState }

func (GatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil))
}

func (o GatewayPtrOutput) ToGatewayPtrOutput() GatewayPtrOutput {
	return o
}

func (o GatewayPtrOutput) ToGatewayPtrOutputWithContext(ctx context.Context) GatewayPtrOutput {
	return o
}

func (o GatewayPtrOutput) Elem() GatewayOutput {
	return o.ApplyT(func(v *Gateway) Gateway {
		if v != nil {
			return *v
		}
		var ret Gateway
		return ret
	}).(GatewayOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Gateway)(nil))
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].([]Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Gateway)(nil))
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Gateway {
		return vs[0].(map[string]Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayPtrOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
