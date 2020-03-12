// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFormation
{
    /// <summary>
    /// Manages a CloudFormation StackSet Instance. Instances are managed in the account and region of the StackSet after the target account permissions have been configured. Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).
    /// 
    /// &gt; **NOTE:** All target accounts must have an IAM Role created that matches the name of the execution role configured in the StackSet (the `execution_role_name` argument in the `aws.cloudformation.StackSet` resource) in a trust relationship with the administrative account or administration IAM Role. The execution role must have appropriate permissions to manage resources defined in the template along with those required for StackSets to operate. See the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) for more details.
    /// 
    /// &gt; **NOTE:** To retain the Stack during resource destroy, ensure `retain_stack` has been set to `true` in the state first. This must be completed _before_ a deployment that would destroy the resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudformation_stack_set_instance.html.markdown.
    /// </summary>
    public partial class StackSetInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Key-value map of input parameters to override from the StackSet for this Instance.
        /// </summary>
        [Output("parameterOverrides")]
        public Output<ImmutableDictionary<string, string>?> ParameterOverrides { get; private set; } = null!;

        /// <summary>
        /// Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
        /// </summary>
        [Output("retainStack")]
        public Output<bool?> RetainStack { get; private set; } = null!;

        /// <summary>
        /// Stack identifier
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// Name of the StackSet.
        /// </summary>
        [Output("stackSetName")]
        public Output<string> StackSetName { get; private set; } = null!;


        /// <summary>
        /// Create a StackSetInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackSetInstance(string name, StackSetInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudformation/stackSetInstance:StackSetInstance", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private StackSetInstance(string name, Input<string> id, StackSetInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudformation/stackSetInstance:StackSetInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StackSetInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackSetInstance Get(string name, Input<string> id, StackSetInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new StackSetInstance(name, id, state, options);
        }
    }

    public sealed class StackSetInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("parameterOverrides")]
        private InputMap<string>? _parameterOverrides;

        /// <summary>
        /// Key-value map of input parameters to override from the StackSet for this Instance.
        /// </summary>
        public InputMap<string> ParameterOverrides
        {
            get => _parameterOverrides ?? (_parameterOverrides = new InputMap<string>());
            set => _parameterOverrides = value;
        }

        /// <summary>
        /// Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
        /// </summary>
        [Input("retainStack")]
        public Input<bool>? RetainStack { get; set; }

        /// <summary>
        /// Name of the StackSet.
        /// </summary>
        [Input("stackSetName", required: true)]
        public Input<string> StackSetName { get; set; } = null!;

        public StackSetInstanceArgs()
        {
        }
    }

    public sealed class StackSetInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("parameterOverrides")]
        private InputMap<string>? _parameterOverrides;

        /// <summary>
        /// Key-value map of input parameters to override from the StackSet for this Instance.
        /// </summary>
        public InputMap<string> ParameterOverrides
        {
            get => _parameterOverrides ?? (_parameterOverrides = new InputMap<string>());
            set => _parameterOverrides = value;
        }

        /// <summary>
        /// Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// During resource destroy, remove Instance from StackSet while keeping the Stack and its associated resources. Must be enabled in the state _before_ destroy operation to take effect. You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to `false`.
        /// </summary>
        [Input("retainStack")]
        public Input<bool>? RetainStack { get; set; }

        /// <summary>
        /// Stack identifier
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// Name of the StackSet.
        /// </summary>
        [Input("stackSetName")]
        public Input<string>? StackSetName { get; set; }

        public StackSetInstanceState()
        {
        }
    }
}
