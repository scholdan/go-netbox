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

// VirtualizationVirtualMachinesReadReader is a Reader for the VirtualizationVirtualMachinesRead structure.
type VirtualizationVirtualMachinesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /virtualization/virtual-machines/{id}/] virtualization_virtual-machines_read", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesReadOK creates a VirtualizationVirtualMachinesReadOK with default headers values
func NewVirtualizationVirtualMachinesReadOK() *VirtualizationVirtualMachinesReadOK {
	return &VirtualizationVirtualMachinesReadOK{}
}

/*
VirtualizationVirtualMachinesReadOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesReadOK virtualization virtual machines read o k
*/
type VirtualizationVirtualMachinesReadOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines read o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines read o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines read o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines read o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines read o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual machines read o k response
func (o *VirtualizationVirtualMachinesReadOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualMachinesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesReadOK %s", 200, payload)
}

func (o *VirtualizationVirtualMachinesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesReadOK %s", 200, payload)
}

func (o *VirtualizationVirtualMachinesReadOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}