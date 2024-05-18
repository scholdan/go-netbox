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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersTokensCreateReader is a Reader for the UsersTokensCreate structure.
type UsersTokensCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersTokensCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensCreateCreated creates a UsersTokensCreateCreated with default headers values
func NewUsersTokensCreateCreated() *UsersTokensCreateCreated {
	return &UsersTokensCreateCreated{}
}

/*
UsersTokensCreateCreated describes a response with status code 201, with default header values.

UsersTokensCreateCreated users tokens create created
*/
type UsersTokensCreateCreated struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens create created response has a 2xx status code
func (o *UsersTokensCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens create created response has a 3xx status code
func (o *UsersTokensCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens create created response has a 4xx status code
func (o *UsersTokensCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens create created response has a 5xx status code
func (o *UsersTokensCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens create created response a status code equal to that given
func (o *UsersTokensCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the users tokens create created response
func (o *UsersTokensCreateCreated) Code() int {
	return 201
}

func (o *UsersTokensCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateCreated %s", 201, payload)
}

func (o *UsersTokensCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateCreated %s", 201, payload)
}

func (o *UsersTokensCreateCreated) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensCreateDefault creates a UsersTokensCreateDefault with default headers values
func NewUsersTokensCreateDefault(code int) *UsersTokensCreateDefault {
	return &UsersTokensCreateDefault{
		_statusCode: code,
	}
}

/*
UsersTokensCreateDefault describes a response with status code -1, with default header values.

UsersTokensCreateDefault users tokens create default
*/
type UsersTokensCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users tokens create default response has a 2xx status code
func (o *UsersTokensCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens create default response has a 3xx status code
func (o *UsersTokensCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens create default response has a 4xx status code
func (o *UsersTokensCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens create default response has a 5xx status code
func (o *UsersTokensCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens create default response a status code equal to that given
func (o *UsersTokensCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users tokens create default response
func (o *UsersTokensCreateDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/tokens/][%d] users_tokens_create default %s", o._statusCode, payload)
}

func (o *UsersTokensCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/tokens/][%d] users_tokens_create default %s", o._statusCode, payload)
}

func (o *UsersTokensCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
