// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway working storage.
 *
 * > **NOTE:** The Storage Gateway API provides no method to remove a working storage disk. Destroying this resource does not perform any Storage Gateway actions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.storagegateway.WorkingStorage("example", {
 *     diskId: data.aws_storagegateway_local_disk.example.id,
 *     gatewayArn: aws_storagegateway_gateway.example.arn,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_storagegateway_working_storage` can be imported by using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:storagegateway/workingStorage:WorkingStorage example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
 * ```
 */
export class WorkingStorage extends pulumi.CustomResource {
    /**
     * Get an existing WorkingStorage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkingStorageState, opts?: pulumi.CustomResourceOptions): WorkingStorage {
        return new WorkingStorage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:storagegateway/workingStorage:WorkingStorage';

    /**
     * Returns true if the given object is an instance of WorkingStorage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkingStorage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkingStorage.__pulumiType;
    }

    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    public readonly gatewayArn!: pulumi.Output<string>;

    /**
     * Create a WorkingStorage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkingStorageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkingStorageArgs | WorkingStorageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkingStorageState | undefined;
            inputs["diskId"] = state ? state.diskId : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
        } else {
            const args = argsOrState as WorkingStorageArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.gatewayArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            inputs["diskId"] = args ? args.diskId : undefined;
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WorkingStorage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkingStorage resources.
 */
export interface WorkingStorageState {
    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    diskId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    gatewayArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkingStorage resource.
 */
export interface WorkingStorageArgs {
    /**
     * Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
     */
    diskId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    gatewayArn: pulumi.Input<string>;
}
