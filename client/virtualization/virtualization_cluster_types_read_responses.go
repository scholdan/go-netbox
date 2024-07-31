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

// VirtualizationClusterTypesReadReader is a Reader for the VirtualizationClusterTypesRead structure.
type VirtualizationClusterTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /virtualization/cluster-types/{id}/] virtualization_cluster-types_read", response, response.Code())
	}
}

// NewVirtualizationClusterTypesReadOK creates a VirtualizationClusterTypesReadOK with default headers values
func NewVirtualizationClusterTypesReadOK() *VirtualizationClusterTypesReadOK {
	return &VirtualizationClusterTypesReadOK{}
}

/*
VirtualizationClusterTypesReadOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesReadOK virtualization cluster types read o k
*/
type VirtualizationClusterTypesReadOK struct {
	Payload *models.ClusterType
}

// IsSuccess returns true when this virtualization cluster types read o k response has a 2xx status code
func (o *VirtualizationClusterTypesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster types read o k response has a 3xx status code
func (o *VirtualizationClusterTypesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster types read o k response has a 4xx status code
func (o *VirtualizationClusterTypesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster types read o k response has a 5xx status code
func (o *VirtualizationClusterTypesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster types read o k response a status code equal to that given
func (o *VirtualizationClusterTypesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster types read o k response
func (o *VirtualizationClusterTypesReadOK) Code() int {
	return 200
}

func (o *VirtualizationClusterTypesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesReadOK %s", 200, payload)
}

func (o *VirtualizationClusterTypesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesReadOK %s", 200, payload)
}

func (o *VirtualizationClusterTypesReadOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
