// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

export class Permissions extends pulumi.CustomResource {
    /**
     * Get an existing Permissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PermissionsState, opts?: pulumi.CustomResourceOptions): Permissions {
        return new Permissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lakeformation/permissions:Permissions';

    /**
     * Returns true if the given object is an instance of Permissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Permissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Permissions.__pulumiType;
    }

    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     */
    public readonly catalogId!: pulumi.Output<string | undefined>;
    /**
     * Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     */
    public readonly catalogResource!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration block for a data location resource. Detailed below.
     */
    public readonly dataLocation!: pulumi.Output<outputs.lakeformation.PermissionsDataLocation>;
    /**
     * Configuration block for a database resource. Detailed below.
     */
    public readonly database!: pulumi.Output<outputs.lakeformation.PermissionsDatabase>;
    /**
     * List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * Subset of `permissions` which the principal can pass.
     */
    public readonly permissionsWithGrantOptions!: pulumi.Output<string[]>;
    /**
     * Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    public readonly principal!: pulumi.Output<string>;
    /**
     * Configuration block for a table resource. Detailed below.
     */
    public readonly table!: pulumi.Output<outputs.lakeformation.PermissionsTable>;
    /**
     * Configuration block for a table with columns resource. Detailed below.
     */
    public readonly tableWithColumns!: pulumi.Output<outputs.lakeformation.PermissionsTableWithColumns>;

    /**
     * Create a Permissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PermissionsArgs | PermissionsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PermissionsState | undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["catalogResource"] = state ? state.catalogResource : undefined;
            inputs["dataLocation"] = state ? state.dataLocation : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["permissionsWithGrantOptions"] = state ? state.permissionsWithGrantOptions : undefined;
            inputs["principal"] = state ? state.principal : undefined;
            inputs["table"] = state ? state.table : undefined;
            inputs["tableWithColumns"] = state ? state.tableWithColumns : undefined;
        } else {
            const args = argsOrState as PermissionsArgs | undefined;
            if ((!args || args.permissions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissions'");
            }
            if ((!args || args.principal === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principal'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["catalogResource"] = args ? args.catalogResource : undefined;
            inputs["dataLocation"] = args ? args.dataLocation : undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["permissionsWithGrantOptions"] = args ? args.permissionsWithGrantOptions : undefined;
            inputs["principal"] = args ? args.principal : undefined;
            inputs["table"] = args ? args.table : undefined;
            inputs["tableWithColumns"] = args ? args.tableWithColumns : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Permissions.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Permissions resources.
 */
export interface PermissionsState {
    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     */
    catalogResource?: pulumi.Input<boolean>;
    /**
     * Configuration block for a data location resource. Detailed below.
     */
    dataLocation?: pulumi.Input<inputs.lakeformation.PermissionsDataLocation>;
    /**
     * Configuration block for a database resource. Detailed below.
     */
    database?: pulumi.Input<inputs.lakeformation.PermissionsDatabase>;
    /**
     * List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subset of `permissions` which the principal can pass.
     */
    permissionsWithGrantOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    principal?: pulumi.Input<string>;
    /**
     * Configuration block for a table resource. Detailed below.
     */
    table?: pulumi.Input<inputs.lakeformation.PermissionsTable>;
    /**
     * Configuration block for a table with columns resource. Detailed below.
     */
    tableWithColumns?: pulumi.Input<inputs.lakeformation.PermissionsTableWithColumns>;
}

/**
 * The set of arguments for constructing a Permissions resource.
 */
export interface PermissionsArgs {
    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
     */
    catalogResource?: pulumi.Input<boolean>;
    /**
     * Configuration block for a data location resource. Detailed below.
     */
    dataLocation?: pulumi.Input<inputs.lakeformation.PermissionsDataLocation>;
    /**
     * Configuration block for a database resource. Detailed below.
     */
    database?: pulumi.Input<inputs.lakeformation.PermissionsDatabase>;
    /**
     * List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    permissions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subset of `permissions` which the principal can pass.
     */
    permissionsWithGrantOptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
     */
    principal: pulumi.Input<string>;
    /**
     * Configuration block for a table resource. Detailed below.
     */
    table?: pulumi.Input<inputs.lakeformation.PermissionsTable>;
    /**
     * Configuration block for a table with columns resource. Detailed below.
     */
    tableWithColumns?: pulumi.Input<inputs.lakeformation.PermissionsTableWithColumns>;
}
