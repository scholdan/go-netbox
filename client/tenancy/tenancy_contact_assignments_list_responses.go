// Code generated by go-swagger; DO NOT EDIT.

package tenancy

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

// TenancyContactAssignmentsListReader is a Reader for the TenancyContactAssignmentsList structure.
type TenancyContactAssignmentsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /tenancy/contact-assignments/] tenancy_contact-assignments_list", response, response.Code())
	}
}

// NewTenancyContactAssignmentsListOK creates a TenancyContactAssignmentsListOK with default headers values
func NewTenancyContactAssignmentsListOK() *TenancyContactAssignmentsListOK {
	return &TenancyContactAssignmentsListOK{}
}

/*
TenancyContactAssignmentsListOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsListOK tenancy contact assignments list o k
*/
type TenancyContactAssignmentsListOK struct {
	Payload *TenancyContactAssignmentsListOKBody
}

// IsSuccess returns true when this tenancy contact assignments list o k response has a 2xx status code
func (o *TenancyContactAssignmentsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments list o k response has a 3xx status code
func (o *TenancyContactAssignmentsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments list o k response has a 4xx status code
func (o *TenancyContactAssignmentsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments list o k response has a 5xx status code
func (o *TenancyContactAssignmentsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments list o k response a status code equal to that given
func (o *TenancyContactAssignmentsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact assignments list o k response
func (o *TenancyContactAssignmentsListOK) Code() int {
	return 200
}

func (o *TenancyContactAssignmentsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/][%d] tenancyContactAssignmentsListOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/][%d] tenancyContactAssignmentsListOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsListOK) GetPayload() *TenancyContactAssignmentsListOKBody {
	return o.Payload
}

func (o *TenancyContactAssignmentsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TenancyContactAssignmentsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
TenancyContactAssignmentsListOKBody tenancy contact assignments list o k body
swagger:model TenancyContactAssignmentsListOKBody
*/
type TenancyContactAssignmentsListOKBody struct {

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
	Results []*models.ContactAssignment `json:"results"`
}

// Validate validates this tenancy contact assignments list o k body
func (o *TenancyContactAssignmentsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *TenancyContactAssignmentsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactAssignmentsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactAssignmentsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactAssignmentsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactAssignmentsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactAssignmentsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactAssignmentsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactAssignmentsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactAssignmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactAssignmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenancy contact assignments list o k body based on the context it is used
func (o *TenancyContactAssignmentsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TenancyContactAssignmentsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactAssignmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactAssignmentsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TenancyContactAssignmentsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TenancyContactAssignmentsListOKBody) UnmarshalBinary(b []byte) error {
	var res TenancyContactAssignmentsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
