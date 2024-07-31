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

// DcimSiteGroupsCreateReader is a Reader for the DcimSiteGroupsCreate structure.
type DcimSiteGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimSiteGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsCreateCreated creates a DcimSiteGroupsCreateCreated with default headers values
func NewDcimSiteGroupsCreateCreated() *DcimSiteGroupsCreateCreated {
	return &DcimSiteGroupsCreateCreated{}
}

/*
DcimSiteGroupsCreateCreated describes a response with status code 201, with default header values.

DcimSiteGroupsCreateCreated dcim site groups create created
*/
type DcimSiteGroupsCreateCreated struct {
	Payload *models.SiteGroup
}

// IsSuccess returns true when this dcim site groups create created response has a 2xx status code
func (o *DcimSiteGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim site groups create created response has a 3xx status code
func (o *DcimSiteGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim site groups create created response has a 4xx status code
func (o *DcimSiteGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim site groups create created response has a 5xx status code
func (o *DcimSiteGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim site groups create created response a status code equal to that given
func (o *DcimSiteGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim site groups create created response
func (o *DcimSiteGroupsCreateCreated) Code() int {
	return 201
}

func (o *DcimSiteGroupsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/site-groups/][%d] dcimSiteGroupsCreateCreated %s", 201, payload)
}

func (o *DcimSiteGroupsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/site-groups/][%d] dcimSiteGroupsCreateCreated %s", 201, payload)
}

func (o *DcimSiteGroupsCreateCreated) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsCreateDefault creates a DcimSiteGroupsCreateDefault with default headers values
func NewDcimSiteGroupsCreateDefault(code int) *DcimSiteGroupsCreateDefault {
	return &DcimSiteGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimSiteGroupsCreateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsCreateDefault dcim site groups create default
*/
type DcimSiteGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim site groups create default response has a 2xx status code
func (o *DcimSiteGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim site groups create default response has a 3xx status code
func (o *DcimSiteGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim site groups create default response has a 4xx status code
func (o *DcimSiteGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim site groups create default response has a 5xx status code
func (o *DcimSiteGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim site groups create default response a status code equal to that given
func (o *DcimSiteGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim site groups create default response
func (o *DcimSiteGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/site-groups/][%d] dcim_site-groups_create default %s", o._statusCode, payload)
}

func (o *DcimSiteGroupsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /dcim/site-groups/][%d] dcim_site-groups_create default %s", o._statusCode, payload)
}

func (o *DcimSiteGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
