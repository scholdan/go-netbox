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
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyContactAssignmentsDeleteReader is a Reader for the TenancyContactAssignmentsDelete structure.
type TenancyContactAssignmentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactAssignmentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsDeleteNoContent creates a TenancyContactAssignmentsDeleteNoContent with default headers values
func NewTenancyContactAssignmentsDeleteNoContent() *TenancyContactAssignmentsDeleteNoContent {
	return &TenancyContactAssignmentsDeleteNoContent{}
}

/*
TenancyContactAssignmentsDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactAssignmentsDeleteNoContent tenancy contact assignments delete no content
*/
type TenancyContactAssignmentsDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy contact assignments delete no content response has a 2xx status code
func (o *TenancyContactAssignmentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments delete no content response has a 3xx status code
func (o *TenancyContactAssignmentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments delete no content response has a 4xx status code
func (o *TenancyContactAssignmentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments delete no content response has a 5xx status code
func (o *TenancyContactAssignmentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments delete no content response a status code equal to that given
func (o *TenancyContactAssignmentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the tenancy contact assignments delete no content response
func (o *TenancyContactAssignmentsDeleteNoContent) Code() int {
	return 204
}

func (o *TenancyContactAssignmentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsDeleteNoContent", 204)
}

func (o *TenancyContactAssignmentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsDeleteNoContent", 204)
}

func (o *TenancyContactAssignmentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactAssignmentsDeleteDefault creates a TenancyContactAssignmentsDeleteDefault with default headers values
func NewTenancyContactAssignmentsDeleteDefault(code int) *TenancyContactAssignmentsDeleteDefault {
	return &TenancyContactAssignmentsDeleteDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsDeleteDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsDeleteDefault tenancy contact assignments delete default
*/
type TenancyContactAssignmentsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact assignments delete default response has a 2xx status code
func (o *TenancyContactAssignmentsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments delete default response has a 3xx status code
func (o *TenancyContactAssignmentsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments delete default response has a 4xx status code
func (o *TenancyContactAssignmentsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments delete default response has a 5xx status code
func (o *TenancyContactAssignmentsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments delete default response a status code equal to that given
func (o *TenancyContactAssignmentsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact assignments delete default response
func (o *TenancyContactAssignmentsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactAssignmentsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_delete default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_delete default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
