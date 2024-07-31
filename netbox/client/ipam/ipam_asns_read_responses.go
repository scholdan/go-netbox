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

// IpamAsnsReadReader is a Reader for the IpamAsnsRead structure.
type IpamAsnsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/asns/{id}/] ipam_asns_read", response, response.Code())
	}
}

// NewIpamAsnsReadOK creates a IpamAsnsReadOK with default headers values
func NewIpamAsnsReadOK() *IpamAsnsReadOK {
	return &IpamAsnsReadOK{}
}

/*
IpamAsnsReadOK describes a response with status code 200, with default header values.

IpamAsnsReadOK ipam asns read o k
*/
type IpamAsnsReadOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns read o k response has a 2xx status code
func (o *IpamAsnsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns read o k response has a 3xx status code
func (o *IpamAsnsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns read o k response has a 4xx status code
func (o *IpamAsnsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns read o k response has a 5xx status code
func (o *IpamAsnsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns read o k response a status code equal to that given
func (o *IpamAsnsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam asns read o k response
func (o *IpamAsnsReadOK) Code() int {
	return 200
}

func (o *IpamAsnsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/asns/{id}/][%d] ipamAsnsReadOK %s", 200, payload)
}

func (o *IpamAsnsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/asns/{id}/][%d] ipamAsnsReadOK %s", 200, payload)
}

func (o *IpamAsnsReadOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
