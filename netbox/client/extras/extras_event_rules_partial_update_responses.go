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

package extras

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

// ExtrasEventRulesPartialUpdateReader is a Reader for the ExtrasEventRulesPartialUpdate structure.
type ExtrasEventRulesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasEventRulesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasEventRulesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasEventRulesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasEventRulesPartialUpdateOK creates a ExtrasEventRulesPartialUpdateOK with default headers values
func NewExtrasEventRulesPartialUpdateOK() *ExtrasEventRulesPartialUpdateOK {
	return &ExtrasEventRulesPartialUpdateOK{}
}

/*
ExtrasEventRulesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasEventRulesPartialUpdateOK extras event rules partial update o k
*/
type ExtrasEventRulesPartialUpdateOK struct {
	Payload *models.EventRule
}

// IsSuccess returns true when this extras event rules partial update o k response has a 2xx status code
func (o *ExtrasEventRulesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras event rules partial update o k response has a 3xx status code
func (o *ExtrasEventRulesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras event rules partial update o k response has a 4xx status code
func (o *ExtrasEventRulesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras event rules partial update o k response has a 5xx status code
func (o *ExtrasEventRulesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras event rules partial update o k response a status code equal to that given
func (o *ExtrasEventRulesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras event rules partial update o k response
func (o *ExtrasEventRulesPartialUpdateOK) Code() int {
	return 200
}

func (o *ExtrasEventRulesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/event-rules/{id}/][%d] extrasEventRulesPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasEventRulesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/event-rules/{id}/][%d] extrasEventRulesPartialUpdateOK %s", 200, payload)
}

func (o *ExtrasEventRulesPartialUpdateOK) GetPayload() *models.EventRule {
	return o.Payload
}

func (o *ExtrasEventRulesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasEventRulesPartialUpdateDefault creates a ExtrasEventRulesPartialUpdateDefault with default headers values
func NewExtrasEventRulesPartialUpdateDefault(code int) *ExtrasEventRulesPartialUpdateDefault {
	return &ExtrasEventRulesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasEventRulesPartialUpdateDefault describes a response with status code -1, with default header values.

ExtrasEventRulesPartialUpdateDefault extras event rules partial update default
*/
type ExtrasEventRulesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras event rules partial update default response has a 2xx status code
func (o *ExtrasEventRulesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras event rules partial update default response has a 3xx status code
func (o *ExtrasEventRulesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras event rules partial update default response has a 4xx status code
func (o *ExtrasEventRulesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras event rules partial update default response has a 5xx status code
func (o *ExtrasEventRulesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras event rules partial update default response a status code equal to that given
func (o *ExtrasEventRulesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras event rules partial update default response
func (o *ExtrasEventRulesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasEventRulesPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/event-rules/{id}/][%d] extras_event_rules_partial_update default %s", o._statusCode, payload)
}

func (o *ExtrasEventRulesPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /extras/event-rules/{id}/][%d] extras_event_rules_partial_update default %s", o._statusCode, payload)
}

func (o *ExtrasEventRulesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasEventRulesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
