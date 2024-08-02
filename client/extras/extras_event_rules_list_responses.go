// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/scholdan/go-netbox/models"
)

// ExtrasEventRulesListReader is a Reader for the ExtrasEventRulesList structure.
type ExtrasEventRulesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasEventRulesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/event-rules/] extras_event_rules_list", response, response.Code())
	}
}

// NewExtrasEventRulesListOK creates a ExtrasEventRulesListOK with default headers values
func NewExtrasEventRulesListOK() *ExtrasEventRulesListOK {
	return &ExtrasEventRulesListOK{}
}

/*
ExtrasEventRulesListOK describes a response with status code 200, with default header values.

ExtrasEventRulesListOK extras event rules list o k
*/
type ExtrasEventRulesListOK struct {
	Payload *ExtrasEventRulesListOKBody
}

// IsSuccess returns true when this extras event rules list o k response has a 2xx status code
func (o *ExtrasEventRulesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules list o k response has a 3xx status code
func (o *ExtrasEventRulesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules list o k response has a 4xx status code
func (o *ExtrasEventRulesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules list o k response has a 5xx status code
func (o *ExtrasEventRulesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules list o k response a status code equal to that given
func (o *ExtrasEventRulesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras event rules list o k response
func (o *ExtrasEventRulesListOK) Code() int {
	return 200
}

func (o *ExtrasEventRulesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/event-rules/][%d] extrasEventRulesListOK %s", 200, payload)
}

func (o *ExtrasEventRulesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/event-rules/][%d] extrasEventRulesListOK %s", 200, payload)
}

func (o *ExtrasEventRulesListOK) GetPayload() *ExtrasEventRulesListOKBody {
	return o.Payload
}

func (o *ExtrasEventRulesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasEventRulesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasEventRulesListOKBody extras event rules list o k body
swagger:model ExtrasEventRulesListOKBody
*/
type ExtrasEventRulesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.EventRule `json:"results"`
}

// Validate validates this extras event rules list o k body
func (o *ExtrasEventRulesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasEventRulesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasEventRulesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasEventRulesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasEventRulesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasEventRulesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasEventRulesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasEventRulesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasEventRulesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasEventRulesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasEventRulesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras event rules list o k body based on the context it is used
func (o *ExtrasEventRulesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasEventRulesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasEventRulesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasEventRulesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasEventRulesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasEventRulesListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasEventRulesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
