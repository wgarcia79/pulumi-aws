// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeStarConnections
{
    /// <summary>
    /// Provides a CodeStar Host.
    /// 
    /// &gt; **NOTE:** The `aws.codestarconnections.Host` resource is created in the state `PENDING`. Authentication with the host provider must be completed in the AWS Console. For more information visit [Set up a pending host](https://docs.aws.amazon.com/dtconsole/latest/userguide/connections-host-setup.html).
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
    ///         var example = new Aws.CodeStarConnections.Host("example", new Aws.CodeStarConnections.HostArgs
    ///         {
    ///             ProviderEndpoint = "https://example.com",
    ///             ProviderType = "GitHubEnterpriseServer",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeStar Host can be imported using the ARN, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:codestarconnections/host:Host example-host arn:aws:codestar-connections:us-west-1:0123456789:host/79d4d357-a2ee-41e4-b350-2fe39ae59448
    /// ```
    /// </summary>
    [AwsResourceType("aws:codestarconnections/host:Host")]
    public partial class Host : Pulumi.CustomResource
    {
        /// <summary>
        /// The CodeStar Host ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the host to be created. The name must be unique in the calling AWS account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the infrastructure to be represented by the host after it is created.
        /// </summary>
        [Output("providerEndpoint")]
        public Output<string> ProviderEndpoint { get; private set; } = null!;

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured.
        /// </summary>
        [Output("providerType")]
        public Output<string> ProviderType { get; private set; } = null!;

        /// <summary>
        /// The CodeStar Host status. Possible values are `PENDING`, `AVAILABLE`, `VPC_CONFIG_DELETING`, `VPC_CONFIG_INITIALIZING`, and `VPC_CONFIG_FAILED_INITIALIZATION`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
        /// </summary>
        [Output("vpcConfiguration")]
        public Output<Outputs.HostVpcConfiguration?> VpcConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a Host resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Host(string name, HostArgs args, CustomResourceOptions? options = null)
            : base("aws:codestarconnections/host:Host", name, args ?? new HostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Host(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
            : base("aws:codestarconnections/host:Host", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Host resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Host Get(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
        {
            return new Host(name, id, state, options);
        }
    }

    public sealed class HostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the host to be created. The name must be unique in the calling AWS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The endpoint of the infrastructure to be represented by the host after it is created.
        /// </summary>
        [Input("providerEndpoint", required: true)]
        public Input<string> ProviderEndpoint { get; set; } = null!;

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured.
        /// </summary>
        [Input("providerType", required: true)]
        public Input<string> ProviderType { get; set; } = null!;

        /// <summary>
        /// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
        /// </summary>
        [Input("vpcConfiguration")]
        public Input<Inputs.HostVpcConfigurationArgs>? VpcConfiguration { get; set; }

        public HostArgs()
        {
        }
    }

    public sealed class HostState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CodeStar Host ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the host to be created. The name must be unique in the calling AWS account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The endpoint of the infrastructure to be represented by the host after it is created.
        /// </summary>
        [Input("providerEndpoint")]
        public Input<string>? ProviderEndpoint { get; set; }

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured.
        /// </summary>
        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        /// <summary>
        /// The CodeStar Host status. Possible values are `PENDING`, `AVAILABLE`, `VPC_CONFIG_DELETING`, `VPC_CONFIG_INITIALIZING`, and `VPC_CONFIG_FAILED_INITIALIZATION`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The VPC configuration to be provisioned for the host. A VPC must be configured, and the infrastructure to be represented by the host must already be connected to the VPC.
        /// </summary>
        [Input("vpcConfiguration")]
        public Input<Inputs.HostVpcConfigurationGetArgs>? VpcConfiguration { get; set; }

        public HostState()
        {
        }
    }
}
