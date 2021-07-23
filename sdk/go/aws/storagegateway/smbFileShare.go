// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AWS Storage Gateway SMB File Share.
//
// ## Example Usage
// ### Active Directory Authentication
//
// > **NOTE:** The gateway must have already joined the Active Directory domain prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbActiveDirectorySettings` in the `storagegateway.Gateway` resource.
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
// 		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
// 			Authentication: pulumi.String("ActiveDirectory"),
// 			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
// 			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
// 			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Guest Authentication
//
// > **NOTE:** The gateway must have already had the SMB guest password set prior to SMB file share creation. e.g. via "SMB Settings" in the AWS Storage Gateway console or `smbGuestPassword` in the `storagegateway.Gateway` resource.
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
// 		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
// 			Authentication: pulumi.String("GuestAccess"),
// 			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
// 			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
// 			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
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
// `aws_storagegateway_smb_file_share` can be imported by using the SMB File Share Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:storagegateway/smbFileShare:SmbFileShare example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
// ```
type SmbFileShare struct {
	pulumi.CustomResourceState

	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration pulumi.BoolPtrOutput `pulumi:"accessBasedEnumeration"`
	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists pulumi.StringArrayOutput `pulumi:"adminUserLists"`
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn pulumi.StringPtrOutput `pulumi:"auditDestinationArn"`
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrOutput `pulumi:"authentication"`
	// The region of the S3 buck used by the file share. Required when specifying a `vpcEndpointDnsName`.
	BucketRegion pulumi.StringPtrOutput `pulumi:"bucketRegion"`
	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes SmbFileShareCacheAttributesPtrOutput `pulumi:"cacheAttributes"`
	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity pulumi.StringPtrOutput `pulumi:"caseSensitivity"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrOutput `pulumi:"defaultStorageClass"`
	// The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
	FileShareName pulumi.StringOutput `pulumi:"fileShareName"`
	// ID of the SMB File Share.
	FileshareId pulumi.StringOutput `pulumi:"fileshareId"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrOutput `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayOutput `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrOutput `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrOutput `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringOutput `pulumi:"locationArn"`
	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy pulumi.StringPtrOutput `pulumi:"notificationPolicy"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrOutput `pulumi:"objectAcl"`
	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled pulumi.BoolPtrOutput `pulumi:"oplocksEnabled"`
	// File share path used by the NFS client to identify the mount point.
	Path pulumi.StringOutput `pulumi:"path"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrOutput `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrOutput `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled pulumi.BoolPtrOutput `pulumi:"smbAclEnabled"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayOutput `pulumi:"validUserLists"`
	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName pulumi.StringPtrOutput `pulumi:"vpcEndpointDnsName"`
}

// NewSmbFileShare registers a new resource with the given unique name, arguments, and options.
func NewSmbFileShare(ctx *pulumi.Context,
	name string, args *SmbFileShareArgs, opts ...pulumi.ResourceOption) (*SmbFileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayArn == nil {
		return nil, errors.New("invalid value for required argument 'GatewayArn'")
	}
	if args.LocationArn == nil {
		return nil, errors.New("invalid value for required argument 'LocationArn'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	var resource SmbFileShare
	err := ctx.RegisterResource("aws:storagegateway/smbFileShare:SmbFileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmbFileShare gets an existing SmbFileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmbFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmbFileShareState, opts ...pulumi.ResourceOption) (*SmbFileShare, error) {
	var resource SmbFileShare
	err := ctx.ReadResource("aws:storagegateway/smbFileShare:SmbFileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmbFileShare resources.
type smbFileShareState struct {
	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration *bool `pulumi:"accessBasedEnumeration"`
	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists []string `pulumi:"adminUserLists"`
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn *string `pulumi:"arn"`
	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn *string `pulumi:"auditDestinationArn"`
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication *string `pulumi:"authentication"`
	// The region of the S3 buck used by the file share. Required when specifying a `vpcEndpointDnsName`.
	BucketRegion *string `pulumi:"bucketRegion"`
	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes *SmbFileShareCacheAttributes `pulumi:"cacheAttributes"`
	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass *string `pulumi:"defaultStorageClass"`
	// The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
	FileShareName *string `pulumi:"fileShareName"`
	// ID of the SMB File Share.
	FileshareId *string `pulumi:"fileshareId"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled *bool `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists []string `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn *string `pulumi:"locationArn"`
	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy *string `pulumi:"notificationPolicy"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl *string `pulumi:"objectAcl"`
	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled *bool `pulumi:"oplocksEnabled"`
	// File share path used by the NFS client to identify the mount point.
	Path *string `pulumi:"path"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays *bool `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn *string `pulumi:"roleArn"`
	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled *bool `pulumi:"smbAclEnabled"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists []string `pulumi:"validUserLists"`
	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName *string `pulumi:"vpcEndpointDnsName"`
}

type SmbFileShareState struct {
	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration pulumi.BoolPtrInput
	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists pulumi.StringArrayInput
	// Amazon Resource Name (ARN) of the SMB File Share.
	Arn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn pulumi.StringPtrInput
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrInput
	// The region of the S3 buck used by the file share. Required when specifying a `vpcEndpointDnsName`.
	BucketRegion pulumi.StringPtrInput
	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes SmbFileShareCacheAttributesPtrInput
	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity pulumi.StringPtrInput
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrInput
	// The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
	FileShareName pulumi.StringPtrInput
	// ID of the SMB File Share.
	FileshareId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringPtrInput
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrInput
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayInput
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrInput
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringPtrInput
	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy pulumi.StringPtrInput
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrInput
	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled pulumi.BoolPtrInput
	// File share path used by the NFS client to identify the mount point.
	Path pulumi.StringPtrInput
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrInput
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringPtrInput
	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled pulumi.BoolPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayInput
	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName pulumi.StringPtrInput
}

func (SmbFileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*smbFileShareState)(nil)).Elem()
}

type smbFileShareArgs struct {
	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration *bool `pulumi:"accessBasedEnumeration"`
	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists []string `pulumi:"adminUserLists"`
	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn *string `pulumi:"auditDestinationArn"`
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication *string `pulumi:"authentication"`
	// The region of the S3 buck used by the file share. Required when specifying a `vpcEndpointDnsName`.
	BucketRegion *string `pulumi:"bucketRegion"`
	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes *SmbFileShareCacheAttributes `pulumi:"cacheAttributes"`
	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass *string `pulumi:"defaultStorageClass"`
	// The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
	FileShareName *string `pulumi:"fileShareName"`
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn string `pulumi:"gatewayArn"`
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled *bool `pulumi:"guessMimeTypeEnabled"`
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists []string `pulumi:"invalidUserLists"`
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The ARN of the backed storage used for storing file data.
	LocationArn string `pulumi:"locationArn"`
	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy *string `pulumi:"notificationPolicy"`
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl *string `pulumi:"objectAcl"`
	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled *bool `pulumi:"oplocksEnabled"`
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly *bool `pulumi:"readOnly"`
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays *bool `pulumi:"requesterPays"`
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn string `pulumi:"roleArn"`
	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled *bool `pulumi:"smbAclEnabled"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists []string `pulumi:"validUserLists"`
	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName *string `pulumi:"vpcEndpointDnsName"`
}

// The set of arguments for constructing a SmbFileShare resource.
type SmbFileShareArgs struct {
	// The files and folders on this share will only be visible to users with read access. Default value is `false`.
	AccessBasedEnumeration pulumi.BoolPtrInput
	// A list of users in the Active Directory that have admin access to the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	AdminUserLists pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) of the CloudWatch Log Group used for the audit logs.
	AuditDestinationArn pulumi.StringPtrInput
	// The authentication method that users use to access the file share. Defaults to `ActiveDirectory`. Valid values: `ActiveDirectory`, `GuestAccess`.
	Authentication pulumi.StringPtrInput
	// The region of the S3 buck used by the file share. Required when specifying a `vpcEndpointDnsName`.
	BucketRegion pulumi.StringPtrInput
	// Refresh cache information. see Cache Attributes for more details.
	CacheAttributes SmbFileShareCacheAttributesPtrInput
	// The case of an object name in an Amazon S3 bucket. For `ClientSpecified`, the client determines the case sensitivity. For `CaseSensitive`, the gateway determines the case sensitivity. The default value is `ClientSpecified`.
	CaseSensitivity pulumi.StringPtrInput
	// The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
	DefaultStorageClass pulumi.StringPtrInput
	// The name of the file share. Must be set if an S3 prefix name is set in `locationArn`.
	FileShareName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the file gateway.
	GatewayArn pulumi.StringInput
	// Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
	GuessMimeTypeEnabled pulumi.BoolPtrInput
	// A list of users in the Active Directory that are not allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	InvalidUserLists pulumi.StringArrayInput
	// Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
	KmsEncrypted pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is true.
	KmsKeyArn pulumi.StringPtrInput
	// The ARN of the backed storage used for storing file data.
	LocationArn pulumi.StringInput
	// The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
	NotificationPolicy pulumi.StringPtrInput
	// Access Control List permission for S3 bucket objects. Defaults to `private`.
	ObjectAcl pulumi.StringPtrInput
	// Boolean to indicate Opportunistic lock (oplock) status. Defaults to `true`.
	OplocksEnabled pulumi.BoolPtrInput
	// Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
	ReadOnly pulumi.BoolPtrInput
	// Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
	RequesterPays pulumi.BoolPtrInput
	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
	RoleArn pulumi.StringInput
	// Set this value to `true` to enable ACL (access control list) on the SMB fileshare. Set it to `false` to map file and directory permissions to the POSIX permissions. This setting applies only to `ActiveDirectory` authentication type.
	SmbAclEnabled pulumi.BoolPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// A list of users in the Active Directory that are allowed to access the file share. Only valid if `authentication` is set to `ActiveDirectory`.
	ValidUserLists pulumi.StringArrayInput
	// The DNS name of the VPC endpoint for S3 private link.
	VpcEndpointDnsName pulumi.StringPtrInput
}

func (SmbFileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smbFileShareArgs)(nil)).Elem()
}

type SmbFileShareInput interface {
	pulumi.Input

	ToSmbFileShareOutput() SmbFileShareOutput
	ToSmbFileShareOutputWithContext(ctx context.Context) SmbFileShareOutput
}

func (*SmbFileShare) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbFileShare)(nil))
}

func (i *SmbFileShare) ToSmbFileShareOutput() SmbFileShareOutput {
	return i.ToSmbFileShareOutputWithContext(context.Background())
}

func (i *SmbFileShare) ToSmbFileShareOutputWithContext(ctx context.Context) SmbFileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbFileShareOutput)
}

func (i *SmbFileShare) ToSmbFileSharePtrOutput() SmbFileSharePtrOutput {
	return i.ToSmbFileSharePtrOutputWithContext(context.Background())
}

func (i *SmbFileShare) ToSmbFileSharePtrOutputWithContext(ctx context.Context) SmbFileSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbFileSharePtrOutput)
}

type SmbFileSharePtrInput interface {
	pulumi.Input

	ToSmbFileSharePtrOutput() SmbFileSharePtrOutput
	ToSmbFileSharePtrOutputWithContext(ctx context.Context) SmbFileSharePtrOutput
}

type smbFileSharePtrType SmbFileShareArgs

func (*smbFileSharePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbFileShare)(nil))
}

func (i *smbFileSharePtrType) ToSmbFileSharePtrOutput() SmbFileSharePtrOutput {
	return i.ToSmbFileSharePtrOutputWithContext(context.Background())
}

func (i *smbFileSharePtrType) ToSmbFileSharePtrOutputWithContext(ctx context.Context) SmbFileSharePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbFileSharePtrOutput)
}

// SmbFileShareArrayInput is an input type that accepts SmbFileShareArray and SmbFileShareArrayOutput values.
// You can construct a concrete instance of `SmbFileShareArrayInput` via:
//
//          SmbFileShareArray{ SmbFileShareArgs{...} }
type SmbFileShareArrayInput interface {
	pulumi.Input

	ToSmbFileShareArrayOutput() SmbFileShareArrayOutput
	ToSmbFileShareArrayOutputWithContext(context.Context) SmbFileShareArrayOutput
}

type SmbFileShareArray []SmbFileShareInput

func (SmbFileShareArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SmbFileShare)(nil))
}

func (i SmbFileShareArray) ToSmbFileShareArrayOutput() SmbFileShareArrayOutput {
	return i.ToSmbFileShareArrayOutputWithContext(context.Background())
}

func (i SmbFileShareArray) ToSmbFileShareArrayOutputWithContext(ctx context.Context) SmbFileShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbFileShareArrayOutput)
}

// SmbFileShareMapInput is an input type that accepts SmbFileShareMap and SmbFileShareMapOutput values.
// You can construct a concrete instance of `SmbFileShareMapInput` via:
//
//          SmbFileShareMap{ "key": SmbFileShareArgs{...} }
type SmbFileShareMapInput interface {
	pulumi.Input

	ToSmbFileShareMapOutput() SmbFileShareMapOutput
	ToSmbFileShareMapOutputWithContext(context.Context) SmbFileShareMapOutput
}

type SmbFileShareMap map[string]SmbFileShareInput

func (SmbFileShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SmbFileShare)(nil))
}

func (i SmbFileShareMap) ToSmbFileShareMapOutput() SmbFileShareMapOutput {
	return i.ToSmbFileShareMapOutputWithContext(context.Background())
}

func (i SmbFileShareMap) ToSmbFileShareMapOutputWithContext(ctx context.Context) SmbFileShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmbFileShareMapOutput)
}

type SmbFileShareOutput struct {
	*pulumi.OutputState
}

func (SmbFileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmbFileShare)(nil))
}

func (o SmbFileShareOutput) ToSmbFileShareOutput() SmbFileShareOutput {
	return o
}

func (o SmbFileShareOutput) ToSmbFileShareOutputWithContext(ctx context.Context) SmbFileShareOutput {
	return o
}

func (o SmbFileShareOutput) ToSmbFileSharePtrOutput() SmbFileSharePtrOutput {
	return o.ToSmbFileSharePtrOutputWithContext(context.Background())
}

func (o SmbFileShareOutput) ToSmbFileSharePtrOutputWithContext(ctx context.Context) SmbFileSharePtrOutput {
	return o.ApplyT(func(v SmbFileShare) *SmbFileShare {
		return &v
	}).(SmbFileSharePtrOutput)
}

type SmbFileSharePtrOutput struct {
	*pulumi.OutputState
}

func (SmbFileSharePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmbFileShare)(nil))
}

func (o SmbFileSharePtrOutput) ToSmbFileSharePtrOutput() SmbFileSharePtrOutput {
	return o
}

func (o SmbFileSharePtrOutput) ToSmbFileSharePtrOutputWithContext(ctx context.Context) SmbFileSharePtrOutput {
	return o
}

type SmbFileShareArrayOutput struct{ *pulumi.OutputState }

func (SmbFileShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmbFileShare)(nil))
}

func (o SmbFileShareArrayOutput) ToSmbFileShareArrayOutput() SmbFileShareArrayOutput {
	return o
}

func (o SmbFileShareArrayOutput) ToSmbFileShareArrayOutputWithContext(ctx context.Context) SmbFileShareArrayOutput {
	return o
}

func (o SmbFileShareArrayOutput) Index(i pulumi.IntInput) SmbFileShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SmbFileShare {
		return vs[0].([]SmbFileShare)[vs[1].(int)]
	}).(SmbFileShareOutput)
}

type SmbFileShareMapOutput struct{ *pulumi.OutputState }

func (SmbFileShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SmbFileShare)(nil))
}

func (o SmbFileShareMapOutput) ToSmbFileShareMapOutput() SmbFileShareMapOutput {
	return o
}

func (o SmbFileShareMapOutput) ToSmbFileShareMapOutputWithContext(ctx context.Context) SmbFileShareMapOutput {
	return o
}

func (o SmbFileShareMapOutput) MapIndex(k pulumi.StringInput) SmbFileShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SmbFileShare {
		return vs[0].(map[string]SmbFileShare)[vs[1].(string)]
	}).(SmbFileShareOutput)
}

func init() {
	pulumi.RegisterOutputType(SmbFileShareOutput{})
	pulumi.RegisterOutputType(SmbFileSharePtrOutput{})
	pulumi.RegisterOutputType(SmbFileShareArrayOutput{})
	pulumi.RegisterOutputType(SmbFileShareMapOutput{})
}
