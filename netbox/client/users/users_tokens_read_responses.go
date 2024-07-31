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

// UsersTokensReadReader is a Reader for the UsersTokensRead structure.
type UsersTokensReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersTokensReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensReadOK creates a UsersTokensReadOK with default headers values
func NewUsersTokensReadOK() *UsersTokensReadOK {
	return &UsersTokensReadOK{}
}

/*
UsersTokensReadOK describes a response with status code 200, with default header values.

UsersTokensReadOK users tokens read o k
*/
type UsersTokensReadOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens read o k response has a 2xx status code
func (o *UsersTokensReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens read o k response has a 3xx status code
func (o *UsersTokensReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens read o k response has a 4xx status code
func (o *UsersTokensReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens read o k response has a 5xx status code
func (o *UsersTokensReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens read o k response a status code equal to that given
func (o *UsersTokensReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users tokens read o k response
func (o *UsersTokensReadOK) Code() int {
	return 200
}

func (o *UsersTokensReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadOK %s", 200, payload)
}

func (o *UsersTokensReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadOK %s", 200, payload)
}

func (o *UsersTokensReadOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensReadDefault creates a UsersTokensReadDefault with default headers values
func NewUsersTokensReadDefault(code int) *UsersTokensReadDefault {
	return &UsersTokensReadDefault{
		_statusCode: code,
	}
}

/*
UsersTokensReadDefault describes a response with status code -1, with default header values.

UsersTokensReadDefault users tokens read default
*/
type UsersTokensReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this users tokens read default response has a 2xx status code
func (o *UsersTokensReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens read default response has a 3xx status code
func (o *UsersTokensReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens read default response has a 4xx status code
func (o *UsersTokensReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens read default response has a 5xx status code
func (o *UsersTokensReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens read default response a status code equal to that given
func (o *UsersTokensReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users tokens read default response
func (o *UsersTokensReadDefault) Code() int {
	return o._statusCode
}

func (o *UsersTokensReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] users_tokens_read default %s", o._statusCode, payload)
}

func (o *UsersTokensReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] users_tokens_read default %s", o._statusCode, payload)
}

func (o *UsersTokensReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
