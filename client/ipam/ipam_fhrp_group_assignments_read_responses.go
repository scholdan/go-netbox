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

// IpamFhrpGroupAssignmentsReadReader is a Reader for the IpamFhrpGroupAssignmentsRead structure.
type IpamFhrpGroupAssignmentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/fhrp-group-assignments/{id}/] ipam_fhrp-group-assignments_read", response, response.Code())
	}
}

// NewIpamFhrpGroupAssignmentsReadOK creates a IpamFhrpGroupAssignmentsReadOK with default headers values
func NewIpamFhrpGroupAssignmentsReadOK() *IpamFhrpGroupAssignmentsReadOK {
	return &IpamFhrpGroupAssignmentsReadOK{}
}

/*
IpamFhrpGroupAssignmentsReadOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsReadOK ipam fhrp group assignments read o k
*/
type IpamFhrpGroupAssignmentsReadOK struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments read o k response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments read o k response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments read o k response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments read o k response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments read o k response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp group assignments read o k response
func (o *IpamFhrpGroupAssignmentsReadOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupAssignmentsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsReadOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsReadOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsReadOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
