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

// DcimCablesReadReader is a Reader for the DcimCablesRead structure.
type DcimCablesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesReadOK creates a DcimCablesReadOK with default headers values
func NewDcimCablesReadOK() *DcimCablesReadOK {
	return &DcimCablesReadOK{}
}

/* DcimCablesReadOK describes a response with status code 200, with default header values.

DcimCablesReadOK dcim cables read o k
*/
type DcimCablesReadOK struct {
	Payload *models.Cable
}

func (o *DcimCablesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/{id}/][%d] dcimCablesReadOK  %+v", 200, o.Payload)
}
func (o *DcimCablesReadOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesReadDefault creates a DcimCablesReadDefault with default headers values
func NewDcimCablesReadDefault(code int) *DcimCablesReadDefault {
	return &DcimCablesReadDefault{
		_statusCode: code,
	}
}

/* DcimCablesReadDefault describes a response with status code -1, with default header values.

DcimCablesReadDefault dcim cables read default
*/
type DcimCablesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables read default response
func (o *DcimCablesReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimCablesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/{id}/][%d] dcim_cables_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimCablesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
