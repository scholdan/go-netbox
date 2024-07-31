// Code generated by go-swagger; DO NOT EDIT.

package virtualization

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

// NewVirtualizationVirtualDisksBulkUpdateParams creates a new VirtualizationVirtualDisksBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualDisksBulkUpdateParams() *VirtualizationVirtualDisksBulkUpdateParams {
	return &VirtualizationVirtualDisksBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualDisksBulkUpdateParamsWithTimeout creates a new VirtualizationVirtualDisksBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualDisksBulkUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualDisksBulkUpdateParams {
	return &VirtualizationVirtualDisksBulkUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualDisksBulkUpdateParamsWithContext creates a new VirtualizationVirtualDisksBulkUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualDisksBulkUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualDisksBulkUpdateParams {
	return &VirtualizationVirtualDisksBulkUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualDisksBulkUpdateParamsWithHTTPClient creates a new VirtualizationVirtualDisksBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualDisksBulkUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualDisksBulkUpdateParams {
	return &VirtualizationVirtualDisksBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
VirtualizationVirtualDisksBulkUpdateParams contains all the parameters to send to the API endpoint

	for the virtualization virtual disks bulk update operation.

	Typically these are written to a http.Request.
*/
type VirtualizationVirtualDisksBulkUpdateParams struct {

	// Data.
	Data *models.WritableVirtualDisk

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual disks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksBulkUpdateParams) WithDefaults() *VirtualizationVirtualDisksBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual disks bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualDisksBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualDisksBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualDisksBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) WithData(data *models.WritableVirtualDisk) *VirtualizationVirtualDisksBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual disks bulk update params
func (o *VirtualizationVirtualDisksBulkUpdateParams) SetData(data *models.WritableVirtualDisk) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualDisksBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
