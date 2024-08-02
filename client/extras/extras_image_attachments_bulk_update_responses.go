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

// ExtrasImageAttachmentsBulkUpdateReader is a Reader for the ExtrasImageAttachmentsBulkUpdate structure.
type ExtrasImageAttachmentsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/image-attachments/] extras_image-attachments_bulk_update", response, response.Code())
	}
}

// NewExtrasImageAttachmentsBulkUpdateOK creates a ExtrasImageAttachmentsBulkUpdateOK with default headers values
func NewExtrasImageAttachmentsBulkUpdateOK() *ExtrasImageAttachmentsBulkUpdateOK {
	return &ExtrasImageAttachmentsBulkUpdateOK{}
}

/*
ExtrasImageAttachmentsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsBulkUpdateOK extras image attachments bulk update o k
*/
type ExtrasImageAttachmentsBulkUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments bulk update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments bulk update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments bulk update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments bulk update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments bulk update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras image attachments bulk update o k response
func (o *ExtrasImageAttachmentsBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extrasImageAttachmentsBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/image-attachments/][%d] extrasImageAttachmentsBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
