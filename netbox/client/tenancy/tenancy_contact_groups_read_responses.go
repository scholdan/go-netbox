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

// TenancyContactGroupsReadReader is a Reader for the TenancyContactGroupsRead structure.
type TenancyContactGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /tenancy/contact-groups/{id}/] tenancy_contact-groups_read", response, response.Code())
	}
}

// NewTenancyContactGroupsReadOK creates a TenancyContactGroupsReadOK with default headers values
func NewTenancyContactGroupsReadOK() *TenancyContactGroupsReadOK {
	return &TenancyContactGroupsReadOK{}
}

/*
TenancyContactGroupsReadOK describes a response with status code 200, with default header values.

TenancyContactGroupsReadOK tenancy contact groups read o k
*/
type TenancyContactGroupsReadOK struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups read o k response has a 2xx status code
func (o *TenancyContactGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups read o k response has a 3xx status code
func (o *TenancyContactGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups read o k response has a 4xx status code
func (o *TenancyContactGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups read o k response has a 5xx status code
func (o *TenancyContactGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups read o k response a status code equal to that given
func (o *TenancyContactGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact groups read o k response
func (o *TenancyContactGroupsReadOK) Code() int {
	return 200
}

func (o *TenancyContactGroupsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadOK %s", 200, payload)
}

func (o *TenancyContactGroupsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadOK %s", 200, payload)
}

func (o *TenancyContactGroupsReadOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
