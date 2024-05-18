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
)

// VirtualizationClusterGroupsDeleteReader is a Reader for the VirtualizationClusterGroupsDelete structure.
type VirtualizationClusterGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClusterGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsDeleteNoContent creates a VirtualizationClusterGroupsDeleteNoContent with default headers values
func NewVirtualizationClusterGroupsDeleteNoContent() *VirtualizationClusterGroupsDeleteNoContent {
	return &VirtualizationClusterGroupsDeleteNoContent{}
}

/*
VirtualizationClusterGroupsDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClusterGroupsDeleteNoContent virtualization cluster groups delete no content
*/
type VirtualizationClusterGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization cluster groups delete no content response has a 2xx status code
func (o *VirtualizationClusterGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups delete no content response has a 3xx status code
func (o *VirtualizationClusterGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups delete no content response has a 4xx status code
func (o *VirtualizationClusterGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups delete no content response has a 5xx status code
func (o *VirtualizationClusterGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups delete no content response a status code equal to that given
func (o *VirtualizationClusterGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the virtualization cluster groups delete no content response
func (o *VirtualizationClusterGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *VirtualizationClusterGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsDeleteNoContent", 204)
}

func (o *VirtualizationClusterGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsDeleteNoContent", 204)
}

func (o *VirtualizationClusterGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationClusterGroupsDeleteDefault creates a VirtualizationClusterGroupsDeleteDefault with default headers values
func NewVirtualizationClusterGroupsDeleteDefault(code int) *VirtualizationClusterGroupsDeleteDefault {
	return &VirtualizationClusterGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClusterGroupsDeleteDefault describes a response with status code -1, with default header values.

VirtualizationClusterGroupsDeleteDefault virtualization cluster groups delete default
*/
type VirtualizationClusterGroupsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups delete default response has a 2xx status code
func (o *VirtualizationClusterGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization cluster groups delete default response has a 3xx status code
func (o *VirtualizationClusterGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization cluster groups delete default response has a 4xx status code
func (o *VirtualizationClusterGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization cluster groups delete default response has a 5xx status code
func (o *VirtualizationClusterGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization cluster groups delete default response a status code equal to that given
func (o *VirtualizationClusterGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization cluster groups delete default response
func (o *VirtualizationClusterGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_delete default %s", o._statusCode, payload)
}

func (o *VirtualizationClusterGroupsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_delete default %s", o._statusCode, payload)
}

func (o *VirtualizationClusterGroupsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
