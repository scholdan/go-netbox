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

// VirtualizationClusterGroupsBulkUpdateReader is a Reader for the VirtualizationClusterGroupsBulkUpdate structure.
type VirtualizationClusterGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /virtualization/cluster-groups/] virtualization_cluster-groups_bulk_update", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsBulkUpdateOK creates a VirtualizationClusterGroupsBulkUpdateOK with default headers values
func NewVirtualizationClusterGroupsBulkUpdateOK() *VirtualizationClusterGroupsBulkUpdateOK {
	return &VirtualizationClusterGroupsBulkUpdateOK{}
}

/*
VirtualizationClusterGroupsBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsBulkUpdateOK virtualization cluster groups bulk update o k
*/
type VirtualizationClusterGroupsBulkUpdateOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups bulk update o k response has a 2xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups bulk update o k response has a 3xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups bulk update o k response has a 4xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups bulk update o k response has a 5xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups bulk update o k response a status code equal to that given
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization cluster groups bulk update o k response
func (o *VirtualizationClusterGroupsBulkUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK %s", 200, payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
