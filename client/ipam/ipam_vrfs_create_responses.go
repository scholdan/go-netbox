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

// IpamVrfsCreateReader is a Reader for the IpamVrfsCreate structure.
type IpamVrfsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVrfsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/vrfs/] ipam_vrfs_create", response, response.Code())
	}
}

// NewIpamVrfsCreateCreated creates a IpamVrfsCreateCreated with default headers values
func NewIpamVrfsCreateCreated() *IpamVrfsCreateCreated {
	return &IpamVrfsCreateCreated{}
}

/*
IpamVrfsCreateCreated describes a response with status code 201, with default header values.

IpamVrfsCreateCreated ipam vrfs create created
*/
type IpamVrfsCreateCreated struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs create created response has a 2xx status code
func (o *IpamVrfsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs create created response has a 3xx status code
func (o *IpamVrfsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs create created response has a 4xx status code
func (o *IpamVrfsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs create created response has a 5xx status code
func (o *IpamVrfsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs create created response a status code equal to that given
func (o *IpamVrfsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam vrfs create created response
func (o *IpamVrfsCreateCreated) Code() int {
	return 201
}

func (o *IpamVrfsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated %s", 201, payload)
}

func (o *IpamVrfsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/vrfs/][%d] ipamVrfsCreateCreated %s", 201, payload)
}

func (o *IpamVrfsCreateCreated) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
