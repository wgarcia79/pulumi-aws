// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of cognito user pools.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cognito"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		selectedRestApi, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
// 			Name: _var.Api_gateway_name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		selectedUserPools, err := cognito.GetUserPools(ctx, &cognito.GetUserPoolsArgs{
// 			Name: _var.Cognito_user_pool_name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewAuthorizer(ctx, "cognito", &apigateway.AuthorizerArgs{
// 			Type:         pulumi.String("COGNITO_USER_POOLS"),
// 			RestApi:      pulumi.String(selectedRestApi.Id),
// 			ProviderArns: interface{}(selectedUserPools.Arns),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetUserPools(ctx *pulumi.Context, args *GetUserPoolsArgs, opts ...pulumi.InvokeOption) (*GetUserPoolsResult, error) {
	var rv GetUserPoolsResult
	err := ctx.Invoke("aws:cognito/getUserPools:getUserPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserPools.
type GetUserPoolsArgs struct {
	// Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getUserPools.
type GetUserPoolsResult struct {
	Arns []string `pulumi:"arns"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of cognito user pool ids.
	Ids  []string `pulumi:"ids"`
	Name string   `pulumi:"name"`
}
