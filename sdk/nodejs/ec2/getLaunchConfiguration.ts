// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides information about a Launch Configuration.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_launch_configuration_ubuntu = pulumi.output(aws.ec2.getLaunchConfiguration({
 *     name: "test-launch-config",
 * }));
 * ```
 */
export function getLaunchConfiguration(args: GetLaunchConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchConfigurationResult> {
    return pulumi.runtime.invoke("aws:ec2/getLaunchConfiguration:getLaunchConfiguration", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLaunchConfiguration.
 */
export interface GetLaunchConfigurationArgs {
    /**
     * The name of the launch configuration.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getLaunchConfiguration.
 */
export interface GetLaunchConfigurationResult {
    /**
     * Whether a Public IP address is associated with the instance.
     */
    readonly associatePublicIpAddress: boolean;
    /**
     * The EBS Block Devices attached to the instance.
     */
    readonly ebsBlockDevices: { deleteOnTermination: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeSize: number, volumeType: string }[];
    /**
     * Whether the launched EC2 instance will be EBS-optimized.
     */
    readonly ebsOptimized: boolean;
    /**
     * Whether Detailed Monitoring is Enabled.
     */
    readonly enableMonitoring: boolean;
    /**
     * The Ephemeral volumes on the instance.
     */
    readonly ephemeralBlockDevices: { deviceName: string, virtualName: string }[];
    /**
     * The IAM Instance Profile to associate with launched instances.
     */
    readonly iamInstanceProfile: string;
    /**
     * The EC2 Image ID of the instance.
     */
    readonly imageId: string;
    /**
     * The Instance Type of the instance to launch.
     */
    readonly instanceType: string;
    /**
     * The Key Name that should be used for the instance.
     */
    readonly keyName: string;
    /**
     * The Tenancy of the instance.
     */
    readonly placementTenancy: string;
    /**
     * The Root Block Device of the instance.
     */
    readonly rootBlockDevices: { deleteOnTermination: boolean, iops: number, volumeSize: number, volumeType: string }[];
    /**
     * A list of associated Security Group IDS.
     */
    readonly securityGroups: string[];
    /**
     * The Price to use for reserving Spot instances.
     */
    readonly spotPrice: string;
    /**
     * The User Data of the instance.
     */
    readonly userData: string;
    /**
     * The ID of a ClassicLink-enabled VPC.
     */
    readonly vpcClassicLinkId: string;
    /**
     * The IDs of one or more Security Groups for the specified ClassicLink-enabled VPC.
     */
    readonly vpcClassicLinkSecurityGroups: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
