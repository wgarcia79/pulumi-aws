# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkAclRuleArgs', 'NetworkAclRule']

@pulumi.input_type
class NetworkAclRuleArgs:
    def __init__(__self__, *,
                 network_acl_id: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 rule_action: pulumi.Input[str],
                 rule_number: pulumi.Input[int],
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[bool]] = None,
                 from_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[str]] = None,
                 icmp_type: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 to_port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a NetworkAclRule resource.
        :param pulumi.Input[str] network_acl_id: The ID of the network ACL.
        :param pulumi.Input[str] protocol: The protocol. A value of -1 means all protocols.
        :param pulumi.Input[str] rule_action: Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        :param pulumi.Input[int] rule_number: The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        :param pulumi.Input[str] cidr_block: The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        :param pulumi.Input[bool] egress: Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        :param pulumi.Input[int] from_port: The from port to match.
        :param pulumi.Input[str] icmp_code: ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] icmp_type: ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block to allow or deny.
        :param pulumi.Input[int] to_port: The to port to match.
        """
        pulumi.set(__self__, "network_acl_id", network_acl_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "rule_action", rule_action)
        pulumi.set(__self__, "rule_number", rule_number)
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if from_port is not None:
            pulumi.set(__self__, "from_port", from_port)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if ipv6_cidr_block is not None:
            pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if to_port is not None:
            pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Input[str]:
        """
        The ID of the network ACL.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_acl_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The protocol. A value of -1 means all protocols.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Input[str]:
        """
        Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> pulumi.Input[int]:
        """
        The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        """
        return pulumi.get(self, "rule_number")

    @rule_number.setter
    def rule_number(self, value: pulumi.Input[int]):
        pulumi.set(self, "rule_number", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> Optional[pulumi.Input[int]]:
        """
        The from port to match.
        """
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[str]]:
        """
        ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[str]]:
        """
        ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 CIDR block to allow or deny.
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> Optional[pulumi.Input[int]]:
        """
        The to port to match.
        """
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to_port", value)


@pulumi.input_type
class _NetworkAclRuleState:
    def __init__(__self__, *,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[bool]] = None,
                 from_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[str]] = None,
                 icmp_type: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_number: Optional[pulumi.Input[int]] = None,
                 to_port: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering NetworkAclRule resources.
        :param pulumi.Input[str] cidr_block: The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        :param pulumi.Input[bool] egress: Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        :param pulumi.Input[int] from_port: The from port to match.
        :param pulumi.Input[str] icmp_code: ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] icmp_type: ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block to allow or deny.
        :param pulumi.Input[str] network_acl_id: The ID of the network ACL.
        :param pulumi.Input[str] protocol: The protocol. A value of -1 means all protocols.
        :param pulumi.Input[str] rule_action: Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        :param pulumi.Input[int] rule_number: The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        :param pulumi.Input[int] to_port: The to port to match.
        """
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if egress is not None:
            pulumi.set(__self__, "egress", egress)
        if from_port is not None:
            pulumi.set(__self__, "from_port", from_port)
        if icmp_code is not None:
            pulumi.set(__self__, "icmp_code", icmp_code)
        if icmp_type is not None:
            pulumi.set(__self__, "icmp_type", icmp_type)
        if ipv6_cidr_block is not None:
            pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if network_acl_id is not None:
            pulumi.set(__self__, "network_acl_id", network_acl_id)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if rule_action is not None:
            pulumi.set(__self__, "rule_action", rule_action)
        if rule_number is not None:
            pulumi.set(__self__, "rule_number", rule_number)
        if to_port is not None:
            pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        """
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def egress(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        """
        return pulumi.get(self, "egress")

    @egress.setter
    def egress(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "egress", value)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> Optional[pulumi.Input[int]]:
        """
        The from port to match.
        """
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> Optional[pulumi.Input[str]]:
        """
        ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_code")

    @icmp_code.setter
    def icmp_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_code", value)

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> Optional[pulumi.Input[str]]:
        """
        ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_type")

    @icmp_type.setter
    def icmp_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_type", value)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 CIDR block to allow or deny.
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the network ACL.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol. A value of -1 means all protocols.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> Optional[pulumi.Input[int]]:
        """
        The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        """
        return pulumi.get(self, "rule_number")

    @rule_number.setter
    def rule_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_number", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> Optional[pulumi.Input[int]]:
        """
        The to port to match.
        """
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to_port", value)


class NetworkAclRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[bool]] = None,
                 from_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[str]] = None,
                 icmp_type: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_number: Optional[pulumi.Input[int]] = None,
                 to_port: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an entry (a rule) in a network ACL with the specified rule number.

        > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
        provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
        defined in-line. At this time you cannot use a Network ACL with in-line rules
        in conjunction with any Network ACL Rule resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar_network_acl = aws.ec2.NetworkAcl("barNetworkAcl", vpc_id=aws_vpc["foo"]["id"])
        bar_network_acl_rule = aws.ec2.NetworkAclRule("barNetworkAclRule",
            network_acl_id=bar_network_acl.id,
            rule_number=200,
            egress=False,
            protocol="tcp",
            rule_action="allow",
            cidr_block=aws_vpc["foo"]["cidr_block"],
            from_port=22,
            to_port=22)
        ```

        > **Note:** One of either `cidr_block` or `ipv6_cidr_block` is required.

        ## Import

        Individual rules can be imported using `NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS`, where `PROTOCOL` can be a decimal (e.g. 6) or string (e.g. tcp) value. If importing a rule previously provisioned by the provider, the `PROTOCOL` must be the input value used at creation time. For more information on protocol numbers and keywords, see herehttps://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml For example, import a network ACL Rule with an argument like thisconsole

        ```sh
         $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:tcp:false
        ```

         Or by the procotol's decimal valueconsole

        ```sh
         $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:6:false
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        :param pulumi.Input[bool] egress: Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        :param pulumi.Input[int] from_port: The from port to match.
        :param pulumi.Input[str] icmp_code: ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] icmp_type: ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block to allow or deny.
        :param pulumi.Input[str] network_acl_id: The ID of the network ACL.
        :param pulumi.Input[str] protocol: The protocol. A value of -1 means all protocols.
        :param pulumi.Input[str] rule_action: Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        :param pulumi.Input[int] rule_number: The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        :param pulumi.Input[int] to_port: The to port to match.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkAclRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an entry (a rule) in a network ACL with the specified rule number.

        > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
        provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
        defined in-line. At this time you cannot use a Network ACL with in-line rules
        in conjunction with any Network ACL Rule resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar_network_acl = aws.ec2.NetworkAcl("barNetworkAcl", vpc_id=aws_vpc["foo"]["id"])
        bar_network_acl_rule = aws.ec2.NetworkAclRule("barNetworkAclRule",
            network_acl_id=bar_network_acl.id,
            rule_number=200,
            egress=False,
            protocol="tcp",
            rule_action="allow",
            cidr_block=aws_vpc["foo"]["cidr_block"],
            from_port=22,
            to_port=22)
        ```

        > **Note:** One of either `cidr_block` or `ipv6_cidr_block` is required.

        ## Import

        Individual rules can be imported using `NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS`, where `PROTOCOL` can be a decimal (e.g. 6) or string (e.g. tcp) value. If importing a rule previously provisioned by the provider, the `PROTOCOL` must be the input value used at creation time. For more information on protocol numbers and keywords, see herehttps://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml For example, import a network ACL Rule with an argument like thisconsole

        ```sh
         $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:tcp:false
        ```

         Or by the procotol's decimal valueconsole

        ```sh
         $ pulumi import aws:ec2/networkAclRule:NetworkAclRule my_rule acl-7aaabd18:100:6:false
        ```

        :param str resource_name: The name of the resource.
        :param NetworkAclRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkAclRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 egress: Optional[pulumi.Input[bool]] = None,
                 from_port: Optional[pulumi.Input[int]] = None,
                 icmp_code: Optional[pulumi.Input[str]] = None,
                 icmp_type: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 rule_action: Optional[pulumi.Input[str]] = None,
                 rule_number: Optional[pulumi.Input[int]] = None,
                 to_port: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkAclRuleArgs.__new__(NetworkAclRuleArgs)

            __props__.__dict__["cidr_block"] = cidr_block
            __props__.__dict__["egress"] = egress
            __props__.__dict__["from_port"] = from_port
            __props__.__dict__["icmp_code"] = icmp_code
            __props__.__dict__["icmp_type"] = icmp_type
            __props__.__dict__["ipv6_cidr_block"] = ipv6_cidr_block
            if network_acl_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_acl_id'")
            __props__.__dict__["network_acl_id"] = network_acl_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if rule_action is None and not opts.urn:
                raise TypeError("Missing required property 'rule_action'")
            __props__.__dict__["rule_action"] = rule_action
            if rule_number is None and not opts.urn:
                raise TypeError("Missing required property 'rule_number'")
            __props__.__dict__["rule_number"] = rule_number
            __props__.__dict__["to_port"] = to_port
        super(NetworkAclRule, __self__).__init__(
            'aws:ec2/networkAclRule:NetworkAclRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            egress: Optional[pulumi.Input[bool]] = None,
            from_port: Optional[pulumi.Input[int]] = None,
            icmp_code: Optional[pulumi.Input[str]] = None,
            icmp_type: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            network_acl_id: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            rule_action: Optional[pulumi.Input[str]] = None,
            rule_number: Optional[pulumi.Input[int]] = None,
            to_port: Optional[pulumi.Input[int]] = None) -> 'NetworkAclRule':
        """
        Get an existing NetworkAclRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr_block: The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        :param pulumi.Input[bool] egress: Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        :param pulumi.Input[int] from_port: The from port to match.
        :param pulumi.Input[str] icmp_code: ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] icmp_type: ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 CIDR block to allow or deny.
        :param pulumi.Input[str] network_acl_id: The ID of the network ACL.
        :param pulumi.Input[str] protocol: The protocol. A value of -1 means all protocols.
        :param pulumi.Input[str] rule_action: Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        :param pulumi.Input[int] rule_number: The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        :param pulumi.Input[int] to_port: The to port to match.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkAclRuleState.__new__(_NetworkAclRuleState)

        __props__.__dict__["cidr_block"] = cidr_block
        __props__.__dict__["egress"] = egress
        __props__.__dict__["from_port"] = from_port
        __props__.__dict__["icmp_code"] = icmp_code
        __props__.__dict__["icmp_type"] = icmp_type
        __props__.__dict__["ipv6_cidr_block"] = ipv6_cidr_block
        __props__.__dict__["network_acl_id"] = network_acl_id
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["rule_action"] = rule_action
        __props__.__dict__["rule_number"] = rule_number
        __props__.__dict__["to_port"] = to_port
        return NetworkAclRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def egress(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        """
        return pulumi.get(self, "egress")

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> pulumi.Output[Optional[int]]:
        """
        The from port to match.
        """
        return pulumi.get(self, "from_port")

    @property
    @pulumi.getter(name="icmpCode")
    def icmp_code(self) -> pulumi.Output[Optional[str]]:
        """
        ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_code")

    @property
    @pulumi.getter(name="icmpType")
    def icmp_type(self) -> pulumi.Output[Optional[str]]:
        """
        ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        """
        return pulumi.get(self, "icmp_type")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv6 CIDR block to allow or deny.
        """
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Output[str]:
        """
        The ID of the network ACL.
        """
        return pulumi.get(self, "network_acl_id")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol. A value of -1 means all protocols.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Output[str]:
        """
        Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        """
        return pulumi.get(self, "rule_action")

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> pulumi.Output[int]:
        """
        The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        """
        return pulumi.get(self, "rule_number")

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> pulumi.Output[Optional[int]]:
        """
        The to port to match.
        """
        return pulumi.get(self, "to_port")

