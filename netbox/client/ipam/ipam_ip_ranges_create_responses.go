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

// IpamIPRangesCreateReader is a Reader for the IpamIPRangesCreate structure.
type IpamIPRangesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamIPRangesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesCreateCreated creates a IpamIPRangesCreateCreated with default headers values
func NewIpamIPRangesCreateCreated() *IpamIPRangesCreateCreated {
	return &IpamIPRangesCreateCreated{}
}

/*
IpamIPRangesCreateCreated describes a response with status code 201, with default header values.

IpamIPRangesCreateCreated ipam Ip ranges create created
*/
type IpamIPRangesCreateCreated struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges create created response has a 2xx status code
func (o *IpamIPRangesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges create created response has a 3xx status code
func (o *IpamIPRangesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges create created response has a 4xx status code
func (o *IpamIPRangesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges create created response has a 5xx status code
func (o *IpamIPRangesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges create created response a status code equal to that given
func (o *IpamIPRangesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam Ip ranges create created response
func (o *IpamIPRangesCreateCreated) Code() int {
	return 201
}

func (o *IpamIPRangesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated %s", 201, payload)
}

func (o *IpamIPRangesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated %s", 201, payload)
}

func (o *IpamIPRangesCreateCreated) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesCreateDefault creates a IpamIPRangesCreateDefault with default headers values
func NewIpamIPRangesCreateDefault(code int) *IpamIPRangesCreateDefault {
	return &IpamIPRangesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesCreateDefault describes a response with status code -1, with default header values.

IpamIPRangesCreateDefault ipam ip ranges create default
*/
type IpamIPRangesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam ip ranges create default response has a 2xx status code
func (o *IpamIPRangesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges create default response has a 3xx status code
func (o *IpamIPRangesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges create default response has a 4xx status code
func (o *IpamIPRangesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges create default response has a 5xx status code
func (o *IpamIPRangesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges create default response a status code equal to that given
func (o *IpamIPRangesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam ip ranges create default response
func (o *IpamIPRangesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPRangesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipam_ip-ranges_create default %s", o._statusCode, payload)
}

func (o *IpamIPRangesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipam_ip-ranges_create default %s", o._statusCode, payload)
}

func (o *IpamIPRangesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
