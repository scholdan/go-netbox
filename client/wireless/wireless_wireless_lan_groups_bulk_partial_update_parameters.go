// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// NewWirelessWirelessLanGroupsBulkPartialUpdateParams creates a new WirelessWirelessLanGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLanGroupsBulkPartialUpdateParams() *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	return &WirelessWirelessLanGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithTimeout creates a new WirelessWirelessLanGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	return &WirelessWirelessLanGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithContext creates a new WirelessWirelessLanGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	return &WirelessWirelessLanGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithHTTPClient creates a new WirelessWirelessLanGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLanGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	return &WirelessWirelessLanGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLanGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the wireless wireless lan groups bulk partial update operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLanGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableWirelessLANGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless lan groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WithDefaults() *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless lan groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WithData(data *models.WritableWirelessLANGroup) *WirelessWirelessLanGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the wireless wireless lan groups bulk partial update params
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) SetData(data *models.WritableWirelessLANGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLanGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
