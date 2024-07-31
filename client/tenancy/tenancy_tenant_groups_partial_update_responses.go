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

	"github.com/fbreckle/go-netbox/models"
)

// TenancyTenantGroupsPartialUpdateReader is a Reader for the TenancyTenantGroupsPartialUpdate structure.
type TenancyTenantGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /tenancy/tenant-groups/{id}/] tenancy_tenant-groups_partial_update", response, response.Code())
	}
}

// NewTenancyTenantGroupsPartialUpdateOK creates a TenancyTenantGroupsPartialUpdateOK with default headers values
func NewTenancyTenantGroupsPartialUpdateOK() *TenancyTenantGroupsPartialUpdateOK {
	return &TenancyTenantGroupsPartialUpdateOK{}
}

/*
TenancyTenantGroupsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsPartialUpdateOK tenancy tenant groups partial update o k
*/
type TenancyTenantGroupsPartialUpdateOK struct {
	Payload *models.TenantGroup
}

// IsSuccess returns true when this tenancy tenant groups partial update o k response has a 2xx status code
func (o *TenancyTenantGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenant groups partial update o k response has a 3xx status code
func (o *TenancyTenantGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups partial update o k response has a 4xx status code
func (o *TenancyTenantGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenant groups partial update o k response has a 5xx status code
func (o *TenancyTenantGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups partial update o k response a status code equal to that given
func (o *TenancyTenantGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy tenant groups partial update o k response
func (o *TenancyTenantGroupsPartialUpdateOK) Code() int {
	return 200
}

func (o *TenancyTenantGroupsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsPartialUpdateOK %s", 200, payload)
}

func (o *TenancyTenantGroupsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsPartialUpdateOK %s", 200, payload)
}

func (o *TenancyTenantGroupsPartialUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
