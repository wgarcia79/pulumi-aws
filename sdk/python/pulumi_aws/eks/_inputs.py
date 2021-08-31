# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterCertificateAuthorityArgs',
    'ClusterEncryptionConfigArgs',
    'ClusterEncryptionConfigProviderArgs',
    'ClusterIdentityArgs',
    'ClusterIdentityOidcArgs',
    'ClusterKubernetesNetworkConfigArgs',
    'ClusterVpcConfigArgs',
    'FargateProfileSelectorArgs',
    'IdentityProviderConfigOidcArgs',
    'NodeGroupLaunchTemplateArgs',
    'NodeGroupRemoteAccessArgs',
    'NodeGroupResourceArgs',
    'NodeGroupResourceAutoscalingGroupArgs',
    'NodeGroupScalingConfigArgs',
    'NodeGroupTaintArgs',
    'NodeGroupUpdateConfigArgs',
]

@pulumi.input_type
class ClusterCertificateAuthorityArgs:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] data: Base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
        """
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class ClusterEncryptionConfigArgs:
    def __init__(__self__, *,
                 provider: pulumi.Input['ClusterEncryptionConfigProviderArgs'],
                 resources: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input['ClusterEncryptionConfigProviderArgs'] provider: Configuration block with provider for encryption. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: List of strings with resources to be encrypted. Valid values: `secrets`.
        """
        pulumi.set(__self__, "provider", provider)
        pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Input['ClusterEncryptionConfigProviderArgs']:
        """
        Configuration block with provider for encryption. Detailed below.
        """
        return pulumi.get(self, "provider")

    @provider.setter
    def provider(self, value: pulumi.Input['ClusterEncryptionConfigProviderArgs']):
        pulumi.set(self, "provider", value)

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of strings with resources to be encrypted. Valid values: `secrets`.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "resources", value)


@pulumi.input_type
class ClusterEncryptionConfigProviderArgs:
    def __init__(__self__, *,
                 key_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key_arn: ARN of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).
        """
        pulumi.set(__self__, "key_arn", key_arn)

    @property
    @pulumi.getter(name="keyArn")
    def key_arn(self) -> pulumi.Input[str]:
        """
        ARN of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).
        """
        return pulumi.get(self, "key_arn")

    @key_arn.setter
    def key_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_arn", value)


@pulumi.input_type
class ClusterIdentityArgs:
    def __init__(__self__, *,
                 oidcs: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterIdentityOidcArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['ClusterIdentityOidcArgs']]] oidcs: Nested block containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster. Detailed below.
        """
        if oidcs is not None:
            pulumi.set(__self__, "oidcs", oidcs)

    @property
    @pulumi.getter
    def oidcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterIdentityOidcArgs']]]]:
        """
        Nested block containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster. Detailed below.
        """
        return pulumi.get(self, "oidcs")

    @oidcs.setter
    def oidcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterIdentityOidcArgs']]]]):
        pulumi.set(self, "oidcs", value)


@pulumi.input_type
class ClusterIdentityOidcArgs:
    def __init__(__self__, *,
                 issuer: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] issuer: Issuer URL for the OpenID Connect identity provider.
        """
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        Issuer URL for the OpenID Connect identity provider.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)


@pulumi.input_type
class ClusterKubernetesNetworkConfigArgs:
    def __init__(__self__, *,
                 service_ipv4_cidr: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] service_ipv4_cidr: The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. You can only specify a custom CIDR block when you create a cluster, changing this value will force a new cluster to be created. The block must meet the following requirements:
        """
        if service_ipv4_cidr is not None:
            pulumi.set(__self__, "service_ipv4_cidr", service_ipv4_cidr)

    @property
    @pulumi.getter(name="serviceIpv4Cidr")
    def service_ipv4_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. You can only specify a custom CIDR block when you create a cluster, changing this value will force a new cluster to be created. The block must meet the following requirements:
        """
        return pulumi.get(self, "service_ipv4_cidr")

    @service_ipv4_cidr.setter
    def service_ipv4_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_ipv4_cidr", value)


@pulumi.input_type
class ClusterVpcConfigArgs:
    def __init__(__self__, *,
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cluster_security_group_id: Optional[pulumi.Input[str]] = None,
                 endpoint_private_access: Optional[pulumi.Input[bool]] = None,
                 endpoint_public_access: Optional[pulumi.Input[bool]] = None,
                 public_access_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
        :param pulumi.Input[str] cluster_security_group_id: Cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control-plane-to-data-plane communication.
        :param pulumi.Input[bool] endpoint_private_access: Whether the Amazon EKS private API server endpoint is enabled. Default is `false`.
        :param pulumi.Input[bool] endpoint_public_access: Whether the Amazon EKS public API server endpoint is enabled. Default is `true`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_access_cidrs: List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. This provider will only perform drift detection of its value when present in a configuration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
        :param pulumi.Input[str] vpc_id: ID of the VPC associated with your cluster.
        """
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if cluster_security_group_id is not None:
            pulumi.set(__self__, "cluster_security_group_id", cluster_security_group_id)
        if endpoint_private_access is not None:
            pulumi.set(__self__, "endpoint_private_access", endpoint_private_access)
        if endpoint_public_access is not None:
            pulumi.set(__self__, "endpoint_public_access", endpoint_public_access)
        if public_access_cidrs is not None:
            pulumi.set(__self__, "public_access_cidrs", public_access_cidrs)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="clusterSecurityGroupId")
    def cluster_security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control-plane-to-data-plane communication.
        """
        return pulumi.get(self, "cluster_security_group_id")

    @cluster_security_group_id.setter
    def cluster_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_security_group_id", value)

    @property
    @pulumi.getter(name="endpointPrivateAccess")
    def endpoint_private_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Amazon EKS private API server endpoint is enabled. Default is `false`.
        """
        return pulumi.get(self, "endpoint_private_access")

    @endpoint_private_access.setter
    def endpoint_private_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "endpoint_private_access", value)

    @property
    @pulumi.getter(name="endpointPublicAccess")
    def endpoint_public_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Amazon EKS public API server endpoint is enabled. Default is `true`.
        """
        return pulumi.get(self, "endpoint_public_access")

    @endpoint_public_access.setter
    def endpoint_public_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "endpoint_public_access", value)

    @property
    @pulumi.getter(name="publicAccessCidrs")
    def public_access_cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. This provider will only perform drift detection of its value when present in a configuration.
        """
        return pulumi.get(self, "public_access_cidrs")

    @public_access_cidrs.setter
    def public_access_cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "public_access_cidrs", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the VPC associated with your cluster.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class FargateProfileSelectorArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] namespace: Kubernetes namespace for selection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Key-value map of Kubernetes labels for selection.
        """
        pulumi.set(__self__, "namespace", namespace)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        Kubernetes namespace for selection.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of Kubernetes labels for selection.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)


@pulumi.input_type
class IdentityProviderConfigOidcArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 identity_provider_config_name: pulumi.Input[str],
                 issuer_url: pulumi.Input[str],
                 groups_claim: Optional[pulumi.Input[str]] = None,
                 groups_prefix: Optional[pulumi.Input[str]] = None,
                 required_claims: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 username_claim: Optional[pulumi.Input[str]] = None,
                 username_prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] client_id: Client ID for the OpenID Connect identity provider.
        :param pulumi.Input[str] identity_provider_config_name: The name of the identity provider config.
        :param pulumi.Input[str] issuer_url: Issuer URL for the OpenID Connect identity provider.
        :param pulumi.Input[str] groups_claim: The JWT claim that the provider will use to return groups.
        :param pulumi.Input[str] groups_prefix: A prefix that is prepended to group claims e.g. `oidc:`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] required_claims: The key value pairs that describe required claims in the identity token.
        :param pulumi.Input[str] username_claim: The JWT claim that the provider will use as the username.
        :param pulumi.Input[str] username_prefix: A prefix that is prepended to username claims.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "identity_provider_config_name", identity_provider_config_name)
        pulumi.set(__self__, "issuer_url", issuer_url)
        if groups_claim is not None:
            pulumi.set(__self__, "groups_claim", groups_claim)
        if groups_prefix is not None:
            pulumi.set(__self__, "groups_prefix", groups_prefix)
        if required_claims is not None:
            pulumi.set(__self__, "required_claims", required_claims)
        if username_claim is not None:
            pulumi.set(__self__, "username_claim", username_claim)
        if username_prefix is not None:
            pulumi.set(__self__, "username_prefix", username_prefix)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        Client ID for the OpenID Connect identity provider.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="identityProviderConfigName")
    def identity_provider_config_name(self) -> pulumi.Input[str]:
        """
        The name of the identity provider config.
        """
        return pulumi.get(self, "identity_provider_config_name")

    @identity_provider_config_name.setter
    def identity_provider_config_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_provider_config_name", value)

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> pulumi.Input[str]:
        """
        Issuer URL for the OpenID Connect identity provider.
        """
        return pulumi.get(self, "issuer_url")

    @issuer_url.setter
    def issuer_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer_url", value)

    @property
    @pulumi.getter(name="groupsClaim")
    def groups_claim(self) -> Optional[pulumi.Input[str]]:
        """
        The JWT claim that the provider will use to return groups.
        """
        return pulumi.get(self, "groups_claim")

    @groups_claim.setter
    def groups_claim(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "groups_claim", value)

    @property
    @pulumi.getter(name="groupsPrefix")
    def groups_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix that is prepended to group claims e.g. `oidc:`.
        """
        return pulumi.get(self, "groups_prefix")

    @groups_prefix.setter
    def groups_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "groups_prefix", value)

    @property
    @pulumi.getter(name="requiredClaims")
    def required_claims(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The key value pairs that describe required claims in the identity token.
        """
        return pulumi.get(self, "required_claims")

    @required_claims.setter
    def required_claims(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "required_claims", value)

    @property
    @pulumi.getter(name="usernameClaim")
    def username_claim(self) -> Optional[pulumi.Input[str]]:
        """
        The JWT claim that the provider will use as the username.
        """
        return pulumi.get(self, "username_claim")

    @username_claim.setter
    def username_claim(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_claim", value)

    @property
    @pulumi.getter(name="usernamePrefix")
    def username_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix that is prepended to username claims.
        """
        return pulumi.get(self, "username_prefix")

    @username_prefix.setter
    def username_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_prefix", value)


@pulumi.input_type
class NodeGroupLaunchTemplateArgs:
    def __init__(__self__, *,
                 version: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] version: EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `ec2.LaunchTemplate` resource or data source is recommended for this argument.
        :param pulumi.Input[str] id: Identifier of the EC2 Launch Template. Conflicts with `name`.
        :param pulumi.Input[str] name: Name of the EC2 Launch Template. Conflicts with `id`.
        """
        pulumi.set(__self__, "version", version)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g. `1`) on read and This provider will show a difference on next plan. Using the `default_version` or `latest_version` attribute of the `ec2.LaunchTemplate` resource or data source is recommended for this argument.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the EC2 Launch Template. Conflicts with `name`.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the EC2 Launch Template. Conflicts with `id`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class NodeGroupRemoteAccessArgs:
    def __init__(__self__, *,
                 ec2_ssh_key: Optional[pulumi.Input[str]] = None,
                 source_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] ec2_ssh_key: EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] source_security_group_ids: Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
        """
        if ec2_ssh_key is not None:
            pulumi.set(__self__, "ec2_ssh_key", ec2_ssh_key)
        if source_security_group_ids is not None:
            pulumi.set(__self__, "source_security_group_ids", source_security_group_ids)

    @property
    @pulumi.getter(name="ec2SshKey")
    def ec2_ssh_key(self) -> Optional[pulumi.Input[str]]:
        """
        EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `source_security_group_ids` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
        """
        return pulumi.get(self, "ec2_ssh_key")

    @ec2_ssh_key.setter
    def ec2_ssh_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ec2_ssh_key", value)

    @property
    @pulumi.getter(name="sourceSecurityGroupIds")
    def source_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2_ssh_key`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
        """
        return pulumi.get(self, "source_security_group_ids")

    @source_security_group_ids.setter
    def source_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "source_security_group_ids", value)


@pulumi.input_type
class NodeGroupResourceArgs:
    def __init__(__self__, *,
                 autoscaling_groups: Optional[pulumi.Input[Sequence[pulumi.Input['NodeGroupResourceAutoscalingGroupArgs']]]] = None,
                 remote_access_security_group_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['NodeGroupResourceAutoscalingGroupArgs']]] autoscaling_groups: List of objects containing information about AutoScaling Groups.
        :param pulumi.Input[str] remote_access_security_group_id: Identifier of the remote access EC2 Security Group.
        """
        if autoscaling_groups is not None:
            pulumi.set(__self__, "autoscaling_groups", autoscaling_groups)
        if remote_access_security_group_id is not None:
            pulumi.set(__self__, "remote_access_security_group_id", remote_access_security_group_id)

    @property
    @pulumi.getter(name="autoscalingGroups")
    def autoscaling_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NodeGroupResourceAutoscalingGroupArgs']]]]:
        """
        List of objects containing information about AutoScaling Groups.
        """
        return pulumi.get(self, "autoscaling_groups")

    @autoscaling_groups.setter
    def autoscaling_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NodeGroupResourceAutoscalingGroupArgs']]]]):
        pulumi.set(self, "autoscaling_groups", value)

    @property
    @pulumi.getter(name="remoteAccessSecurityGroupId")
    def remote_access_security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the remote access EC2 Security Group.
        """
        return pulumi.get(self, "remote_access_security_group_id")

    @remote_access_security_group_id.setter
    def remote_access_security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_access_security_group_id", value)


@pulumi.input_type
class NodeGroupResourceAutoscalingGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Name of the EC2 Launch Template. Conflicts with `id`.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the EC2 Launch Template. Conflicts with `id`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class NodeGroupScalingConfigArgs:
    def __init__(__self__, *,
                 desired_size: pulumi.Input[int],
                 max_size: pulumi.Input[int],
                 min_size: pulumi.Input[int]):
        """
        :param pulumi.Input[int] desired_size: Desired number of worker nodes.
        :param pulumi.Input[int] max_size: Maximum number of worker nodes.
        :param pulumi.Input[int] min_size: Minimum number of worker nodes.
        """
        pulumi.set(__self__, "desired_size", desired_size)
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)

    @property
    @pulumi.getter(name="desiredSize")
    def desired_size(self) -> pulumi.Input[int]:
        """
        Desired number of worker nodes.
        """
        return pulumi.get(self, "desired_size")

    @desired_size.setter
    def desired_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_size", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Input[int]:
        """
        Maximum number of worker nodes.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Input[int]:
        """
        Minimum number of worker nodes.
        """
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_size", value)


@pulumi.input_type
class NodeGroupTaintArgs:
    def __init__(__self__, *,
                 effect: pulumi.Input[str],
                 key: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] effect: The effect of the taint. Valid values: `NO_SCHEDULE`, `NO_EXECUTE`, `PREFER_NO_SCHEDULE`.
        :param pulumi.Input[str] key: The key of the taint. Maximum length of 63.
        :param pulumi.Input[str] value: The value of the taint. Maximum length of 63.
        """
        pulumi.set(__self__, "effect", effect)
        pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def effect(self) -> pulumi.Input[str]:
        """
        The effect of the taint. Valid values: `NO_SCHEDULE`, `NO_EXECUTE`, `PREFER_NO_SCHEDULE`.
        """
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: pulumi.Input[str]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key of the taint. Maximum length of 63.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the taint. Maximum length of 63.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NodeGroupUpdateConfigArgs:
    def __init__(__self__, *,
                 max_unavailable: Optional[pulumi.Input[int]] = None,
                 max_unavailable_percentage: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] max_unavailable: Desired max number of unavailable worker nodes during node group update.
        :param pulumi.Input[int] max_unavailable_percentage: Desired max percentage of unavailable worker nodes during node group update.
        """
        if max_unavailable is not None:
            pulumi.set(__self__, "max_unavailable", max_unavailable)
        if max_unavailable_percentage is not None:
            pulumi.set(__self__, "max_unavailable_percentage", max_unavailable_percentage)

    @property
    @pulumi.getter(name="maxUnavailable")
    def max_unavailable(self) -> Optional[pulumi.Input[int]]:
        """
        Desired max number of unavailable worker nodes during node group update.
        """
        return pulumi.get(self, "max_unavailable")

    @max_unavailable.setter
    def max_unavailable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_unavailable", value)

    @property
    @pulumi.getter(name="maxUnavailablePercentage")
    def max_unavailable_percentage(self) -> Optional[pulumi.Input[int]]:
        """
        Desired max percentage of unavailable worker nodes during node group update.
        """
        return pulumi.get(self, "max_unavailable_percentage")

    @max_unavailable_percentage.setter
    def max_unavailable_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_unavailable_percentage", value)


