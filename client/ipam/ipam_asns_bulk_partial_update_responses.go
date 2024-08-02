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

// IpamAsnsBulkPartialUpdateReader is a Reader for the IpamAsnsBulkPartialUpdate structure.
type IpamAsnsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/asns/] ipam_asns_bulk_partial_update", response, response.Code())
	}
}

// NewIpamAsnsBulkPartialUpdateOK creates a IpamAsnsBulkPartialUpdateOK with default headers values
func NewIpamAsnsBulkPartialUpdateOK() *IpamAsnsBulkPartialUpdateOK {
	return &IpamAsnsBulkPartialUpdateOK{}
}

/*
IpamAsnsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamAsnsBulkPartialUpdateOK ipam asns bulk partial update o k
*/
type IpamAsnsBulkPartialUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns bulk partial update o k response has a 2xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns bulk partial update o k response has a 3xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns bulk partial update o k response has a 4xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns bulk partial update o k response has a 5xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns bulk partial update o k response a status code equal to that given
func (o *IpamAsnsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam asns bulk partial update o k response
func (o *IpamAsnsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamAsnsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamAsnsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamAsnsBulkPartialUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}