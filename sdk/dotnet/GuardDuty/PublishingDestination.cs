// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Provides a resource to manage a GuardDuty PublishingDestination. Requires an existing GuardDuty Detector.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var currentCallerIdentity = Output.Create(Aws.GetCallerIdentity.InvokeAsync());
    ///         var currentRegion = Output.Create(Aws.GetRegion.InvokeAsync());
    ///         var gdBucket = new Aws.S3.Bucket("gdBucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///             ForceDestroy = true,
    ///         });
    ///         var bucketPol = Output.Tuple(gdBucket.Arn, gdBucket.Arn).Apply(values =&gt;
    ///         {
    ///             var gdBucketArn = values.Item1;
    ///             var gdBucketArn1 = values.Item2;
    ///             return Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///             {
    ///                 Statements = 
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                     {
    ///                         Sid = "Allow PutObject",
    ///                         Actions = 
    ///                         {
    ///                             "s3:PutObject",
    ///                         },
    ///                         Resources = 
    ///                         {
    ///                             $"{gdBucketArn}/*",
    ///                         },
    ///                         Principals = 
    ///                         {
    ///                             new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                             {
    ///                                 Type = "Service",
    ///                                 Identifiers = 
    ///                                 {
    ///                                     "guardduty.amazonaws.com",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                     {
    ///                         Sid = "Allow GetBucketLocation",
    ///                         Actions = 
    ///                         {
    ///                             "s3:GetBucketLocation",
    ///                         },
    ///                         Resources = 
    ///                         {
    ///                             gdBucketArn1,
    ///                         },
    ///                         Principals = 
    ///                         {
    ///                             new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                             {
    ///                                 Type = "Service",
    ///                                 Identifiers = 
    ///                                 {
    ///                                     "guardduty.amazonaws.com",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             });
    ///         });
    ///         var kmsPol = Output.Tuple(currentRegion, currentCallerIdentity, currentRegion, currentCallerIdentity, currentCallerIdentity).Apply(values =&gt;
    ///         {
    ///             var currentRegion = values.Item1;
    ///             var currentCallerIdentity = values.Item2;
    ///             var currentRegion1 = values.Item3;
    ///             var currentCallerIdentity1 = values.Item4;
    ///             var currentCallerIdentity2 = values.Item5;
    ///             return Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///             {
    ///                 Statements = 
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                     {
    ///                         Sid = "Allow GuardDuty to encrypt findings",
    ///                         Actions = 
    ///                         {
    ///                             "kms:GenerateDataKey",
    ///                         },
    ///                         Resources = 
    ///                         {
    ///                             $"arn:aws:kms:{currentRegion.Name}:{currentCallerIdentity.AccountId}:key/*",
    ///                         },
    ///                         Principals = 
    ///                         {
    ///                             new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                             {
    ///                                 Type = "Service",
    ///                                 Identifiers = 
    ///                                 {
    ///                                     "guardduty.amazonaws.com",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                     {
    ///                         Sid = "Allow all users to modify/delete key (test only)",
    ///                         Actions = 
    ///                         {
    ///                             "kms:*",
    ///                         },
    ///                         Resources = 
    ///                         {
    ///                             $"arn:aws:kms:{currentRegion1.Name}:{currentCallerIdentity1.AccountId}:key/*",
    ///                         },
    ///                         Principals = 
    ///                         {
    ///                             new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                             {
    ///                                 Type = "AWS",
    ///                                 Identifiers = 
    ///                                 {
    ///                                     $"arn:aws:iam::{currentCallerIdentity2.AccountId}:root",
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             }));
    ///         });
    ///         var testGd = new Aws.GuardDuty.Detector("testGd", new Aws.GuardDuty.DetectorArgs
    ///         {
    ///             Enable = true,
    ///         });
    ///         var gdBucketPolicy = new Aws.S3.BucketPolicy("gdBucketPolicy", new Aws.S3.BucketPolicyArgs
    ///         {
    ///             Bucket = gdBucket.Id,
    ///             Policy = bucketPol.Apply(bucketPol =&gt; bucketPol.Json),
    ///         });
    ///         var gdKey = new Aws.Kms.Key("gdKey", new Aws.Kms.KeyArgs
    ///         {
    ///             Description = "Temporary key for AccTest of TF",
    ///             DeletionWindowInDays = 7,
    ///             Policy = kmsPol.Apply(kmsPol =&gt; kmsPol.Json),
    ///         });
    ///         var test = new Aws.GuardDuty.PublishingDestination("test", new Aws.GuardDuty.PublishingDestinationArgs
    ///         {
    ///             DetectorId = testGd.Id,
    ///             DestinationArn = gdBucket.Arn,
    ///             KmsKeyArn = gdKey.Arn,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 gdBucketPolicy,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &gt; **Note:** Please do not use this simple example for Bucket-Policy and KMS Key Policy in a production environment. It is much too open for such a use-case. Refer to the AWS documentation here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_exportfindings.html
    /// 
    /// ## Import
    /// 
    /// GuardDuty PublishingDestination can be imported using the the master GuardDuty detector ID and PublishingDestinationID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:guardduty/publishingDestination:PublishingDestination test a4b86f26fa42e7e7cf0d1c333ea77777:a4b86f27a0e464e4a7e0516d242f1234
    /// ```
    /// </summary>
    [AwsResourceType("aws:guardduty/publishingDestination:PublishingDestination")]
    public partial class PublishingDestination : Pulumi.CustomResource
    {
        /// <summary>
        /// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        /// </summary>
        [Output("destinationArn")]
        public Output<string> DestinationArn { get; private set; } = null!;

        /// <summary>
        /// Currently there is only "S3" available as destination type which is also the default value
        /// </summary>
        [Output("destinationType")]
        public Output<string?> DestinationType { get; private set; } = null!;

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string> KmsKeyArn { get; private set; } = null!;


        /// <summary>
        /// Create a PublishingDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublishingDestination(string name, PublishingDestinationArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/publishingDestination:PublishingDestination", name, args ?? new PublishingDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublishingDestination(string name, Input<string> id, PublishingDestinationState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/publishingDestination:PublishingDestination", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PublishingDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublishingDestination Get(string name, Input<string> id, PublishingDestinationState? state = null, CustomResourceOptions? options = null)
        {
            return new PublishingDestination(name, id, state, options);
        }
    }

    public sealed class PublishingDestinationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        /// </summary>
        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        /// <summary>
        /// Currently there is only "S3" available as destination type which is also the default value
        /// </summary>
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        /// <summary>
        /// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
        /// </summary>
        [Input("kmsKeyArn", required: true)]
        public Input<string> KmsKeyArn { get; set; } = null!;

        public PublishingDestinationArgs()
        {
        }
    }

    public sealed class PublishingDestinationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        /// </summary>
        [Input("destinationArn")]
        public Input<string>? DestinationArn { get; set; }

        /// <summary>
        /// Currently there is only "S3" available as destination type which is also the default value
        /// </summary>
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        /// <summary>
        /// The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        public PublishingDestinationState()
        {
        }
    }
}
