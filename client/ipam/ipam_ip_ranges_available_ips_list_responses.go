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

// IpamIPRangesAvailableIpsListReader is a Reader for the IpamIPRangesAvailableIpsList structure.
type IpamIPRangesAvailableIpsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesAvailableIpsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesAvailableIpsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/ip-ranges/{id}/available-ips/] ipam_ip-ranges_available-ips_list", response, response.Code())
	}
}

// NewIpamIPRangesAvailableIpsListOK creates a IpamIPRangesAvailableIpsListOK with default headers values
func NewIpamIPRangesAvailableIpsListOK() *IpamIPRangesAvailableIpsListOK {
	return &IpamIPRangesAvailableIpsListOK{}
}

/*
IpamIPRangesAvailableIpsListOK describes a response with status code 200, with default header values.

IpamIPRangesAvailableIpsListOK ipam Ip ranges available ips list o k
*/
type IpamIPRangesAvailableIpsListOK struct {
	Payload []*models.AvailableIP
}

// IsSuccess returns true when this ipam Ip ranges available ips list o k response has a 2xx status code
func (o *IpamIPRangesAvailableIpsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges available ips list o k response has a 3xx status code
func (o *IpamIPRangesAvailableIpsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges available ips list o k response has a 4xx status code
func (o *IpamIPRangesAvailableIpsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges available ips list o k response has a 5xx status code
func (o *IpamIPRangesAvailableIpsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges available ips list o k response a status code equal to that given
func (o *IpamIPRangesAvailableIpsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip ranges available ips list o k response
func (o *IpamIPRangesAvailableIpsListOK) Code() int {
	return 200
}

func (o *IpamIPRangesAvailableIpsListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/available-ips/][%d] ipamIpRangesAvailableIpsListOK %s", 200, payload)
}

func (o *IpamIPRangesAvailableIpsListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/ip-ranges/{id}/available-ips/][%d] ipamIpRangesAvailableIpsListOK %s", 200, payload)
}

func (o *IpamIPRangesAvailableIpsListOK) GetPayload() []*models.AvailableIP {
	return o.Payload
}

func (o *IpamIPRangesAvailableIpsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
