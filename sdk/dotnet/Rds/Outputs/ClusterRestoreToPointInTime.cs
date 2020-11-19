// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Outputs
{

    [OutputType]
    public sealed class ClusterRestoreToPointInTime
    {
        /// <summary>
        /// Date and time in UTC format to restore the database cluster to. Conflicts with `use_latest_restorable_time`.
        /// </summary>
        public readonly string? RestoreToTime;
        /// <summary>
        /// Type of restore to be performed.
        /// Valid options are `full-copy` (default) and `copy-on-write`.
        /// </summary>
        public readonly string? RestoreType;
        /// <summary>
        /// The identifier of the source database cluster from which to restore.
        /// </summary>
        public readonly string SourceClusterIdentifier;
        /// <summary>
        /// Set to true to restore the database cluster to the latest restorable backup time. Defaults to false. Conflicts with `restore_to_time`.
        /// </summary>
        public readonly bool? UseLatestRestorableTime;

        [OutputConstructor]
        private ClusterRestoreToPointInTime(
            string? restoreToTime,

            string? restoreType,

            string sourceClusterIdentifier,

            bool? useLatestRestorableTime)
        {
            RestoreToTime = restoreToTime;
            RestoreType = restoreType;
            SourceClusterIdentifier = sourceClusterIdentifier;
            UseLatestRestorableTime = useLatestRestorableTime;
        }
    }
}