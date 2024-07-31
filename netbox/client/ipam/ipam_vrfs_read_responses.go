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

// IpamVrfsReadReader is a Reader for the IpamVrfsRead structure.
type IpamVrfsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsReadOK creates a IpamVrfsReadOK with default headers values
func NewIpamVrfsReadOK() *IpamVrfsReadOK {
	return &IpamVrfsReadOK{}
}

/*
IpamVrfsReadOK describes a response with status code 200, with default header values.

IpamVrfsReadOK ipam vrfs read o k
*/
type IpamVrfsReadOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs read o k response has a 2xx status code
func (o *IpamVrfsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs read o k response has a 3xx status code
func (o *IpamVrfsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs read o k response has a 4xx status code
func (o *IpamVrfsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs read o k response has a 5xx status code
func (o *IpamVrfsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs read o k response a status code equal to that given
func (o *IpamVrfsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vrfs read o k response
func (o *IpamVrfsReadOK) Code() int {
	return 200
}

func (o *IpamVrfsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipamVrfsReadOK %s", 200, payload)
}

func (o *IpamVrfsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipamVrfsReadOK %s", 200, payload)
}

func (o *IpamVrfsReadOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsReadDefault creates a IpamVrfsReadDefault with default headers values
func NewIpamVrfsReadDefault(code int) *IpamVrfsReadDefault {
	return &IpamVrfsReadDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsReadDefault describes a response with status code -1, with default header values.

IpamVrfsReadDefault ipam vrfs read default
*/
type IpamVrfsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vrfs read default response has a 2xx status code
func (o *IpamVrfsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs read default response has a 3xx status code
func (o *IpamVrfsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs read default response has a 4xx status code
func (o *IpamVrfsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs read default response has a 5xx status code
func (o *IpamVrfsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs read default response a status code equal to that given
func (o *IpamVrfsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vrfs read default response
func (o *IpamVrfsReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipam_vrfs_read default %s", o._statusCode, payload)
}

func (o *IpamVrfsReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipam_vrfs_read default %s", o._statusCode, payload)
}

func (o *IpamVrfsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
