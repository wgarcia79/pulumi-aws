// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Imports a disk image from S3 as a Snapshot.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ebs.SnapshotImport("example", {
 *     diskContainer: {
 *         format: "VHD",
 *         userBucket: {
 *             s3Bucket: "disk-images",
 *             s3Key: "source.vhd",
 *         },
 *     },
 *     roleName: "disk-image-import",
 *     tags: {
 *         Name: "HelloWorld",
 *     },
 * });
 * ```
 */
export class SnapshotImport extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotImport resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotImportState, opts?: pulumi.CustomResourceOptions): SnapshotImport {
        return new SnapshotImport(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ebs/snapshotImport:SnapshotImport';

    /**
     * Returns true if the given object is an instance of SnapshotImport.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotImport {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotImport.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the EBS Snapshot.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The client-specific data. Detailed below.
     */
    public readonly clientData!: pulumi.Output<outputs.ebs.SnapshotImportClientData | undefined>;
    /**
     * The data encryption key identifier for the snapshot.
     */
    public /*out*/ readonly dataEncryptionKeyId!: pulumi.Output<string>;
    /**
     * The description of the disk image being imported.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Information about the disk container. Detailed below.
     */
    public readonly diskContainer!: pulumi.Output<outputs.ebs.SnapshotImportDiskContainer>;
    /**
     * Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
     */
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    /**
     * An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     */
    public /*out*/ readonly ownerAlias!: pulumi.Output<string>;
    /**
     * The AWS account ID of the EBS snapshot owner.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
     */
    public readonly roleName!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the snapshot.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The size of the drive in GiBs.
     */
    public /*out*/ readonly volumeSize!: pulumi.Output<number>;

    /**
     * Create a SnapshotImport resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotImportArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotImportArgs | SnapshotImportState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotImportState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["clientData"] = state ? state.clientData : undefined;
            inputs["dataEncryptionKeyId"] = state ? state.dataEncryptionKeyId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskContainer"] = state ? state.diskContainer : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["ownerAlias"] = state ? state.ownerAlias : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["volumeSize"] = state ? state.volumeSize : undefined;
        } else {
            const args = argsOrState as SnapshotImportArgs | undefined;
            if ((!args || args.diskContainer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskContainer'");
            }
            inputs["clientData"] = args ? args.clientData : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskContainer"] = args ? args.diskContainer : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dataEncryptionKeyId"] = undefined /*out*/;
            inputs["ownerAlias"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
            inputs["volumeSize"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SnapshotImport.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotImport resources.
 */
export interface SnapshotImportState {
    /**
     * Amazon Resource Name (ARN) of the EBS Snapshot.
     */
    arn?: pulumi.Input<string>;
    /**
     * The client-specific data. Detailed below.
     */
    clientData?: pulumi.Input<inputs.ebs.SnapshotImportClientData>;
    /**
     * The data encryption key identifier for the snapshot.
     */
    dataEncryptionKeyId?: pulumi.Input<string>;
    /**
     * The description of the disk image being imported.
     */
    description?: pulumi.Input<string>;
    /**
     * Information about the disk container. Detailed below.
     */
    diskContainer?: pulumi.Input<inputs.ebs.SnapshotImportDiskContainer>;
    /**
     * Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     */
    ownerAlias?: pulumi.Input<string>;
    /**
     * The AWS account ID of the EBS snapshot owner.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
     */
    roleName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the snapshot.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The size of the drive in GiBs.
     */
    volumeSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SnapshotImport resource.
 */
export interface SnapshotImportArgs {
    /**
     * The client-specific data. Detailed below.
     */
    clientData?: pulumi.Input<inputs.ebs.SnapshotImportClientData>;
    /**
     * The description of the disk image being imported.
     */
    description?: pulumi.Input<string>;
    /**
     * Information about the disk container. Detailed below.
     */
    diskContainer: pulumi.Input<inputs.ebs.SnapshotImportDiskContainer>;
    /**
     * Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
     */
    encrypted?: pulumi.Input<boolean>;
    /**
     * An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
     */
    roleName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the snapshot.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
