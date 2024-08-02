// Code generated by go-swagger; DO NOT EDIT.

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyContactGroupsBulkDeleteReader is a Reader for the TenancyContactGroupsBulkDelete structure.
type TenancyContactGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /tenancy/contact-groups/] tenancy_contact-groups_bulk_delete", response, response.Code())
	}
}

// NewTenancyContactGroupsBulkDeleteNoContent creates a TenancyContactGroupsBulkDeleteNoContent with default headers values
func NewTenancyContactGroupsBulkDeleteNoContent() *TenancyContactGroupsBulkDeleteNoContent {
	return &TenancyContactGroupsBulkDeleteNoContent{}
}

/*
TenancyContactGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactGroupsBulkDeleteNoContent tenancy contact groups bulk delete no content
*/
type TenancyContactGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy contact groups bulk delete no content response has a 2xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups bulk delete no content response has a 3xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups bulk delete no content response has a 4xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups bulk delete no content response has a 5xx status code
func (o *TenancyContactGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups bulk delete no content response a status code equal to that given
func (o *TenancyContactGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the tenancy contact groups bulk delete no content response
func (o *TenancyContactGroupsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *TenancyContactGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancyContactGroupsBulkDeleteNoContent", 204)
}

func (o *TenancyContactGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-groups/][%d] tenancyContactGroupsBulkDeleteNoContent", 204)
}

func (o *TenancyContactGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
