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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimModuleBayTemplatesCreateReader is a Reader for the DcimModuleBayTemplatesCreate structure.
type DcimModuleBayTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimModuleBayTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesCreateCreated creates a DcimModuleBayTemplatesCreateCreated with default headers values
func NewDcimModuleBayTemplatesCreateCreated() *DcimModuleBayTemplatesCreateCreated {
	return &DcimModuleBayTemplatesCreateCreated{}
}

/* DcimModuleBayTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimModuleBayTemplatesCreateCreated dcim module bay templates create created
*/
type DcimModuleBayTemplatesCreateCreated struct {
	Payload *models.ModuleBayTemplate
}

func (o *DcimModuleBayTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/module-bay-templates/][%d] dcimModuleBayTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimModuleBayTemplatesCreateCreated) GetPayload() *models.ModuleBayTemplate {
	return o.Payload
}

func (o *DcimModuleBayTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBayTemplatesCreateDefault creates a DcimModuleBayTemplatesCreateDefault with default headers values
func NewDcimModuleBayTemplatesCreateDefault(code int) *DcimModuleBayTemplatesCreateDefault {
	return &DcimModuleBayTemplatesCreateDefault{
		_statusCode: code,
	}
}

/* DcimModuleBayTemplatesCreateDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesCreateDefault dcim module bay templates create default
*/
type DcimModuleBayTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates create default response
func (o *DcimModuleBayTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleBayTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/module-bay-templates/][%d] dcim_module-bay-templates_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimModuleBayTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}