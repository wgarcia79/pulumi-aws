// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Aws.Ec2
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
        /// e.g., to allow easier migration from another management solution
        /// or to make it easier for an operator to connect through bastion host(s).
        /// 
        /// &gt; **Note:** It's strongly discouraged to use this data source for querying ephemeral
        /// instances (e.g., managed via autoscaling group), as the output may change at any time
        /// and you'd need to re-run `apply` every time an instance comes up or dies.
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("aws:ec2/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
        /// e.g., to allow easier migration from another management solution
        /// or to make it easier for an operator to connect through bastion host(s).
        /// 
        /// &gt; **Note:** It's strongly discouraged to use this data source for querying ephemeral
        /// instances (e.g., managed via autoscaling group), as the output may change at any time
        /// and you'd need to re-run `apply` every time an instance comes up or dies.
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("aws:ec2/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetInstancesFilterArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [describe-instances in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetInstancesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstancesFilterArgs>());
            set => _filters = value;
        }

        [Input("instanceStateNames")]
        private List<string>? _instanceStateNames;

        /// <summary>
        /// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
        /// </summary>
        public List<string> InstanceStateNames
        {
            get => _instanceStateNames ?? (_instanceStateNames = new List<string>());
            set => _instanceStateNames = value;
        }

        [Input("instanceTags")]
        private Dictionary<string, string>? _instanceTags;

        /// <summary>
        /// A map of tags, each pair of which must
        /// exactly match a pair on desired instances.
        /// </summary>
        public Dictionary<string, string> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new Dictionary<string, string>());
            set => _instanceTags = value;
        }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetInstancesFilterInputArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to use as filters. There are
        /// several valid keys, for a full reference, check out
        /// [describe-instances in the AWS CLI reference][1].
        /// </summary>
        public InputList<Inputs.GetInstancesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetInstancesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("instanceStateNames")]
        private InputList<string>? _instanceStateNames;

        /// <summary>
        /// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
        /// </summary>
        public InputList<string> InstanceStateNames
        {
            get => _instanceStateNames ?? (_instanceStateNames = new InputList<string>());
            set => _instanceStateNames = value;
        }

        [Input("instanceTags")]
        private InputMap<string>? _instanceTags;

        /// <summary>
        /// A map of tags, each pair of which must
        /// exactly match a pair on desired instances.
        /// </summary>
        public InputMap<string> InstanceTags
        {
            get => _instanceTags ?? (_instanceTags = new InputMap<string>());
            set => _instanceTags = value;
        }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        public readonly ImmutableArray<Outputs.GetInstancesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IDs of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> InstanceStateNames;
        public readonly ImmutableDictionary<string, string> InstanceTags;
        /// <summary>
        /// Private IP addresses of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> PrivateIps;
        /// <summary>
        /// Public IP addresses of instances found through the filter
        /// </summary>
        public readonly ImmutableArray<string> PublicIps;

        [OutputConstructor]
        private GetInstancesResult(
            ImmutableArray<Outputs.GetInstancesFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> instanceStateNames,

            ImmutableDictionary<string, string> instanceTags,

            ImmutableArray<string> privateIps,

            ImmutableArray<string> publicIps)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            InstanceStateNames = instanceStateNames;
            InstanceTags = instanceTags;
            PrivateIps = privateIps;
            PublicIps = publicIps;
        }
    }
}
