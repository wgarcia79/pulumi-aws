// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudfront

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Amazon CloudFront web distribution.
//
// For information about CloudFront distributions, see the
// [Amazon CloudFront Developer Guide][1]. For specific information about creating
// CloudFront web distributions, see the [POST Distribution][2] page in the Amazon
// CloudFront API Reference.
//
// > **NOTE:** CloudFront distributions take about 15 minutes to a deployed state
// after creation or modification. During this time, deletes to resources will be
// blocked. If you need to delete a distribution that is enabled and you do not
// want to wait, you need to use the `retainOnDelete` flag.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudfront_distribution.html.markdown.
type Distribution struct {
	pulumi.CustomResourceState

	// The key pair IDs that CloudFront is aware of for
	// each trusted signer, if the distribution is set up to serve private content
	// with signed URLs.
	ActiveTrustedSigners pulumi.MapOutput `pulumi:"activeTrustedSigners"`
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases pulumi.StringArrayOutput `pulumi:"aliases"`
	// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Internal value used by CloudFront to allow future
	// updates to the distribution configuration.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// Any comments you want to include about the
	// distribution.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses DistributionCustomErrorResponseArrayOutput `pulumi:"customErrorResponses"`
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior DistributionDefaultCacheBehaviorOutput `pulumi:"defaultCacheBehavior"`
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject pulumi.StringPtrOutput `pulumi:"defaultRootObject"`
	// The DNS domain name of either the S3 bucket, or
	// web site of your custom origin.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The current version of the distribution's information. For example:
	// `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The CloudFront Route 53 zone ID that can be used to
	// route an [Alias Resource Record Set][7] to. This attribute is simply an
	// alias for the zone ID `Z2FDTNDATAQYW2`.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion pulumi.StringPtrOutput `pulumi:"httpVersion"`
	// The number of invalidation batches
	// currently in progress.
	InProgressValidationBatches pulumi.IntOutput `pulumi:"inProgressValidationBatches"`
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled pulumi.BoolPtrOutput `pulumi:"isIpv6Enabled"`
	// The date and time the distribution was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig DistributionLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors DistributionOrderedCacheBehaviorArrayOutput `pulumi:"orderedCacheBehaviors"`
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups DistributionOriginGroupArrayOutput `pulumi:"originGroups"`
	// One or more origins for this
	// distribution (multiples allowed).
	Origins DistributionOriginArrayOutput `pulumi:"origins"`
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass pulumi.StringPtrOutput `pulumi:"priceClass"`
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions DistributionRestrictionsOutput `pulumi:"restrictions"`
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete pulumi.BoolPtrOutput `pulumi:"retainOnDelete"`
	// The current status of the distribution. `Deployed` if the
	// distribution's information is fully propagated throughout the Amazon
	// CloudFront system.
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate DistributionViewerCertificateOutput `pulumi:"viewerCertificate"`
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment pulumi.BoolPtrOutput `pulumi:"waitForDeployment"`
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId pulumi.StringPtrOutput `pulumi:"webAclId"`
}

// NewDistribution registers a new resource with the given unique name, arguments, and options.
func NewDistribution(ctx *pulumi.Context,
	name string, args *DistributionArgs, opts ...pulumi.ResourceOption) (*Distribution, error) {
	if args == nil || args.DefaultCacheBehavior == nil {
		return nil, errors.New("missing required argument 'DefaultCacheBehavior'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.Origins == nil {
		return nil, errors.New("missing required argument 'Origins'")
	}
	if args == nil || args.Restrictions == nil {
		return nil, errors.New("missing required argument 'Restrictions'")
	}
	if args == nil || args.ViewerCertificate == nil {
		return nil, errors.New("missing required argument 'ViewerCertificate'")
	}
	if args == nil {
		args = &DistributionArgs{}
	}
	var resource Distribution
	err := ctx.RegisterResource("aws:cloudfront/distribution:Distribution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDistribution gets an existing Distribution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDistribution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DistributionState, opts ...pulumi.ResourceOption) (*Distribution, error) {
	var resource Distribution
	err := ctx.ReadResource("aws:cloudfront/distribution:Distribution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Distribution resources.
type distributionState struct {
	// The key pair IDs that CloudFront is aware of for
	// each trusted signer, if the distribution is set up to serve private content
	// with signed URLs.
	ActiveTrustedSigners map[string]interface{} `pulumi:"activeTrustedSigners"`
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases []string `pulumi:"aliases"`
	// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
	Arn *string `pulumi:"arn"`
	// Internal value used by CloudFront to allow future
	// updates to the distribution configuration.
	CallerReference *string `pulumi:"callerReference"`
	// Any comments you want to include about the
	// distribution.
	Comment *string `pulumi:"comment"`
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses []DistributionCustomErrorResponse `pulumi:"customErrorResponses"`
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior *DistributionDefaultCacheBehavior `pulumi:"defaultCacheBehavior"`
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject *string `pulumi:"defaultRootObject"`
	// The DNS domain name of either the S3 bucket, or
	// web site of your custom origin.
	DomainName *string `pulumi:"domainName"`
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled *bool `pulumi:"enabled"`
	// The current version of the distribution's information. For example:
	// `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// The CloudFront Route 53 zone ID that can be used to
	// route an [Alias Resource Record Set][7] to. This attribute is simply an
	// alias for the zone ID `Z2FDTNDATAQYW2`.
	HostedZoneId *string `pulumi:"hostedZoneId"`
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion *string `pulumi:"httpVersion"`
	// The number of invalidation batches
	// currently in progress.
	InProgressValidationBatches *int `pulumi:"inProgressValidationBatches"`
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled *bool `pulumi:"isIpv6Enabled"`
	// The date and time the distribution was last modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig *DistributionLoggingConfig `pulumi:"loggingConfig"`
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors []DistributionOrderedCacheBehavior `pulumi:"orderedCacheBehaviors"`
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups []DistributionOriginGroup `pulumi:"originGroups"`
	// One or more origins for this
	// distribution (multiples allowed).
	Origins []DistributionOrigin `pulumi:"origins"`
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass *string `pulumi:"priceClass"`
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions *DistributionRestrictions `pulumi:"restrictions"`
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete *bool `pulumi:"retainOnDelete"`
	// The current status of the distribution. `Deployed` if the
	// distribution's information is fully propagated throughout the Amazon
	// CloudFront system.
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate *DistributionViewerCertificate `pulumi:"viewerCertificate"`
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment *bool `pulumi:"waitForDeployment"`
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId *string `pulumi:"webAclId"`
}

type DistributionState struct {
	// The key pair IDs that CloudFront is aware of for
	// each trusted signer, if the distribution is set up to serve private content
	// with signed URLs.
	ActiveTrustedSigners pulumi.MapInput
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases pulumi.StringArrayInput
	// The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
	Arn pulumi.StringPtrInput
	// Internal value used by CloudFront to allow future
	// updates to the distribution configuration.
	CallerReference pulumi.StringPtrInput
	// Any comments you want to include about the
	// distribution.
	Comment pulumi.StringPtrInput
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses DistributionCustomErrorResponseArrayInput
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior DistributionDefaultCacheBehaviorPtrInput
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject pulumi.StringPtrInput
	// The DNS domain name of either the S3 bucket, or
	// web site of your custom origin.
	DomainName pulumi.StringPtrInput
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled pulumi.BoolPtrInput
	// The current version of the distribution's information. For example:
	// `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// The CloudFront Route 53 zone ID that can be used to
	// route an [Alias Resource Record Set][7] to. This attribute is simply an
	// alias for the zone ID `Z2FDTNDATAQYW2`.
	HostedZoneId pulumi.StringPtrInput
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion pulumi.StringPtrInput
	// The number of invalidation batches
	// currently in progress.
	InProgressValidationBatches pulumi.IntPtrInput
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled pulumi.BoolPtrInput
	// The date and time the distribution was last modified.
	LastModifiedTime pulumi.StringPtrInput
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig DistributionLoggingConfigPtrInput
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors DistributionOrderedCacheBehaviorArrayInput
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups DistributionOriginGroupArrayInput
	// One or more origins for this
	// distribution (multiples allowed).
	Origins DistributionOriginArrayInput
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass pulumi.StringPtrInput
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions DistributionRestrictionsPtrInput
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete pulumi.BoolPtrInput
	// The current status of the distribution. `Deployed` if the
	// distribution's information is fully propagated throughout the Amazon
	// CloudFront system.
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate DistributionViewerCertificatePtrInput
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment pulumi.BoolPtrInput
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId pulumi.StringPtrInput
}

func (DistributionState) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionState)(nil)).Elem()
}

type distributionArgs struct {
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases []string `pulumi:"aliases"`
	// Any comments you want to include about the
	// distribution.
	Comment *string `pulumi:"comment"`
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses []DistributionCustomErrorResponse `pulumi:"customErrorResponses"`
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior DistributionDefaultCacheBehavior `pulumi:"defaultCacheBehavior"`
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject *string `pulumi:"defaultRootObject"`
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled bool `pulumi:"enabled"`
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion *string `pulumi:"httpVersion"`
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled *bool `pulumi:"isIpv6Enabled"`
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig *DistributionLoggingConfig `pulumi:"loggingConfig"`
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors []DistributionOrderedCacheBehavior `pulumi:"orderedCacheBehaviors"`
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups []DistributionOriginGroup `pulumi:"originGroups"`
	// One or more origins for this
	// distribution (multiples allowed).
	Origins []DistributionOrigin `pulumi:"origins"`
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass *string `pulumi:"priceClass"`
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions DistributionRestrictions `pulumi:"restrictions"`
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete *bool `pulumi:"retainOnDelete"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate DistributionViewerCertificate `pulumi:"viewerCertificate"`
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment *bool `pulumi:"waitForDeployment"`
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId *string `pulumi:"webAclId"`
}

// The set of arguments for constructing a Distribution resource.
type DistributionArgs struct {
	// Extra CNAMEs (alternate domain names), if any, for
	// this distribution.
	Aliases pulumi.StringArrayInput
	// Any comments you want to include about the
	// distribution.
	Comment pulumi.StringPtrInput
	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses DistributionCustomErrorResponseArrayInput
	// The default cache behavior for this distribution (maximum
	// one).
	DefaultCacheBehavior DistributionDefaultCacheBehaviorInput
	// The object that you want CloudFront to
	// return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject pulumi.StringPtrInput
	// Whether the distribution is enabled to accept end
	// user requests for content.
	Enabled pulumi.BoolInput
	// The maximum HTTP version to support on the
	// distribution. Allowed values are `http1.1` and `http2`. The default is
	// `http2`.
	HttpVersion pulumi.StringPtrInput
	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled pulumi.BoolPtrInput
	// The logging
	// configuration that controls how logs are written
	// to your distribution (maximum one).
	LoggingConfig DistributionLoggingConfigPtrInput
	// An ordered list of cache behaviors
	// resource for this distribution. List from top to bottom
	// in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors DistributionOrderedCacheBehaviorArrayInput
	// One or more originGroup for this
	// distribution (multiples allowed).
	OriginGroups DistributionOriginGroupArrayInput
	// One or more origins for this
	// distribution (multiples allowed).
	Origins DistributionOriginArrayInput
	// The price class for this distribution. One of
	// `PriceClass_All`, `PriceClass_200`, `PriceClass_100`
	PriceClass pulumi.StringPtrInput
	// The restriction
	// configuration for this distribution (maximum one).
	Restrictions DistributionRestrictionsInput
	// Disables the distribution instead of
	// deleting it when destroying the resource. If this is set,
	// the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The SSL
	// configuration for this distribution (maximum
	// one).
	ViewerCertificate DistributionViewerCertificateInput
	// If enabled, the resource will wait for
	// the distribution status to change from `InProgress` to `Deployed`. Setting
	// this to`false` will skip the process. Default: `true`.
	WaitForDeployment pulumi.BoolPtrInput
	// If you're using AWS WAF to filter CloudFront
	// requests, the Id of the AWS WAF web ACL that is associated with the
	// distribution. The WAF Web ACL must exist in the WAF Global (CloudFront)
	// region and the credentials configuring this argument must have
	// `waf:GetWebACL` permissions assigned.
	WebAclId pulumi.StringPtrInput
}

func (DistributionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*distributionArgs)(nil)).Elem()
}

