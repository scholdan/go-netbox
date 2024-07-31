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

// DcimFrontPortsReadReader is a Reader for the DcimFrontPortsRead structure.
type DcimFrontPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/front-ports/{id}/] dcim_front-ports_read", response, response.Code())
	}
}

// NewDcimFrontPortsReadOK creates a DcimFrontPortsReadOK with default headers values
func NewDcimFrontPortsReadOK() *DcimFrontPortsReadOK {
	return &DcimFrontPortsReadOK{}
}

/*
DcimFrontPortsReadOK describes a response with status code 200, with default header values.

DcimFrontPortsReadOK dcim front ports read o k
*/
type DcimFrontPortsReadOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports read o k response has a 2xx status code
func (o *DcimFrontPortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports read o k response has a 3xx status code
func (o *DcimFrontPortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports read o k response has a 4xx status code
func (o *DcimFrontPortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports read o k response has a 5xx status code
func (o *DcimFrontPortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports read o k response a status code equal to that given
func (o *DcimFrontPortsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim front ports read o k response
func (o *DcimFrontPortsReadOK) Code() int {
	return 200
}

func (o *DcimFrontPortsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcimFrontPortsReadOK %s", 200, payload)
}

func (o *DcimFrontPortsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/front-ports/{id}/][%d] dcimFrontPortsReadOK %s", 200, payload)
}

func (o *DcimFrontPortsReadOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
