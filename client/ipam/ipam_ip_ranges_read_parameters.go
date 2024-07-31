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

// NewIpamIPRangesReadParams creates a new IpamIPRangesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPRangesReadParams() *IpamIPRangesReadParams {
	return &IpamIPRangesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPRangesReadParamsWithTimeout creates a new IpamIPRangesReadParams object
// with the ability to set a timeout on a request.
func NewIpamIPRangesReadParamsWithTimeout(timeout time.Duration) *IpamIPRangesReadParams {
	return &IpamIPRangesReadParams{
		timeout: timeout,
	}
}

// NewIpamIPRangesReadParamsWithContext creates a new IpamIPRangesReadParams object
// with the ability to set a context for a request.
func NewIpamIPRangesReadParamsWithContext(ctx context.Context) *IpamIPRangesReadParams {
	return &IpamIPRangesReadParams{
		Context: ctx,
	}
}

// NewIpamIPRangesReadParamsWithHTTPClient creates a new IpamIPRangesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPRangesReadParamsWithHTTPClient(client *http.Client) *IpamIPRangesReadParams {
	return &IpamIPRangesReadParams{
		HTTPClient: client,
	}
}

/*
IpamIPRangesReadParams contains all the parameters to send to the API endpoint

	for the ipam ip ranges read operation.

	Typically these are written to a http.Request.
*/
type IpamIPRangesReadParams struct {

	/* ID.

	   A unique integer value identifying this IP range.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip ranges read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesReadParams) WithDefaults() *IpamIPRangesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip ranges read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPRangesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) WithTimeout(timeout time.Duration) *IpamIPRangesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) WithContext(ctx context.Context) *IpamIPRangesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) WithHTTPClient(client *http.Client) *IpamIPRangesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) WithID(id int64) *IpamIPRangesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip ranges read params
func (o *IpamIPRangesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPRangesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
