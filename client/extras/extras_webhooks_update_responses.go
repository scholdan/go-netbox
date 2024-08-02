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

// ExtrasWebhooksUpdateReader is a Reader for the ExtrasWebhooksUpdate structure.
type ExtrasWebhooksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/webhooks/{id}/] extras_webhooks_update", response, response.Code())
	}
}

// NewExtrasWebhooksUpdateOK creates a ExtrasWebhooksUpdateOK with default headers values
func NewExtrasWebhooksUpdateOK() *ExtrasWebhooksUpdateOK {
	return &ExtrasWebhooksUpdateOK{}
}

/*
ExtrasWebhooksUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksUpdateOK extras webhooks update o k
*/
type ExtrasWebhooksUpdateOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this extras webhooks update o k response has a 2xx status code
func (o *ExtrasWebhooksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks update o k response has a 3xx status code
func (o *ExtrasWebhooksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks update o k response has a 4xx status code
func (o *ExtrasWebhooksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks update o k response has a 5xx status code
func (o *ExtrasWebhooksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks update o k response a status code equal to that given
func (o *ExtrasWebhooksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras webhooks update o k response
func (o *ExtrasWebhooksUpdateOK) Code() int {
	return 200
}

func (o *ExtrasWebhooksUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extrasWebhooksUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extrasWebhooksUpdateOK %s", 200, payload)
}

func (o *ExtrasWebhooksUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}