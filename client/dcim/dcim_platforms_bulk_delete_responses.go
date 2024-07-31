// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPlatformsBulkDeleteReader is a Reader for the DcimPlatformsBulkDelete structure.
type DcimPlatformsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPlatformsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /dcim/platforms/] dcim_platforms_bulk_delete", response, response.Code())
	}
}

// NewDcimPlatformsBulkDeleteNoContent creates a DcimPlatformsBulkDeleteNoContent with default headers values
func NewDcimPlatformsBulkDeleteNoContent() *DcimPlatformsBulkDeleteNoContent {
	return &DcimPlatformsBulkDeleteNoContent{}
}

/*
DcimPlatformsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPlatformsBulkDeleteNoContent dcim platforms bulk delete no content
*/
type DcimPlatformsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim platforms bulk delete no content response has a 2xx status code
func (o *DcimPlatformsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms bulk delete no content response has a 3xx status code
func (o *DcimPlatformsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms bulk delete no content response has a 4xx status code
func (o *DcimPlatformsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms bulk delete no content response has a 5xx status code
func (o *DcimPlatformsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms bulk delete no content response a status code equal to that given
func (o *DcimPlatformsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim platforms bulk delete no content response
func (o *DcimPlatformsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimPlatformsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/][%d] dcimPlatformsBulkDeleteNoContent", 204)
}

func (o *DcimPlatformsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/][%d] dcimPlatformsBulkDeleteNoContent", 204)
}

func (o *DcimPlatformsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
