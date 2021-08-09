// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs.Inputs
{

    public sealed class SnapshotImportClientDataGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-defined comment about the disk upload.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The time that the disk upload ends.
        /// </summary>
        [Input("uploadEnd")]
        public Input<string>? UploadEnd { get; set; }

        /// <summary>
        /// The size of the uploaded disk image, in GiB.
        /// </summary>
        [Input("uploadSize")]
        public Input<double>? UploadSize { get; set; }

        /// <summary>
        /// The time that the disk upload starts.
        /// </summary>
        [Input("uploadStart")]
        public Input<string>? UploadStart { get; set; }

        public SnapshotImportClientDataGetArgs()
        {
        }
    }
}
