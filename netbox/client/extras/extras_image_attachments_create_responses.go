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

// ExtrasImageAttachmentsCreateReader is a Reader for the ExtrasImageAttachmentsCreate structure.
type ExtrasImageAttachmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasImageAttachmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /extras/image-attachments/] extras_image-attachments_create", response, response.Code())
	}
}

// NewExtrasImageAttachmentsCreateCreated creates a ExtrasImageAttachmentsCreateCreated with default headers values
func NewExtrasImageAttachmentsCreateCreated() *ExtrasImageAttachmentsCreateCreated {
	return &ExtrasImageAttachmentsCreateCreated{}
}

/*
ExtrasImageAttachmentsCreateCreated describes a response with status code 201, with default header values.

ExtrasImageAttachmentsCreateCreated extras image attachments create created
*/
type ExtrasImageAttachmentsCreateCreated struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments create created response has a 2xx status code
func (o *ExtrasImageAttachmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments create created response has a 3xx status code
func (o *ExtrasImageAttachmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments create created response has a 4xx status code
func (o *ExtrasImageAttachmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments create created response has a 5xx status code
func (o *ExtrasImageAttachmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments create created response a status code equal to that given
func (o *ExtrasImageAttachmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the extras image attachments create created response
func (o *ExtrasImageAttachmentsCreateCreated) Code() int {
	return 201
}

func (o *ExtrasImageAttachmentsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/image-attachments/][%d] extrasImageAttachmentsCreateCreated %s", 201, payload)
}

func (o *ExtrasImageAttachmentsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/image-attachments/][%d] extrasImageAttachmentsCreateCreated %s", 201, payload)
}

func (o *ExtrasImageAttachmentsCreateCreated) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
