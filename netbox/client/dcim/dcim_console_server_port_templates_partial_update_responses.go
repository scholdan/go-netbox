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

// DcimConsoleServerPortTemplatesPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesPartialUpdate structure.
type DcimConsoleServerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateOK creates a DcimConsoleServerPortTemplatesPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateOK() *DcimConsoleServerPortTemplatesPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesPartialUpdateOK{}
}

/*
DcimConsoleServerPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesPartialUpdateOK dcim console server port templates partial update o k
*/
type DcimConsoleServerPortTemplatesPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates partial update o k response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates partial update o k response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates partial update o k response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates partial update o k response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates partial update o k response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server port templates partial update o k response
func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK %s", 200, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesPartialUpdateDefault creates a DcimConsoleServerPortTemplatesPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateDefault(code int) *DcimConsoleServerPortTemplatesPartialUpdateDefault {
	return &DcimConsoleServerPortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesPartialUpdateDefault dcim console server port templates partial update default
*/
type DcimConsoleServerPortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console server port templates partial update default response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server port templates partial update default response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server port templates partial update default response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server port templates partial update default response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server port templates partial update default response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console server port templates partial update default response
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_partial_update default %s", o._statusCode, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_partial_update default %s", o._statusCode, payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
