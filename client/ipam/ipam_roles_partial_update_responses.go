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

// IpamRolesPartialUpdateReader is a Reader for the IpamRolesPartialUpdate structure.
type IpamRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/roles/{id}/] ipam_roles_partial_update", response, response.Code())
	}
}

// NewIpamRolesPartialUpdateOK creates a IpamRolesPartialUpdateOK with default headers values
func NewIpamRolesPartialUpdateOK() *IpamRolesPartialUpdateOK {
	return &IpamRolesPartialUpdateOK{}
}

/*
IpamRolesPartialUpdateOK describes a response with status code 200, with default header values.

IpamRolesPartialUpdateOK ipam roles partial update o k
*/
type IpamRolesPartialUpdateOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles partial update o k response has a 2xx status code
func (o *IpamRolesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles partial update o k response has a 3xx status code
func (o *IpamRolesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles partial update o k response has a 4xx status code
func (o *IpamRolesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles partial update o k response has a 5xx status code
func (o *IpamRolesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles partial update o k response a status code equal to that given
func (o *IpamRolesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam roles partial update o k response
func (o *IpamRolesPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRolesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipamRolesPartialUpdateOK %s", 200, payload)
}

func (o *IpamRolesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/roles/{id}/][%d] ipamRolesPartialUpdateOK %s", 200, payload)
}

func (o *IpamRolesPartialUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
