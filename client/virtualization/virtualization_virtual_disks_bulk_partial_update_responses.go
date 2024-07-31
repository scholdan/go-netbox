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

	"github.com/fbreckle/go-netbox/models"
)

// VirtualizationVirtualDisksBulkPartialUpdateReader is a Reader for the VirtualizationVirtualDisksBulkPartialUpdate structure.
type VirtualizationVirtualDisksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualDisksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualDisksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /virtualization/virtual-disks/] virtualization_virtual-disks_bulk_partial_update", response, response.Code())
	}
}

// NewVirtualizationVirtualDisksBulkPartialUpdateOK creates a VirtualizationVirtualDisksBulkPartialUpdateOK with default headers values
func NewVirtualizationVirtualDisksBulkPartialUpdateOK() *VirtualizationVirtualDisksBulkPartialUpdateOK {
	return &VirtualizationVirtualDisksBulkPartialUpdateOK{}
}

/*
VirtualizationVirtualDisksBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualDisksBulkPartialUpdateOK virtualization virtual disks bulk partial update o k
*/
type VirtualizationVirtualDisksBulkPartialUpdateOK struct {
	Payload *models.VirtualDisk
}

// IsSuccess returns true when this virtualization virtual disks bulk partial update o k response has a 2xx status code
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual disks bulk partial update o k response has a 3xx status code
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual disks bulk partial update o k response has a 4xx status code
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual disks bulk partial update o k response has a 5xx status code
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual disks bulk partial update o k response a status code equal to that given
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization virtual disks bulk partial update o k response
func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/virtual-disks/][%d] virtualizationVirtualDisksBulkPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/virtual-disks/][%d] virtualizationVirtualDisksBulkPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) GetPayload() *models.VirtualDisk {
	return o.Payload
}

func (o *VirtualizationVirtualDisksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDisk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
