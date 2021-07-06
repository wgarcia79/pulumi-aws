// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information on Service Catalog Portfolio Constraints.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.servicecatalog.getPortfolioConstraints({
 *     portfolioId: "port-3lli3b3an",
 * }));
 * ```
 */
export function getPortfolioConstraints(args: GetPortfolioConstraintsArgs, opts?: pulumi.InvokeOptions): Promise<GetPortfolioConstraintsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:servicecatalog/getPortfolioConstraints:getPortfolioConstraints", {
        "acceptLanguage": args.acceptLanguage,
        "portfolioId": args.portfolioId,
        "productId": args.productId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortfolioConstraints.
 */
export interface GetPortfolioConstraintsArgs {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    acceptLanguage?: string;
    /**
     * Portfolio identifier.
     */
    portfolioId: string;
    /**
     * Product identifier.
     */
    productId?: string;
}

/**
 * A collection of values returned by getPortfolioConstraints.
 */
export interface GetPortfolioConstraintsResult {
    readonly acceptLanguage?: string;
    /**
     * List of information about the constraints. See details below.
     */
    readonly details: outputs.servicecatalog.GetPortfolioConstraintsDetail[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Identifier of the portfolio the product resides in. The constraint applies only to the instance of the product that lives within this portfolio.
     */
    readonly portfolioId: string;
    /**
     * Identifier of the product the constraint applies to. A constraint applies to a specific instance of a product within a certain portfolio.
     */
    readonly productId?: string;
}
