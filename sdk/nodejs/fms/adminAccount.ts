// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to associate/disassociate an AWS Firewall Manager administrator account. This operation must be performed in the `us-east-1` region.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.fms.AdminAccount("example", {});
 * ```
 *
 * ## Import
 *
 * Firewall Manager administrator account association can be imported using the account ID, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:fms/adminAccount:AdminAccount example 123456789012
 * ```
 */
export class AdminAccount extends pulumi.CustomResource {
    /**
     * Get an existing AdminAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminAccountState, opts?: pulumi.CustomResourceOptions): AdminAccount {
        return new AdminAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fms/adminAccount:AdminAccount';

    /**
     * Returns true if the given object is an instance of AdminAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdminAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdminAccount.__pulumiType;
    }

    /**
     * The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
     */
    public readonly accountId!: pulumi.Output<string>;

    /**
     * Create a AdminAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AdminAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminAccountArgs | AdminAccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdminAccountState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
        } else {
            const args = argsOrState as AdminAccountArgs | undefined;
            inputs["accountId"] = args ? args.accountId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AdminAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminAccount resources.
 */
export interface AdminAccountState {
    /**
     * The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
     */
    accountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdminAccount resource.
 */
export interface AdminAccountArgs {
    /**
     * The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
     */
    accountId?: pulumi.Input<string>;
}
