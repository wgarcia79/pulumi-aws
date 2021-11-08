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

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 image_recipe_arn: pulumi.Input[str],
                 infrastructure_configuration_arn: pulumi.Input[str],
                 distribution_configuration_arn: Optional[pulumi.Input[str]] = None,
                 enhanced_image_metadata_enabled: Optional[pulumi.Input[bool]] = None,
                 image_tests_configuration: Optional[pulumi.Input['ImageImageTestsConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[str] image_recipe_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        :param pulumi.Input[str] infrastructure_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        :param pulumi.Input[str] distribution_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        :param pulumi.Input[bool] enhanced_image_metadata_enabled: Whether additional information about the image being created is collected. Defaults to `true`.
        :param pulumi.Input['ImageImageTestsConfigurationArgs'] image_tests_configuration: Configuration block with image tests configuration. Detailed below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "image_recipe_arn", image_recipe_arn)
        pulumi.set(__self__, "infrastructure_configuration_arn", infrastructure_configuration_arn)
        if distribution_configuration_arn is not None:
            pulumi.set(__self__, "distribution_configuration_arn", distribution_configuration_arn)
        if enhanced_image_metadata_enabled is not None:
            pulumi.set(__self__, "enhanced_image_metadata_enabled", enhanced_image_metadata_enabled)
        if image_tests_configuration is not None:
            pulumi.set(__self__, "image_tests_configuration", image_tests_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="imageRecipeArn")
    def image_recipe_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        """
        return pulumi.get(self, "image_recipe_arn")

    @image_recipe_arn.setter
    def image_recipe_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "image_recipe_arn", value)

    @property
    @pulumi.getter(name="infrastructureConfigurationArn")
    def infrastructure_configuration_arn(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        """
        return pulumi.get(self, "infrastructure_configuration_arn")

    @infrastructure_configuration_arn.setter
    def infrastructure_configuration_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "infrastructure_configuration_arn", value)

    @property
    @pulumi.getter(name="distributionConfigurationArn")
    def distribution_configuration_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        """
        return pulumi.get(self, "distribution_configuration_arn")

    @distribution_configuration_arn.setter
    def distribution_configuration_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distribution_configuration_arn", value)

    @property
    @pulumi.getter(name="enhancedImageMetadataEnabled")
    def enhanced_image_metadata_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether additional information about the image being created is collected. Defaults to `true`.
        """
        return pulumi.get(self, "enhanced_image_metadata_enabled")

    @enhanced_image_metadata_enabled.setter
    def enhanced_image_metadata_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enhanced_image_metadata_enabled", value)

    @property
    @pulumi.getter(name="imageTestsConfiguration")
    def image_tests_configuration(self) -> Optional[pulumi.Input['ImageImageTestsConfigurationArgs']]:
        """
        Configuration block with image tests configuration. Detailed below.
        """
        return pulumi.get(self, "image_tests_configuration")

    @image_tests_configuration.setter
    def image_tests_configuration(self, value: Optional[pulumi.Input['ImageImageTestsConfigurationArgs']]):
        pulumi.set(self, "image_tests_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ImageState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 date_created: Optional[pulumi.Input[str]] = None,
                 distribution_configuration_arn: Optional[pulumi.Input[str]] = None,
                 enhanced_image_metadata_enabled: Optional[pulumi.Input[bool]] = None,
                 image_recipe_arn: Optional[pulumi.Input[str]] = None,
                 image_tests_configuration: Optional[pulumi.Input['ImageImageTestsConfigurationArgs']] = None,
                 infrastructure_configuration_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_version: Optional[pulumi.Input[str]] = None,
                 output_resources: Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceArgs']]]] = None,
                 platform: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Image resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the image.
        :param pulumi.Input[str] date_created: Date the image was created.
        :param pulumi.Input[str] distribution_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        :param pulumi.Input[bool] enhanced_image_metadata_enabled: Whether additional information about the image being created is collected. Defaults to `true`.
        :param pulumi.Input[str] image_recipe_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        :param pulumi.Input['ImageImageTestsConfigurationArgs'] image_tests_configuration: Configuration block with image tests configuration. Detailed below.
        :param pulumi.Input[str] infrastructure_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        :param pulumi.Input[str] name: Name of the AMI.
        :param pulumi.Input[str] os_version: Operating System version of the image.
        :param pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceArgs']]] output_resources: List of objects with resources created by the image.
        :param pulumi.Input[str] platform: Platform of the image.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] version: Version of the image.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if date_created is not None:
            pulumi.set(__self__, "date_created", date_created)
        if distribution_configuration_arn is not None:
            pulumi.set(__self__, "distribution_configuration_arn", distribution_configuration_arn)
        if enhanced_image_metadata_enabled is not None:
            pulumi.set(__self__, "enhanced_image_metadata_enabled", enhanced_image_metadata_enabled)
        if image_recipe_arn is not None:
            pulumi.set(__self__, "image_recipe_arn", image_recipe_arn)
        if image_tests_configuration is not None:
            pulumi.set(__self__, "image_tests_configuration", image_tests_configuration)
        if infrastructure_configuration_arn is not None:
            pulumi.set(__self__, "infrastructure_configuration_arn", infrastructure_configuration_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_version is not None:
            pulumi.set(__self__, "os_version", os_version)
        if output_resources is not None:
            pulumi.set(__self__, "output_resources", output_resources)
        if platform is not None:
            pulumi.set(__self__, "platform", platform)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the image.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> Optional[pulumi.Input[str]]:
        """
        Date the image was created.
        """
        return pulumi.get(self, "date_created")

    @date_created.setter
    def date_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_created", value)

    @property
    @pulumi.getter(name="distributionConfigurationArn")
    def distribution_configuration_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        """
        return pulumi.get(self, "distribution_configuration_arn")

    @distribution_configuration_arn.setter
    def distribution_configuration_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distribution_configuration_arn", value)

    @property
    @pulumi.getter(name="enhancedImageMetadataEnabled")
    def enhanced_image_metadata_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether additional information about the image being created is collected. Defaults to `true`.
        """
        return pulumi.get(self, "enhanced_image_metadata_enabled")

    @enhanced_image_metadata_enabled.setter
    def enhanced_image_metadata_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enhanced_image_metadata_enabled", value)

    @property
    @pulumi.getter(name="imageRecipeArn")
    def image_recipe_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        """
        return pulumi.get(self, "image_recipe_arn")

    @image_recipe_arn.setter
    def image_recipe_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_recipe_arn", value)

    @property
    @pulumi.getter(name="imageTestsConfiguration")
    def image_tests_configuration(self) -> Optional[pulumi.Input['ImageImageTestsConfigurationArgs']]:
        """
        Configuration block with image tests configuration. Detailed below.
        """
        return pulumi.get(self, "image_tests_configuration")

    @image_tests_configuration.setter
    def image_tests_configuration(self, value: Optional[pulumi.Input['ImageImageTestsConfigurationArgs']]):
        pulumi.set(self, "image_tests_configuration", value)

    @property
    @pulumi.getter(name="infrastructureConfigurationArn")
    def infrastructure_configuration_arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        """
        return pulumi.get(self, "infrastructure_configuration_arn")

    @infrastructure_configuration_arn.setter
    def infrastructure_configuration_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "infrastructure_configuration_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the AMI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> Optional[pulumi.Input[str]]:
        """
        Operating System version of the image.
        """
        return pulumi.get(self, "os_version")

    @os_version.setter
    def os_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_version", value)

    @property
    @pulumi.getter(name="outputResources")
    def output_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceArgs']]]]:
        """
        List of objects with resources created by the image.
        """
        return pulumi.get(self, "output_resources")

    @output_resources.setter
    def output_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ImageOutputResourceArgs']]]]):
        pulumi.set(self, "output_resources", value)

    @property
    @pulumi.getter
    def platform(self) -> Optional[pulumi.Input[str]]:
        """
        Platform of the image.
        """
        return pulumi.get(self, "platform")

    @platform.setter
    def platform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the image.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distribution_configuration_arn: Optional[pulumi.Input[str]] = None,
                 enhanced_image_metadata_enabled: Optional[pulumi.Input[bool]] = None,
                 image_recipe_arn: Optional[pulumi.Input[str]] = None,
                 image_tests_configuration: Optional[pulumi.Input[pulumi.InputType['ImageImageTestsConfigurationArgs']]] = None,
                 infrastructure_configuration_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages an Image Builder Image.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.Image("example",
            distribution_configuration_arn=aws_imagebuilder_distribution_configuration["example"]["arn"],
            image_recipe_arn=aws_imagebuilder_image_recipe["example"]["arn"],
            infrastructure_configuration_arn=aws_imagebuilder_infrastructure_configuration["example"]["arn"])
        ```

        ## Import

        `aws_imagebuilder_image` resources can be imported using the Amazon Resource Name (ARN), e.g.,

        ```sh
         $ pulumi import aws:imagebuilder/image:Image example arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distribution_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        :param pulumi.Input[bool] enhanced_image_metadata_enabled: Whether additional information about the image being created is collected. Defaults to `true`.
        :param pulumi.Input[str] image_recipe_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        :param pulumi.Input[pulumi.InputType['ImageImageTestsConfigurationArgs']] image_tests_configuration: Configuration block with image tests configuration. Detailed below.
        :param pulumi.Input[str] infrastructure_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Image Builder Image.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.imagebuilder.Image("example",
            distribution_configuration_arn=aws_imagebuilder_distribution_configuration["example"]["arn"],
            image_recipe_arn=aws_imagebuilder_image_recipe["example"]["arn"],
            infrastructure_configuration_arn=aws_imagebuilder_infrastructure_configuration["example"]["arn"])
        ```

        ## Import

        `aws_imagebuilder_image` resources can be imported using the Amazon Resource Name (ARN), e.g.,

        ```sh
         $ pulumi import aws:imagebuilder/image:Image example arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
        ```

        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distribution_configuration_arn: Optional[pulumi.Input[str]] = None,
                 enhanced_image_metadata_enabled: Optional[pulumi.Input[bool]] = None,
                 image_recipe_arn: Optional[pulumi.Input[str]] = None,
                 image_tests_configuration: Optional[pulumi.Input[pulumi.InputType['ImageImageTestsConfigurationArgs']]] = None,
                 infrastructure_configuration_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["distribution_configuration_arn"] = distribution_configuration_arn
            __props__.__dict__["enhanced_image_metadata_enabled"] = enhanced_image_metadata_enabled
            if image_recipe_arn is None and not opts.urn:
                raise TypeError("Missing required property 'image_recipe_arn'")
            __props__.__dict__["image_recipe_arn"] = image_recipe_arn
            __props__.__dict__["image_tests_configuration"] = image_tests_configuration
            if infrastructure_configuration_arn is None and not opts.urn:
                raise TypeError("Missing required property 'infrastructure_configuration_arn'")
            __props__.__dict__["infrastructure_configuration_arn"] = infrastructure_configuration_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["date_created"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["os_version"] = None
            __props__.__dict__["output_resources"] = None
            __props__.__dict__["platform"] = None
            __props__.__dict__["tags_all"] = None
            __props__.__dict__["version"] = None
        super(Image, __self__).__init__(
            'aws:imagebuilder/image:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            date_created: Optional[pulumi.Input[str]] = None,
            distribution_configuration_arn: Optional[pulumi.Input[str]] = None,
            enhanced_image_metadata_enabled: Optional[pulumi.Input[bool]] = None,
            image_recipe_arn: Optional[pulumi.Input[str]] = None,
            image_tests_configuration: Optional[pulumi.Input[pulumi.InputType['ImageImageTestsConfigurationArgs']]] = None,
            infrastructure_configuration_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            os_version: Optional[pulumi.Input[str]] = None,
            output_resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageOutputResourceArgs']]]]] = None,
            platform: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Image':
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the image.
        :param pulumi.Input[str] date_created: Date the image was created.
        :param pulumi.Input[str] distribution_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        :param pulumi.Input[bool] enhanced_image_metadata_enabled: Whether additional information about the image being created is collected. Defaults to `true`.
        :param pulumi.Input[str] image_recipe_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        :param pulumi.Input[pulumi.InputType['ImageImageTestsConfigurationArgs']] image_tests_configuration: Configuration block with image tests configuration. Detailed below.
        :param pulumi.Input[str] infrastructure_configuration_arn: Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        :param pulumi.Input[str] name: Name of the AMI.
        :param pulumi.Input[str] os_version: Operating System version of the image.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ImageOutputResourceArgs']]]] output_resources: List of objects with resources created by the image.
        :param pulumi.Input[str] platform: Platform of the image.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] version: Version of the image.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageState.__new__(_ImageState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["date_created"] = date_created
        __props__.__dict__["distribution_configuration_arn"] = distribution_configuration_arn
        __props__.__dict__["enhanced_image_metadata_enabled"] = enhanced_image_metadata_enabled
        __props__.__dict__["image_recipe_arn"] = image_recipe_arn
        __props__.__dict__["image_tests_configuration"] = image_tests_configuration
        __props__.__dict__["infrastructure_configuration_arn"] = infrastructure_configuration_arn
        __props__.__dict__["name"] = name
        __props__.__dict__["os_version"] = os_version
        __props__.__dict__["output_resources"] = output_resources
        __props__.__dict__["platform"] = platform
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["version"] = version
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the image.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dateCreated")
    def date_created(self) -> pulumi.Output[str]:
        """
        Date the image was created.
        """
        return pulumi.get(self, "date_created")

    @property
    @pulumi.getter(name="distributionConfigurationArn")
    def distribution_configuration_arn(self) -> pulumi.Output[Optional[str]]:
        """
        Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
        """
        return pulumi.get(self, "distribution_configuration_arn")

    @property
    @pulumi.getter(name="enhancedImageMetadataEnabled")
    def enhanced_image_metadata_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether additional information about the image being created is collected. Defaults to `true`.
        """
        return pulumi.get(self, "enhanced_image_metadata_enabled")

    @property
    @pulumi.getter(name="imageRecipeArn")
    def image_recipe_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
        """
        return pulumi.get(self, "image_recipe_arn")

    @property
    @pulumi.getter(name="imageTestsConfiguration")
    def image_tests_configuration(self) -> pulumi.Output['outputs.ImageImageTestsConfiguration']:
        """
        Configuration block with image tests configuration. Detailed below.
        """
        return pulumi.get(self, "image_tests_configuration")

    @property
    @pulumi.getter(name="infrastructureConfigurationArn")
    def infrastructure_configuration_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
        """
        return pulumi.get(self, "infrastructure_configuration_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the AMI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> pulumi.Output[str]:
        """
        Operating System version of the image.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="outputResources")
    def output_resources(self) -> pulumi.Output[Sequence['outputs.ImageOutputResource']]:
        """
        List of objects with resources created by the image.
        """
        return pulumi.get(self, "output_resources")

    @property
    @pulumi.getter
    def platform(self) -> pulumi.Output[str]:
        """
        Platform of the image.
        """
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags for the Image Builder Image. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Version of the image.
        """
        return pulumi.get(self, "version")

