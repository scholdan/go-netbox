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

// IpamFhrpGroupAssignmentsPartialUpdateReader is a Reader for the IpamFhrpGroupAssignmentsPartialUpdate structure.
type IpamFhrpGroupAssignmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/fhrp-group-assignments/{id}/] ipam_fhrp-group-assignments_partial_update", response, response.Code())
	}
}

// NewIpamFhrpGroupAssignmentsPartialUpdateOK creates a IpamFhrpGroupAssignmentsPartialUpdateOK with default headers values
func NewIpamFhrpGroupAssignmentsPartialUpdateOK() *IpamFhrpGroupAssignmentsPartialUpdateOK {
	return &IpamFhrpGroupAssignmentsPartialUpdateOK{}
}

/*
IpamFhrpGroupAssignmentsPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsPartialUpdateOK ipam fhrp group assignments partial update o k
*/
type IpamFhrpGroupAssignmentsPartialUpdateOK struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments partial update o k response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments partial update o k response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments partial update o k response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments partial update o k response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments partial update o k response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp group assignments partial update o k response
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
