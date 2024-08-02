// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsolePortTemplatesBulkDeleteReader is a Reader for the DcimConsolePortTemplatesBulkDelete structure.
type DcimConsolePortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/console-port-templates/] dcim_console-port-templates_bulk_delete", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesBulkDeleteNoContent creates a DcimConsolePortTemplatesBulkDeleteNoContent with default headers values
func NewDcimConsolePortTemplatesBulkDeleteNoContent() *DcimConsolePortTemplatesBulkDeleteNoContent {
	return &DcimConsolePortTemplatesBulkDeleteNoContent{}
}

/*
DcimConsolePortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortTemplatesBulkDeleteNoContent dcim console port templates bulk delete no content
*/
type DcimConsolePortTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console port templates bulk delete no content response has a 2xx status code
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates bulk delete no content response has a 3xx status code
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates bulk delete no content response has a 4xx status code
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates bulk delete no content response has a 5xx status code
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates bulk delete no content response a status code equal to that given
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim console port templates bulk delete no content response
func (o *DcimConsolePortTemplatesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimConsolePortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimConsolePortTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimConsolePortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}