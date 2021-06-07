// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceCatalog
{
    /// <summary>
    /// Manages Service Catalog AWS Organizations Access, a portfolio sharing feature through AWS Organizations. This allows Service Catalog to receive updates on your organization in order to sync your shares with the current structure. This resource will prompt AWS to set `organizations:EnableAWSServiceAccess` on your behalf so that your shares can be in sync with any changes in your AWS Organizations structure.
    /// 
    /// &gt; **NOTE:** This resource can only be used by the management account in the organization. In other words, a delegated administrator is not authorized to use the resource.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ServiceCatalog.OrganizationsAccess("example", new Aws.ServiceCatalog.OrganizationsAccessArgs
    ///         {
    ///             Enabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:servicecatalog/organizationsAccess:OrganizationsAccess")]
    public partial class OrganizationsAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable AWS Organizations access.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationsAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationsAccess(string name, OrganizationsAccessArgs args, CustomResourceOptions? options = null)
            : base("aws:servicecatalog/organizationsAccess:OrganizationsAccess", name, args ?? new OrganizationsAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationsAccess(string name, Input<string> id, OrganizationsAccessState? state = null, CustomResourceOptions? options = null)
            : base("aws:servicecatalog/organizationsAccess:OrganizationsAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationsAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationsAccess Get(string name, Input<string> id, OrganizationsAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationsAccess(name, id, state, options);
        }
    }

    public sealed class OrganizationsAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable AWS Organizations access.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public OrganizationsAccessArgs()
        {
        }
    }

    public sealed class OrganizationsAccessState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable AWS Organizations access.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public OrganizationsAccessState()
        {
        }
    }
}