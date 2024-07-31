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

	"github.com/fbreckle/go-netbox/models"
)

// NewDcimModulesCreateParams creates a new DcimModulesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesCreateParams() *DcimModulesCreateParams {
	return &DcimModulesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesCreateParamsWithTimeout creates a new DcimModulesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimModulesCreateParamsWithTimeout(timeout time.Duration) *DcimModulesCreateParams {
	return &DcimModulesCreateParams{
		timeout: timeout,
	}
}

// NewDcimModulesCreateParamsWithContext creates a new DcimModulesCreateParams object
// with the ability to set a context for a request.
func NewDcimModulesCreateParamsWithContext(ctx context.Context) *DcimModulesCreateParams {
	return &DcimModulesCreateParams{
		Context: ctx,
	}
}

// NewDcimModulesCreateParamsWithHTTPClient creates a new DcimModulesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesCreateParamsWithHTTPClient(client *http.Client) *DcimModulesCreateParams {
	return &DcimModulesCreateParams{
		HTTPClient: client,
	}
}

/*
DcimModulesCreateParams contains all the parameters to send to the API endpoint

	for the dcim modules create operation.

	Typically these are written to a http.Request.
*/
type DcimModulesCreateParams struct {

	// Data.
	Data *models.WritableModule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesCreateParams) WithDefaults() *DcimModulesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules create params
func (o *DcimModulesCreateParams) WithTimeout(timeout time.Duration) *DcimModulesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules create params
func (o *DcimModulesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules create params
func (o *DcimModulesCreateParams) WithContext(ctx context.Context) *DcimModulesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules create params
func (o *DcimModulesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules create params
func (o *DcimModulesCreateParams) WithHTTPClient(client *http.Client) *DcimModulesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules create params
func (o *DcimModulesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim modules create params
func (o *DcimModulesCreateParams) WithData(data *models.WritableModule) *DcimModulesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim modules create params
func (o *DcimModulesCreateParams) SetData(data *models.WritableModule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
