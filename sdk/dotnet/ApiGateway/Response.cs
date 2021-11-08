// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides an API Gateway Gateway Response for a REST API Gateway.
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
    ///         var main = new Aws.ApiGateway.RestApi("main", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///         });
    ///         var test = new Aws.ApiGateway.Response("test", new Aws.ApiGateway.ResponseArgs
    ///         {
    ///             RestApiId = main.Id,
    ///             StatusCode = "401",
    ///             ResponseType = "UNAUTHORIZED",
    ///             ResponseTemplates = 
    ///             {
    ///                 { "application/json", "{\"message\":$context.error.messageString}" },
    ///             },
    ///             ResponseParameters = 
    ///             {
    ///                 { "gatewayresponse.header.Authorization", "'Basic'" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_api_gateway_gateway_response` can be imported using `REST-API-ID/RESPONSE-TYPE`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigateway/response:Response example 12345abcde/UNAUTHORIZED
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigateway/response:Response")]
    public partial class Response : Pulumi.CustomResource
    {
        /// <summary>
        /// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
        /// </summary>
        [Output("responseParameters")]
        public Output<ImmutableDictionary<string, string>?> ResponseParameters { get; private set; } = null!;

        /// <summary>
        /// A map specifying the templates used to transform the response body.
        /// </summary>
        [Output("responseTemplates")]
        public Output<ImmutableDictionary<string, string>?> ResponseTemplates { get; private set; } = null!;

        /// <summary>
        /// The response type of the associated GatewayResponse.
        /// </summary>
        [Output("responseType")]
        public Output<string> ResponseType { get; private set; } = null!;

        /// <summary>
        /// The string identifier of the associated REST API.
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;

        /// <summary>
        /// The HTTP status code of the Gateway Response.
        /// </summary>
        [Output("statusCode")]
        public Output<string?> StatusCode { get; private set; } = null!;


        /// <summary>
        /// Create a Response resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Response(string name, ResponseArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/response:Response", name, args ?? new ResponseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Response(string name, Input<string> id, ResponseState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/response:Response", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Response resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Response Get(string name, Input<string> id, ResponseState? state = null, CustomResourceOptions? options = null)
        {
            return new Response(name, id, state, options);
        }
    }

    public sealed class ResponseArgs : Pulumi.ResourceArgs
    {
        [Input("responseParameters")]
        private InputMap<string>? _responseParameters;

        /// <summary>
        /// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
        /// </summary>
        public InputMap<string> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<string>());
            set => _responseParameters = value;
        }

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map specifying the templates used to transform the response body.
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The response type of the associated GatewayResponse.
        /// </summary>
        [Input("responseType", required: true)]
        public Input<string> ResponseType { get; set; } = null!;

        /// <summary>
        /// The string identifier of the associated REST API.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        /// <summary>
        /// The HTTP status code of the Gateway Response.
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        public ResponseArgs()
        {
        }
    }

    public sealed class ResponseState : Pulumi.ResourceArgs
    {
        [Input("responseParameters")]
        private InputMap<string>? _responseParameters;

        /// <summary>
        /// A map specifying the parameters (paths, query strings and headers) of the Gateway Response.
        /// </summary>
        public InputMap<string> ResponseParameters
        {
            get => _responseParameters ?? (_responseParameters = new InputMap<string>());
            set => _responseParameters = value;
        }

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map specifying the templates used to transform the response body.
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The response type of the associated GatewayResponse.
        /// </summary>
        [Input("responseType")]
        public Input<string>? ResponseType { get; set; }

        /// <summary>
        /// The string identifier of the associated REST API.
        /// </summary>
        [Input("restApiId")]
        public Input<string>? RestApiId { get; set; }

        /// <summary>
        /// The HTTP status code of the Gateway Response.
        /// </summary>
        [Input("statusCode")]
        public Input<string>? StatusCode { get; set; }

        public ResponseState()
        {
        }
    }
}
