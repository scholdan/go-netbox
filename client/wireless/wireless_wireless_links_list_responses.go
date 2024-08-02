// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// WirelessWirelessLinksListReader is a Reader for the WirelessWirelessLinksList structure.
type WirelessWirelessLinksListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLinksListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /wireless/wireless-links/] wireless_wireless-links_list", response, response.Code())
	}
}

// NewWirelessWirelessLinksListOK creates a WirelessWirelessLinksListOK with default headers values
func NewWirelessWirelessLinksListOK() *WirelessWirelessLinksListOK {
	return &WirelessWirelessLinksListOK{}
}

/*
WirelessWirelessLinksListOK describes a response with status code 200, with default header values.

WirelessWirelessLinksListOK wireless wireless links list o k
*/
type WirelessWirelessLinksListOK struct {
	Payload *WirelessWirelessLinksListOKBody
}

// IsSuccess returns true when this wireless wireless links list o k response has a 2xx status code
func (o *WirelessWirelessLinksListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links list o k response has a 3xx status code
func (o *WirelessWirelessLinksListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links list o k response has a 4xx status code
func (o *WirelessWirelessLinksListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links list o k response has a 5xx status code
func (o *WirelessWirelessLinksListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links list o k response a status code equal to that given
func (o *WirelessWirelessLinksListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless links list o k response
func (o *WirelessWirelessLinksListOK) Code() int {
	return 200
}

func (o *WirelessWirelessLinksListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-links/][%d] wirelessWirelessLinksListOK %s", 200, payload)
}

func (o *WirelessWirelessLinksListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-links/][%d] wirelessWirelessLinksListOK %s", 200, payload)
}

func (o *WirelessWirelessLinksListOK) GetPayload() *WirelessWirelessLinksListOKBody {
	return o.Payload
}

func (o *WirelessWirelessLinksListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(WirelessWirelessLinksListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
WirelessWirelessLinksListOKBody wireless wireless links list o k body
swagger:model WirelessWirelessLinksListOKBody
*/
type WirelessWirelessLinksListOKBody struct {

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
	Results []*models.WirelessLink `json:"results"`
}

// Validate validates this wireless wireless links list o k body
func (o *WirelessWirelessLinksListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *WirelessWirelessLinksListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLinksListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLinksListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLinksListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLinksListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLinksListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLinksListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLinksListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this wireless wireless links list o k body based on the context it is used
func (o *WirelessWirelessLinksListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WirelessWirelessLinksListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLinksListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WirelessWirelessLinksListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WirelessWirelessLinksListOKBody) UnmarshalBinary(b []byte) error {
	var res WirelessWirelessLinksListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}