// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

	"github.com/scholdan/go-netbox/netbox/models"
)

// DcimPowerOutletsListReader is a Reader for the DcimPowerOutletsList structure.
type DcimPowerOutletsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsListOK creates a DcimPowerOutletsListOK with default headers values
func NewDcimPowerOutletsListOK() *DcimPowerOutletsListOK {
	return &DcimPowerOutletsListOK{}
}

/*
DcimPowerOutletsListOK describes a response with status code 200, with default header values.

DcimPowerOutletsListOK dcim power outlets list o k
*/
type DcimPowerOutletsListOK struct {
	Payload *DcimPowerOutletsListOKBody
}

// IsSuccess returns true when this dcim power outlets list o k response has a 2xx status code
func (o *DcimPowerOutletsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets list o k response has a 3xx status code
func (o *DcimPowerOutletsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets list o k response has a 4xx status code
func (o *DcimPowerOutletsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets list o k response has a 5xx status code
func (o *DcimPowerOutletsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets list o k response a status code equal to that given
func (o *DcimPowerOutletsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlets list o k response
func (o *DcimPowerOutletsListOK) Code() int {
	return 200
}

func (o *DcimPowerOutletsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-outlets/][%d] dcimPowerOutletsListOK %s", 200, payload)
}

func (o *DcimPowerOutletsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-outlets/][%d] dcimPowerOutletsListOK %s", 200, payload)
}

func (o *DcimPowerOutletsListOK) GetPayload() *DcimPowerOutletsListOKBody {
	return o.Payload
}

func (o *DcimPowerOutletsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimPowerOutletsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsListDefault creates a DcimPowerOutletsListDefault with default headers values
func NewDcimPowerOutletsListDefault(code int) *DcimPowerOutletsListDefault {
	return &DcimPowerOutletsListDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsListDefault describes a response with status code -1, with default header values.

DcimPowerOutletsListDefault dcim power outlets list default
*/
type DcimPowerOutletsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets list default response has a 2xx status code
func (o *DcimPowerOutletsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets list default response has a 3xx status code
func (o *DcimPowerOutletsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets list default response has a 4xx status code
func (o *DcimPowerOutletsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets list default response has a 5xx status code
func (o *DcimPowerOutletsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets list default response a status code equal to that given
func (o *DcimPowerOutletsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlets list default response
func (o *DcimPowerOutletsListDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletsListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-outlets/][%d] dcim_power-outlets_list default %s", o._statusCode, payload)
}

func (o *DcimPowerOutletsListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-outlets/][%d] dcim_power-outlets_list default %s", o._statusCode, payload)
}

func (o *DcimPowerOutletsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimPowerOutletsListOKBody dcim power outlets list o k body
swagger:model DcimPowerOutletsListOKBody
*/
type DcimPowerOutletsListOKBody struct {

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
	Results []*models.PowerOutlet `json:"results"`
}

// Validate validates this dcim power outlets list o k body
func (o *DcimPowerOutletsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimPowerOutletsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerOutletsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerOutletsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerOutletsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerOutletsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerOutletsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerOutletsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerOutletsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerOutletsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerOutletsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim power outlets list o k body based on the context it is used
func (o *DcimPowerOutletsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimPowerOutletsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerOutletsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerOutletsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimPowerOutletsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimPowerOutletsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimPowerOutletsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
