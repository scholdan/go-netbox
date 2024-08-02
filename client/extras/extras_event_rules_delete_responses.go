// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasEventRulesDeleteReader is a Reader for the ExtrasEventRulesDelete structure.
type ExtrasEventRulesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasEventRulesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /extras/event-rules/{id}/] extras_event_rules_delete", response, response.Code())
	}
}

// NewExtrasEventRulesDeleteNoContent creates a ExtrasEventRulesDeleteNoContent with default headers values
func NewExtrasEventRulesDeleteNoContent() *ExtrasEventRulesDeleteNoContent {
	return &ExtrasEventRulesDeleteNoContent{}
}

/*
ExtrasEventRulesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasEventRulesDeleteNoContent extras event rules delete no content
*/
type ExtrasEventRulesDeleteNoContent struct {
}

// IsSuccess returns true when this extras event rules delete no content response has a 2xx status code
func (o *ExtrasEventRulesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules delete no content response has a 3xx status code
func (o *ExtrasEventRulesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules delete no content response has a 4xx status code
func (o *ExtrasEventRulesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules delete no content response has a 5xx status code
func (o *ExtrasEventRulesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules delete no content response a status code equal to that given
func (o *ExtrasEventRulesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the extras event rules delete no content response
func (o *ExtrasEventRulesDeleteNoContent) Code() int {
	return 204
}

func (o *ExtrasEventRulesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/event-rules/{id}/][%d] extrasEventRulesDeleteNoContent", 204)
}

func (o *ExtrasEventRulesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/event-rules/{id}/][%d] extrasEventRulesDeleteNoContent", 204)
}

func (o *ExtrasEventRulesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}