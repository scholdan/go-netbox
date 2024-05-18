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

// DcimInventoryItemsCreateReader is a Reader for the DcimInventoryItemsCreate structure.
type DcimInventoryItemsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInventoryItemsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsCreateCreated creates a DcimInventoryItemsCreateCreated with default headers values
func NewDcimInventoryItemsCreateCreated() *DcimInventoryItemsCreateCreated {
	return &DcimInventoryItemsCreateCreated{}
}

/*
DcimInventoryItemsCreateCreated describes a response with status code 201, with default header values.

DcimInventoryItemsCreateCreated dcim inventory items create created
*/
type DcimInventoryItemsCreateCreated struct {
	Payload *models.InventoryItem
}

// IsSuccess returns true when this dcim inventory items create created response has a 2xx status code
func (o *DcimInventoryItemsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items create created response has a 3xx status code
func (o *DcimInventoryItemsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items create created response has a 4xx status code
func (o *DcimInventoryItemsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items create created response has a 5xx status code
func (o *DcimInventoryItemsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items create created response a status code equal to that given
func (o *DcimInventoryItemsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim inventory items create created response
func (o *DcimInventoryItemsCreateCreated) Code() int {
	return 201
}

func (o *DcimInventoryItemsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated %s", 201, payload)
}

func (o *DcimInventoryItemsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated %s", 201, payload)
}

func (o *DcimInventoryItemsCreateCreated) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsCreateDefault creates a DcimInventoryItemsCreateDefault with default headers values
func NewDcimInventoryItemsCreateDefault(code int) *DcimInventoryItemsCreateDefault {
	return &DcimInventoryItemsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsCreateDefault describes a response with status code -1, with default header values.

DcimInventoryItemsCreateDefault dcim inventory items create default
*/
type DcimInventoryItemsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim inventory items create default response has a 2xx status code
func (o *DcimInventoryItemsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items create default response has a 3xx status code
func (o *DcimInventoryItemsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items create default response has a 4xx status code
func (o *DcimInventoryItemsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items create default response has a 5xx status code
func (o *DcimInventoryItemsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items create default response a status code equal to that given
func (o *DcimInventoryItemsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim inventory items create default response
func (o *DcimInventoryItemsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcim_inventory-items_create default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcim_inventory-items_create default %s", o._statusCode, payload)
}

func (o *DcimInventoryItemsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
