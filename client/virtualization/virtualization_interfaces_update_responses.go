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

// VirtualizationInterfacesUpdateReader is a Reader for the VirtualizationInterfacesUpdate structure.
type VirtualizationInterfacesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /virtualization/interfaces/{id}/] virtualization_interfaces_update", response, response.Code())
	}
}

// NewVirtualizationInterfacesUpdateOK creates a VirtualizationInterfacesUpdateOK with default headers values
func NewVirtualizationInterfacesUpdateOK() *VirtualizationInterfacesUpdateOK {
	return &VirtualizationInterfacesUpdateOK{}
}

/*
VirtualizationInterfacesUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesUpdateOK virtualization interfaces update o k
*/
type VirtualizationInterfacesUpdateOK struct {
	Payload *models.VMInterface
}

// IsSuccess returns true when this virtualization interfaces update o k response has a 2xx status code
func (o *VirtualizationInterfacesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization interfaces update o k response has a 3xx status code
func (o *VirtualizationInterfacesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization interfaces update o k response has a 4xx status code
func (o *VirtualizationInterfacesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization interfaces update o k response has a 5xx status code
func (o *VirtualizationInterfacesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization interfaces update o k response a status code equal to that given
func (o *VirtualizationInterfacesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization interfaces update o k response
func (o *VirtualizationInterfacesUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationInterfacesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/interfaces/{id}/][%d] virtualizationInterfacesUpdateOK %s", 200, payload)
}

func (o *VirtualizationInterfacesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/interfaces/{id}/][%d] virtualizationInterfacesUpdateOK %s", 200, payload)
}

func (o *VirtualizationInterfacesUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
