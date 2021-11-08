// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an RDS DB proxy default target group resource.
 *
 * The `aws.rds.ProxyDefaultTargetGroup` behaves differently from normal resources, in that the provider does not _create_ or _destroy_ this resource, since it implicitly exists as part of an RDS DB Proxy. On the provider resource creation it is automatically imported and on resource destruction, the provider performs no actions in RDS.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleProxy = new aws.rds.Proxy("exampleProxy", {
 *     debugLogging: false,
 *     engineFamily: "MYSQL",
 *     idleClientTimeout: 1800,
 *     requireTls: true,
 *     roleArn: aws_iam_role.example.arn,
 *     vpcSecurityGroupIds: [aws_security_group.example.id],
 *     vpcSubnetIds: [aws_subnet.example.id],
 *     auths: [{
 *         authScheme: "SECRETS",
 *         description: "example",
 *         iamAuth: "DISABLED",
 *         secretArn: aws_secretsmanager_secret.example.arn,
 *     }],
 *     tags: {
 *         Name: "example",
 *         Key: "value",
 *     },
 * });
 * const exampleProxyDefaultTargetGroup = new aws.rds.ProxyDefaultTargetGroup("exampleProxyDefaultTargetGroup", {
 *     dbProxyName: exampleProxy.name,
 *     connectionPoolConfig: {
 *         connectionBorrowTimeout: 120,
 *         initQuery: "SET x=1, y=2",
 *         maxConnectionsPercent: 100,
 *         maxIdleConnectionsPercent: 50,
 *         sessionPinningFilters: ["EXCLUDE_VARIABLE_SETS"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * DB proxy default target groups can be imported using the `db_proxy_name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:rds/proxyDefaultTargetGroup:ProxyDefaultTargetGroup example example
 * ```
 */
export class ProxyDefaultTargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing ProxyDefaultTargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyDefaultTargetGroupState, opts?: pulumi.CustomResourceOptions): ProxyDefaultTargetGroup {
        return new ProxyDefaultTargetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/proxyDefaultTargetGroup:ProxyDefaultTargetGroup';

    /**
     * Returns true if the given object is an instance of ProxyDefaultTargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProxyDefaultTargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProxyDefaultTargetGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) representing the target group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The settings that determine the size and behavior of the connection pool for the target group.
     */
    public readonly connectionPoolConfig!: pulumi.Output<outputs.rds.ProxyDefaultTargetGroupConnectionPoolConfig>;
    /**
     * Name of the RDS DB Proxy.
     */
    public readonly dbProxyName!: pulumi.Output<string>;
    /**
     * The name of the default target group.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;

    /**
     * Create a ProxyDefaultTargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyDefaultTargetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyDefaultTargetGroupArgs | ProxyDefaultTargetGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyDefaultTargetGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["connectionPoolConfig"] = state ? state.connectionPoolConfig : undefined;
            inputs["dbProxyName"] = state ? state.dbProxyName : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProxyDefaultTargetGroupArgs | undefined;
            if ((!args || args.dbProxyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbProxyName'");
            }
            inputs["connectionPoolConfig"] = args ? args.connectionPoolConfig : undefined;
            inputs["dbProxyName"] = args ? args.dbProxyName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProxyDefaultTargetGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProxyDefaultTargetGroup resources.
 */
export interface ProxyDefaultTargetGroupState {
    /**
     * The Amazon Resource Name (ARN) representing the target group.
     */
    arn?: pulumi.Input<string>;
    /**
     * The settings that determine the size and behavior of the connection pool for the target group.
     */
    connectionPoolConfig?: pulumi.Input<inputs.rds.ProxyDefaultTargetGroupConnectionPoolConfig>;
    /**
     * Name of the RDS DB Proxy.
     */
    dbProxyName?: pulumi.Input<string>;
    /**
     * The name of the default target group.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProxyDefaultTargetGroup resource.
 */
export interface ProxyDefaultTargetGroupArgs {
    /**
     * The settings that determine the size and behavior of the connection pool for the target group.
     */
    connectionPoolConfig?: pulumi.Input<inputs.rds.ProxyDefaultTargetGroupConnectionPoolConfig>;
    /**
     * Name of the RDS DB Proxy.
     */
    dbProxyName: pulumi.Input<string>;
}
