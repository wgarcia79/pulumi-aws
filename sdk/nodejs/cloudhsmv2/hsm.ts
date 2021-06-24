// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an HSM module in Amazon CloudHSM v2 cluster.
 *
 * ## Example Usage
 *
 * The following example below creates an HSM module in CloudHSM cluster.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const cluster = aws.cloudhsmv2.getCluster({
 *     clusterId: _var.cloudhsm_cluster_id,
 * });
 * const cloudhsmV2Hsm = new aws.cloudhsmv2.Hsm("cloudhsmV2Hsm", {
 *     subnetId: cluster.then(cluster => cluster.subnetIds?[0]),
 *     clusterId: cluster.then(cluster => cluster.clusterId),
 * });
 * ```
 *
 * ## Import
 *
 * HSM modules can be imported using their HSM ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cloudhsmv2/hsm:Hsm bar hsm-quo8dahtaca
 * ```
 */
export class Hsm extends pulumi.CustomResource {
    /**
     * Get an existing Hsm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HsmState, opts?: pulumi.CustomResourceOptions): Hsm {
        return new Hsm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudhsmv2/hsm:Hsm';

    /**
     * Returns true if the given object is an instance of Hsm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hsm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hsm.__pulumiType;
    }

    /**
     * The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The ID of Cloud HSM v2 cluster to which HSM will be added.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The id of the ENI interface allocated for HSM module.
     */
    public /*out*/ readonly hsmEniId!: pulumi.Output<string>;
    /**
     * The id of the HSM module.
     */
    public /*out*/ readonly hsmId!: pulumi.Output<string>;
    /**
     * The state of the HSM module.
     */
    public /*out*/ readonly hsmState!: pulumi.Output<string>;
    /**
     * The IP address of HSM module. Must be within the CIDR of selected subnet.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The ID of subnet in which HSM module will be located.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a Hsm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HsmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HsmArgs | HsmState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HsmState | undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["hsmEniId"] = state ? state.hsmEniId : undefined;
            inputs["hsmId"] = state ? state.hsmId : undefined;
            inputs["hsmState"] = state ? state.hsmState : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as HsmArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["hsmEniId"] = undefined /*out*/;
            inputs["hsmId"] = undefined /*out*/;
            inputs["hsmState"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Hsm.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Hsm resources.
 */
export interface HsmState {
    /**
     * The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The ID of Cloud HSM v2 cluster to which HSM will be added.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The id of the ENI interface allocated for HSM module.
     */
    hsmEniId?: pulumi.Input<string>;
    /**
     * The id of the HSM module.
     */
    hsmId?: pulumi.Input<string>;
    /**
     * The state of the HSM module.
     */
    hsmState?: pulumi.Input<string>;
    /**
     * The IP address of HSM module. Must be within the CIDR of selected subnet.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The ID of subnet in which HSM module will be located.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Hsm resource.
 */
export interface HsmArgs {
    /**
     * The IDs of AZ in which HSM module will be located. Do not use together with subnet_id.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The ID of Cloud HSM v2 cluster to which HSM will be added.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The IP address of HSM module. Must be within the CIDR of selected subnet.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The ID of subnet in which HSM module will be located.
     */
    subnetId?: pulumi.Input<string>;
}
