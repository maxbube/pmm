// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddExternalExporterParams creates a new AddExternalExporterParams object
// with the default values initialized.
func NewAddExternalExporterParams() *AddExternalExporterParams {
	var ()
	return &AddExternalExporterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddExternalExporterParamsWithTimeout creates a new AddExternalExporterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddExternalExporterParamsWithTimeout(timeout time.Duration) *AddExternalExporterParams {
	var ()
	return &AddExternalExporterParams{

		timeout: timeout,
	}
}

// NewAddExternalExporterParamsWithContext creates a new AddExternalExporterParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddExternalExporterParamsWithContext(ctx context.Context) *AddExternalExporterParams {
	var ()
	return &AddExternalExporterParams{

		Context: ctx,
	}
}

// NewAddExternalExporterParamsWithHTTPClient creates a new AddExternalExporterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddExternalExporterParamsWithHTTPClient(client *http.Client) *AddExternalExporterParams {
	var ()
	return &AddExternalExporterParams{
		HTTPClient: client,
	}
}

/*AddExternalExporterParams contains all the parameters to send to the API endpoint
for the add external exporter operation typically these are written to a http.Request
*/
type AddExternalExporterParams struct {

	/*Body*/
	Body AddExternalExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add external exporter params
func (o *AddExternalExporterParams) WithTimeout(timeout time.Duration) *AddExternalExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add external exporter params
func (o *AddExternalExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add external exporter params
func (o *AddExternalExporterParams) WithContext(ctx context.Context) *AddExternalExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add external exporter params
func (o *AddExternalExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add external exporter params
func (o *AddExternalExporterParams) WithHTTPClient(client *http.Client) *AddExternalExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add external exporter params
func (o *AddExternalExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add external exporter params
func (o *AddExternalExporterParams) WithBody(body AddExternalExporterBody) *AddExternalExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add external exporter params
func (o *AddExternalExporterParams) SetBody(body AddExternalExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddExternalExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}