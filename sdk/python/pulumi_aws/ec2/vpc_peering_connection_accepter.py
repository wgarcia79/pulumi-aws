# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VpcPeeringConnectionAccepterArgs', 'VpcPeeringConnectionAccepter']

@pulumi.input_type
class VpcPeeringConnectionAccepterArgs:
    def __init__(__self__, *,
                 vpc_peering_connection_id: pulumi.Input[str],
                 accepter: Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']] = None,
                 auto_accept: Optional[pulumi.Input[bool]] = None,
                 requester: Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VpcPeeringConnectionAccepter resource.
        :param pulumi.Input[str] vpc_peering_connection_id: The VPC Peering Connection ID to manage.
        :param pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs'] accepter: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        :param pulumi.Input[bool] auto_accept: Whether or not to accept the peering request. Defaults to `false`.
        :param pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs'] requester: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "vpc_peering_connection_id", vpc_peering_connection_id)
        if accepter is not None:
            pulumi.set(__self__, "accepter", accepter)
        if auto_accept is not None:
            pulumi.set(__self__, "auto_accept", auto_accept)
        if requester is not None:
            pulumi.set(__self__, "requester", requester)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Input[str]:
        """
        The VPC Peering Connection ID to manage.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    @vpc_peering_connection_id.setter
    def vpc_peering_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_peering_connection_id", value)

    @property
    @pulumi.getter
    def accepter(self) -> Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']]:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        """
        return pulumi.get(self, "accepter")

    @accepter.setter
    def accepter(self, value: Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']]):
        pulumi.set(self, "accepter", value)

    @property
    @pulumi.getter(name="autoAccept")
    def auto_accept(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to accept the peering request. Defaults to `false`.
        """
        return pulumi.get(self, "auto_accept")

    @auto_accept.setter
    def auto_accept(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_accept", value)

    @property
    @pulumi.getter
    def requester(self) -> Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']]:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        """
        return pulumi.get(self, "requester")

    @requester.setter
    def requester(self, value: Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']]):
        pulumi.set(self, "requester", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _VpcPeeringConnectionAccepterState:
    def __init__(__self__, *,
                 accept_status: Optional[pulumi.Input[str]] = None,
                 accepter: Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']] = None,
                 auto_accept: Optional[pulumi.Input[bool]] = None,
                 peer_owner_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 requester: Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPeeringConnectionAccepter resources.
        :param pulumi.Input[str] accept_status: The status of the VPC Peering Connection request.
        :param pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs'] accepter: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        :param pulumi.Input[bool] auto_accept: Whether or not to accept the peering request. Defaults to `false`.
        :param pulumi.Input[str] peer_owner_id: The AWS account ID of the owner of the requester VPC.
        :param pulumi.Input[str] peer_region: The region of the accepter VPC.
        :param pulumi.Input[str] peer_vpc_id: The ID of the requester VPC.
        :param pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs'] requester: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] vpc_id: The ID of the accepter VPC.
        :param pulumi.Input[str] vpc_peering_connection_id: The VPC Peering Connection ID to manage.
        """
        if accept_status is not None:
            pulumi.set(__self__, "accept_status", accept_status)
        if accepter is not None:
            pulumi.set(__self__, "accepter", accepter)
        if auto_accept is not None:
            pulumi.set(__self__, "auto_accept", auto_accept)
        if peer_owner_id is not None:
            pulumi.set(__self__, "peer_owner_id", peer_owner_id)
        if peer_region is not None:
            pulumi.set(__self__, "peer_region", peer_region)
        if peer_vpc_id is not None:
            pulumi.set(__self__, "peer_vpc_id", peer_vpc_id)
        if requester is not None:
            pulumi.set(__self__, "requester", requester)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_peering_connection_id is not None:
            pulumi.set(__self__, "vpc_peering_connection_id", vpc_peering_connection_id)

    @property
    @pulumi.getter(name="acceptStatus")
    def accept_status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPC Peering Connection request.
        """
        return pulumi.get(self, "accept_status")

    @accept_status.setter
    def accept_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accept_status", value)

    @property
    @pulumi.getter
    def accepter(self) -> Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']]:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        """
        return pulumi.get(self, "accepter")

    @accepter.setter
    def accepter(self, value: Optional[pulumi.Input['VpcPeeringConnectionAccepterAccepterArgs']]):
        pulumi.set(self, "accepter", value)

    @property
    @pulumi.getter(name="autoAccept")
    def auto_accept(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to accept the peering request. Defaults to `false`.
        """
        return pulumi.get(self, "auto_accept")

    @auto_accept.setter
    def auto_accept(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_accept", value)

    @property
    @pulumi.getter(name="peerOwnerId")
    def peer_owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS account ID of the owner of the requester VPC.
        """
        return pulumi.get(self, "peer_owner_id")

    @peer_owner_id.setter
    def peer_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_owner_id", value)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the accepter VPC.
        """
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_region", value)

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the requester VPC.
        """
        return pulumi.get(self, "peer_vpc_id")

    @peer_vpc_id.setter
    def peer_vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_vpc_id", value)

    @property
    @pulumi.getter
    def requester(self) -> Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']]:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        """
        return pulumi.get(self, "requester")

    @requester.setter
    def requester(self, value: Optional[pulumi.Input['VpcPeeringConnectionAccepterRequesterArgs']]):
        pulumi.set(self, "requester", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the accepter VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC Peering Connection ID to manage.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    @vpc_peering_connection_id.setter
    def vpc_peering_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_peering_connection_id", value)


class VpcPeeringConnectionAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterAccepterArgs']]] = None,
                 auto_accept: Optional[pulumi.Input[bool]] = None,
                 requester: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterRequesterArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage the accepter's side of a VPC Peering Connection.

        When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
        VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
        accepter's account.
        The requester can use the `ec2.VpcPeeringConnection` resource to manage its side of the connection
        and the accepter can use the `ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
        connection into management.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        peer = pulumi.providers.Aws("peer", region="us-west-2")
        # Accepter's credentials.
        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        peer_vpc = aws.ec2.Vpc("peerVpc", cidr_block="10.1.0.0/16",
        opts=pulumi.ResourceOptions(provider=aws["peer"]))
        peer_caller_identity = aws.get_caller_identity()
        # Requester's side of the connection.
        peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
            vpc_id=main.id,
            peer_vpc_id=peer_vpc.id,
            peer_owner_id=peer_caller_identity.account_id,
            peer_region="us-west-2",
            auto_accept=False,
            tags={
                "Side": "Requester",
            })
        # Accepter's side of the connection.
        peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
            vpc_peering_connection_id=peer_vpc_peering_connection.id,
            auto_accept=True,
            tags={
                "Side": "Accepter",
            },
            opts=pulumi.ResourceOptions(provider=aws["peer"]))
        ```

        ## Import

        VPC Peering Connection Accepters can be imported by using the Peering Connection ID, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
        ```

         Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the provider configuration on an imported resource, this provder will always show a difference. To workaround this behavior, either omit the argument from the configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. terraform resource "aws_vpc_peering_connection_accepter" "example" {

        # ... other configuration ...

        # There is no AWS EC2 API for reading auto_accept

         lifecycle {

         ignore_changes = [auto_accept]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterAccepterArgs']] accepter: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        :param pulumi.Input[bool] auto_accept: Whether or not to accept the peering request. Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterRequesterArgs']] requester: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] vpc_peering_connection_id: The VPC Peering Connection ID to manage.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPeeringConnectionAccepterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage the accepter's side of a VPC Peering Connection.

        When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
        VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
        accepter's account.
        The requester can use the `ec2.VpcPeeringConnection` resource to manage its side of the connection
        and the accepter can use the `ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
        connection into management.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        peer = pulumi.providers.Aws("peer", region="us-west-2")
        # Accepter's credentials.
        main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
        peer_vpc = aws.ec2.Vpc("peerVpc", cidr_block="10.1.0.0/16",
        opts=pulumi.ResourceOptions(provider=aws["peer"]))
        peer_caller_identity = aws.get_caller_identity()
        # Requester's side of the connection.
        peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
            vpc_id=main.id,
            peer_vpc_id=peer_vpc.id,
            peer_owner_id=peer_caller_identity.account_id,
            peer_region="us-west-2",
            auto_accept=False,
            tags={
                "Side": "Requester",
            })
        # Accepter's side of the connection.
        peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
            vpc_peering_connection_id=peer_vpc_peering_connection.id,
            auto_accept=True,
            tags={
                "Side": "Accepter",
            },
            opts=pulumi.ResourceOptions(provider=aws["peer"]))
        ```

        ## Import

        VPC Peering Connection Accepters can be imported by using the Peering Connection ID, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
        ```

         Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the provider configuration on an imported resource, this provder will always show a difference. To workaround this behavior, either omit the argument from the configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. terraform resource "aws_vpc_peering_connection_accepter" "example" {

        # ... other configuration ...

        # There is no AWS EC2 API for reading auto_accept

         lifecycle {

         ignore_changes = [auto_accept]

         } }

        :param str resource_name: The name of the resource.
        :param VpcPeeringConnectionAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcPeeringConnectionAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterAccepterArgs']]] = None,
                 auto_accept: Optional[pulumi.Input[bool]] = None,
                 requester: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterRequesterArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPeeringConnectionAccepterArgs.__new__(VpcPeeringConnectionAccepterArgs)

            __props__.__dict__["accepter"] = accepter
            __props__.__dict__["auto_accept"] = auto_accept
            __props__.__dict__["requester"] = requester
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            if vpc_peering_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_peering_connection_id'")
            __props__.__dict__["vpc_peering_connection_id"] = vpc_peering_connection_id
            __props__.__dict__["accept_status"] = None
            __props__.__dict__["peer_owner_id"] = None
            __props__.__dict__["peer_region"] = None
            __props__.__dict__["peer_vpc_id"] = None
            __props__.__dict__["vpc_id"] = None
        super(VpcPeeringConnectionAccepter, __self__).__init__(
            'aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_status: Optional[pulumi.Input[str]] = None,
            accepter: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterAccepterArgs']]] = None,
            auto_accept: Optional[pulumi.Input[bool]] = None,
            peer_owner_id: Optional[pulumi.Input[str]] = None,
            peer_region: Optional[pulumi.Input[str]] = None,
            peer_vpc_id: Optional[pulumi.Input[str]] = None,
            requester: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterRequesterArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_peering_connection_id: Optional[pulumi.Input[str]] = None) -> 'VpcPeeringConnectionAccepter':
        """
        Get an existing VpcPeeringConnectionAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_status: The status of the VPC Peering Connection request.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterAccepterArgs']] accepter: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        :param pulumi.Input[bool] auto_accept: Whether or not to accept the peering request. Defaults to `false`.
        :param pulumi.Input[str] peer_owner_id: The AWS account ID of the owner of the requester VPC.
        :param pulumi.Input[str] peer_region: The region of the accepter VPC.
        :param pulumi.Input[str] peer_vpc_id: The ID of the requester VPC.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectionAccepterRequesterArgs']] requester: A configuration block that describes [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] vpc_id: The ID of the accepter VPC.
        :param pulumi.Input[str] vpc_peering_connection_id: The VPC Peering Connection ID to manage.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPeeringConnectionAccepterState.__new__(_VpcPeeringConnectionAccepterState)

        __props__.__dict__["accept_status"] = accept_status
        __props__.__dict__["accepter"] = accepter
        __props__.__dict__["auto_accept"] = auto_accept
        __props__.__dict__["peer_owner_id"] = peer_owner_id
        __props__.__dict__["peer_region"] = peer_region
        __props__.__dict__["peer_vpc_id"] = peer_vpc_id
        __props__.__dict__["requester"] = requester
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_peering_connection_id"] = vpc_peering_connection_id
        return VpcPeeringConnectionAccepter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptStatus")
    def accept_status(self) -> pulumi.Output[str]:
        """
        The status of the VPC Peering Connection request.
        """
        return pulumi.get(self, "accept_status")

    @property
    @pulumi.getter
    def accepter(self) -> pulumi.Output['outputs.VpcPeeringConnectionAccepterAccepter']:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
        """
        return pulumi.get(self, "accepter")

    @property
    @pulumi.getter(name="autoAccept")
    def auto_accept(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to accept the peering request. Defaults to `false`.
        """
        return pulumi.get(self, "auto_accept")

    @property
    @pulumi.getter(name="peerOwnerId")
    def peer_owner_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID of the owner of the requester VPC.
        """
        return pulumi.get(self, "peer_owner_id")

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Output[str]:
        """
        The region of the accepter VPC.
        """
        return pulumi.get(self, "peer_region")

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the requester VPC.
        """
        return pulumi.get(self, "peer_vpc_id")

    @property
    @pulumi.getter
    def requester(self) -> pulumi.Output['outputs.VpcPeeringConnectionAccepterRequester']:
        """
        A configuration block that describes [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
        """
        return pulumi.get(self, "requester")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the accepter VPC.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Output[str]:
        """
        The VPC Peering Connection ID to manage.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

