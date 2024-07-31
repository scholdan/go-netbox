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

package virtualization

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

// VirtualizationInterfacesReadReader is a Reader for the VirtualizationInterfacesRead structure.
type VirtualizationInterfacesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesReadOK creates a VirtualizationInterfacesReadOK with default headers values
func NewVirtualizationInterfacesReadOK() *VirtualizationInterfacesReadOK {
	return &VirtualizationInterfacesReadOK{}
}

/*
VirtualizationInterfacesReadOK describes a response with status code 200, with default header values.

VirtualizationInterfacesReadOK virtualization interfaces read o k
*/
type VirtualizationInterfacesReadOK struct {
	Payload *models.VMInterface
}

// IsSuccess returns true when this virtualization interfaces read o k response has a 2xx status code
func (o *VirtualizationInterfacesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization interfaces read o k response has a 3xx status code
func (o *VirtualizationInterfacesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization interfaces read o k response has a 4xx status code
func (o *VirtualizationInterfacesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization interfaces read o k response has a 5xx status code
func (o *VirtualizationInterfacesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization interfaces read o k response a status code equal to that given
func (o *VirtualizationInterfacesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization interfaces read o k response
func (o *VirtualizationInterfacesReadOK) Code() int {
	return 200
}

func (o *VirtualizationInterfacesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/interfaces/{id}/][%d] virtualizationInterfacesReadOK %s", 200, payload)
}

func (o *VirtualizationInterfacesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/interfaces/{id}/][%d] virtualizationInterfacesReadOK %s", 200, payload)
}

func (o *VirtualizationInterfacesReadOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesReadDefault creates a VirtualizationInterfacesReadDefault with default headers values
func NewVirtualizationInterfacesReadDefault(code int) *VirtualizationInterfacesReadDefault {
	return &VirtualizationInterfacesReadDefault{
		_statusCode: code,
	}
}

/*
VirtualizationInterfacesReadDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesReadDefault virtualization interfaces read default
*/
type VirtualizationInterfacesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization interfaces read default response has a 2xx status code
func (o *VirtualizationInterfacesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization interfaces read default response has a 3xx status code
func (o *VirtualizationInterfacesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization interfaces read default response has a 4xx status code
func (o *VirtualizationInterfacesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization interfaces read default response has a 5xx status code
func (o *VirtualizationInterfacesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization interfaces read default response a status code equal to that given
func (o *VirtualizationInterfacesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization interfaces read default response
func (o *VirtualizationInterfacesReadDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/interfaces/{id}/][%d] virtualization_interfaces_read default %s", o._statusCode, payload)
}

func (o *VirtualizationInterfacesReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/interfaces/{id}/][%d] virtualization_interfaces_read default %s", o._statusCode, payload)
}

func (o *VirtualizationInterfacesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
