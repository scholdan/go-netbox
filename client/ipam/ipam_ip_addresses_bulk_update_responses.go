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

// IpamIPAddressesBulkUpdateReader is a Reader for the IpamIPAddressesBulkUpdate structure.
type IpamIPAddressesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/ip-addresses/] ipam_ip-addresses_bulk_update", response, response.Code())
	}
}

// NewIpamIPAddressesBulkUpdateOK creates a IpamIPAddressesBulkUpdateOK with default headers values
func NewIpamIPAddressesBulkUpdateOK() *IpamIPAddressesBulkUpdateOK {
	return &IpamIPAddressesBulkUpdateOK{}
}

/*
IpamIPAddressesBulkUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesBulkUpdateOK ipam Ip addresses bulk update o k
*/
type IpamIPAddressesBulkUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses bulk update o k response has a 2xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses bulk update o k response has a 3xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses bulk update o k response has a 4xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses bulk update o k response has a 5xx status code
func (o *IpamIPAddressesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses bulk update o k response a status code equal to that given
func (o *IpamIPAddressesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip addresses bulk update o k response
func (o *IpamIPAddressesBulkUpdateOK) Code() int {
	return 200
}

func (o *IpamIPAddressesBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipamIpAddressesBulkUpdateOK %s", 200, payload)
}

func (o *IpamIPAddressesBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/ip-addresses/][%d] ipamIpAddressesBulkUpdateOK %s", 200, payload)
}

func (o *IpamIPAddressesBulkUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}