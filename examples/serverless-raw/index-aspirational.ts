// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

// An example of what the API could look like
// See comments marked with "mock"

import * as aws from "@pulumi/aws";
import { asset } from "@pulumi/pulumi";
import { Output } from "@pulumi/pulumi";
import { AWSLambdaFullAccess } from "@pulumi/aws/iam";

let region = aws.config.requireRegion();
let stageName = "dev";

// DynamoDB Table
let usersTable = new aws.dynamodb.Table("UsersDynamoTable", {
    attribute: [
        { name: "userId", type: "S" },
    ],
    hashKey: "userId",
    readCapacity: 1,
    writeCapacity: 1,
});

// Lambda Function
let lambda = new aws.lambda.Function("mylambda", {
    role: AWSLambdaFullAccess, // Mock of what API could be
    handler: "index.handler",
    runtime: aws.lambda.NodeJS6d10Runtime,
    code: new asset.AssetArchive({ // TODO: this can probably be simplified by just referencing the ".app" folder
        "index.js": new asset.FileAsset("./app/index.js"),
        "node_modules": new asset.FileArchive("./app/node_modules")
    }),
    environment:  [ { 
        variables: {
            "USERS_TABLE": usersTable.name
        }
    } ]
});

// APIGateway RestAPI

// Mocked out API of what it could look like
// This wrapper would set up the routes and add invoke permission from API Gateway to the Lambda
let restApi = new RestApiWrapper("my-api", lambda, stageName,
    routes: [ 
        { "ANY", "/" }
        { "ANY", "{proxy+}" }
    ]
);

// export the final API URL
export let endpoint = restApi.invokeUrl.apply(url => url + stageName); // ideally, "stageName" would already be included
