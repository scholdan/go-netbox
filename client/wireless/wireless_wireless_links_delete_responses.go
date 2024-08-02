// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLinksDeleteReader is a Reader for the WirelessWirelessLinksDelete structure.
type WirelessWirelessLinksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLinksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLinksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /wireless/wireless-links/{id}/] wireless_wireless-links_delete", response, response.Code())
	}
}

// NewWirelessWirelessLinksDeleteNoContent creates a WirelessWirelessLinksDeleteNoContent with default headers values
func NewWirelessWirelessLinksDeleteNoContent() *WirelessWirelessLinksDeleteNoContent {
	return &WirelessWirelessLinksDeleteNoContent{}
}

/*
WirelessWirelessLinksDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLinksDeleteNoContent wireless wireless links delete no content
*/
type WirelessWirelessLinksDeleteNoContent struct {
}

// IsSuccess returns true when this wireless wireless links delete no content response has a 2xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless links delete no content response has a 3xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless links delete no content response has a 4xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless links delete no content response has a 5xx status code
func (o *WirelessWirelessLinksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless links delete no content response a status code equal to that given
func (o *WirelessWirelessLinksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the wireless wireless links delete no content response
func (o *WirelessWirelessLinksDeleteNoContent) Code() int {
	return 204
}

func (o *WirelessWirelessLinksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksDeleteNoContent", 204)
}

func (o *WirelessWirelessLinksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-links/{id}/][%d] wirelessWirelessLinksDeleteNoContent", 204)
}

func (o *WirelessWirelessLinksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}