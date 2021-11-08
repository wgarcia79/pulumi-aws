// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Fms
{
    /// <summary>
    /// Provides a resource to associate/disassociate an AWS Firewall Manager administrator account. This operation must be performed in the `us-east-1` region.
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
    ///         var example = new Aws.Fms.AdminAccount("example", new Aws.Fms.AdminAccountArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Manager administrator account association can be imported using the account ID, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:fms/adminAccount:AdminAccount example 123456789012
    /// ```
    /// </summary>
    [AwsResourceType("aws:fms/adminAccount:AdminAccount")]
    public partial class AdminAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;


        /// <summary>
        /// Create a AdminAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdminAccount(string name, AdminAccountArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:fms/adminAccount:AdminAccount", name, args ?? new AdminAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdminAccount(string name, Input<string> id, AdminAccountState? state = null, CustomResourceOptions? options = null)
            : base("aws:fms/adminAccount:AdminAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AdminAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdminAccount Get(string name, Input<string> id, AdminAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new AdminAccount(name, id, state, options);
        }
    }

    public sealed class AdminAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        public AdminAccountArgs()
        {
        }
    }

    public sealed class AdminAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        public AdminAccountState()
        {
        }
    }
}
