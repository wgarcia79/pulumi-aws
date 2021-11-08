// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Service Discovery Instance resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
 *     cidrBlock: "10.0.0.0/16",
 *     enableDnsSupport: true,
 *     enableDnsHostnames: true,
 * });
 * const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", {
 *     description: "example",
 *     vpc: exampleVpc.id,
 * });
 * const exampleService = new aws.servicediscovery.Service("exampleService", {
 *     dnsConfig: {
 *         namespaceId: examplePrivateDnsNamespace.id,
 *         dnsRecords: [{
 *             ttl: 10,
 *             type: "A",
 *         }],
 *         routingPolicy: "MULTIVALUE",
 *     },
 *     healthCheckCustomConfig: {
 *         failureThreshold: 1,
 *     },
 * });
 * const exampleInstance = new aws.servicediscovery.Instance("exampleInstance", {
 *     instanceId: "example-instance-id",
 *     serviceId: exampleService.id,
 *     attributes: {
 *         AWS_INSTANCE_IPV4: "172.18.0.1",
 *         custom_attribute: "custom",
 *     },
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleHttpNamespace = new aws.servicediscovery.HttpNamespace("exampleHttpNamespace", {description: "example"});
 * const exampleService = new aws.servicediscovery.Service("exampleService", {namespaceId: exampleHttpNamespace.id});
 * const exampleInstance = new aws.servicediscovery.Instance("exampleInstance", {
 *     instanceId: "example-instance-id",
 *     serviceId: exampleService.id,
 *     attributes: {
 *         AWS_EC2_INSTANCE_ID: "i-0abdg374kd892cj6dl",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Service Discovery Instance can be imported using the service ID and instance ID, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:servicediscovery/instance:Instance example 0123456789/i-0123
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicediscovery/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * A map contains the attributes of the instance. Check the [doc](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html#API_RegisterInstance_RequestSyntax) for the supported attributes and syntax.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: string}>;
    /**
     * The ID of the service instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the service that you want to use to create the instance.
     */
    public readonly serviceId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.attributes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attributes'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * A map contains the attributes of the instance. Check the [doc](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html#API_RegisterInstance_RequestSyntax) for the supported attributes and syntax.
     */
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the service instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The ID of the service that you want to use to create the instance.
     */
    serviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * A map contains the attributes of the instance. Check the [doc](https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html#API_RegisterInstance_RequestSyntax) for the supported attributes and syntax.
     */
    attributes: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the service instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The ID of the service that you want to use to create the instance.
     */
    serviceId: pulumi.Input<string>;
}
