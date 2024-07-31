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

// NewDcimInventoryItemRolesCreateParams creates a new DcimInventoryItemRolesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemRolesCreateParams() *DcimInventoryItemRolesCreateParams {
	return &DcimInventoryItemRolesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemRolesCreateParamsWithTimeout creates a new DcimInventoryItemRolesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemRolesCreateParamsWithTimeout(timeout time.Duration) *DcimInventoryItemRolesCreateParams {
	return &DcimInventoryItemRolesCreateParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemRolesCreateParamsWithContext creates a new DcimInventoryItemRolesCreateParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemRolesCreateParamsWithContext(ctx context.Context) *DcimInventoryItemRolesCreateParams {
	return &DcimInventoryItemRolesCreateParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemRolesCreateParamsWithHTTPClient creates a new DcimInventoryItemRolesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemRolesCreateParamsWithHTTPClient(client *http.Client) *DcimInventoryItemRolesCreateParams {
	return &DcimInventoryItemRolesCreateParams{
		HTTPClient: client,
	}
}

/*
DcimInventoryItemRolesCreateParams contains all the parameters to send to the API endpoint

	for the dcim inventory item roles create operation.

	Typically these are written to a http.Request.
*/
type DcimInventoryItemRolesCreateParams struct {

	// Data.
	Data *models.InventoryItemRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory item roles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemRolesCreateParams) WithDefaults() *DcimInventoryItemRolesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory item roles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemRolesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) WithTimeout(timeout time.Duration) *DcimInventoryItemRolesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) WithContext(ctx context.Context) *DcimInventoryItemRolesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) WithHTTPClient(client *http.Client) *DcimInventoryItemRolesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) WithData(data *models.InventoryItemRole) *DcimInventoryItemRolesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim inventory item roles create params
func (o *DcimInventoryItemRolesCreateParams) SetData(data *models.InventoryItemRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemRolesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
