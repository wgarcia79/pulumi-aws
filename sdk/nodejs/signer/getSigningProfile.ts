// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a Signer Signing Profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const productionSigningProfile = pulumi.output(aws.signer.getSigningProfile({
 *     name: "prod_profile_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK",
 * }, { async: true }));
 * ```
 */
export function getSigningProfile(args: GetSigningProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetSigningProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:signer/getSigningProfile:getSigningProfile", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getSigningProfile.
 */
export interface GetSigningProfileArgs {
    /**
     * The name of the target signing profile.
     */
    readonly name: string;
    /**
     * A list of tags associated with the signing profile.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getSigningProfile.
 */
export interface GetSigningProfileResult {
    /**
     * The Amazon Resource Name (ARN) for the signing profile.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * A human-readable name for the signing platform associated with the signing profile.
     */
    readonly platformDisplayName: string;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    readonly platformId: string;
    /**
     * Revocation information for a signing profile.
     */
    readonly revocationRecords: outputs.signer.GetSigningProfileRevocationRecord[];
    /**
     * The validity period for a signing job.
     */
    readonly signatureValidityPeriods: outputs.signer.GetSigningProfileSignatureValidityPeriod[];
    /**
     * The status of the target signing profile.
     */
    readonly status: string;
    /**
     * A list of tags associated with the signing profile.
     */
    readonly tags: {[key: string]: string};
    /**
     * The current version of the signing profile.
     */
    readonly version: string;
    /**
     * The signing profile ARN, including the profile version.
     */
    readonly versionArn: string;
}
