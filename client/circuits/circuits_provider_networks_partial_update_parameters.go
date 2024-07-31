// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// NewCircuitsProviderNetworksPartialUpdateParams creates a new CircuitsProviderNetworksPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProviderNetworksPartialUpdateParams() *CircuitsProviderNetworksPartialUpdateParams {
	return &CircuitsProviderNetworksPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProviderNetworksPartialUpdateParamsWithTimeout creates a new CircuitsProviderNetworksPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProviderNetworksPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProviderNetworksPartialUpdateParams {
	return &CircuitsProviderNetworksPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProviderNetworksPartialUpdateParamsWithContext creates a new CircuitsProviderNetworksPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProviderNetworksPartialUpdateParamsWithContext(ctx context.Context) *CircuitsProviderNetworksPartialUpdateParams {
	return &CircuitsProviderNetworksPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProviderNetworksPartialUpdateParamsWithHTTPClient creates a new CircuitsProviderNetworksPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProviderNetworksPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProviderNetworksPartialUpdateParams {
	return &CircuitsProviderNetworksPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
CircuitsProviderNetworksPartialUpdateParams contains all the parameters to send to the API endpoint

	for the circuits provider networks partial update operation.

	Typically these are written to a http.Request.
*/
type CircuitsProviderNetworksPartialUpdateParams struct {

	// Data.
	Data *models.WritableProviderNetwork

	/* ID.

	   A unique integer value identifying this provider network.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits provider networks partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksPartialUpdateParams) WithDefaults() *CircuitsProviderNetworksPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits provider networks partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProviderNetworksPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) WithContext(ctx context.Context) *CircuitsProviderNetworksPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProviderNetworksPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) WithData(data *models.WritableProviderNetwork) *CircuitsProviderNetworksPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) SetData(data *models.WritableProviderNetwork) {
	o.Data = data
}

// WithID adds the id to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) WithID(id int64) *CircuitsProviderNetworksPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits provider networks partial update params
func (o *CircuitsProviderNetworksPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProviderNetworksPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
