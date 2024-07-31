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

// IpamFhrpGroupsPartialUpdateReader is a Reader for the IpamFhrpGroupsPartialUpdate structure.
type IpamFhrpGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/fhrp-groups/{id}/] ipam_fhrp-groups_partial_update", response, response.Code())
	}
}

// NewIpamFhrpGroupsPartialUpdateOK creates a IpamFhrpGroupsPartialUpdateOK with default headers values
func NewIpamFhrpGroupsPartialUpdateOK() *IpamFhrpGroupsPartialUpdateOK {
	return &IpamFhrpGroupsPartialUpdateOK{}
}

/*
IpamFhrpGroupsPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupsPartialUpdateOK ipam fhrp groups partial update o k
*/
type IpamFhrpGroupsPartialUpdateOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups partial update o k response has a 2xx status code
func (o *IpamFhrpGroupsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups partial update o k response has a 3xx status code
func (o *IpamFhrpGroupsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups partial update o k response has a 4xx status code
func (o *IpamFhrpGroupsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups partial update o k response has a 5xx status code
func (o *IpamFhrpGroupsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups partial update o k response a status code equal to that given
func (o *IpamFhrpGroupsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp groups partial update o k response
func (o *IpamFhrpGroupsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupsPartialUpdateOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
