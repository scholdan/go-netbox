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

// IpamServicesReadReader is a Reader for the IpamServicesRead structure.
type IpamServicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesReadOK creates a IpamServicesReadOK with default headers values
func NewIpamServicesReadOK() *IpamServicesReadOK {
	return &IpamServicesReadOK{}
}

/*
IpamServicesReadOK describes a response with status code 200, with default header values.

IpamServicesReadOK ipam services read o k
*/
type IpamServicesReadOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services read o k response has a 2xx status code
func (o *IpamServicesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services read o k response has a 3xx status code
func (o *IpamServicesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services read o k response has a 4xx status code
func (o *IpamServicesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services read o k response has a 5xx status code
func (o *IpamServicesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services read o k response a status code equal to that given
func (o *IpamServicesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam services read o k response
func (o *IpamServicesReadOK) Code() int {
	return 200
}

func (o *IpamServicesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipamServicesReadOK %s", 200, payload)
}

func (o *IpamServicesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipamServicesReadOK %s", 200, payload)
}

func (o *IpamServicesReadOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesReadDefault creates a IpamServicesReadDefault with default headers values
func NewIpamServicesReadDefault(code int) *IpamServicesReadDefault {
	return &IpamServicesReadDefault{
		_statusCode: code,
	}
}

/*
IpamServicesReadDefault describes a response with status code -1, with default header values.

IpamServicesReadDefault ipam services read default
*/
type IpamServicesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam services read default response has a 2xx status code
func (o *IpamServicesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services read default response has a 3xx status code
func (o *IpamServicesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services read default response has a 4xx status code
func (o *IpamServicesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services read default response has a 5xx status code
func (o *IpamServicesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services read default response a status code equal to that given
func (o *IpamServicesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam services read default response
func (o *IpamServicesReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipam_services_read default %s", o._statusCode, payload)
}

func (o *IpamServicesReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/services/{id}/][%d] ipam_services_read default %s", o._statusCode, payload)
}

func (o *IpamServicesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
