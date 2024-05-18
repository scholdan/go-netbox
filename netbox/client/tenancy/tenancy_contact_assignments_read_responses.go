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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyContactAssignmentsReadReader is a Reader for the TenancyContactAssignmentsRead structure.
type TenancyContactAssignmentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsReadOK creates a TenancyContactAssignmentsReadOK with default headers values
func NewTenancyContactAssignmentsReadOK() *TenancyContactAssignmentsReadOK {
	return &TenancyContactAssignmentsReadOK{}
}

/*
TenancyContactAssignmentsReadOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsReadOK tenancy contact assignments read o k
*/
type TenancyContactAssignmentsReadOK struct {
	Payload *models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments read o k response has a 2xx status code
func (o *TenancyContactAssignmentsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments read o k response has a 3xx status code
func (o *TenancyContactAssignmentsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments read o k response has a 4xx status code
func (o *TenancyContactAssignmentsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments read o k response has a 5xx status code
func (o *TenancyContactAssignmentsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments read o k response a status code equal to that given
func (o *TenancyContactAssignmentsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenancy contact assignments read o k response
func (o *TenancyContactAssignmentsReadOK) Code() int {
	return 200
}

func (o *TenancyContactAssignmentsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsReadOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsReadOK %s", 200, payload)
}

func (o *TenancyContactAssignmentsReadOK) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsReadDefault creates a TenancyContactAssignmentsReadDefault with default headers values
func NewTenancyContactAssignmentsReadDefault(code int) *TenancyContactAssignmentsReadDefault {
	return &TenancyContactAssignmentsReadDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsReadDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsReadDefault tenancy contact assignments read default
*/
type TenancyContactAssignmentsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact assignments read default response has a 2xx status code
func (o *TenancyContactAssignmentsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments read default response has a 3xx status code
func (o *TenancyContactAssignmentsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments read default response has a 4xx status code
func (o *TenancyContactAssignmentsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments read default response has a 5xx status code
func (o *TenancyContactAssignmentsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments read default response a status code equal to that given
func (o *TenancyContactAssignmentsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact assignments read default response
func (o *TenancyContactAssignmentsReadDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactAssignmentsReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_read default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_read default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
