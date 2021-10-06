# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterClusterNodeArgs',
    'ClusterLoggingArgs',
    'ClusterSnapshotCopyArgs',
    'ParameterGroupParameterArgs',
    'ScheduledActionTargetActionArgs',
    'ScheduledActionTargetActionPauseClusterArgs',
    'ScheduledActionTargetActionResizeClusterArgs',
    'ScheduledActionTargetActionResumeClusterArgs',
    'SecurityGroupIngressArgs',
]

@pulumi.input_type
class ClusterClusterNodeArgs:
    def __init__(__self__, *,
                 node_role: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 public_ip_address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] node_role: Whether the node is a leader node or a compute node
        :param pulumi.Input[str] private_ip_address: The private IP address of a node within a cluster
        :param pulumi.Input[str] public_ip_address: The public IP address of a node within a cluster
        """
        if node_role is not None:
            pulumi.set(__self__, "node_role", node_role)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if public_ip_address is not None:
            pulumi.set(__self__, "public_ip_address", public_ip_address)

    @property
    @pulumi.getter(name="nodeRole")
    def node_role(self) -> Optional[pulumi.Input[str]]:
        """
        Whether the node is a leader node or a compute node
        """
        return pulumi.get(self, "node_role")

    @node_role.setter
    def node_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_role", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The private IP address of a node within a cluster
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The public IP address of a node within a cluster
        """
        return pulumi.get(self, "public_ip_address")

    @public_ip_address.setter
    def public_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_ip_address", value)


@pulumi.input_type
class ClusterLoggingArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[bool],
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] enable: Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
        :param pulumi.Input[str] bucket_name: The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
               For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
        :param pulumi.Input[str] s3_key_prefix: The prefix applied to the log file names.
        """
        pulumi.set(__self__, "enable", enable)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if s3_key_prefix is not None:
            pulumi.set(__self__, "s3_key_prefix", s3_key_prefix)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        """
        Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
        For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="s3KeyPrefix")
    def s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix applied to the log file names.
        """
        return pulumi.get(self, "s3_key_prefix")

    @s3_key_prefix.setter
    def s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_key_prefix", value)


@pulumi.input_type
class ClusterSnapshotCopyArgs:
    def __init__(__self__, *,
                 destination_region: pulumi.Input[str],
                 grant_name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] destination_region: The destination region that you want to copy snapshots to.
        :param pulumi.Input[str] grant_name: The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        :param pulumi.Input[int] retention_period: The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
        """
        pulumi.set(__self__, "destination_region", destination_region)
        if grant_name is not None:
            pulumi.set(__self__, "grant_name", grant_name)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)

    @property
    @pulumi.getter(name="destinationRegion")
    def destination_region(self) -> pulumi.Input[str]:
        """
        The destination region that you want to copy snapshots to.
        """
        return pulumi.get(self, "destination_region")

    @destination_region.setter
    def destination_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_region", value)

    @property
    @pulumi.getter(name="grantName")
    def grant_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
        """
        return pulumi.get(self, "grant_name")

    @grant_name.setter
    def grant_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_name", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        The number of days to retain automated snapshots in the destination region after they are copied from the source region. Defaults to `7`.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period", value)


@pulumi.input_type
class ParameterGroupParameterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: The name of the Redshift parameter.
        :param pulumi.Input[str] value: The value of the Redshift parameter.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the Redshift parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the Redshift parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ScheduledActionTargetActionArgs:
    def __init__(__self__, *,
                 pause_cluster: Optional[pulumi.Input['ScheduledActionTargetActionPauseClusterArgs']] = None,
                 resize_cluster: Optional[pulumi.Input['ScheduledActionTargetActionResizeClusterArgs']] = None,
                 resume_cluster: Optional[pulumi.Input['ScheduledActionTargetActionResumeClusterArgs']] = None):
        """
        :param pulumi.Input['ScheduledActionTargetActionPauseClusterArgs'] pause_cluster: An action that runs a `PauseCluster` API operation. Documented below.
        :param pulumi.Input['ScheduledActionTargetActionResizeClusterArgs'] resize_cluster: An action that runs a `ResizeCluster` API operation. Documented below.
        :param pulumi.Input['ScheduledActionTargetActionResumeClusterArgs'] resume_cluster: An action that runs a `ResumeCluster` API operation. Documented below.
        """
        if pause_cluster is not None:
            pulumi.set(__self__, "pause_cluster", pause_cluster)
        if resize_cluster is not None:
            pulumi.set(__self__, "resize_cluster", resize_cluster)
        if resume_cluster is not None:
            pulumi.set(__self__, "resume_cluster", resume_cluster)

    @property
    @pulumi.getter(name="pauseCluster")
    def pause_cluster(self) -> Optional[pulumi.Input['ScheduledActionTargetActionPauseClusterArgs']]:
        """
        An action that runs a `PauseCluster` API operation. Documented below.
        """
        return pulumi.get(self, "pause_cluster")

    @pause_cluster.setter
    def pause_cluster(self, value: Optional[pulumi.Input['ScheduledActionTargetActionPauseClusterArgs']]):
        pulumi.set(self, "pause_cluster", value)

    @property
    @pulumi.getter(name="resizeCluster")
    def resize_cluster(self) -> Optional[pulumi.Input['ScheduledActionTargetActionResizeClusterArgs']]:
        """
        An action that runs a `ResizeCluster` API operation. Documented below.
        """
        return pulumi.get(self, "resize_cluster")

    @resize_cluster.setter
    def resize_cluster(self, value: Optional[pulumi.Input['ScheduledActionTargetActionResizeClusterArgs']]):
        pulumi.set(self, "resize_cluster", value)

    @property
    @pulumi.getter(name="resumeCluster")
    def resume_cluster(self) -> Optional[pulumi.Input['ScheduledActionTargetActionResumeClusterArgs']]:
        """
        An action that runs a `ResumeCluster` API operation. Documented below.
        """
        return pulumi.get(self, "resume_cluster")

    @resume_cluster.setter
    def resume_cluster(self, value: Optional[pulumi.Input['ScheduledActionTargetActionResumeClusterArgs']]):
        pulumi.set(self, "resume_cluster", value)


@pulumi.input_type
class ScheduledActionTargetActionPauseClusterArgs:
    def __init__(__self__, *,
                 cluster_identifier: pulumi.Input[str]):
        """
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster to be resumed.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the cluster to be resumed.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)


@pulumi.input_type
class ScheduledActionTargetActionResizeClusterArgs:
    def __init__(__self__, *,
                 cluster_identifier: pulumi.Input[str],
                 classic: Optional[pulumi.Input[bool]] = None,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 number_of_nodes: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster to be resumed.
        :param pulumi.Input[bool] classic: A boolean value indicating whether the resize operation is using the classic resize process. Default: `false`.
        :param pulumi.Input[str] cluster_type: The new cluster type for the specified cluster.
        :param pulumi.Input[str] node_type: The new node type for the nodes you are adding.
        :param pulumi.Input[int] number_of_nodes: The new number of nodes for the cluster.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        if classic is not None:
            pulumi.set(__self__, "classic", classic)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if number_of_nodes is not None:
            pulumi.set(__self__, "number_of_nodes", number_of_nodes)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the cluster to be resumed.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter
    def classic(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean value indicating whether the resize operation is using the classic resize process. Default: `false`.
        """
        return pulumi.get(self, "classic")

    @classic.setter
    def classic(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "classic", value)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[pulumi.Input[str]]:
        """
        The new cluster type for the specified cluster.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        The new node type for the nodes you are adding.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="numberOfNodes")
    def number_of_nodes(self) -> Optional[pulumi.Input[int]]:
        """
        The new number of nodes for the cluster.
        """
        return pulumi.get(self, "number_of_nodes")

    @number_of_nodes.setter
    def number_of_nodes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number_of_nodes", value)


@pulumi.input_type
class ScheduledActionTargetActionResumeClusterArgs:
    def __init__(__self__, *,
                 cluster_identifier: pulumi.Input[str]):
        """
        :param pulumi.Input[str] cluster_identifier: The identifier of the cluster to be resumed.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the cluster to be resumed.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)


@pulumi.input_type
class SecurityGroupIngressArgs:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 security_group_name: Optional[pulumi.Input[str]] = None,
                 security_group_owner_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cidr: The CIDR block to accept
        :param pulumi.Input[str] security_group_name: The name of the security group to authorize
        :param pulumi.Input[str] security_group_owner_id: The owner Id of the security group provided
               by `security_group_name`.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if security_group_name is not None:
            pulumi.set(__self__, "security_group_name", security_group_name)
        if security_group_owner_id is not None:
            pulumi.set(__self__, "security_group_owner_id", security_group_owner_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block to accept
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="securityGroupName")
    def security_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security group to authorize
        """
        return pulumi.get(self, "security_group_name")

    @security_group_name.setter
    def security_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_name", value)

    @property
    @pulumi.getter(name="securityGroupOwnerId")
    def security_group_owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner Id of the security group provided
        by `security_group_name`.
        """
        return pulumi.get(self, "security_group_owner_id")

    @security_group_owner_id.setter
    def security_group_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_owner_id", value)


