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

	"github.com/fbreckle/go-netbox/models"
)

// ExtrasImageAttachmentsReadReader is a Reader for the ExtrasImageAttachmentsRead structure.
type ExtrasImageAttachmentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/image-attachments/{id}/] extras_image-attachments_read", response, response.Code())
	}
}

// NewExtrasImageAttachmentsReadOK creates a ExtrasImageAttachmentsReadOK with default headers values
func NewExtrasImageAttachmentsReadOK() *ExtrasImageAttachmentsReadOK {
	return &ExtrasImageAttachmentsReadOK{}
}

/*
ExtrasImageAttachmentsReadOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsReadOK extras image attachments read o k
*/
type ExtrasImageAttachmentsReadOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments read o k response has a 2xx status code
func (o *ExtrasImageAttachmentsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments read o k response has a 3xx status code
func (o *ExtrasImageAttachmentsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments read o k response has a 4xx status code
func (o *ExtrasImageAttachmentsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments read o k response has a 5xx status code
func (o *ExtrasImageAttachmentsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments read o k response a status code equal to that given
func (o *ExtrasImageAttachmentsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras image attachments read o k response
func (o *ExtrasImageAttachmentsReadOK) Code() int {
	return 200
}

func (o *ExtrasImageAttachmentsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/image-attachments/{id}/][%d] extrasImageAttachmentsReadOK %s", 200, payload)
}

func (o *ExtrasImageAttachmentsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/image-attachments/{id}/][%d] extrasImageAttachmentsReadOK %s", 200, payload)
}

func (o *ExtrasImageAttachmentsReadOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
