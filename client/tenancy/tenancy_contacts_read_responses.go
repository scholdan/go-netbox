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

// TenancyContactsReadReader is a Reader for the TenancyContactsRead structure.
type TenancyContactsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /tenancy/contacts/{id}/] tenancy_contacts_read", response, response.Code())
	}
}

// NewTenancyContactsReadOK creates a TenancyContactsReadOK with default headers values
func NewTenancyContactsReadOK() *TenancyContactsReadOK {
	return &TenancyContactsReadOK{}
}

/*
TenancyContactsReadOK describes a response with status code 200, with default header values.

TenancyContactsReadOK tenancy contacts read o k
*/
type TenancyContactsReadOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts read o k response has a 2xx status code
func (o *TenancyContactsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts read o k response has a 3xx status code
func (o *TenancyContactsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts read o k response has a 4xx status code
func (o *TenancyContactsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts read o k response has a 5xx status code
func (o *TenancyContactsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts read o k response a status code equal to that given
func (o *TenancyContactsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contacts read o k response
func (o *TenancyContactsReadOK) Code() int {
	return 200
}

func (o *TenancyContactsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancyContactsReadOK %s", 200, payload)
}

func (o *TenancyContactsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contacts/{id}/][%d] tenancyContactsReadOK %s", 200, payload)
}

func (o *TenancyContactsReadOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
