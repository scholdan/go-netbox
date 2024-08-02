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

// ExtrasTagsBulkPartialUpdateReader is a Reader for the ExtrasTagsBulkPartialUpdate structure.
type ExtrasTagsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/tags/] extras_tags_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasTagsBulkPartialUpdateOK creates a ExtrasTagsBulkPartialUpdateOK with default headers values
func NewExtrasTagsBulkPartialUpdateOK() *ExtrasTagsBulkPartialUpdateOK {
	return &ExtrasTagsBulkPartialUpdateOK{}
}

/*
ExtrasTagsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsBulkPartialUpdateOK extras tags bulk partial update o k
*/
type ExtrasTagsBulkPartialUpdateOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this extras tags bulk partial update o k response has a 2xx status code
func (o *ExtrasTagsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags bulk partial update o k response has a 3xx status code
func (o *ExtrasTagsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags bulk partial update o k response has a 4xx status code
func (o *ExtrasTagsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags bulk partial update o k response has a 5xx status code
func (o *ExtrasTagsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags bulk partial update o k response a status code equal to that given
func (o *ExtrasTagsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras tags bulk partial update o k response
func (o *ExtrasTagsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasTagsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/tags/][%d] extrasTagsBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasTagsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/tags/][%d] extrasTagsBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasTagsBulkPartialUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}