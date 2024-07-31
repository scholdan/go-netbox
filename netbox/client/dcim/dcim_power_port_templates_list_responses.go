// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// DcimPowerPortTemplatesListReader is a Reader for the DcimPowerPortTemplatesList structure.
type DcimPowerPortTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /dcim/power-port-templates/] dcim_power-port-templates_list", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesListOK creates a DcimPowerPortTemplatesListOK with default headers values
func NewDcimPowerPortTemplatesListOK() *DcimPowerPortTemplatesListOK {
	return &DcimPowerPortTemplatesListOK{}
}

/*
DcimPowerPortTemplatesListOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesListOK dcim power port templates list o k
*/
type DcimPowerPortTemplatesListOK struct {
	Payload *DcimPowerPortTemplatesListOKBody
}

// IsSuccess returns true when this dcim power port templates list o k response has a 2xx status code
func (o *DcimPowerPortTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates list o k response has a 3xx status code
func (o *DcimPowerPortTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates list o k response has a 4xx status code
func (o *DcimPowerPortTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates list o k response has a 5xx status code
func (o *DcimPowerPortTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates list o k response a status code equal to that given
func (o *DcimPowerPortTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power port templates list o k response
func (o *DcimPowerPortTemplatesListOK) Code() int {
	return 200
}

func (o *DcimPowerPortTemplatesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-port-templates/][%d] dcimPowerPortTemplatesListOK %s", 200, payload)
}

func (o *DcimPowerPortTemplatesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-port-templates/][%d] dcimPowerPortTemplatesListOK %s", 200, payload)
}

func (o *DcimPowerPortTemplatesListOK) GetPayload() *DcimPowerPortTemplatesListOKBody {
	return o.Payload
}

func (o *DcimPowerPortTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimPowerPortTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimPowerPortTemplatesListOKBody dcim power port templates list o k body
swagger:model DcimPowerPortTemplatesListOKBody
*/
type DcimPowerPortTemplatesListOKBody struct {

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
	Results []*models.PowerPortTemplate `json:"results"`
}

// Validate validates this dcim power port templates list o k body
func (o *DcimPowerPortTemplatesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimPowerPortTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPortTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPortTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPortTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPortTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPortTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim power port templates list o k body based on the context it is used
func (o *DcimPowerPortTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimPowerPortTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimPowerPortTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimPowerPortTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimPowerPortTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
