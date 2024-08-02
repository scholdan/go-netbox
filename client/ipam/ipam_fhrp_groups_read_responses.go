// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// IpamFhrpGroupsReadReader is a Reader for the IpamFhrpGroupsRead structure.
type IpamFhrpGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/fhrp-groups/{id}/] ipam_fhrp-groups_read", response, response.Code())
	}
}

// NewIpamFhrpGroupsReadOK creates a IpamFhrpGroupsReadOK with default headers values
func NewIpamFhrpGroupsReadOK() *IpamFhrpGroupsReadOK {
	return &IpamFhrpGroupsReadOK{}
}

/*
IpamFhrpGroupsReadOK describes a response with status code 200, with default header values.

IpamFhrpGroupsReadOK ipam fhrp groups read o k
*/
type IpamFhrpGroupsReadOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups read o k response has a 2xx status code
func (o *IpamFhrpGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups read o k response has a 3xx status code
func (o *IpamFhrpGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups read o k response has a 4xx status code
func (o *IpamFhrpGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups read o k response has a 5xx status code
func (o *IpamFhrpGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups read o k response a status code equal to that given
func (o *IpamFhrpGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp groups read o k response
func (o *IpamFhrpGroupsReadOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsReadOK %s", 200, payload)
}

func (o *IpamFhrpGroupsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsReadOK %s", 200, payload)
}

func (o *IpamFhrpGroupsReadOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
