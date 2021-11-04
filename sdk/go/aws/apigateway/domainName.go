// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
// can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
//
// This resource just establishes ownership of and the TLS settings for
// a particular domain name. An API can be attached to a particular path
// under the registered domain name using
// the `apigateway.BasePathMapping` resource.
//
// API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
// API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
// addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
// (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfrontDomainName`
// attribute.
//
// In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
// a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
// given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
// the `regionalDomainName` attribute.
//
// > **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the `acm.Certificate` resource.
//
// > **Note:** The `apigateway.DomainName` resource expects dependency on the `acm.CertificateValidation` as
// only verified certificates can be used. This can be made either explicitly by adding the
// `dependsOn = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN
// from the validation resource where it will be available after the resource creation:
// `regionalCertificateArn = aws_acm_certificate_validation.cert.certificate_arn`.
//
// ## Example Usage
// ### Edge Optimized (ACM Certificate)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
// 			CertificateArn: pulumi.Any(aws_acm_certificate_validation.Example.Certificate_arn),
// 			DomainName:     pulumi.String("api.example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
// 			Name:   exampleDomainName.DomainName,
// 			Type:   pulumi.String("A"),
// 			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					EvaluateTargetHealth: pulumi.Bool(true),
// 					Name:                 exampleDomainName.CloudfrontDomainName,
// 					ZoneId:               exampleDomainName.CloudfrontZoneId,
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Edge Optimized (IAM Certificate)
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
// 			DomainName:            pulumi.String("api.example.com"),
// 			CertificateName:       pulumi.String("example-api"),
// 			CertificateBody:       readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/example.crt")),
// 			CertificateChain:      readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/ca.crt")),
// 			CertificatePrivateKey: readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/example.key")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
// 			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
// 			Name:   exampleDomainName.DomainName,
// 			Type:   pulumi.String("A"),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					Name:                 exampleDomainName.CloudfrontDomainName,
// 					ZoneId:               exampleDomainName.CloudfrontZoneId,
// 					EvaluateTargetHealth: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Regional (ACM Certificate)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
// 			DomainName:             pulumi.String("api.example.com"),
// 			RegionalCertificateArn: pulumi.Any(aws_acm_certificate_validation.Example.Certificate_arn),
// 			EndpointConfiguration: &apigateway.DomainNameEndpointConfigurationArgs{
// 				Types: pulumi.String{
// 					"REGIONAL",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
// 			Name:   exampleDomainName.DomainName,
// 			Type:   pulumi.String("A"),
// 			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					EvaluateTargetHealth: pulumi.Bool(true),
// 					Name:                 exampleDomainName.RegionalDomainName,
// 					ZoneId:               exampleDomainName.RegionalZoneId,
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Regional (IAM Certificate)
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
// 			CertificateBody:         readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/example.crt")),
// 			CertificateChain:        readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/ca.crt")),
// 			CertificatePrivateKey:   readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/example.com/example.key")),
// 			DomainName:              pulumi.String("api.example.com"),
// 			RegionalCertificateName: pulumi.String("example-api"),
// 			EndpointConfiguration: &apigateway.DomainNameEndpointConfigurationArgs{
// 				Types: pulumi.String{
// 					"REGIONAL",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
// 			Name:   exampleDomainName.DomainName,
// 			Type:   pulumi.String("A"),
// 			ZoneId: pulumi.Any(aws_route53_zone.Example.Id),
// 			Aliases: route53.RecordAliasArray{
// 				&route53.RecordAliasArgs{
// 					EvaluateTargetHealth: pulumi.Bool(true),
// 					Name:                 exampleDomainName.RegionalDomainName,
// 					ZoneId:               exampleDomainName.RegionalZoneId,
// 				},
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
// API Gateway domain names can be imported using their `name`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/domainName:DomainName example dev.example.com
// ```
type DomainName struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringPtrOutput `pulumi:"certificateArn"`
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringPtrOutput `pulumi:"certificateBody"`
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringPtrOutput `pulumi:"certificateChain"`
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringPtrOutput `pulumi:"certificateName"`
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringPtrOutput `pulumi:"certificatePrivateKey"`
	// The upload date associated with the domain certificate.
	CertificateUploadDate pulumi.StringOutput `pulumi:"certificateUploadDate"`
	// The hostname created by Cloudfront to represent
	// the distribution that implements this domain name mapping.
	CloudfrontDomainName pulumi.StringOutput `pulumi:"cloudfrontDomainName"`
	// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
	// that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneId pulumi.StringOutput `pulumi:"cloudfrontZoneId"`
	// The fully-qualified domain name to register
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationOutput `pulumi:"endpointConfiguration"`
	// The mutual TLS authentication configuration for the domain name. Defined below.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrOutput `pulumi:"mutualTlsAuthentication"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringPtrOutput `pulumi:"regionalCertificateArn"`
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringPtrOutput `pulumi:"regionalCertificateName"`
	// The hostname for the custom domain's regional endpoint.
	RegionalDomainName pulumi.StringOutput `pulumi:"regionalDomainName"`
	// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneId pulumi.StringOutput `pulumi:"regionalZoneId"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringOutput `pulumi:"securityPolicy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDomainName registers a new resource with the given unique name, arguments, and options.
func NewDomainName(ctx *pulumi.Context,
	name string, args *DomainNameArgs, opts ...pulumi.ResourceOption) (*DomainName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	var resource DomainName
	err := ctx.RegisterResource("aws:apigateway/domainName:DomainName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainName gets an existing DomainName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNameState, opts ...pulumi.ResourceOption) (*DomainName, error) {
	var resource DomainName
	err := ctx.ReadResource("aws:apigateway/domainName:DomainName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainName resources.
type domainNameState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn *string `pulumi:"certificateArn"`
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody *string `pulumi:"certificateBody"`
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain *string `pulumi:"certificateChain"`
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName *string `pulumi:"certificateName"`
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey *string `pulumi:"certificatePrivateKey"`
	// The upload date associated with the domain certificate.
	CertificateUploadDate *string `pulumi:"certificateUploadDate"`
	// The hostname created by Cloudfront to represent
	// the distribution that implements this domain name mapping.
	CloudfrontDomainName *string `pulumi:"cloudfrontDomainName"`
	// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
	// that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneId *string `pulumi:"cloudfrontZoneId"`
	// The fully-qualified domain name to register
	DomainName *string `pulumi:"domainName"`
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration *DomainNameEndpointConfiguration `pulumi:"endpointConfiguration"`
	// The mutual TLS authentication configuration for the domain name. Defined below.
	MutualTlsAuthentication *DomainNameMutualTlsAuthentication `pulumi:"mutualTlsAuthentication"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn *string `pulumi:"regionalCertificateArn"`
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName *string `pulumi:"regionalCertificateName"`
	// The hostname for the custom domain's regional endpoint.
	RegionalDomainName *string `pulumi:"regionalDomainName"`
	// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneId *string `pulumi:"regionalZoneId"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy *string `pulumi:"securityPolicy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DomainNameState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringPtrInput
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringPtrInput
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringPtrInput
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringPtrInput
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringPtrInput
	// The upload date associated with the domain certificate.
	CertificateUploadDate pulumi.StringPtrInput
	// The hostname created by Cloudfront to represent
	// the distribution that implements this domain name mapping.
	CloudfrontDomainName pulumi.StringPtrInput
	// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
	// that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneId pulumi.StringPtrInput
	// The fully-qualified domain name to register
	DomainName pulumi.StringPtrInput
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationPtrInput
	// The mutual TLS authentication configuration for the domain name. Defined below.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrInput
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringPtrInput
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringPtrInput
	// The hostname for the custom domain's regional endpoint.
	RegionalDomainName pulumi.StringPtrInput
	// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneId pulumi.StringPtrInput
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
}

func (DomainNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameState)(nil)).Elem()
}

type domainNameArgs struct {
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn *string `pulumi:"certificateArn"`
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody *string `pulumi:"certificateBody"`
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain *string `pulumi:"certificateChain"`
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName *string `pulumi:"certificateName"`
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey *string `pulumi:"certificatePrivateKey"`
	// The fully-qualified domain name to register
	DomainName string `pulumi:"domainName"`
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration *DomainNameEndpointConfiguration `pulumi:"endpointConfiguration"`
	// The mutual TLS authentication configuration for the domain name. Defined below.
	MutualTlsAuthentication *DomainNameMutualTlsAuthentication `pulumi:"mutualTlsAuthentication"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn *string `pulumi:"regionalCertificateArn"`
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName *string `pulumi:"regionalCertificateName"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy *string `pulumi:"securityPolicy"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DomainName resource.
type DomainNameArgs struct {
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringPtrInput
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringPtrInput
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringPtrInput
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringPtrInput
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringPtrInput
	// The fully-qualified domain name to register
	DomainName pulumi.StringInput
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationPtrInput
	// The mutual TLS authentication configuration for the domain name. Defined below.
	MutualTlsAuthentication DomainNameMutualTlsAuthenticationPtrInput
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringPtrInput
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringPtrInput
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (DomainNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainNameArgs)(nil)).Elem()
}

type DomainNameInput interface {
	pulumi.Input

	ToDomainNameOutput() DomainNameOutput
	ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput
}

func (*DomainName) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainName)(nil))
}

func (i *DomainName) ToDomainNameOutput() DomainNameOutput {
	return i.ToDomainNameOutputWithContext(context.Background())
}

func (i *DomainName) ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameOutput)
}

func (i *DomainName) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return i.ToDomainNamePtrOutputWithContext(context.Background())
}

func (i *DomainName) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNamePtrOutput)
}

type DomainNamePtrInput interface {
	pulumi.Input

	ToDomainNamePtrOutput() DomainNamePtrOutput
	ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput
}

type domainNamePtrType DomainNameArgs

func (*domainNamePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainName)(nil))
}

func (i *domainNamePtrType) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return i.ToDomainNamePtrOutputWithContext(context.Background())
}

func (i *domainNamePtrType) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNamePtrOutput)
}

// DomainNameArrayInput is an input type that accepts DomainNameArray and DomainNameArrayOutput values.
// You can construct a concrete instance of `DomainNameArrayInput` via:
//
//          DomainNameArray{ DomainNameArgs{...} }
type DomainNameArrayInput interface {
	pulumi.Input

	ToDomainNameArrayOutput() DomainNameArrayOutput
	ToDomainNameArrayOutputWithContext(context.Context) DomainNameArrayOutput
}

type DomainNameArray []DomainNameInput

func (DomainNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainName)(nil)).Elem()
}

func (i DomainNameArray) ToDomainNameArrayOutput() DomainNameArrayOutput {
	return i.ToDomainNameArrayOutputWithContext(context.Background())
}

func (i DomainNameArray) ToDomainNameArrayOutputWithContext(ctx context.Context) DomainNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameArrayOutput)
}

// DomainNameMapInput is an input type that accepts DomainNameMap and DomainNameMapOutput values.
// You can construct a concrete instance of `DomainNameMapInput` via:
//
//          DomainNameMap{ "key": DomainNameArgs{...} }
type DomainNameMapInput interface {
	pulumi.Input

	ToDomainNameMapOutput() DomainNameMapOutput
	ToDomainNameMapOutputWithContext(context.Context) DomainNameMapOutput
}

type DomainNameMap map[string]DomainNameInput

func (DomainNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainName)(nil)).Elem()
}

func (i DomainNameMap) ToDomainNameMapOutput() DomainNameMapOutput {
	return i.ToDomainNameMapOutputWithContext(context.Background())
}

func (i DomainNameMap) ToDomainNameMapOutputWithContext(ctx context.Context) DomainNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainNameMapOutput)
}

type DomainNameOutput struct{ *pulumi.OutputState }

func (DomainNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainName)(nil))
}

func (o DomainNameOutput) ToDomainNameOutput() DomainNameOutput {
	return o
}

func (o DomainNameOutput) ToDomainNameOutputWithContext(ctx context.Context) DomainNameOutput {
	return o
}

func (o DomainNameOutput) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return o.ToDomainNamePtrOutputWithContext(context.Background())
}

func (o DomainNameOutput) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainName) *DomainName {
		return &v
	}).(DomainNamePtrOutput)
}

type DomainNamePtrOutput struct{ *pulumi.OutputState }

func (DomainNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainName)(nil))
}

func (o DomainNamePtrOutput) ToDomainNamePtrOutput() DomainNamePtrOutput {
	return o
}

func (o DomainNamePtrOutput) ToDomainNamePtrOutputWithContext(ctx context.Context) DomainNamePtrOutput {
	return o
}

func (o DomainNamePtrOutput) Elem() DomainNameOutput {
	return o.ApplyT(func(v *DomainName) DomainName {
		if v != nil {
			return *v
		}
		var ret DomainName
		return ret
	}).(DomainNameOutput)
}

type DomainNameArrayOutput struct{ *pulumi.OutputState }

func (DomainNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainName)(nil))
}

func (o DomainNameArrayOutput) ToDomainNameArrayOutput() DomainNameArrayOutput {
	return o
}

func (o DomainNameArrayOutput) ToDomainNameArrayOutputWithContext(ctx context.Context) DomainNameArrayOutput {
	return o
}

func (o DomainNameArrayOutput) Index(i pulumi.IntInput) DomainNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainName {
		return vs[0].([]DomainName)[vs[1].(int)]
	}).(DomainNameOutput)
}

type DomainNameMapOutput struct{ *pulumi.OutputState }

func (DomainNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DomainName)(nil))
}

func (o DomainNameMapOutput) ToDomainNameMapOutput() DomainNameMapOutput {
	return o
}

func (o DomainNameMapOutput) ToDomainNameMapOutputWithContext(ctx context.Context) DomainNameMapOutput {
	return o
}

func (o DomainNameMapOutput) MapIndex(k pulumi.StringInput) DomainNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DomainName {
		return vs[0].(map[string]DomainName)[vs[1].(string)]
	}).(DomainNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNameInput)(nil)).Elem(), &DomainName{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNamePtrInput)(nil)).Elem(), &DomainName{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNameArrayInput)(nil)).Elem(), DomainNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainNameMapInput)(nil)).Elem(), DomainNameMap{})
	pulumi.RegisterOutputType(DomainNameOutput{})
	pulumi.RegisterOutputType(DomainNamePtrOutput{})
	pulumi.RegisterOutputType(DomainNameArrayOutput{})
	pulumi.RegisterOutputType(DomainNameMapOutput{})
}
