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

// DcimConsoleServerPortsTraceReader is a Reader for the DcimConsoleServerPortsTrace structure.
type DcimConsoleServerPortsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/console-server-ports/{id}/trace/] dcim_console-server-ports_trace", response, response.Code())
	}
}

// NewDcimConsoleServerPortsTraceOK creates a DcimConsoleServerPortsTraceOK with default headers values
func NewDcimConsoleServerPortsTraceOK() *DcimConsoleServerPortsTraceOK {
	return &DcimConsoleServerPortsTraceOK{}
}

/*
DcimConsoleServerPortsTraceOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsTraceOK dcim console server ports trace o k
*/
type DcimConsoleServerPortsTraceOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports trace o k response has a 2xx status code
func (o *DcimConsoleServerPortsTraceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports trace o k response has a 3xx status code
func (o *DcimConsoleServerPortsTraceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports trace o k response has a 4xx status code
func (o *DcimConsoleServerPortsTraceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports trace o k response has a 5xx status code
func (o *DcimConsoleServerPortsTraceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports trace o k response a status code equal to that given
func (o *DcimConsoleServerPortsTraceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports trace o k response
func (o *DcimConsoleServerPortsTraceOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsTraceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/trace/][%d] dcimConsoleServerPortsTraceOK %s", 200, payload)
}

func (o *DcimConsoleServerPortsTraceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/console-server-ports/{id}/trace/][%d] dcimConsoleServerPortsTraceOK %s", 200, payload)
}

func (o *DcimConsoleServerPortsTraceOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
