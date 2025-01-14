// Code generated by go-swagger; DO NOT EDIT.

package postgresql_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new postgresql clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for postgresql clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePostgresqlCluster(params *CreatePostgresqlClusterParams, opts ...ClientOption) (*CreatePostgresqlClusterOK, error)

	GetPostgresqlClusterCredentials(params *GetPostgresqlClusterCredentialsParams, opts ...ClientOption) (*GetPostgresqlClusterCredentialsOK, error)

	GetPostgresqlClusterResources(params *GetPostgresqlClusterResourcesParams, opts ...ClientOption) (*GetPostgresqlClusterResourcesOK, error)

	UpdatePostgresqlCluster(params *UpdatePostgresqlClusterParams, opts ...ClientOption) (*UpdatePostgresqlClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePostgresqlCluster creates postgresql cluster creates a new postgresql cluster
*/
func (a *Client) CreatePostgresqlCluster(params *CreatePostgresqlClusterParams, opts ...ClientOption) (*CreatePostgresqlClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostgresqlClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreatePostgresqlCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/PostgresqlCluster/Create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePostgresqlClusterReader{formats: a.formats},
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
	success, ok := result.(*CreatePostgresqlClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreatePostgresqlClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPostgresqlClusterCredentials gets postgresql cluster credentials returns a postgresql cluster credentials by cluster name
*/
func (a *Client) GetPostgresqlClusterCredentials(params *GetPostgresqlClusterCredentialsParams, opts ...ClientOption) (*GetPostgresqlClusterCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresqlClusterCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPostgresqlClusterCredentials",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/PostgresqlCluster/GetCredentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPostgresqlClusterCredentialsReader{formats: a.formats},
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
	success, ok := result.(*GetPostgresqlClusterCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresqlClusterCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetPostgresqlClusterResources gets postgresql cluster resources returns expected resources to be consumed by the cluster
*/
func (a *Client) GetPostgresqlClusterResources(params *GetPostgresqlClusterResourcesParams, opts ...ClientOption) (*GetPostgresqlClusterResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostgresqlClusterResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPostgresqlClusterResources",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/PostgresqlCluster/Resources/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPostgresqlClusterResourcesReader{formats: a.formats},
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
	success, ok := result.(*GetPostgresqlClusterResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPostgresqlClusterResourcesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdatePostgresqlCluster updates postgresql cluster updates existing postgresql cluster
*/
func (a *Client) UpdatePostgresqlCluster(params *UpdatePostgresqlClusterParams, opts ...ClientOption) (*UpdatePostgresqlClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePostgresqlClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdatePostgresqlCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/PostgresqlCluster/Update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePostgresqlClusterReader{formats: a.formats},
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
	success, ok := result.(*UpdatePostgresqlClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePostgresqlClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
