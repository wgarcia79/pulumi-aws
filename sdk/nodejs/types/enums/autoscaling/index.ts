// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const Metric = {
    GroupMinSize: "GroupMinSize",
    GroupMaxSize: "GroupMaxSize",
    GroupDesiredCapacity: "GroupDesiredCapacity",
    GroupInServiceInstances: "GroupInServiceInstances",
    GroupPendingInstances: "GroupPendingInstances",
    GroupStandbyInstances: "GroupStandbyInstances",
    GroupTerminatingInstances: "GroupTerminatingInstances",
    GroupTotalInstances: "GroupTotalInstances",
} as const;

/**
 * See https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html
 */
export type Metric = (typeof Metric)[keyof typeof Metric];

export const MetricsGranularity = {
    OneMinute: "1Minute",
} as const;

/**
 * See https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_EnableMetricsCollection.html
 */
export type MetricsGranularity = (typeof MetricsGranularity)[keyof typeof MetricsGranularity];

export const NotificationType = {
    InstanceLaunch: "autoscaling:EC2_INSTANCE_LAUNCH",
    InstanceTerminate: "autoscaling:EC2_INSTANCE_TERMINATE",
    InstanceLaunchError: "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
    InstanceTerminateError: "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    TestNotification: "autoscaling:TEST_NOTIFICATION",
} as const;

/**
 * See https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_NotificationConfiguration.html
 */
export type NotificationType = (typeof NotificationType)[keyof typeof NotificationType];