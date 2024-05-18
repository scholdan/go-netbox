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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimCablesListReader is a Reader for the DcimCablesList structure.
type DcimCablesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesListOK creates a DcimCablesListOK with default headers values
func NewDcimCablesListOK() *DcimCablesListOK {
	return &DcimCablesListOK{}
}

/*
DcimCablesListOK describes a response with status code 200, with default header values.

DcimCablesListOK dcim cables list o k
*/
type DcimCablesListOK struct {
	Payload *DcimCablesListOKBody
}

// IsSuccess returns true when this dcim cables list o k response has a 2xx status code
func (o *DcimCablesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables list o k response has a 3xx status code
func (o *DcimCablesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables list o k response has a 4xx status code
func (o *DcimCablesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables list o k response has a 5xx status code
func (o *DcimCablesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables list o k response a status code equal to that given
func (o *DcimCablesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim cables list o k response
func (o *DcimCablesListOK) Code() int {
	return 200
}

func (o *DcimCablesListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListOK %s", 200, payload)
}

func (o *DcimCablesListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListOK %s", 200, payload)
}

func (o *DcimCablesListOK) GetPayload() *DcimCablesListOKBody {
	return o.Payload
}

func (o *DcimCablesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimCablesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesListDefault creates a DcimCablesListDefault with default headers values
func NewDcimCablesListDefault(code int) *DcimCablesListDefault {
	return &DcimCablesListDefault{
		_statusCode: code,
	}
}

/*
DcimCablesListDefault describes a response with status code -1, with default header values.

DcimCablesListDefault dcim cables list default
*/
type DcimCablesListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim cables list default response has a 2xx status code
func (o *DcimCablesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables list default response has a 3xx status code
func (o *DcimCablesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables list default response has a 4xx status code
func (o *DcimCablesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables list default response has a 5xx status code
func (o *DcimCablesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables list default response a status code equal to that given
func (o *DcimCablesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim cables list default response
func (o *DcimCablesListDefault) Code() int {
	return o._statusCode
}

func (o *DcimCablesListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcim_cables_list default %s", o._statusCode, payload)
}

func (o *DcimCablesListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcim_cables_list default %s", o._statusCode, payload)
}

func (o *DcimCablesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimCablesListOKBody dcim cables list o k body
swagger:model DcimCablesListOKBody
*/
type DcimCablesListOKBody struct {

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
	Results []*models.Cable `json:"results"`
}

// Validate validates this dcim cables list o k body
func (o *DcimCablesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimCablesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimCablesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCablesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCablesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimCablesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim cables list o k body based on the context it is used
func (o *DcimCablesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCablesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimCablesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimCablesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimCablesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
