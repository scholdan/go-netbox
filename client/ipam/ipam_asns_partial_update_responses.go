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

// IpamAsnsPartialUpdateReader is a Reader for the IpamAsnsPartialUpdate structure.
type IpamAsnsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/asns/{id}/] ipam_asns_partial_update", response, response.Code())
	}
}

// NewIpamAsnsPartialUpdateOK creates a IpamAsnsPartialUpdateOK with default headers values
func NewIpamAsnsPartialUpdateOK() *IpamAsnsPartialUpdateOK {
	return &IpamAsnsPartialUpdateOK{}
}

/*
IpamAsnsPartialUpdateOK describes a response with status code 200, with default header values.

IpamAsnsPartialUpdateOK ipam asns partial update o k
*/
type IpamAsnsPartialUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns partial update o k response has a 2xx status code
func (o *IpamAsnsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns partial update o k response has a 3xx status code
func (o *IpamAsnsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns partial update o k response has a 4xx status code
func (o *IpamAsnsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns partial update o k response has a 5xx status code
func (o *IpamAsnsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns partial update o k response a status code equal to that given
func (o *IpamAsnsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam asns partial update o k response
func (o *IpamAsnsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamAsnsPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipamAsnsPartialUpdateOK %s", 200, payload)
}

func (o *IpamAsnsPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/asns/{id}/][%d] ipamAsnsPartialUpdateOK %s", 200, payload)
}

func (o *IpamAsnsPartialUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
