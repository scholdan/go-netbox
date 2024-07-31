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

// NewDcimDevicesBulkUpdateParams creates a new DcimDevicesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesBulkUpdateParams() *DcimDevicesBulkUpdateParams {
	return &DcimDevicesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesBulkUpdateParamsWithTimeout creates a new DcimDevicesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimDevicesBulkUpdateParams {
	return &DcimDevicesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDevicesBulkUpdateParamsWithContext creates a new DcimDevicesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimDevicesBulkUpdateParamsWithContext(ctx context.Context) *DcimDevicesBulkUpdateParams {
	return &DcimDevicesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimDevicesBulkUpdateParamsWithHTTPClient creates a new DcimDevicesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimDevicesBulkUpdateParams {
	return &DcimDevicesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimDevicesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim devices bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimDevicesBulkUpdateParams struct {

	// Data.
	Data *models.WritableDeviceWithConfigContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesBulkUpdateParams) WithDefaults() *DcimDevicesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimDevicesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) WithContext(ctx context.Context) *DcimDevicesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimDevicesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) WithData(data *models.WritableDeviceWithConfigContext) *DcimDevicesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim devices bulk update params
func (o *DcimDevicesBulkUpdateParams) SetData(data *models.WritableDeviceWithConfigContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
