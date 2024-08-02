// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationClustersDeleteReader is a Reader for the VirtualizationClustersDelete structure.
type VirtualizationClustersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClustersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /virtualization/clusters/{id}/] virtualization_clusters_delete", response, response.Code())
	}
}

// NewVirtualizationClustersDeleteNoContent creates a VirtualizationClustersDeleteNoContent with default headers values
func NewVirtualizationClustersDeleteNoContent() *VirtualizationClustersDeleteNoContent {
	return &VirtualizationClustersDeleteNoContent{}
}

/*
VirtualizationClustersDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClustersDeleteNoContent virtualization clusters delete no content
*/
type VirtualizationClustersDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization clusters delete no content response has a 2xx status code
func (o *VirtualizationClustersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters delete no content response has a 3xx status code
func (o *VirtualizationClustersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters delete no content response has a 4xx status code
func (o *VirtualizationClustersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters delete no content response has a 5xx status code
func (o *VirtualizationClustersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters delete no content response a status code equal to that given
func (o *VirtualizationClustersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the virtualization clusters delete no content response
func (o *VirtualizationClustersDeleteNoContent) Code() int {
	return 204
}

func (o *VirtualizationClustersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent", 204)
}

func (o *VirtualizationClustersDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent", 204)
}

func (o *VirtualizationClustersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}