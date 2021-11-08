// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Amp
{
    /// <summary>
    /// Manages an Amazon Managed Service for Prometheus (AMP) Alert Manager Definition
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
    ///         var demoWorkspace = new Aws.Amp.Workspace("demoWorkspace", new Aws.Amp.WorkspaceArgs
    ///         {
    ///         });
    ///         var demoAlertManagerDefinition = new Aws.Amp.AlertManagerDefinition("demoAlertManagerDefinition", new Aws.Amp.AlertManagerDefinitionArgs
    ///         {
    ///             WorkspaceId = demoWorkspace.Id,
    ///             Definition = @"alertmanager_config: |
    ///   route:
    ///     receiver: 'default'
    ///   receivers:
    ///     - name: 'default'
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// The prometheus alert manager definition can be imported using the workspace identifier, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:amp/alertManagerDefinition:AlertManagerDefinition demo ws-C6DCB907-F2D7-4D96-957B-66691F865D8B
    /// ```
    /// </summary>
    [AwsResourceType("aws:amp/alertManagerDefinition:AlertManagerDefinition")]
    public partial class AlertManagerDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
        /// </summary>
        [Output("definition")]
        public Output<string> Definition { get; private set; } = null!;

        /// <summary>
        /// The id of the prometheus workspace the alert manager definition should be linked to
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a AlertManagerDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertManagerDefinition(string name, AlertManagerDefinitionArgs args, CustomResourceOptions? options = null)
            : base("aws:amp/alertManagerDefinition:AlertManagerDefinition", name, args ?? new AlertManagerDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertManagerDefinition(string name, Input<string> id, AlertManagerDefinitionState? state = null, CustomResourceOptions? options = null)
            : base("aws:amp/alertManagerDefinition:AlertManagerDefinition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertManagerDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertManagerDefinition Get(string name, Input<string> id, AlertManagerDefinitionState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertManagerDefinition(name, id, state, options);
        }
    }

    public sealed class AlertManagerDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// The id of the prometheus workspace the alert manager definition should be linked to
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public AlertManagerDefinitionArgs()
        {
        }
    }

    public sealed class AlertManagerDefinitionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
        /// </summary>
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// The id of the prometheus workspace the alert manager definition should be linked to
        /// </summary>
        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public AlertManagerDefinitionState()
        {
        }
    }
}
