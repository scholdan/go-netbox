// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// NewExtrasConfigTemplatesBulkDeleteParams creates a new ExtrasConfigTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigTemplatesBulkDeleteParams() *ExtrasConfigTemplatesBulkDeleteParams {
	return &ExtrasConfigTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigTemplatesBulkDeleteParamsWithTimeout creates a new ExtrasConfigTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasConfigTemplatesBulkDeleteParams {
	return &ExtrasConfigTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasConfigTemplatesBulkDeleteParamsWithContext creates a new ExtrasConfigTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasConfigTemplatesBulkDeleteParamsWithContext(ctx context.Context) *ExtrasConfigTemplatesBulkDeleteParams {
	return &ExtrasConfigTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasConfigTemplatesBulkDeleteParamsWithHTTPClient creates a new ExtrasConfigTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasConfigTemplatesBulkDeleteParams {
	return &ExtrasConfigTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
ExtrasConfigTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint

	for the extras config templates bulk delete operation.

	Typically these are written to a http.Request.
*/
type ExtrasConfigTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigTemplatesBulkDeleteParams) WithDefaults() *ExtrasConfigTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasConfigTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) WithContext(ctx context.Context) *ExtrasConfigTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasConfigTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config templates bulk delete params
func (o *ExtrasConfigTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
