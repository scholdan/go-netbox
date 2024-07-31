// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

	"github.com/scholdan/go-netbox/netbox/models"
)

// NewDcimPowerPanelsPartialUpdateParams creates a new DcimPowerPanelsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsPartialUpdateParams() *DcimPowerPanelsPartialUpdateParams {
	return &DcimPowerPanelsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsPartialUpdateParamsWithTimeout creates a new DcimPowerPanelsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsPartialUpdateParams {
	return &DcimPowerPanelsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsPartialUpdateParamsWithContext creates a new DcimPowerPanelsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerPanelsPartialUpdateParams {
	return &DcimPowerPanelsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsPartialUpdateParamsWithHTTPClient creates a new DcimPowerPanelsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsPartialUpdateParams {
	return &DcimPowerPanelsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimPowerPanelsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim power panels partial update operation.

	Typically these are written to a http.Request.
*/
type DcimPowerPanelsPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerPanel

	/* ID.

	   A unique integer value identifying this power panel.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsPartialUpdateParams) WithDefaults() *DcimPowerPanelsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerPanelsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) WithData(data *models.WritablePowerPanel) *DcimPowerPanelsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) SetData(data *models.WritablePowerPanel) {
	o.Data = data
}

// WithID adds the id to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) WithID(id int64) *DcimPowerPanelsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels partial update params
func (o *DcimPowerPanelsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
