// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListGatewaysParams creates a new ListGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGatewaysParams() *ListGatewaysParams {
	return &ListGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGatewaysParamsWithTimeout creates a new ListGatewaysParams object
// with the ability to set a timeout on a request.
func NewListGatewaysParamsWithTimeout(timeout time.Duration) *ListGatewaysParams {
	return &ListGatewaysParams{
		timeout: timeout,
	}
}

// NewListGatewaysParamsWithContext creates a new ListGatewaysParams object
// with the ability to set a context for a request.
func NewListGatewaysParamsWithContext(ctx context.Context) *ListGatewaysParams {
	return &ListGatewaysParams{
		Context: ctx,
	}
}

// NewListGatewaysParamsWithHTTPClient creates a new ListGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGatewaysParamsWithHTTPClient(client *http.Client) *ListGatewaysParams {
	return &ListGatewaysParams{
		HTTPClient: client,
	}
}

/*
ListGatewaysParams contains all the parameters to send to the API endpoint

	for the list gateways operation.

	Typically these are written to a http.Request.
*/
type ListGatewaysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGatewaysParams) WithDefaults() *ListGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGatewaysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list gateways params
func (o *ListGatewaysParams) WithTimeout(timeout time.Duration) *ListGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list gateways params
func (o *ListGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list gateways params
func (o *ListGatewaysParams) WithContext(ctx context.Context) *ListGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list gateways params
func (o *ListGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list gateways params
func (o *ListGatewaysParams) WithHTTPClient(client *http.Client) *ListGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list gateways params
func (o *ListGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}