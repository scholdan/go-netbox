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

// DcimDeviceTypesDeleteReader is a Reader for the DcimDeviceTypesDelete structure.
type DcimDeviceTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesDeleteNoContent creates a DcimDeviceTypesDeleteNoContent with default headers values
func NewDcimDeviceTypesDeleteNoContent() *DcimDeviceTypesDeleteNoContent {
	return &DcimDeviceTypesDeleteNoContent{}
}

/*
DcimDeviceTypesDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceTypesDeleteNoContent dcim device types delete no content
*/
type DcimDeviceTypesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device types delete no content response has a 2xx status code
func (o *DcimDeviceTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types delete no content response has a 3xx status code
func (o *DcimDeviceTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types delete no content response has a 4xx status code
func (o *DcimDeviceTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types delete no content response has a 5xx status code
func (o *DcimDeviceTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types delete no content response a status code equal to that given
func (o *DcimDeviceTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim device types delete no content response
func (o *DcimDeviceTypesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimDeviceTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcimDeviceTypesDeleteNoContent", 204)
}

func (o *DcimDeviceTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcimDeviceTypesDeleteNoContent", 204)
}

func (o *DcimDeviceTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceTypesDeleteDefault creates a DcimDeviceTypesDeleteDefault with default headers values
func NewDcimDeviceTypesDeleteDefault(code int) *DcimDeviceTypesDeleteDefault {
	return &DcimDeviceTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesDeleteDefault describes a response with status code -1, with default header values.

DcimDeviceTypesDeleteDefault dcim device types delete default
*/
type DcimDeviceTypesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim device types delete default response has a 2xx status code
func (o *DcimDeviceTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types delete default response has a 3xx status code
func (o *DcimDeviceTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types delete default response has a 4xx status code
func (o *DcimDeviceTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types delete default response has a 5xx status code
func (o *DcimDeviceTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types delete default response a status code equal to that given
func (o *DcimDeviceTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim device types delete default response
func (o *DcimDeviceTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceTypesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcim_device-types_delete default %s", o._statusCode, payload)
}

func (o *DcimDeviceTypesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcim_device-types_delete default %s", o._statusCode, payload)
}

func (o *DcimDeviceTypesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
