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

// DcimModuleTypesUpdateReader is a Reader for the DcimModuleTypesUpdate structure.
type DcimModuleTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesUpdateOK creates a DcimModuleTypesUpdateOK with default headers values
func NewDcimModuleTypesUpdateOK() *DcimModuleTypesUpdateOK {
	return &DcimModuleTypesUpdateOK{}
}

/*
DcimModuleTypesUpdateOK describes a response with status code 200, with default header values.

DcimModuleTypesUpdateOK dcim module types update o k
*/
type DcimModuleTypesUpdateOK struct {
	Payload *models.ModuleType
}

// IsSuccess returns true when this dcim module types update o k response has a 2xx status code
func (o *DcimModuleTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types update o k response has a 3xx status code
func (o *DcimModuleTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types update o k response has a 4xx status code
func (o *DcimModuleTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types update o k response has a 5xx status code
func (o *DcimModuleTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types update o k response a status code equal to that given
func (o *DcimModuleTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim module types update o k response
func (o *DcimModuleTypesUpdateOK) Code() int {
	return 200
}

func (o *DcimModuleTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-types/{id}/][%d] dcimModuleTypesUpdateOK %s", 200, payload)
}

func (o *DcimModuleTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-types/{id}/][%d] dcimModuleTypesUpdateOK %s", 200, payload)
}

func (o *DcimModuleTypesUpdateOK) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesUpdateDefault creates a DcimModuleTypesUpdateDefault with default headers values
func NewDcimModuleTypesUpdateDefault(code int) *DcimModuleTypesUpdateDefault {
	return &DcimModuleTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesUpdateDefault describes a response with status code -1, with default header values.

DcimModuleTypesUpdateDefault dcim module types update default
*/
type DcimModuleTypesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim module types update default response has a 2xx status code
func (o *DcimModuleTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types update default response has a 3xx status code
func (o *DcimModuleTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types update default response has a 4xx status code
func (o *DcimModuleTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types update default response has a 5xx status code
func (o *DcimModuleTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types update default response a status code equal to that given
func (o *DcimModuleTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim module types update default response
func (o *DcimModuleTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-types/{id}/][%d] dcim_module-types_update default %s", o._statusCode, payload)
}

func (o *DcimModuleTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /dcim/module-types/{id}/][%d] dcim_module-types_update default %s", o._statusCode, payload)
}

func (o *DcimModuleTypesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
