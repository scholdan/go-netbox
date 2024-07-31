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

// IpamFhrpGroupsCreateReader is a Reader for the IpamFhrpGroupsCreate structure.
type IpamFhrpGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamFhrpGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/fhrp-groups/] ipam_fhrp-groups_create", response, response.Code())
	}
}

// NewIpamFhrpGroupsCreateCreated creates a IpamFhrpGroupsCreateCreated with default headers values
func NewIpamFhrpGroupsCreateCreated() *IpamFhrpGroupsCreateCreated {
	return &IpamFhrpGroupsCreateCreated{}
}

/*
IpamFhrpGroupsCreateCreated describes a response with status code 201, with default header values.

IpamFhrpGroupsCreateCreated ipam fhrp groups create created
*/
type IpamFhrpGroupsCreateCreated struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups create created response has a 2xx status code
func (o *IpamFhrpGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups create created response has a 3xx status code
func (o *IpamFhrpGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups create created response has a 4xx status code
func (o *IpamFhrpGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups create created response has a 5xx status code
func (o *IpamFhrpGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups create created response a status code equal to that given
func (o *IpamFhrpGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam fhrp groups create created response
func (o *IpamFhrpGroupsCreateCreated) Code() int {
	return 201
}

func (o *IpamFhrpGroupsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipamFhrpGroupsCreateCreated %s", 201, payload)
}

func (o *IpamFhrpGroupsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipamFhrpGroupsCreateCreated %s", 201, payload)
}

func (o *IpamFhrpGroupsCreateCreated) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
