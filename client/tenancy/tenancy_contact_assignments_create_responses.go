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

// TenancyContactAssignmentsCreateReader is a Reader for the TenancyContactAssignmentsCreate structure.
type TenancyContactAssignmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyContactAssignmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /tenancy/contact-assignments/] tenancy_contact-assignments_create", response, response.Code())
	}
}

// NewTenancyContactAssignmentsCreateCreated creates a TenancyContactAssignmentsCreateCreated with default headers values
func NewTenancyContactAssignmentsCreateCreated() *TenancyContactAssignmentsCreateCreated {
	return &TenancyContactAssignmentsCreateCreated{}
}

/*
TenancyContactAssignmentsCreateCreated describes a response with status code 201, with default header values.

TenancyContactAssignmentsCreateCreated tenancy contact assignments create created
*/
type TenancyContactAssignmentsCreateCreated struct {
	Payload *models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments create created response has a 2xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments create created response has a 3xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments create created response has a 4xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments create created response has a 5xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments create created response a status code equal to that given
func (o *TenancyContactAssignmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy contact assignments create created response
func (o *TenancyContactAssignmentsCreateCreated) Code() int {
	return 201
}

func (o *TenancyContactAssignmentsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancyContactAssignmentsCreateCreated %s", 201, payload)
}

func (o *TenancyContactAssignmentsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancyContactAssignmentsCreateCreated %s", 201, payload)
}

func (o *TenancyContactAssignmentsCreateCreated) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
