// Code generated by go-swagger; DO NOT EDIT.

package object_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new object details API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for object details API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLabels gets labels gets list of labels for object details
*/
func (a *Client) GetLabels(params *GetLabelsParams) (*GetLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLabels",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetLabels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetMetrics gets metrics gets map of metrics for specific filtering
*/
func (a *Client) GetMetrics(params *GetMetricsParams) (*GetMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetricsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMetrics",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetMetrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMetricsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMetricsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetQueryExample gets query example gets list of query examples
*/
func (a *Client) GetQueryExample(params *GetQueryExampleParams) (*GetQueryExampleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueryExampleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetQueryExample",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetQueryExample",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQueryExampleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetQueryExampleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetQueryExampleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
