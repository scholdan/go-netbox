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

// VirtualizationClustersBulkPartialUpdateReader is a Reader for the VirtualizationClustersBulkPartialUpdate structure.
type VirtualizationClustersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /virtualization/clusters/] virtualization_clusters_bulk_partial_update", response, response.Code())
	}
}

// NewVirtualizationClustersBulkPartialUpdateOK creates a VirtualizationClustersBulkPartialUpdateOK with default headers values
func NewVirtualizationClustersBulkPartialUpdateOK() *VirtualizationClustersBulkPartialUpdateOK {
	return &VirtualizationClustersBulkPartialUpdateOK{}
}

/*
VirtualizationClustersBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersBulkPartialUpdateOK virtualization clusters bulk partial update o k
*/
type VirtualizationClustersBulkPartialUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters bulk partial update o k response has a 2xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters bulk partial update o k response has a 3xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters bulk partial update o k response has a 4xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters bulk partial update o k response has a 5xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters bulk partial update o k response a status code equal to that given
func (o *VirtualizationClustersBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the virtualization clusters bulk partial update o k response
func (o *VirtualizationClustersBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *VirtualizationClustersBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualizationClustersBulkPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationClustersBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualizationClustersBulkPartialUpdateOK %s", 200, payload)
}

func (o *VirtualizationClustersBulkPartialUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}