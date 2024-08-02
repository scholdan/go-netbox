// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// NewDcimManufacturersBulkDeleteParams creates a new DcimManufacturersBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersBulkDeleteParams() *DcimManufacturersBulkDeleteParams {
	return &DcimManufacturersBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersBulkDeleteParamsWithTimeout creates a new DcimManufacturersBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimManufacturersBulkDeleteParams {
	return &DcimManufacturersBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersBulkDeleteParamsWithContext creates a new DcimManufacturersBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimManufacturersBulkDeleteParamsWithContext(ctx context.Context) *DcimManufacturersBulkDeleteParams {
	return &DcimManufacturersBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimManufacturersBulkDeleteParamsWithHTTPClient creates a new DcimManufacturersBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimManufacturersBulkDeleteParams {
	return &DcimManufacturersBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimManufacturersBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim manufacturers bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimManufacturersBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersBulkDeleteParams) WithDefaults() *DcimManufacturersBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimManufacturersBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) WithContext(ctx context.Context) *DcimManufacturersBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimManufacturersBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers bulk delete params
func (o *DcimManufacturersBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
