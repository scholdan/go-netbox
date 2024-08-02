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

// IpamIPRangesCreateReader is a Reader for the IpamIPRangesCreate structure.
type IpamIPRangesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamIPRangesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ipam/ip-ranges/] ipam_ip-ranges_create", response, response.Code())
	}
}

// NewIpamIPRangesCreateCreated creates a IpamIPRangesCreateCreated with default headers values
func NewIpamIPRangesCreateCreated() *IpamIPRangesCreateCreated {
	return &IpamIPRangesCreateCreated{}
}

/*
IpamIPRangesCreateCreated describes a response with status code 201, with default header values.

IpamIPRangesCreateCreated ipam Ip ranges create created
*/
type IpamIPRangesCreateCreated struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges create created response has a 2xx status code
func (o *IpamIPRangesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges create created response has a 3xx status code
func (o *IpamIPRangesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges create created response has a 4xx status code
func (o *IpamIPRangesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges create created response has a 5xx status code
func (o *IpamIPRangesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges create created response a status code equal to that given
func (o *IpamIPRangesCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the ipam Ip ranges create created response
func (o *IpamIPRangesCreateCreated) Code() int {
	return 201
}

func (o *IpamIPRangesCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated %s", 201, payload)
}

func (o *IpamIPRangesCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated %s", 201, payload)
}

func (o *IpamIPRangesCreateCreated) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}