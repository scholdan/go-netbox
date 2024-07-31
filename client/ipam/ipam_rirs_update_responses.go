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

// IpamRirsUpdateReader is a Reader for the IpamRirsUpdate structure.
type IpamRirsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/rirs/{id}/] ipam_rirs_update", response, response.Code())
	}
}

// NewIpamRirsUpdateOK creates a IpamRirsUpdateOK with default headers values
func NewIpamRirsUpdateOK() *IpamRirsUpdateOK {
	return &IpamRirsUpdateOK{}
}

/*
IpamRirsUpdateOK describes a response with status code 200, with default header values.

IpamRirsUpdateOK ipam rirs update o k
*/
type IpamRirsUpdateOK struct {
	Payload *models.RIR
}

// IsSuccess returns true when this ipam rirs update o k response has a 2xx status code
func (o *IpamRirsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs update o k response has a 3xx status code
func (o *IpamRirsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs update o k response has a 4xx status code
func (o *IpamRirsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs update o k response has a 5xx status code
func (o *IpamRirsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs update o k response a status code equal to that given
func (o *IpamRirsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam rirs update o k response
func (o *IpamRirsUpdateOK) Code() int {
	return 200
}

func (o *IpamRirsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipamRirsUpdateOK %s", 200, payload)
}

func (o *IpamRirsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipamRirsUpdateOK %s", 200, payload)
}

func (o *IpamRirsUpdateOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
