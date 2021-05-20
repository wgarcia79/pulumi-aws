// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an NFS Location within AWS DataSync.
 *
 * > **NOTE:** The DataSync Agents must be available before creating this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.datasync.NfsLocation("example", {
 *     serverHostname: "nfs.example.com",
 *     subdirectory: "/exported/path",
 *     onPremConfig: {
 *         agentArns: [aws_datasync_agent.example.arn],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_datasync_location_nfs` can be imported by using the DataSync Task Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:datasync/nfsLocation:NfsLocation example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
 * ```
 */
export class NfsLocation extends pulumi.CustomResource {
    /**
     * Get an existing NfsLocation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NfsLocationState, opts?: pulumi.CustomResourceOptions): NfsLocation {
        return new NfsLocation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:datasync/nfsLocation:NfsLocation';

    /**
     * Returns true if the given object is an instance of NfsLocation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NfsLocation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NfsLocation.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block containing information for connecting to the NFS File System.
     */
    public readonly onPremConfig!: pulumi.Output<outputs.datasync.NfsLocationOnPremConfig>;
    /**
     * Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
     */
    public readonly serverHostname!: pulumi.Output<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    public readonly subdirectory!: pulumi.Output<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly uri!: pulumi.Output<string>;

    /**
     * Create a NfsLocation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NfsLocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NfsLocationArgs | NfsLocationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NfsLocationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["onPremConfig"] = state ? state.onPremConfig : undefined;
            inputs["serverHostname"] = state ? state.serverHostname : undefined;
            inputs["subdirectory"] = state ? state.subdirectory : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as NfsLocationArgs | undefined;
            if ((!args || args.onPremConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'onPremConfig'");
            }
            if ((!args || args.serverHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverHostname'");
            }
            if ((!args || args.subdirectory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subdirectory'");
            }
            inputs["onPremConfig"] = args ? args.onPremConfig : undefined;
            inputs["serverHostname"] = args ? args.serverHostname : undefined;
            inputs["subdirectory"] = args ? args.subdirectory : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NfsLocation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NfsLocation resources.
 */
export interface NfsLocationState {
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Configuration block containing information for connecting to the NFS File System.
     */
    readonly onPremConfig?: pulumi.Input<inputs.datasync.NfsLocationOnPremConfig>;
    /**
     * Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
     */
    readonly serverHostname?: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    readonly subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NfsLocation resource.
 */
export interface NfsLocationArgs {
    /**
     * Configuration block containing information for connecting to the NFS File System.
     */
    readonly onPremConfig: pulumi.Input<inputs.datasync.NfsLocationOnPremConfig>;
    /**
     * Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
     */
    readonly serverHostname: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    readonly subdirectory: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    readonly tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
