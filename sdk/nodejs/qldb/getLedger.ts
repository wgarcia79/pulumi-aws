// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to fetch information about a Quantum Ledger Database.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.qldb.getLedger({
 *     name: "an_example_ledger",
 * }));
 * ```
 */
export function getLedger(args: GetLedgerArgs, opts?: pulumi.InvokeOptions): Promise<GetLedgerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:qldb/getLedger:getLedger", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLedger.
 */
export interface GetLedgerArgs {
    /**
     * The friendly name of the ledger to match.
     */
    name: string;
}

/**
 * A collection of values returned by getLedger.
 */
export interface GetLedgerResult {
    readonly arn: string;
    readonly deletionProtection: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly permissionsMode: string;
}
