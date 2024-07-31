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

	"github.com/scholdan/go-netbox/netbox/models"
)

// ExtrasExportTemplatesCreateReader is a Reader for the ExtrasExportTemplatesCreate structure.
type ExtrasExportTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasExportTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasExportTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesCreateCreated creates a ExtrasExportTemplatesCreateCreated with default headers values
func NewExtrasExportTemplatesCreateCreated() *ExtrasExportTemplatesCreateCreated {
	return &ExtrasExportTemplatesCreateCreated{}
}

/*
ExtrasExportTemplatesCreateCreated describes a response with status code 201, with default header values.

ExtrasExportTemplatesCreateCreated extras export templates create created
*/
type ExtrasExportTemplatesCreateCreated struct {
	Payload *models.ExportTemplate
}

// IsSuccess returns true when this extras export templates create created response has a 2xx status code
func (o *ExtrasExportTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates create created response has a 3xx status code
func (o *ExtrasExportTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates create created response has a 4xx status code
func (o *ExtrasExportTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates create created response has a 5xx status code
func (o *ExtrasExportTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates create created response a status code equal to that given
func (o *ExtrasExportTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the extras export templates create created response
func (o *ExtrasExportTemplatesCreateCreated) Code() int {
	return 201
}

func (o *ExtrasExportTemplatesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extrasExportTemplatesCreateCreated %s", 201, payload)
}

func (o *ExtrasExportTemplatesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extrasExportTemplatesCreateCreated %s", 201, payload)
}

func (o *ExtrasExportTemplatesCreateCreated) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesCreateDefault creates a ExtrasExportTemplatesCreateDefault with default headers values
func NewExtrasExportTemplatesCreateDefault(code int) *ExtrasExportTemplatesCreateDefault {
	return &ExtrasExportTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesCreateDefault describes a response with status code -1, with default header values.

ExtrasExportTemplatesCreateDefault extras export templates create default
*/
type ExtrasExportTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras export templates create default response has a 2xx status code
func (o *ExtrasExportTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates create default response has a 3xx status code
func (o *ExtrasExportTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates create default response has a 4xx status code
func (o *ExtrasExportTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates create default response has a 5xx status code
func (o *ExtrasExportTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates create default response a status code equal to that given
func (o *ExtrasExportTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras export templates create default response
func (o *ExtrasExportTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasExportTemplatesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extras_export-templates_create default %s", o._statusCode, payload)
}

func (o *ExtrasExportTemplatesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extras_export-templates_create default %s", o._statusCode, payload)
}

func (o *ExtrasExportTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
