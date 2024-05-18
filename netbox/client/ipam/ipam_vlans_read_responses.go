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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamVlansReadReader is a Reader for the IpamVlansRead structure.
type IpamVlansReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansReadOK creates a IpamVlansReadOK with default headers values
func NewIpamVlansReadOK() *IpamVlansReadOK {
	return &IpamVlansReadOK{}
}

/*
IpamVlansReadOK describes a response with status code 200, with default header values.

IpamVlansReadOK ipam vlans read o k
*/
type IpamVlansReadOK struct {
	Payload *models.VLAN
}

// IsSuccess returns true when this ipam vlans read o k response has a 2xx status code
func (o *IpamVlansReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlans read o k response has a 3xx status code
func (o *IpamVlansReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlans read o k response has a 4xx status code
func (o *IpamVlansReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlans read o k response has a 5xx status code
func (o *IpamVlansReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlans read o k response a status code equal to that given
func (o *IpamVlansReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vlans read o k response
func (o *IpamVlansReadOK) Code() int {
	return 200
}

func (o *IpamVlansReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlans/{id}/][%d] ipamVlansReadOK %s", 200, payload)
}

func (o *IpamVlansReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlans/{id}/][%d] ipamVlansReadOK %s", 200, payload)
}

func (o *IpamVlansReadOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlansReadDefault creates a IpamVlansReadDefault with default headers values
func NewIpamVlansReadDefault(code int) *IpamVlansReadDefault {
	return &IpamVlansReadDefault{
		_statusCode: code,
	}
}

/*
IpamVlansReadDefault describes a response with status code -1, with default header values.

IpamVlansReadDefault ipam vlans read default
*/
type IpamVlansReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vlans read default response has a 2xx status code
func (o *IpamVlansReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlans read default response has a 3xx status code
func (o *IpamVlansReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlans read default response has a 4xx status code
func (o *IpamVlansReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlans read default response has a 5xx status code
func (o *IpamVlansReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlans read default response a status code equal to that given
func (o *IpamVlansReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vlans read default response
func (o *IpamVlansReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlans/{id}/][%d] ipam_vlans_read default %s", o._statusCode, payload)
}

func (o *IpamVlansReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlans/{id}/][%d] ipam_vlans_read default %s", o._statusCode, payload)
}

func (o *IpamVlansReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
