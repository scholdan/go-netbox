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

package ipam

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

// IpamPrefixesCreateReader is a Reader for the IpamPrefixesCreate structure.
type IpamPrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesCreateCreated creates a IpamPrefixesCreateCreated with default headers values
func NewIpamPrefixesCreateCreated() *IpamPrefixesCreateCreated {
	return &IpamPrefixesCreateCreated{}
}

/*
IpamPrefixesCreateCreated describes a response with status code 201, with default header values.

IpamPrefixesCreateCreated ipam prefixes create created
*/
type IpamPrefixesCreateCreated struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes create created response has a 2xx status code
func (o *IpamPrefixesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes create created response has a 3xx status code
func (o *IpamPrefixesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes create created response has a 4xx status code
func (o *IpamPrefixesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes create created response has a 5xx status code
func (o *IpamPrefixesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes create created response a status code equal to that given
func (o *IpamPrefixesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam prefixes create created response
func (o *IpamPrefixesCreateCreated) Code() int {
	return 201
}

func (o *IpamPrefixesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipamPrefixesCreateCreated %s", 201, payload)
}

func (o *IpamPrefixesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipamPrefixesCreateCreated %s", 201, payload)
}

func (o *IpamPrefixesCreateCreated) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesCreateDefault creates a IpamPrefixesCreateDefault with default headers values
func NewIpamPrefixesCreateDefault(code int) *IpamPrefixesCreateDefault {
	return &IpamPrefixesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesCreateDefault describes a response with status code -1, with default header values.

IpamPrefixesCreateDefault ipam prefixes create default
*/
type IpamPrefixesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes create default response has a 2xx status code
func (o *IpamPrefixesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes create default response has a 3xx status code
func (o *IpamPrefixesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes create default response has a 4xx status code
func (o *IpamPrefixesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes create default response has a 5xx status code
func (o *IpamPrefixesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes create default response a status code equal to that given
func (o *IpamPrefixesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam prefixes create default response
func (o *IpamPrefixesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipam_prefixes_create default %s", o._statusCode, payload)
}

func (o *IpamPrefixesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/prefixes/][%d] ipam_prefixes_create default %s", o._statusCode, payload)
}

func (o *IpamPrefixesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
