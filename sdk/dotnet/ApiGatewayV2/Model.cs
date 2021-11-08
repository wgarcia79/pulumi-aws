// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    /// <summary>
    /// Manages an Amazon API Gateway Version 2 [model](https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html#models-mappings-models).
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ApiGatewayV2.Model("example", new Aws.ApiGatewayV2.ModelArgs
    ///         {
    ///             ApiId = aws_apigatewayv2_api.Example.Id,
    ///             ContentType = "application/json",
    ///             Schema = @"{
    ///   ""$schema"": ""http://json-schema.org/draft-04/schema#"",
    ///   ""title"": ""ExampleModel"",
    ///   ""type"": ""object"",
    ///   ""properties"": {
    ///     ""id"": { ""type"": ""string"" }
    ///   }
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_apigatewayv2_model` can be imported by using the API identifier and model identifier, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigatewayv2/model:Model example aabbccddee/1122334
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigatewayv2/model:Model")]
    public partial class Model : Pulumi.CustomResource
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The content-type for the model, for example, `application/json`. Must be between 1 and 256 characters in length.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// The description of the model. Must be between 1 and 128 characters in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the model. Must be alphanumeric. Must be between 1 and 128 characters in length.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The schema for the model. This should be a [JSON schema draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Must be less than or equal to 32768 characters in length.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;


        /// <summary>
        /// Create a Model resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Model(string name, ModelArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/model:Model", name, args ?? new ModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Model(string name, Input<string> id, ModelState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/model:Model", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Model resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Model Get(string name, Input<string> id, ModelState? state = null, CustomResourceOptions? options = null)
        {
            return new Model(name, id, state, options);
        }
    }

    public sealed class ModelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The content-type for the model, for example, `application/json`. Must be between 1 and 256 characters in length.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// The description of the model. Must be between 1 and 128 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the model. Must be alphanumeric. Must be between 1 and 128 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schema for the model. This should be a [JSON schema draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Must be less than or equal to 32768 characters in length.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        public ModelArgs()
        {
        }
    }

    public sealed class ModelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The content-type for the model, for example, `application/json`. Must be between 1 and 256 characters in length.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// The description of the model. Must be between 1 and 128 characters in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the model. Must be alphanumeric. Must be between 1 and 128 characters in length.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schema for the model. This should be a [JSON schema draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Must be less than or equal to 32768 characters in length.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public ModelState()
        {
        }
    }
}
