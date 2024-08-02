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

// IpamPrefixesAvailableIpsListReader is a Reader for the IpamPrefixesAvailableIpsList structure.
type IpamPrefixesAvailableIpsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesAvailableIpsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/prefixes/{id}/available-ips/] ipam_prefixes_available-ips_list", response, response.Code())
	}
}

// NewIpamPrefixesAvailableIpsListOK creates a IpamPrefixesAvailableIpsListOK with default headers values
func NewIpamPrefixesAvailableIpsListOK() *IpamPrefixesAvailableIpsListOK {
	return &IpamPrefixesAvailableIpsListOK{}
}

/*
IpamPrefixesAvailableIpsListOK describes a response with status code 200, with default header values.

IpamPrefixesAvailableIpsListOK ipam prefixes available ips list o k
*/
type IpamPrefixesAvailableIpsListOK struct {
	Payload []*models.AvailableIP
}

// IsSuccess returns true when this ipam prefixes available ips list o k response has a 2xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes available ips list o k response has a 3xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes available ips list o k response has a 4xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes available ips list o k response has a 5xx status code
func (o *IpamPrefixesAvailableIpsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes available ips list o k response a status code equal to that given
func (o *IpamPrefixesAvailableIpsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam prefixes available ips list o k response
func (o *IpamPrefixesAvailableIpsListOK) Code() int {
	return 200
}

func (o *IpamPrefixesAvailableIpsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsListOK %s", 200, payload)
}

func (o *IpamPrefixesAvailableIpsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsListOK %s", 200, payload)
}

func (o *IpamPrefixesAvailableIpsListOK) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}