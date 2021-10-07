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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimPowerPortsTraceReader is a Reader for the DcimPowerPortsTrace structure.
type DcimPowerPortsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsTraceOK creates a DcimPowerPortsTraceOK with default headers values
func NewDcimPowerPortsTraceOK() *DcimPowerPortsTraceOK {
	return &DcimPowerPortsTraceOK{}
}

/* DcimPowerPortsTraceOK describes a response with status code 200, with default header values.

DcimPowerPortsTraceOK dcim power ports trace o k
*/
type DcimPowerPortsTraceOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcimPowerPortsTraceOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortsTraceOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsTraceDefault creates a DcimPowerPortsTraceDefault with default headers values
func NewDcimPowerPortsTraceDefault(code int) *DcimPowerPortsTraceDefault {
	return &DcimPowerPortsTraceDefault{
		_statusCode: code,
	}
}

/* DcimPowerPortsTraceDefault describes a response with status code -1, with default header values.

DcimPowerPortsTraceDefault dcim power ports trace default
*/
type DcimPowerPortsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power ports trace default response
func (o *DcimPowerPortsTraceDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsTraceDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcim_power-ports_trace default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerPortsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
