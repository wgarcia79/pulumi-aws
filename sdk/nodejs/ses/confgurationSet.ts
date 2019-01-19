// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SES configuration set resource
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_ses_configuration_set_test = new aws.ses.ConfgurationSet("test", {
 *     name: "some-configuration-set-test",
 * });
 * ```
 */
export class ConfgurationSet extends pulumi.CustomResource {
    /**
     * Get an existing ConfgurationSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfgurationSetState, opts?: pulumi.CustomResourceOptions): ConfgurationSet {
        return new ConfgurationSet(name, <any>state, { ...opts, id: id });
    }

    /**
     * The name of the configuration set
     */
    public readonly name: pulumi.Output<string>;

    /**
     * Create a ConfgurationSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfgurationSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfgurationSetArgs | ConfgurationSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ConfgurationSetState = argsOrState as ConfgurationSetState | undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ConfgurationSetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("aws:ses/confgurationSet:ConfgurationSet", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfgurationSet resources.
 */
export interface ConfgurationSetState {
    /**
     * The name of the configuration set
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfgurationSet resource.
 */
export interface ConfgurationSetArgs {
    /**
     * The name of the configuration set
     */
    readonly name?: pulumi.Input<string>;
}
