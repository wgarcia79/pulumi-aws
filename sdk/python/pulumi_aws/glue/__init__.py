# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .catalog_database import *
from .catalog_table import *
from .classifier import *
from .connection import *
from .crawler import *
from .data_catalog_encryption_settings import *
from .dev_endpoint import *
from .get_connection import *
from .get_data_catalog_encryption_settings import *
from .get_script import *
from .job import *
from .ml_transform import *
from .partition import *
from .registry import *
from .resource_policy import *
from .schema import *
from .security_configuration import *
from .trigger import *
from .user_defined_function import *
from .workflow import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:glue/catalogDatabase:CatalogDatabase":
                return CatalogDatabase(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/catalogTable:CatalogTable":
                return CatalogTable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/classifier:Classifier":
                return Classifier(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/connection:Connection":
                return Connection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/crawler:Crawler":
                return Crawler(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings":
                return DataCatalogEncryptionSettings(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/devEndpoint:DevEndpoint":
                return DevEndpoint(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/job:Job":
                return Job(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/mLTransform:MLTransform":
                return MLTransform(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/partition:Partition":
                return Partition(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/registry:Registry":
                return Registry(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/resourcePolicy:ResourcePolicy":
                return ResourcePolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/schema:Schema":
                return Schema(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/securityConfiguration:SecurityConfiguration":
                return SecurityConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/trigger:Trigger":
                return Trigger(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/userDefinedFunction:UserDefinedFunction":
                return UserDefinedFunction(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:glue/workflow:Workflow":
                return Workflow(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "glue/catalogDatabase", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/catalogTable", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/classifier", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/connection", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/crawler", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/dataCatalogEncryptionSettings", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/devEndpoint", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/job", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/mLTransform", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/partition", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/registry", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/resourcePolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/schema", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/securityConfiguration", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/trigger", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/userDefinedFunction", _module_instance)
    pulumi.runtime.register_resource_module("aws", "glue/workflow", _module_instance)

_register_module()
