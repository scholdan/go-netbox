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

	"github.com/scholdan/go-netbox/netbox/models"
)

// TenancyContactGroupsListReader is a Reader for the TenancyContactGroupsList structure.
type TenancyContactGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactGroupsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsListOK creates a TenancyContactGroupsListOK with default headers values
func NewTenancyContactGroupsListOK() *TenancyContactGroupsListOK {
	return &TenancyContactGroupsListOK{}
}

/*
TenancyContactGroupsListOK describes a response with status code 200, with default header values.

TenancyContactGroupsListOK tenancy contact groups list o k
*/
type TenancyContactGroupsListOK struct {
	Payload *TenancyContactGroupsListOKBody
}

// IsSuccess returns true when this tenancy contact groups list o k response has a 2xx status code
func (o *TenancyContactGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups list o k response has a 3xx status code
func (o *TenancyContactGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups list o k response has a 4xx status code
func (o *TenancyContactGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups list o k response has a 5xx status code
func (o *TenancyContactGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups list o k response a status code equal to that given
func (o *TenancyContactGroupsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact groups list o k response
func (o *TenancyContactGroupsListOK) Code() int {
	return 200
}

func (o *TenancyContactGroupsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/][%d] tenancyContactGroupsListOK %s", 200, payload)
}

func (o *TenancyContactGroupsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/][%d] tenancyContactGroupsListOK %s", 200, payload)
}

func (o *TenancyContactGroupsListOK) GetPayload() *TenancyContactGroupsListOKBody {
	return o.Payload
}

func (o *TenancyContactGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TenancyContactGroupsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsListDefault creates a TenancyContactGroupsListDefault with default headers values
func NewTenancyContactGroupsListDefault(code int) *TenancyContactGroupsListDefault {
	return &TenancyContactGroupsListDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsListDefault describes a response with status code -1, with default header values.

TenancyContactGroupsListDefault tenancy contact groups list default
*/
type TenancyContactGroupsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact groups list default response has a 2xx status code
func (o *TenancyContactGroupsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups list default response has a 3xx status code
func (o *TenancyContactGroupsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups list default response has a 4xx status code
func (o *TenancyContactGroupsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups list default response has a 5xx status code
func (o *TenancyContactGroupsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups list default response a status code equal to that given
func (o *TenancyContactGroupsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact groups list default response
func (o *TenancyContactGroupsListDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactGroupsListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/][%d] tenancy_contact-groups_list default %s", o._statusCode, payload)
}

func (o *TenancyContactGroupsListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-groups/][%d] tenancy_contact-groups_list default %s", o._statusCode, payload)
}

func (o *TenancyContactGroupsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
TenancyContactGroupsListOKBody tenancy contact groups list o k body
swagger:model TenancyContactGroupsListOKBody
*/
type TenancyContactGroupsListOKBody struct {

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
	Results []*models.ContactGroup `json:"results"`
}

// Validate validates this tenancy contact groups list o k body
func (o *TenancyContactGroupsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *TenancyContactGroupsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactGroupsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactGroupsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactGroupsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactGroupsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("tenancyContactGroupsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *TenancyContactGroupsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("tenancyContactGroupsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tenancy contact groups list o k body based on the context it is used
func (o *TenancyContactGroupsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TenancyContactGroupsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {

			if swag.IsZero(o.Results[i]) { // not required
				return nil
			}

			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenancyContactGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenancyContactGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TenancyContactGroupsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TenancyContactGroupsListOKBody) UnmarshalBinary(b []byte) error {
	var res TenancyContactGroupsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
