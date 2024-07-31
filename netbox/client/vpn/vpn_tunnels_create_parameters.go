// Code generated by go-swagger; DO NOT EDIT.

package vpn

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

// NewVpnTunnelsCreateParams creates a new VpnTunnelsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVpnTunnelsCreateParams() *VpnTunnelsCreateParams {
	return &VpnTunnelsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVpnTunnelsCreateParamsWithTimeout creates a new VpnTunnelsCreateParams object
// with the ability to set a timeout on a request.
func NewVpnTunnelsCreateParamsWithTimeout(timeout time.Duration) *VpnTunnelsCreateParams {
	return &VpnTunnelsCreateParams{
		timeout: timeout,
	}
}

// NewVpnTunnelsCreateParamsWithContext creates a new VpnTunnelsCreateParams object
// with the ability to set a context for a request.
func NewVpnTunnelsCreateParamsWithContext(ctx context.Context) *VpnTunnelsCreateParams {
	return &VpnTunnelsCreateParams{
		Context: ctx,
	}
}

// NewVpnTunnelsCreateParamsWithHTTPClient creates a new VpnTunnelsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVpnTunnelsCreateParamsWithHTTPClient(client *http.Client) *VpnTunnelsCreateParams {
	return &VpnTunnelsCreateParams{
		HTTPClient: client,
	}
}

/*
VpnTunnelsCreateParams contains all the parameters to send to the API endpoint

	for the vpn tunnels create operation.

	Typically these are written to a http.Request.
*/
type VpnTunnelsCreateParams struct {

	// Data.
	Data *models.WritableTunnel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vpn tunnels create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsCreateParams) WithDefaults() *VpnTunnelsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vpn tunnels create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VpnTunnelsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) WithTimeout(timeout time.Duration) *VpnTunnelsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) WithContext(ctx context.Context) *VpnTunnelsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) WithHTTPClient(client *http.Client) *VpnTunnelsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) WithData(data *models.WritableTunnel) *VpnTunnelsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the vpn tunnels create params
func (o *VpnTunnelsCreateParams) SetData(data *models.WritableTunnel) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VpnTunnelsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
