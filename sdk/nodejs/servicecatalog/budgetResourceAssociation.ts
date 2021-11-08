// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Service Catalog Budget Resource Association.
 *
 * > **Tip:** A "resource" is either a Service Catalog portfolio or product.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicecatalog.BudgetResourceAssociation("example", {
 *     budgetName: "budget-pjtvyakdlyo3m",
 *     resourceId: "prod-dnigbtea24ste",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_servicecatalog_budget_resource_association` can be imported using the budget name and resource ID, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation example budget-pjtvyakdlyo3m:prod-dnigbtea24ste
 * ```
 */
export class BudgetResourceAssociation extends pulumi.CustomResource {
    /**
     * Get an existing BudgetResourceAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BudgetResourceAssociationState, opts?: pulumi.CustomResourceOptions): BudgetResourceAssociation {
        return new BudgetResourceAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicecatalog/budgetResourceAssociation:BudgetResourceAssociation';

    /**
     * Returns true if the given object is an instance of BudgetResourceAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BudgetResourceAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BudgetResourceAssociation.__pulumiType;
    }

    /**
     * Budget name.
     */
    public readonly budgetName!: pulumi.Output<string>;
    /**
     * Resource identifier.
     */
    public readonly resourceId!: pulumi.Output<string>;

    /**
     * Create a BudgetResourceAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BudgetResourceAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BudgetResourceAssociationArgs | BudgetResourceAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BudgetResourceAssociationState | undefined;
            inputs["budgetName"] = state ? state.budgetName : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
        } else {
            const args = argsOrState as BudgetResourceAssociationArgs | undefined;
            if ((!args || args.budgetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'budgetName'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            inputs["budgetName"] = args ? args.budgetName : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BudgetResourceAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BudgetResourceAssociation resources.
 */
export interface BudgetResourceAssociationState {
    /**
     * Budget name.
     */
    budgetName?: pulumi.Input<string>;
    /**
     * Resource identifier.
     */
    resourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BudgetResourceAssociation resource.
 */
export interface BudgetResourceAssociationArgs {
    /**
     * Budget name.
     */
    budgetName: pulumi.Input<string>;
    /**
     * Resource identifier.
     */
    resourceId: pulumi.Input<string>;
}
