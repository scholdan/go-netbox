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

package dcim

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

// DcimDeviceRolesUpdateReader is a Reader for the DcimDeviceRolesUpdate structure.
type DcimDeviceRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesUpdateOK creates a DcimDeviceRolesUpdateOK with default headers values
func NewDcimDeviceRolesUpdateOK() *DcimDeviceRolesUpdateOK {
	return &DcimDeviceRolesUpdateOK{}
}

/*
DcimDeviceRolesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesUpdateOK dcim device roles update o k
*/
type DcimDeviceRolesUpdateOK struct {
	Payload *models.DeviceRole
}

// IsSuccess returns true when this dcim device roles update o k response has a 2xx status code
func (o *DcimDeviceRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles update o k response has a 3xx status code
func (o *DcimDeviceRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles update o k response has a 4xx status code
func (o *DcimDeviceRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles update o k response has a 5xx status code
func (o *DcimDeviceRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles update o k response a status code equal to that given
func (o *DcimDeviceRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim device roles update o k response
func (o *DcimDeviceRolesUpdateOK) Code() int {
	return 200
}

func (o *DcimDeviceRolesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-roles/{id}/][%d] dcimDeviceRolesUpdateOK %s", 200, payload)
}

func (o *DcimDeviceRolesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-roles/{id}/][%d] dcimDeviceRolesUpdateOK %s", 200, payload)
}

func (o *DcimDeviceRolesUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesUpdateDefault creates a DcimDeviceRolesUpdateDefault with default headers values
func NewDcimDeviceRolesUpdateDefault(code int) *DcimDeviceRolesUpdateDefault {
	return &DcimDeviceRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceRolesUpdateDefault dcim device roles update default
*/
type DcimDeviceRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device roles update default response has a 2xx status code
func (o *DcimDeviceRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles update default response has a 3xx status code
func (o *DcimDeviceRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles update default response has a 4xx status code
func (o *DcimDeviceRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles update default response has a 5xx status code
func (o *DcimDeviceRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles update default response a status code equal to that given
func (o *DcimDeviceRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device roles update default response
func (o *DcimDeviceRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-roles/{id}/][%d] dcim_device-roles_update default %s", o._statusCode, payload)
}

func (o *DcimDeviceRolesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/device-roles/{id}/][%d] dcim_device-roles_update default %s", o._statusCode, payload)
}

func (o *DcimDeviceRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
