// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a member account in the current organization.
 * 
 * > **Note:** Account management must be done from the organization's master account.
 * 
 * !> **WARNING:** Deleting this Terraform resource will only remove an AWS account from an organization. Terraform will not close the account. The member account must be prepared to be a standalone account beforehand. See the [AWS Organizations documentation](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html) for more information.
 * 
 * ## Example Usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_organizations_account_account = new aws.organizations.Account("account", {
 *     email: "john@doe.org",
 *     name: "my_new_account",
 * });
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN for this account.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
     */
    public readonly email: pulumi.Output<string>;
    /**
     * If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
     */
    public readonly iamUserAccessToBilling: pulumi.Output<string | undefined>;
    public /*out*/ readonly joinedMethod: pulumi.Output<string>;
    public /*out*/ readonly joinedTimestamp: pulumi.Output<string>;
    /**
     * A friendly name for the member account.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account.
     */
    public readonly roleName: pulumi.Output<string | undefined>;
    public /*out*/ readonly status: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AccountState = argsOrState as AccountState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["iamUserAccessToBilling"] = state ? state.iamUserAccessToBilling : undefined;
            inputs["joinedMethod"] = state ? state.joinedMethod : undefined;
            inputs["joinedTimestamp"] = state ? state.joinedTimestamp : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if (!args || args.email === undefined) {
                throw new Error("Missing required property 'email'");
            }
            inputs["email"] = args ? args.email : undefined;
            inputs["iamUserAccessToBilling"] = args ? args.iamUserAccessToBilling : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["joinedMethod"] = undefined /*out*/;
            inputs["joinedTimestamp"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        super("aws:organizations/account:Account", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * The ARN for this account.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
     */
    readonly iamUserAccessToBilling?: pulumi.Input<string>;
    readonly joinedMethod?: pulumi.Input<string>;
    readonly joinedTimestamp?: pulumi.Input<string>;
    /**
     * A friendly name for the member account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account.
     */
    readonly roleName?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
     */
    readonly email: pulumi.Input<string>;
    /**
     * If set to `ALLOW`, the new account enables IAM users to access account billing information if they have the required permissions. If set to `DENY`, then only the root user of the new account can access account billing information.
     */
    readonly iamUserAccessToBilling?: pulumi.Input<string>;
    /**
     * A friendly name for the member account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the master account, allowing users in the master account to assume the role, as permitted by the master account administrator. The role has administrator permissions in the new member account.
     */
    readonly roleName?: pulumi.Input<string>;
}
