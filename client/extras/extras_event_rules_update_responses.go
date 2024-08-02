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

// ExtrasEventRulesUpdateReader is a Reader for the ExtrasEventRulesUpdate structure.
type ExtrasEventRulesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasEventRulesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /extras/event-rules/{id}/] extras_event_rules_update", response, response.Code())
	}
}

// NewExtrasEventRulesUpdateOK creates a ExtrasEventRulesUpdateOK with default headers values
func NewExtrasEventRulesUpdateOK() *ExtrasEventRulesUpdateOK {
	return &ExtrasEventRulesUpdateOK{}
}

/*
ExtrasEventRulesUpdateOK describes a response with status code 200, with default header values.

ExtrasEventRulesUpdateOK extras event rules update o k
*/
type ExtrasEventRulesUpdateOK struct {
	Payload *models.EventRule
}

// IsSuccess returns true when this extras event rules update o k response has a 2xx status code
func (o *ExtrasEventRulesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules update o k response has a 3xx status code
func (o *ExtrasEventRulesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules update o k response has a 4xx status code
func (o *ExtrasEventRulesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules update o k response has a 5xx status code
func (o *ExtrasEventRulesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules update o k response a status code equal to that given
func (o *ExtrasEventRulesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras event rules update o k response
func (o *ExtrasEventRulesUpdateOK) Code() int {
	return 200
}

func (o *ExtrasEventRulesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/event-rules/{id}/][%d] extrasEventRulesUpdateOK %s", 200, payload)
}

func (o *ExtrasEventRulesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /extras/event-rules/{id}/][%d] extrasEventRulesUpdateOK %s", 200, payload)
}

func (o *ExtrasEventRulesUpdateOK) GetPayload() *models.EventRule {
	return o.Payload
}

func (o *ExtrasEventRulesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}