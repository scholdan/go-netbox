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

// TenancyContactAssignmentsCreateReader is a Reader for the TenancyContactAssignmentsCreate structure.
type TenancyContactAssignmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyContactAssignmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsCreateCreated creates a TenancyContactAssignmentsCreateCreated with default headers values
func NewTenancyContactAssignmentsCreateCreated() *TenancyContactAssignmentsCreateCreated {
	return &TenancyContactAssignmentsCreateCreated{}
}

/*
TenancyContactAssignmentsCreateCreated describes a response with status code 201, with default header values.

TenancyContactAssignmentsCreateCreated tenancy contact assignments create created
*/
type TenancyContactAssignmentsCreateCreated struct {
	Payload *models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments create created response has a 2xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments create created response has a 3xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments create created response has a 4xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments create created response has a 5xx status code
func (o *TenancyContactAssignmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments create created response a status code equal to that given
func (o *TenancyContactAssignmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy contact assignments create created response
func (o *TenancyContactAssignmentsCreateCreated) Code() int {
	return 201
}

func (o *TenancyContactAssignmentsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancyContactAssignmentsCreateCreated %s", 201, payload)
}

func (o *TenancyContactAssignmentsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancyContactAssignmentsCreateCreated %s", 201, payload)
}

func (o *TenancyContactAssignmentsCreateCreated) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsCreateDefault creates a TenancyContactAssignmentsCreateDefault with default headers values
func NewTenancyContactAssignmentsCreateDefault(code int) *TenancyContactAssignmentsCreateDefault {
	return &TenancyContactAssignmentsCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsCreateDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsCreateDefault tenancy contact assignments create default
*/
type TenancyContactAssignmentsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contact assignments create default response has a 2xx status code
func (o *TenancyContactAssignmentsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments create default response has a 3xx status code
func (o *TenancyContactAssignmentsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments create default response has a 4xx status code
func (o *TenancyContactAssignmentsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments create default response has a 5xx status code
func (o *TenancyContactAssignmentsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments create default response a status code equal to that given
func (o *TenancyContactAssignmentsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contact assignments create default response
func (o *TenancyContactAssignmentsCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactAssignmentsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancy_contact-assignments_create default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenancy/contact-assignments/][%d] tenancy_contact-assignments_create default %s", o._statusCode, payload)
}

func (o *TenancyContactAssignmentsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
