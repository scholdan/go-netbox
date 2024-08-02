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

// NewDcimRearPortTemplatesBulkPartialUpdateParams creates a new DcimRearPortTemplatesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesBulkPartialUpdateParams() *DcimRearPortTemplatesBulkPartialUpdateParams {
	return &DcimRearPortTemplatesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesBulkPartialUpdateParamsWithTimeout creates a new DcimRearPortTemplatesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkPartialUpdateParams {
	return &DcimRearPortTemplatesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesBulkPartialUpdateParamsWithContext creates a new DcimRearPortTemplatesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRearPortTemplatesBulkPartialUpdateParams {
	return &DcimRearPortTemplatesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesBulkPartialUpdateParamsWithHTTPClient creates a new DcimRearPortTemplatesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkPartialUpdateParams {
	return &DcimRearPortTemplatesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimRearPortTemplatesBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim rear port templates bulk partial update operation.

	Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRearPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WithDefaults() *DcimRearPortTemplatesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRearPortTemplatesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WithData(data *models.WritableRearPortTemplate) *DcimRearPortTemplatesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear port templates bulk partial update params
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) SetData(data *models.WritableRearPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
