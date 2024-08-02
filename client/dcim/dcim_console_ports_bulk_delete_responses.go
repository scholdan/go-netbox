// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsolePortsBulkDeleteReader is a Reader for the DcimConsolePortsBulkDelete structure.
type DcimConsolePortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/console-ports/] dcim_console-ports_bulk_delete", response, response.Code())
	}
}

// NewDcimConsolePortsBulkDeleteNoContent creates a DcimConsolePortsBulkDeleteNoContent with default headers values
func NewDcimConsolePortsBulkDeleteNoContent() *DcimConsolePortsBulkDeleteNoContent {
	return &DcimConsolePortsBulkDeleteNoContent{}
}

/*
DcimConsolePortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortsBulkDeleteNoContent dcim console ports bulk delete no content
*/
type DcimConsolePortsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console ports bulk delete no content response has a 2xx status code
func (o *DcimConsolePortsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports bulk delete no content response has a 3xx status code
func (o *DcimConsolePortsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports bulk delete no content response has a 4xx status code
func (o *DcimConsolePortsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports bulk delete no content response has a 5xx status code
func (o *DcimConsolePortsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports bulk delete no content response a status code equal to that given
func (o *DcimConsolePortsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim console ports bulk delete no content response
func (o *DcimConsolePortsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimConsolePortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/][%d] dcimConsolePortsBulkDeleteNoContent", 204)
}

func (o *DcimConsolePortsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/][%d] dcimConsolePortsBulkDeleteNoContent", 204)
}

func (o *DcimConsolePortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}