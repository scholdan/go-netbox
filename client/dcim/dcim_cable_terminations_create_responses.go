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

// DcimCableTerminationsCreateReader is a Reader for the DcimCableTerminationsCreate structure.
type DcimCableTerminationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimCableTerminationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/cable-terminations/] dcim_cable-terminations_create", response, response.Code())
	}
}

// NewDcimCableTerminationsCreateCreated creates a DcimCableTerminationsCreateCreated with default headers values
func NewDcimCableTerminationsCreateCreated() *DcimCableTerminationsCreateCreated {
	return &DcimCableTerminationsCreateCreated{}
}

/*
DcimCableTerminationsCreateCreated describes a response with status code 201, with default header values.

DcimCableTerminationsCreateCreated dcim cable terminations create created
*/
type DcimCableTerminationsCreateCreated struct {
	Payload *models.CableTermination
}

// IsSuccess returns true when this dcim cable terminations create created response has a 2xx status code
func (o *DcimCableTerminationsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations create created response has a 3xx status code
func (o *DcimCableTerminationsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations create created response has a 4xx status code
func (o *DcimCableTerminationsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations create created response has a 5xx status code
func (o *DcimCableTerminationsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations create created response a status code equal to that given
func (o *DcimCableTerminationsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim cable terminations create created response
func (o *DcimCableTerminationsCreateCreated) Code() int {
	return 201
}

func (o *DcimCableTerminationsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcimCableTerminationsCreateCreated %s", 201, payload)
}

func (o *DcimCableTerminationsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcimCableTerminationsCreateCreated %s", 201, payload)
}

func (o *DcimCableTerminationsCreateCreated) GetPayload() *models.CableTermination {
	return o.Payload
}

func (o *DcimCableTerminationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CableTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
