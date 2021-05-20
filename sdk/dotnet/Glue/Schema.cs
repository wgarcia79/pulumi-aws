// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Provides a Glue Schema resource.
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
    ///         var example = new Aws.Glue.Schema("example", new Aws.Glue.SchemaArgs
    ///         {
    ///             SchemaName = "example",
    ///             RegistryArn = aws_glue_registry.Test.Arn,
    ///             DataFormat = "AVRO",
    ///             Compatibility = "NONE",
    ///             SchemaDefinition = "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Glue Registries can be imported using `arn`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:glue/schema:Schema example arn:aws:glue:us-west-2:123456789012:schema/example/example
    /// ```
    /// </summary>
    [AwsResourceType("aws:glue/schema:Schema")]
    public partial class Schema : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the schema.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        /// </summary>
        [Output("compatibility")]
        public Output<string> Compatibility { get; private set; } = null!;

        /// <summary>
        /// The data format of the schema definition. Currently only `AVRO` is supported.
        /// </summary>
        [Output("dataFormat")]
        public Output<string> DataFormat { get; private set; } = null!;

        /// <summary>
        /// A description of the schema.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The latest version of the schema associated with the returned schema definition.
        /// </summary>
        [Output("latestSchemaVersion")]
        public Output<int> LatestSchemaVersion { get; private set; } = null!;

        /// <summary>
        /// The next version of the schema associated with the returned schema definition.
        /// </summary>
        [Output("nextSchemaVersion")]
        public Output<int> NextSchemaVersion { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Glue Registry to create the schema in.
        /// </summary>
        [Output("registryArn")]
        public Output<string> RegistryArn { get; private set; } = null!;

        /// <summary>
        /// The name of the Glue Registry.
        /// </summary>
        [Output("registryName")]
        public Output<string> RegistryName { get; private set; } = null!;

        /// <summary>
        /// The version number of the checkpoint (the last time the compatibility mode was changed).
        /// </summary>
        [Output("schemaCheckpoint")]
        public Output<int> SchemaCheckpoint { get; private set; } = null!;

        /// <summary>
        /// The schema definition using the `data_format` setting for `schema_name`.
        /// </summary>
        [Output("schemaDefinition")]
        public Output<string> SchemaDefinition { get; private set; } = null!;

        /// <summary>
        /// The Name of the schema.
        /// </summary>
        [Output("schemaName")]
        public Output<string> SchemaName { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Schema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schema(string name, SchemaArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/schema:Schema", name, args ?? new SchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schema(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/schema:Schema", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Schema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schema Get(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new Schema(name, id, state, options);
        }
    }

    public sealed class SchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        /// </summary>
        [Input("compatibility", required: true)]
        public Input<string> Compatibility { get; set; } = null!;

        /// <summary>
        /// The data format of the schema definition. Currently only `AVRO` is supported.
        /// </summary>
        [Input("dataFormat", required: true)]
        public Input<string> DataFormat { get; set; } = null!;

        /// <summary>
        /// A description of the schema.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ARN of the Glue Registry to create the schema in.
        /// </summary>
        [Input("registryArn")]
        public Input<string>? RegistryArn { get; set; }

        /// <summary>
        /// The schema definition using the `data_format` setting for `schema_name`.
        /// </summary>
        [Input("schemaDefinition", required: true)]
        public Input<string> SchemaDefinition { get; set; } = null!;

        /// <summary>
        /// The Name of the schema.
        /// </summary>
        [Input("schemaName", required: true)]
        public Input<string> SchemaName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public SchemaArgs()
        {
        }
    }

    public sealed class SchemaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the schema.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
        /// </summary>
        [Input("compatibility")]
        public Input<string>? Compatibility { get; set; }

        /// <summary>
        /// The data format of the schema definition. Currently only `AVRO` is supported.
        /// </summary>
        [Input("dataFormat")]
        public Input<string>? DataFormat { get; set; }

        /// <summary>
        /// A description of the schema.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The latest version of the schema associated with the returned schema definition.
        /// </summary>
        [Input("latestSchemaVersion")]
        public Input<int>? LatestSchemaVersion { get; set; }

        /// <summary>
        /// The next version of the schema associated with the returned schema definition.
        /// </summary>
        [Input("nextSchemaVersion")]
        public Input<int>? NextSchemaVersion { get; set; }

        /// <summary>
        /// The ARN of the Glue Registry to create the schema in.
        /// </summary>
        [Input("registryArn")]
        public Input<string>? RegistryArn { get; set; }

        /// <summary>
        /// The name of the Glue Registry.
        /// </summary>
        [Input("registryName")]
        public Input<string>? RegistryName { get; set; }

        /// <summary>
        /// The version number of the checkpoint (the last time the compatibility mode was changed).
        /// </summary>
        [Input("schemaCheckpoint")]
        public Input<int>? SchemaCheckpoint { get; set; }

        /// <summary>
        /// The schema definition using the `data_format` setting for `schema_name`.
        /// </summary>
        [Input("schemaDefinition")]
        public Input<string>? SchemaDefinition { get; set; }

        /// <summary>
        /// The Name of the schema.
        /// </summary>
        [Input("schemaName")]
        public Input<string>? SchemaName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public SchemaState()
        {
        }
    }
}
