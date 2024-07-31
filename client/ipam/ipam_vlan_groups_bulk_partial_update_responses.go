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

// IpamVlanGroupsBulkPartialUpdateReader is a Reader for the IpamVlanGroupsBulkPartialUpdate structure.
type IpamVlanGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/vlan-groups/] ipam_vlan-groups_bulk_partial_update", response, response.Code())
	}
}

// NewIpamVlanGroupsBulkPartialUpdateOK creates a IpamVlanGroupsBulkPartialUpdateOK with default headers values
func NewIpamVlanGroupsBulkPartialUpdateOK() *IpamVlanGroupsBulkPartialUpdateOK {
	return &IpamVlanGroupsBulkPartialUpdateOK{}
}

/*
IpamVlanGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkPartialUpdateOK ipam vlan groups bulk partial update o k
*/
type IpamVlanGroupsBulkPartialUpdateOK struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups bulk partial update o k response has a 2xx status code
func (o *IpamVlanGroupsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups bulk partial update o k response has a 3xx status code
func (o *IpamVlanGroupsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups bulk partial update o k response has a 4xx status code
func (o *IpamVlanGroupsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups bulk partial update o k response has a 5xx status code
func (o *IpamVlanGroupsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups bulk partial update o k response a status code equal to that given
func (o *IpamVlanGroupsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam vlan groups bulk partial update o k response
func (o *IpamVlanGroupsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/][%d] ipamVlanGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/][%d] ipamVlanGroupsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
