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

// NewDcimModuleBayTemplatesPartialUpdateParams creates a new DcimModuleBayTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBayTemplatesPartialUpdateParams() *DcimModuleBayTemplatesPartialUpdateParams {
	return &DcimModuleBayTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBayTemplatesPartialUpdateParamsWithTimeout creates a new DcimModuleBayTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBayTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimModuleBayTemplatesPartialUpdateParams {
	return &DcimModuleBayTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModuleBayTemplatesPartialUpdateParamsWithContext creates a new DcimModuleBayTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimModuleBayTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimModuleBayTemplatesPartialUpdateParams {
	return &DcimModuleBayTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimModuleBayTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimModuleBayTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBayTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimModuleBayTemplatesPartialUpdateParams {
	return &DcimModuleBayTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimModuleBayTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim module bay templates partial update operation.

	Typically these are written to a http.Request.
*/
type DcimModuleBayTemplatesPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim module bay templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithDefaults() *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bay templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithData(data *models.WritableModuleBayTemplate) *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetData(data *models.WritableModuleBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) WithID(id int64) *DcimModuleBayTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim module bay templates partial update params
func (o *DcimModuleBayTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBayTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
