// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Amazon Managed Service for Prometheus (AMP) Alert Manager Definition
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const demoWorkspace = new aws.amp.Workspace("demoWorkspace", {});
 * const demoAlertManagerDefinition = new aws.amp.AlertManagerDefinition("demoAlertManagerDefinition", {
 *     workspaceId: demoWorkspace.id,
 *     definition: `alertmanager_config: |
 *   route:
 *     receiver: 'default'
 *   receivers:
 *     - name: 'default'
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * The prometheus alert manager definition can be imported using the workspace identifier, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:amp/alertManagerDefinition:AlertManagerDefinition demo ws-C6DCB907-F2D7-4D96-957B-66691F865D8B
 * ```
 */
export class AlertManagerDefinition extends pulumi.CustomResource {
    /**
     * Get an existing AlertManagerDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlertManagerDefinitionState, opts?: pulumi.CustomResourceOptions): AlertManagerDefinition {
        return new AlertManagerDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:amp/alertManagerDefinition:AlertManagerDefinition';

    /**
     * Returns true if the given object is an instance of AlertManagerDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertManagerDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertManagerDefinition.__pulumiType;
    }

    /**
     * the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
     */
    public readonly definition!: pulumi.Output<string>;
    /**
     * The id of the prometheus workspace the alert manager definition should be linked to
     */
    public readonly workspaceId!: pulumi.Output<string>;

    /**
     * Create a AlertManagerDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertManagerDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlertManagerDefinitionArgs | AlertManagerDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlertManagerDefinitionState | undefined;
            inputs["definition"] = state ? state.definition : undefined;
            inputs["workspaceId"] = state ? state.workspaceId : undefined;
        } else {
            const args = argsOrState as AlertManagerDefinitionArgs | undefined;
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.workspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceId'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["workspaceId"] = args ? args.workspaceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AlertManagerDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AlertManagerDefinition resources.
 */
export interface AlertManagerDefinitionState {
    /**
     * the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
     */
    definition?: pulumi.Input<string>;
    /**
     * The id of the prometheus workspace the alert manager definition should be linked to
     */
    workspaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AlertManagerDefinition resource.
 */
export interface AlertManagerDefinitionArgs {
    /**
     * the alert manager definition that you want to be applied. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html).
     */
    definition: pulumi.Input<string>;
    /**
     * The id of the prometheus workspace the alert manager definition should be linked to
     */
    workspaceId: pulumi.Input<string>;
}
