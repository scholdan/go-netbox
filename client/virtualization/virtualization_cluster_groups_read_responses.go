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

	"github.com/fbreckle/go-netbox/models"
)

// VirtualizationClusterGroupsReadReader is a Reader for the VirtualizationClusterGroupsRead structure.
type VirtualizationClusterGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /virtualization/cluster-groups/{id}/] virtualization_cluster-groups_read", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsReadOK creates a VirtualizationClusterGroupsReadOK with default headers values
func NewVirtualizationClusterGroupsReadOK() *VirtualizationClusterGroupsReadOK {
	return &VirtualizationClusterGroupsReadOK{}
}

/*
VirtualizationClusterGroupsReadOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsReadOK virtualization cluster groups read o k
*/
type VirtualizationClusterGroupsReadOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups read o k response has a 2xx status code
func (o *VirtualizationClusterGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups read o k response has a 3xx status code
func (o *VirtualizationClusterGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups read o k response has a 4xx status code
func (o *VirtualizationClusterGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups read o k response has a 5xx status code
func (o *VirtualizationClusterGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups read o k response a status code equal to that given
func (o *VirtualizationClusterGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster groups read o k response
func (o *VirtualizationClusterGroupsReadOK) Code() int {
	return 200
}

func (o *VirtualizationClusterGroupsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsReadOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
