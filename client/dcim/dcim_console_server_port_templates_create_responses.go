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

// DcimConsoleServerPortTemplatesCreateReader is a Reader for the DcimConsoleServerPortTemplatesCreate structure.
type DcimConsoleServerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsoleServerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/console-server-port-templates/] dcim_console-server-port-templates_create", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesCreateCreated creates a DcimConsoleServerPortTemplatesCreateCreated with default headers values
func NewDcimConsoleServerPortTemplatesCreateCreated() *DcimConsoleServerPortTemplatesCreateCreated {
	return &DcimConsoleServerPortTemplatesCreateCreated{}
}

/*
DcimConsoleServerPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimConsoleServerPortTemplatesCreateCreated dcim console server port templates create created
*/
type DcimConsoleServerPortTemplatesCreateCreated struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates create created response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates create created response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates create created response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates create created response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates create created response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim console server port templates create created response
func (o *DcimConsoleServerPortTemplatesCreateCreated) Code() int {
	return 201
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
