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

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StatusListReader is a Reader for the StatusList structure.
type StatusListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusListOK creates a StatusListOK with default headers values
func NewStatusListOK() *StatusListOK {
	return &StatusListOK{}
}

/*
StatusListOK describes a response with status code 200, with default header values.

StatusListOK status list o k
*/
type StatusListOK struct {
	Payload interface{}
}

// IsSuccess returns true when this status list o k response has a 2xx status code
func (o *StatusListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this status list o k response has a 3xx status code
func (o *StatusListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this status list o k response has a 4xx status code
func (o *StatusListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this status list o k response has a 5xx status code
func (o *StatusListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this status list o k response a status code equal to that given
func (o *StatusListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the status list o k response
func (o *StatusListOK) Code() int {
	return 200
}

func (o *StatusListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/][%d] statusListOK %s", 200, payload)
}

func (o *StatusListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/][%d] statusListOK %s", 200, payload)
}

func (o *StatusListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StatusListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusListDefault creates a StatusListDefault with default headers values
func NewStatusListDefault(code int) *StatusListDefault {
	return &StatusListDefault{
		_statusCode: code,
	}
}

/*
StatusListDefault describes a response with status code -1, with default header values.

StatusListDefault status list default
*/
type StatusListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this status list default response has a 2xx status code
func (o *StatusListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this status list default response has a 3xx status code
func (o *StatusListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this status list default response has a 4xx status code
func (o *StatusListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this status list default response has a 5xx status code
func (o *StatusListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this status list default response a status code equal to that given
func (o *StatusListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the status list default response
func (o *StatusListDefault) Code() int {
	return o._statusCode
}

func (o *StatusListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/][%d] status_list default %s", o._statusCode, payload)
}

func (o *StatusListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /status/][%d] status_list default %s", o._statusCode, payload)
}

func (o *StatusListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *StatusListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
