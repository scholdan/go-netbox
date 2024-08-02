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

// IpamAggregatesReadReader is a Reader for the IpamAggregatesRead structure.
type IpamAggregatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAggregatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ipam/aggregates/{id}/] ipam_aggregates_read", response, response.Code())
	}
}

// NewIpamAggregatesReadOK creates a IpamAggregatesReadOK with default headers values
func NewIpamAggregatesReadOK() *IpamAggregatesReadOK {
	return &IpamAggregatesReadOK{}
}

/*
IpamAggregatesReadOK describes a response with status code 200, with default header values.

IpamAggregatesReadOK ipam aggregates read o k
*/
type IpamAggregatesReadOK struct {
	Payload *models.Aggregate
}

// IsSuccess returns true when this ipam aggregates read o k response has a 2xx status code
func (o *IpamAggregatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates read o k response has a 3xx status code
func (o *IpamAggregatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates read o k response has a 4xx status code
func (o *IpamAggregatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates read o k response has a 5xx status code
func (o *IpamAggregatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates read o k response a status code equal to that given
func (o *IpamAggregatesReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam aggregates read o k response
func (o *IpamAggregatesReadOK) Code() int {
	return 200
}

func (o *IpamAggregatesReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipamAggregatesReadOK %s", 200, payload)
}

func (o *IpamAggregatesReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ipam/aggregates/{id}/][%d] ipamAggregatesReadOK %s", 200, payload)
}

func (o *IpamAggregatesReadOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *IpamAggregatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
