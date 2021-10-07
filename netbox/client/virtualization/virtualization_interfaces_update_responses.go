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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationInterfacesUpdateReader is a Reader for the VirtualizationInterfacesUpdate structure.
type VirtualizationInterfacesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesUpdateOK creates a VirtualizationInterfacesUpdateOK with default headers values
func NewVirtualizationInterfacesUpdateOK() *VirtualizationInterfacesUpdateOK {
	return &VirtualizationInterfacesUpdateOK{}
}

/* VirtualizationInterfacesUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesUpdateOK virtualization interfaces update o k
*/
type VirtualizationInterfacesUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/interfaces/{id}/][%d] virtualizationInterfacesUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesUpdateDefault creates a VirtualizationInterfacesUpdateDefault with default headers values
func NewVirtualizationInterfacesUpdateDefault(code int) *VirtualizationInterfacesUpdateDefault {
	return &VirtualizationInterfacesUpdateDefault{
		_statusCode: code,
	}
}

/* VirtualizationInterfacesUpdateDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesUpdateDefault virtualization interfaces update default
*/
type VirtualizationInterfacesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces update default response
func (o *VirtualizationInterfacesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/interfaces/{id}/][%d] virtualization_interfaces_update default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationInterfacesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
