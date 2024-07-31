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

// NewDcimDeviceRolesBulkUpdateParams creates a new DcimDeviceRolesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceRolesBulkUpdateParams() *DcimDeviceRolesBulkUpdateParams {
	return &DcimDeviceRolesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesBulkUpdateParamsWithTimeout creates a new DcimDeviceRolesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceRolesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesBulkUpdateParams {
	return &DcimDeviceRolesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceRolesBulkUpdateParamsWithContext creates a new DcimDeviceRolesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceRolesBulkUpdateParamsWithContext(ctx context.Context) *DcimDeviceRolesBulkUpdateParams {
	return &DcimDeviceRolesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceRolesBulkUpdateParamsWithHTTPClient creates a new DcimDeviceRolesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceRolesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesBulkUpdateParams {
	return &DcimDeviceRolesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimDeviceRolesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim device roles bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimDeviceRolesBulkUpdateParams struct {

	// Data.
	Data *models.DeviceRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesBulkUpdateParams) WithDefaults() *DcimDeviceRolesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) WithContext(ctx context.Context) *DcimDeviceRolesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) WithData(data *models.DeviceRole) *DcimDeviceRolesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device roles bulk update params
func (o *DcimDeviceRolesBulkUpdateParams) SetData(data *models.DeviceRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
