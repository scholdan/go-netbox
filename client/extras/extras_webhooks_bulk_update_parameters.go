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

	"github.com/fbreckle/go-netbox/models"
)

// NewExtrasWebhooksBulkUpdateParams creates a new ExtrasWebhooksBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksBulkUpdateParams() *ExtrasWebhooksBulkUpdateParams {
	return &ExtrasWebhooksBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksBulkUpdateParamsWithTimeout creates a new ExtrasWebhooksBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksBulkUpdateParams {
	return &ExtrasWebhooksBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksBulkUpdateParamsWithContext creates a new ExtrasWebhooksBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksBulkUpdateParamsWithContext(ctx context.Context) *ExtrasWebhooksBulkUpdateParams {
	return &ExtrasWebhooksBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksBulkUpdateParamsWithHTTPClient creates a new ExtrasWebhooksBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksBulkUpdateParams {
	return &ExtrasWebhooksBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasWebhooksBulkUpdateParams contains all the parameters to send to the API endpoint

	for the extras webhooks bulk update operation.

	Typically these are written to a http.Request.
*/
type ExtrasWebhooksBulkUpdateParams struct {

	// Data.
	Data *models.Webhook

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkUpdateParams) WithDefaults() *ExtrasWebhooksBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) WithContext(ctx context.Context) *ExtrasWebhooksBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) WithData(data *models.Webhook) *ExtrasWebhooksBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras webhooks bulk update params
func (o *ExtrasWebhooksBulkUpdateParams) SetData(data *models.Webhook) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
