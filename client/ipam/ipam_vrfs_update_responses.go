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

// IpamVrfsUpdateReader is a Reader for the IpamVrfsUpdate structure.
type IpamVrfsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/vrfs/{id}/] ipam_vrfs_update", response, response.Code())
	}
}

// NewIpamVrfsUpdateOK creates a IpamVrfsUpdateOK with default headers values
func NewIpamVrfsUpdateOK() *IpamVrfsUpdateOK {
	return &IpamVrfsUpdateOK{}
}

/*
IpamVrfsUpdateOK describes a response with status code 200, with default header values.

IpamVrfsUpdateOK ipam vrfs update o k
*/
type IpamVrfsUpdateOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs update o k response has a 2xx status code
func (o *IpamVrfsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs update o k response has a 3xx status code
func (o *IpamVrfsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs update o k response has a 4xx status code
func (o *IpamVrfsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs update o k response has a 5xx status code
func (o *IpamVrfsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs update o k response a status code equal to that given
func (o *IpamVrfsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vrfs update o k response
func (o *IpamVrfsUpdateOK) Code() int {
	return 200
}

func (o *IpamVrfsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipamVrfsUpdateOK %s", 200, payload)
}

func (o *IpamVrfsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipamVrfsUpdateOK %s", 200, payload)
}

func (o *IpamVrfsUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
