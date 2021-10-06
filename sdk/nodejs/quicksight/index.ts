// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./dataSource";
export * from "./group";
export * from "./groupMembership";
export * from "./user";

// Import resources to register:
import { DataSource } from "./dataSource";
import { Group } from "./group";
import { GroupMembership } from "./groupMembership";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:quicksight/dataSource:DataSource":
                return new DataSource(name, <any>undefined, { urn })
            case "aws:quicksight/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "aws:quicksight/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "aws:quicksight/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "quicksight/dataSource", _module)
pulumi.runtime.registerResourceModule("aws", "quicksight/group", _module)
pulumi.runtime.registerResourceModule("aws", "quicksight/groupMembership", _module)
pulumi.runtime.registerResourceModule("aws", "quicksight/user", _module)
