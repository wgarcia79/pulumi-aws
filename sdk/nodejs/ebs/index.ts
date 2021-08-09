// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./defaultKmsKey";
export * from "./encryptionByDefault";
export * from "./getDefaultKmsKey";
export * from "./getEbsVolumes";
export * from "./getEncryptionByDefault";
export * from "./getSnapshot";
export * from "./getSnapshotIds";
export * from "./getVolume";
export * from "./snapshot";
export * from "./snapshotCopy";
export * from "./snapshotImport";
export * from "./volume";

// Import resources to register:
import { DefaultKmsKey } from "./defaultKmsKey";
import { EncryptionByDefault } from "./encryptionByDefault";
import { Snapshot } from "./snapshot";
import { SnapshotCopy } from "./snapshotCopy";
import { SnapshotImport } from "./snapshotImport";
import { Volume } from "./volume";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:ebs/defaultKmsKey:DefaultKmsKey":
                return new DefaultKmsKey(name, <any>undefined, { urn })
            case "aws:ebs/encryptionByDefault:EncryptionByDefault":
                return new EncryptionByDefault(name, <any>undefined, { urn })
            case "aws:ebs/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "aws:ebs/snapshotCopy:SnapshotCopy":
                return new SnapshotCopy(name, <any>undefined, { urn })
            case "aws:ebs/snapshotImport:SnapshotImport":
                return new SnapshotImport(name, <any>undefined, { urn })
            case "aws:ebs/volume:Volume":
                return new Volume(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "ebs/defaultKmsKey", _module)
pulumi.runtime.registerResourceModule("aws", "ebs/encryptionByDefault", _module)
pulumi.runtime.registerResourceModule("aws", "ebs/snapshot", _module)
pulumi.runtime.registerResourceModule("aws", "ebs/snapshotCopy", _module)
pulumi.runtime.registerResourceModule("aws", "ebs/snapshotImport", _module)
pulumi.runtime.registerResourceModule("aws", "ebs/volume", _module)
