# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['ParameterArgs', 'Parameter']

@pulumi.input_type
class ParameterArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[Union[str, 'ParameterType']],
                 value: pulumi.Input[str],
                 allowed_pattern: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overwrite: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Parameter resource.
        :param pulumi.Input[Union[str, 'ParameterType']] type: The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        :param pulumi.Input[str] value: The value of the parameter.
        :param pulumi.Input[str] allowed_pattern: A regular expression used to validate the parameter value.
        :param pulumi.Input[str] arn: The ARN of the parameter.
        :param pulumi.Input[str] data_type: The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
               ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] key_id: The KMS key id or arn for encrypting a SecureString.
        :param pulumi.Input[str] name: The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        :param pulumi.Input[bool] overwrite: Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] tier: The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)
        if allowed_pattern is not None:
            pulumi.set(__self__, "allowed_pattern", allowed_pattern)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overwrite is not None:
            pulumi.set(__self__, "overwrite", overwrite)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[Union[str, 'ParameterType']]:
        """
        The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[Union[str, 'ParameterType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="allowedPattern")
    def allowed_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        A regular expression used to validate the parameter value.
        """
        return pulumi.get(self, "allowed_pattern")

    @allowed_pattern.setter
    def allowed_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_pattern", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the parameter.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[pulumi.Input[str]]:
        """
        The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
        ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        """
        return pulumi.get(self, "data_type")

    @data_type.setter
    def data_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The KMS key id or arn for encrypting a SecureString.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overwrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        """
        return pulumi.get(self, "overwrite")

    @overwrite.setter
    def overwrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the object.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class _ParameterState:
    def __init__(__self__, *,
                 allowed_pattern: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overwrite: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[Union[str, 'ParameterType']]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Parameter resources.
        :param pulumi.Input[str] allowed_pattern: A regular expression used to validate the parameter value.
        :param pulumi.Input[str] arn: The ARN of the parameter.
        :param pulumi.Input[str] data_type: The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
               ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] key_id: The KMS key id or arn for encrypting a SecureString.
        :param pulumi.Input[str] name: The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        :param pulumi.Input[bool] overwrite: Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] tier: The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        :param pulumi.Input[Union[str, 'ParameterType']] type: The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        :param pulumi.Input[str] value: The value of the parameter.
        :param pulumi.Input[int] version: The version of the parameter.
        """
        if allowed_pattern is not None:
            pulumi.set(__self__, "allowed_pattern", allowed_pattern)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overwrite is not None:
            pulumi.set(__self__, "overwrite", overwrite)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="allowedPattern")
    def allowed_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        A regular expression used to validate the parameter value.
        """
        return pulumi.get(self, "allowed_pattern")

    @allowed_pattern.setter
    def allowed_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_pattern", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the parameter.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[pulumi.Input[str]]:
        """
        The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
        ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        """
        return pulumi.get(self, "data_type")

    @data_type.setter
    def data_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the parameter.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The KMS key id or arn for encrypting a SecureString.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overwrite(self) -> Optional[pulumi.Input[bool]]:
        """
        Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        """
        return pulumi.get(self, "overwrite")

    @overwrite.setter
    def overwrite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the object.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[Union[str, 'ParameterType']]]:
        """
        The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[Union[str, 'ParameterType']]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        The version of the parameter.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class Parameter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_pattern: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overwrite: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[Union[str, 'ParameterType']]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SSM Parameter resource.

        ## Example Usage

        To store a basic string parameter:

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ssm.Parameter("foo",
            type="String",
            value="bar")
        ```

        To store an encrypted string using the default SSM KMS key:

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.Instance("default",
            allocated_storage=10,
            storage_type="gp2",
            engine="mysql",
            engine_version="5.7.16",
            instance_class="db.t2.micro",
            name="mydb",
            username="foo",
            password=var["database_master_password"],
            db_subnet_group_name="my_database_subnet_group",
            parameter_group_name="default.mysql5.7")
        secret = aws.ssm.Parameter("secret",
            description="The parameter description",
            type="SecureString",
            value=var["database_master_password"],
            tags={
                "environment": "production",
            })
        ```

        ## Import

        SSM Parameters can be imported using the `parameter store name`, e.g.

        ```sh
         $ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_pattern: A regular expression used to validate the parameter value.
        :param pulumi.Input[str] arn: The ARN of the parameter.
        :param pulumi.Input[str] data_type: The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
               ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] key_id: The KMS key id or arn for encrypting a SecureString.
        :param pulumi.Input[str] name: The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        :param pulumi.Input[bool] overwrite: Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] tier: The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        :param pulumi.Input[Union[str, 'ParameterType']] type: The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        :param pulumi.Input[str] value: The value of the parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ParameterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an SSM Parameter resource.

        ## Example Usage

        To store a basic string parameter:

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ssm.Parameter("foo",
            type="String",
            value="bar")
        ```

        To store an encrypted string using the default SSM KMS key:

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.Instance("default",
            allocated_storage=10,
            storage_type="gp2",
            engine="mysql",
            engine_version="5.7.16",
            instance_class="db.t2.micro",
            name="mydb",
            username="foo",
            password=var["database_master_password"],
            db_subnet_group_name="my_database_subnet_group",
            parameter_group_name="default.mysql5.7")
        secret = aws.ssm.Parameter("secret",
            description="The parameter description",
            type="SecureString",
            value=var["database_master_password"],
            tags={
                "environment": "production",
            })
        ```

        ## Import

        SSM Parameters can be imported using the `parameter store name`, e.g.

        ```sh
         $ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
        ```

        :param str resource_name: The name of the resource.
        :param ParameterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParameterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_pattern: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 data_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overwrite: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[Union[str, 'ParameterType']]] = None,
                 value: Optional[pulumi.Input[str]] = None,
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
            __props__ = ParameterArgs.__new__(ParameterArgs)

            __props__.__dict__["allowed_pattern"] = allowed_pattern
            __props__.__dict__["arn"] = arn
            __props__.__dict__["data_type"] = data_type
            __props__.__dict__["description"] = description
            __props__.__dict__["key_id"] = key_id
            __props__.__dict__["name"] = name
            __props__.__dict__["overwrite"] = overwrite
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tier"] = tier
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["version"] = None
        super(Parameter, __self__).__init__(
            'aws:ssm/parameter:Parameter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_pattern: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            data_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            overwrite: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tier: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[Union[str, 'ParameterType']]] = None,
            value: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'Parameter':
        """
        Get an existing Parameter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_pattern: A regular expression used to validate the parameter value.
        :param pulumi.Input[str] arn: The ARN of the parameter.
        :param pulumi.Input[str] data_type: The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
               ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        :param pulumi.Input[str] description: The description of the parameter.
        :param pulumi.Input[str] key_id: The KMS key id or arn for encrypting a SecureString.
        :param pulumi.Input[str] name: The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        :param pulumi.Input[bool] overwrite: Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the object.
        :param pulumi.Input[str] tier: The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        :param pulumi.Input[Union[str, 'ParameterType']] type: The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        :param pulumi.Input[str] value: The value of the parameter.
        :param pulumi.Input[int] version: The version of the parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ParameterState.__new__(_ParameterState)

        __props__.__dict__["allowed_pattern"] = allowed_pattern
        __props__.__dict__["arn"] = arn
        __props__.__dict__["data_type"] = data_type
        __props__.__dict__["description"] = description
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["name"] = name
        __props__.__dict__["overwrite"] = overwrite
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tier"] = tier
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        __props__.__dict__["version"] = version
        return Parameter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedPattern")
    def allowed_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        A regular expression used to validate the parameter value.
        """
        return pulumi.get(self, "allowed_pattern")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the parameter.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> pulumi.Output[str]:
        """
        The data_type of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
        ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the parameter.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The KMS key id or arn for encrypting a SecureString.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the parameter. If the name contains a path (e.g. any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def overwrite(self) -> pulumi.Output[Optional[bool]]:
        """
        Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
        """
        return pulumi.get(self, "overwrite")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the object.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[Optional[str]]:
        """
        The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the parameter.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        The version of the parameter.
        """
        return pulumi.get(self, "version")

