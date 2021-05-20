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
    /// Provides a Glue ML Transform resource.
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
    ///         var testCatalogDatabase = new Aws.Glue.CatalogDatabase("testCatalogDatabase", new Aws.Glue.CatalogDatabaseArgs
    ///         {
    ///             Name = "example",
    ///         });
    ///         var testCatalogTable = new Aws.Glue.CatalogTable("testCatalogTable", new Aws.Glue.CatalogTableArgs
    ///         {
    ///             Name = "example",
    ///             DatabaseName = testCatalogDatabase.Name,
    ///             Owner = "my_owner",
    ///             Retention = 1,
    ///             TableType = "VIRTUAL_VIEW",
    ///             ViewExpandedText = "view_expanded_text_1",
    ///             ViewOriginalText = "view_original_text_1",
    ///             StorageDescriptor = new Aws.Glue.Inputs.CatalogTableStorageDescriptorArgs
    ///             {
    ///                 BucketColumns = 
    ///                 {
    ///                     "bucket_column_1",
    ///                 },
    ///                 Compressed = false,
    ///                 InputFormat = "SequenceFileInputFormat",
    ///                 Location = "my_location",
    ///                 NumberOfBuckets = 1,
    ///                 OutputFormat = "SequenceFileInputFormat",
    ///                 StoredAsSubDirectories = false,
    ///                 Parameters = 
    ///                 {
    ///                     { "param1", "param1_val" },
    ///                 },
    ///                 Columns = 
    ///                 {
    ///                     new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
    ///                     {
    ///                         Name = "my_column_1",
    ///                         Type = "int",
    ///                         Comment = "my_column1_comment",
    ///                     },
    ///                     new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
    ///                     {
    ///                         Name = "my_column_2",
    ///                         Type = "string",
    ///                         Comment = "my_column2_comment",
    ///                     },
    ///                 },
    ///                 SerDeInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSerDeInfoArgs
    ///                 {
    ///                     Name = "ser_de_name",
    ///                     Parameters = 
    ///                     {
    ///                         { "param1", "param_val_1" },
    ///                     },
    ///                     SerializationLibrary = "org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
    ///                 },
    ///                 SortColumns = 
    ///                 {
    ///                     new Aws.Glue.Inputs.CatalogTableStorageDescriptorSortColumnArgs
    ///                     {
    ///                         Column = "my_column_1",
    ///                         SortOrder = 1,
    ///                     },
    ///                 },
    ///                 SkewedInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSkewedInfoArgs
    ///                 {
    ///                     SkewedColumnNames = 
    ///                     {
    ///                         "my_column_1",
    ///                     },
    ///                     SkewedColumnValueLocationMaps = 
    ///                     {
    ///                         { "my_column_1", "my_column_1_val_loc_map" },
    ///                     },
    ///                     SkewedColumnValues = 
    ///                     {
    ///                         "skewed_val_1",
    ///                     },
    ///                 },
    ///             },
    ///             PartitionKeys = 
    ///             {
    ///                 new Aws.Glue.Inputs.CatalogTablePartitionKeyArgs
    ///                 {
    ///                     Name = "my_column_1",
    ///                     Type = "int",
    ///                     Comment = "my_column_1_comment",
    ///                 },
    ///                 new Aws.Glue.Inputs.CatalogTablePartitionKeyArgs
    ///                 {
    ///                     Name = "my_column_2",
    ///                     Type = "string",
    ///                     Comment = "my_column_2_comment",
    ///                 },
    ///             },
    ///             Parameters = 
    ///             {
    ///                 { "param1", "param1_val" },
    ///             },
    ///         });
    ///         var testMLTransform = new Aws.Glue.MLTransform("testMLTransform", new Aws.Glue.MLTransformArgs
    ///         {
    ///             RoleArn = aws_iam_role.Test.Arn,
    ///             InputRecordTables = 
    ///             {
    ///                 new Aws.Glue.Inputs.MLTransformInputRecordTableArgs
    ///                 {
    ///                     DatabaseName = testCatalogTable.DatabaseName,
    ///                     TableName = testCatalogTable.Name,
    ///                 },
    ///             },
    ///             Parameters = new Aws.Glue.Inputs.MLTransformParametersArgs
    ///             {
    ///                 TransformType = "FIND_MATCHES",
    ///                 FindMatchesParameters = new Aws.Glue.Inputs.MLTransformParametersFindMatchesParametersArgs
    ///                 {
    ///                     PrimaryKeyColumnName = "my_column_1",
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 aws_iam_role_policy_attachment.Test,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Glue ML Transforms can be imported using `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:glue/mLTransform:MLTransform example tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
    /// ```
    /// </summary>
    [AwsResourceType("aws:glue/mLTransform:MLTransform")]
    public partial class MLTransform : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Glue ML Transform.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the ML Transform.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Output("glueVersion")]
        public Output<string> GlueVersion { get; private set; } = null!;

        /// <summary>
        /// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
        /// </summary>
        [Output("inputRecordTables")]
        public Output<ImmutableArray<Outputs.MLTransformInputRecordTable>> InputRecordTables { get; private set; } = null!;

        /// <summary>
        /// The number of labels available for this transform.
        /// </summary>
        [Output("labelCount")]
        public Output<int> LabelCount { get; private set; } = null!;

        /// <summary>
        /// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `max_capacity` is a mutually exclusive option with `number_of_workers` and `worker_type`.
        /// </summary>
        [Output("maxCapacity")]
        public Output<double> MaxCapacity { get; private set; } = null!;

        /// <summary>
        /// The maximum number of times to retry this ML Transform if it fails.
        /// </summary>
        [Output("maxRetries")]
        public Output<int?> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// The name you assign to this ML Transform. It must be unique in your account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of workers of a defined `worker_type` that are allocated when an ML Transform runs. Required with `worker_type`.
        /// </summary>
        [Output("numberOfWorkers")]
        public Output<int?> NumberOfWorkers { get; private set; } = null!;

        /// <summary>
        /// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
        /// </summary>
        [Output("parameters")]
        public Output<Outputs.MLTransformParameters> Parameters { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role associated with this ML Transform.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The object that represents the schema that this transform accepts. see Schema.
        /// </summary>
        [Output("schemas")]
        public Output<ImmutableArray<Outputs.MLTransformSchema>> Schemas { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `number_of_workers`.
        /// </summary>
        [Output("workerType")]
        public Output<string?> WorkerType { get; private set; } = null!;


        /// <summary>
        /// Create a MLTransform resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MLTransform(string name, MLTransformArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/mLTransform:MLTransform", name, args ?? new MLTransformArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MLTransform(string name, Input<string> id, MLTransformState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/mLTransform:MLTransform", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MLTransform resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MLTransform Get(string name, Input<string> id, MLTransformState? state = null, CustomResourceOptions? options = null)
        {
            return new MLTransform(name, id, state, options);
        }
    }

    public sealed class MLTransformArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the ML Transform.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Input("glueVersion")]
        public Input<string>? GlueVersion { get; set; }

        [Input("inputRecordTables", required: true)]
        private InputList<Inputs.MLTransformInputRecordTableArgs>? _inputRecordTables;

        /// <summary>
        /// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
        /// </summary>
        public InputList<Inputs.MLTransformInputRecordTableArgs> InputRecordTables
        {
            get => _inputRecordTables ?? (_inputRecordTables = new InputList<Inputs.MLTransformInputRecordTableArgs>());
            set => _inputRecordTables = value;
        }

        /// <summary>
        /// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `max_capacity` is a mutually exclusive option with `number_of_workers` and `worker_type`.
        /// </summary>
        [Input("maxCapacity")]
        public Input<double>? MaxCapacity { get; set; }

        /// <summary>
        /// The maximum number of times to retry this ML Transform if it fails.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name you assign to this ML Transform. It must be unique in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of workers of a defined `worker_type` that are allocated when an ML Transform runs. Required with `worker_type`.
        /// </summary>
        [Input("numberOfWorkers")]
        public Input<int>? NumberOfWorkers { get; set; }

        /// <summary>
        /// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<Inputs.MLTransformParametersArgs> Parameters { get; set; } = null!;

        /// <summary>
        /// The ARN of the IAM role associated with this ML Transform.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

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

        /// <summary>
        /// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `number_of_workers`.
        /// </summary>
        [Input("workerType")]
        public Input<string>? WorkerType { get; set; }

        public MLTransformArgs()
        {
        }
    }

    public sealed class MLTransformState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of Glue ML Transform.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the ML Transform.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
        /// </summary>
        [Input("glueVersion")]
        public Input<string>? GlueVersion { get; set; }

        [Input("inputRecordTables")]
        private InputList<Inputs.MLTransformInputRecordTableGetArgs>? _inputRecordTables;

        /// <summary>
        /// A list of AWS Glue table definitions used by the transform. see Input Record Tables.
        /// </summary>
        public InputList<Inputs.MLTransformInputRecordTableGetArgs> InputRecordTables
        {
            get => _inputRecordTables ?? (_inputRecordTables = new InputList<Inputs.MLTransformInputRecordTableGetArgs>());
            set => _inputRecordTables = value;
        }

        /// <summary>
        /// The number of labels available for this transform.
        /// </summary>
        [Input("labelCount")]
        public Input<int>? LabelCount { get; set; }

        /// <summary>
        /// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `max_capacity` is a mutually exclusive option with `number_of_workers` and `worker_type`.
        /// </summary>
        [Input("maxCapacity")]
        public Input<double>? MaxCapacity { get; set; }

        /// <summary>
        /// The maximum number of times to retry this ML Transform if it fails.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The name you assign to this ML Transform. It must be unique in your account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of workers of a defined `worker_type` that are allocated when an ML Transform runs. Required with `worker_type`.
        /// </summary>
        [Input("numberOfWorkers")]
        public Input<int>? NumberOfWorkers { get; set; }

        /// <summary>
        /// The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
        /// </summary>
        [Input("parameters")]
        public Input<Inputs.MLTransformParametersGetArgs>? Parameters { get; set; }

        /// <summary>
        /// The ARN of the IAM role associated with this ML Transform.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("schemas")]
        private InputList<Inputs.MLTransformSchemaGetArgs>? _schemas;

        /// <summary>
        /// The object that represents the schema that this transform accepts. see Schema.
        /// </summary>
        public InputList<Inputs.MLTransformSchemaGetArgs> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<Inputs.MLTransformSchemaGetArgs>());
            set => _schemas = value;
        }

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

        /// <summary>
        /// The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `number_of_workers`.
        /// </summary>
        [Input("workerType")]
        public Input<string>? WorkerType { get; set; }

        public MLTransformState()
        {
        }
    }
}
