// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import { asset } from "@pulumi/pulumi";
import { Output } from "@pulumi/pulumi";

let region = aws.config.requireRegion();
let stageName = "dev";

///////////////////
// DynamoDB Table
///////////////////
let usersTable = new aws.dynamodb.Table("UsersDynamoTable", {
    attribute: [
        { name: "userId", type: "S" },
    ],
    hashKey: "userId",
    readCapacity: 1,
    writeCapacity: 1,
});

///////////////////
// Lambda Function
///////////////////
let policy = {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "lambda.amazonaws.com",
            },
            "Effect": "Allow",
            "Sid": "",
        },
    ],
};

let role = new aws.iam.Role("mylambda-role", {
    assumeRolePolicy: JSON.stringify(policy),
});
let fullAccess = new aws.iam.RolePolicyAttachment("mylambda-access", {
    role: role,
    policyArn: aws.iam.AWSLambdaFullAccess,
});

// create lambda with sources in ./app folder, set environment variable
let lambda = new aws.lambda.Function("mylambda", {
    role: role.arn,
    handler: "index.handler",
    runtime: aws.lambda.NodeJS6d10Runtime,
    code: new asset.AssetArchive({
        "index.js": new asset.FileAsset("./app/index.js"),
        "node_modules": new asset.FileArchive("./app/node_modules")
    }),
    environment:  [ { 
        variables: {
            "USERS_TABLE": usersTable.name
        }
    } ]
});

///////////////////
// APIGateway RestAPI
///////////////////

function swaggerSpecForProxy(lambda: aws.lambda.Function): Output<string> {
    let swaggerSpec = 
        lambda.arn.apply(arn => 
            JSON.stringify({
                swagger: "2.0",
                info: { title: "donnapulumiapi", version: "1.0" },
                paths: {
                    "/{proxy+}": {
                        "x-amazon-apigateway-any-method": {
                            "x-amazon-apigateway-integration": {
                                uri: arn && "arn:aws:apigateway:" + region +
                                        ":lambda:path/2015-03-31/functions/" + arn + "/invocations",
                                passthroughBehavior: "when_no_match",
                                httpMethod: "POST",
                                type: "aws_proxy",
                            },
                        },
                    },
                },
            })
        );

    return swaggerSpec;
}

// create the API Gateway REST API, using a swagger spec
let restApi = new aws.apigateway.RestApi("api", {
    body: swaggerSpecForProxy(lambda),
});

// create a deployment
let deployment = new aws.apigateway.Deployment("serverless-deployment-dev", {
    restApi: restApi,
    description: "my deployment",
    stageName: "", // must be specified, but *must* be blank
});

// create a stage
let stage = new aws.apigateway.Stage("serverless-stage-dev", { 
    restApi: restApi,
    deployment: deployment,
    stageName: stageName
});

// give permissions from API Gateway to the Lambda
let invokePermission = new aws.lambda.Permission("serverlessLambdaPermission", {
    action: "lambda:invokeFunction",
    function: lambda,
    principal: "apigateway.amazonaws.com",
    sourceArn: deployment.executionArn.apply(arn => arn + "*/*"), // TODO: fix hardcoded path
});

// export the final API URL
export let endpoint = deployment.invokeUrl.apply(url => url + stageName);
