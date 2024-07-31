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

// DcimRearPortsCreateReader is a Reader for the DcimRearPortsCreate structure.
type DcimRearPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/rear-ports/] dcim_rear-ports_create", response, response.Code())
	}
}

// NewDcimRearPortsCreateCreated creates a DcimRearPortsCreateCreated with default headers values
func NewDcimRearPortsCreateCreated() *DcimRearPortsCreateCreated {
	return &DcimRearPortsCreateCreated{}
}

/*
DcimRearPortsCreateCreated describes a response with status code 201, with default header values.

DcimRearPortsCreateCreated dcim rear ports create created
*/
type DcimRearPortsCreateCreated struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports create created response has a 2xx status code
func (o *DcimRearPortsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports create created response has a 3xx status code
func (o *DcimRearPortsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports create created response has a 4xx status code
func (o *DcimRearPortsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports create created response has a 5xx status code
func (o *DcimRearPortsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports create created response a status code equal to that given
func (o *DcimRearPortsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim rear ports create created response
func (o *DcimRearPortsCreateCreated) Code() int {
	return 201
}

func (o *DcimRearPortsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcimRearPortsCreateCreated %s", 201, payload)
}

func (o *DcimRearPortsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/rear-ports/][%d] dcimRearPortsCreateCreated %s", 201, payload)
}

func (o *DcimRearPortsCreateCreated) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
