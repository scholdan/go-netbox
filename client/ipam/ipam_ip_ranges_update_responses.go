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

// IpamIPRangesUpdateReader is a Reader for the IpamIPRangesUpdate structure.
type IpamIPRangesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/ip-ranges/{id}/] ipam_ip-ranges_update", response, response.Code())
	}
}

// NewIpamIPRangesUpdateOK creates a IpamIPRangesUpdateOK with default headers values
func NewIpamIPRangesUpdateOK() *IpamIPRangesUpdateOK {
	return &IpamIPRangesUpdateOK{}
}

/*
IpamIPRangesUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesUpdateOK ipam Ip ranges update o k
*/
type IpamIPRangesUpdateOK struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges update o k response has a 2xx status code
func (o *IpamIPRangesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges update o k response has a 3xx status code
func (o *IpamIPRangesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges update o k response has a 4xx status code
func (o *IpamIPRangesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges update o k response has a 5xx status code
func (o *IpamIPRangesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges update o k response a status code equal to that given
func (o *IpamIPRangesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip ranges update o k response
func (o *IpamIPRangesUpdateOK) Code() int {
	return 200
}

func (o *IpamIPRangesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/ip-ranges/{id}/][%d] ipamIpRangesUpdateOK %s", 200, payload)
}

func (o *IpamIPRangesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/ip-ranges/{id}/][%d] ipamIpRangesUpdateOK %s", 200, payload)
}

func (o *IpamIPRangesUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
