// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// NewExtrasExportTemplatesBulkUpdateParams creates a new ExtrasExportTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesBulkUpdateParams() *ExtrasExportTemplatesBulkUpdateParams {
	return &ExtrasExportTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesBulkUpdateParamsWithTimeout creates a new ExtrasExportTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesBulkUpdateParams {
	return &ExtrasExportTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesBulkUpdateParamsWithContext creates a new ExtrasExportTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesBulkUpdateParamsWithContext(ctx context.Context) *ExtrasExportTemplatesBulkUpdateParams {
	return &ExtrasExportTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesBulkUpdateParamsWithHTTPClient creates a new ExtrasExportTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesBulkUpdateParams {
	return &ExtrasExportTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasExportTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the extras export templates bulk update operation.

	Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.ExportTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesBulkUpdateParams) WithDefaults() *ExtrasExportTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) WithContext(ctx context.Context) *ExtrasExportTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) WithData(data *models.ExportTemplate) *ExtrasExportTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras export templates bulk update params
func (o *ExtrasExportTemplatesBulkUpdateParams) SetData(data *models.ExportTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
