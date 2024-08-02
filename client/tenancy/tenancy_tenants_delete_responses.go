// Code generated by go-swagger; DO NOT EDIT.

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantsDeleteReader is a Reader for the TenancyTenantsDelete structure.
type TenancyTenantsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /tenancy/tenants/{id}/] tenancy_tenants_delete", response, response.Code())
	}
}

// NewTenancyTenantsDeleteNoContent creates a TenancyTenantsDeleteNoContent with default headers values
func NewTenancyTenantsDeleteNoContent() *TenancyTenantsDeleteNoContent {
	return &TenancyTenantsDeleteNoContent{}
}

/*
TenancyTenantsDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantsDeleteNoContent tenancy tenants delete no content
*/
type TenancyTenantsDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy tenants delete no content response has a 2xx status code
func (o *TenancyTenantsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants delete no content response has a 3xx status code
func (o *TenancyTenantsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants delete no content response has a 4xx status code
func (o *TenancyTenantsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants delete no content response has a 5xx status code
func (o *TenancyTenantsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants delete no content response a status code equal to that given
func (o *TenancyTenantsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the tenancy tenants delete no content response
func (o *TenancyTenantsDeleteNoContent) Code() int {
	return 204
}

func (o *TenancyTenantsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/{id}/][%d] tenancyTenantsDeleteNoContent", 204)
}

func (o *TenancyTenantsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/{id}/][%d] tenancyTenantsDeleteNoContent", 204)
}

func (o *TenancyTenantsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}