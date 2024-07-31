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

	"github.com/scholdan/go-netbox/netbox/models"
)

// DcimVirtualChassisUpdateReader is a Reader for the DcimVirtualChassisUpdate structure.
type DcimVirtualChassisUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisUpdateOK creates a DcimVirtualChassisUpdateOK with default headers values
func NewDcimVirtualChassisUpdateOK() *DcimVirtualChassisUpdateOK {
	return &DcimVirtualChassisUpdateOK{}
}

/*
DcimVirtualChassisUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisUpdateOK dcim virtual chassis update o k
*/
type DcimVirtualChassisUpdateOK struct {
	Payload *models.VirtualChassis
}

// IsSuccess returns true when this dcim virtual chassis update o k response has a 2xx status code
func (o *DcimVirtualChassisUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis update o k response has a 3xx status code
func (o *DcimVirtualChassisUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis update o k response has a 4xx status code
func (o *DcimVirtualChassisUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis update o k response has a 5xx status code
func (o *DcimVirtualChassisUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis update o k response a status code equal to that given
func (o *DcimVirtualChassisUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim virtual chassis update o k response
func (o *DcimVirtualChassisUpdateOK) Code() int {
	return 200
}

func (o *DcimVirtualChassisUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateOK %s", 200, payload)
}

func (o *DcimVirtualChassisUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateOK %s", 200, payload)
}

func (o *DcimVirtualChassisUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisUpdateDefault creates a DcimVirtualChassisUpdateDefault with default headers values
func NewDcimVirtualChassisUpdateDefault(code int) *DcimVirtualChassisUpdateDefault {
	return &DcimVirtualChassisUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualChassisUpdateDefault describes a response with status code -1, with default header values.

DcimVirtualChassisUpdateDefault dcim virtual chassis update default
*/
type DcimVirtualChassisUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim virtual chassis update default response has a 2xx status code
func (o *DcimVirtualChassisUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual chassis update default response has a 3xx status code
func (o *DcimVirtualChassisUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual chassis update default response has a 4xx status code
func (o *DcimVirtualChassisUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual chassis update default response has a 5xx status code
func (o *DcimVirtualChassisUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual chassis update default response a status code equal to that given
func (o *DcimVirtualChassisUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim virtual chassis update default response
func (o *DcimVirtualChassisUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcim_virtual-chassis_update default %s", o._statusCode, payload)
}

func (o *DcimVirtualChassisUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcim_virtual-chassis_update default %s", o._statusCode, payload)
}

func (o *DcimVirtualChassisUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
