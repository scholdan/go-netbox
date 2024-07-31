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
	"github.com/go-openapi/swag"
)

// NewDcimPowerFeedsDeleteParams creates a new DcimPowerFeedsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsDeleteParams() *DcimPowerFeedsDeleteParams {
	return &DcimPowerFeedsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsDeleteParamsWithTimeout creates a new DcimPowerFeedsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsDeleteParams {
	return &DcimPowerFeedsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsDeleteParamsWithContext creates a new DcimPowerFeedsDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsDeleteParamsWithContext(ctx context.Context) *DcimPowerFeedsDeleteParams {
	return &DcimPowerFeedsDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsDeleteParamsWithHTTPClient creates a new DcimPowerFeedsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsDeleteParams {
	return &DcimPowerFeedsDeleteParams{
		HTTPClient: client,
	}
}

/*
DcimPowerFeedsDeleteParams contains all the parameters to send to the API endpoint

	for the dcim power feeds delete operation.

	Typically these are written to a http.Request.
*/
type DcimPowerFeedsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this power feed.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsDeleteParams) WithDefaults() *DcimPowerFeedsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) WithContext(ctx context.Context) *DcimPowerFeedsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) WithID(id int64) *DcimPowerFeedsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds delete params
func (o *DcimPowerFeedsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
