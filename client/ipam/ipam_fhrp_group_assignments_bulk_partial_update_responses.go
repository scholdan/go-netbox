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

// IpamFhrpGroupAssignmentsBulkPartialUpdateReader is a Reader for the IpamFhrpGroupAssignmentsBulkPartialUpdate structure.
type IpamFhrpGroupAssignmentsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/fhrp-group-assignments/] ipam_fhrp-group-assignments_bulk_partial_update", response, response.Code())
	}
}

// NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK creates a IpamFhrpGroupAssignmentsBulkPartialUpdateOK with default headers values
func NewIpamFhrpGroupAssignmentsBulkPartialUpdateOK() *IpamFhrpGroupAssignmentsBulkPartialUpdateOK {
	return &IpamFhrpGroupAssignmentsBulkPartialUpdateOK{}
}

/*
IpamFhrpGroupAssignmentsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsBulkPartialUpdateOK ipam fhrp group assignments bulk partial update o k
*/
type IpamFhrpGroupAssignmentsBulkPartialUpdateOK struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments bulk partial update o k response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments bulk partial update o k response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments bulk partial update o k response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments bulk partial update o k response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments bulk partial update o k response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam fhrp group assignments bulk partial update o k response
func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
