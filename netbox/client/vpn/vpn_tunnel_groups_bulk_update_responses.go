// Code generated by go-swagger; DO NOT EDIT.

package vpn

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

// VpnTunnelGroupsBulkUpdateReader is a Reader for the VpnTunnelGroupsBulkUpdate structure.
type VpnTunnelGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVpnTunnelGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /vpn/tunnel-groups/] vpn_tunnel-groups_bulk_update", response, response.Code())
	}
}

// NewVpnTunnelGroupsBulkUpdateOK creates a VpnTunnelGroupsBulkUpdateOK with default headers values
func NewVpnTunnelGroupsBulkUpdateOK() *VpnTunnelGroupsBulkUpdateOK {
	return &VpnTunnelGroupsBulkUpdateOK{}
}

/*
VpnTunnelGroupsBulkUpdateOK describes a response with status code 200, with default header values.

VpnTunnelGroupsBulkUpdateOK vpn tunnel groups bulk update o k
*/
type VpnTunnelGroupsBulkUpdateOK struct {
	Payload *models.TunnelGroup
}

// IsSuccess returns true when this vpn tunnel groups bulk update o k response has a 2xx status code
func (o *VpnTunnelGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel groups bulk update o k response has a 3xx status code
func (o *VpnTunnelGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel groups bulk update o k response has a 4xx status code
func (o *VpnTunnelGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel groups bulk update o k response has a 5xx status code
func (o *VpnTunnelGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel groups bulk update o k response a status code equal to that given
func (o *VpnTunnelGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vpn tunnel groups bulk update o k response
func (o *VpnTunnelGroupsBulkUpdateOK) Code() int {
	return 200
}

func (o *VpnTunnelGroupsBulkUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-groups/][%d] vpnTunnelGroupsBulkUpdateOK %s", 200, payload)
}

func (o *VpnTunnelGroupsBulkUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnel-groups/][%d] vpnTunnelGroupsBulkUpdateOK %s", 200, payload)
}

func (o *VpnTunnelGroupsBulkUpdateOK) GetPayload() *models.TunnelGroup {
	return o.Payload
}

func (o *VpnTunnelGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TunnelGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}