# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceResult:
    """
    A collection of values returned by getInstance.
    """
    def __init__(__self__, ami=None, arn=None, associate_public_ip_address=None, availability_zone=None, credit_specifications=None, disable_api_termination=None, ebs_block_devices=None, ebs_optimized=None, ephemeral_block_devices=None, filters=None, get_password_data=None, get_user_data=None, host_id=None, iam_instance_profile=None, id=None, instance_id=None, instance_state=None, instance_tags=None, instance_type=None, key_name=None, monitoring=None, network_interface_id=None, password_data=None, placement_group=None, private_dns=None, private_ip=None, public_dns=None, public_ip=None, root_block_devices=None, security_groups=None, source_dest_check=None, subnet_id=None, tags=None, tenancy=None, user_data=None, user_data_base64=None, vpc_security_group_ids=None):
        if ami and not isinstance(ami, str):
            raise TypeError("Expected argument 'ami' to be a str")
        __self__.ami = ami
        """
        The ID of the AMI used to launch the instance.
        """
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN of the instance.
        """
        if associate_public_ip_address and not isinstance(associate_public_ip_address, bool):
            raise TypeError("Expected argument 'associate_public_ip_address' to be a bool")
        __self__.associate_public_ip_address = associate_public_ip_address
        """
        Whether or not the Instance is associated with a public IP address or not (Boolean).
        """
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        """
        The availability zone of the Instance.
        """
        if credit_specifications and not isinstance(credit_specifications, list):
            raise TypeError("Expected argument 'credit_specifications' to be a list")
        __self__.credit_specifications = credit_specifications
        """
        The credit specification of the Instance.
        """
        if disable_api_termination and not isinstance(disable_api_termination, bool):
            raise TypeError("Expected argument 'disable_api_termination' to be a bool")
        __self__.disable_api_termination = disable_api_termination
        if ebs_block_devices and not isinstance(ebs_block_devices, list):
            raise TypeError("Expected argument 'ebs_block_devices' to be a list")
        __self__.ebs_block_devices = ebs_block_devices
        """
        The EBS block device mappings of the Instance.
        """
        if ebs_optimized and not isinstance(ebs_optimized, bool):
            raise TypeError("Expected argument 'ebs_optimized' to be a bool")
        __self__.ebs_optimized = ebs_optimized
        """
        Whether the Instance is EBS optimized or not (Boolean).
        """
        if ephemeral_block_devices and not isinstance(ephemeral_block_devices, list):
            raise TypeError("Expected argument 'ephemeral_block_devices' to be a list")
        __self__.ephemeral_block_devices = ephemeral_block_devices
        """
        The ephemeral block device mappings of the Instance.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if get_password_data and not isinstance(get_password_data, bool):
            raise TypeError("Expected argument 'get_password_data' to be a bool")
        __self__.get_password_data = get_password_data
        if get_user_data and not isinstance(get_user_data, bool):
            raise TypeError("Expected argument 'get_user_data' to be a bool")
        __self__.get_user_data = get_user_data
        if host_id and not isinstance(host_id, str):
            raise TypeError("Expected argument 'host_id' to be a str")
        __self__.host_id = host_id
        """
        The Id of the dedicated host the instance will be assigned to.
        """
        if iam_instance_profile and not isinstance(iam_instance_profile, str):
            raise TypeError("Expected argument 'iam_instance_profile' to be a str")
        __self__.iam_instance_profile = iam_instance_profile
        """
        The name of the instance profile associated with the Instance.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        if instance_state and not isinstance(instance_state, str):
            raise TypeError("Expected argument 'instance_state' to be a str")
        __self__.instance_state = instance_state
        """
        The state of the instance. One of: `pending`, `running`, `shutting-down`, `terminated`, `stopping`, `stopped`. See [Instance Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html) for more information.
        """
        if instance_tags and not isinstance(instance_tags, dict):
            raise TypeError("Expected argument 'instance_tags' to be a dict")
        __self__.instance_tags = instance_tags
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        __self__.instance_type = instance_type
        """
        The type of the Instance.
        """
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        __self__.key_name = key_name
        """
        The key name of the Instance.
        """
        if monitoring and not isinstance(monitoring, bool):
            raise TypeError("Expected argument 'monitoring' to be a bool")
        __self__.monitoring = monitoring
        """
        Whether detailed monitoring is enabled or disabled for the Instance (Boolean).
        """
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        __self__.network_interface_id = network_interface_id
        """
        The ID of the network interface that was created with the Instance.
        """
        if password_data and not isinstance(password_data, str):
            raise TypeError("Expected argument 'password_data' to be a str")
        __self__.password_data = password_data
        """
        Base-64 encoded encrypted password data for the instance.
        Useful for getting the administrator password for instances running Microsoft Windows.
        This attribute is only exported if `get_password_data` is true.
        See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
        """
        if placement_group and not isinstance(placement_group, str):
            raise TypeError("Expected argument 'placement_group' to be a str")
        __self__.placement_group = placement_group
        """
        The placement group of the Instance.
        """
        if private_dns and not isinstance(private_dns, str):
            raise TypeError("Expected argument 'private_dns' to be a str")
        __self__.private_dns = private_dns
        """
        The private DNS name assigned to the Instance. Can only be
        used inside the Amazon EC2, and only available if you've enabled DNS hostnames
        for your VPC.
        """
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        __self__.private_ip = private_ip
        """
        The private IP address assigned to the Instance.
        """
        if public_dns and not isinstance(public_dns, str):
            raise TypeError("Expected argument 'public_dns' to be a str")
        __self__.public_dns = public_dns
        """
        The public DNS name assigned to the Instance. For EC2-VPC, this
        is only available if you've enabled DNS hostnames for your VPC.
        """
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        __self__.public_ip = public_ip
        """
        The public IP address assigned to the Instance, if applicable. **NOTE**: If you are using an [`ec2.Eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
        """
        if root_block_devices and not isinstance(root_block_devices, list):
            raise TypeError("Expected argument 'root_block_devices' to be a list")
        __self__.root_block_devices = root_block_devices
        """
        The root block device mappings of the Instance
        """
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        __self__.security_groups = security_groups
        """
        The associated security groups.
        """
        if source_dest_check and not isinstance(source_dest_check, bool):
            raise TypeError("Expected argument 'source_dest_check' to be a bool")
        __self__.source_dest_check = source_dest_check
        """
        Whether the network interface performs source/destination checking (Boolean).
        """
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        __self__.subnet_id = subnet_id
        """
        The VPC subnet ID.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the Instance.
        """
        if tenancy and not isinstance(tenancy, str):
            raise TypeError("Expected argument 'tenancy' to be a str")
        __self__.tenancy = tenancy
        """
        The tenancy of the instance: `dedicated`, `default`, `host`.
        """
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        __self__.user_data = user_data
        """
        SHA-1 hash of User Data supplied to the Instance.
        """
        if user_data_base64 and not isinstance(user_data_base64, str):
            raise TypeError("Expected argument 'user_data_base64' to be a str")
        __self__.user_data_base64 = user_data_base64
        """
        Base64 encoded contents of User Data supplied to the Instance. Valid UTF-8 contents can be decoded with the [`base64decode` function](https://www.terraform.io/docs/configuration/functions/base64decode.html). This attribute is only exported if `get_user_data` is true.
        """
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError("Expected argument 'vpc_security_group_ids' to be a list")
        __self__.vpc_security_group_ids = vpc_security_group_ids
        """
        The associated security groups in a non-default VPC.
        """
class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            ami=self.ami,
            arn=self.arn,
            associate_public_ip_address=self.associate_public_ip_address,
            availability_zone=self.availability_zone,
            credit_specifications=self.credit_specifications,
            disable_api_termination=self.disable_api_termination,
            ebs_block_devices=self.ebs_block_devices,
            ebs_optimized=self.ebs_optimized,
            ephemeral_block_devices=self.ephemeral_block_devices,
            filters=self.filters,
            get_password_data=self.get_password_data,
            get_user_data=self.get_user_data,
            host_id=self.host_id,
            iam_instance_profile=self.iam_instance_profile,
            id=self.id,
            instance_id=self.instance_id,
            instance_state=self.instance_state,
            instance_tags=self.instance_tags,
            instance_type=self.instance_type,
            key_name=self.key_name,
            monitoring=self.monitoring,
            network_interface_id=self.network_interface_id,
            password_data=self.password_data,
            placement_group=self.placement_group,
            private_dns=self.private_dns,
            private_ip=self.private_ip,
            public_dns=self.public_dns,
            public_ip=self.public_ip,
            root_block_devices=self.root_block_devices,
            security_groups=self.security_groups,
            source_dest_check=self.source_dest_check,
            subnet_id=self.subnet_id,
            tags=self.tags,
            tenancy=self.tenancy,
            user_data=self.user_data,
            user_data_base64=self.user_data_base64,
            vpc_security_group_ids=self.vpc_security_group_ids)

def get_instance(filters=None,get_password_data=None,get_user_data=None,instance_id=None,instance_tags=None,tags=None,opts=None):
    """
    Use this data source to get the ID of an Amazon EC2 Instance for use in other
    resources.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/instance.html.markdown.


    :param list filters: One or more name/value pairs to use as filters. There are
           several valid keys, for a full reference, check out
           [describe-instances in the AWS CLI reference][1].
    :param bool get_password_data: If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
    :param bool get_user_data: Retrieve Base64 encoded User Data contents into the `user_data_base64` attribute. A SHA-1 hash of the User Data contents will always be present in the `user_data` attribute. Defaults to `false`.
    :param str instance_id: Specify the exact Instance ID with which to populate the data source.
    :param dict instance_tags: A mapping of tags, each pair of which must
           exactly match a pair on the desired Instance.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['getPasswordData'] = get_password_data
    __args__['getUserData'] = get_user_data
    __args__['instanceId'] = instance_id
    __args__['instanceTags'] = instance_tags
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getInstance:getInstance', __args__, opts=opts).value

    return AwaitableGetInstanceResult(
        ami=__ret__.get('ami'),
        arn=__ret__.get('arn'),
        associate_public_ip_address=__ret__.get('associatePublicIpAddress'),
        availability_zone=__ret__.get('availabilityZone'),
        credit_specifications=__ret__.get('creditSpecifications'),
        disable_api_termination=__ret__.get('disableApiTermination'),
        ebs_block_devices=__ret__.get('ebsBlockDevices'),
        ebs_optimized=__ret__.get('ebsOptimized'),
        ephemeral_block_devices=__ret__.get('ephemeralBlockDevices'),
        filters=__ret__.get('filters'),
        get_password_data=__ret__.get('getPasswordData'),
        get_user_data=__ret__.get('getUserData'),
        host_id=__ret__.get('hostId'),
        iam_instance_profile=__ret__.get('iamInstanceProfile'),
        id=__ret__.get('id'),
        instance_id=__ret__.get('instanceId'),
        instance_state=__ret__.get('instanceState'),
        instance_tags=__ret__.get('instanceTags'),
        instance_type=__ret__.get('instanceType'),
        key_name=__ret__.get('keyName'),
        monitoring=__ret__.get('monitoring'),
        network_interface_id=__ret__.get('networkInterfaceId'),
        password_data=__ret__.get('passwordData'),
        placement_group=__ret__.get('placementGroup'),
        private_dns=__ret__.get('privateDns'),
        private_ip=__ret__.get('privateIp'),
        public_dns=__ret__.get('publicDns'),
        public_ip=__ret__.get('publicIp'),
        root_block_devices=__ret__.get('rootBlockDevices'),
        security_groups=__ret__.get('securityGroups'),
        source_dest_check=__ret__.get('sourceDestCheck'),
        subnet_id=__ret__.get('subnetId'),
        tags=__ret__.get('tags'),
        tenancy=__ret__.get('tenancy'),
        user_data=__ret__.get('userData'),
        user_data_base64=__ret__.get('userDataBase64'),
        vpc_security_group_ids=__ret__.get('vpcSecurityGroupIds'))
