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

	"github.com/fbreckle/go-netbox/models"
)

// VpnTunnelGroupsReadReader is a Reader for the VpnTunnelGroupsRead structure.
type VpnTunnelGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVpnTunnelGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /vpn/tunnel-groups/{id}/] vpn_tunnel-groups_read", response, response.Code())
	}
}

// NewVpnTunnelGroupsReadOK creates a VpnTunnelGroupsReadOK with default headers values
func NewVpnTunnelGroupsReadOK() *VpnTunnelGroupsReadOK {
	return &VpnTunnelGroupsReadOK{}
}

/*
VpnTunnelGroupsReadOK describes a response with status code 200, with default header values.

VpnTunnelGroupsReadOK vpn tunnel groups read o k
*/
type VpnTunnelGroupsReadOK struct {
	Payload *models.TunnelGroup
}

// IsSuccess returns true when this vpn tunnel groups read o k response has a 2xx status code
func (o *VpnTunnelGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel groups read o k response has a 3xx status code
func (o *VpnTunnelGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel groups read o k response has a 4xx status code
func (o *VpnTunnelGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel groups read o k response has a 5xx status code
func (o *VpnTunnelGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel groups read o k response a status code equal to that given
func (o *VpnTunnelGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vpn tunnel groups read o k response
func (o *VpnTunnelGroupsReadOK) Code() int {
	return 200
}

func (o *VpnTunnelGroupsReadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vpn/tunnel-groups/{id}/][%d] vpnTunnelGroupsReadOK %s", 200, payload)
}

func (o *VpnTunnelGroupsReadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /vpn/tunnel-groups/{id}/][%d] vpnTunnelGroupsReadOK %s", 200, payload)
}

func (o *VpnTunnelGroupsReadOK) GetPayload() *models.TunnelGroup {
	return o.Payload
}

func (o *VpnTunnelGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TunnelGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
