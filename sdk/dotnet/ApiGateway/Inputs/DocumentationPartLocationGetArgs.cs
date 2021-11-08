// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway.Inputs
{

    public sealed class DocumentationPartLocationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTTP verb of a method. The default value is `*` for any method.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// The name of the targeted API entity.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL path of the target. The default value is `/` for the root resource.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The HTTP status code of a response. The default value is `*` for any status code.
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        /// <summary>
        /// The type of API entity to which the documentation content appliesE.g., `API`, `METHOD` or `REQUEST_BODY`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DocumentationPartLocationGetArgs()
        {
        }
    }
}
