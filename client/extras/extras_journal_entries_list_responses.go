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

	"github.com/fbreckle/go-netbox/models"
)

// ExtrasJournalEntriesListReader is a Reader for the ExtrasJournalEntriesList structure.
type ExtrasJournalEntriesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /extras/journal-entries/] extras_journal-entries_list", response, response.Code())
	}
}

// NewExtrasJournalEntriesListOK creates a ExtrasJournalEntriesListOK with default headers values
func NewExtrasJournalEntriesListOK() *ExtrasJournalEntriesListOK {
	return &ExtrasJournalEntriesListOK{}
}

/*
ExtrasJournalEntriesListOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesListOK extras journal entries list o k
*/
type ExtrasJournalEntriesListOK struct {
	Payload *ExtrasJournalEntriesListOKBody
}

// IsSuccess returns true when this extras journal entries list o k response has a 2xx status code
func (o *ExtrasJournalEntriesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries list o k response has a 3xx status code
func (o *ExtrasJournalEntriesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries list o k response has a 4xx status code
func (o *ExtrasJournalEntriesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries list o k response has a 5xx status code
func (o *ExtrasJournalEntriesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries list o k response a status code equal to that given
func (o *ExtrasJournalEntriesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras journal entries list o k response
func (o *ExtrasJournalEntriesListOK) Code() int {
	return 200
}

func (o *ExtrasJournalEntriesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/journal-entries/][%d] extrasJournalEntriesListOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /extras/journal-entries/][%d] extrasJournalEntriesListOK %s", 200, payload)
}

func (o *ExtrasJournalEntriesListOK) GetPayload() *ExtrasJournalEntriesListOKBody {
	return o.Payload
}

func (o *ExtrasJournalEntriesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasJournalEntriesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExtrasJournalEntriesListOKBody extras journal entries list o k body
swagger:model ExtrasJournalEntriesListOKBody
*/
type ExtrasJournalEntriesListOKBody struct {

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
	Results []*models.JournalEntry `json:"results"`
}

// Validate validates this extras journal entries list o k body
func (o *ExtrasJournalEntriesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExtrasJournalEntriesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasJournalEntriesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasJournalEntriesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasJournalEntriesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasJournalEntriesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasJournalEntriesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasJournalEntriesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasJournalEntriesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasJournalEntriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasJournalEntriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras journal entries list o k body based on the context it is used
func (o *ExtrasJournalEntriesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasJournalEntriesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasJournalEntriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extrasJournalEntriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasJournalEntriesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasJournalEntriesListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasJournalEntriesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
