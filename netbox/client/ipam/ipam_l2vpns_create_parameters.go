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

package ipam

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

// NewIpamL2vpnsCreateParams creates a new IpamL2vpnsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnsCreateParams() *IpamL2vpnsCreateParams {
	return &IpamL2vpnsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnsCreateParamsWithTimeout creates a new IpamL2vpnsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnsCreateParamsWithTimeout(timeout time.Duration) *IpamL2vpnsCreateParams {
	return &IpamL2vpnsCreateParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnsCreateParamsWithContext creates a new IpamL2vpnsCreateParams object
// with the ability to set a context for a request.
func NewIpamL2vpnsCreateParamsWithContext(ctx context.Context) *IpamL2vpnsCreateParams {
	return &IpamL2vpnsCreateParams{
		Context: ctx,
	}
}

// NewIpamL2vpnsCreateParamsWithHTTPClient creates a new IpamL2vpnsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnsCreateParamsWithHTTPClient(client *http.Client) *IpamL2vpnsCreateParams {
	return &IpamL2vpnsCreateParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnsCreateParams contains all the parameters to send to the API endpoint

	for the ipam l2vpns create operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnsCreateParams struct {

	// Data.
	Data *models.WritableL2VPN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpns create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsCreateParams) WithDefaults() *IpamL2vpnsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpns create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) WithTimeout(timeout time.Duration) *IpamL2vpnsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) WithContext(ctx context.Context) *IpamL2vpnsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) WithHTTPClient(client *http.Client) *IpamL2vpnsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) WithData(data *models.WritableL2VPN) *IpamL2vpnsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam l2vpns create params
func (o *IpamL2vpnsCreateParams) SetData(data *models.WritableL2VPN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
