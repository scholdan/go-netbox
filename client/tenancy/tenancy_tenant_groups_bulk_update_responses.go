// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// TenancyTenantGroupsBulkUpdateReader is a Reader for the TenancyTenantGroupsBulkUpdate structure.
type TenancyTenantGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /tenancy/tenant-groups/] tenancy_tenant-groups_bulk_update", response, response.Code())
	}
}

// NewTenancyTenantGroupsBulkUpdateOK creates a TenancyTenantGroupsBulkUpdateOK with default headers values
func NewTenancyTenantGroupsBulkUpdateOK() *TenancyTenantGroupsBulkUpdateOK {
	return &TenancyTenantGroupsBulkUpdateOK{}
}

/*
TenancyTenantGroupsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsBulkUpdateOK tenancy tenant groups bulk update o k
*/
type TenancyTenantGroupsBulkUpdateOK struct {
	Payload *models.TenantGroup
}

// IsSuccess returns true when this tenancy tenant groups bulk update o k response has a 2xx status code
func (o *TenancyTenantGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenant groups bulk update o k response has a 3xx status code
func (o *TenancyTenantGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups bulk update o k response has a 4xx status code
func (o *TenancyTenantGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenant groups bulk update o k response has a 5xx status code
func (o *TenancyTenantGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups bulk update o k response a status code equal to that given
func (o *TenancyTenantGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy tenant groups bulk update o k response
func (o *TenancyTenantGroupsBulkUpdateOK) Code() int {
	return 200
}

func (o *TenancyTenantGroupsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkUpdateOK %s", 200, payload)
}

func (o *TenancyTenantGroupsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkUpdateOK %s", 200, payload)
}

func (o *TenancyTenantGroupsBulkUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}