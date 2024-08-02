// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGateway(params *CreateGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGatewayCreated, error)

	DeleteGateway(params *DeleteGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGatewayNoContent, error)

	GetGateway(params *GetGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGatewayOK, error)

	ListGateways(params *ListGatewaysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGatewaysOK, error)

	UpdateGateway(params *UpdateGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGatewayOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGateway creates a new gateway
*/
func (a *Client) CreateGateway(params *CreateGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGatewayCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createGateway",
		Method:             "POST",
		PathPattern:        "/plugins/nb-gateways/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGatewayCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteGateway deletes a specific gateway
*/
func (a *Client) DeleteGateway(params *DeleteGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGatewayNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteGateway",
		Method:             "DELETE",
		PathPattern:        "/gateways/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGatewayNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGateway retrieves a specific gateway
*/
func (a *Client) GetGateway(params *GetGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGateway",
		Method:             "GET",
		PathPattern:        "/gateways/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListGateways lists all gateways
*/
func (a *Client) ListGateways(params *ListGatewaysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGatewaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGateways",
		Method:             "GET",
		PathPattern:        "/plugins/nb-gateways/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listGateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateGateway updates a specific gateway
*/
func (a *Client) UpdateGateway(params *UpdateGatewayParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateGateway",
		Method:             "PUT",
		PathPattern:        "/gateways/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
