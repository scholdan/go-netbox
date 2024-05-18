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

// IpamVlanGroupsDeleteReader is a Reader for the IpamVlanGroupsDelete structure.
type IpamVlanGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlanGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsDeleteNoContent creates a IpamVlanGroupsDeleteNoContent with default headers values
func NewIpamVlanGroupsDeleteNoContent() *IpamVlanGroupsDeleteNoContent {
	return &IpamVlanGroupsDeleteNoContent{}
}

/*
IpamVlanGroupsDeleteNoContent describes a response with status code 204, with default header values.

IpamVlanGroupsDeleteNoContent ipam vlan groups delete no content
*/
type IpamVlanGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam vlan groups delete no content response has a 2xx status code
func (o *IpamVlanGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups delete no content response has a 3xx status code
func (o *IpamVlanGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups delete no content response has a 4xx status code
func (o *IpamVlanGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups delete no content response has a 5xx status code
func (o *IpamVlanGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups delete no content response a status code equal to that given
func (o *IpamVlanGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the ipam vlan groups delete no content response
func (o *IpamVlanGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *IpamVlanGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsDeleteNoContent", 204)
}

func (o *IpamVlanGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsDeleteNoContent", 204)
}

func (o *IpamVlanGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVlanGroupsDeleteDefault creates a IpamVlanGroupsDeleteDefault with default headers values
func NewIpamVlanGroupsDeleteDefault(code int) *IpamVlanGroupsDeleteDefault {
	return &IpamVlanGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsDeleteDefault describes a response with status code -1, with default header values.

IpamVlanGroupsDeleteDefault ipam vlan groups delete default
*/
type IpamVlanGroupsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vlan groups delete default response has a 2xx status code
func (o *IpamVlanGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups delete default response has a 3xx status code
func (o *IpamVlanGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups delete default response has a 4xx status code
func (o *IpamVlanGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups delete default response has a 5xx status code
func (o *IpamVlanGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups delete default response a status code equal to that given
func (o *IpamVlanGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vlan groups delete default response
func (o *IpamVlanGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlanGroupsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_delete default %s", o._statusCode, payload)
}

func (o *IpamVlanGroupsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_delete default %s", o._statusCode, payload)
}

func (o *IpamVlanGroupsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
