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

	"github.com/scholdan/go-netbox/netbox/models"
)

// NewDcimVirtualDeviceContextsCreateParams creates a new DcimVirtualDeviceContextsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualDeviceContextsCreateParams() *DcimVirtualDeviceContextsCreateParams {
	return &DcimVirtualDeviceContextsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualDeviceContextsCreateParamsWithTimeout creates a new DcimVirtualDeviceContextsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualDeviceContextsCreateParamsWithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsCreateParams {
	return &DcimVirtualDeviceContextsCreateParams{
		timeout: timeout,
	}
}

// NewDcimVirtualDeviceContextsCreateParamsWithContext creates a new DcimVirtualDeviceContextsCreateParams object
// with the ability to set a context for a request.
func NewDcimVirtualDeviceContextsCreateParamsWithContext(ctx context.Context) *DcimVirtualDeviceContextsCreateParams {
	return &DcimVirtualDeviceContextsCreateParams{
		Context: ctx,
	}
}

// NewDcimVirtualDeviceContextsCreateParamsWithHTTPClient creates a new DcimVirtualDeviceContextsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualDeviceContextsCreateParamsWithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsCreateParams {
	return &DcimVirtualDeviceContextsCreateParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualDeviceContextsCreateParams contains all the parameters to send to the API endpoint

	for the dcim virtual device contexts create operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualDeviceContextsCreateParams struct {

	// Data.
	Data *models.WritableVirtualDeviceContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual device contexts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsCreateParams) WithDefaults() *DcimVirtualDeviceContextsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual device contexts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) WithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) WithContext(ctx context.Context) *DcimVirtualDeviceContextsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) WithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) WithData(data *models.WritableVirtualDeviceContext) *DcimVirtualDeviceContextsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual device contexts create params
func (o *DcimVirtualDeviceContextsCreateParams) SetData(data *models.WritableVirtualDeviceContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualDeviceContextsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
