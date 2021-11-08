// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy
{
    /// <summary>
    /// Provides a CodeDeploy application to be used as a basis for deployments
    /// 
    /// ## Example Usage
    /// ### ECS Application
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.CodeDeploy.Application("example", new Aws.CodeDeploy.ApplicationArgs
    ///         {
    ///             ComputePlatform = "ECS",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Lambda Application
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.CodeDeploy.Application("example", new Aws.CodeDeploy.ApplicationArgs
    ///         {
    ///             ComputePlatform = "Lambda",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Server Application
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.CodeDeploy.Application("example", new Aws.CodeDeploy.ApplicationArgs
    ///         {
    ///             ComputePlatform = "Server",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeDeploy Applications can be imported using the `name`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:codedeploy/application:Application example my-application
    /// ```
    /// </summary>
    [AwsResourceType("aws:codedeploy/application:Application")]
    public partial class Application : Pulumi.CustomResource
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the CodeDeploy application.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        /// </summary>
        [Output("computePlatform")]
        public Output<string?> ComputePlatform { get; private set; } = null!;

        /// <summary>
        /// The name for a connection to a GitHub account.
        /// </summary>
        [Output("githubAccountName")]
        public Output<string> GithubAccountName { get; private set; } = null!;

        /// <summary>
        /// Whether the user has authenticated with GitHub for the specified application.
        /// </summary>
        [Output("linkedToGithub")]
        public Output<bool> LinkedToGithub { get; private set; } = null!;

        /// <summary>
        /// The name of the application.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:codedeploy/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("aws:codedeploy/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        /// </summary>
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        /// <summary>
        /// The name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public ApplicationArgs()
        {
        }
    }

    public sealed class ApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The ARN of the CodeDeploy application.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
        /// </summary>
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        /// <summary>
        /// The name for a connection to a GitHub account.
        /// </summary>
        [Input("githubAccountName")]
        public Input<string>? GithubAccountName { get; set; }

        /// <summary>
        /// Whether the user has authenticated with GitHub for the specified application.
        /// </summary>
        [Input("linkedToGithub")]
        public Input<bool>? LinkedToGithub { get; set; }

        /// <summary>
        /// The name of the application.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public ApplicationState()
        {
        }
    }
}
