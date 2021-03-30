// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/storagegateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := aws_volume_attachment.Test.Device_name
// 		_, err := storagegateway.GetLocalDisk(ctx, &storagegateway.GetLocalDiskArgs{
// 			DiskPath:   &opt0,
// 			GatewayArn: aws_storagegateway_gateway.Test.Arn,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLocalDisk(ctx *pulumi.Context, args *GetLocalDiskArgs, opts ...pulumi.InvokeOption) (*GetLocalDiskResult, error) {
	var rv GetLocalDiskResult
	err := ctx.Invoke("aws:storagegateway/getLocalDisk:getLocalDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalDisk.
type GetLocalDiskArgs struct {
	// The device node of the local disk to retrieve. For example, `/dev/sdb`.
	DiskNode *string `pulumi:"diskNode"`
	// The device path of the local disk to retrieve. For example, `/dev/xvdb` or `/dev/nvme1n1`.
	DiskPath *string `pulumi:"diskPath"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
}

// A collection of values returned by getLocalDisk.
type GetLocalDiskResult struct {
	// The disk identifier. e.g. `pci-0000:03:00.0-scsi-0:0:0:0`
	DiskId     string `pulumi:"diskId"`
	DiskNode   string `pulumi:"diskNode"`
	DiskPath   string `pulumi:"diskPath"`
	GatewayArn string `pulumi:"gatewayArn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
