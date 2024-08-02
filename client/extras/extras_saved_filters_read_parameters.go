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
)

// NewExtrasSavedFiltersReadParams creates a new ExtrasSavedFiltersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasSavedFiltersReadParams() *ExtrasSavedFiltersReadParams {
	return &ExtrasSavedFiltersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasSavedFiltersReadParamsWithTimeout creates a new ExtrasSavedFiltersReadParams object
// with the ability to set a timeout on a request.
func NewExtrasSavedFiltersReadParamsWithTimeout(timeout time.Duration) *ExtrasSavedFiltersReadParams {
	return &ExtrasSavedFiltersReadParams{
		timeout: timeout,
	}
}

// NewExtrasSavedFiltersReadParamsWithContext creates a new ExtrasSavedFiltersReadParams object
// with the ability to set a context for a request.
func NewExtrasSavedFiltersReadParamsWithContext(ctx context.Context) *ExtrasSavedFiltersReadParams {
	return &ExtrasSavedFiltersReadParams{
		Context: ctx,
	}
}

// NewExtrasSavedFiltersReadParamsWithHTTPClient creates a new ExtrasSavedFiltersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasSavedFiltersReadParamsWithHTTPClient(client *http.Client) *ExtrasSavedFiltersReadParams {
	return &ExtrasSavedFiltersReadParams{
		HTTPClient: client,
	}
}

/*
ExtrasSavedFiltersReadParams contains all the parameters to send to the API endpoint

	for the extras saved filters read operation.

	Typically these are written to a http.Request.
*/
type ExtrasSavedFiltersReadParams struct {

	/* ID.

	   A unique integer value identifying this saved filter.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras saved filters read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersReadParams) WithDefaults() *ExtrasSavedFiltersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras saved filters read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasSavedFiltersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) WithTimeout(timeout time.Duration) *ExtrasSavedFiltersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) WithContext(ctx context.Context) *ExtrasSavedFiltersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) WithHTTPClient(client *http.Client) *ExtrasSavedFiltersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) WithID(id int64) *ExtrasSavedFiltersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras saved filters read params
func (o *ExtrasSavedFiltersReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasSavedFiltersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
