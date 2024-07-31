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

// ExtrasWebhooksBulkUpdateReader is a Reader for the ExtrasWebhooksBulkUpdate structure.
type ExtrasWebhooksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/webhooks/] extras_webhooks_bulk_update", response, response.Code())
	}
}

// NewExtrasWebhooksBulkUpdateOK creates a ExtrasWebhooksBulkUpdateOK with default headers values
func NewExtrasWebhooksBulkUpdateOK() *ExtrasWebhooksBulkUpdateOK {
	return &ExtrasWebhooksBulkUpdateOK{}
}

/*
ExtrasWebhooksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkUpdateOK extras webhooks bulk update o k
*/
type ExtrasWebhooksBulkUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks bulk update o k response has a 2xx status code
func (o *ExtrasWebhooksBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks bulk update o k response has a 3xx status code
func (o *ExtrasWebhooksBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks bulk update o k response has a 4xx status code
func (o *ExtrasWebhooksBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks bulk update o k response has a 5xx status code
func (o *ExtrasWebhooksBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks bulk update o k response a status code equal to that given
func (o *ExtrasWebhooksBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks bulk update o k response
func (o *ExtrasWebhooksBulkUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/webhooks/][%d] extrasWebhooksBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/webhooks/][%d] extrasWebhooksBulkUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksBulkUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
