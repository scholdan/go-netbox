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

// IpamPrefixesPartialUpdateReader is a Reader for the IpamPrefixesPartialUpdate structure.
type IpamPrefixesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesPartialUpdateOK creates a IpamPrefixesPartialUpdateOK with default headers values
func NewIpamPrefixesPartialUpdateOK() *IpamPrefixesPartialUpdateOK {
	return &IpamPrefixesPartialUpdateOK{}
}

/*
IpamPrefixesPartialUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesPartialUpdateOK ipam prefixes partial update o k
*/
type IpamPrefixesPartialUpdateOK struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes partial update o k response has a 2xx status code
func (o *IpamPrefixesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes partial update o k response has a 3xx status code
func (o *IpamPrefixesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes partial update o k response has a 4xx status code
func (o *IpamPrefixesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes partial update o k response has a 5xx status code
func (o *IpamPrefixesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes partial update o k response a status code equal to that given
func (o *IpamPrefixesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam prefixes partial update o k response
func (o *IpamPrefixesPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamPrefixesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesPartialUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesPartialUpdateDefault creates a IpamPrefixesPartialUpdateDefault with default headers values
func NewIpamPrefixesPartialUpdateDefault(code int) *IpamPrefixesPartialUpdateDefault {
	return &IpamPrefixesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamPrefixesPartialUpdateDefault ipam prefixes partial update default
*/
type IpamPrefixesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes partial update default response has a 2xx status code
func (o *IpamPrefixesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes partial update default response has a 3xx status code
func (o *IpamPrefixesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes partial update default response has a 4xx status code
func (o *IpamPrefixesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes partial update default response has a 5xx status code
func (o *IpamPrefixesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes partial update default response a status code equal to that given
func (o *IpamPrefixesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam prefixes partial update default response
func (o *IpamPrefixesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamPrefixesPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipam_prefixes_partial_update default %s", o._statusCode, payload)
}

func (o *IpamPrefixesPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipam_prefixes_partial_update default %s", o._statusCode, payload)
}

func (o *IpamPrefixesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
