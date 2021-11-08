// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Cognito User Pool resource.
 *
 * ## Example Usage
 * ### Basic configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const pool = new aws.cognito.UserPool("pool", {});
 * ```
 * ### Enabling SMS and Software Token Multi-Factor Authentication
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configuration ...
 * const example = new aws.cognito.UserPool("example", {
 *     mfaConfiguration: "ON",
 *     smsAuthenticationMessage: "Your code is {####}",
 *     smsConfiguration: {
 *         externalId: "example",
 *         snsCallerArn: aws_iam_role.example.arn,
 *     },
 *     softwareTokenMfaConfiguration: {
 *         enabled: true,
 *     },
 * });
 * ```
 * ### Using Account Recovery Setting
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.cognito.UserPool("test", {
 *     accountRecoverySetting: {
 *         recoveryMechanisms: [
 *             {
 *                 name: "verified_email",
 *                 priority: 1,
 *             },
 *             {
 *                 name: "verified_phone_number",
 *                 priority: 2,
 *             },
 *         ],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Cognito User Pools can be imported using the `id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cognito/userPool:UserPool pool <id>
 * ```
 */
export class UserPool extends pulumi.CustomResource {
    /**
     * Get an existing UserPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPoolState, opts?: pulumi.CustomResourceOptions): UserPool {
        return new UserPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/userPool:UserPool';

    /**
     * Returns true if the given object is an instance of UserPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPool.__pulumiType;
    }

    /**
     * Configuration block to define which verified available method a user can use to recover their forgotten password. Detailed below.
     */
    public readonly accountRecoverySetting!: pulumi.Output<outputs.cognito.UserPoolAccountRecoverySetting | undefined>;
    /**
     * Configuration block for creating a new user profile. Detailed below.
     */
    public readonly adminCreateUserConfig!: pulumi.Output<outputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Valid values: `phoneNumber`, `email`, or `preferredUsername`. Conflicts with `usernameAttributes`.
     */
    public readonly aliasAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * ARN of the user pool.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Attributes to be auto-verified. Valid values: `email`, `phoneNumber`.
     */
    public readonly autoVerifiedAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * Date the user pool was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * A custom domain name that you provide to Amazon Cognito. This parameter applies only if you use a custom domain to host the sign-up and sign-in pages for your application. For example: `auth.example.com`.
     */
    public /*out*/ readonly customDomain!: pulumi.Output<string>;
    /**
     * Configuration block for the user pool's device tracking. Detailed below.
     */
    public readonly deviceConfiguration!: pulumi.Output<outputs.cognito.UserPoolDeviceConfiguration | undefined>;
    /**
     * Holds the domain prefix if the user pool has a domain associated with it.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * Configuration block for configuring email. Detailed below.
     */
    public readonly emailConfiguration!: pulumi.Output<outputs.cognito.UserPoolEmailConfiguration | undefined>;
    /**
     * String representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    public readonly emailVerificationMessage!: pulumi.Output<string>;
    /**
     * String representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    public readonly emailVerificationSubject!: pulumi.Output<string>;
    /**
     * Endpoint name of the user pool. Example format: `cognito-idp.REGION.amazonaws.com/xxxx_yyyyy`
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * A number estimating the size of the user pool.
     */
    public /*out*/ readonly estimatedNumberOfUsers!: pulumi.Output<number>;
    /**
     * Configuration block for the AWS Lambda triggers associated with the user pool. Detailed below.
     */
    public readonly lambdaConfig!: pulumi.Output<outputs.cognito.UserPoolLambdaConfig | undefined>;
    /**
     * Date the user pool was last modified.
     */
    public /*out*/ readonly lastModifiedDate!: pulumi.Output<string>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values are `OFF` (MFA Tokens are not required), `ON` (MFA is required for all users to sign in; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured), or `OPTIONAL` (MFA Will be required only for individual users who have MFA Enabled; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured).
     */
    public readonly mfaConfiguration!: pulumi.Output<string | undefined>;
    /**
     * Name of the attribute.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration blocked for information about the user pool password policy. Detailed below.
     */
    public readonly passwordPolicy!: pulumi.Output<outputs.cognito.UserPoolPasswordPolicy>;
    /**
     * Configuration block for the schema attributes of a user pool. Detailed below. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Attributes can be added, but not modified or removed. Maximum of 50 attributes.
     */
    public readonly schemas!: pulumi.Output<outputs.cognito.UserPoolSchema[] | undefined>;
    /**
     * String representing the SMS authentication message. The Message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    public readonly smsAuthenticationMessage!: pulumi.Output<string | undefined>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection.
     */
    public readonly smsConfiguration!: pulumi.Output<outputs.cognito.UserPoolSmsConfiguration>;
    /**
     * String representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    public readonly smsVerificationMessage!: pulumi.Output<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    public readonly softwareTokenMfaConfiguration!: pulumi.Output<outputs.cognito.UserPoolSoftwareTokenMfaConfiguration | undefined>;
    /**
     * Map of tags to assign to the User Pool. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features. Detailed below.
     */
    public readonly userPoolAddOns!: pulumi.Output<outputs.cognito.UserPoolUserPoolAddOns | undefined>;
    /**
     * Whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    public readonly usernameAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * Configuration block for username configuration. Detailed below.
     */
    public readonly usernameConfiguration!: pulumi.Output<outputs.cognito.UserPoolUsernameConfiguration | undefined>;
    /**
     * Configuration block for verification message templates. Detailed below.
     */
    public readonly verificationMessageTemplate!: pulumi.Output<outputs.cognito.UserPoolVerificationMessageTemplate>;

    /**
     * Create a UserPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPoolArgs | UserPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserPoolState | undefined;
            inputs["accountRecoverySetting"] = state ? state.accountRecoverySetting : undefined;
            inputs["adminCreateUserConfig"] = state ? state.adminCreateUserConfig : undefined;
            inputs["aliasAttributes"] = state ? state.aliasAttributes : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoVerifiedAttributes"] = state ? state.autoVerifiedAttributes : undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["customDomain"] = state ? state.customDomain : undefined;
            inputs["deviceConfiguration"] = state ? state.deviceConfiguration : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["emailConfiguration"] = state ? state.emailConfiguration : undefined;
            inputs["emailVerificationMessage"] = state ? state.emailVerificationMessage : undefined;
            inputs["emailVerificationSubject"] = state ? state.emailVerificationSubject : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["estimatedNumberOfUsers"] = state ? state.estimatedNumberOfUsers : undefined;
            inputs["lambdaConfig"] = state ? state.lambdaConfig : undefined;
            inputs["lastModifiedDate"] = state ? state.lastModifiedDate : undefined;
            inputs["mfaConfiguration"] = state ? state.mfaConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["passwordPolicy"] = state ? state.passwordPolicy : undefined;
            inputs["schemas"] = state ? state.schemas : undefined;
            inputs["smsAuthenticationMessage"] = state ? state.smsAuthenticationMessage : undefined;
            inputs["smsConfiguration"] = state ? state.smsConfiguration : undefined;
            inputs["smsVerificationMessage"] = state ? state.smsVerificationMessage : undefined;
            inputs["softwareTokenMfaConfiguration"] = state ? state.softwareTokenMfaConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["userPoolAddOns"] = state ? state.userPoolAddOns : undefined;
            inputs["usernameAttributes"] = state ? state.usernameAttributes : undefined;
            inputs["usernameConfiguration"] = state ? state.usernameConfiguration : undefined;
            inputs["verificationMessageTemplate"] = state ? state.verificationMessageTemplate : undefined;
        } else {
            const args = argsOrState as UserPoolArgs | undefined;
            inputs["accountRecoverySetting"] = args ? args.accountRecoverySetting : undefined;
            inputs["adminCreateUserConfig"] = args ? args.adminCreateUserConfig : undefined;
            inputs["aliasAttributes"] = args ? args.aliasAttributes : undefined;
            inputs["autoVerifiedAttributes"] = args ? args.autoVerifiedAttributes : undefined;
            inputs["deviceConfiguration"] = args ? args.deviceConfiguration : undefined;
            inputs["emailConfiguration"] = args ? args.emailConfiguration : undefined;
            inputs["emailVerificationMessage"] = args ? args.emailVerificationMessage : undefined;
            inputs["emailVerificationSubject"] = args ? args.emailVerificationSubject : undefined;
            inputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            inputs["mfaConfiguration"] = args ? args.mfaConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["passwordPolicy"] = args ? args.passwordPolicy : undefined;
            inputs["schemas"] = args ? args.schemas : undefined;
            inputs["smsAuthenticationMessage"] = args ? args.smsAuthenticationMessage : undefined;
            inputs["smsConfiguration"] = args ? args.smsConfiguration : undefined;
            inputs["smsVerificationMessage"] = args ? args.smsVerificationMessage : undefined;
            inputs["softwareTokenMfaConfiguration"] = args ? args.softwareTokenMfaConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userPoolAddOns"] = args ? args.userPoolAddOns : undefined;
            inputs["usernameAttributes"] = args ? args.usernameAttributes : undefined;
            inputs["usernameConfiguration"] = args ? args.usernameConfiguration : undefined;
            inputs["verificationMessageTemplate"] = args ? args.verificationMessageTemplate : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["customDomain"] = undefined /*out*/;
            inputs["domain"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["estimatedNumberOfUsers"] = undefined /*out*/;
            inputs["lastModifiedDate"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPool resources.
 */
export interface UserPoolState {
    /**
     * Configuration block to define which verified available method a user can use to recover their forgotten password. Detailed below.
     */
    accountRecoverySetting?: pulumi.Input<inputs.cognito.UserPoolAccountRecoverySetting>;
    /**
     * Configuration block for creating a new user profile. Detailed below.
     */
    adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Valid values: `phoneNumber`, `email`, or `preferredUsername`. Conflicts with `usernameAttributes`.
     */
    aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ARN of the user pool.
     */
    arn?: pulumi.Input<string>;
    /**
     * Attributes to be auto-verified. Valid values: `email`, `phoneNumber`.
     */
    autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date the user pool was created.
     */
    creationDate?: pulumi.Input<string>;
    /**
     * A custom domain name that you provide to Amazon Cognito. This parameter applies only if you use a custom domain to host the sign-up and sign-in pages for your application. For example: `auth.example.com`.
     */
    customDomain?: pulumi.Input<string>;
    /**
     * Configuration block for the user pool's device tracking. Detailed below.
     */
    deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    /**
     * Holds the domain prefix if the user pool has a domain associated with it.
     */
    domain?: pulumi.Input<string>;
    /**
     * Configuration block for configuring email. Detailed below.
     */
    emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    /**
     * String representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    emailVerificationMessage?: pulumi.Input<string>;
    /**
     * String representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    emailVerificationSubject?: pulumi.Input<string>;
    /**
     * Endpoint name of the user pool. Example format: `cognito-idp.REGION.amazonaws.com/xxxx_yyyyy`
     */
    endpoint?: pulumi.Input<string>;
    /**
     * A number estimating the size of the user pool.
     */
    estimatedNumberOfUsers?: pulumi.Input<number>;
    /**
     * Configuration block for the AWS Lambda triggers associated with the user pool. Detailed below.
     */
    lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    /**
     * Date the user pool was last modified.
     */
    lastModifiedDate?: pulumi.Input<string>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values are `OFF` (MFA Tokens are not required), `ON` (MFA is required for all users to sign in; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured), or `OPTIONAL` (MFA Will be required only for individual users who have MFA Enabled; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured).
     */
    mfaConfiguration?: pulumi.Input<string>;
    /**
     * Name of the attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration blocked for information about the user pool password policy. Detailed below.
     */
    passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    /**
     * Configuration block for the schema attributes of a user pool. Detailed below. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Attributes can be added, but not modified or removed. Maximum of 50 attributes.
     */
    schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    /**
     * String representing the SMS authentication message. The Message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    smsAuthenticationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection.
     */
    smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    /**
     * String representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    smsVerificationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    /**
     * Map of tags to assign to the User Pool. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features. Detailed below.
     */
    userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    /**
     * Whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block for username configuration. Detailed below.
     */
    usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    /**
     * Configuration block for verification message templates. Detailed below.
     */
    verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}

/**
 * The set of arguments for constructing a UserPool resource.
 */
export interface UserPoolArgs {
    /**
     * Configuration block to define which verified available method a user can use to recover their forgotten password. Detailed below.
     */
    accountRecoverySetting?: pulumi.Input<inputs.cognito.UserPoolAccountRecoverySetting>;
    /**
     * Configuration block for creating a new user profile. Detailed below.
     */
    adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Valid values: `phoneNumber`, `email`, or `preferredUsername`. Conflicts with `usernameAttributes`.
     */
    aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Attributes to be auto-verified. Valid values: `email`, `phoneNumber`.
     */
    autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block for the user pool's device tracking. Detailed below.
     */
    deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    /**
     * Configuration block for configuring email. Detailed below.
     */
    emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    /**
     * String representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    emailVerificationMessage?: pulumi.Input<string>;
    /**
     * String representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    emailVerificationSubject?: pulumi.Input<string>;
    /**
     * Configuration block for the AWS Lambda triggers associated with the user pool. Detailed below.
     */
    lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values are `OFF` (MFA Tokens are not required), `ON` (MFA is required for all users to sign in; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured), or `OPTIONAL` (MFA Will be required only for individual users who have MFA Enabled; requires at least one of `smsConfiguration` or `softwareTokenMfaConfiguration` to be configured).
     */
    mfaConfiguration?: pulumi.Input<string>;
    /**
     * Name of the attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration blocked for information about the user pool password policy. Detailed below.
     */
    passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    /**
     * Configuration block for the schema attributes of a user pool. Detailed below. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Attributes can be added, but not modified or removed. Maximum of 50 attributes.
     */
    schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    /**
     * String representing the SMS authentication message. The Message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    smsAuthenticationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection.
     */
    smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    /**
     * String representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    smsVerificationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    /**
     * Map of tags to assign to the User Pool. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features. Detailed below.
     */
    userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    /**
     * Whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block for username configuration. Detailed below.
     */
    usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    /**
     * Configuration block for verification message templates. Detailed below.
     */
    verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}
