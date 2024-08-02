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

// ExtrasCustomLinksBulkUpdateReader is a Reader for the ExtrasCustomLinksBulkUpdate structure.
type ExtrasCustomLinksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomLinksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/custom-links/] extras_custom-links_bulk_update", response, response.Code())
	}
}

// NewExtrasCustomLinksBulkUpdateOK creates a ExtrasCustomLinksBulkUpdateOK with default headers values
func NewExtrasCustomLinksBulkUpdateOK() *ExtrasCustomLinksBulkUpdateOK {
	return &ExtrasCustomLinksBulkUpdateOK{}
}

/*
ExtrasCustomLinksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomLinksBulkUpdateOK extras custom links bulk update o k
*/
type ExtrasCustomLinksBulkUpdateOK struct {
	Payload *models.CustomLink
}

// IsSuccess returns true when this extras custom links bulk update o k response has a 2xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom links bulk update o k response has a 3xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom links bulk update o k response has a 4xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom links bulk update o k response has a 5xx status code
func (o *ExtrasCustomLinksBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom links bulk update o k response a status code equal to that given
func (o *ExtrasCustomLinksBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom links bulk update o k response
func (o *ExtrasCustomLinksBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasCustomLinksBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extrasCustomLinksBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomLinksBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/custom-links/][%d] extrasCustomLinksBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasCustomLinksBulkUpdateOK) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
