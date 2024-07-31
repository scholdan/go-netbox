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

	"github.com/scholdan/go-netbox/netbox/models"
)

// VpnTunnelTerminationsUpdateReader is a Reader for the VpnTunnelTerminationsUpdate structure.
type VpnTunnelTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVpnTunnelTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVpnTunnelTerminationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVpnTunnelTerminationsUpdateOK creates a VpnTunnelTerminationsUpdateOK with default headers values
func NewVpnTunnelTerminationsUpdateOK() *VpnTunnelTerminationsUpdateOK {
	return &VpnTunnelTerminationsUpdateOK{}
}

/*
VpnTunnelTerminationsUpdateOK describes a response with status code 200, with default header values.

VpnTunnelTerminationsUpdateOK vpn tunnel terminations update o k
*/
type VpnTunnelTerminationsUpdateOK struct {
	Payload *models.TunnelTermination
}

// IsSuccess returns true when this vpn tunnel terminations update o k response has a 2xx status code
func (o *VpnTunnelTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel terminations update o k response has a 3xx status code
func (o *VpnTunnelTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel terminations update o k response has a 4xx status code
func (o *VpnTunnelTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel terminations update o k response has a 5xx status code
func (o *VpnTunnelTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel terminations update o k response a status code equal to that given
func (o *VpnTunnelTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vpn tunnel terminations update o k response
func (o *VpnTunnelTerminationsUpdateOK) Code() int {
	return 200
}

func (o *VpnTunnelTerminationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-terminations/{id}/][%d] vpnTunnelTerminationsUpdateOK %s", 200, payload)
}

func (o *VpnTunnelTerminationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-terminations/{id}/][%d] vpnTunnelTerminationsUpdateOK %s", 200, payload)
}

func (o *VpnTunnelTerminationsUpdateOK) GetPayload() *models.TunnelTermination {
	return o.Payload
}

func (o *VpnTunnelTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TunnelTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVpnTunnelTerminationsUpdateDefault creates a VpnTunnelTerminationsUpdateDefault with default headers values
func NewVpnTunnelTerminationsUpdateDefault(code int) *VpnTunnelTerminationsUpdateDefault {
	return &VpnTunnelTerminationsUpdateDefault{
		_statusCode: code,
	}
}

/*
VpnTunnelTerminationsUpdateDefault describes a response with status code -1, with default header values.

VpnTunnelTerminationsUpdateDefault vpn tunnel terminations update default
*/
type VpnTunnelTerminationsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this vpn tunnel terminations update default response has a 2xx status code
func (o *VpnTunnelTerminationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vpn tunnel terminations update default response has a 3xx status code
func (o *VpnTunnelTerminationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vpn tunnel terminations update default response has a 4xx status code
func (o *VpnTunnelTerminationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vpn tunnel terminations update default response has a 5xx status code
func (o *VpnTunnelTerminationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vpn tunnel terminations update default response a status code equal to that given
func (o *VpnTunnelTerminationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vpn tunnel terminations update default response
func (o *VpnTunnelTerminationsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VpnTunnelTerminationsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-terminations/{id}/][%d] vpn_tunnel-terminations_update default %s", o._statusCode, payload)
}

func (o *VpnTunnelTerminationsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-terminations/{id}/][%d] vpn_tunnel-terminations_update default %s", o._statusCode, payload)
}

func (o *VpnTunnelTerminationsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VpnTunnelTerminationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
