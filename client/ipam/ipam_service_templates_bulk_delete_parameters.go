// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIpamServiceTemplatesBulkDeleteParams creates a new IpamServiceTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServiceTemplatesBulkDeleteParams() *IpamServiceTemplatesBulkDeleteParams {
	return &IpamServiceTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServiceTemplatesBulkDeleteParamsWithTimeout creates a new IpamServiceTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamServiceTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamServiceTemplatesBulkDeleteParams {
	return &IpamServiceTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamServiceTemplatesBulkDeleteParamsWithContext creates a new IpamServiceTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamServiceTemplatesBulkDeleteParamsWithContext(ctx context.Context) *IpamServiceTemplatesBulkDeleteParams {
	return &IpamServiceTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamServiceTemplatesBulkDeleteParamsWithHTTPClient creates a new IpamServiceTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServiceTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamServiceTemplatesBulkDeleteParams {
	return &IpamServiceTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
IpamServiceTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint

	for the ipam service templates bulk delete operation.

	Typically these are written to a http.Request.
*/
type IpamServiceTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam service templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesBulkDeleteParams) WithDefaults() *IpamServiceTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam service templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamServiceTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) WithContext(ctx context.Context) *IpamServiceTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamServiceTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam service templates bulk delete params
func (o *IpamServiceTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServiceTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
