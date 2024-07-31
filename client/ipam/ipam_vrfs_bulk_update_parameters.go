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

	"github.com/scholdan/go-netbox/models"
)

// NewIpamVrfsBulkUpdateParams creates a new IpamVrfsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsBulkUpdateParams() *IpamVrfsBulkUpdateParams {
	return &IpamVrfsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsBulkUpdateParamsWithTimeout creates a new IpamVrfsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamVrfsBulkUpdateParams {
	return &IpamVrfsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVrfsBulkUpdateParamsWithContext creates a new IpamVrfsBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamVrfsBulkUpdateParamsWithContext(ctx context.Context) *IpamVrfsBulkUpdateParams {
	return &IpamVrfsBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamVrfsBulkUpdateParamsWithHTTPClient creates a new IpamVrfsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamVrfsBulkUpdateParams {
	return &IpamVrfsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamVrfsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam vrfs bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamVrfsBulkUpdateParams struct {

	// Data.
	Data *models.WritableVRF

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkUpdateParams) WithDefaults() *IpamVrfsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamVrfsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) WithContext(ctx context.Context) *IpamVrfsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamVrfsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) WithData(data *models.WritableVRF) *IpamVrfsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs bulk update params
func (o *IpamVrfsBulkUpdateParams) SetData(data *models.WritableVRF) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
