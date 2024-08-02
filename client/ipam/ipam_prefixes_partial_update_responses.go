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

// IpamPrefixesPartialUpdateReader is a Reader for the IpamPrefixesPartialUpdate structure.
type IpamPrefixesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/prefixes/{id}/] ipam_prefixes_partial_update", response, response.Code())
	}
}

// NewIpamPrefixesPartialUpdateOK creates a IpamPrefixesPartialUpdateOK with default headers values
func NewIpamPrefixesPartialUpdateOK() *IpamPrefixesPartialUpdateOK {
	return &IpamPrefixesPartialUpdateOK{}
}

/*
IpamPrefixesPartialUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesPartialUpdateOK ipam prefixes partial update o k
*/
type IpamPrefixesPartialUpdateOK struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes partial update o k response has a 2xx status code
func (o *IpamPrefixesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes partial update o k response has a 3xx status code
func (o *IpamPrefixesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes partial update o k response has a 4xx status code
func (o *IpamPrefixesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes partial update o k response has a 5xx status code
func (o *IpamPrefixesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes partial update o k response a status code equal to that given
func (o *IpamPrefixesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam prefixes partial update o k response
func (o *IpamPrefixesPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamPrefixesPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesPartialUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}