// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVrfsDeleteReader is a Reader for the IpamVrfsDelete structure.
type IpamVrfsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ipam/vrfs/{id}/] ipam_vrfs_delete", response, response.Code())
	}
}

// NewIpamVrfsDeleteNoContent creates a IpamVrfsDeleteNoContent with default headers values
func NewIpamVrfsDeleteNoContent() *IpamVrfsDeleteNoContent {
	return &IpamVrfsDeleteNoContent{}
}

/*
IpamVrfsDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsDeleteNoContent ipam vrfs delete no content
*/
type IpamVrfsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam vrfs delete no content response has a 2xx status code
func (o *IpamVrfsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs delete no content response has a 3xx status code
func (o *IpamVrfsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs delete no content response has a 4xx status code
func (o *IpamVrfsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs delete no content response has a 5xx status code
func (o *IpamVrfsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs delete no content response a status code equal to that given
func (o *IpamVrfsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the ipam vrfs delete no content response
func (o *IpamVrfsDeleteNoContent) Code() int {
	return 204
}

func (o *IpamVrfsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/{id}/][%d] ipamVrfsDeleteNoContent", 204)
}

func (o *IpamVrfsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/{id}/][%d] ipamVrfsDeleteNoContent", 204)
}

func (o *IpamVrfsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}