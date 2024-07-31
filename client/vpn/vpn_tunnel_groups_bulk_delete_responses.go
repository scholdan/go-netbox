// Code generated by go-swagger; DO NOT EDIT.

package vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VpnTunnelGroupsBulkDeleteReader is a Reader for the VpnTunnelGroupsBulkDelete structure.
type VpnTunnelGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VpnTunnelGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVpnTunnelGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /vpn/tunnel-groups/] vpn_tunnel-groups_bulk_delete", response, response.Code())
	}
}

// NewVpnTunnelGroupsBulkDeleteNoContent creates a VpnTunnelGroupsBulkDeleteNoContent with default headers values
func NewVpnTunnelGroupsBulkDeleteNoContent() *VpnTunnelGroupsBulkDeleteNoContent {
	return &VpnTunnelGroupsBulkDeleteNoContent{}
}

/*
VpnTunnelGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

VpnTunnelGroupsBulkDeleteNoContent vpn tunnel groups bulk delete no content
*/
type VpnTunnelGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this vpn tunnel groups bulk delete no content response has a 2xx status code
func (o *VpnTunnelGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vpn tunnel groups bulk delete no content response has a 3xx status code
func (o *VpnTunnelGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vpn tunnel groups bulk delete no content response has a 4xx status code
func (o *VpnTunnelGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this vpn tunnel groups bulk delete no content response has a 5xx status code
func (o *VpnTunnelGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this vpn tunnel groups bulk delete no content response a status code equal to that given
func (o *VpnTunnelGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the vpn tunnel groups bulk delete no content response
func (o *VpnTunnelGroupsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *VpnTunnelGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpn/tunnel-groups/][%d] vpnTunnelGroupsBulkDeleteNoContent", 204)
}

func (o *VpnTunnelGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /vpn/tunnel-groups/][%d] vpnTunnelGroupsBulkDeleteNoContent", 204)
}

func (o *VpnTunnelGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
