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

// DcimInventoryItemsPartialUpdateReader is a Reader for the DcimInventoryItemsPartialUpdate structure.
type DcimInventoryItemsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsPartialUpdateOK creates a DcimInventoryItemsPartialUpdateOK with default headers values
func NewDcimInventoryItemsPartialUpdateOK() *DcimInventoryItemsPartialUpdateOK {
	return &DcimInventoryItemsPartialUpdateOK{}
}

/*
DcimInventoryItemsPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemsPartialUpdateOK dcim inventory items partial update o k
*/
type DcimInventoryItemsPartialUpdateOK struct {
	Payload *models.InventoryItem
}

// IsSuccess returns true when this dcim inventory items partial update o k response has a 2xx status code
func (o *DcimInventoryItemsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items partial update o k response has a 3xx status code
func (o *DcimInventoryItemsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items partial update o k response has a 4xx status code
func (o *DcimInventoryItemsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items partial update o k response has a 5xx status code
func (o *DcimInventoryItemsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items partial update o k response a status code equal to that given
func (o *DcimInventoryItemsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim inventory items partial update o k response
func (o *DcimInventoryItemsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimInventoryItemsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcimInventoryItemsPartialUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcimInventoryItemsPartialUpdateOK %s", 200, payload)
}

func (o *DcimInventoryItemsPartialUpdateOK) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsPartialUpdateDefault creates a DcimInventoryItemsPartialUpdateDefault with default headers values
func NewDcimInventoryItemsPartialUpdateDefault(code int) *DcimInventoryItemsPartialUpdateDefault {
	return &DcimInventoryItemsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemsPartialUpdateDefault dcim inventory items partial update default
*/
type DcimInventoryItemsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim inventory items partial update default response has a 2xx status code
func (o *DcimInventoryItemsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items partial update default response has a 3xx status code
func (o *DcimInventoryItemsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items partial update default response has a 4xx status code
func (o *DcimInventoryItemsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items partial update default response has a 5xx status code
func (o *DcimInventoryItemsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items partial update default response a status code equal to that given
func (o *DcimInventoryItemsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim inventory items partial update default response
func (o *DcimInventoryItemsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcim_inventory-items_partial_update default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcim_inventory-items_partial_update default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
