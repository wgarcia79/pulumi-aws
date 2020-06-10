module github.com/pulumi/pulumi-aws/provider/v2

go 1.14

require (
	github.com/hashicorp/aws-sdk-go-base v0.4.0
	github.com/hashicorp/terraform-plugin-sdk v1.13.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.4.1-0.20200608011815-6feeb51f2d39
	github.com/pulumi/pulumi/sdk/v2 v2.3.1-0.20200608163628-76aede990ebd
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20191010190908-1261a98537f2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pkg
	github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
	github.com/pulumi/tf2pulumi => ../../tf2pulumi
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20200605145502-fb12b583285e
)
