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

// ExtrasCustomLinksBulkPartialUpdateReader is a Reader for the ExtrasCustomLinksBulkPartialUpdate structure.
type ExtrasCustomLinksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/custom-links/] extras_custom-links_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasCustomLinksBulkPartialUpdateOK creates a ExtrasCustomLinksBulkPartialUpdateOK with default headers values
func NewExtrasCustomLinksBulkPartialUpdateOK() *ExtrasCustomLinksBulkPartialUpdateOK {
	return &ExtrasCustomLinksBulkPartialUpdateOK{}
}

/*
ExtrasCustomLinksBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksBulkPartialUpdateOK extras custom links bulk partial update o k
*/
type ExtrasCustomLinksBulkPartialUpdateOK struct {
	Payload *models.CustomLink
}

// IsSuccess returns true when this extras custom links bulk partial update o k response has a 2xx status code
func (o *ExtrasCustomLinksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom links bulk partial update o k response has a 3xx status code
func (o *ExtrasCustomLinksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom links bulk partial update o k response has a 4xx status code
func (o *ExtrasCustomLinksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom links bulk partial update o k response has a 5xx status code
func (o *ExtrasCustomLinksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom links bulk partial update o k response a status code equal to that given
func (o *ExtrasCustomLinksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom links bulk partial update o k response
func (o *ExtrasCustomLinksBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/custom-links/][%d] extrasCustomLinksBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/custom-links/][%d] extrasCustomLinksBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}