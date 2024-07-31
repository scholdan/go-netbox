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

// IpamPrefixesBulkUpdateReader is a Reader for the IpamPrefixesBulkUpdate structure.
type IpamPrefixesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/prefixes/] ipam_prefixes_bulk_update", response, response.Code())
	}
}

// NewIpamPrefixesBulkUpdateOK creates a IpamPrefixesBulkUpdateOK with default headers values
func NewIpamPrefixesBulkUpdateOK() *IpamPrefixesBulkUpdateOK {
	return &IpamPrefixesBulkUpdateOK{}
}

/*
IpamPrefixesBulkUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesBulkUpdateOK ipam prefixes bulk update o k
*/
type IpamPrefixesBulkUpdateOK struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes bulk update o k response has a 2xx status code
func (o *IpamPrefixesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes bulk update o k response has a 3xx status code
func (o *IpamPrefixesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes bulk update o k response has a 4xx status code
func (o *IpamPrefixesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes bulk update o k response has a 5xx status code
func (o *IpamPrefixesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes bulk update o k response a status code equal to that given
func (o *IpamPrefixesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam prefixes bulk update o k response
func (o *IpamPrefixesBulkUpdateOK) Code() int {
	return 200
}

func (o *IpamPrefixesBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/prefixes/][%d] ipamPrefixesBulkUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/prefixes/][%d] ipamPrefixesBulkUpdateOK %s", 200, payload)
}

func (o *IpamPrefixesBulkUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
