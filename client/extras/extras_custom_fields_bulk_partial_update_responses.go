// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// ExtrasCustomFieldsBulkPartialUpdateReader is a Reader for the ExtrasCustomFieldsBulkPartialUpdate structure.
type ExtrasCustomFieldsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/custom-fields/] extras_custom-fields_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasCustomFieldsBulkPartialUpdateOK creates a ExtrasCustomFieldsBulkPartialUpdateOK with default headers values
func NewExtrasCustomFieldsBulkPartialUpdateOK() *ExtrasCustomFieldsBulkPartialUpdateOK {
	return &ExtrasCustomFieldsBulkPartialUpdateOK{}
}

/*
ExtrasCustomFieldsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsBulkPartialUpdateOK extras custom fields bulk partial update o k
*/
type ExtrasCustomFieldsBulkPartialUpdateOK struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this extras custom fields bulk partial update o k response has a 2xx status code
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields bulk partial update o k response has a 3xx status code
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields bulk partial update o k response has a 4xx status code
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields bulk partial update o k response has a 5xx status code
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields bulk partial update o k response a status code equal to that given
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom fields bulk partial update o k response
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/custom-fields/][%d] extrasCustomFieldsBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/custom-fields/][%d] extrasCustomFieldsBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}