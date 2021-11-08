// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an AppSync DataSource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleTable = new aws.dynamodb.Table("exampleTable", {
 *     readCapacity: 1,
 *     writeCapacity: 1,
 *     hashKey: "UserId",
 *     attributes: [{
 *         name: "UserId",
 *         type: "S",
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "appsync.amazonaws.com"
 *       },
 *       "Effect": "Allow"
 *     }
 *   ]
 * }
 * `});
 * const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
 *     role: exampleRole.id,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "dynamodb:*"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${exampleTable.arn}"
 *       ]
 *     }
 *   ]
 * }
 * `,
 * });
 * const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {authenticationType: "API_KEY"});
 * const exampleDataSource = new aws.appsync.DataSource("exampleDataSource", {
 *     apiId: exampleGraphQLApi.id,
 *     name: "tf_appsync_example",
 *     serviceRoleArn: exampleRole.arn,
 *     type: "AMAZON_DYNAMODB",
 *     dynamodbConfig: {
 *         tableName: exampleTable.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_appsync_datasource` can be imported with their `api_id`, a hyphen, and `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:appsync/dataSource:DataSource example abcdef123456-example
 * ```
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceState, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appsync/dataSource:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * The API ID for the GraphQL API for the DataSource.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The ARN
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the DataSource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * DynamoDB settings. See below
     */
    public readonly dynamodbConfig!: pulumi.Output<outputs.appsync.DataSourceDynamodbConfig | undefined>;
    /**
     * Amazon Elasticsearch settings. See below
     */
    public readonly elasticsearchConfig!: pulumi.Output<outputs.appsync.DataSourceElasticsearchConfig | undefined>;
    /**
     * HTTP settings. See below
     */
    public readonly httpConfig!: pulumi.Output<outputs.appsync.DataSourceHttpConfig | undefined>;
    /**
     * AWS Lambda settings. See below
     */
    public readonly lambdaConfig!: pulumi.Output<outputs.appsync.DataSourceLambdaConfig | undefined>;
    /**
     * A user-supplied name for the DataSource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The IAM service role ARN for the data source.
     */
    public readonly serviceRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceArgs | DataSourceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataSourceState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamodbConfig"] = state ? state.dynamodbConfig : undefined;
            inputs["elasticsearchConfig"] = state ? state.elasticsearchConfig : undefined;
            inputs["httpConfig"] = state ? state.httpConfig : undefined;
            inputs["lambdaConfig"] = state ? state.lambdaConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DataSourceArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamodbConfig"] = args ? args.dynamodbConfig : undefined;
            inputs["elasticsearchConfig"] = args ? args.elasticsearchConfig : undefined;
            inputs["httpConfig"] = args ? args.httpConfig : undefined;
            inputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataSource.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSource resources.
 */
export interface DataSourceState {
    /**
     * The API ID for the GraphQL API for the DataSource.
     */
    apiId?: pulumi.Input<string>;
    /**
     * The ARN
     */
    arn?: pulumi.Input<string>;
    /**
     * A description of the DataSource.
     */
    description?: pulumi.Input<string>;
    /**
     * DynamoDB settings. See below
     */
    dynamodbConfig?: pulumi.Input<inputs.appsync.DataSourceDynamodbConfig>;
    /**
     * Amazon Elasticsearch settings. See below
     */
    elasticsearchConfig?: pulumi.Input<inputs.appsync.DataSourceElasticsearchConfig>;
    /**
     * HTTP settings. See below
     */
    httpConfig?: pulumi.Input<inputs.appsync.DataSourceHttpConfig>;
    /**
     * AWS Lambda settings. See below
     */
    lambdaConfig?: pulumi.Input<inputs.appsync.DataSourceLambdaConfig>;
    /**
     * A user-supplied name for the DataSource.
     */
    name?: pulumi.Input<string>;
    /**
     * The IAM service role ARN for the data source.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * The API ID for the GraphQL API for the DataSource.
     */
    apiId: pulumi.Input<string>;
    /**
     * A description of the DataSource.
     */
    description?: pulumi.Input<string>;
    /**
     * DynamoDB settings. See below
     */
    dynamodbConfig?: pulumi.Input<inputs.appsync.DataSourceDynamodbConfig>;
    /**
     * Amazon Elasticsearch settings. See below
     */
    elasticsearchConfig?: pulumi.Input<inputs.appsync.DataSourceElasticsearchConfig>;
    /**
     * HTTP settings. See below
     */
    httpConfig?: pulumi.Input<inputs.appsync.DataSourceHttpConfig>;
    /**
     * AWS Lambda settings. See below
     */
    lambdaConfig?: pulumi.Input<inputs.appsync.DataSourceLambdaConfig>;
    /**
     * A user-supplied name for the DataSource.
     */
    name?: pulumi.Input<string>;
    /**
     * The IAM service role ARN for the data source.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * The type of the DataSource. Valid values: `AWS_LAMBDA`, `AMAZON_DYNAMODB`, `AMAZON_ELASTICSEARCH`, `HTTP`, `NONE`.
     */
    type: pulumi.Input<string>;
}
