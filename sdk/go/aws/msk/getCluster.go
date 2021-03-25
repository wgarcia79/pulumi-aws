// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on an Amazon MSK Cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/msk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := msk.LookupCluster(ctx, &msk.LookupClusterArgs{
// 			ClusterName: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("aws:msk/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// Map of key-value pairs assigned to the cluster.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	Arn string `pulumi:"arn"`
	// A comma separated list of one or more hostname:port pairs of Kafka brokers suitable to boostrap connectivity to the Kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `PLAINTEXT` or `TLS_PLAINTEXT`. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	BootstrapBrokers string `pulumi:"bootstrapBrokers"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity using SASL/SCRAM to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS` and `clientAuthentication` is set to `sasl`. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	BootstrapBrokersSaslScram string `pulumi:"bootstrapBrokersSaslScram"`
	// A comma separated list of one or more DNS names (or IPs) and TLS port pairs kafka brokers suitable to boostrap connectivity to the kafka cluster. Only contains value if `clientBroker` encryption in transit is set to `TLS_PLAINTEXT` or `TLS`. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	BootstrapBrokersTls string `pulumi:"bootstrapBrokersTls"`
	ClusterName         string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Apache Kafka version.
	KafkaVersion string `pulumi:"kafkaVersion"`
	// Number of broker nodes in the cluster.
	NumberOfBrokerNodes int `pulumi:"numberOfBrokerNodes"`
	// Map of key-value pairs assigned to the cluster.
	Tags map[string]string `pulumi:"tags"`
	// A comma separated list of one or more hostname:port pairs to use to connect to the Apache Zookeeper cluster. The returned values are sorted alphbetically. The AWS API may not return all endpoints, so this value is not guaranteed to be stable across applies.
	ZookeeperConnectString string `pulumi:"zookeeperConnectString"`
}
