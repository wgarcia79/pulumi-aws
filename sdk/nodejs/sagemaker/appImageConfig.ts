// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Sagemaker App Image Config resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.sagemaker.AppImageConfig("test", {
 *     appImageConfigName: "example",
 *     kernelGatewayImageConfig: {
 *         kernelSpec: {
 *             name: "example",
 *         },
 *     },
 * });
 * ```
 * ### Default File System Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.sagemaker.AppImageConfig("test", {
 *     appImageConfigName: "example",
 *     kernelGatewayImageConfig: {
 *         fileSystemConfig: {},
 *         kernelSpec: {
 *             name: "example",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Sagemaker App Image Configs can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sagemaker/appImageConfig:AppImageConfig example example
 * ```
 */
export class AppImageConfig extends pulumi.CustomResource {
    /**
     * Get an existing AppImageConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppImageConfigState, opts?: pulumi.CustomResourceOptions): AppImageConfig {
        return new AppImageConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/appImageConfig:AppImageConfig';

    /**
     * Returns true if the given object is an instance of AppImageConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppImageConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppImageConfig.__pulumiType;
    }

    /**
     * The name of the App Image Config.
     */
    public readonly appImageConfigName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
     */
    public readonly kernelGatewayImageConfig!: pulumi.Output<outputs.sagemaker.AppImageConfigKernelGatewayImageConfig | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a AppImageConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppImageConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppImageConfigArgs | AppImageConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppImageConfigState | undefined;
            inputs["appImageConfigName"] = state ? state.appImageConfigName : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["kernelGatewayImageConfig"] = state ? state.kernelGatewayImageConfig : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AppImageConfigArgs | undefined;
            if ((!args || args.appImageConfigName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appImageConfigName'");
            }
            inputs["appImageConfigName"] = args ? args.appImageConfigName : undefined;
            inputs["kernelGatewayImageConfig"] = args ? args.kernelGatewayImageConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AppImageConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppImageConfig resources.
 */
export interface AppImageConfigState {
    /**
     * The name of the App Image Config.
     */
    appImageConfigName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this App Image Config.
     */
    arn?: pulumi.Input<string>;
    /**
     * The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
     */
    kernelGatewayImageConfig?: pulumi.Input<inputs.sagemaker.AppImageConfigKernelGatewayImageConfig>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AppImageConfig resource.
 */
export interface AppImageConfigArgs {
    /**
     * The name of the App Image Config.
     */
    appImageConfigName: pulumi.Input<string>;
    /**
     * The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. See Kernel Gateway Image Config details below.
     */
    kernelGatewayImageConfig?: pulumi.Input<inputs.sagemaker.AppImageConfigKernelGatewayImageConfig>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
