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

__all__ = ['TableArgs', 'Table']

@pulumi.input_type
class TableArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 table_name: pulumi.Input[str],
                 retention_properties: Optional[pulumi.Input['TableRetentionPropertiesArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Table resource.
        :param pulumi.Input[str] database_name: The name of the Timestream database.
        :param pulumi.Input[str] table_name: The name of the Timestream table.
        :param pulumi.Input['TableRetentionPropertiesArgs'] retention_properties: The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "table_name", table_name)
        if retention_properties is not None:
            pulumi.set(__self__, "retention_properties", retention_properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The name of the Timestream database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        The name of the Timestream table.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="retentionProperties")
    def retention_properties(self) -> Optional[pulumi.Input['TableRetentionPropertiesArgs']]:
        """
        The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        """
        return pulumi.get(self, "retention_properties")

    @retention_properties.setter
    def retention_properties(self, value: Optional[pulumi.Input['TableRetentionPropertiesArgs']]):
        pulumi.set(self, "retention_properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TableState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 retention_properties: Optional[pulumi.Input['TableRetentionPropertiesArgs']] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Table resources.
        :param pulumi.Input[str] arn: The ARN that uniquely identifies this table.
        :param pulumi.Input[str] database_name: The name of the Timestream database.
        :param pulumi.Input['TableRetentionPropertiesArgs'] retention_properties: The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        :param pulumi.Input[str] table_name: The name of the Timestream table.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if retention_properties is not None:
            pulumi.set(__self__, "retention_properties", retention_properties)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN that uniquely identifies this table.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Timestream database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="retentionProperties")
    def retention_properties(self) -> Optional[pulumi.Input['TableRetentionPropertiesArgs']]:
        """
        The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        """
        return pulumi.get(self, "retention_properties")

    @retention_properties.setter
    def retention_properties(self, value: Optional[pulumi.Input['TableRetentionPropertiesArgs']]):
        pulumi.set(self, "retention_properties", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Timestream table.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class Table(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 retention_properties: Optional[pulumi.Input[pulumi.InputType['TableRetentionPropertiesArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a Timestream table resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.timestreamwrite.Table("example",
            database_name=aws_timestreamwrite_database["example"]["database_name"],
            table_name="example")
        ```
        ### Full usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.timestreamwrite.Table("example",
            database_name=aws_timestreamwrite_database["example"]["database_name"],
            table_name="example",
            retention_properties=aws.timestreamwrite.TableRetentionPropertiesArgs(
                magnetic_store_retention_period_in_days=30,
                memory_store_retention_period_in_hours=8,
            ),
            tags={
                "Name": "example-timestream-table",
            })
        ```

        ## Import

        Timestream tables can be imported using the `table_name` and `database_name` separate by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:timestreamwrite/table:Table example ExampleTable:ExampleDatabase
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the Timestream database.
        :param pulumi.Input[pulumi.InputType['TableRetentionPropertiesArgs']] retention_properties: The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        :param pulumi.Input[str] table_name: The name of the Timestream table.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Timestream table resource.

        ## Example Usage
        ### Basic usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.timestreamwrite.Table("example",
            database_name=aws_timestreamwrite_database["example"]["database_name"],
            table_name="example")
        ```
        ### Full usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.timestreamwrite.Table("example",
            database_name=aws_timestreamwrite_database["example"]["database_name"],
            table_name="example",
            retention_properties=aws.timestreamwrite.TableRetentionPropertiesArgs(
                magnetic_store_retention_period_in_days=30,
                memory_store_retention_period_in_hours=8,
            ),
            tags={
                "Name": "example-timestream-table",
            })
        ```

        ## Import

        Timestream tables can be imported using the `table_name` and `database_name` separate by a colon (`:`), e.g.,

        ```sh
         $ pulumi import aws:timestreamwrite/table:Table example ExampleTable:ExampleDatabase
        ```

        :param str resource_name: The name of the resource.
        :param TableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 retention_properties: Optional[pulumi.Input[pulumi.InputType['TableRetentionPropertiesArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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
            __props__ = TableArgs.__new__(TableArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["retention_properties"] = retention_properties
            if table_name is None and not opts.urn:
                raise TypeError("Missing required property 'table_name'")
            __props__.__dict__["table_name"] = table_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["tags_all"] = None
        super(Table, __self__).__init__(
            'aws:timestreamwrite/table:Table',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            retention_properties: Optional[pulumi.Input[pulumi.InputType['TableRetentionPropertiesArgs']]] = None,
            table_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Table':
        """
        Get an existing Table resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN that uniquely identifies this table.
        :param pulumi.Input[str] database_name: The name of the Timestream database.
        :param pulumi.Input[pulumi.InputType['TableRetentionPropertiesArgs']] retention_properties: The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        :param pulumi.Input[str] table_name: The name of the Timestream table.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TableState.__new__(_TableState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["retention_properties"] = retention_properties
        __props__.__dict__["table_name"] = table_name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return Table(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN that uniquely identifies this table.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the Timestream database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="retentionProperties")
    def retention_properties(self) -> pulumi.Output['outputs.TableRetentionProperties']:
        """
        The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
        """
        return pulumi.get(self, "retention_properties")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[str]:
        """
        The name of the Timestream table.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

