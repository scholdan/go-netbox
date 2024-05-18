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

// DcimPowerOutletTemplatesUpdateReader is a Reader for the DcimPowerOutletTemplatesUpdate structure.
type DcimPowerOutletTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesUpdateOK creates a DcimPowerOutletTemplatesUpdateOK with default headers values
func NewDcimPowerOutletTemplatesUpdateOK() *DcimPowerOutletTemplatesUpdateOK {
	return &DcimPowerOutletTemplatesUpdateOK{}
}

/*
DcimPowerOutletTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesUpdateOK dcim power outlet templates update o k
*/
type DcimPowerOutletTemplatesUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

// IsSuccess returns true when this dcim power outlet templates update o k response has a 2xx status code
func (o *DcimPowerOutletTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates update o k response has a 3xx status code
func (o *DcimPowerOutletTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates update o k response has a 4xx status code
func (o *DcimPowerOutletTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates update o k response has a 5xx status code
func (o *DcimPowerOutletTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates update o k response a status code equal to that given
func (o *DcimPowerOutletTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlet templates update o k response
func (o *DcimPowerOutletTemplatesUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletTemplatesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletTemplatesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesUpdateOK %s", 200, payload)
}

func (o *DcimPowerOutletTemplatesUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesUpdateDefault creates a DcimPowerOutletTemplatesUpdateDefault with default headers values
func NewDcimPowerOutletTemplatesUpdateDefault(code int) *DcimPowerOutletTemplatesUpdateDefault {
	return &DcimPowerOutletTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesUpdateDefault dcim power outlet templates update default
*/
type DcimPowerOutletTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlet templates update default response has a 2xx status code
func (o *DcimPowerOutletTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates update default response has a 3xx status code
func (o *DcimPowerOutletTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates update default response has a 4xx status code
func (o *DcimPowerOutletTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates update default response has a 5xx status code
func (o *DcimPowerOutletTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates update default response a status code equal to that given
func (o *DcimPowerOutletTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlet templates update default response
func (o *DcimPowerOutletTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletTemplatesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_update default %s", o._statusCode, payload)
}

func (o *DcimPowerOutletTemplatesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_update default %s", o._statusCode, payload)
}

func (o *DcimPowerOutletTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
