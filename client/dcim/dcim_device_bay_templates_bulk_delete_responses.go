// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBayTemplatesBulkDeleteReader is a Reader for the DcimDeviceBayTemplatesBulkDelete structure.
type DcimDeviceBayTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBayTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/device-bay-templates/] dcim_device-bay-templates_bulk_delete", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesBulkDeleteNoContent creates a DcimDeviceBayTemplatesBulkDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesBulkDeleteNoContent() *DcimDeviceBayTemplatesBulkDeleteNoContent {
	return &DcimDeviceBayTemplatesBulkDeleteNoContent{}
}

/*
DcimDeviceBayTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBayTemplatesBulkDeleteNoContent dcim device bay templates bulk delete no content
*/
type DcimDeviceBayTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device bay templates bulk delete no content response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates bulk delete no content response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates bulk delete no content response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates bulk delete no content response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates bulk delete no content response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim device bay templates bulk delete no content response
func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkDeleteNoContent", 204)
}

func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}