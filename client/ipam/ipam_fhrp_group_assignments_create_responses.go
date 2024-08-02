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

// IpamFhrpGroupAssignmentsCreateReader is a Reader for the IpamFhrpGroupAssignmentsCreate structure.
type IpamFhrpGroupAssignmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamFhrpGroupAssignmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/fhrp-group-assignments/] ipam_fhrp-group-assignments_create", response, response.Code())
	}
}

// NewIpamFhrpGroupAssignmentsCreateCreated creates a IpamFhrpGroupAssignmentsCreateCreated with default headers values
func NewIpamFhrpGroupAssignmentsCreateCreated() *IpamFhrpGroupAssignmentsCreateCreated {
	return &IpamFhrpGroupAssignmentsCreateCreated{}
}

/*
IpamFhrpGroupAssignmentsCreateCreated describes a response with status code 201, with default header values.

IpamFhrpGroupAssignmentsCreateCreated ipam fhrp group assignments create created
*/
type IpamFhrpGroupAssignmentsCreateCreated struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments create created response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments create created response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments create created response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments create created response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments create created response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam fhrp group assignments create created response
func (o *IpamFhrpGroupAssignmentsCreateCreated) Code() int {
	return 201
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated %s", 201, payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated %s", 201, payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}