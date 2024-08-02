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

// NewDcimVirtualChassisBulkDeleteParams creates a new DcimVirtualChassisBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisBulkDeleteParams() *DcimVirtualChassisBulkDeleteParams {
	return &DcimVirtualChassisBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisBulkDeleteParamsWithTimeout creates a new DcimVirtualChassisBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisBulkDeleteParams {
	return &DcimVirtualChassisBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisBulkDeleteParamsWithContext creates a new DcimVirtualChassisBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisBulkDeleteParamsWithContext(ctx context.Context) *DcimVirtualChassisBulkDeleteParams {
	return &DcimVirtualChassisBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisBulkDeleteParamsWithHTTPClient creates a new DcimVirtualChassisBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisBulkDeleteParams {
	return &DcimVirtualChassisBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualChassisBulkDeleteParams contains all the parameters to send to the API endpoint

	for the dcim virtual chassis bulk delete operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualChassisBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisBulkDeleteParams) WithDefaults() *DcimVirtualChassisBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) WithContext(ctx context.Context) *DcimVirtualChassisBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis bulk delete params
func (o *DcimVirtualChassisBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}