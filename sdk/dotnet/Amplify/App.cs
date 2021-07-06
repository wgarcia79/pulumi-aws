// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Amplify
{
    /// <summary>
    /// ## Import
    /// 
    /// Amplify App can be imported using Amplify App ID (appId), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:amplify/app:App example d2ypk4k47z8u6
    /// ```
    /// 
    ///  App ID can be obtained from App ARN (e.g. `arn:aws:amplify:us-east-1:12345678:apps/d2ypk4k47z8u6`).
    /// </summary>
    [AwsResourceType("aws:amplify/app:App")]
    public partial class App : Pulumi.CustomResource
    {
        /// <summary>
        /// The personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
        /// </summary>
        [Output("accessToken")]
        public Output<string?> AccessToken { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amplify app.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The automated branch creation configuration for an Amplify app. An `auto_branch_creation_config` block is documented below.
        /// </summary>
        [Output("autoBranchCreationConfig")]
        public Output<Outputs.AppAutoBranchCreationConfig> AutoBranchCreationConfig { get; private set; } = null!;

        /// <summary>
        /// The automated branch creation glob patterns for an Amplify app.
        /// </summary>
        [Output("autoBranchCreationPatterns")]
        public Output<ImmutableArray<string>> AutoBranchCreationPatterns { get; private set; } = null!;

        /// <summary>
        /// The credentials for basic authorization for an Amplify app.
        /// </summary>
        [Output("basicAuthCredentials")]
        public Output<string?> BasicAuthCredentials { get; private set; } = null!;

        /// <summary>
        /// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
        /// </summary>
        [Output("buildSpec")]
        public Output<string> BuildSpec { get; private set; } = null!;

        /// <summary>
        /// The custom rewrite and redirect rules for an Amplify app. A `custom_rule` block is documented below.
        /// </summary>
        [Output("customRules")]
        public Output<ImmutableArray<Outputs.AppCustomRule>> CustomRules { get; private set; } = null!;

        /// <summary>
        /// The default domain for the Amplify app.
        /// </summary>
        [Output("defaultDomain")]
        public Output<string> DefaultDomain { get; private set; } = null!;

        /// <summary>
        /// The description for an Amplify app.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Enables automated branch creation for an Amplify app.
        /// </summary>
        [Output("enableAutoBranchCreation")]
        public Output<bool?> EnableAutoBranchCreation { get; private set; } = null!;

        /// <summary>
        /// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
        /// </summary>
        [Output("enableBasicAuth")]
        public Output<bool?> EnableBasicAuth { get; private set; } = null!;

        /// <summary>
        /// Enables auto-building of branches for the Amplify App.
        /// </summary>
        [Output("enableBranchAutoBuild")]
        public Output<bool?> EnableBranchAutoBuild { get; private set; } = null!;

        /// <summary>
        /// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
        /// </summary>
        [Output("enableBranchAutoDeletion")]
        public Output<bool?> EnableBranchAutoDeletion { get; private set; } = null!;

        /// <summary>
        /// The environment variables map for an Amplify app.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, string>?> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// The AWS Identity and Access Management (IAM) service role for an Amplify app.
        /// </summary>
        [Output("iamServiceRoleArn")]
        public Output<string?> IamServiceRoleArn { get; private set; } = null!;

        /// <summary>
        /// The name for an Amplify app.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
        /// </summary>
        [Output("oauthToken")]
        public Output<string?> OauthToken { get; private set; } = null!;

        /// <summary>
        /// The platform or framework for an Amplify app. Valid values: `WEB`.
        /// </summary>
        [Output("platform")]
        public Output<string?> Platform { get; private set; } = null!;

        /// <summary>
        /// Describes the information about a production branch for an Amplify app. A `production_branch` block is documented below.
        /// </summary>
        [Output("productionBranches")]
        public Output<ImmutableArray<Outputs.AppProductionBranch>> ProductionBranches { get; private set; } = null!;

        /// <summary>
        /// The repository for an Amplify app.
        /// </summary>
        [Output("repository")]
        public Output<string?> Repository { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:amplify/app:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
            : base("aws:amplify/app:App", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, AppState? state = null, CustomResourceOptions? options = null)
        {
            return new App(name, id, state, options);
        }
    }

    public sealed class AppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// The automated branch creation configuration for an Amplify app. An `auto_branch_creation_config` block is documented below.
        /// </summary>
        [Input("autoBranchCreationConfig")]
        public Input<Inputs.AppAutoBranchCreationConfigArgs>? AutoBranchCreationConfig { get; set; }

        [Input("autoBranchCreationPatterns")]
        private InputList<string>? _autoBranchCreationPatterns;

        /// <summary>
        /// The automated branch creation glob patterns for an Amplify app.
        /// </summary>
        public InputList<string> AutoBranchCreationPatterns
        {
            get => _autoBranchCreationPatterns ?? (_autoBranchCreationPatterns = new InputList<string>());
            set => _autoBranchCreationPatterns = value;
        }

        /// <summary>
        /// The credentials for basic authorization for an Amplify app.
        /// </summary>
        [Input("basicAuthCredentials")]
        public Input<string>? BasicAuthCredentials { get; set; }

        /// <summary>
        /// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
        /// </summary>
        [Input("buildSpec")]
        public Input<string>? BuildSpec { get; set; }

        [Input("customRules")]
        private InputList<Inputs.AppCustomRuleArgs>? _customRules;

        /// <summary>
        /// The custom rewrite and redirect rules for an Amplify app. A `custom_rule` block is documented below.
        /// </summary>
        public InputList<Inputs.AppCustomRuleArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.AppCustomRuleArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// The description for an Amplify app.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables automated branch creation for an Amplify app.
        /// </summary>
        [Input("enableAutoBranchCreation")]
        public Input<bool>? EnableAutoBranchCreation { get; set; }

        /// <summary>
        /// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
        /// </summary>
        [Input("enableBasicAuth")]
        public Input<bool>? EnableBasicAuth { get; set; }

        /// <summary>
        /// Enables auto-building of branches for the Amplify App.
        /// </summary>
        [Input("enableBranchAutoBuild")]
        public Input<bool>? EnableBranchAutoBuild { get; set; }

        /// <summary>
        /// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
        /// </summary>
        [Input("enableBranchAutoDeletion")]
        public Input<bool>? EnableBranchAutoDeletion { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables map for an Amplify app.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The AWS Identity and Access Management (IAM) service role for an Amplify app.
        /// </summary>
        [Input("iamServiceRoleArn")]
        public Input<string>? IamServiceRoleArn { get; set; }

        /// <summary>
        /// The name for an Amplify app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
        /// </summary>
        [Input("oauthToken")]
        public Input<string>? OauthToken { get; set; }

        /// <summary>
        /// The platform or framework for an Amplify app. Valid values: `WEB`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The repository for an Amplify app.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public AppArgs()
        {
        }
    }

    public sealed class AppState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The personal access token for a third-party source control system for an Amplify app. The personal access token is used to create a webhook and a read-only deploy key. The token is not stored.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amplify app.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The automated branch creation configuration for an Amplify app. An `auto_branch_creation_config` block is documented below.
        /// </summary>
        [Input("autoBranchCreationConfig")]
        public Input<Inputs.AppAutoBranchCreationConfigGetArgs>? AutoBranchCreationConfig { get; set; }

        [Input("autoBranchCreationPatterns")]
        private InputList<string>? _autoBranchCreationPatterns;

        /// <summary>
        /// The automated branch creation glob patterns for an Amplify app.
        /// </summary>
        public InputList<string> AutoBranchCreationPatterns
        {
            get => _autoBranchCreationPatterns ?? (_autoBranchCreationPatterns = new InputList<string>());
            set => _autoBranchCreationPatterns = value;
        }

        /// <summary>
        /// The credentials for basic authorization for an Amplify app.
        /// </summary>
        [Input("basicAuthCredentials")]
        public Input<string>? BasicAuthCredentials { get; set; }

        /// <summary>
        /// The [build specification](https://docs.aws.amazon.com/amplify/latest/userguide/build-settings.html) (build spec) for an Amplify app.
        /// </summary>
        [Input("buildSpec")]
        public Input<string>? BuildSpec { get; set; }

        [Input("customRules")]
        private InputList<Inputs.AppCustomRuleGetArgs>? _customRules;

        /// <summary>
        /// The custom rewrite and redirect rules for an Amplify app. A `custom_rule` block is documented below.
        /// </summary>
        public InputList<Inputs.AppCustomRuleGetArgs> CustomRules
        {
            get => _customRules ?? (_customRules = new InputList<Inputs.AppCustomRuleGetArgs>());
            set => _customRules = value;
        }

        /// <summary>
        /// The default domain for the Amplify app.
        /// </summary>
        [Input("defaultDomain")]
        public Input<string>? DefaultDomain { get; set; }

        /// <summary>
        /// The description for an Amplify app.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables automated branch creation for an Amplify app.
        /// </summary>
        [Input("enableAutoBranchCreation")]
        public Input<bool>? EnableAutoBranchCreation { get; set; }

        /// <summary>
        /// Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
        /// </summary>
        [Input("enableBasicAuth")]
        public Input<bool>? EnableBasicAuth { get; set; }

        /// <summary>
        /// Enables auto-building of branches for the Amplify App.
        /// </summary>
        [Input("enableBranchAutoBuild")]
        public Input<bool>? EnableBranchAutoBuild { get; set; }

        /// <summary>
        /// Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
        /// </summary>
        [Input("enableBranchAutoDeletion")]
        public Input<bool>? EnableBranchAutoDeletion { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables map for an Amplify app.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The AWS Identity and Access Management (IAM) service role for an Amplify app.
        /// </summary>
        [Input("iamServiceRoleArn")]
        public Input<string>? IamServiceRoleArn { get; set; }

        /// <summary>
        /// The name for an Amplify app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key. The OAuth token is not stored.
        /// </summary>
        [Input("oauthToken")]
        public Input<string>? OauthToken { get; set; }

        /// <summary>
        /// The platform or framework for an Amplify app. Valid values: `WEB`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        [Input("productionBranches")]
        private InputList<Inputs.AppProductionBranchGetArgs>? _productionBranches;

        /// <summary>
        /// Describes the information about a production branch for an Amplify app. A `production_branch` block is documented below.
        /// </summary>
        public InputList<Inputs.AppProductionBranchGetArgs> ProductionBranches
        {
            get => _productionBranches ?? (_productionBranches = new InputList<Inputs.AppProductionBranchGetArgs>());
            set => _productionBranches = value;
        }

        /// <summary>
        /// The repository for an Amplify app.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        public AppState()
        {
        }
    }
}
