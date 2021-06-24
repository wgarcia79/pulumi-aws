// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages AWS Managed Streaming for Kafka cluster
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const vpc = new aws.ec2.Vpc("vpc", {cidrBlock: "192.168.0.0/22"});
 * const azs = aws.getAvailabilityZones({
 *     state: "available",
 * });
 * const subnetAz1 = new aws.ec2.Subnet("subnetAz1", {
 *     availabilityZone: azs.then(azs => azs.names?[0]),
 *     cidrBlock: "192.168.0.0/24",
 *     vpcId: vpc.id,
 * });
 * const subnetAz2 = new aws.ec2.Subnet("subnetAz2", {
 *     availabilityZone: azs.then(azs => azs.names?[1]),
 *     cidrBlock: "192.168.1.0/24",
 *     vpcId: vpc.id,
 * });
 * const subnetAz3 = new aws.ec2.Subnet("subnetAz3", {
 *     availabilityZone: azs.then(azs => azs.names?[2]),
 *     cidrBlock: "192.168.2.0/24",
 *     vpcId: vpc.id,
 * });
 * const sg = new aws.ec2.SecurityGroup("sg", {vpcId: vpc.id});
 * const kms = new aws.kms.Key("kms", {description: "example"});
 * const test = new aws.cloudwatch.LogGroup("test", {});
 * const bucket = new aws.s3.Bucket("bucket", {acl: "private"});
 * const firehoseRole = new aws.iam.Role("firehoseRole", {assumeRolePolicy: `{
 * "Version": "2012-10-17",
 * "Statement": [
 *   {
 *     "Action": "sts:AssumeRole",
 *     "Principal": {
 *       "Service": "firehose.amazonaws.com"
 *     },
 *     "Effect": "Allow",
 *     "Sid": ""
 *   }
 *   ]
 * }
 * `});
 * const testStream = new aws.kinesis.FirehoseDeliveryStream("testStream", {
 *     destination: "s3",
 *     s3Configuration: {
 *         roleArn: firehoseRole.arn,
 *         bucketArn: bucket.arn,
 *     },
 *     tags: {
 *         LogDeliveryEnabled: "placeholder",
 *     },
 * });
 * const example = new aws.msk.Cluster("example", {
 *     kafkaVersion: "2.4.1",
 *     numberOfBrokerNodes: 3,
 *     brokerNodeGroupInfo: {
 *         instanceType: "kafka.m5.large",
 *         ebsVolumeSize: 1000,
 *         clientSubnets: [
 *             subnetAz1.id,
 *             subnetAz2.id,
 *             subnetAz3.id,
 *         ],
 *         securityGroups: [sg.id],
 *     },
 *     encryptionInfo: {
 *         encryptionAtRestKmsKeyArn: kms.arn,
 *     },
 *     openMonitoring: {
 *         prometheus: {
 *             jmxExporter: {
 *                 enabledInBroker: true,
 *             },
 *             nodeExporter: {
 *                 enabledInBroker: true,
 *             },
 *         },
 *     },
 *     loggingInfo: {
 *         brokerLogs: {
 *             cloudwatchLogs: {
 *                 enabled: true,
 *                 logGroup: test.name,
 *             },
 *             firehose: {
 *                 enabled: true,
 *                 deliveryStream: testStream.name,
 *             },
 *             s3: {
 *                 enabled: true,
 *                 bucket: bucket.id,
 *                 prefix: "logs/msk-",
 *             },
 *         },
 *     },
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * export const zookeeperConnectString = example.zookeeperConnectString;
 * export const bootstrapBrokersTls = example.bootstrapBrokersTls;
 * ```
 *
 * ## Import
 *
 * MSK clusters can be imported using the cluster `arn`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:msk/cluster:Cluster example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:msk/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
     */
    public /*out*/ readonly bootstrapBrokers!: pulumi.Output<string>;
    /**
     * One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    public /*out*/ readonly bootstrapBrokersSaslIam!: pulumi.Output<string>;
    /**
     * One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    public /*out*/ readonly bootstrapBrokersSaslScram!: pulumi.Output<string>;
    /**
     * One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    public /*out*/ readonly bootstrapBrokersTls!: pulumi.Output<string>;
    /**
     * Configuration block for the broker nodes of the Kafka cluster.
     */
    public readonly brokerNodeGroupInfo!: pulumi.Output<outputs.msk.ClusterBrokerNodeGroupInfo>;
    /**
     * Configuration block for specifying a client authentication. See below.
     */
    public readonly clientAuthentication!: pulumi.Output<outputs.msk.ClusterClientAuthentication | undefined>;
    /**
     * Name of the MSK cluster.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
     */
    public readonly configurationInfo!: pulumi.Output<outputs.msk.ClusterConfigurationInfo | undefined>;
    /**
     * Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
     * * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
     */
    public /*out*/ readonly currentVersion!: pulumi.Output<string>;
    /**
     * Configuration block for specifying encryption. See below.
     */
    public readonly encryptionInfo!: pulumi.Output<outputs.msk.ClusterEncryptionInfo | undefined>;
    /**
     * Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
     */
    public readonly enhancedMonitoring!: pulumi.Output<string | undefined>;
    /**
     * Specify the desired Kafka software version.
     */
    public readonly kafkaVersion!: pulumi.Output<string>;
    /**
     * Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
     */
    public readonly loggingInfo!: pulumi.Output<outputs.msk.ClusterLoggingInfo | undefined>;
    /**
     * The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
     */
    public readonly numberOfBrokerNodes!: pulumi.Output<number>;
    /**
     * Configuration block for JMX and Node monitoring for the MSK cluster. See below.
     */
    public readonly openMonitoring!: pulumi.Output<outputs.msk.ClusterOpenMonitoring | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     */
    public /*out*/ readonly zookeeperConnectString!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bootstrapBrokers"] = state ? state.bootstrapBrokers : undefined;
            inputs["bootstrapBrokersSaslIam"] = state ? state.bootstrapBrokersSaslIam : undefined;
            inputs["bootstrapBrokersSaslScram"] = state ? state.bootstrapBrokersSaslScram : undefined;
            inputs["bootstrapBrokersTls"] = state ? state.bootstrapBrokersTls : undefined;
            inputs["brokerNodeGroupInfo"] = state ? state.brokerNodeGroupInfo : undefined;
            inputs["clientAuthentication"] = state ? state.clientAuthentication : undefined;
            inputs["clusterName"] = state ? state.clusterName : undefined;
            inputs["configurationInfo"] = state ? state.configurationInfo : undefined;
            inputs["currentVersion"] = state ? state.currentVersion : undefined;
            inputs["encryptionInfo"] = state ? state.encryptionInfo : undefined;
            inputs["enhancedMonitoring"] = state ? state.enhancedMonitoring : undefined;
            inputs["kafkaVersion"] = state ? state.kafkaVersion : undefined;
            inputs["loggingInfo"] = state ? state.loggingInfo : undefined;
            inputs["numberOfBrokerNodes"] = state ? state.numberOfBrokerNodes : undefined;
            inputs["openMonitoring"] = state ? state.openMonitoring : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["zookeeperConnectString"] = state ? state.zookeeperConnectString : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.brokerNodeGroupInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'brokerNodeGroupInfo'");
            }
            if ((!args || args.kafkaVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaVersion'");
            }
            if ((!args || args.numberOfBrokerNodes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'numberOfBrokerNodes'");
            }
            inputs["brokerNodeGroupInfo"] = args ? args.brokerNodeGroupInfo : undefined;
            inputs["clientAuthentication"] = args ? args.clientAuthentication : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["configurationInfo"] = args ? args.configurationInfo : undefined;
            inputs["encryptionInfo"] = args ? args.encryptionInfo : undefined;
            inputs["enhancedMonitoring"] = args ? args.enhancedMonitoring : undefined;
            inputs["kafkaVersion"] = args ? args.kafkaVersion : undefined;
            inputs["loggingInfo"] = args ? args.loggingInfo : undefined;
            inputs["numberOfBrokerNodes"] = args ? args.numberOfBrokerNodes : undefined;
            inputs["openMonitoring"] = args ? args.openMonitoring : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["bootstrapBrokers"] = undefined /*out*/;
            inputs["bootstrapBrokersSaslIam"] = undefined /*out*/;
            inputs["bootstrapBrokersSaslScram"] = undefined /*out*/;
            inputs["bootstrapBrokersTls"] = undefined /*out*/;
            inputs["currentVersion"] = undefined /*out*/;
            inputs["zookeeperConnectString"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Amazon Resource Name (ARN) of the MSK Configuration to use in the cluster.
     */
    arn?: pulumi.Input<string>;
    /**
     * Comma separated list of one or more hostname:port pairs of kafka brokers suitable to bootstrap connectivity to the kafka cluster. Contains a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The resource sorts values alphabetically. AWS may not always return all endpoints so this value is not guaranteed to be stable across applies.
     */
    bootstrapBrokers?: pulumi.Input<string>;
    /**
     * One or more DNS names (or IP addresses) and SASL IAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9098`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.iam` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    bootstrapBrokersSaslIam?: pulumi.Input<string>;
    /**
     * One or more DNS names (or IP addresses) and SASL SCRAM port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9096`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS` and `client_authentication.0.sasl.0.scram` is set to `true`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    bootstrapBrokersSaslScram?: pulumi.Input<string>;
    /**
     * One or more DNS names (or IP addresses) and TLS port pairs. For example, `b-1.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-2.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094,b-3.exampleClusterName.abcde.c2.kafka.us-east-1.amazonaws.com:9094`. This attribute will have a value if `encryption_info.0.encryption_in_transit.0.client_broker` is set to `TLS_PLAINTEXT` or `TLS`. The resource sorts the list alphabetically. AWS may not always return all endpoints so the values may not be stable across applies.
     */
    bootstrapBrokersTls?: pulumi.Input<string>;
    /**
     * Configuration block for the broker nodes of the Kafka cluster.
     */
    brokerNodeGroupInfo?: pulumi.Input<inputs.msk.ClusterBrokerNodeGroupInfo>;
    /**
     * Configuration block for specifying a client authentication. See below.
     */
    clientAuthentication?: pulumi.Input<inputs.msk.ClusterClientAuthentication>;
    /**
     * Name of the MSK cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
     */
    configurationInfo?: pulumi.Input<inputs.msk.ClusterConfigurationInfo>;
    /**
     * Current version of the MSK Cluster used for updates, e.g. `K13V1IB3VIYZZH`
     * * `encryption_info.0.encryption_at_rest_kms_key_arn` - The ARN of the KMS key used for encryption at rest of the broker data volumes.
     */
    currentVersion?: pulumi.Input<string>;
    /**
     * Configuration block for specifying encryption. See below.
     */
    encryptionInfo?: pulumi.Input<inputs.msk.ClusterEncryptionInfo>;
    /**
     * Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
     */
    enhancedMonitoring?: pulumi.Input<string>;
    /**
     * Specify the desired Kafka software version.
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
     */
    loggingInfo?: pulumi.Input<inputs.msk.ClusterLoggingInfo>;
    /**
     * The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
     */
    numberOfBrokerNodes?: pulumi.Input<number>;
    /**
     * Configuration block for JMX and Node monitoring for the MSK cluster. See below.
     */
    openMonitoring?: pulumi.Input<inputs.msk.ClusterOpenMonitoring>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
     */
    zookeeperConnectString?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Configuration block for the broker nodes of the Kafka cluster.
     */
    brokerNodeGroupInfo: pulumi.Input<inputs.msk.ClusterBrokerNodeGroupInfo>;
    /**
     * Configuration block for specifying a client authentication. See below.
     */
    clientAuthentication?: pulumi.Input<inputs.msk.ClusterClientAuthentication>;
    /**
     * Name of the MSK cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Configuration block for specifying a MSK Configuration to attach to Kafka brokers. See below.
     */
    configurationInfo?: pulumi.Input<inputs.msk.ClusterConfigurationInfo>;
    /**
     * Configuration block for specifying encryption. See below.
     */
    encryptionInfo?: pulumi.Input<inputs.msk.ClusterEncryptionInfo>;
    /**
     * Specify the desired enhanced MSK CloudWatch monitoring level.  See [Monitoring Amazon MSK with Amazon CloudWatch](https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html)
     */
    enhancedMonitoring?: pulumi.Input<string>;
    /**
     * Specify the desired Kafka software version.
     */
    kafkaVersion: pulumi.Input<string>;
    /**
     * Configuration block for streaming broker logs to Cloudwatch/S3/Kinesis Firehose. See below.
     */
    loggingInfo?: pulumi.Input<inputs.msk.ClusterLoggingInfo>;
    /**
     * The desired total number of broker nodes in the kafka cluster.  It must be a multiple of the number of specified client subnets.
     */
    numberOfBrokerNodes: pulumi.Input<number>;
    /**
     * Configuration block for JMX and Node monitoring for the MSK cluster. See below.
     */
    openMonitoring?: pulumi.Input<inputs.msk.ClusterOpenMonitoring>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
