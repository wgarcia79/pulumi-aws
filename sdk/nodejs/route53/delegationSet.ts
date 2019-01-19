// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a [Route53 Delegation Set](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html) resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_route53_delegation_set_main = new aws.route53.DelegationSet("main", {
 *     referenceName: "DynDNS",
 * });
 * const aws_route53_zone_primary = new aws.route53.Zone("primary", {
 *     delegationSetId: aws_route53_delegation_set_main.id,
 *     name: "hashicorp.com",
 * });
 * const aws_route53_zone_secondary = new aws.route53.Zone("secondary", {
 *     delegationSetId: aws_route53_delegation_set_main.id,
 *     name: "terraform.io",
 * });
 * ```
 */
export class DelegationSet extends pulumi.CustomResource {
    /**
     * Get an existing DelegationSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DelegationSetState, opts?: pulumi.CustomResourceOptions): DelegationSet {
        return new DelegationSet(name, <any>state, { ...opts, id: id });
    }

    /**
     * A list of authoritative name servers for the hosted zone
     * (effectively a list of NS records).
     */
    public /*out*/ readonly nameServers: pulumi.Output<string[]>;
    /**
     * This is a reference name used in Caller Reference
     * (helpful for identifying single delegation set amongst others)
     */
    public readonly referenceName: pulumi.Output<string | undefined>;

    /**
     * Create a DelegationSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DelegationSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DelegationSetArgs | DelegationSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DelegationSetState = argsOrState as DelegationSetState | undefined;
            inputs["nameServers"] = state ? state.nameServers : undefined;
            inputs["referenceName"] = state ? state.referenceName : undefined;
        } else {
            const args = argsOrState as DelegationSetArgs | undefined;
            inputs["referenceName"] = args ? args.referenceName : undefined;
            inputs["nameServers"] = undefined /*out*/;
        }
        super("aws:route53/delegationSet:DelegationSet", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DelegationSet resources.
 */
export interface DelegationSetState {
    /**
     * A list of authoritative name servers for the hosted zone
     * (effectively a list of NS records).
     */
    readonly nameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This is a reference name used in Caller Reference
     * (helpful for identifying single delegation set amongst others)
     */
    readonly referenceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DelegationSet resource.
 */
export interface DelegationSetArgs {
    /**
     * This is a reference name used in Caller Reference
     * (helpful for identifying single delegation set amongst others)
     */
    readonly referenceName?: pulumi.Input<string>;
}
