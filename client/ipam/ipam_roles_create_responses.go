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

// IpamRolesCreateReader is a Reader for the IpamRolesCreate structure.
type IpamRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/roles/] ipam_roles_create", response, response.Code())
	}
}

// NewIpamRolesCreateCreated creates a IpamRolesCreateCreated with default headers values
func NewIpamRolesCreateCreated() *IpamRolesCreateCreated {
	return &IpamRolesCreateCreated{}
}

/*
IpamRolesCreateCreated describes a response with status code 201, with default header values.

IpamRolesCreateCreated ipam roles create created
*/
type IpamRolesCreateCreated struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles create created response has a 2xx status code
func (o *IpamRolesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles create created response has a 3xx status code
func (o *IpamRolesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles create created response has a 4xx status code
func (o *IpamRolesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles create created response has a 5xx status code
func (o *IpamRolesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles create created response a status code equal to that given
func (o *IpamRolesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam roles create created response
func (o *IpamRolesCreateCreated) Code() int {
	return 201
}

func (o *IpamRolesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/roles/][%d] ipamRolesCreateCreated %s", 201, payload)
}

func (o *IpamRolesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/roles/][%d] ipamRolesCreateCreated %s", 201, payload)
}

func (o *IpamRolesCreateCreated) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
