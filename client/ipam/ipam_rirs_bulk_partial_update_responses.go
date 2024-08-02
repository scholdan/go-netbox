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

// IpamRirsBulkPartialUpdateReader is a Reader for the IpamRirsBulkPartialUpdate structure.
type IpamRirsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/rirs/] ipam_rirs_bulk_partial_update", response, response.Code())
	}
}

// NewIpamRirsBulkPartialUpdateOK creates a IpamRirsBulkPartialUpdateOK with default headers values
func NewIpamRirsBulkPartialUpdateOK() *IpamRirsBulkPartialUpdateOK {
	return &IpamRirsBulkPartialUpdateOK{}
}

/*
IpamRirsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamRirsBulkPartialUpdateOK ipam rirs bulk partial update o k
*/
type IpamRirsBulkPartialUpdateOK struct {
	Payload *models.RIR
}

// IsSuccess returns true when this ipam rirs bulk partial update o k response has a 2xx status code
func (o *IpamRirsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs bulk partial update o k response has a 3xx status code
func (o *IpamRirsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs bulk partial update o k response has a 4xx status code
func (o *IpamRirsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs bulk partial update o k response has a 5xx status code
func (o *IpamRirsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs bulk partial update o k response a status code equal to that given
func (o *IpamRirsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam rirs bulk partial update o k response
func (o *IpamRirsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRirsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/rirs/][%d] ipamRirsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamRirsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/rirs/][%d] ipamRirsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamRirsBulkPartialUpdateOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}