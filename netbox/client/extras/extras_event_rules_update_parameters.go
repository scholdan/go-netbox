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
	"github.com/go-openapi/swag"

	"github.com/scholdan/go-netbox/models"
)

// NewExtrasEventRulesUpdateParams creates a new ExtrasEventRulesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasEventRulesUpdateParams() *ExtrasEventRulesUpdateParams {
	return &ExtrasEventRulesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasEventRulesUpdateParamsWithTimeout creates a new ExtrasEventRulesUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasEventRulesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasEventRulesUpdateParams {
	return &ExtrasEventRulesUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasEventRulesUpdateParamsWithContext creates a new ExtrasEventRulesUpdateParams object
// with the ability to set a context for a request.
func NewExtrasEventRulesUpdateParamsWithContext(ctx context.Context) *ExtrasEventRulesUpdateParams {
	return &ExtrasEventRulesUpdateParams{
		Context: ctx,
	}
}

// NewExtrasEventRulesUpdateParamsWithHTTPClient creates a new ExtrasEventRulesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasEventRulesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasEventRulesUpdateParams {
	return &ExtrasEventRulesUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasEventRulesUpdateParams contains all the parameters to send to the API endpoint

	for the extras event rules update operation.

	Typically these are written to a http.Request.
*/
type ExtrasEventRulesUpdateParams struct {

	// Data.
	Data *models.WritableEventRule

	/* ID.

	   A unique integer value identifying this event_rule.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras event rules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasEventRulesUpdateParams) WithDefaults() *ExtrasEventRulesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras event rules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasEventRulesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasEventRulesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) WithContext(ctx context.Context) *ExtrasEventRulesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasEventRulesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) WithData(data *models.WritableEventRule) *ExtrasEventRulesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) SetData(data *models.WritableEventRule) {
	o.Data = data
}

// WithID adds the id to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) WithID(id int64) *ExtrasEventRulesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras event rules update params
func (o *ExtrasEventRulesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasEventRulesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
