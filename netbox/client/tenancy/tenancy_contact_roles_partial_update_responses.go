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

// TenancyContactRolesPartialUpdateReader is a Reader for the TenancyContactRolesPartialUpdate structure.
type TenancyContactRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /tenancy/contact-roles/{id}/] tenancy_contact-roles_partial_update", response, response.Code())
	}
}

// NewTenancyContactRolesPartialUpdateOK creates a TenancyContactRolesPartialUpdateOK with default headers values
func NewTenancyContactRolesPartialUpdateOK() *TenancyContactRolesPartialUpdateOK {
	return &TenancyContactRolesPartialUpdateOK{}
}

/*
TenancyContactRolesPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactRolesPartialUpdateOK tenancy contact roles partial update o k
*/
type TenancyContactRolesPartialUpdateOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles partial update o k response has a 2xx status code
func (o *TenancyContactRolesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles partial update o k response has a 3xx status code
func (o *TenancyContactRolesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles partial update o k response has a 4xx status code
func (o *TenancyContactRolesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles partial update o k response has a 5xx status code
func (o *TenancyContactRolesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles partial update o k response a status code equal to that given
func (o *TenancyContactRolesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact roles partial update o k response
func (o *TenancyContactRolesPartialUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactRolesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/{id}/][%d] tenancyContactRolesPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactRolesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contact-roles/{id}/][%d] tenancyContactRolesPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactRolesPartialUpdateOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
