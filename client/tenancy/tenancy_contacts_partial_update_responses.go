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

// TenancyContactsPartialUpdateReader is a Reader for the TenancyContactsPartialUpdate structure.
type TenancyContactsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /tenancy/contacts/{id}/] tenancy_contacts_partial_update", response, response.Code())
	}
}

// NewTenancyContactsPartialUpdateOK creates a TenancyContactsPartialUpdateOK with default headers values
func NewTenancyContactsPartialUpdateOK() *TenancyContactsPartialUpdateOK {
	return &TenancyContactsPartialUpdateOK{}
}

/*
TenancyContactsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactsPartialUpdateOK tenancy contacts partial update o k
*/
type TenancyContactsPartialUpdateOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts partial update o k response has a 2xx status code
func (o *TenancyContactsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts partial update o k response has a 3xx status code
func (o *TenancyContactsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts partial update o k response has a 4xx status code
func (o *TenancyContactsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts partial update o k response has a 5xx status code
func (o *TenancyContactsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts partial update o k response a status code equal to that given
func (o *TenancyContactsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contacts partial update o k response
func (o *TenancyContactsPartialUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancyContactsPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contacts/{id}/][%d] tenancyContactsPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactsPartialUpdateOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
