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

	"github.com/fbreckle/go-netbox/models"
)

// NewIpamRirsBulkPartialUpdateParams creates a new IpamRirsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsBulkPartialUpdateParams() *IpamRirsBulkPartialUpdateParams {
	return &IpamRirsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsBulkPartialUpdateParamsWithTimeout creates a new IpamRirsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRirsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRirsBulkPartialUpdateParams {
	return &IpamRirsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRirsBulkPartialUpdateParamsWithContext creates a new IpamRirsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamRirsBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamRirsBulkPartialUpdateParams {
	return &IpamRirsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamRirsBulkPartialUpdateParamsWithHTTPClient creates a new IpamRirsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRirsBulkPartialUpdateParams {
	return &IpamRirsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamRirsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam rirs bulk partial update operation.

	Typically these are written to a http.Request.
*/
type IpamRirsBulkPartialUpdateParams struct {

	// Data.
	Data *models.RIR

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsBulkPartialUpdateParams) WithDefaults() *IpamRirsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRirsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamRirsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRirsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) WithData(data *models.RIR) *IpamRirsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam rirs bulk partial update params
func (o *IpamRirsBulkPartialUpdateParams) SetData(data *models.RIR) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
