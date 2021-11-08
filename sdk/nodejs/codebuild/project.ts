// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a CodeBuild Project resource. See also the `aws.codebuild.Webhook` resource, which manages the webhook to the source (e.g., the "rebuild every time a code change is pushed" option in the CodeBuild web console).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucket = new aws.s3.Bucket("exampleBucket", {acl: "private"});
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "codebuild.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `});
 * const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
 *     role: exampleRole.name,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Resource": [
 *         "*"
 *       ],
 *       "Action": [
 *         "logs:CreateLogGroup",
 *         "logs:CreateLogStream",
 *         "logs:PutLogEvents"
 *       ]
 *     },
 *     {
 *       "Effect": "Allow",
 *       "Action": [
 *         "ec2:CreateNetworkInterface",
 *         "ec2:DescribeDhcpOptions",
 *         "ec2:DescribeNetworkInterfaces",
 *         "ec2:DeleteNetworkInterface",
 *         "ec2:DescribeSubnets",
 *         "ec2:DescribeSecurityGroups",
 *         "ec2:DescribeVpcs"
 *       ],
 *       "Resource": "*"
 *     },
 *     {
 *       "Effect": "Allow",
 *       "Action": [
 *         "ec2:CreateNetworkInterfacePermission"
 *       ],
 *       "Resource": [
 *         "arn:aws:ec2:us-east-1:123456789012:network-interface/*"
 *       ],
 *       "Condition": {
 *         "StringEquals": {
 *           "ec2:Subnet": [
 *             "${aws_subnet.example1.arn}",
 *             "${aws_subnet.example2.arn}"
 *           ],
 *           "ec2:AuthorizedService": "codebuild.amazonaws.com"
 *         }
 *       }
 *     },
 *     {
 *       "Effect": "Allow",
 *       "Action": [
 *         "s3:*"
 *       ],
 *       "Resource": [
 *         "${exampleBucket.arn}",
 *         "${exampleBucket.arn}/*"
 *       ]
 *     }
 *   ]
 * }
 * `,
 * });
 * const exampleProject = new aws.codebuild.Project("exampleProject", {
 *     description: "test_codebuild_project",
 *     buildTimeout: "5",
 *     serviceRole: exampleRole.arn,
 *     artifacts: {
 *         type: "NO_ARTIFACTS",
 *     },
 *     cache: {
 *         type: "S3",
 *         location: exampleBucket.bucket,
 *     },
 *     environment: {
 *         computeType: "BUILD_GENERAL1_SMALL",
 *         image: "aws/codebuild/standard:1.0",
 *         type: "LINUX_CONTAINER",
 *         imagePullCredentialsType: "CODEBUILD",
 *         environmentVariables: [
 *             {
 *                 name: "SOME_KEY1",
 *                 value: "SOME_VALUE1",
 *             },
 *             {
 *                 name: "SOME_KEY2",
 *                 value: "SOME_VALUE2",
 *                 type: "PARAMETER_STORE",
 *             },
 *         ],
 *     },
 *     logsConfig: {
 *         cloudwatchLogs: {
 *             groupName: "log-group",
 *             streamName: "log-stream",
 *         },
 *         s3Logs: {
 *             status: "ENABLED",
 *             location: pulumi.interpolate`${exampleBucket.id}/build-log`,
 *         },
 *     },
 *     source: {
 *         type: "GITHUB",
 *         location: "https://github.com/mitchellh/packer.git",
 *         gitCloneDepth: 1,
 *         gitSubmodulesConfig: {
 *             fetchSubmodules: true,
 *         },
 *     },
 *     sourceVersion: "master",
 *     vpcConfig: {
 *         vpcId: aws_vpc.example.id,
 *         subnets: [
 *             aws_subnet.example1.id,
 *             aws_subnet.example2.id,
 *         ],
 *         securityGroupIds: [
 *             aws_security_group.example1.id,
 *             aws_security_group.example2.id,
 *         ],
 *     },
 *     tags: {
 *         Environment: "Test",
 *     },
 * });
 * const project_with_cache = new aws.codebuild.Project("project-with-cache", {
 *     description: "test_codebuild_project_cache",
 *     buildTimeout: "5",
 *     queuedTimeout: "5",
 *     serviceRole: exampleRole.arn,
 *     artifacts: {
 *         type: "NO_ARTIFACTS",
 *     },
 *     cache: {
 *         type: "LOCAL",
 *         modes: [
 *             "LOCAL_DOCKER_LAYER_CACHE",
 *             "LOCAL_SOURCE_CACHE",
 *         ],
 *     },
 *     environment: {
 *         computeType: "BUILD_GENERAL1_SMALL",
 *         image: "aws/codebuild/standard:1.0",
 *         type: "LINUX_CONTAINER",
 *         imagePullCredentialsType: "CODEBUILD",
 *         environmentVariables: [{
 *             name: "SOME_KEY1",
 *             value: "SOME_VALUE1",
 *         }],
 *     },
 *     source: {
 *         type: "GITHUB",
 *         location: "https://github.com/mitchellh/packer.git",
 *         gitCloneDepth: 1,
 *     },
 *     tags: {
 *         Environment: "Test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * CodeBuild Project can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:codebuild/project:Project name project-name
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codebuild/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * ARN of the CodeBuild project.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly artifacts!: pulumi.Output<outputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    public readonly badgeEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * URL of the build badge when `badgeEnabled` is enabled.
     */
    public /*out*/ readonly badgeUrl!: pulumi.Output<string>;
    /**
     * Defines the batch build options for the project.
     */
    public readonly buildBatchConfig!: pulumi.Output<outputs.codebuild.ProjectBuildBatchConfig | undefined>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    public readonly buildTimeout!: pulumi.Output<number | undefined>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly cache!: pulumi.Output<outputs.codebuild.ProjectCache | undefined>;
    /**
     * Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
     */
    public readonly concurrentBuildLimit!: pulumi.Output<number | undefined>;
    /**
     * Short description of the project.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    public readonly encryptionKey!: pulumi.Output<string>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly environment!: pulumi.Output<outputs.codebuild.ProjectEnvironment>;
    /**
     * A set of file system locations to to mount inside the build. File system locations are documented below.
     */
    public readonly fileSystemLocations!: pulumi.Output<outputs.codebuild.ProjectFileSystemLocation[] | undefined>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly logsConfig!: pulumi.Output<outputs.codebuild.ProjectLogsConfig | undefined>;
    /**
     * Name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    public readonly queuedTimeout!: pulumi.Output<number | undefined>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly secondaryArtifacts!: pulumi.Output<outputs.codebuild.ProjectSecondaryArtifact[] | undefined>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly secondarySources!: pulumi.Output<outputs.codebuild.ProjectSecondarySource[] | undefined>;
    /**
     * Specifies the service role ARN for the batch build project.
     */
    public readonly serviceRole!: pulumi.Output<string>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly source!: pulumi.Output<outputs.codebuild.ProjectSource>;
    /**
     * Version of the build input to be built for this project. If not specified, the latest version is used.
     */
    public readonly sourceVersion!: pulumi.Output<string | undefined>;
    /**
     * Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block. Detailed below.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.codebuild.ProjectVpcConfig | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["artifacts"] = state ? state.artifacts : undefined;
            inputs["badgeEnabled"] = state ? state.badgeEnabled : undefined;
            inputs["badgeUrl"] = state ? state.badgeUrl : undefined;
            inputs["buildBatchConfig"] = state ? state.buildBatchConfig : undefined;
            inputs["buildTimeout"] = state ? state.buildTimeout : undefined;
            inputs["cache"] = state ? state.cache : undefined;
            inputs["concurrentBuildLimit"] = state ? state.concurrentBuildLimit : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["fileSystemLocations"] = state ? state.fileSystemLocations : undefined;
            inputs["logsConfig"] = state ? state.logsConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["queuedTimeout"] = state ? state.queuedTimeout : undefined;
            inputs["secondaryArtifacts"] = state ? state.secondaryArtifacts : undefined;
            inputs["secondarySources"] = state ? state.secondarySources : undefined;
            inputs["serviceRole"] = state ? state.serviceRole : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["sourceVersion"] = state ? state.sourceVersion : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.artifacts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'artifacts'");
            }
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.serviceRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceRole'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            inputs["artifacts"] = args ? args.artifacts : undefined;
            inputs["badgeEnabled"] = args ? args.badgeEnabled : undefined;
            inputs["buildBatchConfig"] = args ? args.buildBatchConfig : undefined;
            inputs["buildTimeout"] = args ? args.buildTimeout : undefined;
            inputs["cache"] = args ? args.cache : undefined;
            inputs["concurrentBuildLimit"] = args ? args.concurrentBuildLimit : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["fileSystemLocations"] = args ? args.fileSystemLocations : undefined;
            inputs["logsConfig"] = args ? args.logsConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["queuedTimeout"] = args ? args.queuedTimeout : undefined;
            inputs["secondaryArtifacts"] = args ? args.secondaryArtifacts : undefined;
            inputs["secondarySources"] = args ? args.secondarySources : undefined;
            inputs["serviceRole"] = args ? args.serviceRole : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["sourceVersion"] = args ? args.sourceVersion : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["badgeUrl"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * ARN of the CodeBuild project.
     */
    arn?: pulumi.Input<string>;
    /**
     * Configuration block. Detailed below.
     */
    artifacts?: pulumi.Input<inputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    badgeEnabled?: pulumi.Input<boolean>;
    /**
     * URL of the build badge when `badgeEnabled` is enabled.
     */
    badgeUrl?: pulumi.Input<string>;
    /**
     * Defines the batch build options for the project.
     */
    buildBatchConfig?: pulumi.Input<inputs.codebuild.ProjectBuildBatchConfig>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    buildTimeout?: pulumi.Input<number>;
    /**
     * Configuration block. Detailed below.
     */
    cache?: pulumi.Input<inputs.codebuild.ProjectCache>;
    /**
     * Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
     */
    concurrentBuildLimit?: pulumi.Input<number>;
    /**
     * Short description of the project.
     */
    description?: pulumi.Input<string>;
    /**
     * AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * Configuration block. Detailed below.
     */
    environment?: pulumi.Input<inputs.codebuild.ProjectEnvironment>;
    /**
     * A set of file system locations to to mount inside the build. File system locations are documented below.
     */
    fileSystemLocations?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectFileSystemLocation>[]>;
    /**
     * Configuration block. Detailed below.
     */
    logsConfig?: pulumi.Input<inputs.codebuild.ProjectLogsConfig>;
    /**
     * Name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    name?: pulumi.Input<string>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    queuedTimeout?: pulumi.Input<number>;
    /**
     * Configuration block. Detailed below.
     */
    secondaryArtifacts?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondaryArtifact>[]>;
    /**
     * Configuration block. Detailed below.
     */
    secondarySources?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondarySource>[]>;
    /**
     * Specifies the service role ARN for the batch build project.
     */
    serviceRole?: pulumi.Input<string>;
    /**
     * Configuration block. Detailed below.
     */
    source?: pulumi.Input<inputs.codebuild.ProjectSource>;
    /**
     * Version of the build input to be built for this project. If not specified, the latest version is used.
     */
    sourceVersion?: pulumi.Input<string>;
    /**
     * Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block. Detailed below.
     */
    vpcConfig?: pulumi.Input<inputs.codebuild.ProjectVpcConfig>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Configuration block. Detailed below.
     */
    artifacts: pulumi.Input<inputs.codebuild.ProjectArtifacts>;
    /**
     * Generates a publicly-accessible URL for the projects build badge. Available as `badgeUrl` attribute when enabled.
     */
    badgeEnabled?: pulumi.Input<boolean>;
    /**
     * Defines the batch build options for the project.
     */
    buildBatchConfig?: pulumi.Input<inputs.codebuild.ProjectBuildBatchConfig>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
     */
    buildTimeout?: pulumi.Input<number>;
    /**
     * Configuration block. Detailed below.
     */
    cache?: pulumi.Input<inputs.codebuild.ProjectCache>;
    /**
     * Specify a maximum number of concurrent builds for the project. The value specified must be greater than 0 and less than the account concurrent running builds limit.
     */
    concurrentBuildLimit?: pulumi.Input<number>;
    /**
     * Short description of the project.
     */
    description?: pulumi.Input<string>;
    /**
     * AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
     */
    encryptionKey?: pulumi.Input<string>;
    /**
     * Configuration block. Detailed below.
     */
    environment: pulumi.Input<inputs.codebuild.ProjectEnvironment>;
    /**
     * A set of file system locations to to mount inside the build. File system locations are documented below.
     */
    fileSystemLocations?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectFileSystemLocation>[]>;
    /**
     * Configuration block. Detailed below.
     */
    logsConfig?: pulumi.Input<inputs.codebuild.ProjectLogsConfig>;
    /**
     * Name of the project. If `type` is set to `S3`, this is the name of the output artifact object
     */
    name?: pulumi.Input<string>;
    /**
     * Number of minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
     */
    queuedTimeout?: pulumi.Input<number>;
    /**
     * Configuration block. Detailed below.
     */
    secondaryArtifacts?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondaryArtifact>[]>;
    /**
     * Configuration block. Detailed below.
     */
    secondarySources?: pulumi.Input<pulumi.Input<inputs.codebuild.ProjectSecondarySource>[]>;
    /**
     * Specifies the service role ARN for the batch build project.
     */
    serviceRole: pulumi.Input<string>;
    /**
     * Configuration block. Detailed below.
     */
    source: pulumi.Input<inputs.codebuild.ProjectSource>;
    /**
     * Version of the build input to be built for this project. If not specified, the latest version is used.
     */
    sourceVersion?: pulumi.Input<string>;
    /**
     * Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block. Detailed below.
     */
    vpcConfig?: pulumi.Input<inputs.codebuild.ProjectVpcConfig>;
}
