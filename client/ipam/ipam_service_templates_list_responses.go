// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// IpamServiceTemplatesListReader is a Reader for the IpamServiceTemplatesList structure.
type IpamServiceTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServiceTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServiceTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/service-templates/] ipam_service-templates_list", response, response.Code())
	}
}

// NewIpamServiceTemplatesListOK creates a IpamServiceTemplatesListOK with default headers values
func NewIpamServiceTemplatesListOK() *IpamServiceTemplatesListOK {
	return &IpamServiceTemplatesListOK{}
}

/*
IpamServiceTemplatesListOK describes a response with status code 200, with default header values.

IpamServiceTemplatesListOK ipam service templates list o k
*/
type IpamServiceTemplatesListOK struct {
	Payload *IpamServiceTemplatesListOKBody
}

// IsSuccess returns true when this ipam service templates list o k response has a 2xx status code
func (o *IpamServiceTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam service templates list o k response has a 3xx status code
func (o *IpamServiceTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam service templates list o k response has a 4xx status code
func (o *IpamServiceTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam service templates list o k response has a 5xx status code
func (o *IpamServiceTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam service templates list o k response a status code equal to that given
func (o *IpamServiceTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam service templates list o k response
func (o *IpamServiceTemplatesListOK) Code() int {
	return 200
}

func (o *IpamServiceTemplatesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/service-templates/][%d] ipamServiceTemplatesListOK %s", 200, payload)
}

func (o *IpamServiceTemplatesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/service-templates/][%d] ipamServiceTemplatesListOK %s", 200, payload)
}

func (o *IpamServiceTemplatesListOK) GetPayload() *IpamServiceTemplatesListOKBody {
	return o.Payload
}

func (o *IpamServiceTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamServiceTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IpamServiceTemplatesListOKBody ipam service templates list o k body
swagger:model IpamServiceTemplatesListOKBody
*/
type IpamServiceTemplatesListOKBody struct {

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
	Results []*models.ServiceTemplate `json:"results"`
}

// Validate validates this ipam service templates list o k body
func (o *IpamServiceTemplatesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *IpamServiceTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamServiceTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamServiceTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamServiceTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamServiceTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamServiceTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam service templates list o k body based on the context it is used
func (o *IpamServiceTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamServiceTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamServiceTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamServiceTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamServiceTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamServiceTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}