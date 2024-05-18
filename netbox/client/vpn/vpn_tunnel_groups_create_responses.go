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

package vpn

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

// VpnTunnelGroupsCreateReader is a Reader for the VpnTunnelGroupsCreate structure.
type VpnTunnelGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVpnTunnelGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelGroupsCreateCreated creates a VpnTunnelGroupsCreateCreated with default headers values
func NewVpnTunnelGroupsCreateCreated() *VpnTunnelGroupsCreateCreated {
	return &VpnTunnelGroupsCreateCreated{}
}

/*
VpnTunnelGroupsCreateCreated describes a response with status code 201, with default header values.

VpnTunnelGroupsCreateCreated vpn tunnel groups create created
*/
type VpnTunnelGroupsCreateCreated struct {
	Payload *models.TunnelGroup
}

// IsSuccess returns true when this vpn tunnel groups create created response has a 2xx status code
func (o *VpnTunnelGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel groups create created response has a 3xx status code
func (o *VpnTunnelGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel groups create created response has a 4xx status code
func (o *VpnTunnelGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel groups create created response has a 5xx status code
func (o *VpnTunnelGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel groups create created response a status code equal to that given
func (o *VpnTunnelGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the vpn tunnel groups create created response
func (o *VpnTunnelGroupsCreateCreated) Code() int {
	return 201
}

func (o *VpnTunnelGroupsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnel-groups/][%d] vpnTunnelGroupsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelGroupsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnel-groups/][%d] vpnTunnelGroupsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelGroupsCreateCreated) GetPayload() *models.TunnelGroup {
	return o.Payload
}

func (o *VpnTunnelGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TunnelGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVpnTunnelGroupsCreateDefault creates a VpnTunnelGroupsCreateDefault with default headers values
func NewVpnTunnelGroupsCreateDefault(code int) *VpnTunnelGroupsCreateDefault {
	return &VpnTunnelGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelGroupsCreateDefault describes a response with status code -1, with default header values.

VpnTunnelGroupsCreateDefault vpn tunnel groups create default
*/
type VpnTunnelGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnel groups create default response has a 2xx status code
func (o *VpnTunnelGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnel groups create default response has a 3xx status code
func (o *VpnTunnelGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnel groups create default response has a 4xx status code
func (o *VpnTunnelGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnel groups create default response has a 5xx status code
func (o *VpnTunnelGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnel groups create default response a status code equal to that given
func (o *VpnTunnelGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnel groups create default response
func (o *VpnTunnelGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelGroupsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnel-groups/][%d] vpn_tunnel-groups_create default %s", o._statusCode, payload)
}

func (o *VpnTunnelGroupsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnel-groups/][%d] vpn_tunnel-groups_create default %s", o._statusCode, payload)
}

func (o *VpnTunnelGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
