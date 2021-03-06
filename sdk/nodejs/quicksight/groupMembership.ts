// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing QuickSight Group Membership
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.quicksight.GroupMembership("example", {
 *     groupName: "all-access-users",
 *     memberName: "john_smith",
 * });
 * ```
 *
 * ## Import
 *
 * QuickSight Group membership can be imported using the AWS account ID, namespace, group name and member name separated by `/`.
 *
 * ```sh
 *  $ pulumi import aws:quicksight/groupMembership:GroupMembership example 123456789123/default/all-access-users/john_smith
 * ```
 */
export class GroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipState, opts?: pulumi.CustomResourceOptions): GroupMembership {
        return new GroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:quicksight/groupMembership:GroupMembership';

    /**
     * Returns true if the given object is an instance of GroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembership.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * The name of the group in which the member will be added.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The name of the member to add to the group.
     */
    public readonly memberName!: pulumi.Output<string>;
    /**
     * The namespace. Defaults to `default`. Currently only `default` is supported.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembershipArgs | GroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupMembershipState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsAccountId"] = state ? state.awsAccountId : undefined;
            inputs["groupName"] = state ? state.groupName : undefined;
            inputs["memberName"] = state ? state.memberName : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
        } else {
            const args = argsOrState as GroupMembershipArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.memberName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberName'");
            }
            inputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["memberName"] = args ? args.memberName : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupMembership.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembership resources.
 */
export interface GroupMembershipState {
    arn?: pulumi.Input<string>;
    /**
     * The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The name of the group in which the member will be added.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The name of the member to add to the group.
     */
    memberName?: pulumi.Input<string>;
    /**
     * The namespace. Defaults to `default`. Currently only `default` is supported.
     */
    namespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupMembership resource.
 */
export interface GroupMembershipArgs {
    /**
     * The ID for the AWS account that the group is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
     */
    awsAccountId?: pulumi.Input<string>;
    /**
     * The name of the group in which the member will be added.
     */
    groupName: pulumi.Input<string>;
    /**
     * The name of the member to add to the group.
     */
    memberName: pulumi.Input<string>;
    /**
     * The namespace. Defaults to `default`. Currently only `default` is supported.
     */
    namespace?: pulumi.Input<string>;
}
