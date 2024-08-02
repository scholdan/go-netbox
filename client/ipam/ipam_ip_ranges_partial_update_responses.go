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

// IpamIPRangesPartialUpdateReader is a Reader for the IpamIPRangesPartialUpdate structure.
type IpamIPRangesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/ip-ranges/{id}/] ipam_ip-ranges_partial_update", response, response.Code())
	}
}

// NewIpamIPRangesPartialUpdateOK creates a IpamIPRangesPartialUpdateOK with default headers values
func NewIpamIPRangesPartialUpdateOK() *IpamIPRangesPartialUpdateOK {
	return &IpamIPRangesPartialUpdateOK{}
}

/*
IpamIPRangesPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesPartialUpdateOK ipam Ip ranges partial update o k
*/
type IpamIPRangesPartialUpdateOK struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges partial update o k response has a 2xx status code
func (o *IpamIPRangesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges partial update o k response has a 3xx status code
func (o *IpamIPRangesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges partial update o k response has a 4xx status code
func (o *IpamIPRangesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges partial update o k response has a 5xx status code
func (o *IpamIPRangesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges partial update o k response a status code equal to that given
func (o *IpamIPRangesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam Ip ranges partial update o k response
func (o *IpamIPRangesPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamIPRangesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/ip-ranges/{id}/][%d] ipamIpRangesPartialUpdateOK %s", 200, payload)
}

func (o *IpamIPRangesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/ip-ranges/{id}/][%d] ipamIpRangesPartialUpdateOK %s", 200, payload)
}

func (o *IpamIPRangesPartialUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
