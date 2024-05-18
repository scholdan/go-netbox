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

// DcimPowerPortsDeleteReader is a Reader for the DcimPowerPortsDelete structure.
type DcimPowerPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsDeleteNoContent creates a DcimPowerPortsDeleteNoContent with default headers values
func NewDcimPowerPortsDeleteNoContent() *DcimPowerPortsDeleteNoContent {
	return &DcimPowerPortsDeleteNoContent{}
}

/*
DcimPowerPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortsDeleteNoContent dcim power ports delete no content
*/
type DcimPowerPortsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power ports delete no content response has a 2xx status code
func (o *DcimPowerPortsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports delete no content response has a 3xx status code
func (o *DcimPowerPortsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports delete no content response has a 4xx status code
func (o *DcimPowerPortsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports delete no content response has a 5xx status code
func (o *DcimPowerPortsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports delete no content response a status code equal to that given
func (o *DcimPowerPortsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim power ports delete no content response
func (o *DcimPowerPortsDeleteNoContent) Code() int {
	return 204
}

func (o *DcimPowerPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcimPowerPortsDeleteNoContent", 204)
}

func (o *DcimPowerPortsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcimPowerPortsDeleteNoContent", 204)
}

func (o *DcimPowerPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerPortsDeleteDefault creates a DcimPowerPortsDeleteDefault with default headers values
func NewDcimPowerPortsDeleteDefault(code int) *DcimPowerPortsDeleteDefault {
	return &DcimPowerPortsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsDeleteDefault describes a response with status code -1, with default header values.

DcimPowerPortsDeleteDefault dcim power ports delete default
*/
type DcimPowerPortsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power ports delete default response has a 2xx status code
func (o *DcimPowerPortsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports delete default response has a 3xx status code
func (o *DcimPowerPortsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports delete default response has a 4xx status code
func (o *DcimPowerPortsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports delete default response has a 5xx status code
func (o *DcimPowerPortsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports delete default response a status code equal to that given
func (o *DcimPowerPortsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power ports delete default response
func (o *DcimPowerPortsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcim_power-ports_delete default %s", o._statusCode, payload)
}

func (o *DcimPowerPortsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcim_power-ports_delete default %s", o._statusCode, payload)
}

func (o *DcimPowerPortsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
