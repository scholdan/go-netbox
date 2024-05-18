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
)

// IpamRirsDeleteReader is a Reader for the IpamRirsDelete structure.
type IpamRirsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRirsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsDeleteNoContent creates a IpamRirsDeleteNoContent with default headers values
func NewIpamRirsDeleteNoContent() *IpamRirsDeleteNoContent {
	return &IpamRirsDeleteNoContent{}
}

/*
IpamRirsDeleteNoContent describes a response with status code 204, with default header values.

IpamRirsDeleteNoContent ipam rirs delete no content
*/
type IpamRirsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam rirs delete no content response has a 2xx status code
func (o *IpamRirsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs delete no content response has a 3xx status code
func (o *IpamRirsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs delete no content response has a 4xx status code
func (o *IpamRirsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs delete no content response has a 5xx status code
func (o *IpamRirsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs delete no content response a status code equal to that given
func (o *IpamRirsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the ipam rirs delete no content response
func (o *IpamRirsDeleteNoContent) Code() int {
	return 204
}

func (o *IpamRirsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipamRirsDeleteNoContent", 204)
}

func (o *IpamRirsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipamRirsDeleteNoContent", 204)
}

func (o *IpamRirsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRirsDeleteDefault creates a IpamRirsDeleteDefault with default headers values
func NewIpamRirsDeleteDefault(code int) *IpamRirsDeleteDefault {
	return &IpamRirsDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamRirsDeleteDefault describes a response with status code -1, with default header values.

IpamRirsDeleteDefault ipam rirs delete default
*/
type IpamRirsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam rirs delete default response has a 2xx status code
func (o *IpamRirsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs delete default response has a 3xx status code
func (o *IpamRirsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs delete default response has a 4xx status code
func (o *IpamRirsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs delete default response has a 5xx status code
func (o *IpamRirsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs delete default response a status code equal to that given
func (o *IpamRirsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam rirs delete default response
func (o *IpamRirsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamRirsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipam_rirs_delete default %s", o._statusCode, payload)
}

func (o *IpamRirsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ipam/rirs/{id}/][%d] ipam_rirs_delete default %s", o._statusCode, payload)
}

func (o *IpamRirsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
