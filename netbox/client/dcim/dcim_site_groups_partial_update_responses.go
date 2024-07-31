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

// DcimSiteGroupsPartialUpdateReader is a Reader for the DcimSiteGroupsPartialUpdate structure.
type DcimSiteGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSiteGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsPartialUpdateOK creates a DcimSiteGroupsPartialUpdateOK with default headers values
func NewDcimSiteGroupsPartialUpdateOK() *DcimSiteGroupsPartialUpdateOK {
	return &DcimSiteGroupsPartialUpdateOK{}
}

/*
DcimSiteGroupsPartialUpdateOK describes a response with status code 200, with default header values.

DcimSiteGroupsPartialUpdateOK dcim site groups partial update o k
*/
type DcimSiteGroupsPartialUpdateOK struct {
	Payload *models.SiteGroup
}

// IsSuccess returns true when this dcim site groups partial update o k response has a 2xx status code
func (o *DcimSiteGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim site groups partial update o k response has a 3xx status code
func (o *DcimSiteGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim site groups partial update o k response has a 4xx status code
func (o *DcimSiteGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim site groups partial update o k response has a 5xx status code
func (o *DcimSiteGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim site groups partial update o k response a status code equal to that given
func (o *DcimSiteGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim site groups partial update o k response
func (o *DcimSiteGroupsPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimSiteGroupsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcimSiteGroupsPartialUpdateOK %s", 200, payload)
}

func (o *DcimSiteGroupsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcimSiteGroupsPartialUpdateOK %s", 200, payload)
}

func (o *DcimSiteGroupsPartialUpdateOK) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsPartialUpdateDefault creates a DcimSiteGroupsPartialUpdateDefault with default headers values
func NewDcimSiteGroupsPartialUpdateDefault(code int) *DcimSiteGroupsPartialUpdateDefault {
	return &DcimSiteGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimSiteGroupsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsPartialUpdateDefault dcim site groups partial update default
*/
type DcimSiteGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim site groups partial update default response has a 2xx status code
func (o *DcimSiteGroupsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim site groups partial update default response has a 3xx status code
func (o *DcimSiteGroupsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim site groups partial update default response has a 4xx status code
func (o *DcimSiteGroupsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim site groups partial update default response has a 5xx status code
func (o *DcimSiteGroupsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim site groups partial update default response a status code equal to that given
func (o *DcimSiteGroupsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim site groups partial update default response
func (o *DcimSiteGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcim_site-groups_partial_update default %s", o._statusCode, payload)
}

func (o *DcimSiteGroupsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcim_site-groups_partial_update default %s", o._statusCode, payload)
}

func (o *DcimSiteGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
