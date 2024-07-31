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

// ExtrasEventRulesReadReader is a Reader for the ExtrasEventRulesRead structure.
type ExtrasEventRulesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasEventRulesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/event-rules/{id}/] extras_event_rules_read", response, response.Code())
	}
}

// NewExtrasEventRulesReadOK creates a ExtrasEventRulesReadOK with default headers values
func NewExtrasEventRulesReadOK() *ExtrasEventRulesReadOK {
	return &ExtrasEventRulesReadOK{}
}

/*
ExtrasEventRulesReadOK describes a response with status code 200, with default header values.

ExtrasEventRulesReadOK extras event rules read o k
*/
type ExtrasEventRulesReadOK struct {
	Payload *models.EventRule
}

// IsSuccess returns true when this extras event rules read o k response has a 2xx status code
func (o *ExtrasEventRulesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules read o k response has a 3xx status code
func (o *ExtrasEventRulesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules read o k response has a 4xx status code
func (o *ExtrasEventRulesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules read o k response has a 5xx status code
func (o *ExtrasEventRulesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules read o k response a status code equal to that given
func (o *ExtrasEventRulesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras event rules read o k response
func (o *ExtrasEventRulesReadOK) Code() int {
	return 200
}

func (o *ExtrasEventRulesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extrasEventRulesReadOK %s", 200, payload)
}

func (o *ExtrasEventRulesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/event-rules/{id}/][%d] extrasEventRulesReadOK %s", 200, payload)
}

func (o *ExtrasEventRulesReadOK) GetPayload() *models.EventRule {
	return o.Payload
}

func (o *ExtrasEventRulesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
