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

// DcimRacksUpdateReader is a Reader for the DcimRacksUpdate structure.
type DcimRacksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksUpdateOK creates a DcimRacksUpdateOK with default headers values
func NewDcimRacksUpdateOK() *DcimRacksUpdateOK {
	return &DcimRacksUpdateOK{}
}

/*
DcimRacksUpdateOK describes a response with status code 200, with default header values.

DcimRacksUpdateOK dcim racks update o k
*/
type DcimRacksUpdateOK struct {
	Payload *models.Rack
}

// IsSuccess returns true when this dcim racks update o k response has a 2xx status code
func (o *DcimRacksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks update o k response has a 3xx status code
func (o *DcimRacksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks update o k response has a 4xx status code
func (o *DcimRacksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks update o k response has a 5xx status code
func (o *DcimRacksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks update o k response a status code equal to that given
func (o *DcimRacksUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim racks update o k response
func (o *DcimRacksUpdateOK) Code() int {
	return 200
}

func (o *DcimRacksUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/racks/{id}/][%d] dcimRacksUpdateOK %s", 200, payload)
}

func (o *DcimRacksUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/racks/{id}/][%d] dcimRacksUpdateOK %s", 200, payload)
}

func (o *DcimRacksUpdateOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksUpdateDefault creates a DcimRacksUpdateDefault with default headers values
func NewDcimRacksUpdateDefault(code int) *DcimRacksUpdateDefault {
	return &DcimRacksUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRacksUpdateDefault describes a response with status code -1, with default header values.

DcimRacksUpdateDefault dcim racks update default
*/
type DcimRacksUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim racks update default response has a 2xx status code
func (o *DcimRacksUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks update default response has a 3xx status code
func (o *DcimRacksUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks update default response has a 4xx status code
func (o *DcimRacksUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks update default response has a 5xx status code
func (o *DcimRacksUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks update default response a status code equal to that given
func (o *DcimRacksUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim racks update default response
func (o *DcimRacksUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRacksUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/racks/{id}/][%d] dcim_racks_update default %s", o._statusCode, payload)
}

func (o *DcimRacksUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/racks/{id}/][%d] dcim_racks_update default %s", o._statusCode, payload)
}

func (o *DcimRacksUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
