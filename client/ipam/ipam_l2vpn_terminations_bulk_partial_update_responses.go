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

// IpamL2vpnTerminationsBulkPartialUpdateReader is a Reader for the IpamL2vpnTerminationsBulkPartialUpdate structure.
type IpamL2vpnTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /ipam/l2vpn-terminations/] ipam_l2vpn-terminations_bulk_partial_update", response, response.Code())
	}
}

// NewIpamL2vpnTerminationsBulkPartialUpdateOK creates a IpamL2vpnTerminationsBulkPartialUpdateOK with default headers values
func NewIpamL2vpnTerminationsBulkPartialUpdateOK() *IpamL2vpnTerminationsBulkPartialUpdateOK {
	return &IpamL2vpnTerminationsBulkPartialUpdateOK{}
}

/*
IpamL2vpnTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnTerminationsBulkPartialUpdateOK ipam l2vpn terminations bulk partial update o k
*/
type IpamL2vpnTerminationsBulkPartialUpdateOK struct {
	Payload *models.L2VPNTermination
}

// IsSuccess returns true when this ipam l2vpn terminations bulk partial update o k response has a 2xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpn terminations bulk partial update o k response has a 3xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpn terminations bulk partial update o k response has a 4xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpn terminations bulk partial update o k response has a 5xx status code
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpn terminations bulk partial update o k response a status code equal to that given
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam l2vpn terminations bulk partial update o k response
func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /ipam/l2vpn-terminations/][%d] ipamL2vpnTerminationsBulkPartialUpdateOK %s", 200, payload)
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) GetPayload() *models.L2VPNTermination {
	return o.Payload
}

func (o *IpamL2vpnTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPNTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
