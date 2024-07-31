// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// CircuitsCircuitTypesListReader is a Reader for the CircuitsCircuitTypesList structure.
type CircuitsCircuitTypesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /circuits/circuit-types/] circuits_circuit-types_list", response, response.Code())
	}
}

// NewCircuitsCircuitTypesListOK creates a CircuitsCircuitTypesListOK with default headers values
func NewCircuitsCircuitTypesListOK() *CircuitsCircuitTypesListOK {
	return &CircuitsCircuitTypesListOK{}
}

/*
CircuitsCircuitTypesListOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesListOK circuits circuit types list o k
*/
type CircuitsCircuitTypesListOK struct {
	Payload *CircuitsCircuitTypesListOKBody
}

// IsSuccess returns true when this circuits circuit types list o k response has a 2xx status code
func (o *CircuitsCircuitTypesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types list o k response has a 3xx status code
func (o *CircuitsCircuitTypesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types list o k response has a 4xx status code
func (o *CircuitsCircuitTypesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types list o k response has a 5xx status code
func (o *CircuitsCircuitTypesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types list o k response a status code equal to that given
func (o *CircuitsCircuitTypesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the circuits circuit types list o k response
func (o *CircuitsCircuitTypesListOK) Code() int {
	return 200
}

func (o *CircuitsCircuitTypesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-types/][%d] circuitsCircuitTypesListOK %s", 200, payload)
}

func (o *CircuitsCircuitTypesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /circuits/circuit-types/][%d] circuitsCircuitTypesListOK %s", 200, payload)
}

func (o *CircuitsCircuitTypesListOK) GetPayload() *CircuitsCircuitTypesListOKBody {
	return o.Payload
}

func (o *CircuitsCircuitTypesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CircuitsCircuitTypesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CircuitsCircuitTypesListOKBody circuits circuit types list o k body
swagger:model CircuitsCircuitTypesListOKBody
*/
type CircuitsCircuitTypesListOKBody struct {

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
	Results []*models.CircuitType `json:"results"`
}

// Validate validates this circuits circuit types list o k body
func (o *CircuitsCircuitTypesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *CircuitsCircuitTypesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("circuitsCircuitTypesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTypesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("circuitsCircuitTypesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTypesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("circuitsCircuitTypesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CircuitsCircuitTypesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("circuitsCircuitTypesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("circuitsCircuitTypesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("circuitsCircuitTypesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this circuits circuit types list o k body based on the context it is used
func (o *CircuitsCircuitTypesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CircuitsCircuitTypesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("circuitsCircuitTypesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("circuitsCircuitTypesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CircuitsCircuitTypesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CircuitsCircuitTypesListOKBody) UnmarshalBinary(b []byte) error {
	var res CircuitsCircuitTypesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
