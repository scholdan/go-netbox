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

// DcimConsoleServerPortTemplatesReadReader is a Reader for the DcimConsoleServerPortTemplatesRead structure.
type DcimConsoleServerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/console-server-port-templates/{id}/] dcim_console-server-port-templates_read", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesReadOK creates a DcimConsoleServerPortTemplatesReadOK with default headers values
func NewDcimConsoleServerPortTemplatesReadOK() *DcimConsoleServerPortTemplatesReadOK {
	return &DcimConsoleServerPortTemplatesReadOK{}
}

/*
DcimConsoleServerPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesReadOK dcim console server port templates read o k
*/
type DcimConsoleServerPortTemplatesReadOK struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates read o k response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates read o k response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates read o k response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates read o k response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates read o k response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server port templates read o k response
func (o *DcimConsoleServerPortTemplatesReadOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortTemplatesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesReadOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
