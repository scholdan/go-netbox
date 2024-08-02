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

// VirtualizationClustersCreateReader is a Reader for the VirtualizationClustersCreate structure.
type VirtualizationClustersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClustersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /virtualization/clusters/] virtualization_clusters_create", response, response.Code())
	}
}

// NewVirtualizationClustersCreateCreated creates a VirtualizationClustersCreateCreated with default headers values
func NewVirtualizationClustersCreateCreated() *VirtualizationClustersCreateCreated {
	return &VirtualizationClustersCreateCreated{}
}

/*
VirtualizationClustersCreateCreated describes a response with status code 201, with default header values.

VirtualizationClustersCreateCreated virtualization clusters create created
*/
type VirtualizationClustersCreateCreated struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters create created response has a 2xx status code
func (o *VirtualizationClustersCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters create created response has a 3xx status code
func (o *VirtualizationClustersCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters create created response has a 4xx status code
func (o *VirtualizationClustersCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters create created response has a 5xx status code
func (o *VirtualizationClustersCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters create created response a status code equal to that given
func (o *VirtualizationClustersCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the virtualization clusters create created response
func (o *VirtualizationClustersCreateCreated) Code() int {
	return 201
}

func (o *VirtualizationClustersCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /virtualization/clusters/][%d] virtualizationClustersCreateCreated %s", 201, payload)
}

func (o *VirtualizationClustersCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /virtualization/clusters/][%d] virtualizationClustersCreateCreated %s", 201, payload)
}

func (o *VirtualizationClustersCreateCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}