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

// DcimConsolePortTemplatesCreateReader is a Reader for the DcimConsolePortTemplatesCreate structure.
type DcimConsolePortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsolePortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /dcim/console-port-templates/] dcim_console-port-templates_create", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesCreateCreated creates a DcimConsolePortTemplatesCreateCreated with default headers values
func NewDcimConsolePortTemplatesCreateCreated() *DcimConsolePortTemplatesCreateCreated {
	return &DcimConsolePortTemplatesCreateCreated{}
}

/*
DcimConsolePortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimConsolePortTemplatesCreateCreated dcim console port templates create created
*/
type DcimConsolePortTemplatesCreateCreated struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates create created response has a 2xx status code
func (o *DcimConsolePortTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates create created response has a 3xx status code
func (o *DcimConsolePortTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates create created response has a 4xx status code
func (o *DcimConsolePortTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates create created response has a 5xx status code
func (o *DcimConsolePortTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates create created response a status code equal to that given
func (o *DcimConsolePortTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim console port templates create created response
func (o *DcimConsolePortTemplatesCreateCreated) Code() int {
	return 201
}

func (o *DcimConsolePortTemplatesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/console-port-templates/][%d] dcimConsolePortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimConsolePortTemplatesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/console-port-templates/][%d] dcimConsolePortTemplatesCreateCreated %s", 201, payload)
}

func (o *DcimConsolePortTemplatesCreateCreated) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
