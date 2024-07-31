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

	"github.com/fbreckle/go-netbox/models"
)

// NewDcimRackRolesPartialUpdateParams creates a new DcimRackRolesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesPartialUpdateParams() *DcimRackRolesPartialUpdateParams {
	return &DcimRackRolesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesPartialUpdateParamsWithTimeout creates a new DcimRackRolesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackRolesPartialUpdateParams {
	return &DcimRackRolesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesPartialUpdateParamsWithContext creates a new DcimRackRolesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackRolesPartialUpdateParamsWithContext(ctx context.Context) *DcimRackRolesPartialUpdateParams {
	return &DcimRackRolesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackRolesPartialUpdateParamsWithHTTPClient creates a new DcimRackRolesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackRolesPartialUpdateParams {
	return &DcimRackRolesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimRackRolesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim rack roles partial update operation.

	Typically these are written to a http.Request.
*/
type DcimRackRolesPartialUpdateParams struct {

	// Data.
	Data *models.RackRole

	/* ID.

	   A unique integer value identifying this rack role.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesPartialUpdateParams) WithDefaults() *DcimRackRolesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackRolesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) WithContext(ctx context.Context) *DcimRackRolesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackRolesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) WithData(data *models.RackRole) *DcimRackRolesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) SetData(data *models.RackRole) {
	o.Data = data
}

// WithID adds the id to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) WithID(id int64) *DcimRackRolesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack roles partial update params
func (o *DcimRackRolesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
