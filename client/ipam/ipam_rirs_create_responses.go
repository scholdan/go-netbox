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

// IpamRirsCreateReader is a Reader for the IpamRirsCreate structure.
type IpamRirsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRirsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/rirs/] ipam_rirs_create", response, response.Code())
	}
}

// NewIpamRirsCreateCreated creates a IpamRirsCreateCreated with default headers values
func NewIpamRirsCreateCreated() *IpamRirsCreateCreated {
	return &IpamRirsCreateCreated{}
}

/*
IpamRirsCreateCreated describes a response with status code 201, with default header values.

IpamRirsCreateCreated ipam rirs create created
*/
type IpamRirsCreateCreated struct {
	Payload *models.RIR
}

// IsSuccess returns true when this ipam rirs create created response has a 2xx status code
func (o *IpamRirsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs create created response has a 3xx status code
func (o *IpamRirsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs create created response has a 4xx status code
func (o *IpamRirsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs create created response has a 5xx status code
func (o *IpamRirsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs create created response a status code equal to that given
func (o *IpamRirsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam rirs create created response
func (o *IpamRirsCreateCreated) Code() int {
	return 201
}

func (o *IpamRirsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/rirs/][%d] ipamRirsCreateCreated %s", 201, payload)
}

func (o *IpamRirsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/rirs/][%d] ipamRirsCreateCreated %s", 201, payload)
}

func (o *IpamRirsCreateCreated) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}