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

// DcimCableTerminationsUpdateReader is a Reader for the DcimCableTerminationsUpdate structure.
type DcimCableTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCableTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/cable-terminations/{id}/] dcim_cable-terminations_update", response, response.Code())
	}
}

// NewDcimCableTerminationsUpdateOK creates a DcimCableTerminationsUpdateOK with default headers values
func NewDcimCableTerminationsUpdateOK() *DcimCableTerminationsUpdateOK {
	return &DcimCableTerminationsUpdateOK{}
}

/*
DcimCableTerminationsUpdateOK describes a response with status code 200, with default header values.

DcimCableTerminationsUpdateOK dcim cable terminations update o k
*/
type DcimCableTerminationsUpdateOK struct {
	Payload *models.CableTermination
}

// IsSuccess returns true when this dcim cable terminations update o k response has a 2xx status code
func (o *DcimCableTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations update o k response has a 3xx status code
func (o *DcimCableTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations update o k response has a 4xx status code
func (o *DcimCableTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations update o k response has a 5xx status code
func (o *DcimCableTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations update o k response a status code equal to that given
func (o *DcimCableTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cable terminations update o k response
func (o *DcimCableTerminationsUpdateOK) Code() int {
	return 200
}

func (o *DcimCableTerminationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateOK %s", 200, payload)
}

func (o *DcimCableTerminationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateOK %s", 200, payload)
}

func (o *DcimCableTerminationsUpdateOK) GetPayload() *models.CableTermination {
	return o.Payload
}

func (o *DcimCableTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CableTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
