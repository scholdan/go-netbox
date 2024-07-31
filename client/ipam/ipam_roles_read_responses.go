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

	"github.com/fbreckle/go-netbox/models"
)

// IpamRolesReadReader is a Reader for the IpamRolesRead structure.
type IpamRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/roles/{id}/] ipam_roles_read", response, response.Code())
	}
}

// NewIpamRolesReadOK creates a IpamRolesReadOK with default headers values
func NewIpamRolesReadOK() *IpamRolesReadOK {
	return &IpamRolesReadOK{}
}

/*
IpamRolesReadOK describes a response with status code 200, with default header values.

IpamRolesReadOK ipam roles read o k
*/
type IpamRolesReadOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles read o k response has a 2xx status code
func (o *IpamRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles read o k response has a 3xx status code
func (o *IpamRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles read o k response has a 4xx status code
func (o *IpamRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles read o k response has a 5xx status code
func (o *IpamRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles read o k response a status code equal to that given
func (o *IpamRolesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam roles read o k response
func (o *IpamRolesReadOK) Code() int {
	return 200
}

func (o *IpamRolesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK %s", 200, payload)
}

func (o *IpamRolesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK %s", 200, payload)
}

func (o *IpamRolesReadOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
