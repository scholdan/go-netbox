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
)

// DcimDevicesDeleteReader is a Reader for the DcimDevicesDelete structure.
type DcimDevicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesDeleteNoContent creates a DcimDevicesDeleteNoContent with default headers values
func NewDcimDevicesDeleteNoContent() *DcimDevicesDeleteNoContent {
	return &DcimDevicesDeleteNoContent{}
}

/*
DcimDevicesDeleteNoContent describes a response with status code 204, with default header values.

DcimDevicesDeleteNoContent dcim devices delete no content
*/
type DcimDevicesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim devices delete no content response has a 2xx status code
func (o *DcimDevicesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices delete no content response has a 3xx status code
func (o *DcimDevicesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices delete no content response has a 4xx status code
func (o *DcimDevicesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices delete no content response has a 5xx status code
func (o *DcimDevicesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices delete no content response a status code equal to that given
func (o *DcimDevicesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim devices delete no content response
func (o *DcimDevicesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimDevicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcimDevicesDeleteNoContent", 204)
}

func (o *DcimDevicesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcimDevicesDeleteNoContent", 204)
}

func (o *DcimDevicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDevicesDeleteDefault creates a DcimDevicesDeleteDefault with default headers values
func NewDcimDevicesDeleteDefault(code int) *DcimDevicesDeleteDefault {
	return &DcimDevicesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesDeleteDefault describes a response with status code -1, with default header values.

DcimDevicesDeleteDefault dcim devices delete default
*/
type DcimDevicesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim devices delete default response has a 2xx status code
func (o *DcimDevicesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices delete default response has a 3xx status code
func (o *DcimDevicesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices delete default response has a 4xx status code
func (o *DcimDevicesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices delete default response has a 5xx status code
func (o *DcimDevicesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices delete default response a status code equal to that given
func (o *DcimDevicesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim devices delete default response
func (o *DcimDevicesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcim_devices_delete default %s", o._statusCode, payload)
}

func (o *DcimDevicesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcim_devices_delete default %s", o._statusCode, payload)
}

func (o *DcimDevicesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
