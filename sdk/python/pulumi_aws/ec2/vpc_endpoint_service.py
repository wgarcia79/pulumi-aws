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

__all__ = ['VpcEndpointServiceArgs', 'VpcEndpointService']

@pulumi.input_type
class VpcEndpointServiceArgs:
    def __init__(__self__, *,
                 acceptance_required: pulumi.Input[bool],
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 gateway_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 network_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_dns_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VpcEndpointService resource.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateway_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[str] private_dns_name: The private DNS name for the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "acceptance_required", acceptance_required)
        if allowed_principals is not None:
            pulumi.set(__self__, "allowed_principals", allowed_principals)
        if gateway_load_balancer_arns is not None:
            pulumi.set(__self__, "gateway_load_balancer_arns", gateway_load_balancer_arns)
        if network_load_balancer_arns is not None:
            pulumi.set(__self__, "network_load_balancer_arns", network_load_balancer_arns)
        if private_dns_name is not None:
            pulumi.set(__self__, "private_dns_name", private_dns_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> pulumi.Input[bool]:
        """
        Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        """
        return pulumi.get(self, "acceptance_required")

    @acceptance_required.setter
    def acceptance_required(self, value: pulumi.Input[bool]):
        pulumi.set(self, "acceptance_required", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ARNs of one or more principals allowed to discover the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter(name="gatewayLoadBalancerArns")
    def gateway_load_balancer_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "gateway_load_balancer_arns")

    @gateway_load_balancer_arns.setter
    def gateway_load_balancer_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "gateway_load_balancer_arns", value)

    @property
    @pulumi.getter(name="networkLoadBalancerArns")
    def network_load_balancer_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "network_load_balancer_arns")

    @network_load_balancer_arns.setter
    def network_load_balancer_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "network_load_balancer_arns", value)

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The private DNS name for the service.
        """
        return pulumi.get(self, "private_dns_name")

    @private_dns_name.setter
    def private_dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_dns_name", value)

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
class _VpcEndpointServiceState:
    def __init__(__self__, *,
                 acceptance_required: Optional[pulumi.Input[bool]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 base_endpoint_dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 gateway_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 manages_vpc_endpoints: Optional[pulumi.Input[bool]] = None,
                 network_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_dns_name: Optional[pulumi.Input[str]] = None,
                 private_dns_name_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['VpcEndpointServicePrivateDnsNameConfigurationArgs']]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 service_type: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointService resources.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the VPC endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] availability_zones: The Availability Zones in which the service is available.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] base_endpoint_dns_names: The DNS names for the service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateway_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        :param pulumi.Input[bool] manages_vpc_endpoints: Whether or not the service manages its VPC endpoints - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[str] private_dns_name: The private DNS name for the service.
        :param pulumi.Input[Sequence[pulumi.Input['VpcEndpointServicePrivateDnsNameConfigurationArgs']]] private_dns_name_configurations: List of objects containing information about the endpoint service private DNS name configuration.
        :param pulumi.Input[str] service_name: The service name.
        :param pulumi.Input[str] service_type: The service type, `Gateway` or `Interface`.
        :param pulumi.Input[str] state: Verification state of the VPC endpoint service. Consumers of the endpoint service can use the private name only when the state is `verified`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        if acceptance_required is not None:
            pulumi.set(__self__, "acceptance_required", acceptance_required)
        if allowed_principals is not None:
            pulumi.set(__self__, "allowed_principals", allowed_principals)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if base_endpoint_dns_names is not None:
            pulumi.set(__self__, "base_endpoint_dns_names", base_endpoint_dns_names)
        if gateway_load_balancer_arns is not None:
            pulumi.set(__self__, "gateway_load_balancer_arns", gateway_load_balancer_arns)
        if manages_vpc_endpoints is not None:
            pulumi.set(__self__, "manages_vpc_endpoints", manages_vpc_endpoints)
        if network_load_balancer_arns is not None:
            pulumi.set(__self__, "network_load_balancer_arns", network_load_balancer_arns)
        if private_dns_name is not None:
            pulumi.set(__self__, "private_dns_name", private_dns_name)
        if private_dns_name_configurations is not None:
            pulumi.set(__self__, "private_dns_name_configurations", private_dns_name_configurations)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        """
        return pulumi.get(self, "acceptance_required")

    @acceptance_required.setter
    def acceptance_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "acceptance_required", value)

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ARNs of one or more principals allowed to discover the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @allowed_principals.setter
    def allowed_principals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_principals", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the VPC endpoint service.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The Availability Zones in which the service is available.
        """
        return pulumi.get(self, "availability_zones")

    @availability_zones.setter
    def availability_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "availability_zones", value)

    @property
    @pulumi.getter(name="baseEndpointDnsNames")
    def base_endpoint_dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The DNS names for the service.
        """
        return pulumi.get(self, "base_endpoint_dns_names")

    @base_endpoint_dns_names.setter
    def base_endpoint_dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "base_endpoint_dns_names", value)

    @property
    @pulumi.getter(name="gatewayLoadBalancerArns")
    def gateway_load_balancer_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "gateway_load_balancer_arns")

    @gateway_load_balancer_arns.setter
    def gateway_load_balancer_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "gateway_load_balancer_arns", value)

    @property
    @pulumi.getter(name="managesVpcEndpoints")
    def manages_vpc_endpoints(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the service manages its VPC endpoints - `true` or `false`.
        """
        return pulumi.get(self, "manages_vpc_endpoints")

    @manages_vpc_endpoints.setter
    def manages_vpc_endpoints(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "manages_vpc_endpoints", value)

    @property
    @pulumi.getter(name="networkLoadBalancerArns")
    def network_load_balancer_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "network_load_balancer_arns")

    @network_load_balancer_arns.setter
    def network_load_balancer_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "network_load_balancer_arns", value)

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> Optional[pulumi.Input[str]]:
        """
        The private DNS name for the service.
        """
        return pulumi.get(self, "private_dns_name")

    @private_dns_name.setter
    def private_dns_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_dns_name", value)

    @property
    @pulumi.getter(name="privateDnsNameConfigurations")
    def private_dns_name_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcEndpointServicePrivateDnsNameConfigurationArgs']]]]:
        """
        List of objects containing information about the endpoint service private DNS name configuration.
        """
        return pulumi.get(self, "private_dns_name_configurations")

    @private_dns_name_configurations.setter
    def private_dns_name_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcEndpointServicePrivateDnsNameConfigurationArgs']]]]):
        pulumi.set(self, "private_dns_name_configurations", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input[str]]:
        """
        The service type, `Gateway` or `Interface`.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Verification state of the VPC endpoint service. Consumers of the endpoint service can use the private name only when the state is `verified`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

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


class VpcEndpointService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptance_required: Optional[pulumi.Input[bool]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 gateway_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 network_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_dns_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a VPC Endpoint Service resource.
        Service consumers can create an _Interface_ VPC Endpoint to connect to the service.

        > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
        both a standalone VPC Endpoint Service Allowed Principal resource
        and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
        a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
        and will overwrite the association.

        ## Example Usage
        ### Network Load Balancers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["example"]["arn"]])
        ```
        ### Gateway Load Balancers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            gateway_load_balancer_arns=[aws_lb["example"]["arn"]])
        ```

        ## Import

        VPC Endpoint Services can be imported using the `VPC endpoint service id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcEndpointService:VpcEndpointService foo vpce-svc-0f97a19d3fa8220bc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateway_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[str] private_dns_name: The private DNS name for the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Endpoint Service resource.
        Service consumers can create an _Interface_ VPC Endpoint to connect to the service.

        > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
        both a standalone VPC Endpoint Service Allowed Principal resource
        and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
        a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
        and will overwrite the association.

        ## Example Usage
        ### Network Load Balancers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["example"]["arn"]])
        ```
        ### Gateway Load Balancers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpcEndpointService("example",
            acceptance_required=False,
            gateway_load_balancer_arns=[aws_lb["example"]["arn"]])
        ```

        ## Import

        VPC Endpoint Services can be imported using the `VPC endpoint service id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcEndpointService:VpcEndpointService foo vpce-svc-0f97a19d3fa8220bc
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acceptance_required: Optional[pulumi.Input[bool]] = None,
                 allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 gateway_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 network_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_dns_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = VpcEndpointServiceArgs.__new__(VpcEndpointServiceArgs)

            if acceptance_required is None and not opts.urn:
                raise TypeError("Missing required property 'acceptance_required'")
            __props__.__dict__["acceptance_required"] = acceptance_required
            __props__.__dict__["allowed_principals"] = allowed_principals
            __props__.__dict__["gateway_load_balancer_arns"] = gateway_load_balancer_arns
            __props__.__dict__["network_load_balancer_arns"] = network_load_balancer_arns
            __props__.__dict__["private_dns_name"] = private_dns_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            __props__.__dict__["arn"] = None
            __props__.__dict__["availability_zones"] = None
            __props__.__dict__["base_endpoint_dns_names"] = None
            __props__.__dict__["manages_vpc_endpoints"] = None
            __props__.__dict__["private_dns_name_configurations"] = None
            __props__.__dict__["service_name"] = None
            __props__.__dict__["service_type"] = None
            __props__.__dict__["state"] = None
        super(VpcEndpointService, __self__).__init__(
            'aws:ec2/vpcEndpointService:VpcEndpointService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acceptance_required: Optional[pulumi.Input[bool]] = None,
            allowed_principals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            base_endpoint_dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            gateway_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            manages_vpc_endpoints: Optional[pulumi.Input[bool]] = None,
            network_load_balancer_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            private_dns_name: Optional[pulumi.Input[str]] = None,
            private_dns_name_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServicePrivateDnsNameConfigurationArgs']]]]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            service_type: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VpcEndpointService':
        """
        Get an existing VpcEndpointService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] acceptance_required: Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_principals: The ARNs of one or more principals allowed to discover the endpoint service.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the VPC endpoint service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] availability_zones: The Availability Zones in which the service is available.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] base_endpoint_dns_names: The DNS names for the service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] gateway_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        :param pulumi.Input[bool] manages_vpc_endpoints: Whether or not the service manages its VPC endpoints - `true` or `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] network_load_balancer_arns: Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        :param pulumi.Input[str] private_dns_name: The private DNS name for the service.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcEndpointServicePrivateDnsNameConfigurationArgs']]]] private_dns_name_configurations: List of objects containing information about the endpoint service private DNS name configuration.
        :param pulumi.Input[str] service_name: The service name.
        :param pulumi.Input[str] service_type: The service type, `Gateway` or `Interface`.
        :param pulumi.Input[str] state: Verification state of the VPC endpoint service. Consumers of the endpoint service can use the private name only when the state is `verified`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointServiceState.__new__(_VpcEndpointServiceState)

        __props__.__dict__["acceptance_required"] = acceptance_required
        __props__.__dict__["allowed_principals"] = allowed_principals
        __props__.__dict__["arn"] = arn
        __props__.__dict__["availability_zones"] = availability_zones
        __props__.__dict__["base_endpoint_dns_names"] = base_endpoint_dns_names
        __props__.__dict__["gateway_load_balancer_arns"] = gateway_load_balancer_arns
        __props__.__dict__["manages_vpc_endpoints"] = manages_vpc_endpoints
        __props__.__dict__["network_load_balancer_arns"] = network_load_balancer_arns
        __props__.__dict__["private_dns_name"] = private_dns_name
        __props__.__dict__["private_dns_name_configurations"] = private_dns_name_configurations
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["service_type"] = service_type
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return VpcEndpointService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> pulumi.Output[bool]:
        """
        Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
        """
        return pulumi.get(self, "acceptance_required")

    @property
    @pulumi.getter(name="allowedPrincipals")
    def allowed_principals(self) -> pulumi.Output[Sequence[str]]:
        """
        The ARNs of one or more principals allowed to discover the endpoint service.
        """
        return pulumi.get(self, "allowed_principals")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the VPC endpoint service.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> pulumi.Output[Sequence[str]]:
        """
        The Availability Zones in which the service is available.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="baseEndpointDnsNames")
    def base_endpoint_dns_names(self) -> pulumi.Output[Sequence[str]]:
        """
        The DNS names for the service.
        """
        return pulumi.get(self, "base_endpoint_dns_names")

    @property
    @pulumi.getter(name="gatewayLoadBalancerArns")
    def gateway_load_balancer_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Amazon Resource Names (ARNs) of one or more Gateway Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "gateway_load_balancer_arns")

    @property
    @pulumi.getter(name="managesVpcEndpoints")
    def manages_vpc_endpoints(self) -> pulumi.Output[bool]:
        """
        Whether or not the service manages its VPC endpoints - `true` or `false`.
        """
        return pulumi.get(self, "manages_vpc_endpoints")

    @property
    @pulumi.getter(name="networkLoadBalancerArns")
    def network_load_balancer_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Amazon Resource Names (ARNs) of one or more Network Load Balancers for the endpoint service.
        """
        return pulumi.get(self, "network_load_balancer_arns")

    @property
    @pulumi.getter(name="privateDnsName")
    def private_dns_name(self) -> pulumi.Output[str]:
        """
        The private DNS name for the service.
        """
        return pulumi.get(self, "private_dns_name")

    @property
    @pulumi.getter(name="privateDnsNameConfigurations")
    def private_dns_name_configurations(self) -> pulumi.Output[Sequence['outputs.VpcEndpointServicePrivateDnsNameConfiguration']]:
        """
        List of objects containing information about the endpoint service private DNS name configuration.
        """
        return pulumi.get(self, "private_dns_name_configurations")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The service name.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> pulumi.Output[str]:
        """
        The service type, `Gateway` or `Interface`.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Verification state of the VPC endpoint service. Consumers of the endpoint service can use the private name only when the state is `verified`.
        """
        return pulumi.get(self, "state")

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

