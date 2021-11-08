// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {RestApi} from "./index";

/**
 * Manages an API Gateway Request Validator.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigateway.RequestValidator("example", {
 *     restApi: aws_api_gateway_rest_api.example.id,
 *     validateRequestBody: true,
 *     validateRequestParameters: true,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_api_gateway_request_validator` can be imported using `REST-API-ID/REQUEST-VALIDATOR-ID`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:apigateway/requestValidator:RequestValidator example 12345abcde/67890fghij
 * ```
 */
export class RequestValidator extends pulumi.CustomResource {
    /**
     * Get an existing RequestValidator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RequestValidatorState, opts?: pulumi.CustomResourceOptions): RequestValidator {
        return new RequestValidator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigateway/requestValidator:RequestValidator';

    /**
     * Returns true if the given object is an instance of RequestValidator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RequestValidator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RequestValidator.__pulumiType;
    }

    /**
     * The name of the request validator
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the associated Rest API
     */
    public readonly restApi!: pulumi.Output<string>;
    /**
     * Boolean whether to validate request body. Defaults to `false`.
     */
    public readonly validateRequestBody!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean whether to validate request parameters. Defaults to `false`.
     */
    public readonly validateRequestParameters!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RequestValidator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RequestValidatorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RequestValidatorArgs | RequestValidatorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RequestValidatorState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["restApi"] = state ? state.restApi : undefined;
            inputs["validateRequestBody"] = state ? state.validateRequestBody : undefined;
            inputs["validateRequestParameters"] = state ? state.validateRequestParameters : undefined;
        } else {
            const args = argsOrState as RequestValidatorArgs | undefined;
            if ((!args || args.restApi === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restApi'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["restApi"] = args ? args.restApi : undefined;
            inputs["validateRequestBody"] = args ? args.validateRequestBody : undefined;
            inputs["validateRequestParameters"] = args ? args.validateRequestParameters : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RequestValidator.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RequestValidator resources.
 */
export interface RequestValidatorState {
    /**
     * The name of the request validator
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the associated Rest API
     */
    restApi?: pulumi.Input<string | RestApi>;
    /**
     * Boolean whether to validate request body. Defaults to `false`.
     */
    validateRequestBody?: pulumi.Input<boolean>;
    /**
     * Boolean whether to validate request parameters. Defaults to `false`.
     */
    validateRequestParameters?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RequestValidator resource.
 */
export interface RequestValidatorArgs {
    /**
     * The name of the request validator
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the associated Rest API
     */
    restApi: pulumi.Input<string | RestApi>;
    /**
     * Boolean whether to validate request body. Defaults to `false`.
     */
    validateRequestBody?: pulumi.Input<boolean>;
    /**
     * Boolean whether to validate request parameters. Defaults to `false`.
     */
    validateRequestParameters?: pulumi.Input<boolean>;
}
