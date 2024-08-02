// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scholdan/go-netbox/models"
)

// VirtualizationVirtualMachinesBulkUpdateReader is a Reader for the VirtualizationVirtualMachinesBulkUpdate structure.
type VirtualizationVirtualMachinesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /virtualization/virtual-machines/] virtualization_virtual-machines_bulk_update", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesBulkUpdateOK creates a VirtualizationVirtualMachinesBulkUpdateOK with default headers values
func NewVirtualizationVirtualMachinesBulkUpdateOK() *VirtualizationVirtualMachinesBulkUpdateOK {
	return &VirtualizationVirtualMachinesBulkUpdateOK{}
}

/*
VirtualizationVirtualMachinesBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesBulkUpdateOK virtualization virtual machines bulk update o k
*/
type VirtualizationVirtualMachinesBulkUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines bulk update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines bulk update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines bulk update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines bulk update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines bulk update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual machines bulk update o k response
func (o *VirtualizationVirtualMachinesBulkUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
