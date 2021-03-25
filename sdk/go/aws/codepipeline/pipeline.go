// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codepipeline

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodePipeline.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codepipeline"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codestarconnections"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := codestarconnections.NewConnection(ctx, "example", &codestarconnections.ConnectionArgs{
// 			ProviderType: pulumi.String("GitHub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		codepipelineBucket, err := s3.NewBucket(ctx, "codepipelineBucket", &s3.BucketArgs{
// 			Acl: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		codepipelineRole, err := iam.NewRole(ctx, "codepipelineRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"codepipeline.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		s3kmskey, err := kms.LookupAlias(ctx, &kms.LookupAliasArgs{
// 			Name: "alias/myKmsKey",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = codepipeline.NewPipeline(ctx, "codepipeline", &codepipeline.PipelineArgs{
// 			RoleArn: codepipelineRole.Arn,
// 			ArtifactStore: &codepipeline.PipelineArtifactStoreArgs{
// 				Location: codepipelineBucket.Bucket,
// 				Type:     pulumi.String("S3"),
// 				EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
// 					Id:   pulumi.String(s3kmskey.Arn),
// 					Type: pulumi.String("KMS"),
// 				},
// 			},
// 			Stages: codepipeline.PipelineStageArray{
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Source"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Source"),
// 							Category: pulumi.String("Source"),
// 							Owner:    pulumi.String("AWS"),
// 							Provider: pulumi.String("CodeStarSourceConnection"),
// 							Version:  pulumi.String("1"),
// 							OutputArtifacts: pulumi.StringArray{
// 								pulumi.String("source_output"),
// 							},
// 							Configuration: pulumi.StringMap{
// 								"ConnectionArn":    example.Arn,
// 								"FullRepositoryId": pulumi.String("my-organization/example"),
// 								"BranchName":       pulumi.String("main"),
// 							},
// 						},
// 					},
// 				},
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Build"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Build"),
// 							Category: pulumi.String("Build"),
// 							Owner:    pulumi.String("AWS"),
// 							Provider: pulumi.String("CodeBuild"),
// 							InputArtifacts: pulumi.StringArray{
// 								pulumi.String("source_output"),
// 							},
// 							OutputArtifacts: pulumi.StringArray{
// 								pulumi.String("build_output"),
// 							},
// 							Version: pulumi.String("1"),
// 							Configuration: pulumi.StringMap{
// 								"ProjectName": pulumi.String("test"),
// 							},
// 						},
// 					},
// 				},
// 				&codepipeline.PipelineStageArgs{
// 					Name: pulumi.String("Deploy"),
// 					Actions: codepipeline.PipelineStageActionArray{
// 						&codepipeline.PipelineStageActionArgs{
// 							Name:     pulumi.String("Deploy"),
// 							Category: pulumi.String("Deploy"),
// 							Owner:    pulumi.String("AWS"),
// 							Provider: pulumi.String("CloudFormation"),
// 							InputArtifacts: pulumi.StringArray{
// 								pulumi.String("build_output"),
// 							},
// 							Version: pulumi.String("1"),
// 							Configuration: pulumi.StringMap{
// 								"ActionMode":     pulumi.String("REPLACE_ON_FAILURE"),
// 								"Capabilities":   pulumi.String("CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM"),
// 								"OutputFileName": pulumi.String("CreateStackOutput.json"),
// 								"StackName":      pulumi.String("MyStack"),
// 								"TemplatePath":   pulumi.String("build_output::sam-templated.yaml"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "codepipelinePolicy", &iam.RolePolicyArgs{
// 			Role: codepipelineRole.ID(),
// 			Policy: pulumi.All(codepipelineBucket.Arn, codepipelineBucket.Arn).ApplyT(func(_args []interface{}) (string, error) {
// 				codepipelineBucketArn := _args[0].(string)
// 				codepipelineBucketArn1 := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\":\"Allow\",\n", "      \"Action\": [\n", "        \"s3:GetObject\",\n", "        \"s3:GetObjectVersion\",\n", "        \"s3:GetBucketVersioning\",\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Resource\": [\n", "        \"", codepipelineBucketArn, "\",\n", "        \"", codepipelineBucketArn1, "/*\"\n", "      ]\n", "    },\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"codebuild:BatchGetBuilds\",\n", "        \"codebuild:StartBuild\"\n", "      ],\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// CodePipelines can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:codepipeline/pipeline:Pipeline foo example
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	// The codepipeline ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// One or more artifactStore blocks. Artifact stores are documented below.
	ArtifactStore PipelineArtifactStoreOutput `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A stage block. Stages are documented below.
	Stages PipelineStageArrayOutput `pulumi:"stages"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArtifactStore == nil {
		return nil, errors.New("invalid value for required argument 'ArtifactStore'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.Stages == nil {
		return nil, errors.New("invalid value for required argument 'Stages'")
	}
	var resource Pipeline
	err := ctx.RegisterResource("aws:codepipeline/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("aws:codepipeline/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// The codepipeline ARN.
	Arn *string `pulumi:"arn"`
	// One or more artifactStore blocks. Artifact stores are documented below.
	ArtifactStore *PipelineArtifactStore `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// A stage block. Stages are documented below.
	Stages []PipelineStage `pulumi:"stages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type PipelineState struct {
	// The codepipeline ARN.
	Arn pulumi.StringPtrInput
	// One or more artifactStore blocks. Artifact stores are documented below.
	ArtifactStore PipelineArtifactStorePtrInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringPtrInput
	// A stage block. Stages are documented below.
	Stages PipelineStageArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// One or more artifactStore blocks. Artifact stores are documented below.
	ArtifactStore PipelineArtifactStore `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// A stage block. Stages are documented below.
	Stages []PipelineStage `pulumi:"stages"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// One or more artifactStore blocks. Artifact stores are documented below.
	ArtifactStore PipelineArtifactStoreInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringInput
	// A stage block. Stages are documented below.
	Stages PipelineStageArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

func (i *Pipeline) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePtrOutput)
}

type PipelinePtrInput interface {
	pulumi.Input

	ToPipelinePtrOutput() PipelinePtrOutput
	ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput
}

type pipelinePtrType PipelineArgs

func (*pipelinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (i *pipelinePtrType) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *pipelinePtrType) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePtrOutput)
}

// PipelineArrayInput is an input type that accepts PipelineArray and PipelineArrayOutput values.
// You can construct a concrete instance of `PipelineArrayInput` via:
//
//          PipelineArray{ PipelineArgs{...} }
type PipelineArrayInput interface {
	pulumi.Input

	ToPipelineArrayOutput() PipelineArrayOutput
	ToPipelineArrayOutputWithContext(context.Context) PipelineArrayOutput
}

type PipelineArray []PipelineInput

func (PipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Pipeline)(nil))
}

func (i PipelineArray) ToPipelineArrayOutput() PipelineArrayOutput {
	return i.ToPipelineArrayOutputWithContext(context.Background())
}

func (i PipelineArray) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineArrayOutput)
}

// PipelineMapInput is an input type that accepts PipelineMap and PipelineMapOutput values.
// You can construct a concrete instance of `PipelineMapInput` via:
//
//          PipelineMap{ "key": PipelineArgs{...} }
type PipelineMapInput interface {
	pulumi.Input

	ToPipelineMapOutput() PipelineMapOutput
	ToPipelineMapOutputWithContext(context.Context) PipelineMapOutput
}

type PipelineMap map[string]PipelineInput

func (PipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Pipeline)(nil))
}

func (i PipelineMap) ToPipelineMapOutput() PipelineMapOutput {
	return i.ToPipelineMapOutputWithContext(context.Background())
}

func (i PipelineMap) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineMapOutput)
}

type PipelineOutput struct {
	*pulumi.OutputState
}

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o.ToPipelinePtrOutputWithContext(context.Background())
}

func (o PipelineOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o.ApplyT(func(v Pipeline) *Pipeline {
		return &v
	}).(PipelinePtrOutput)
}

type PipelinePtrOutput struct {
	*pulumi.OutputState
}

func (PipelinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (o PipelinePtrOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o
}

func (o PipelinePtrOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o
}

type PipelineArrayOutput struct{ *pulumi.OutputState }

func (PipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipeline)(nil))
}

func (o PipelineArrayOutput) ToPipelineArrayOutput() PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) Index(i pulumi.IntInput) PipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].([]Pipeline)[vs[1].(int)]
	}).(PipelineOutput)
}

type PipelineMapOutput struct{ *pulumi.OutputState }

func (PipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipeline)(nil))
}

func (o PipelineMapOutput) ToPipelineMapOutput() PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) MapIndex(k pulumi.StringInput) PipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].(map[string]Pipeline)[vs[1].(string)]
	}).(PipelineOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
	pulumi.RegisterOutputType(PipelinePtrOutput{})
	pulumi.RegisterOutputType(PipelineArrayOutput{})
	pulumi.RegisterOutputType(PipelineMapOutput{})
}
