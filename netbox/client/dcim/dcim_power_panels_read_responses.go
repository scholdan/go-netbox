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

// DcimPowerPanelsReadReader is a Reader for the DcimPowerPanelsRead structure.
type DcimPowerPanelsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsReadOK creates a DcimPowerPanelsReadOK with default headers values
func NewDcimPowerPanelsReadOK() *DcimPowerPanelsReadOK {
	return &DcimPowerPanelsReadOK{}
}

/*
DcimPowerPanelsReadOK describes a response with status code 200, with default header values.

DcimPowerPanelsReadOK dcim power panels read o k
*/
type DcimPowerPanelsReadOK struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels read o k response has a 2xx status code
func (o *DcimPowerPanelsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels read o k response has a 3xx status code
func (o *DcimPowerPanelsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels read o k response has a 4xx status code
func (o *DcimPowerPanelsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels read o k response has a 5xx status code
func (o *DcimPowerPanelsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels read o k response a status code equal to that given
func (o *DcimPowerPanelsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power panels read o k response
func (o *DcimPowerPanelsReadOK) Code() int {
	return 200
}

func (o *DcimPowerPanelsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcimPowerPanelsReadOK %s", 200, payload)
}

func (o *DcimPowerPanelsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcimPowerPanelsReadOK %s", 200, payload)
}

func (o *DcimPowerPanelsReadOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsReadDefault creates a DcimPowerPanelsReadDefault with default headers values
func NewDcimPowerPanelsReadDefault(code int) *DcimPowerPanelsReadDefault {
	return &DcimPowerPanelsReadDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsReadDefault describes a response with status code -1, with default header values.

DcimPowerPanelsReadDefault dcim power panels read default
*/
type DcimPowerPanelsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power panels read default response has a 2xx status code
func (o *DcimPowerPanelsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels read default response has a 3xx status code
func (o *DcimPowerPanelsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels read default response has a 4xx status code
func (o *DcimPowerPanelsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels read default response has a 5xx status code
func (o *DcimPowerPanelsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels read default response a status code equal to that given
func (o *DcimPowerPanelsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power panels read default response
func (o *DcimPowerPanelsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPanelsReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcim_power-panels_read default %s", o._statusCode, payload)
}

func (o *DcimPowerPanelsReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/power-panels/{id}/][%d] dcim_power-panels_read default %s", o._statusCode, payload)
}

func (o *DcimPowerPanelsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
