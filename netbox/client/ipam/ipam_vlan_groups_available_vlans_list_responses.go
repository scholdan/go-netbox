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

// IpamVlanGroupsAvailableVlansListReader is a Reader for the IpamVlanGroupsAvailableVlansList structure.
type IpamVlanGroupsAvailableVlansListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsAvailableVlansListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsAvailableVlansListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsAvailableVlansListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsAvailableVlansListOK creates a IpamVlanGroupsAvailableVlansListOK with default headers values
func NewIpamVlanGroupsAvailableVlansListOK() *IpamVlanGroupsAvailableVlansListOK {
	return &IpamVlanGroupsAvailableVlansListOK{}
}

/*
IpamVlanGroupsAvailableVlansListOK describes a response with status code 200, with default header values.

IpamVlanGroupsAvailableVlansListOK ipam vlan groups available vlans list o k
*/
type IpamVlanGroupsAvailableVlansListOK struct {
	Payload []*models.AvailableVLAN
}

// IsSuccess returns true when this ipam vlan groups available vlans list o k response has a 2xx status code
func (o *IpamVlanGroupsAvailableVlansListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups available vlans list o k response has a 3xx status code
func (o *IpamVlanGroupsAvailableVlansListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups available vlans list o k response has a 4xx status code
func (o *IpamVlanGroupsAvailableVlansListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups available vlans list o k response has a 5xx status code
func (o *IpamVlanGroupsAvailableVlansListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups available vlans list o k response a status code equal to that given
func (o *IpamVlanGroupsAvailableVlansListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vlan groups available vlans list o k response
func (o *IpamVlanGroupsAvailableVlansListOK) Code() int {
	return 200
}

func (o *IpamVlanGroupsAvailableVlansListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/available-vlans/][%d] ipamVlanGroupsAvailableVlansListOK %s", 200, payload)
}

func (o *IpamVlanGroupsAvailableVlansListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/available-vlans/][%d] ipamVlanGroupsAvailableVlansListOK %s", 200, payload)
}

func (o *IpamVlanGroupsAvailableVlansListOK) GetPayload() []*models.AvailableVLAN {
	return o.Payload
}

func (o *IpamVlanGroupsAvailableVlansListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsAvailableVlansListDefault creates a IpamVlanGroupsAvailableVlansListDefault with default headers values
func NewIpamVlanGroupsAvailableVlansListDefault(code int) *IpamVlanGroupsAvailableVlansListDefault {
	return &IpamVlanGroupsAvailableVlansListDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsAvailableVlansListDefault describes a response with status code -1, with default header values.

IpamVlanGroupsAvailableVlansListDefault ipam vlan groups available vlans list default
*/
type IpamVlanGroupsAvailableVlansListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam vlan groups available vlans list default response has a 2xx status code
func (o *IpamVlanGroupsAvailableVlansListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups available vlans list default response has a 3xx status code
func (o *IpamVlanGroupsAvailableVlansListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups available vlans list default response has a 4xx status code
func (o *IpamVlanGroupsAvailableVlansListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups available vlans list default response has a 5xx status code
func (o *IpamVlanGroupsAvailableVlansListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups available vlans list default response a status code equal to that given
func (o *IpamVlanGroupsAvailableVlansListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam vlan groups available vlans list default response
func (o *IpamVlanGroupsAvailableVlansListDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlanGroupsAvailableVlansListDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/available-vlans/][%d] ipam_vlan-groups_available-vlans_list default %s", o._statusCode, payload)
}

func (o *IpamVlanGroupsAvailableVlansListDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/available-vlans/][%d] ipam_vlan-groups_available-vlans_list default %s", o._statusCode, payload)
}

func (o *IpamVlanGroupsAvailableVlansListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsAvailableVlansListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
