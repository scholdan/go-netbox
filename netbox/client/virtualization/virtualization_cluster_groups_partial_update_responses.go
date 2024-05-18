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

package virtualization

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

// VirtualizationClusterGroupsPartialUpdateReader is a Reader for the VirtualizationClusterGroupsPartialUpdate structure.
type VirtualizationClusterGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsPartialUpdateOK creates a VirtualizationClusterGroupsPartialUpdateOK with default headers values
func NewVirtualizationClusterGroupsPartialUpdateOK() *VirtualizationClusterGroupsPartialUpdateOK {
	return &VirtualizationClusterGroupsPartialUpdateOK{}
}

/*
VirtualizationClusterGroupsPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsPartialUpdateOK virtualization cluster groups partial update o k
*/
type VirtualizationClusterGroupsPartialUpdateOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups partial update o k response has a 2xx status code
func (o *VirtualizationClusterGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups partial update o k response has a 3xx status code
func (o *VirtualizationClusterGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups partial update o k response has a 4xx status code
func (o *VirtualizationClusterGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups partial update o k response has a 5xx status code
func (o *VirtualizationClusterGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups partial update o k response a status code equal to that given
func (o *VirtualizationClusterGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster groups partial update o k response
func (o *VirtualizationClusterGroupsPartialUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsPartialUpdateDefault creates a VirtualizationClusterGroupsPartialUpdateDefault with default headers values
func NewVirtualizationClusterGroupsPartialUpdateDefault(code int) *VirtualizationClusterGroupsPartialUpdateDefault {
	return &VirtualizationClusterGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterGroupsPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClusterGroupsPartialUpdateDefault virtualization cluster groups partial update default
*/
type VirtualizationClusterGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups partial update default response has a 2xx status code
func (o *VirtualizationClusterGroupsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster groups partial update default response has a 3xx status code
func (o *VirtualizationClusterGroupsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster groups partial update default response has a 4xx status code
func (o *VirtualizationClusterGroupsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster groups partial update default response has a 5xx status code
func (o *VirtualizationClusterGroupsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster groups partial update default response a status code equal to that given
func (o *VirtualizationClusterGroupsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization cluster groups partial update default response
func (o *VirtualizationClusterGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_partial_update default %s", o._statusCode, payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_partial_update default %s", o._statusCode, payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
