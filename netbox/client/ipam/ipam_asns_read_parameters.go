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
)

// NewIpamAsnsReadParams creates a new IpamAsnsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAsnsReadParams() *IpamAsnsReadParams {
	return &IpamAsnsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAsnsReadParamsWithTimeout creates a new IpamAsnsReadParams object
// with the ability to set a timeout on a request.
func NewIpamAsnsReadParamsWithTimeout(timeout time.Duration) *IpamAsnsReadParams {
	return &IpamAsnsReadParams{
		timeout: timeout,
	}
}

// NewIpamAsnsReadParamsWithContext creates a new IpamAsnsReadParams object
// with the ability to set a context for a request.
func NewIpamAsnsReadParamsWithContext(ctx context.Context) *IpamAsnsReadParams {
	return &IpamAsnsReadParams{
		Context: ctx,
	}
}

// NewIpamAsnsReadParamsWithHTTPClient creates a new IpamAsnsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAsnsReadParamsWithHTTPClient(client *http.Client) *IpamAsnsReadParams {
	return &IpamAsnsReadParams{
		HTTPClient: client,
	}
}

/*
IpamAsnsReadParams contains all the parameters to send to the API endpoint

	for the ipam asns read operation.

	Typically these are written to a http.Request.
*/
type IpamAsnsReadParams struct {

	/* ID.

	   A unique integer value identifying this ASN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam asns read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsReadParams) WithDefaults() *IpamAsnsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam asns read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam asns read params
func (o *IpamAsnsReadParams) WithTimeout(timeout time.Duration) *IpamAsnsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam asns read params
func (o *IpamAsnsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam asns read params
func (o *IpamAsnsReadParams) WithContext(ctx context.Context) *IpamAsnsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam asns read params
func (o *IpamAsnsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam asns read params
func (o *IpamAsnsReadParams) WithHTTPClient(client *http.Client) *IpamAsnsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam asns read params
func (o *IpamAsnsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam asns read params
func (o *IpamAsnsReadParams) WithID(id int64) *IpamAsnsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam asns read params
func (o *IpamAsnsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAsnsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
