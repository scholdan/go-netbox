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

// IpamL2vpnTerminationsUpdateReader is a Reader for the IpamL2vpnTerminationsUpdate structure.
type IpamL2vpnTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ipam/l2vpn-terminations/{id}/] ipam_l2vpn-terminations_update", response, response.Code())
	}
}

// NewIpamL2vpnTerminationsUpdateOK creates a IpamL2vpnTerminationsUpdateOK with default headers values
func NewIpamL2vpnTerminationsUpdateOK() *IpamL2vpnTerminationsUpdateOK {
	return &IpamL2vpnTerminationsUpdateOK{}
}

/*
IpamL2vpnTerminationsUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsUpdateOK ipam l2vpn terminations update o k
*/
type IpamL2vpnTerminationsUpdateOK struct {
	Payload *models.L2VPNTermination
}

// IsSuccess returns true when this ipam l2vpn terminations update o k response has a 2xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations update o k response has a 3xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations update o k response has a 4xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations update o k response has a 5xx status code
func (o *IpamL2vpnTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations update o k response a status code equal to that given
func (o *IpamL2vpnTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam l2vpn terminations update o k response
func (o *IpamL2vpnTerminationsUpdateOK) Code() int {
	return 200
}

func (o *IpamL2vpnTerminationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsUpdateOK %s", 200, payload)
}

func (o *IpamL2vpnTerminationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ipam/l2vpn-terminations/{id}/][%d] ipamL2vpnTerminationsUpdateOK %s", 200, payload)
}

func (o *IpamL2vpnTerminationsUpdateOK) GetPayload() *models.L2VPNTermination {
	return o.Payload
}

func (o *IpamL2vpnTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPNTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
