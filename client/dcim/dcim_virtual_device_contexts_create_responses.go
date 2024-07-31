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

	"github.com/fbreckle/go-netbox/models"
)

// DcimVirtualDeviceContextsCreateReader is a Reader for the DcimVirtualDeviceContextsCreate structure.
type DcimVirtualDeviceContextsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualDeviceContextsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimVirtualDeviceContextsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/virtual-device-contexts/] dcim_virtual-device-contexts_create", response, response.Code())
	}
}

// NewDcimVirtualDeviceContextsCreateCreated creates a DcimVirtualDeviceContextsCreateCreated with default headers values
func NewDcimVirtualDeviceContextsCreateCreated() *DcimVirtualDeviceContextsCreateCreated {
	return &DcimVirtualDeviceContextsCreateCreated{}
}

/*
DcimVirtualDeviceContextsCreateCreated describes a response with status code 201, with default header values.

DcimVirtualDeviceContextsCreateCreated dcim virtual device contexts create created
*/
type DcimVirtualDeviceContextsCreateCreated struct {
	Payload *models.VirtualDeviceContext
}

// IsSuccess returns true when this dcim virtual device contexts create created response has a 2xx status code
func (o *DcimVirtualDeviceContextsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual device contexts create created response has a 3xx status code
func (o *DcimVirtualDeviceContextsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts create created response has a 4xx status code
func (o *DcimVirtualDeviceContextsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual device contexts create created response has a 5xx status code
func (o *DcimVirtualDeviceContextsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts create created response a status code equal to that given
func (o *DcimVirtualDeviceContextsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim virtual device contexts create created response
func (o *DcimVirtualDeviceContextsCreateCreated) Code() int {
	return 201
}

func (o *DcimVirtualDeviceContextsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/virtual-device-contexts/][%d] dcimVirtualDeviceContextsCreateCreated %s", 201, payload)
}

func (o *DcimVirtualDeviceContextsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/virtual-device-contexts/][%d] dcimVirtualDeviceContextsCreateCreated %s", 201, payload)
}

func (o *DcimVirtualDeviceContextsCreateCreated) GetPayload() *models.VirtualDeviceContext {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDeviceContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
