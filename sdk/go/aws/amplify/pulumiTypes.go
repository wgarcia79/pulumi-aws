// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppAutoBranchCreationConfig struct {
	// The basic authorization credentials for the autocreated branch.
	BasicAuthCredentials *string `pulumi:"basicAuthCredentials"`
	// The build specification (build spec) for the autocreated branch.
	BuildSpec *string `pulumi:"buildSpec"`
	// Enables auto building for the autocreated branch.
	EnableAutoBuild *bool `pulumi:"enableAutoBuild"`
	// Enables basic authorization for the autocreated branch.
	EnableBasicAuth *bool `pulumi:"enableBasicAuth"`
	// Enables performance mode for the branch.
	EnablePerformanceMode *bool `pulumi:"enablePerformanceMode"`
	// Enables pull request previews for the autocreated branch.
	EnablePullRequestPreview *bool `pulumi:"enablePullRequestPreview"`
	// The environment variables for the autocreated branch.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The framework for the autocreated branch.
	Framework *string `pulumi:"framework"`
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string `pulumi:"pullRequestEnvironmentName"`
	// Describes the current stage for the autocreated branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage *string `pulumi:"stage"`
}

// AppAutoBranchCreationConfigInput is an input type that accepts AppAutoBranchCreationConfigArgs and AppAutoBranchCreationConfigOutput values.
// You can construct a concrete instance of `AppAutoBranchCreationConfigInput` via:
//
//          AppAutoBranchCreationConfigArgs{...}
type AppAutoBranchCreationConfigInput interface {
	pulumi.Input

	ToAppAutoBranchCreationConfigOutput() AppAutoBranchCreationConfigOutput
	ToAppAutoBranchCreationConfigOutputWithContext(context.Context) AppAutoBranchCreationConfigOutput
}

type AppAutoBranchCreationConfigArgs struct {
	// The basic authorization credentials for the autocreated branch.
	BasicAuthCredentials pulumi.StringPtrInput `pulumi:"basicAuthCredentials"`
	// The build specification (build spec) for the autocreated branch.
	BuildSpec pulumi.StringPtrInput `pulumi:"buildSpec"`
	// Enables auto building for the autocreated branch.
	EnableAutoBuild pulumi.BoolPtrInput `pulumi:"enableAutoBuild"`
	// Enables basic authorization for the autocreated branch.
	EnableBasicAuth pulumi.BoolPtrInput `pulumi:"enableBasicAuth"`
	// Enables performance mode for the branch.
	EnablePerformanceMode pulumi.BoolPtrInput `pulumi:"enablePerformanceMode"`
	// Enables pull request previews for the autocreated branch.
	EnablePullRequestPreview pulumi.BoolPtrInput `pulumi:"enablePullRequestPreview"`
	// The environment variables for the autocreated branch.
	EnvironmentVariables pulumi.StringMapInput `pulumi:"environmentVariables"`
	// The framework for the autocreated branch.
	Framework pulumi.StringPtrInput `pulumi:"framework"`
	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName pulumi.StringPtrInput `pulumi:"pullRequestEnvironmentName"`
	// Describes the current stage for the autocreated branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
	Stage pulumi.StringPtrInput `pulumi:"stage"`
}

func (AppAutoBranchCreationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAutoBranchCreationConfig)(nil)).Elem()
}

func (i AppAutoBranchCreationConfigArgs) ToAppAutoBranchCreationConfigOutput() AppAutoBranchCreationConfigOutput {
	return i.ToAppAutoBranchCreationConfigOutputWithContext(context.Background())
}

func (i AppAutoBranchCreationConfigArgs) ToAppAutoBranchCreationConfigOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAutoBranchCreationConfigOutput)
}

func (i AppAutoBranchCreationConfigArgs) ToAppAutoBranchCreationConfigPtrOutput() AppAutoBranchCreationConfigPtrOutput {
	return i.ToAppAutoBranchCreationConfigPtrOutputWithContext(context.Background())
}

func (i AppAutoBranchCreationConfigArgs) ToAppAutoBranchCreationConfigPtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAutoBranchCreationConfigOutput).ToAppAutoBranchCreationConfigPtrOutputWithContext(ctx)
}

// AppAutoBranchCreationConfigPtrInput is an input type that accepts AppAutoBranchCreationConfigArgs, AppAutoBranchCreationConfigPtr and AppAutoBranchCreationConfigPtrOutput values.
// You can construct a concrete instance of `AppAutoBranchCreationConfigPtrInput` via:
//
//          AppAutoBranchCreationConfigArgs{...}
//
//  or:
//
//          nil
type AppAutoBranchCreationConfigPtrInput interface {
	pulumi.Input

	ToAppAutoBranchCreationConfigPtrOutput() AppAutoBranchCreationConfigPtrOutput
	ToAppAutoBranchCreationConfigPtrOutputWithContext(context.Context) AppAutoBranchCreationConfigPtrOutput
}

type appAutoBranchCreationConfigPtrType AppAutoBranchCreationConfigArgs

func AppAutoBranchCreationConfigPtr(v *AppAutoBranchCreationConfigArgs) AppAutoBranchCreationConfigPtrInput {
	return (*appAutoBranchCreationConfigPtrType)(v)
}

func (*appAutoBranchCreationConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAutoBranchCreationConfig)(nil)).Elem()
}

func (i *appAutoBranchCreationConfigPtrType) ToAppAutoBranchCreationConfigPtrOutput() AppAutoBranchCreationConfigPtrOutput {
	return i.ToAppAutoBranchCreationConfigPtrOutputWithContext(context.Background())
}

func (i *appAutoBranchCreationConfigPtrType) ToAppAutoBranchCreationConfigPtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAutoBranchCreationConfigPtrOutput)
}

type AppAutoBranchCreationConfigOutput struct{ *pulumi.OutputState }

func (AppAutoBranchCreationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppAutoBranchCreationConfig)(nil)).Elem()
}

func (o AppAutoBranchCreationConfigOutput) ToAppAutoBranchCreationConfigOutput() AppAutoBranchCreationConfigOutput {
	return o
}

func (o AppAutoBranchCreationConfigOutput) ToAppAutoBranchCreationConfigOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigOutput {
	return o
}

func (o AppAutoBranchCreationConfigOutput) ToAppAutoBranchCreationConfigPtrOutput() AppAutoBranchCreationConfigPtrOutput {
	return o.ToAppAutoBranchCreationConfigPtrOutputWithContext(context.Background())
}

func (o AppAutoBranchCreationConfigOutput) ToAppAutoBranchCreationConfigPtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppAutoBranchCreationConfig) *AppAutoBranchCreationConfig {
		return &v
	}).(AppAutoBranchCreationConfigPtrOutput)
}

// The basic authorization credentials for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) BasicAuthCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *string { return v.BasicAuthCredentials }).(pulumi.StringPtrOutput)
}

// The build specification (build spec) for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *string { return v.BuildSpec }).(pulumi.StringPtrOutput)
}

// Enables auto building for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) EnableAutoBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *bool { return v.EnableAutoBuild }).(pulumi.BoolPtrOutput)
}

// Enables basic authorization for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) EnableBasicAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *bool { return v.EnableBasicAuth }).(pulumi.BoolPtrOutput)
}

// Enables performance mode for the branch.
func (o AppAutoBranchCreationConfigOutput) EnablePerformanceMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *bool { return v.EnablePerformanceMode }).(pulumi.BoolPtrOutput)
}

// Enables pull request previews for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) EnablePullRequestPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *bool { return v.EnablePullRequestPreview }).(pulumi.BoolPtrOutput)
}

// The environment variables for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The framework for the autocreated branch.
func (o AppAutoBranchCreationConfigOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *string { return v.Framework }).(pulumi.StringPtrOutput)
}

// The Amplify environment name for the pull request.
func (o AppAutoBranchCreationConfigOutput) PullRequestEnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *string { return v.PullRequestEnvironmentName }).(pulumi.StringPtrOutput)
}

// Describes the current stage for the autocreated branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
func (o AppAutoBranchCreationConfigOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppAutoBranchCreationConfig) *string { return v.Stage }).(pulumi.StringPtrOutput)
}

type AppAutoBranchCreationConfigPtrOutput struct{ *pulumi.OutputState }

func (AppAutoBranchCreationConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAutoBranchCreationConfig)(nil)).Elem()
}

func (o AppAutoBranchCreationConfigPtrOutput) ToAppAutoBranchCreationConfigPtrOutput() AppAutoBranchCreationConfigPtrOutput {
	return o
}

func (o AppAutoBranchCreationConfigPtrOutput) ToAppAutoBranchCreationConfigPtrOutputWithContext(ctx context.Context) AppAutoBranchCreationConfigPtrOutput {
	return o
}

func (o AppAutoBranchCreationConfigPtrOutput) Elem() AppAutoBranchCreationConfigOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) AppAutoBranchCreationConfig {
		if v != nil {
			return *v
		}
		var ret AppAutoBranchCreationConfig
		return ret
	}).(AppAutoBranchCreationConfigOutput)
}

// The basic authorization credentials for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) BasicAuthCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *string {
		if v == nil {
			return nil
		}
		return v.BasicAuthCredentials
	}).(pulumi.StringPtrOutput)
}

// The build specification (build spec) for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *string {
		if v == nil {
			return nil
		}
		return v.BuildSpec
	}).(pulumi.StringPtrOutput)
}

// Enables auto building for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) EnableAutoBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutoBuild
	}).(pulumi.BoolPtrOutput)
}

// Enables basic authorization for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) EnableBasicAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBasicAuth
	}).(pulumi.BoolPtrOutput)
}

// Enables performance mode for the branch.
func (o AppAutoBranchCreationConfigPtrOutput) EnablePerformanceMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePerformanceMode
	}).(pulumi.BoolPtrOutput)
}

// Enables pull request previews for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) EnablePullRequestPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePullRequestPreview
	}).(pulumi.BoolPtrOutput)
}

// The environment variables for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

// The framework for the autocreated branch.
func (o AppAutoBranchCreationConfigPtrOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *string {
		if v == nil {
			return nil
		}
		return v.Framework
	}).(pulumi.StringPtrOutput)
}

// The Amplify environment name for the pull request.
func (o AppAutoBranchCreationConfigPtrOutput) PullRequestEnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *string {
		if v == nil {
			return nil
		}
		return v.PullRequestEnvironmentName
	}).(pulumi.StringPtrOutput)
}

// Describes the current stage for the autocreated branch. Valid values: `PRODUCTION`, `BETA`, `DEVELOPMENT`, `EXPERIMENTAL`, `PULL_REQUEST`.
func (o AppAutoBranchCreationConfigPtrOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAutoBranchCreationConfig) *string {
		if v == nil {
			return nil
		}
		return v.Stage
	}).(pulumi.StringPtrOutput)
}

type AppCustomRule struct {
	// The condition for a URL rewrite or redirect rule, such as a country code.
	Condition *string `pulumi:"condition"`
	// The source pattern for a URL rewrite or redirect rule.
	Source string `pulumi:"source"`
	// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status *string `pulumi:"status"`
	// The target pattern for a URL rewrite or redirect rule.
	Target string `pulumi:"target"`
}

// AppCustomRuleInput is an input type that accepts AppCustomRuleArgs and AppCustomRuleOutput values.
// You can construct a concrete instance of `AppCustomRuleInput` via:
//
//          AppCustomRuleArgs{...}
type AppCustomRuleInput interface {
	pulumi.Input

	ToAppCustomRuleOutput() AppCustomRuleOutput
	ToAppCustomRuleOutputWithContext(context.Context) AppCustomRuleOutput
}

type AppCustomRuleArgs struct {
	// The condition for a URL rewrite or redirect rule, such as a country code.
	Condition pulumi.StringPtrInput `pulumi:"condition"`
	// The source pattern for a URL rewrite or redirect rule.
	Source pulumi.StringInput `pulumi:"source"`
	// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The target pattern for a URL rewrite or redirect rule.
	Target pulumi.StringInput `pulumi:"target"`
}

func (AppCustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCustomRule)(nil)).Elem()
}

func (i AppCustomRuleArgs) ToAppCustomRuleOutput() AppCustomRuleOutput {
	return i.ToAppCustomRuleOutputWithContext(context.Background())
}

func (i AppCustomRuleArgs) ToAppCustomRuleOutputWithContext(ctx context.Context) AppCustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCustomRuleOutput)
}

// AppCustomRuleArrayInput is an input type that accepts AppCustomRuleArray and AppCustomRuleArrayOutput values.
// You can construct a concrete instance of `AppCustomRuleArrayInput` via:
//
//          AppCustomRuleArray{ AppCustomRuleArgs{...} }
type AppCustomRuleArrayInput interface {
	pulumi.Input

	ToAppCustomRuleArrayOutput() AppCustomRuleArrayOutput
	ToAppCustomRuleArrayOutputWithContext(context.Context) AppCustomRuleArrayOutput
}

type AppCustomRuleArray []AppCustomRuleInput

func (AppCustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppCustomRule)(nil)).Elem()
}

func (i AppCustomRuleArray) ToAppCustomRuleArrayOutput() AppCustomRuleArrayOutput {
	return i.ToAppCustomRuleArrayOutputWithContext(context.Background())
}

func (i AppCustomRuleArray) ToAppCustomRuleArrayOutputWithContext(ctx context.Context) AppCustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCustomRuleArrayOutput)
}

type AppCustomRuleOutput struct{ *pulumi.OutputState }

func (AppCustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCustomRule)(nil)).Elem()
}

func (o AppCustomRuleOutput) ToAppCustomRuleOutput() AppCustomRuleOutput {
	return o
}

func (o AppCustomRuleOutput) ToAppCustomRuleOutputWithContext(ctx context.Context) AppCustomRuleOutput {
	return o
}

// The condition for a URL rewrite or redirect rule, such as a country code.
func (o AppCustomRuleOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppCustomRule) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

// The source pattern for a URL rewrite or redirect rule.
func (o AppCustomRuleOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v AppCustomRule) string { return v.Source }).(pulumi.StringOutput)
}

// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
func (o AppCustomRuleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppCustomRule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The target pattern for a URL rewrite or redirect rule.
func (o AppCustomRuleOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v AppCustomRule) string { return v.Target }).(pulumi.StringOutput)
}

type AppCustomRuleArrayOutput struct{ *pulumi.OutputState }

func (AppCustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppCustomRule)(nil)).Elem()
}

func (o AppCustomRuleArrayOutput) ToAppCustomRuleArrayOutput() AppCustomRuleArrayOutput {
	return o
}

func (o AppCustomRuleArrayOutput) ToAppCustomRuleArrayOutputWithContext(ctx context.Context) AppCustomRuleArrayOutput {
	return o
}

func (o AppCustomRuleArrayOutput) Index(i pulumi.IntInput) AppCustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppCustomRule {
		return vs[0].([]AppCustomRule)[vs[1].(int)]
	}).(AppCustomRuleOutput)
}

type AppProductionBranch struct {
	// The branch name for the production branch.
	BranchName *string `pulumi:"branchName"`
	// The last deploy time of the production branch.
	LastDeployTime *string `pulumi:"lastDeployTime"`
	// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status *string `pulumi:"status"`
	// The thumbnail URL for the production branch.
	ThumbnailUrl *string `pulumi:"thumbnailUrl"`
}

// AppProductionBranchInput is an input type that accepts AppProductionBranchArgs and AppProductionBranchOutput values.
// You can construct a concrete instance of `AppProductionBranchInput` via:
//
//          AppProductionBranchArgs{...}
type AppProductionBranchInput interface {
	pulumi.Input

	ToAppProductionBranchOutput() AppProductionBranchOutput
	ToAppProductionBranchOutputWithContext(context.Context) AppProductionBranchOutput
}

type AppProductionBranchArgs struct {
	// The branch name for the production branch.
	BranchName pulumi.StringPtrInput `pulumi:"branchName"`
	// The last deploy time of the production branch.
	LastDeployTime pulumi.StringPtrInput `pulumi:"lastDeployTime"`
	// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The thumbnail URL for the production branch.
	ThumbnailUrl pulumi.StringPtrInput `pulumi:"thumbnailUrl"`
}

func (AppProductionBranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppProductionBranch)(nil)).Elem()
}

func (i AppProductionBranchArgs) ToAppProductionBranchOutput() AppProductionBranchOutput {
	return i.ToAppProductionBranchOutputWithContext(context.Background())
}

func (i AppProductionBranchArgs) ToAppProductionBranchOutputWithContext(ctx context.Context) AppProductionBranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProductionBranchOutput)
}

// AppProductionBranchArrayInput is an input type that accepts AppProductionBranchArray and AppProductionBranchArrayOutput values.
// You can construct a concrete instance of `AppProductionBranchArrayInput` via:
//
//          AppProductionBranchArray{ AppProductionBranchArgs{...} }
type AppProductionBranchArrayInput interface {
	pulumi.Input

	ToAppProductionBranchArrayOutput() AppProductionBranchArrayOutput
	ToAppProductionBranchArrayOutputWithContext(context.Context) AppProductionBranchArrayOutput
}

type AppProductionBranchArray []AppProductionBranchInput

func (AppProductionBranchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppProductionBranch)(nil)).Elem()
}

func (i AppProductionBranchArray) ToAppProductionBranchArrayOutput() AppProductionBranchArrayOutput {
	return i.ToAppProductionBranchArrayOutputWithContext(context.Background())
}

func (i AppProductionBranchArray) ToAppProductionBranchArrayOutputWithContext(ctx context.Context) AppProductionBranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppProductionBranchArrayOutput)
}

type AppProductionBranchOutput struct{ *pulumi.OutputState }

func (AppProductionBranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppProductionBranch)(nil)).Elem()
}

func (o AppProductionBranchOutput) ToAppProductionBranchOutput() AppProductionBranchOutput {
	return o
}

func (o AppProductionBranchOutput) ToAppProductionBranchOutputWithContext(ctx context.Context) AppProductionBranchOutput {
	return o
}

// The branch name for the production branch.
func (o AppProductionBranchOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppProductionBranch) *string { return v.BranchName }).(pulumi.StringPtrOutput)
}

// The last deploy time of the production branch.
func (o AppProductionBranchOutput) LastDeployTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppProductionBranch) *string { return v.LastDeployTime }).(pulumi.StringPtrOutput)
}

// The status code for a URL rewrite or redirect rule. Valid values: `200`, `301`, `302`, `404`, `404-200`.
func (o AppProductionBranchOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppProductionBranch) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The thumbnail URL for the production branch.
func (o AppProductionBranchOutput) ThumbnailUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppProductionBranch) *string { return v.ThumbnailUrl }).(pulumi.StringPtrOutput)
}

type AppProductionBranchArrayOutput struct{ *pulumi.OutputState }

func (AppProductionBranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppProductionBranch)(nil)).Elem()
}

func (o AppProductionBranchArrayOutput) ToAppProductionBranchArrayOutput() AppProductionBranchArrayOutput {
	return o
}

func (o AppProductionBranchArrayOutput) ToAppProductionBranchArrayOutputWithContext(ctx context.Context) AppProductionBranchArrayOutput {
	return o
}

func (o AppProductionBranchArrayOutput) Index(i pulumi.IntInput) AppProductionBranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppProductionBranch {
		return vs[0].([]AppProductionBranch)[vs[1].(int)]
	}).(AppProductionBranchOutput)
}

type DomainAssociationSubDomain struct {
	// The branch name setting for the subdomain.
	BranchName string `pulumi:"branchName"`
	// The DNS record for the subdomain.
	DnsRecord *string `pulumi:"dnsRecord"`
	// The prefix setting for the subdomain.
	Prefix string `pulumi:"prefix"`
	// The verified status of the subdomain.
	Verified *bool `pulumi:"verified"`
}

// DomainAssociationSubDomainInput is an input type that accepts DomainAssociationSubDomainArgs and DomainAssociationSubDomainOutput values.
// You can construct a concrete instance of `DomainAssociationSubDomainInput` via:
//
//          DomainAssociationSubDomainArgs{...}
type DomainAssociationSubDomainInput interface {
	pulumi.Input

	ToDomainAssociationSubDomainOutput() DomainAssociationSubDomainOutput
	ToDomainAssociationSubDomainOutputWithContext(context.Context) DomainAssociationSubDomainOutput
}

type DomainAssociationSubDomainArgs struct {
	// The branch name setting for the subdomain.
	BranchName pulumi.StringInput `pulumi:"branchName"`
	// The DNS record for the subdomain.
	DnsRecord pulumi.StringPtrInput `pulumi:"dnsRecord"`
	// The prefix setting for the subdomain.
	Prefix pulumi.StringInput `pulumi:"prefix"`
	// The verified status of the subdomain.
	Verified pulumi.BoolPtrInput `pulumi:"verified"`
}

func (DomainAssociationSubDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAssociationSubDomain)(nil)).Elem()
}

func (i DomainAssociationSubDomainArgs) ToDomainAssociationSubDomainOutput() DomainAssociationSubDomainOutput {
	return i.ToDomainAssociationSubDomainOutputWithContext(context.Background())
}

func (i DomainAssociationSubDomainArgs) ToDomainAssociationSubDomainOutputWithContext(ctx context.Context) DomainAssociationSubDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationSubDomainOutput)
}

// DomainAssociationSubDomainArrayInput is an input type that accepts DomainAssociationSubDomainArray and DomainAssociationSubDomainArrayOutput values.
// You can construct a concrete instance of `DomainAssociationSubDomainArrayInput` via:
//
//          DomainAssociationSubDomainArray{ DomainAssociationSubDomainArgs{...} }
type DomainAssociationSubDomainArrayInput interface {
	pulumi.Input

	ToDomainAssociationSubDomainArrayOutput() DomainAssociationSubDomainArrayOutput
	ToDomainAssociationSubDomainArrayOutputWithContext(context.Context) DomainAssociationSubDomainArrayOutput
}

type DomainAssociationSubDomainArray []DomainAssociationSubDomainInput

func (DomainAssociationSubDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainAssociationSubDomain)(nil)).Elem()
}

func (i DomainAssociationSubDomainArray) ToDomainAssociationSubDomainArrayOutput() DomainAssociationSubDomainArrayOutput {
	return i.ToDomainAssociationSubDomainArrayOutputWithContext(context.Background())
}

func (i DomainAssociationSubDomainArray) ToDomainAssociationSubDomainArrayOutputWithContext(ctx context.Context) DomainAssociationSubDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAssociationSubDomainArrayOutput)
}

type DomainAssociationSubDomainOutput struct{ *pulumi.OutputState }

func (DomainAssociationSubDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainAssociationSubDomain)(nil)).Elem()
}

func (o DomainAssociationSubDomainOutput) ToDomainAssociationSubDomainOutput() DomainAssociationSubDomainOutput {
	return o
}

func (o DomainAssociationSubDomainOutput) ToDomainAssociationSubDomainOutputWithContext(ctx context.Context) DomainAssociationSubDomainOutput {
	return o
}

// The branch name setting for the subdomain.
func (o DomainAssociationSubDomainOutput) BranchName() pulumi.StringOutput {
	return o.ApplyT(func(v DomainAssociationSubDomain) string { return v.BranchName }).(pulumi.StringOutput)
}

// The DNS record for the subdomain.
func (o DomainAssociationSubDomainOutput) DnsRecord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainAssociationSubDomain) *string { return v.DnsRecord }).(pulumi.StringPtrOutput)
}

// The prefix setting for the subdomain.
func (o DomainAssociationSubDomainOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v DomainAssociationSubDomain) string { return v.Prefix }).(pulumi.StringOutput)
}

// The verified status of the subdomain.
func (o DomainAssociationSubDomainOutput) Verified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DomainAssociationSubDomain) *bool { return v.Verified }).(pulumi.BoolPtrOutput)
}

type DomainAssociationSubDomainArrayOutput struct{ *pulumi.OutputState }

func (DomainAssociationSubDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DomainAssociationSubDomain)(nil)).Elem()
}

func (o DomainAssociationSubDomainArrayOutput) ToDomainAssociationSubDomainArrayOutput() DomainAssociationSubDomainArrayOutput {
	return o
}

func (o DomainAssociationSubDomainArrayOutput) ToDomainAssociationSubDomainArrayOutputWithContext(ctx context.Context) DomainAssociationSubDomainArrayOutput {
	return o
}

func (o DomainAssociationSubDomainArrayOutput) Index(i pulumi.IntInput) DomainAssociationSubDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DomainAssociationSubDomain {
		return vs[0].([]DomainAssociationSubDomain)[vs[1].(int)]
	}).(DomainAssociationSubDomainOutput)
}

func init() {
	pulumi.RegisterOutputType(AppAutoBranchCreationConfigOutput{})
	pulumi.RegisterOutputType(AppAutoBranchCreationConfigPtrOutput{})
	pulumi.RegisterOutputType(AppCustomRuleOutput{})
	pulumi.RegisterOutputType(AppCustomRuleArrayOutput{})
	pulumi.RegisterOutputType(AppProductionBranchOutput{})
	pulumi.RegisterOutputType(AppProductionBranchArrayOutput{})
	pulumi.RegisterOutputType(DomainAssociationSubDomainOutput{})
	pulumi.RegisterOutputType(DomainAssociationSubDomainArrayOutput{})
}
