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

	"github.com/scholdan/go-netbox/models"
)

// NewCreateGatewayParams creates a new CreateGatewayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGatewayParams() *CreateGatewayParams {
	return &CreateGatewayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGatewayParamsWithTimeout creates a new CreateGatewayParams object
// with the ability to set a timeout on a request.
func NewCreateGatewayParamsWithTimeout(timeout time.Duration) *CreateGatewayParams {
	return &CreateGatewayParams{
		timeout: timeout,
	}
}

// NewCreateGatewayParamsWithContext creates a new CreateGatewayParams object
// with the ability to set a context for a request.
func NewCreateGatewayParamsWithContext(ctx context.Context) *CreateGatewayParams {
	return &CreateGatewayParams{
		Context: ctx,
	}
}

// NewCreateGatewayParamsWithHTTPClient creates a new CreateGatewayParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGatewayParamsWithHTTPClient(client *http.Client) *CreateGatewayParams {
	return &CreateGatewayParams{
		HTTPClient: client,
	}
}

/*
CreateGatewayParams contains all the parameters to send to the API endpoint

	for the create gateway operation.

	Typically these are written to a http.Request.
*/
type CreateGatewayParams struct {

	// Body.
	Body *models.Gateway

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGatewayParams) WithDefaults() *CreateGatewayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGatewayParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create gateway params
func (o *CreateGatewayParams) WithTimeout(timeout time.Duration) *CreateGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gateway params
func (o *CreateGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gateway params
func (o *CreateGatewayParams) WithContext(ctx context.Context) *CreateGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gateway params
func (o *CreateGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gateway params
func (o *CreateGatewayParams) WithHTTPClient(client *http.Client) *CreateGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gateway params
func (o *CreateGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create gateway params
func (o *CreateGatewayParams) WithBody(body *models.Gateway) *CreateGatewayParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gateway params
func (o *CreateGatewayParams) SetBody(body *models.Gateway) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
