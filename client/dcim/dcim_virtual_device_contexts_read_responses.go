// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// DcimVirtualDeviceContextsReadReader is a Reader for the DcimVirtualDeviceContextsRead structure.
type DcimVirtualDeviceContextsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualDeviceContextsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualDeviceContextsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/virtual-device-contexts/{id}/] dcim_virtual-device-contexts_read", response, response.Code())
	}
}

// NewDcimVirtualDeviceContextsReadOK creates a DcimVirtualDeviceContextsReadOK with default headers values
func NewDcimVirtualDeviceContextsReadOK() *DcimVirtualDeviceContextsReadOK {
	return &DcimVirtualDeviceContextsReadOK{}
}

/*
DcimVirtualDeviceContextsReadOK describes a response with status code 200, with default header values.

DcimVirtualDeviceContextsReadOK dcim virtual device contexts read o k
*/
type DcimVirtualDeviceContextsReadOK struct {
	Payload *models.VirtualDeviceContext
}

// IsSuccess returns true when this dcim virtual device contexts read o k response has a 2xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual device contexts read o k response has a 3xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts read o k response has a 4xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual device contexts read o k response has a 5xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts read o k response a status code equal to that given
func (o *DcimVirtualDeviceContextsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim virtual device contexts read o k response
func (o *DcimVirtualDeviceContextsReadOK) Code() int {
	return 200
}

func (o *DcimVirtualDeviceContextsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadOK %s", 200, payload)
}

func (o *DcimVirtualDeviceContextsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadOK %s", 200, payload)
}

func (o *DcimVirtualDeviceContextsReadOK) GetPayload() *models.VirtualDeviceContext {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDeviceContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}