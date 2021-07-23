// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue.Inputs
{

    public sealed class CrawlerS3TargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        [Input("exclusions")]
        private InputList<string>? _exclusions;

        /// <summary>
        /// A list of glob patterns used to exclude from the crawl.
        /// </summary>
        public InputList<string> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<string>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The path of the Amazon DocumentDB or MongoDB target (database/collection).
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.
        /// </summary>
        [Input("sampleSize")]
        public Input<int>? SampleSize { get; set; }

        public CrawlerS3TargetArgs()
        {
        }
    }
}
