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

// VpnTunnelsUpdateReader is a Reader for the VpnTunnelsUpdate structure.
type VpnTunnelsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVpnTunnelsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /vpn/tunnels/{id}/] vpn_tunnels_update", response, response.Code())
	}
}

// NewVpnTunnelsUpdateOK creates a VpnTunnelsUpdateOK with default headers values
func NewVpnTunnelsUpdateOK() *VpnTunnelsUpdateOK {
	return &VpnTunnelsUpdateOK{}
}

/*
VpnTunnelsUpdateOK describes a response with status code 200, with default header values.

VpnTunnelsUpdateOK vpn tunnels update o k
*/
type VpnTunnelsUpdateOK struct {
	Payload *models.Tunnel
}

// IsSuccess returns true when this vpn tunnels update o k response has a 2xx status code
func (o *VpnTunnelsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnels update o k response has a 3xx status code
func (o *VpnTunnelsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnels update o k response has a 4xx status code
func (o *VpnTunnelsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnels update o k response has a 5xx status code
func (o *VpnTunnelsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnels update o k response a status code equal to that given
func (o *VpnTunnelsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vpn tunnels update o k response
func (o *VpnTunnelsUpdateOK) Code() int {
	return 200
}

func (o *VpnTunnelsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpnTunnelsUpdateOK %s", 200, payload)
}

func (o *VpnTunnelsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /vpn/tunnels/{id}/][%d] vpnTunnelsUpdateOK %s", 200, payload)
}

func (o *VpnTunnelsUpdateOK) GetPayload() *models.Tunnel {
	return o.Payload
}

func (o *VpnTunnelsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tunnel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
