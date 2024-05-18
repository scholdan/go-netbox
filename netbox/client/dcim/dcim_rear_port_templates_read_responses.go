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

// DcimRearPortTemplatesReadReader is a Reader for the DcimRearPortTemplatesRead structure.
type DcimRearPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesReadOK creates a DcimRearPortTemplatesReadOK with default headers values
func NewDcimRearPortTemplatesReadOK() *DcimRearPortTemplatesReadOK {
	return &DcimRearPortTemplatesReadOK{}
}

/*
DcimRearPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesReadOK dcim rear port templates read o k
*/
type DcimRearPortTemplatesReadOK struct {
	Payload *models.RearPortTemplate
}

// IsSuccess returns true when this dcim rear port templates read o k response has a 2xx status code
func (o *DcimRearPortTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear port templates read o k response has a 3xx status code
func (o *DcimRearPortTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear port templates read o k response has a 4xx status code
func (o *DcimRearPortTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear port templates read o k response has a 5xx status code
func (o *DcimRearPortTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear port templates read o k response a status code equal to that given
func (o *DcimRearPortTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim rear port templates read o k response
func (o *DcimRearPortTemplatesReadOK) Code() int {
	return 200
}

func (o *DcimRearPortTemplatesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesReadOK %s", 200, payload)
}

func (o *DcimRearPortTemplatesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesReadOK %s", 200, payload)
}

func (o *DcimRearPortTemplatesReadOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesReadDefault creates a DcimRearPortTemplatesReadDefault with default headers values
func NewDcimRearPortTemplatesReadDefault(code int) *DcimRearPortTemplatesReadDefault {
	return &DcimRearPortTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortTemplatesReadDefault describes a response with status code -1, with default header values.

DcimRearPortTemplatesReadDefault dcim rear port templates read default
*/
type DcimRearPortTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rear port templates read default response has a 2xx status code
func (o *DcimRearPortTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear port templates read default response has a 3xx status code
func (o *DcimRearPortTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear port templates read default response has a 4xx status code
func (o *DcimRearPortTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear port templates read default response has a 5xx status code
func (o *DcimRearPortTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear port templates read default response a status code equal to that given
func (o *DcimRearPortTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rear port templates read default response
func (o *DcimRearPortTemplatesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_read default %s", o._statusCode, payload)
}

func (o *DcimRearPortTemplatesReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_read default %s", o._statusCode, payload)
}

func (o *DcimRearPortTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
