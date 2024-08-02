// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// NewTenancyContactGroupsUpdateParams creates a new TenancyContactGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactGroupsUpdateParams() *TenancyContactGroupsUpdateParams {
	return &TenancyContactGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactGroupsUpdateParamsWithTimeout creates a new TenancyContactGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactGroupsUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactGroupsUpdateParams {
	return &TenancyContactGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactGroupsUpdateParamsWithContext creates a new TenancyContactGroupsUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactGroupsUpdateParamsWithContext(ctx context.Context) *TenancyContactGroupsUpdateParams {
	return &TenancyContactGroupsUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactGroupsUpdateParamsWithHTTPClient creates a new TenancyContactGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactGroupsUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactGroupsUpdateParams {
	return &TenancyContactGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactGroupsUpdateParams contains all the parameters to send to the API endpoint

	for the tenancy contact groups update operation.

	Typically these are written to a http.Request.
*/
type TenancyContactGroupsUpdateParams struct {

	// Data.
	Data *models.WritableContactGroup

	/* ID.

	   A unique integer value identifying this contact group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactGroupsUpdateParams) WithDefaults() *TenancyContactGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) WithContext(ctx context.Context) *TenancyContactGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) WithData(data *models.WritableContactGroup) *TenancyContactGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) SetData(data *models.WritableContactGroup) {
	o.Data = data
}

// WithID adds the id to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) WithID(id int64) *TenancyContactGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy contact groups update params
func (o *TenancyContactGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
