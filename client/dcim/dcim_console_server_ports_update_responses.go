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

// DcimConsoleServerPortsUpdateReader is a Reader for the DcimConsoleServerPortsUpdate structure.
type DcimConsoleServerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/console-server-ports/{id}/] dcim_console-server-ports_update", response, response.Code())
	}
}

// NewDcimConsoleServerPortsUpdateOK creates a DcimConsoleServerPortsUpdateOK with default headers values
func NewDcimConsoleServerPortsUpdateOK() *DcimConsoleServerPortsUpdateOK {
	return &DcimConsoleServerPortsUpdateOK{}
}

/*
DcimConsoleServerPortsUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsUpdateOK dcim console server ports update o k
*/
type DcimConsoleServerPortsUpdateOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports update o k response has a 2xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports update o k response has a 3xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports update o k response has a 4xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports update o k response has a 5xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports update o k response a status code equal to that given
func (o *DcimConsoleServerPortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports update o k response
func (o *DcimConsoleServerPortsUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortsUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
