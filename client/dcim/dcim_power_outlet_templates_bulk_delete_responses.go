// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerOutletTemplatesBulkDeleteReader is a Reader for the DcimPowerOutletTemplatesBulkDelete structure.
type DcimPowerOutletTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/power-outlet-templates/] dcim_power-outlet-templates_bulk_delete", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesBulkDeleteNoContent creates a DcimPowerOutletTemplatesBulkDeleteNoContent with default headers values
func NewDcimPowerOutletTemplatesBulkDeleteNoContent() *DcimPowerOutletTemplatesBulkDeleteNoContent {
	return &DcimPowerOutletTemplatesBulkDeleteNoContent{}
}

/*
DcimPowerOutletTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerOutletTemplatesBulkDeleteNoContent dcim power outlet templates bulk delete no content
*/
type DcimPowerOutletTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power outlet templates bulk delete no content response has a 2xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates bulk delete no content response has a 3xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates bulk delete no content response has a 4xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates bulk delete no content response has a 5xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates bulk delete no content response a status code equal to that given
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim power outlet templates bulk delete no content response
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
