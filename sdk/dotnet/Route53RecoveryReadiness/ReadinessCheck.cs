// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53RecoveryReadiness
{
    /// <summary>
    /// Provides an AWS Route 53 Recovery Readiness Readiness Check.
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
    ///         var example = new Aws.Route53RecoveryReadiness.ReadinessCheck("example", new Aws.Route53RecoveryReadiness.ReadinessCheckArgs
    ///         {
    ///             ReadinessCheckName = my_cw_alarm_check,
    ///             ResourceSetName = my_cw_alarm_set,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Route53 Recovery Readiness readiness checks can be imported via the readiness check name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53recoveryreadiness/readinessCheck:ReadinessCheck my-cw-alarm-check
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53recoveryreadiness/readinessCheck:ReadinessCheck")]
    public partial class ReadinessCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the readiness_check
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Unique name describing the readiness check.
        /// </summary>
        [Output("readinessCheckName")]
        public Output<string> ReadinessCheckName { get; private set; } = null!;

        /// <summary>
        /// Name describing the resource set that will be monitored for readiness.
        /// </summary>
        [Output("resourceSetName")]
        public Output<string> ResourceSetName { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a ReadinessCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReadinessCheck(string name, ReadinessCheckArgs args, CustomResourceOptions? options = null)
            : base("aws:route53recoveryreadiness/readinessCheck:ReadinessCheck", name, args ?? new ReadinessCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReadinessCheck(string name, Input<string> id, ReadinessCheckState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53recoveryreadiness/readinessCheck:ReadinessCheck", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReadinessCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReadinessCheck Get(string name, Input<string> id, ReadinessCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new ReadinessCheck(name, id, state, options);
        }
    }

    public sealed class ReadinessCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name describing the readiness check.
        /// </summary>
        [Input("readinessCheckName", required: true)]
        public Input<string> ReadinessCheckName { get; set; } = null!;

        /// <summary>
        /// Name describing the resource set that will be monitored for readiness.
        /// </summary>
        [Input("resourceSetName", required: true)]
        public Input<string> ResourceSetName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ReadinessCheckArgs()
        {
        }
    }

    public sealed class ReadinessCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the readiness_check
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Unique name describing the readiness check.
        /// </summary>
        [Input("readinessCheckName")]
        public Input<string>? ReadinessCheckName { get; set; }

        /// <summary>
        /// Name describing the resource set that will be monitored for readiness.
        /// </summary>
        [Input("resourceSetName")]
        public Input<string>? ResourceSetName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ReadinessCheckState()
        {
        }
    }
}
