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

// NewIpamFhrpGroupsBulkPartialUpdateParams creates a new IpamFhrpGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupsBulkPartialUpdateParams() *IpamFhrpGroupsBulkPartialUpdateParams {
	return &IpamFhrpGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupsBulkPartialUpdateParamsWithTimeout creates a new IpamFhrpGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupsBulkPartialUpdateParams {
	return &IpamFhrpGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupsBulkPartialUpdateParamsWithContext creates a new IpamFhrpGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamFhrpGroupsBulkPartialUpdateParams {
	return &IpamFhrpGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupsBulkPartialUpdateParamsWithHTTPClient creates a new IpamFhrpGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupsBulkPartialUpdateParams {
	return &IpamFhrpGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamFhrpGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam fhrp groups bulk partial update operation.

	Typically these are written to a http.Request.
*/
type IpamFhrpGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.FHRPGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WithDefaults() *IpamFhrpGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamFhrpGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WithData(data *models.FHRPGroup) *IpamFhrpGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam fhrp groups bulk partial update params
func (o *IpamFhrpGroupsBulkPartialUpdateParams) SetData(data *models.FHRPGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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