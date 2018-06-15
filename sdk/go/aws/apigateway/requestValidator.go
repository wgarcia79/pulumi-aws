// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type RequestValidator struct {
	s *pulumi.ResourceState
}

// NewRequestValidator registers a new resource with the given unique name, arguments, and options.
func NewRequestValidator(ctx *pulumi.Context,
	name string, args *RequestValidatorArgs, opts ...pulumi.ResourceOpt) (*RequestValidator, error) {
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["restApi"] = nil
		inputs["validateRequestBody"] = nil
		inputs["validateRequestParameters"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["restApi"] = args.RestApi
		inputs["validateRequestBody"] = args.ValidateRequestBody
		inputs["validateRequestParameters"] = args.ValidateRequestParameters
	}
	s, err := ctx.RegisterResource("aws:apigateway/requestValidator:RequestValidator", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RequestValidator{s: s}, nil
}

// GetRequestValidator gets an existing RequestValidator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRequestValidator(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RequestValidatorState, opts ...pulumi.ResourceOpt) (*RequestValidator, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["restApi"] = state.RestApi
		inputs["validateRequestBody"] = state.ValidateRequestBody
		inputs["validateRequestParameters"] = state.ValidateRequestParameters
	}
	s, err := ctx.ReadResource("aws:apigateway/requestValidator:RequestValidator", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RequestValidator{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RequestValidator) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RequestValidator) ID() *pulumi.IDOutput {
	return r.s.ID
}

func (r *RequestValidator) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *RequestValidator) RestApi() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["restApi"])
}

func (r *RequestValidator) ValidateRequestBody() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["validateRequestBody"])
}

func (r *RequestValidator) ValidateRequestParameters() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["validateRequestParameters"])
}

// Input properties used for looking up and filtering RequestValidator resources.
type RequestValidatorState struct {
	Name interface{}
	RestApi interface{}
	ValidateRequestBody interface{}
	ValidateRequestParameters interface{}
}

// The set of arguments for constructing a RequestValidator resource.
type RequestValidatorArgs struct {
	Name interface{}
	RestApi interface{}
	ValidateRequestBody interface{}
	ValidateRequestParameters interface{}
}