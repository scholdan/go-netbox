// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamL2vpnsDeleteReader is a Reader for the IpamL2vpnsDelete structure.
type IpamL2vpnsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamL2vpnsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ipam/l2vpns/{id}/] ipam_l2vpns_delete", response, response.Code())
	}
}

// NewIpamL2vpnsDeleteNoContent creates a IpamL2vpnsDeleteNoContent with default headers values
func NewIpamL2vpnsDeleteNoContent() *IpamL2vpnsDeleteNoContent {
	return &IpamL2vpnsDeleteNoContent{}
}

/*
IpamL2vpnsDeleteNoContent describes a response with status code 204, with default header values.

IpamL2vpnsDeleteNoContent ipam l2vpns delete no content
*/
type IpamL2vpnsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam l2vpns delete no content response has a 2xx status code
func (o *IpamL2vpnsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpns delete no content response has a 3xx status code
func (o *IpamL2vpnsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns delete no content response has a 4xx status code
func (o *IpamL2vpnsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpns delete no content response has a 5xx status code
func (o *IpamL2vpnsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns delete no content response a status code equal to that given
func (o *IpamL2vpnsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the ipam l2vpns delete no content response
func (o *IpamL2vpnsDeleteNoContent) Code() int {
	return 204
}

func (o *IpamL2vpnsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/{id}/][%d] ipamL2vpnsDeleteNoContent", 204)
}

func (o *IpamL2vpnsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/{id}/][%d] ipamL2vpnsDeleteNoContent", 204)
}

func (o *IpamL2vpnsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}