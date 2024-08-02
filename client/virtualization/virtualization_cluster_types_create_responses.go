// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scholdan/go-netbox/models"
)

// VirtualizationClusterTypesCreateReader is a Reader for the VirtualizationClusterTypesCreate structure.
type VirtualizationClusterTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClusterTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /virtualization/cluster-types/] virtualization_cluster-types_create", response, response.Code())
	}
}

// NewVirtualizationClusterTypesCreateCreated creates a VirtualizationClusterTypesCreateCreated with default headers values
func NewVirtualizationClusterTypesCreateCreated() *VirtualizationClusterTypesCreateCreated {
	return &VirtualizationClusterTypesCreateCreated{}
}

/*
VirtualizationClusterTypesCreateCreated describes a response with status code 201, with default header values.

VirtualizationClusterTypesCreateCreated virtualization cluster types create created
*/
type VirtualizationClusterTypesCreateCreated struct {
	Payload *models.ClusterType
}

// IsSuccess returns true when this virtualization cluster types create created response has a 2xx status code
func (o *VirtualizationClusterTypesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster types create created response has a 3xx status code
func (o *VirtualizationClusterTypesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster types create created response has a 4xx status code
func (o *VirtualizationClusterTypesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster types create created response has a 5xx status code
func (o *VirtualizationClusterTypesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster types create created response a status code equal to that given
func (o *VirtualizationClusterTypesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the virtualization cluster types create created response
func (o *VirtualizationClusterTypesCreateCreated) Code() int {
	return 201
}

func (o *VirtualizationClusterTypesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /virtualization/cluster-types/][%d] virtualizationClusterTypesCreateCreated %s", 201, payload)
}

func (o *VirtualizationClusterTypesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /virtualization/cluster-types/][%d] virtualizationClusterTypesCreateCreated %s", 201, payload)
}

func (o *VirtualizationClusterTypesCreateCreated) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}