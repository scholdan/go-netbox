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

// DcimConsoleServerPortTemplatesPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesPartialUpdate structure.
type DcimConsoleServerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /dcim/console-server-port-templates/{id}/] dcim_console-server-port-templates_partial_update", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateOK creates a DcimConsoleServerPortTemplatesPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateOK() *DcimConsoleServerPortTemplatesPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesPartialUpdateOK{}
}

/*
DcimConsoleServerPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesPartialUpdateOK dcim console server port templates partial update o k
*/
type DcimConsoleServerPortTemplatesPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates partial update o k response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates partial update o k response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates partial update o k response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates partial update o k response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates partial update o k response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server port templates partial update o k response
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
