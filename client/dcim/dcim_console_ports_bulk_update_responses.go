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

// DcimConsolePortsBulkUpdateReader is a Reader for the DcimConsolePortsBulkUpdate structure.
type DcimConsolePortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /dcim/console-ports/] dcim_console-ports_bulk_update", response, response.Code())
	}
}

// NewDcimConsolePortsBulkUpdateOK creates a DcimConsolePortsBulkUpdateOK with default headers values
func NewDcimConsolePortsBulkUpdateOK() *DcimConsolePortsBulkUpdateOK {
	return &DcimConsolePortsBulkUpdateOK{}
}

/*
DcimConsolePortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsBulkUpdateOK dcim console ports bulk update o k
*/
type DcimConsolePortsBulkUpdateOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports bulk update o k response has a 2xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports bulk update o k response has a 3xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports bulk update o k response has a 4xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports bulk update o k response has a 5xx status code
func (o *DcimConsolePortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports bulk update o k response a status code equal to that given
func (o *DcimConsolePortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console ports bulk update o k response
func (o *DcimConsolePortsBulkUpdateOK) Code() int {
	return 200
}

func (o *DcimConsolePortsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcimConsolePortsBulkUpdateOK %s", 200, payload)
}

func (o *DcimConsolePortsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/console-ports/][%d] dcimConsolePortsBulkUpdateOK %s", 200, payload)
}

func (o *DcimConsolePortsBulkUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
