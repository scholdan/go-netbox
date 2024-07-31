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

// VpnTunnelsCreateReader is a Reader for the VpnTunnelsCreate structure.
type VpnTunnelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVpnTunnelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /vpn/tunnels/] vpn_tunnels_create", response, response.Code())
	}
}

// NewVpnTunnelsCreateCreated creates a VpnTunnelsCreateCreated with default headers values
func NewVpnTunnelsCreateCreated() *VpnTunnelsCreateCreated {
	return &VpnTunnelsCreateCreated{}
}

/*
VpnTunnelsCreateCreated describes a response with status code 201, with default header values.

VpnTunnelsCreateCreated vpn tunnels create created
*/
type VpnTunnelsCreateCreated struct {
	Payload *models.Tunnel
}

// IsSuccess returns true when this vpn tunnels create created response has a 2xx status code
func (o *VpnTunnelsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnels create created response has a 3xx status code
func (o *VpnTunnelsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnels create created response has a 4xx status code
func (o *VpnTunnelsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnels create created response has a 5xx status code
func (o *VpnTunnelsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnels create created response a status code equal to that given
func (o *VpnTunnelsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the vpn tunnels create created response
func (o *VpnTunnelsCreateCreated) Code() int {
	return 201
}

func (o *VpnTunnelsCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpnTunnelsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelsCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /vpn/tunnels/][%d] vpnTunnelsCreateCreated %s", 201, payload)
}

func (o *VpnTunnelsCreateCreated) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *VpnTunnelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
