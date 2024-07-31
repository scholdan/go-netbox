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

// IpamRouteTargetsPartialUpdateReader is a Reader for the IpamRouteTargetsPartialUpdate structure.
type IpamRouteTargetsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsPartialUpdateOK creates a IpamRouteTargetsPartialUpdateOK with default headers values
func NewIpamRouteTargetsPartialUpdateOK() *IpamRouteTargetsPartialUpdateOK {
	return &IpamRouteTargetsPartialUpdateOK{}
}

/*
IpamRouteTargetsPartialUpdateOK describes a response with status code 200, with default header values.

IpamRouteTargetsPartialUpdateOK ipam route targets partial update o k
*/
type IpamRouteTargetsPartialUpdateOK struct {
	Payload *models.RouteTarget
}

// IsSuccess returns true when this ipam route targets partial update o k response has a 2xx status code
func (o *IpamRouteTargetsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets partial update o k response has a 3xx status code
func (o *IpamRouteTargetsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets partial update o k response has a 4xx status code
func (o *IpamRouteTargetsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets partial update o k response has a 5xx status code
func (o *IpamRouteTargetsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets partial update o k response a status code equal to that given
func (o *IpamRouteTargetsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam route targets partial update o k response
func (o *IpamRouteTargetsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRouteTargetsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/route-targets/{id}/][%d] ipamRouteTargetsPartialUpdateOK %s", 200, payload)
}

func (o *IpamRouteTargetsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/route-targets/{id}/][%d] ipamRouteTargetsPartialUpdateOK %s", 200, payload)
}

func (o *IpamRouteTargetsPartialUpdateOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsPartialUpdateDefault creates a IpamRouteTargetsPartialUpdateDefault with default headers values
func NewIpamRouteTargetsPartialUpdateDefault(code int) *IpamRouteTargetsPartialUpdateDefault {
	return &IpamRouteTargetsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamRouteTargetsPartialUpdateDefault ipam route targets partial update default
*/
type IpamRouteTargetsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam route targets partial update default response has a 2xx status code
func (o *IpamRouteTargetsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets partial update default response has a 3xx status code
func (o *IpamRouteTargetsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets partial update default response has a 4xx status code
func (o *IpamRouteTargetsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets partial update default response has a 5xx status code
func (o *IpamRouteTargetsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets partial update default response a status code equal to that given
func (o *IpamRouteTargetsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam route targets partial update default response
func (o *IpamRouteTargetsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/route-targets/{id}/][%d] ipam_route-targets_partial_update default %s", o._statusCode, payload)
}

func (o *IpamRouteTargetsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/route-targets/{id}/][%d] ipam_route-targets_partial_update default %s", o._statusCode, payload)
}

func (o *IpamRouteTargetsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
