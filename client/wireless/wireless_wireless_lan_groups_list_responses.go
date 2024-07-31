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

	"github.com/fbreckle/go-netbox/models"
)

// WirelessWirelessLanGroupsListReader is a Reader for the WirelessWirelessLanGroupsList structure.
type WirelessWirelessLanGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /wireless/wireless-lan-groups/] wireless_wireless-lan-groups_list", response, response.Code())
	}
}

// NewWirelessWirelessLanGroupsListOK creates a WirelessWirelessLanGroupsListOK with default headers values
func NewWirelessWirelessLanGroupsListOK() *WirelessWirelessLanGroupsListOK {
	return &WirelessWirelessLanGroupsListOK{}
}

/*
WirelessWirelessLanGroupsListOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsListOK wireless wireless lan groups list o k
*/
type WirelessWirelessLanGroupsListOK struct {
	Payload *WirelessWirelessLanGroupsListOKBody
}

// IsSuccess returns true when this wireless wireless lan groups list o k response has a 2xx status code
func (o *WirelessWirelessLanGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups list o k response has a 3xx status code
func (o *WirelessWirelessLanGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups list o k response has a 4xx status code
func (o *WirelessWirelessLanGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups list o k response has a 5xx status code
func (o *WirelessWirelessLanGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups list o k response a status code equal to that given
func (o *WirelessWirelessLanGroupsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless lan groups list o k response
func (o *WirelessWirelessLanGroupsListOK) Code() int {
	return 200
}

func (o *WirelessWirelessLanGroupsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsListOK %s", 200, payload)
}

func (o *WirelessWirelessLanGroupsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsListOK %s", 200, payload)
}

func (o *WirelessWirelessLanGroupsListOK) GetPayload() *WirelessWirelessLanGroupsListOKBody {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(WirelessWirelessLanGroupsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
WirelessWirelessLanGroupsListOKBody wireless wireless lan groups list o k body
swagger:model WirelessWirelessLanGroupsListOKBody
*/
type WirelessWirelessLanGroupsListOKBody struct {

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
	Results []*models.WirelessLANGroup `json:"results"`
}

// Validate validates this wireless wireless lan groups list o k body
func (o *WirelessWirelessLanGroupsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *WirelessWirelessLanGroupsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLanGroupsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLanGroupsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLanGroupsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLanGroupsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLanGroupsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLanGroupsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLanGroupsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this wireless wireless lan groups list o k body based on the context it is used
func (o *WirelessWirelessLanGroupsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WirelessWirelessLanGroupsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WirelessWirelessLanGroupsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WirelessWirelessLanGroupsListOKBody) UnmarshalBinary(b []byte) error {
	var res WirelessWirelessLanGroupsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
