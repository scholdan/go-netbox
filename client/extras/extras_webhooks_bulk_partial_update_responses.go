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

// ExtrasWebhooksBulkPartialUpdateReader is a Reader for the ExtrasWebhooksBulkPartialUpdate structure.
type ExtrasWebhooksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /extras/webhooks/] extras_webhooks_bulk_partial_update", response, response.Code())
	}
}

// NewExtrasWebhooksBulkPartialUpdateOK creates a ExtrasWebhooksBulkPartialUpdateOK with default headers values
func NewExtrasWebhooksBulkPartialUpdateOK() *ExtrasWebhooksBulkPartialUpdateOK {
	return &ExtrasWebhooksBulkPartialUpdateOK{}
}

/*
ExtrasWebhooksBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkPartialUpdateOK extras webhooks bulk partial update o k
*/
type ExtrasWebhooksBulkPartialUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks bulk partial update o k response has a 2xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks bulk partial update o k response has a 3xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks bulk partial update o k response has a 4xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks bulk partial update o k response has a 5xx status code
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks bulk partial update o k response a status code equal to that given
func (o *ExtrasWebhooksBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks bulk partial update o k response
func (o *ExtrasWebhooksBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extrasWebhooksBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extrasWebhooksBulkPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
