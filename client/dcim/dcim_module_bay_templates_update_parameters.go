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

// NewDcimModuleBayTemplatesUpdateParams creates a new DcimModuleBayTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBayTemplatesUpdateParams() *DcimModuleBayTemplatesUpdateParams {
	return &DcimModuleBayTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBayTemplatesUpdateParamsWithTimeout creates a new DcimModuleBayTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBayTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimModuleBayTemplatesUpdateParams {
	return &DcimModuleBayTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModuleBayTemplatesUpdateParamsWithContext creates a new DcimModuleBayTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimModuleBayTemplatesUpdateParamsWithContext(ctx context.Context) *DcimModuleBayTemplatesUpdateParams {
	return &DcimModuleBayTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimModuleBayTemplatesUpdateParamsWithHTTPClient creates a new DcimModuleBayTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBayTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimModuleBayTemplatesUpdateParams {
	return &DcimModuleBayTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimModuleBayTemplatesUpdateParams contains all the parameters to send to the API endpoint

	for the dcim module bay templates update operation.

	Typically these are written to a http.Request.
*/
type DcimModuleBayTemplatesUpdateParams struct {

	// Data.
	Data *models.WritableModuleBayTemplate

	/* ID.

	   A unique integer value identifying this module bay template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module bay templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesUpdateParams) WithDefaults() *DcimModuleBayTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bay templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimModuleBayTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) WithContext(ctx context.Context) *DcimModuleBayTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimModuleBayTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) WithData(data *models.WritableModuleBayTemplate) *DcimModuleBayTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) SetData(data *models.WritableModuleBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) WithID(id int64) *DcimModuleBayTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim module bay templates update params
func (o *DcimModuleBayTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBayTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
