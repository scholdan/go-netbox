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

// NewIpamAsnsCreateParams creates a new IpamAsnsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAsnsCreateParams() *IpamAsnsCreateParams {
	return &IpamAsnsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAsnsCreateParamsWithTimeout creates a new IpamAsnsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamAsnsCreateParamsWithTimeout(timeout time.Duration) *IpamAsnsCreateParams {
	return &IpamAsnsCreateParams{
		timeout: timeout,
	}
}

// NewIpamAsnsCreateParamsWithContext creates a new IpamAsnsCreateParams object
// with the ability to set a context for a request.
func NewIpamAsnsCreateParamsWithContext(ctx context.Context) *IpamAsnsCreateParams {
	return &IpamAsnsCreateParams{
		Context: ctx,
	}
}

// NewIpamAsnsCreateParamsWithHTTPClient creates a new IpamAsnsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAsnsCreateParamsWithHTTPClient(client *http.Client) *IpamAsnsCreateParams {
	return &IpamAsnsCreateParams{
		HTTPClient: client,
	}
}

/*
IpamAsnsCreateParams contains all the parameters to send to the API endpoint

	for the ipam asns create operation.

	Typically these are written to a http.Request.
*/
type IpamAsnsCreateParams struct {

	// Data.
	Data *models.WritableASN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam asns create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsCreateParams) WithDefaults() *IpamAsnsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam asns create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam asns create params
func (o *IpamAsnsCreateParams) WithTimeout(timeout time.Duration) *IpamAsnsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam asns create params
func (o *IpamAsnsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam asns create params
func (o *IpamAsnsCreateParams) WithContext(ctx context.Context) *IpamAsnsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam asns create params
func (o *IpamAsnsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam asns create params
func (o *IpamAsnsCreateParams) WithHTTPClient(client *http.Client) *IpamAsnsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam asns create params
func (o *IpamAsnsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam asns create params
func (o *IpamAsnsCreateParams) WithData(data *models.WritableASN) *IpamAsnsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam asns create params
func (o *IpamAsnsCreateParams) SetData(data *models.WritableASN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAsnsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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