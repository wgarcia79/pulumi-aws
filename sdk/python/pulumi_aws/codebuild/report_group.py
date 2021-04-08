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

__all__ = ['ReportGroupArgs', 'ReportGroup']

@pulumi.input_type
class ReportGroupArgs:
    def __init__(__self__, *,
                 export_config: pulumi.Input['ReportGroupExportConfigArgs'],
                 type: pulumi.Input[str],
                 delete_reports: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ReportGroup resource.
        :param pulumi.Input['ReportGroupExportConfigArgs'] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        :param pulumi.Input[bool] delete_reports: If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        """
        pulumi.set(__self__, "export_config", export_config)
        pulumi.set(__self__, "type", type)
        if delete_reports is not None:
            pulumi.set(__self__, "delete_reports", delete_reports)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="exportConfig")
    def export_config(self) -> pulumi.Input['ReportGroupExportConfigArgs']:
        """
        Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        """
        return pulumi.get(self, "export_config")

    @export_config.setter
    def export_config(self, value: pulumi.Input['ReportGroupExportConfigArgs']):
        pulumi.set(self, "export_config", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="deleteReports")
    def delete_reports(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        """
        return pulumi.get(self, "delete_reports")

    @delete_reports.setter
    def delete_reports(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_reports", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a Report Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ReportGroupState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 delete_reports: Optional[pulumi.Input[bool]] = None,
                 export_config: Optional[pulumi.Input['ReportGroupExportConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReportGroup resources.
        :param pulumi.Input[str] arn: The ARN of Report Group.
        :param pulumi.Input[str] created: The date and time this Report Group was created.
        :param pulumi.Input[bool] delete_reports: If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        :param pulumi.Input['ReportGroupExportConfigArgs'] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if delete_reports is not None:
            pulumi.set(__self__, "delete_reports", delete_reports)
        if export_config is not None:
            pulumi.set(__self__, "export_config", export_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of Report Group.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time this Report Group was created.
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="deleteReports")
    def delete_reports(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        """
        return pulumi.get(self, "delete_reports")

    @delete_reports.setter
    def delete_reports(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_reports", value)

    @property
    @pulumi.getter(name="exportConfig")
    def export_config(self) -> Optional[pulumi.Input['ReportGroupExportConfigArgs']]:
        """
        Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        """
        return pulumi.get(self, "export_config")

    @export_config.setter
    def export_config(self, value: Optional[pulumi.Input['ReportGroupExportConfigArgs']]):
        pulumi.set(self, "export_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a Report Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ReportGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_reports: Optional[pulumi.Input[bool]] = None,
                 export_config: Optional[pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CodeBuild Report Groups Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey",
            description="my test kms key",
            deletion_window_in_days=7,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Id": "kms-tf-1",
          "Statement": [
            {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                "AWS": "*"
              },
              "Action": "kms:*",
              "Resource": "*"
            }
          ]
        }
        \"\"\")
        example_bucket = aws.s3.Bucket("exampleBucket")
        example_report_group = aws.codebuild.ReportGroup("exampleReportGroup",
            type="TEST",
            export_config=aws.codebuild.ReportGroupExportConfigArgs(
                type="S3",
                s3_destination={
                    "bucket": example_bucket.id,
                    "encryptionDisabled": False,
                    "encryption_key": example_key.arn,
                    "packaging": "NONE",
                    "path": "/some",
                },
            ))
        ```

        ## Import

        CodeBuild Report Group can be imported using the CodeBuild Report Group arn, e.g.

        ```sh
         $ pulumi import aws:codebuild/reportGroup:ReportGroup example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_reports: If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        :param pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReportGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CodeBuild Report Groups Resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey",
            description="my test kms key",
            deletion_window_in_days=7,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Id": "kms-tf-1",
          "Statement": [
            {
              "Sid": "Enable IAM User Permissions",
              "Effect": "Allow",
              "Principal": {
                "AWS": "*"
              },
              "Action": "kms:*",
              "Resource": "*"
            }
          ]
        }
        \"\"\")
        example_bucket = aws.s3.Bucket("exampleBucket")
        example_report_group = aws.codebuild.ReportGroup("exampleReportGroup",
            type="TEST",
            export_config=aws.codebuild.ReportGroupExportConfigArgs(
                type="S3",
                s3_destination={
                    "bucket": example_bucket.id,
                    "encryptionDisabled": False,
                    "encryption_key": example_key.arn,
                    "packaging": "NONE",
                    "path": "/some",
                },
            ))
        ```

        ## Import

        CodeBuild Report Group can be imported using the CodeBuild Report Group arn, e.g.

        ```sh
         $ pulumi import aws:codebuild/reportGroup:ReportGroup example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
        ```

        :param str resource_name: The name of the resource.
        :param ReportGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReportGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_reports: Optional[pulumi.Input[bool]] = None,
                 export_config: Optional[pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReportGroupArgs.__new__(ReportGroupArgs)

            __props__.__dict__["delete_reports"] = delete_reports
            if export_config is None and not opts.urn:
                raise TypeError("Missing required property 'export_config'")
            __props__.__dict__["export_config"] = export_config
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
            __props__.__dict__["created"] = None
        super(ReportGroup, __self__).__init__(
            'aws:codebuild/reportGroup:ReportGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            created: Optional[pulumi.Input[str]] = None,
            delete_reports: Optional[pulumi.Input[bool]] = None,
            export_config: Optional[pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ReportGroup':
        """
        Get an existing ReportGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of Report Group.
        :param pulumi.Input[str] created: The date and time this Report Group was created.
        :param pulumi.Input[bool] delete_reports: If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        :param pulumi.Input[pulumi.InputType['ReportGroupExportConfigArgs']] export_config: Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        :param pulumi.Input[str] name: The name of a Report Group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] type: The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReportGroupState.__new__(_ReportGroupState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["created"] = created
        __props__.__dict__["delete_reports"] = delete_reports
        __props__.__dict__["export_config"] = export_config
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        return ReportGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of Report Group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        The date and time this Report Group was created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="deleteReports")
    def delete_reports(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, deletes any reports that belong to a report group before deleting the report group. If `false`, you must delete any reports in the report group before deleting it. Default value is `false`.
        """
        return pulumi.get(self, "delete_reports")

    @property
    @pulumi.getter(name="exportConfig")
    def export_config(self) -> pulumi.Output['outputs.ReportGroupExportConfig']:
        """
        Information about the destination where the raw data of this Report Group is exported. see Export Config documented below.
        """
        return pulumi.get(self, "export_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of a Report Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The export configuration type. Valid values are `S3` and `NO_EXPORT`.
        """
        return pulumi.get(self, "type")

