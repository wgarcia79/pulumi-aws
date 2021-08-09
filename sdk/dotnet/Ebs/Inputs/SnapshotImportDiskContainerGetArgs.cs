// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs.Inputs
{

    public sealed class SnapshotImportDiskContainerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the disk image being imported.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The format of the disk image being imported. One of `VHD` or `VMDK`.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// The URL to the Amazon S3-based disk image being imported. It can either be a https URL (https://..) or an Amazon S3 URL (s3://..). One of `url` or `user_bucket` must be set.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The Amazon S3 bucket for the disk image. One of `url` or `user_bucket` must be set. Detailed below.
        /// </summary>
        [Input("userBucket")]
        public Input<Inputs.SnapshotImportDiskContainerUserBucketGetArgs>? UserBucket { get; set; }

        public SnapshotImportDiskContainerGetArgs()
        {
        }
    }
}
