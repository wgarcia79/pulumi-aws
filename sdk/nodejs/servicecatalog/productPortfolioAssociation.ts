// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Service Catalog Product Portfolio Association.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicecatalog.ProductPortfolioAssociation("example", {
 *     portfolioId: "port-68656c6c6f",
 *     productId: "prod-dnigbtea24ste",
 * });
 * ```
 *
 * ## Import
 *
 * `aws_servicecatalog_product_portfolio_association` can be imported using the accept language, portfolio ID, and product ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:servicecatalog/productPortfolioAssociation:ProductPortfolioAssociation example en:port-68656c6c6f:prod-dnigbtea24ste
 * ```
 */
export class ProductPortfolioAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ProductPortfolioAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductPortfolioAssociationState, opts?: pulumi.CustomResourceOptions): ProductPortfolioAssociation {
        return new ProductPortfolioAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicecatalog/productPortfolioAssociation:ProductPortfolioAssociation';

    /**
     * Returns true if the given object is an instance of ProductPortfolioAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductPortfolioAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductPortfolioAssociation.__pulumiType;
    }

    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    /**
     * Portfolio identifier.
     */
    public readonly portfolioId!: pulumi.Output<string>;
    /**
     * Product identifier.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Identifier of the source portfolio.
     */
    public readonly sourcePortfolioId!: pulumi.Output<string | undefined>;

    /**
     * Create a ProductPortfolioAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductPortfolioAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductPortfolioAssociationArgs | ProductPortfolioAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProductPortfolioAssociationState | undefined;
            inputs["acceptLanguage"] = state ? state.acceptLanguage : undefined;
            inputs["portfolioId"] = state ? state.portfolioId : undefined;
            inputs["productId"] = state ? state.productId : undefined;
            inputs["sourcePortfolioId"] = state ? state.sourcePortfolioId : undefined;
        } else {
            const args = argsOrState as ProductPortfolioAssociationArgs | undefined;
            if ((!args || args.portfolioId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portfolioId'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            inputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            inputs["portfolioId"] = args ? args.portfolioId : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["sourcePortfolioId"] = args ? args.sourcePortfolioId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProductPortfolioAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductPortfolioAssociation resources.
 */
export interface ProductPortfolioAssociationState {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    readonly acceptLanguage?: pulumi.Input<string>;
    /**
     * Portfolio identifier.
     */
    readonly portfolioId?: pulumi.Input<string>;
    /**
     * Product identifier.
     */
    readonly productId?: pulumi.Input<string>;
    /**
     * Identifier of the source portfolio.
     */
    readonly sourcePortfolioId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductPortfolioAssociation resource.
 */
export interface ProductPortfolioAssociationArgs {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    readonly acceptLanguage?: pulumi.Input<string>;
    /**
     * Portfolio identifier.
     */
    readonly portfolioId: pulumi.Input<string>;
    /**
     * Product identifier.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * Identifier of the source portfolio.
     */
    readonly sourcePortfolioId?: pulumi.Input<string>;
}
