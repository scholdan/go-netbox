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

// TenancyContactGroupsBulkPartialUpdateReader is a Reader for the TenancyContactGroupsBulkPartialUpdate structure.
type TenancyContactGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /tenancy/contact-groups/] tenancy_contact-groups_bulk_partial_update", response, response.Code())
	}
}

// NewTenancyContactGroupsBulkPartialUpdateOK creates a TenancyContactGroupsBulkPartialUpdateOK with default headers values
func NewTenancyContactGroupsBulkPartialUpdateOK() *TenancyContactGroupsBulkPartialUpdateOK {
	return &TenancyContactGroupsBulkPartialUpdateOK{}
}

/*
TenancyContactGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactGroupsBulkPartialUpdateOK tenancy contact groups bulk partial update o k
*/
type TenancyContactGroupsBulkPartialUpdateOK struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups bulk partial update o k response has a 2xx status code
func (o *TenancyContactGroupsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups bulk partial update o k response has a 3xx status code
func (o *TenancyContactGroupsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups bulk partial update o k response has a 4xx status code
func (o *TenancyContactGroupsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups bulk partial update o k response has a 5xx status code
func (o *TenancyContactGroupsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups bulk partial update o k response a status code equal to that given
func (o *TenancyContactGroupsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact groups bulk partial update o k response
func (o *TenancyContactGroupsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *TenancyContactGroupsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contact-groups/][%d] tenancyContactGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactGroupsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /tenancy/contact-groups/][%d] tenancyContactGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *TenancyContactGroupsBulkPartialUpdateOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
