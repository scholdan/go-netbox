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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimManufacturersPartialUpdateReader is a Reader for the DcimManufacturersPartialUpdate structure.
type DcimManufacturersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersPartialUpdateOK creates a DcimManufacturersPartialUpdateOK with default headers values
func NewDcimManufacturersPartialUpdateOK() *DcimManufacturersPartialUpdateOK {
	return &DcimManufacturersPartialUpdateOK{}
}

/*
DcimManufacturersPartialUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersPartialUpdateOK dcim manufacturers partial update o k
*/
type DcimManufacturersPartialUpdateOK struct {
	Payload *models.Manufacturer
}

// IsSuccess returns true when this dcim manufacturers partial update o k response has a 2xx status code
func (o *DcimManufacturersPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers partial update o k response has a 3xx status code
func (o *DcimManufacturersPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers partial update o k response has a 4xx status code
func (o *DcimManufacturersPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers partial update o k response has a 5xx status code
func (o *DcimManufacturersPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers partial update o k response a status code equal to that given
func (o *DcimManufacturersPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim manufacturers partial update o k response
func (o *DcimManufacturersPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimManufacturersPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK %s", 200, payload)
}

func (o *DcimManufacturersPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK %s", 200, payload)
}

func (o *DcimManufacturersPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersPartialUpdateDefault creates a DcimManufacturersPartialUpdateDefault with default headers values
func NewDcimManufacturersPartialUpdateDefault(code int) *DcimManufacturersPartialUpdateDefault {
	return &DcimManufacturersPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimManufacturersPartialUpdateDefault describes a response with status code -1, with default header values.

DcimManufacturersPartialUpdateDefault dcim manufacturers partial update default
*/
type DcimManufacturersPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim manufacturers partial update default response has a 2xx status code
func (o *DcimManufacturersPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim manufacturers partial update default response has a 3xx status code
func (o *DcimManufacturersPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim manufacturers partial update default response has a 4xx status code
func (o *DcimManufacturersPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim manufacturers partial update default response has a 5xx status code
func (o *DcimManufacturersPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim manufacturers partial update default response a status code equal to that given
func (o *DcimManufacturersPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim manufacturers partial update default response
func (o *DcimManufacturersPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcim_manufacturers_partial_update default %s", o._statusCode, payload)
}

func (o *DcimManufacturersPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcim_manufacturers_partial_update default %s", o._statusCode, payload)
}

func (o *DcimManufacturersPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
