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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsTraceReader is a Reader for the CircuitsCircuitTerminationsTrace structure.
type CircuitsCircuitTerminationsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsTraceOK creates a CircuitsCircuitTerminationsTraceOK with default headers values
func NewCircuitsCircuitTerminationsTraceOK() *CircuitsCircuitTerminationsTraceOK {
	return &CircuitsCircuitTerminationsTraceOK{}
}

/* CircuitsCircuitTerminationsTraceOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsTraceOK circuits circuit terminations trace o k
*/
type CircuitsCircuitTerminationsTraceOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsTraceOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/trace/][%d] circuitsCircuitTerminationsTraceOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsTraceOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsTraceDefault creates a CircuitsCircuitTerminationsTraceDefault with default headers values
func NewCircuitsCircuitTerminationsTraceDefault(code int) *CircuitsCircuitTerminationsTraceDefault {
	return &CircuitsCircuitTerminationsTraceDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitTerminationsTraceDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsTraceDefault circuits circuit terminations trace default
*/
type CircuitsCircuitTerminationsTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations trace default response
func (o *CircuitsCircuitTerminationsTraceDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsTraceDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/trace/][%d] circuits_circuit-terminations_trace default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitTerminationsTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
