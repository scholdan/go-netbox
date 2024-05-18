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

// VirtualizationClustersReadReader is a Reader for the VirtualizationClustersRead structure.
type VirtualizationClustersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersReadOK creates a VirtualizationClustersReadOK with default headers values
func NewVirtualizationClustersReadOK() *VirtualizationClustersReadOK {
	return &VirtualizationClustersReadOK{}
}

/*
VirtualizationClustersReadOK describes a response with status code 200, with default header values.

VirtualizationClustersReadOK virtualization clusters read o k
*/
type VirtualizationClustersReadOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters read o k response has a 2xx status code
func (o *VirtualizationClustersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters read o k response has a 3xx status code
func (o *VirtualizationClustersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters read o k response has a 4xx status code
func (o *VirtualizationClustersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters read o k response has a 5xx status code
func (o *VirtualizationClustersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters read o k response a status code equal to that given
func (o *VirtualizationClustersReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization clusters read o k response
func (o *VirtualizationClustersReadOK) Code() int {
	return 200
}

func (o *VirtualizationClustersReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadOK %s", 200, payload)
}

func (o *VirtualizationClustersReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadOK %s", 200, payload)
}

func (o *VirtualizationClustersReadOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersReadDefault creates a VirtualizationClustersReadDefault with default headers values
func NewVirtualizationClustersReadDefault(code int) *VirtualizationClustersReadDefault {
	return &VirtualizationClustersReadDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersReadDefault describes a response with status code -1, with default header values.

VirtualizationClustersReadDefault virtualization clusters read default
*/
type VirtualizationClustersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization clusters read default response has a 2xx status code
func (o *VirtualizationClustersReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters read default response has a 3xx status code
func (o *VirtualizationClustersReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters read default response has a 4xx status code
func (o *VirtualizationClustersReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters read default response has a 5xx status code
func (o *VirtualizationClustersReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters read default response a status code equal to that given
func (o *VirtualizationClustersReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization clusters read default response
func (o *VirtualizationClustersReadDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClustersReadDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualization_clusters_read default %s", o._statusCode, payload)
}

func (o *VirtualizationClustersReadDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualization_clusters_read default %s", o._statusCode, payload)
}

func (o *VirtualizationClustersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
