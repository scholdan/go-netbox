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
	"github.com/go-openapi/swag"

	"github.com/scholdan/go-netbox/models"
)

// NewIpamIPRangesPartialUpdateParams creates a new IpamIPRangesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPRangesPartialUpdateParams() *IpamIPRangesPartialUpdateParams {
	return &IpamIPRangesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPRangesPartialUpdateParamsWithTimeout creates a new IpamIPRangesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamIPRangesPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamIPRangesPartialUpdateParams {
	return &IpamIPRangesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamIPRangesPartialUpdateParamsWithContext creates a new IpamIPRangesPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamIPRangesPartialUpdateParamsWithContext(ctx context.Context) *IpamIPRangesPartialUpdateParams {
	return &IpamIPRangesPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamIPRangesPartialUpdateParamsWithHTTPClient creates a new IpamIPRangesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPRangesPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamIPRangesPartialUpdateParams {
	return &IpamIPRangesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamIPRangesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam ip ranges partial update operation.

	Typically these are written to a http.Request.
*/
type IpamIPRangesPartialUpdateParams struct {

	// Data.
	Data *models.WritableIPRange

	/* ID.

	   A unique integer value identifying this IP range.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip ranges partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesPartialUpdateParams) WithDefaults() *IpamIPRangesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip ranges partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamIPRangesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) WithContext(ctx context.Context) *IpamIPRangesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamIPRangesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) WithData(data *models.WritableIPRange) *IpamIPRangesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) SetData(data *models.WritableIPRange) {
	o.Data = data
}

// WithID adds the id to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) WithID(id int64) *IpamIPRangesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip ranges partial update params
func (o *IpamIPRangesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPRangesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
