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

// IpamIPAddressesBulkPartialUpdateReader is a Reader for the IpamIPAddressesBulkPartialUpdate structure.
type IpamIPAddressesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/ip-addresses/] ipam_ip-addresses_bulk_partial_update", response, response.Code())
	}
}

// NewIpamIPAddressesBulkPartialUpdateOK creates a IpamIPAddressesBulkPartialUpdateOK with default headers values
func NewIpamIPAddressesBulkPartialUpdateOK() *IpamIPAddressesBulkPartialUpdateOK {
	return &IpamIPAddressesBulkPartialUpdateOK{}
}

/*
IpamIPAddressesBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkPartialUpdateOK ipam Ip addresses bulk partial update o k
*/
type IpamIPAddressesBulkPartialUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses bulk partial update o k response has a 2xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses bulk partial update o k response has a 3xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses bulk partial update o k response has a 4xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses bulk partial update o k response has a 5xx status code
func (o *IpamIPAddressesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses bulk partial update o k response a status code equal to that given
func (o *IpamIPAddressesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip addresses bulk partial update o k response
func (o *IpamIPAddressesBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamIPAddressesBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipamIpAddressesBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamIPAddressesBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/][%d] ipamIpAddressesBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamIPAddressesBulkPartialUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
