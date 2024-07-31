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

// VirtualizationVirtualDisksReadReader is a Reader for the VirtualizationVirtualDisksRead structure.
type VirtualizationVirtualDisksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualDisksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualDisksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /virtualization/virtual-disks/{id}/] virtualization_virtual-disks_read", response, response.Code())
	}
}

// NewVirtualizationVirtualDisksReadOK creates a VirtualizationVirtualDisksReadOK with default headers values
func NewVirtualizationVirtualDisksReadOK() *VirtualizationVirtualDisksReadOK {
	return &VirtualizationVirtualDisksReadOK{}
}

/*
VirtualizationVirtualDisksReadOK describes a response with status code 200, with default header values.

VirtualizationVirtualDisksReadOK virtualization virtual disks read o k
*/
type VirtualizationVirtualDisksReadOK struct {
	Payload *models.VirtualDisk
}

// IsSuccess returns true when this virtualization virtual disks read o k response has a 2xx status code
func (o *VirtualizationVirtualDisksReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual disks read o k response has a 3xx status code
func (o *VirtualizationVirtualDisksReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual disks read o k response has a 4xx status code
func (o *VirtualizationVirtualDisksReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual disks read o k response has a 5xx status code
func (o *VirtualizationVirtualDisksReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual disks read o k response a status code equal to that given
func (o *VirtualizationVirtualDisksReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual disks read o k response
func (o *VirtualizationVirtualDisksReadOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualDisksReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksReadOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksReadOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksReadOK) GetPayload() *models.VirtualDisk {
	return o.Payload
}

func (o *VirtualizationVirtualDisksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDisk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
