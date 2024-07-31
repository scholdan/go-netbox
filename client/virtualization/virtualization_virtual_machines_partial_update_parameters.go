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
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/models"
)

// NewVirtualizationVirtualMachinesPartialUpdateParams creates a new VirtualizationVirtualMachinesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesPartialUpdateParams() *VirtualizationVirtualMachinesPartialUpdateParams {
	return &VirtualizationVirtualMachinesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesPartialUpdateParamsWithTimeout creates a new VirtualizationVirtualMachinesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesPartialUpdateParams {
	return &VirtualizationVirtualMachinesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesPartialUpdateParamsWithContext creates a new VirtualizationVirtualMachinesPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesPartialUpdateParams {
	return &VirtualizationVirtualMachinesPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesPartialUpdateParamsWithHTTPClient creates a new VirtualizationVirtualMachinesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesPartialUpdateParams {
	return &VirtualizationVirtualMachinesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
VirtualizationVirtualMachinesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the virtualization virtual machines partial update operation.

	Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesPartialUpdateParams struct {

	// Data.
	Data *models.WritableVirtualMachineWithConfigContext

	/* ID.

	   A unique integer value identifying this virtual machine.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithDefaults() *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithData(data *models.WritableVirtualMachineWithConfigContext) *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetData(data *models.WritableVirtualMachineWithConfigContext) {
	o.Data = data
}

// WithID adds the id to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WithID(id int64) *VirtualizationVirtualMachinesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines partial update params
func (o *VirtualizationVirtualMachinesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
