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
)

// VpnTunnelTerminationsDeleteReader is a Reader for the VpnTunnelTerminationsDelete structure.
type VpnTunnelTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVpnTunnelTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelTerminationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelTerminationsDeleteNoContent creates a VpnTunnelTerminationsDeleteNoContent with default headers values
func NewVpnTunnelTerminationsDeleteNoContent() *VpnTunnelTerminationsDeleteNoContent {
	return &VpnTunnelTerminationsDeleteNoContent{}
}

/*
VpnTunnelTerminationsDeleteNoContent describes a response with status code 204, with default header values.

VpnTunnelTerminationsDeleteNoContent vpn tunnel terminations delete no content
*/
type VpnTunnelTerminationsDeleteNoContent struct {
}

// IsSuccess returns true when this vpn tunnel terminations delete no content response has a 2xx status code
func (o *VpnTunnelTerminationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel terminations delete no content response has a 3xx status code
func (o *VpnTunnelTerminationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel terminations delete no content response has a 4xx status code
func (o *VpnTunnelTerminationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel terminations delete no content response has a 5xx status code
func (o *VpnTunnelTerminationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel terminations delete no content response a status code equal to that given
func (o *VpnTunnelTerminationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the vpn tunnel terminations delete no content response
func (o *VpnTunnelTerminationsDeleteNoContent) Code() int {
	return 204
}

func (o *VpnTunnelTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpn/tunnel-terminations/{id}/][%d] vpnTunnelTerminationsDeleteNoContent", 204)
}

func (o *VpnTunnelTerminationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /vpn/tunnel-terminations/{id}/][%d] vpnTunnelTerminationsDeleteNoContent", 204)
}

func (o *VpnTunnelTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVpnTunnelTerminationsDeleteDefault creates a VpnTunnelTerminationsDeleteDefault with default headers values
func NewVpnTunnelTerminationsDeleteDefault(code int) *VpnTunnelTerminationsDeleteDefault {
	return &VpnTunnelTerminationsDeleteDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelTerminationsDeleteDefault describes a response with status code -1, with default header values.

VpnTunnelTerminationsDeleteDefault vpn tunnel terminations delete default
*/
type VpnTunnelTerminationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnel terminations delete default response has a 2xx status code
func (o *VpnTunnelTerminationsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnel terminations delete default response has a 3xx status code
func (o *VpnTunnelTerminationsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnel terminations delete default response has a 4xx status code
func (o *VpnTunnelTerminationsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnel terminations delete default response has a 5xx status code
func (o *VpnTunnelTerminationsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnel terminations delete default response a status code equal to that given
func (o *VpnTunnelTerminationsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnel terminations delete default response
func (o *VpnTunnelTerminationsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelTerminationsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /vpn/tunnel-terminations/{id}/][%d] vpn_tunnel-terminations_delete default %s", o._statusCode, payload)
}

func (o *VpnTunnelTerminationsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /vpn/tunnel-terminations/{id}/][%d] vpn_tunnel-terminations_delete default %s", o._statusCode, payload)
}

func (o *VpnTunnelTerminationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelTerminationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
