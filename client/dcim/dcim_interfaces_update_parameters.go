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

	"github.com/scholdan/go-netbox/models"
)

// NewDcimInterfacesUpdateParams creates a new DcimInterfacesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfacesUpdateParams() *DcimInterfacesUpdateParams {
	return &DcimInterfacesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesUpdateParamsWithTimeout creates a new DcimInterfacesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInterfacesUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfacesUpdateParams {
	return &DcimInterfacesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInterfacesUpdateParamsWithContext creates a new DcimInterfacesUpdateParams object
// with the ability to set a context for a request.
func NewDcimInterfacesUpdateParamsWithContext(ctx context.Context) *DcimInterfacesUpdateParams {
	return &DcimInterfacesUpdateParams{
		Context: ctx,
	}
}

// NewDcimInterfacesUpdateParamsWithHTTPClient creates a new DcimInterfacesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfacesUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfacesUpdateParams {
	return &DcimInterfacesUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimInterfacesUpdateParams contains all the parameters to send to the API endpoint

	for the dcim interfaces update operation.

	Typically these are written to a http.Request.
*/
type DcimInterfacesUpdateParams struct {

	// Data.
	Data *models.WritableInterface

	/* ID.

	   A unique integer value identifying this interface.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interfaces update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesUpdateParams) WithDefaults() *DcimInterfacesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interfaces update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfacesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) WithContext(ctx context.Context) *DcimInterfacesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfacesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) WithData(data *models.WritableInterface) *DcimInterfacesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) SetData(data *models.WritableInterface) {
	o.Data = data
}

// WithID adds the id to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) WithID(id int64) *DcimInterfacesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces update params
func (o *DcimInterfacesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
