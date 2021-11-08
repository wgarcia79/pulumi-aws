// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an IP access control group in AWS WorkSpaces Service
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const contractors = new aws.workspaces.IpGroup("contractors", {
 *     description: "Contractors IP access control group",
 *     rules: [
 *         {
 *             description: "NY",
 *             source: "150.24.14.0/24",
 *         },
 *         {
 *             description: "LA",
 *             source: "125.191.14.85/32",
 *         },
 *         {
 *             description: "STL",
 *             source: "44.98.100.0/24",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * WorkSpaces IP groups can be imported using their GroupID, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:workspaces/ipGroup:IpGroup example wsipg-488lrtl3k
 * ```
 */
export class IpGroup extends pulumi.CustomResource {
    /**
     * Get an existing IpGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpGroupState, opts?: pulumi.CustomResourceOptions): IpGroup {
        return new IpGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:workspaces/ipGroup:IpGroup';

    /**
     * Returns true if the given object is an instance of IpGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpGroup.__pulumiType;
    }

    /**
     * The description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the IP group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    public readonly rules!: pulumi.Output<outputs.workspaces.IpGroupRule[] | undefined>;
    /**
     * A map of tags assigned to the WorkSpaces directory. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a IpGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpGroupArgs | IpGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as IpGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IpGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpGroup resources.
 */
export interface IpGroupState {
    /**
     * The description.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the IP group.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.workspaces.IpGroupRule>[]>;
    /**
     * A map of tags assigned to the WorkSpaces directory. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a IpGroup resource.
 */
export interface IpGroupArgs {
    /**
     * The description.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the IP group.
     */
    name?: pulumi.Input<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.workspaces.IpGroupRule>[]>;
    /**
     * A map of tags assigned to the WorkSpaces directory. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
