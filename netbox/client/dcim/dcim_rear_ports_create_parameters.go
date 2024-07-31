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

	"github.com/scholdan/go-netbox/models"
)

// NewDcimRearPortsCreateParams creates a new DcimRearPortsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsCreateParams() *DcimRearPortsCreateParams {
	return &DcimRearPortsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsCreateParamsWithTimeout creates a new DcimRearPortsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsCreateParamsWithTimeout(timeout time.Duration) *DcimRearPortsCreateParams {
	return &DcimRearPortsCreateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsCreateParamsWithContext creates a new DcimRearPortsCreateParams object
// with the ability to set a context for a request.
func NewDcimRearPortsCreateParamsWithContext(ctx context.Context) *DcimRearPortsCreateParams {
	return &DcimRearPortsCreateParams{
		Context: ctx,
	}
}

// NewDcimRearPortsCreateParamsWithHTTPClient creates a new DcimRearPortsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsCreateParamsWithHTTPClient(client *http.Client) *DcimRearPortsCreateParams {
	return &DcimRearPortsCreateParams{
		HTTPClient: client,
	}
}

/*
DcimRearPortsCreateParams contains all the parameters to send to the API endpoint

	for the dcim rear ports create operation.

	Typically these are written to a http.Request.
*/
type DcimRearPortsCreateParams struct {

	// Data.
	Data *models.WritableRearPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsCreateParams) WithDefaults() *DcimRearPortsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) WithTimeout(timeout time.Duration) *DcimRearPortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) WithContext(ctx context.Context) *DcimRearPortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) WithHTTPClient(client *http.Client) *DcimRearPortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) WithData(data *models.WritableRearPort) *DcimRearPortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear ports create params
func (o *DcimRearPortsCreateParams) SetData(data *models.WritableRearPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
