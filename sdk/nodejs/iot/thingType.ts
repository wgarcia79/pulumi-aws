// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates and manages an AWS IoT Thing Type.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.iot.ThingType("foo", {});
 * ```
 *
 * ## Import
 *
 * IOT Thing Types can be imported using the name, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:iot/thingType:ThingType example example
 * ```
 */
export class ThingType extends pulumi.CustomResource {
    /**
     * Get an existing ThingType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ThingTypeState, opts?: pulumi.CustomResourceOptions): ThingType {
        return new ThingType(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iot/thingType:ThingType';

    /**
     * Returns true if the given object is an instance of ThingType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ThingType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ThingType.__pulumiType;
    }

    /**
     * The ARN of the created AWS IoT Thing Type.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether the thing type is deprecated. If true, no new things could be associated with this type.
     */
    public readonly deprecated!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the thing type.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * , Configuration block that can contain the following properties of the thing type:
     */
    public readonly properties!: pulumi.Output<outputs.iot.ThingTypeProperties | undefined>;

    /**
     * Create a ThingType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ThingTypeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ThingTypeArgs | ThingTypeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ThingTypeState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["deprecated"] = state ? state.deprecated : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
        } else {
            const args = argsOrState as ThingTypeArgs | undefined;
            inputs["deprecated"] = args ? args.deprecated : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ThingType.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ThingType resources.
 */
export interface ThingTypeState {
    /**
     * The ARN of the created AWS IoT Thing Type.
     */
    arn?: pulumi.Input<string>;
    /**
     * Whether the thing type is deprecated. If true, no new things could be associated with this type.
     */
    deprecated?: pulumi.Input<boolean>;
    /**
     * The name of the thing type.
     */
    name?: pulumi.Input<string>;
    /**
     * , Configuration block that can contain the following properties of the thing type:
     */
    properties?: pulumi.Input<inputs.iot.ThingTypeProperties>;
}

/**
 * The set of arguments for constructing a ThingType resource.
 */
export interface ThingTypeArgs {
    /**
     * Whether the thing type is deprecated. If true, no new things could be associated with this type.
     */
    deprecated?: pulumi.Input<boolean>;
    /**
     * The name of the thing type.
     */
    name?: pulumi.Input<string>;
    /**
     * , Configuration block that can contain the following properties of the thing type:
     */
    properties?: pulumi.Input<inputs.iot.ThingTypeProperties>;
}
