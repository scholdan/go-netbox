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

// DcimRegionsPartialUpdateReader is a Reader for the DcimRegionsPartialUpdate structure.
type DcimRegionsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsPartialUpdateOK creates a DcimRegionsPartialUpdateOK with default headers values
func NewDcimRegionsPartialUpdateOK() *DcimRegionsPartialUpdateOK {
	return &DcimRegionsPartialUpdateOK{}
}

/* DcimRegionsPartialUpdateOK describes a response with status code 200, with default header values.

DcimRegionsPartialUpdateOK dcim regions partial update o k
*/
type DcimRegionsPartialUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/{id}/][%d] dcimRegionsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsPartialUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsPartialUpdateDefault creates a DcimRegionsPartialUpdateDefault with default headers values
func NewDcimRegionsPartialUpdateDefault(code int) *DcimRegionsPartialUpdateDefault {
	return &DcimRegionsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimRegionsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRegionsPartialUpdateDefault dcim regions partial update default
*/
type DcimRegionsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions partial update default response
func (o *DcimRegionsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRegionsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/{id}/][%d] dcim_regions_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRegionsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
