// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The provider type for the aws package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState

}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.Profile == nil {
		args.Profile = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_PROFILE").(string))
	}
	if args.Region == nil {
		args.Region = pulumi.StringPtr(getEnvOrDefault("", nil, "AWS_REGION", "AWS_DEFAULT_REGION").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:aws", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	AccessKey *string `pulumi:"accessKey"`
	AllowedAccountIds []string `pulumi:"allowedAccountIds"`
	AssumeRole *ProviderAssumeRole `pulumi:"assumeRole"`
	Endpoints []ProviderEndpoint `pulumi:"endpoints"`
	ForbiddenAccountIds []string `pulumi:"forbiddenAccountIds"`
	// Resource tag key prefixes to ignore across all resources.
	IgnoreTagPrefixes []string `pulumi:"ignoreTagPrefixes"`
	// Resource tag keys to ignore across all resources.
	IgnoreTags []string `pulumi:"ignoreTags"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
	Insecure *bool `pulumi:"insecure"`
	// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries *int `pulumi:"maxRetries"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	Profile *string `pulumi:"profile"`
	// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
	Region *string `pulumi:"region"`
	// Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By
	// default, the S3 client will use virtual hosted bucket addressing when possible (http://BUCKET.s3.amazonaws.com/KEY).
	// Specific to the Amazon S3 service.
	S3ForcePathStyle *bool `pulumi:"s3ForcePathStyle"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	SecretKey *string `pulumi:"secretKey"`
	// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
	SharedCredentialsFile *string `pulumi:"sharedCredentialsFile"`
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
	// available/implemented.
	SkipCredentialsValidation *bool `pulumi:"skipCredentialsValidation"`
	// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
	SkipGetEc2Platforms *bool `pulumi:"skipGetEc2Platforms"`
	SkipMetadataApiCheck *bool `pulumi:"skipMetadataApiCheck"`
	// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
	// not public (yet).
	SkipRegionValidation *bool `pulumi:"skipRegionValidation"`
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	SkipRequestingAccountId *bool `pulumi:"skipRequestingAccountId"`
	// session token. A session token is only required if you are using temporary security credentials.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	AccessKey pulumi.StringPtrInput
	AllowedAccountIds pulumi.StringArrayInput
	AssumeRole ProviderAssumeRolePtrInput
	Endpoints ProviderEndpointArrayInput
	ForbiddenAccountIds pulumi.StringArrayInput
	// Resource tag key prefixes to ignore across all resources.
	IgnoreTagPrefixes pulumi.StringArrayInput
	// Resource tag keys to ignore across all resources.
	IgnoreTags pulumi.StringArrayInput
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted,default value is `false`
	Insecure pulumi.BoolPtrInput
	// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
	MaxRetries pulumi.IntPtrInput
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	Profile pulumi.StringPtrInput
	// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
	Region pulumi.StringPtrInput
	// Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By
	// default, the S3 client will use virtual hosted bucket addressing when possible (http://BUCKET.s3.amazonaws.com/KEY).
	// Specific to the Amazon S3 service.
	S3ForcePathStyle pulumi.BoolPtrInput
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	SecretKey pulumi.StringPtrInput
	// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
	SharedCredentialsFile pulumi.StringPtrInput
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
	// available/implemented.
	SkipCredentialsValidation pulumi.BoolPtrInput
	// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
	SkipGetEc2Platforms pulumi.BoolPtrInput
	SkipMetadataApiCheck pulumi.BoolPtrInput
	// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
	// not public (yet).
	SkipRegionValidation pulumi.BoolPtrInput
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	SkipRequestingAccountId pulumi.BoolPtrInput
	// session token. A session token is only required if you are using temporary security credentials.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

