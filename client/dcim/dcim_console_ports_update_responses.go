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

// DcimConsolePortsUpdateReader is a Reader for the DcimConsolePortsUpdate structure.
type DcimConsolePortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/console-ports/{id}/] dcim_console-ports_update", response, response.Code())
	}
}

// NewDcimConsolePortsUpdateOK creates a DcimConsolePortsUpdateOK with default headers values
func NewDcimConsolePortsUpdateOK() *DcimConsolePortsUpdateOK {
	return &DcimConsolePortsUpdateOK{}
}

/*
DcimConsolePortsUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsUpdateOK dcim console ports update o k
*/
type DcimConsolePortsUpdateOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports update o k response has a 2xx status code
func (o *DcimConsolePortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports update o k response has a 3xx status code
func (o *DcimConsolePortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports update o k response has a 4xx status code
func (o *DcimConsolePortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports update o k response has a 5xx status code
func (o *DcimConsolePortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports update o k response a status code equal to that given
func (o *DcimConsolePortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console ports update o k response
func (o *DcimConsolePortsUpdateOK) Code() int {
	return 200
}

func (o *DcimConsolePortsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcimConsolePortsUpdateOK %s", 200, payload)
}

func (o *DcimConsolePortsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-ports/{id}/][%d] dcimConsolePortsUpdateOK %s", 200, payload)
}

func (o *DcimConsolePortsUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
