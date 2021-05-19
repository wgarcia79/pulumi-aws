// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an [AWS Organizations Delegated Administrator](https://docs.aws.amazon.com/organizations/latest/APIReference/API_RegisterDelegatedAdministrator.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.organizations.DelegatedAdministrator("example", {
 *     accountId: "AWS ACCOUNT ID",
 *     servicePrincipal: "Service principal",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_organizations_delegated_administrator` can be imported by using the account ID and its service principal, e.g.
 *
 * ```sh
 *  $ pulumi import aws:organizations/delegatedAdministrator:DelegatedAdministrator example 123456789012/config.amazonaws.com
 * ```
 */
export class DelegatedAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing DelegatedAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DelegatedAdministratorState, opts?: pulumi.CustomResourceOptions): DelegatedAdministrator {
        return new DelegatedAdministrator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:organizations/delegatedAdministrator:DelegatedAdministrator';

    /**
     * Returns true if the given object is an instance of DelegatedAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DelegatedAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DelegatedAdministrator.__pulumiType;
    }

    /**
     * The account ID number of the member account in the organization to register as a delegated administrator.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the delegated administrator's account.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date when the account was made a delegated administrator.
     */
    public /*out*/ readonly delegationEnabledDate!: pulumi.Output<string>;
    /**
     * The email address that is associated with the delegated administrator's AWS account.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The method by which the delegated administrator's account joined the organization.
     */
    public /*out*/ readonly joinedMethod!: pulumi.Output<string>;
    /**
     * The date when the delegated administrator's account became a part of the organization.
     */
    public /*out*/ readonly joinedTimestamp!: pulumi.Output<string>;
    /**
     * The friendly name of the delegated administrator's account.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The service principal of the AWS service for which you want to make the member account a delegated administrator.
     */
    public readonly servicePrincipal!: pulumi.Output<string>;
    /**
     * The status of the delegated administrator's account in the organization.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a DelegatedAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DelegatedAdministratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DelegatedAdministratorArgs | DelegatedAdministratorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DelegatedAdministratorState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["delegationEnabledDate"] = state ? state.delegationEnabledDate : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["joinedMethod"] = state ? state.joinedMethod : undefined;
            inputs["joinedTimestamp"] = state ? state.joinedTimestamp : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["servicePrincipal"] = state ? state.servicePrincipal : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as DelegatedAdministratorArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.servicePrincipal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipal'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["servicePrincipal"] = args ? args.servicePrincipal : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["delegationEnabledDate"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
            inputs["joinedMethod"] = undefined /*out*/;
            inputs["joinedTimestamp"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DelegatedAdministrator.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DelegatedAdministrator resources.
 */
export interface DelegatedAdministratorState {
    /**
     * The account ID number of the member account in the organization to register as a delegated administrator.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the delegated administrator's account.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The date when the account was made a delegated administrator.
     */
    readonly delegationEnabledDate?: pulumi.Input<string>;
    /**
     * The email address that is associated with the delegated administrator's AWS account.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * The method by which the delegated administrator's account joined the organization.
     */
    readonly joinedMethod?: pulumi.Input<string>;
    /**
     * The date when the delegated administrator's account became a part of the organization.
     */
    readonly joinedTimestamp?: pulumi.Input<string>;
    /**
     * The friendly name of the delegated administrator's account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The service principal of the AWS service for which you want to make the member account a delegated administrator.
     */
    readonly servicePrincipal?: pulumi.Input<string>;
    /**
     * The status of the delegated administrator's account in the organization.
     */
    readonly status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DelegatedAdministrator resource.
 */
export interface DelegatedAdministratorArgs {
    /**
     * The account ID number of the member account in the organization to register as a delegated administrator.
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * The service principal of the AWS service for which you want to make the member account a delegated administrator.
     */
    readonly servicePrincipal: pulumi.Input<string>;
}
