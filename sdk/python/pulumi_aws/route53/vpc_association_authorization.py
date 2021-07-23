# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcAssociationAuthorizationArgs', 'VpcAssociationAuthorization']

@pulumi.input_type
class VpcAssociationAuthorizationArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 vpc_region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcAssociationAuthorization resource.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone_id", zone_id)
        if vpc_region is not None:
            pulumi.set(__self__, "vpc_region", vpc_region)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The VPC to authorize for association with the private hosted zone.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC's region. Defaults to the region of the AWS provider.
        """
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_region", value)


@pulumi.input_type
class _VpcAssociationAuthorizationState:
    def __init__(__self__, *,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcAssociationAuthorization resources.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_region is not None:
            pulumi.set(__self__, "vpc_region", vpc_region)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC to authorize for association with the private hosted zone.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC's region. Defaults to the region of the AWS provider.
        """
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_region", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class VpcAssociationAuthorization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Authorizes a VPC in a different account to be associated with a local Route53 Hosted Zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        alternate = pulumi.providers.Aws("alternate")
        example_vpc = aws.ec2.Vpc("exampleVpc",
            cidr_block="10.6.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True)
        example_zone = aws.route53.Zone("exampleZone", vpcs=[aws.route53.ZoneVpcArgs(
            vpc_id=example_vpc.id,
        )])
        alternate_vpc = aws.ec2.Vpc("alternateVpc",
            cidr_block="10.7.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        example_vpc_association_authorization = aws.route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization",
            vpc_id=alternate_vpc.id,
            zone_id=example_zone.id)
        example_zone_association = aws.route53.ZoneAssociation("exampleZoneAssociation",
            vpc_id=example_vpc_association_authorization.vpc_id,
            zone_id=example_vpc_association_authorization.zone_id,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        ```

        ## Import

        Route 53 VPC Association Authorizations can be imported via the Hosted Zone ID and VPC ID, separated by a colon (`:`), e.g.

        ```sh
         $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcAssociationAuthorizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Authorizes a VPC in a different account to be associated with a local Route53 Hosted Zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        alternate = pulumi.providers.Aws("alternate")
        example_vpc = aws.ec2.Vpc("exampleVpc",
            cidr_block="10.6.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True)
        example_zone = aws.route53.Zone("exampleZone", vpcs=[aws.route53.ZoneVpcArgs(
            vpc_id=example_vpc.id,
        )])
        alternate_vpc = aws.ec2.Vpc("alternateVpc",
            cidr_block="10.7.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        example_vpc_association_authorization = aws.route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization",
            vpc_id=alternate_vpc.id,
            zone_id=example_zone.id)
        example_zone_association = aws.route53.ZoneAssociation("exampleZoneAssociation",
            vpc_id=example_vpc_association_authorization.vpc_id,
            zone_id=example_vpc_association_authorization.zone_id,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        ```

        ## Import

        Route 53 VPC Association Authorizations can be imported via the Hosted Zone ID and VPC ID, separated by a colon (`:`), e.g.

        ```sh
         $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
        ```

        :param str resource_name: The name of the resource.
        :param VpcAssociationAuthorizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcAssociationAuthorizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcAssociationAuthorizationArgs.__new__(VpcAssociationAuthorizationArgs)

            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vpc_region"] = vpc_region
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(VpcAssociationAuthorization, __self__).__init__(
            'aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_region: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'VpcAssociationAuthorization':
        """
        Get an existing VpcAssociationAuthorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcAssociationAuthorizationState.__new__(_VpcAssociationAuthorizationState)

        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_region"] = vpc_region
        __props__.__dict__["zone_id"] = zone_id
        return VpcAssociationAuthorization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC to authorize for association with the private hosted zone.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> pulumi.Output[str]:
        """
        The VPC's region. Defaults to the region of the AWS provider.
        """
        return pulumi.get(self, "vpc_region")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        return pulumi.get(self, "zone_id")

