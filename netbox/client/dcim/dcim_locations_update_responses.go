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

// DcimLocationsUpdateReader is a Reader for the DcimLocationsUpdate structure.
type DcimLocationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsUpdateOK creates a DcimLocationsUpdateOK with default headers values
func NewDcimLocationsUpdateOK() *DcimLocationsUpdateOK {
	return &DcimLocationsUpdateOK{}
}

/*
DcimLocationsUpdateOK describes a response with status code 200, with default header values.

DcimLocationsUpdateOK dcim locations update o k
*/
type DcimLocationsUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations update o k response has a 2xx status code
func (o *DcimLocationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations update o k response has a 3xx status code
func (o *DcimLocationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations update o k response has a 4xx status code
func (o *DcimLocationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations update o k response has a 5xx status code
func (o *DcimLocationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations update o k response a status code equal to that given
func (o *DcimLocationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim locations update o k response
func (o *DcimLocationsUpdateOK) Code() int {
	return 200
}

func (o *DcimLocationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK %s", 200, payload)
}

func (o *DcimLocationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK %s", 200, payload)
}

func (o *DcimLocationsUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsUpdateDefault creates a DcimLocationsUpdateDefault with default headers values
func NewDcimLocationsUpdateDefault(code int) *DcimLocationsUpdateDefault {
	return &DcimLocationsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsUpdateDefault describes a response with status code -1, with default header values.

DcimLocationsUpdateDefault dcim locations update default
*/
type DcimLocationsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim locations update default response has a 2xx status code
func (o *DcimLocationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations update default response has a 3xx status code
func (o *DcimLocationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations update default response has a 4xx status code
func (o *DcimLocationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations update default response has a 5xx status code
func (o *DcimLocationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations update default response a status code equal to that given
func (o *DcimLocationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim locations update default response
func (o *DcimLocationsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimLocationsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcim_locations_update default %s", o._statusCode, payload)
}

func (o *DcimLocationsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcim_locations_update default %s", o._statusCode, payload)
}

func (o *DcimLocationsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
