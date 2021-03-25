// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:opsworks/application:Application":
		r, err = NewApplication(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/customLayer:CustomLayer":
		r, err = NewCustomLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/gangliaLayer:GangliaLayer":
		r, err = NewGangliaLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/haproxyLayer:HaproxyLayer":
		r, err = NewHaproxyLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/instance:Instance":
		r, err = NewInstance(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/javaAppLayer:JavaAppLayer":
		r, err = NewJavaAppLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/memcachedLayer:MemcachedLayer":
		r, err = NewMemcachedLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/mysqlLayer:MysqlLayer":
		r, err = NewMysqlLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/nodejsAppLayer:NodejsAppLayer":
		r, err = NewNodejsAppLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/permission:Permission":
		r, err = NewPermission(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/phpAppLayer:PhpAppLayer":
		r, err = NewPhpAppLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/railsAppLayer:RailsAppLayer":
		r, err = NewRailsAppLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/rdsDbInstance:RdsDbInstance":
		r, err = NewRdsDbInstance(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/stack:Stack":
		r, err = NewStack(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/staticWebLayer:StaticWebLayer":
		r, err = NewStaticWebLayer(ctx, name, nil, pulumi.URN_(urn))
	case "aws:opsworks/userProfile:UserProfile":
		r, err = NewUserProfile(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/customLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/gangliaLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/haproxyLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/javaAppLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/memcachedLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/mysqlLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/nodejsAppLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/permission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/phpAppLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/railsAppLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/rdsDbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/stack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/staticWebLayer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"opsworks/userProfile",
		&module{version},
	)
}
