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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scholdan/go-netbox/netbox/models"
)

// UsersUsersCreateReader is a Reader for the UsersUsersCreate structure.
type UsersUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersUsersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersCreateCreated creates a UsersUsersCreateCreated with default headers values
func NewUsersUsersCreateCreated() *UsersUsersCreateCreated {
	return &UsersUsersCreateCreated{}
}

/*
UsersUsersCreateCreated describes a response with status code 201, with default header values.

UsersUsersCreateCreated users users create created
*/
type UsersUsersCreateCreated struct {
	Payload *models.User
}

// IsSuccess returns true when this users users create created response has a 2xx status code
func (o *UsersUsersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users create created response has a 3xx status code
func (o *UsersUsersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users create created response has a 4xx status code
func (o *UsersUsersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users create created response has a 5xx status code
func (o *UsersUsersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users users create created response a status code equal to that given
func (o *UsersUsersCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the users users create created response
func (o *UsersUsersCreateCreated) Code() int {
	return 201
}

func (o *UsersUsersCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] usersUsersCreateCreated %s", 201, payload)
}

func (o *UsersUsersCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] usersUsersCreateCreated %s", 201, payload)
}

func (o *UsersUsersCreateCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersCreateDefault creates a UsersUsersCreateDefault with default headers values
func NewUsersUsersCreateDefault(code int) *UsersUsersCreateDefault {
	return &UsersUsersCreateDefault{
		_statusCode: code,
	}
}

/*
UsersUsersCreateDefault describes a response with status code -1, with default header values.

UsersUsersCreateDefault users users create default
*/
type UsersUsersCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users users create default response has a 2xx status code
func (o *UsersUsersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users create default response has a 3xx status code
func (o *UsersUsersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users create default response has a 4xx status code
func (o *UsersUsersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users create default response has a 5xx status code
func (o *UsersUsersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users create default response a status code equal to that given
func (o *UsersUsersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users users create default response
func (o *UsersUsersCreateDefault) Code() int {
	return o._statusCode
}

func (o *UsersUsersCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] users_users_create default %s", o._statusCode, payload)
}

func (o *UsersUsersCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/users/][%d] users_users_create default %s", o._statusCode, payload)
}

func (o *UsersUsersCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
