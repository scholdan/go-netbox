// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlansDeleteReader is a Reader for the IpamVlansDelete structure.
type IpamVlansDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlansDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ipam/vlans/{id}/] ipam_vlans_delete", response, response.Code())
	}
}

// NewIpamVlansDeleteNoContent creates a IpamVlansDeleteNoContent with default headers values
func NewIpamVlansDeleteNoContent() *IpamVlansDeleteNoContent {
	return &IpamVlansDeleteNoContent{}
}

/*
IpamVlansDeleteNoContent describes a response with status code 204, with default header values.

IpamVlansDeleteNoContent ipam vlans delete no content
*/
type IpamVlansDeleteNoContent struct {
}

// IsSuccess returns true when this ipam vlans delete no content response has a 2xx status code
func (o *IpamVlansDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlans delete no content response has a 3xx status code
func (o *IpamVlansDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlans delete no content response has a 4xx status code
func (o *IpamVlansDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlans delete no content response has a 5xx status code
func (o *IpamVlansDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlans delete no content response a status code equal to that given
func (o *IpamVlansDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the ipam vlans delete no content response
func (o *IpamVlansDeleteNoContent) Code() int {
	return 204
}

func (o *IpamVlansDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/{id}/][%d] ipamVlansDeleteNoContent", 204)
}

func (o *IpamVlansDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/{id}/][%d] ipamVlansDeleteNoContent", 204)
}

func (o *IpamVlansDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}