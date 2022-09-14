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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimCableTerminationsDeleteReader is a Reader for the DcimCableTerminationsDelete structure.
type DcimCableTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCableTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCableTerminationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsDeleteNoContent creates a DcimCableTerminationsDeleteNoContent with default headers values
func NewDcimCableTerminationsDeleteNoContent() *DcimCableTerminationsDeleteNoContent {
	return &DcimCableTerminationsDeleteNoContent{}
}

/* DcimCableTerminationsDeleteNoContent describes a response with status code 204, with default header values.

DcimCableTerminationsDeleteNoContent dcim cable terminations delete no content
*/
type DcimCableTerminationsDeleteNoContent struct {
}

func (o *DcimCableTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteNoContent ", 204)
}

func (o *DcimCableTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimCableTerminationsDeleteDefault creates a DcimCableTerminationsDeleteDefault with default headers values
func NewDcimCableTerminationsDeleteDefault(code int) *DcimCableTerminationsDeleteDefault {
	return &DcimCableTerminationsDeleteDefault{
		_statusCode: code,
	}
}

/* DcimCableTerminationsDeleteDefault describes a response with status code -1, with default header values.

DcimCableTerminationsDeleteDefault dcim cable terminations delete default
*/
type DcimCableTerminationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cable terminations delete default response
func (o *DcimCableTerminationsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimCableTerminationsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_delete default  %+v", o._statusCode, o.Payload)
}
func (o *DcimCableTerminationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}