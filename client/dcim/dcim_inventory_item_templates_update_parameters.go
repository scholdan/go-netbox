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

// NewDcimInventoryItemTemplatesUpdateParams creates a new DcimInventoryItemTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemTemplatesUpdateParams() *DcimInventoryItemTemplatesUpdateParams {
	return &DcimInventoryItemTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemTemplatesUpdateParamsWithTimeout creates a new DcimInventoryItemTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesUpdateParams {
	return &DcimInventoryItemTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemTemplatesUpdateParamsWithContext creates a new DcimInventoryItemTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemTemplatesUpdateParamsWithContext(ctx context.Context) *DcimInventoryItemTemplatesUpdateParams {
	return &DcimInventoryItemTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemTemplatesUpdateParamsWithHTTPClient creates a new DcimInventoryItemTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesUpdateParams {
	return &DcimInventoryItemTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimInventoryItemTemplatesUpdateParams contains all the parameters to send to the API endpoint

	for the dcim inventory item templates update operation.

	Typically these are written to a http.Request.
*/
type DcimInventoryItemTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableInventoryItemTemplate

	/* ID.

	   A unique integer value identifying this inventory item template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory item templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesUpdateParams) WithDefaults() *DcimInventoryItemTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory item templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) WithContext(ctx context.Context) *DcimInventoryItemTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) WithData(data *models.WritableInventoryItemTemplate) *DcimInventoryItemTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) SetData(data *models.WritableInventoryItemTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) WithID(id int64) *DcimInventoryItemTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory item templates update params
func (o *DcimInventoryItemTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
