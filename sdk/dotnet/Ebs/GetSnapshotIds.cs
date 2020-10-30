// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs
{
    public static class GetSnapshotIds
    {
        /// <summary>
        /// Use this data source to get a list of EBS Snapshot IDs matching the specified
        /// criteria.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ebsVolumes = Output.Create(Aws.Ebs.GetSnapshotIds.InvokeAsync(new Aws.Ebs.GetSnapshotIdsArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Ebs.Inputs.GetSnapshotIdsFilterArgs
        ///                 {
        ///                     Name = "volume-size",
        ///                     Values = 
        ///                     {
        ///                         "40",
        ///                     },
        ///                 },
        ///                 new Aws.Ebs.Inputs.GetSnapshotIdsFilterArgs
        ///                 {
        ///                     Name = "tag:Name",
        ///                     Values = 
        ///                     {
        ///                         "Example",
        ///                     },
        ///                 },
        ///             },
        ///             Owners = 
        ///             {
        ///                 "self",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSnapshotIdsResult> InvokeAsync(GetSnapshotIdsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotIdsResult>("aws:ebs/getSnapshotIds:getSnapshotIds", args ?? new GetSnapshotIdsArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotIdsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSnapshotIdsFilterArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to filter off of. There are
        /// several valid keys, for a full reference, check out
        /// [describe-volumes in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetSnapshotIdsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSnapshotIdsFilterArgs>());
            set => _filters = value;
        }

        [Input("owners")]
        private List<string>? _owners;

        /// <summary>
        /// Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
        /// </summary>
        public List<string> Owners
        {
            get => _owners ?? (_owners = new List<string>());
            set => _owners = value;
        }

        [Input("restorableByUserIds")]
        private List<string>? _restorableByUserIds;

        /// <summary>
        /// One or more AWS accounts IDs that can create volumes from the snapshot.
        /// </summary>
        public List<string> RestorableByUserIds
        {
            get => _restorableByUserIds ?? (_restorableByUserIds = new List<string>());
            set => _restorableByUserIds = value;
        }

        public GetSnapshotIdsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotIdsResult
    {
        public readonly ImmutableArray<Outputs.GetSnapshotIdsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of EBS snapshot IDs, sorted by creation time in descending order.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> Owners;
        public readonly ImmutableArray<string> RestorableByUserIds;

        [OutputConstructor]
        private GetSnapshotIdsResult(
            ImmutableArray<Outputs.GetSnapshotIdsFilterResult> filters,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> owners,

            ImmutableArray<string> restorableByUserIds)
        {
            Filters = filters;
            Id = id;
            Ids = ids;
            Owners = owners;
            RestorableByUserIds = restorableByUserIds;
        }
    }
}
